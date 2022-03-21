package vision

import (
	"context"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
)

func (s *visionServer) GetPaymentInfoByPaymentCode(ctx context.Context, request *stark.GetPaymentInfoByPaymentCodeRequest) (*stark.GetPaymentInfoByPaymentCodeReply, error) {
	return s.getPaymentInfoByPaymentCode.Get(ctx, request)
}
