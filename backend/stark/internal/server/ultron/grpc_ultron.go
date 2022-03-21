package ultron

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"gitlab.com/greyhole/myid/pkg/extractor"
	nightkit "gitlab.com/greyhole/night-kit/pkg/api"
	"go.uber.org/zap"

	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/jarvis"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/umo"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/setting"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	ep "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
)

func NewServer(
	service nightkit.Service,
	setting setting.Setting,
	entClient *ent.Client,
	natashaClient natasha.NatashaClient,
	blackWidowClient natasha.BlackWidowClient,
) stark.UltronServer {
	return &ultronServer{
		logger:           service.Logger(),
		stats:            service.Stats(),
		entClient:        entClient,
		jarvisClient:     jarvis.New(service.Logger()),
		natashaClient:    natashaClient,
		blackWidowClient: blackWidowClient,
		umoClient: umo.New(
			service.Logger(),
			setting.GetCryptoConfig().Domain,
			setting.GetCryptoConfig().StoreId,
			setting.GetCryptoConfig().SecretKey,
		),
	}
}

type ultronServer struct {
	stark.UnimplementedUltronServer

	logger           *zap.Logger
	stats            *statsd.Client
	entClient        *ent.Client
	jarvisClient     jarvis.Client
	natashaClient    natasha.NatashaClient
	blackWidowClient natasha.BlackWidowClient
	umoClient        umo.Client
}

const (
	mexUserIDFmt    = "M%s"
	mexUserIDFmt64  = "M%d"
	mexUserIDFmtUMO = "M%d_%d"
	systemUser      = "system"
)

const (
	// 1 USDT = 23000 VND
	exchangeRate = 23000
)

type PaymentTransition struct {
	ExpectedType      stark.PaymentType
	ExpectedMethod    stark.MethodType
	ExpectedStatus    stark.Status
	NextSuccessStatus stark.Status
	NextFailedStatus  stark.Status
}

func (p *PaymentTransition) GetMerchantSuccessNote() string {
	switch p.NextSuccessStatus {
	case stark.Status_CANCELED:
		return "Merchant canceled payment."
	case stark.Status_REJECTED:
		return "Merchant rejected payment."
	case stark.Status_APPROVED:
		return "Merchant approved payment."
	case stark.Status_SUBMITTED:
		return "Merchant submitted payment."
	case stark.Status_COMPLETED:
		return "Merchant completed payment."
	}

	return "Merchant succeeded on handling payment."
}

func convertUSDT2VND(amountInUSDT float64) uint64 {
	return uint64(amountInUSDT * exchangeRate)
}

func getMexIDAudit(mexID string) string {
	return fmt.Sprintf(mexUserIDFmt, mexID)
}

func getMexIDAudit64(mexID int64) string {
	return fmt.Sprintf(mexUserIDFmt64, mexID)
}

func getUserID(ctx context.Context) (int64, string, error) {
	n, err := extractor.New().GetUserID(ctx)
	if err != nil {
		return 0, "", err
	}
	return n, strconv.FormatInt(n, 10), nil
}

func (s *ultronServer) completeCallback(
	request *jarvis.CompletePaymentRequest,
	responseErr error,
	transition PaymentTransition,
) {
	txErr := tx.WithTransaction(context.Background(), s.logger, s.entClient, func(ctx context.Context, tx tx.Tx) error {
		payment, err := tx.Client().Payment.
			Query().
			Where(
				ep.IDEQ(request.PaymentId),
				ep.TypeEQ(int32(transition.ExpectedType)),
				ep.MethodEQ(int32(transition.ExpectedMethod)),
				ep.StatusEQ(int32(transition.ExpectedStatus)),
			).
			ForUpdate().
			First(ctx)
		if err != nil {
			s.logger.Error("get payment error", zap.Error(err))
			return err
		}
		caller := getMexIDAudit64(payment.MerchantID)

		paymentUpdate := tx.Client().Payment.
			UpdateOne(payment).
			SetUpdatedAt(time.Now().UTC())
		revisionCreate := tx.Client().Revision.Create().
			SetCreatedAt(time.Now().UTC()).
			SetCreatedBy(caller).
			SetUpdatedBy(caller).
			SetPaymentID(payment.ID).
			SetPayment(payment)

		if responseErr != nil {
			s.logger.Error("complete payment error", zap.Error(responseErr))
			paymentUpdate.SetStatus(int32(transition.NextFailedStatus))
			revisionCreate.SetStatus(int32(transition.NextFailedStatus))
			revisionCreate.SetNote(responseErr.Error())
		} else {
			paymentUpdate.SetStatus(int32(transition.NextSuccessStatus))
			revisionCreate.SetStatus(int32(transition.NextSuccessStatus))
			revisionCreate.SetNote(transition.GetMerchantSuccessNote())
		}

		_, err = paymentUpdate.Save(ctx)
		if err != nil {
			s.logger.Error("could not update status payment", zap.Error(err))
			return err
		}
		_, err = revisionCreate.Save(ctx)
		if err != nil {
			s.logger.Error("create revision err", zap.Error(err))
			return err
		}

		return nil
	})

	if txErr != nil {
		s.logger.Error("callback complete payment err", zap.Error(txErr))
	}
}

