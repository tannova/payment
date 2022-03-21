package telcopayment

import (
	"context"
	"fmt"
	codes "gitlab.com/mcuc/monorepo/backend/stark/pkg/code"
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
	epayment "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

func (a telcoPayment) ApproveTopUp(ctx context.Context, callerID string, request *stark.ApproveTelcoTopUpRequest) (*stark.ApproveTelcoTopUpReply, error) {
	logger := logging.Logger(ctx)

	metaData, err := utils.GetMetadata(ctx)
	if err != nil {
		metaData = metadata.Pairs("x-user-id", "1")
	}
	outCtx := metadata.NewOutgoingContext(ctx, metaData)

	expectedType := int32(stark.PaymentType_TOPUP)
	expectedMethod := int32(stark.MethodType_P)
	expectedStatus := int32(stark.Status_CREATED)
	nextSuccessStatus := int32(stark.Status_APPROVED)
	nextFailedStatus := int32(stark.Status_REJECTING)

	err = tx.WithTransaction(ctx, logger, a.entClient, func(ctx context.Context, tx tx.Tx) error {
		payment, err := tx.Client().Payment.
			Query().
			Where(
				epayment.IDEQ(request.GetPaymentId()),
				epayment.TypeEQ(expectedType),
				epayment.MethodEQ(expectedMethod),
				epayment.StatusEQ(expectedStatus),
			).
			WithPaymentBankingDetail().
			ForUpdate().
			First(ctx)
		if err != nil {
			logger.Error("failed to get banking top up payment by id", zap.Error(err))
			return err
		}
		paymentDetail, err := payment.QueryPaymentTelcoDetail().Only(ctx)
		if err != nil {
			logger.Error("failed to get banking top up payment detail", zap.Error(err))
			return err
		}

		payment, err = payment.Update().
			SetApprovedBy(callerID).
			SetApprovedAt(time.Now().UTC()).
			Save(ctx)
		if err != nil {
			logger.Error("failed to get update approve by", zap.Error(err))
			return err
		}

		reply, err := a.natashaCli.GetMerchant(outCtx, &natasha.GetMerchantRequest{
			Id: payment.MerchantID,
		})
		if err != nil {
			logger.Error("get merchant err", zap.Error(err))
			return err
		}

		setting, err := a.telcoSetting.GetSettings(ctx, &stark.GetSettingsRequest{})
		if err != nil {
			logger.Error("failed to get setting", zap.Error(err))
			return err
		}
		if setting.EnableThirdParty {
			providers, err := a.getEnableChargeCardProviders(setting.ChargeCardProviders)
			if err != nil {
				return err
			}
			_, err = a.grootCli.ChargeCard(outCtx, &groot.ChargeCardRequest{
				Card: &groot.Card{
					Amount:    paymentDetail.Amount,
					TelcoName: groot.TelcoName(paymentDetail.TelcoName),
					Code:      paymentDetail.CardID,
					Serial:    paymentDetail.SerialNumber,
				},
				PaymentId: fmt.Sprint(payment.ID),
				Providers: providers,
			})
			if err != nil {
				logger.Error("failed to charge card", zap.Error(err))
				payment, err = a.updatePaymentStatus(ctx, tx, payment, callerID, nextFailedStatus, "Card is wrong")

				completeRequest, err := createCompletePaymentRequest(
					ctx,
					payment,
					paymentDetail,
					// No need balance for this case
					// Will add this value when needed
					0,
				)
				if err != nil {
					logger.Error("get merchant err", zap.Error(err))
					return err
				}
				err = a.jarvisClient.CompletePayment(
					ctx,
					reply.Merchant.WebhookUrl,
					completeRequest,
					a.completeRejectTopUpCallback)

				return err
			}

			payment, err = a.updatePaymentStatus(ctx, tx, payment, callerID, nextSuccessStatus, request.Note)
			return err
		}

		makePaymentReply, err := a.blackWidowCli.MakePayment(outCtx, &natasha.MakePaymentRequest{
			MerchantId: payment.MerchantID,
			Amount:     paymentDetail.Amount,
			Type:       natasha.PaymentType_USER_TOP_UP,
		})
		if err != nil {
			logger.Error("failed to make payment", zap.Error(err))
			a.updatePaymentStatus(ctx, tx, payment, callerID, nextFailedStatus, err.Error())
			return err
		}

		payment, err = a.updatePaymentStatus(ctx, tx, payment, callerID, nextSuccessStatus, request.Note)
		if err != nil {
			return err
		}

		completeRequest, err := createCompletePaymentRequest(
			ctx,
			payment,
			paymentDetail,
			makePaymentReply.Balance,
		)
		if err != nil {
			logger.Error("get merchant err", zap.Error(err))
			return err
		}

		err = a.jarvisClient.CompletePayment(
			ctx,
			reply.Merchant.WebhookUrl,
			completeRequest,
			a.completeTopUpCallback)
		if err != nil {
			logger.Error("complete payment err", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, status.Internal(err.Error())
	}

	return &stark.ApproveTelcoTopUpReply{}, nil
}

func (a telcoPayment) completeRejectTopUpCallback(
	request *jarvis.CompletePaymentRequest,
	responseErr error,
) {
	a.completeCallback(
		request,
		responseErr,
		PaymentTransition{
			ExpectedType:      stark.PaymentType_TOPUP,
			ExpectedMethod:    stark.MethodType_P,
			ExpectedStatus:    stark.Status_REJECTING,
			NextSuccessStatus: stark.Status_REJECTED,
			NextFailedStatus:  stark.Status_REJECT_FAILED,
		})
}

func (a telcoPayment) getEnableChargeCardProviders(providers []*stark.Provider) ([]string, error) {
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
		return nil, status.Error(codes.Code_TELCO_CHARGE_CARD_PROVIDER_EMPTY)
	}

	return res, nil
}

func (a telcoPayment) completeTopUpCallback(
	request *jarvis.CompletePaymentRequest,
	responseErr error,
) {
	a.completeCallback(
		request,
		responseErr,
		PaymentTransition{
			ExpectedType:      stark.PaymentType_TOPUP,
			ExpectedMethod:    stark.MethodType_P,
			ExpectedStatus:    stark.Status_APPROVED,
			NextSuccessStatus: stark.Status_COMPLETED,
			NextFailedStatus:  stark.Status_APPROVE_FAILED,
		})
}
