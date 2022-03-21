package morgan

import (
	"context"
	"sort"
	"time"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	groot "gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/jarvis"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	codes "gitlab.com/mcuc/monorepo/backend/stark/pkg/code"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	epayment "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

func (s *morganServer) ApproveTelcoWithdraw(ctx context.Context, request *stark.ApproveTelcoWithdrawRequest) (*stark.ApproveTelcoWithdrawReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	logger := logging.Logger(ctx)
	_, tellerStr, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		logger.Error("can not get teller id", zap.Error(err))
		return nil, status.Unauthenticated
	}

	metaData, err := utils.GetMetadata(ctx)
	if err != nil {
		return nil, err
	}
	outCtx := metadata.NewOutgoingContext(ctx, metaData)

	expectedType := int32(stark.PaymentType_WITHDRAW)
	expectedMethod := int32(stark.MethodType_P)
	expectedStatus := int32(stark.Status_CREATED)
	nextSuccessStatus := int32(stark.Status_APPROVED)
	nextFailedStatus := int32(stark.Status_APPROVE_FAILED)

	var (
		payment       *ent.Payment
		paymentDetail *ent.PaymentTelcoDetail
		merchant      *natasha.Merchant
	)
	err = tx.WithTransaction(ctx, logger, s.entClient, func(ctx context.Context, tx tx.Tx) error {
		payment, err = tx.Client().Payment.
			Query().
			Where(
				epayment.IDEQ(request.GetPaymentId()),
				epayment.TypeEQ(expectedType),
				epayment.MethodEQ(expectedMethod),
				epayment.StatusEQ(expectedStatus),
			).
			ForUpdate().
			First(ctx)
		if err != nil {
			logger.Error("failed to get banking top up payment by id", zap.Error(err))
			return err
		}

		paymentDetail, err = payment.QueryPaymentTelcoDetail().Only(ctx)
		if err != nil {
			logger.Error("failed to get banking withdraw payment detail", zap.Error(err))
			return err
		}

		payment, err = payment.Update().
			SetApprovedBy(tellerStr).
			SetApprovedAt(time.Now().UTC()).
			Save(ctx)
		if err != nil {
			logger.Error("failed to get update approve by", zap.Error(err))
			return err
		}

		payment, err = s.updatePaymentStatus(ctx, tx, payment, tellerStr, nextSuccessStatus, request.GetNote())
		if err != nil {
			logger.Error("failed to update payment status to success", zap.Error(err))
			return err
		}

		// Start network call to MEX to inform
		reply, err := s.natashaClient.GetMerchant(outCtx, &natasha.GetMerchantRequest{
			Id: payment.MerchantID,
		})
		if err != nil {
			logger.Error("failed to get merchant", zap.Error(err))
			return err
		}
		merchant = reply.GetMerchant()

		mexBalanceReply, err := s.blackWidowCli.MakePayment(outCtx, &natasha.MakePaymentRequest{
			MerchantId: payment.MerchantID,
			Amount:     paymentDetail.Amount,
			Type:       natasha.PaymentType_USER_WITHDRAW,
		})
		if err != nil {
			logger.Error("failed to make payment", zap.Error(err))
			return err
		}

		card := &groot.Card{
			Code:      "123123123",
			Serial:    "23123123123",
			TelcoName: groot.TelcoName(paymentDetail.TelcoName),
			Amount:    paymentDetail.Amount,
		}

		setting, err := s.telcoSetting.GetSettings(ctx, &stark.GetSettingsRequest{})
		if err != nil {
			logger.Error("failed to get setting", zap.Error(err))
			return err
		}
		if setting.EnableThirdParty {
			providers, err := s.getEnableGetCardProviders(setting.GetCardProviders)
			if err != nil {
				return err
			}

			card, err = s.grootCli.GetCard(outCtx, &groot.GetCardRequest{
				TelcoName: groot.TelcoName(paymentDetail.TelcoName),
				Amount:    paymentDetail.Amount,
				Providers: providers,
			})
			if err != nil {
				logger.Error("failed to get card", zap.Error(err))
				return err
			}
		}

		paymentDetail, err = paymentDetail.Update().
			SetCardID(card.Code).
			SetSerialNumber(card.Serial).
			Save(ctx)
		if err != nil {
			logger.Error("failed to update payment detail",
				zap.String("cardCode", card.Code),
				zap.Error(err))
			return err
		}

		payment, err = s.updatePaymentStatus(ctx, tx, payment, tellerStr, int32(stark.Status_SUBMITTED), request.GetNote())
		if err != nil {
			logger.Error("failed to update payment status to success", zap.Error(err))
			return err
		}

		completeRequest := createCompletePaymentRequest(
			ctx,
			payment,
			paymentDetail,
			mexBalanceReply.Balance,
		)
		err = s.jarvisClient.CompletePayment(
			ctx,
			merchant.GetWebhookUrl(),
			completeRequest,
			s.completeWithdrawCallback)

		return err
	})
	if err != nil {
		if payment != nil {
			_ = tx.WithTransaction(ctx, logger, s.entClient, func(ctx context.Context, tx tx.Tx) error {
				_, err := s.updatePaymentStatus(ctx, tx, payment, tellerStr, nextFailedStatus, err.Error())
				if err != nil {
					logger.Error("failed to update payment status to failed", zap.Error(err))
				}

				return err
			})
		}

		return nil, status.Internal(err.Error())
	}

	return &stark.ApproveTelcoWithdrawReply{}, nil
}

func (s *morganServer) completeWithdrawCallback(
	request *jarvis.CompletePaymentRequest,
	responseErr error,
) {
	s.completeCallback(
		request,
		responseErr,
		PaymentTransition{
			ExpectedType:      stark.PaymentType_WITHDRAW,
			ExpectedMethod:    stark.MethodType_P,
			ExpectedStatus:    stark.Status_SUBMITTED,
			NextSuccessStatus: stark.Status_COMPLETED,
			NextFailedStatus:  stark.Status_APPROVE_FAILED,
		})
}

func (s *morganServer) getEnableGetCardProviders(providers []*stark.Provider) ([]string, error) {
	sort.Slice(providers, func(i, j int) bool {
		return providers[i].Priority < providers[j].Priority
	})

	res := make([]string, 0)
	for _, p := range providers {
		if p.Enable {
			res = append(res, p.Name)
		}
	}

	if len(res) == 0 {
		return nil, status.Error(codes.Code_TELCO_GET_CARD_PROVIDER_EMPTY)
	}

	return res, nil
}