func createCompletePaymentRequest(
	ctx context.Context,
	payment *ent.Payment,
	paymentDetail *ent.PaymentCryptoDetail,
	currentBalance uint64,
) (*jarvis.CompletePaymentRequest, error) {
	return &jarvis.CompletePaymentRequest{
		PaymentId:     payment.ID,
		PaymentMethod: stark.MethodType(payment.Method),
		PaymentType:   stark.PaymentType(payment.Type),
		PaymentStatus: stark.Status(payment.Status),
		PaymentDetail: &jarvis.PaymentDetail{
			Crypto: &stark.CryptoPaymentDetail{
				CryptoType:        stark.CryptoType(paymentDetail.CryptoType),
				CryptoNetworkType: stark.CryptoNetworkType(paymentDetail.CryptoNetworkType),
				CryptoWalletName:  stark.CryptoWalletName(paymentDetail.CryptoWalletName),
				ReceiverAddress:   paymentDetail.ReceiverAddress,
				SenderAddress:     paymentDetail.SenderAddress,
				Amount:            paymentDetail.Amount,
				ReceivedAmount:    paymentDetail.ReceivedAmount,
				TxHash:            paymentDetail.TxHash,
				Fee:               paymentDetail.Fee,
				ImageUrl:          paymentDetail.ImageURL,
			},
		},
		MexCurrentBalance: currentBalance,
	}, nil
}

const (
	createOrderUMONote             = "Send withdraw request to UMO"
	createOrderUMOFailedNote       = "Send withdraw request to UMO failed"
	createOrderUMOSubmitFailedNote = "Withdraw request summit failed"
)

const (
	UMO_CodeOK = 1
)

func (s *ultronServer) transferCryptoAndSubmitWithdraw(
	payment *ent.Payment,
	paymentDetail *ent.PaymentCryptoDetail,
) {
	// Current status = APPROVED
	// Need to block teller to manual approve by changing status to SUBMITTING
	// Call create order to UMO
	// Waiting for callback
	// Handle submitted
	ctx := context.Background()
	cryptoType := stark.CryptoType(paymentDetail.CryptoType)
	cryptoNetworkType := stark.CryptoNetworkType(paymentDetail.CryptoNetworkType)

	err := s.updatePaymenStatusWithRevision(
		ctx,
		payment.ID,
		stark.Status_SUBMITTING,
		createOrderUMONote,
	)
	if err != nil {
		s.logger.Error("update payment status failed", zap.Error(err))
		return
	}

	hotWallets, err := s.GetSystemCryptoHotWallets(ctx, &stark.GetSystemCryptoHotWalletsRequest{
		CryptoType:        cryptoType,
		CryptoNetworkType: cryptoNetworkType,
		MerchantId:        payment.MerchantID,
	})
	if err != nil || len(hotWallets.Records) == 0 {
		s.logger.Error("get hot wallets failed", zap.Error(err), zap.Any("hot wallets", hotWallets.Records))
		return
	}

	reply, err := s.umoClient.CreateOrder(ctx, &umo.CreateOrderRequest{
		AppTransID: strconv.FormatInt(payment.ID, 10),
		AppUser:    fmt.Sprintf(mexUserIDFmtUMO, payment.MerchantID, payment.MerchantUserID),
		Amount:     paymentDetail.Amount,
		Currency: getCurrency(
			cryptoType,
			cryptoNetworkType,
		),
		Sender:      hotWallets.Records[0].Address,
		Recipient:   paymentDetail.ReceiverAddress,
		Description: "CMGP Withdraw crypto",
	})
	s.logger.Info("create order umo", zap.Any("reply", reply))
	// Handle error from reply
	// {"code":-8014,"message":"Invalid request on store 795237177952942219","data":{}}}
	if err != nil || reply.Code != UMO_CodeOK {
		s.logger.Error("send request to umo failed", zap.Error(err))
		err = s.updatePaymenStatusWithRevision(
			ctx,
			payment.ID,
			stark.Status_SUBMIT_FAILED,
			createOrderUMOFailedNote,
		)
		if err != nil {
			s.logger.Error("update payment status failed", zap.Error(err))
		}
		return
	}
	err = s.entClient.PaymentCryptoDetail.
		UpdateOne(paymentDetail).
		SetTxID(reply.Data.TransID).
		Exec(ctx)
	if err != nil {
		s.logger.Error("update tx_id failed", zap.Error(err))
	}

}

func (s *ultronServer) updatePaymenStatusWithRevision(
	ctx context.Context,
	paymentID int64,
	status stark.Status,
	note string,
) error {
	payment, err := s.entClient.Payment.
		UpdateOneID(paymentID).
		SetStatus(int32(status)).
		Save(ctx)
	if err != nil {
		return err
	}
	err = s.entClient.Revision.
		Create().
		SetCreatedAt(time.Now().UTC()).
		SetCreatedBy(systemUser).
		SetUpdatedBy(systemUser).
		SetPaymentID(paymentID).
		SetStatus(int32(status)).
		SetNote(note).
		SetPayment(payment).
		Exec(ctx)
	return err
}
