package ultron

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	ecw "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/cryptowallet"
	pcd "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentcryptodetail"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

const (
	paymentTypeDeposit  = "DEPOSIT"
	paymentTypeWithdraw = "WITHDRAW"
)

const (
	storeUMO = "UMO_%s"
)

func (s *ultronServer) CallbackTransaction(
	ctx context.Context,
	request *stark.CallbackTransactionRequest,
) (*stark.CallbackTransactionReply, error) {

	logger := logging.Logger(ctx)
	logger.Info("CallbackTransaction", zap.Any("req", request))

	var result error
	switch request.Type {
	case paymentTypeDeposit:
		result = s.processDeposit(ctx, request)
	case paymentTypeWithdraw:
		result = s.processWithdraw(ctx, request)
	}
	if result != nil {
		logger.Error("error when process callback", zap.Error(result))
	}

	return &stark.CallbackTransactionReply{}, nil
}

var (
	errInvalidCurrency = errors.New("invalid currency")
	errInvalidStatus   = errors.New("wallet not in used")
)

func (s *ultronServer) processDeposit(ctx context.Context, request *stark.CallbackTransactionRequest) error {
	logger := logging.Logger(ctx)
	cryptoType, networkType := getCryptoTypes(request.GetCurrency())
	if cryptoType == stark.CryptoType_CRYPTO_TYPE_UNSPECIFIED ||
		networkType == stark.CryptoNetworkType_CRYPTO_NETWORK_TYPE_UNSPECIFIED {
		return errInvalidCurrency
	}
	// check recipient address is valid from the pool address
	receiverAddress := strings.TrimPrefix(request.GetRecipient(), request.GetCurrency()+"_")
	senderAddress := strings.TrimPrefix(request.GetSender(), request.GetCurrency()+"_")
	wallet, err := s.entClient.CryptoWallet.
		Query().
		Where(
			ecw.Address(receiverAddress),
			ecw.CryptoTypeEQ(int32(cryptoType)),
			ecw.CryptoNetworkTypeEQ(int32(networkType)),
		).First(ctx)
	if err != nil {
		return err
	}
	if wallet.Status != int32(stark.CryptoWalletStatus_USED) {
		return errInvalidStatus
	}

	err = tx.WithTransaction(ctx, logger, s.entClient, func(ctx context.Context, tx tx.Tx) error {
		caller := fmt.Sprintf(storeUMO, request.GetStoreId())
		status := int32(stark.Status_CREATED)
		now := time.Now().UTC()

		payment, err := tx.Client().Payment.Create().
			SetCreatedAt(now).
			SetCreatedBy(caller).
			SetUpdatedBy(caller).
			SetUpdatedAt(now).
			SetMerchantUserID(wallet.MerchantUserID).
			SetMerchantID(wallet.MerchantID).
			SetMethod(int32(stark.MethodType_C)).
			SetType(int32(stark.PaymentType_TOPUP)).
			SetStatus(status).
			Save(ctx)
		if err != nil {
			logger.Error("can not create payment", zap.Error(err))
			return err
		}
		_, err = tx.Client().PaymentCryptoDetail.Create().
			SetCreatedAt(now).
			SetCreatedBy(caller).
			SetUpdatedBy(caller).
			SetUpdatedAt(now).
			SetPaymentID(payment.ID).
			SetCryptoType(wallet.CryptoType).
			SetCryptoNetworkType(wallet.CryptoNetworkType).
			SetCryptoWalletName(int32(stark.CryptoWalletName_UMO)).
			SetPayment(payment).
			SetReceiverAddress(receiverAddress).
			SetSenderAddress(senderAddress).
			SetAmount(request.GetAmount()).
			SetReceivedAmount(request.GetReceivedAmount()).
			SetTxHash(request.GetTxHash()).
			SetTxID(request.TransId).
			SetFee(request.GetBcFee()).
			Save(ctx)

		if err != nil {
			logger.Error("can not create crypto payment detail", zap.Error(err))
			return err
		}

		_, err = tx.Client().Revision.Create().
			SetCreatedAt(now).
			SetCreatedBy(caller).
			SetUpdatedBy(caller).
			SetUpdatedAt(now).
			SetPaymentID(payment.ID).
			SetStatus(status).
			SetPayment(payment).
			SetNote(request.GetMessage()).
			Save(ctx)

		if err != nil {
			logger.Error("can not create revision", zap.Error(err))
			return err
		}

		return nil
	})

	return err
}

const (
	transaction_FAILED    = "FAILED"
	transaction_SUBMITTED = "SUBMITTED"
	transaction_SUCCESS   = "SUCCESS"
)

func (s *ultronServer) processWithdraw(ctx context.Context, request *stark.CallbackTransactionRequest) error {
	// logger := logging.Logger(ctx)
	// request.State: FAILED, SUBMITTED, SUCCESS
	paymentDetail, err := s.entClient.PaymentCryptoDetail.
		Query().
		Where(pcd.TxIDEQ(request.AppTransId)).
		Only(ctx)
	if err != nil {
		return err
	}
	payment, err := paymentDetail.QueryPayment().Only(ctx)
	switch request.State {
	case transaction_FAILED:
		err = s.updatePaymenStatusWithRevision(
			ctx,
			payment.ID,
			stark.Status_SUBMIT_FAILED,
			createOrderUMOSubmitFailedNote,
		)
	case transaction_SUBMITTED:
		// only for update status, not need to handle for now
	case transaction_SUCCESS:
		metaData, err := utils.GetMetadata(ctx)
		if err != nil {
			// TODO(tiennv147): setup user
			metaData = metadata.Pairs("x-user-id", "1")
		}
		// TODO(tiennv147): update status to APPROVED again before call
		outCtx := metadata.NewOutgoingContext(ctx, metaData)
		_, _ = s.SubmitCryptoWithdraw(outCtx, &stark.SubmitCryptoWithdrawRequest{
			PaymentId:     payment.ID,
			SenderAddress: request.Sender,
			Amount:        request.Amount,
			Fee:           request.BcFee,
			TxHash:        request.TxHash,
		})
	}
	return err
}

func getCryptoTypes(currency string) (stark.CryptoType, stark.CryptoNetworkType) {
	switch currency {
	case "USD":
		return stark.CryptoType_USDT, stark.CryptoNetworkType_ERC20
	case "ASD":
		return stark.CryptoType_USDT, stark.CryptoNetworkType_TRC20
	case "UNU":
		return stark.CryptoType_USDT, stark.CryptoNetworkType_BEP20
	default:
		return stark.CryptoType_CRYPTO_TYPE_UNSPECIFIED, stark.CryptoNetworkType_CRYPTO_NETWORK_TYPE_UNSPECIFIED
	}
}

func getCurrency(_ stark.CryptoType, cryptoNetworkType stark.CryptoNetworkType) string {
	currency := ""
	// Only support USDT for now
	switch cryptoNetworkType {
	case stark.CryptoNetworkType_BEP20:
		currency = "UNU"
	case stark.CryptoNetworkType_TRC20:
		currency = "ASD"
	}
	return currency
}
