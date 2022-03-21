package banthe247

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	ie "gitlab.com/mcuc/monorepo/backend/groot/internal/error"
)

const (
	_loginPath = "/v2/PayCard/DangNhap"
)


type GetTokenReply struct {
	Token string
}

func (p provider) GetToken(ctx context.Context) (*GetTokenReply, error) {
	logger := logging.Logger(ctx)

	r, err := http.NewRequest("POST", p.cfg.baseURL + _loginPath, nil)
	if err != nil {
		logger.Error("failed to init request", zap.Error(err))
		return nil, err
	}

	client := &http.Client{}
	q := r.URL.Query()
	q.Add("security", p.cfg.security)
	q.Add("userName", p.cfg.userName)
	q.Add("password", p.cfg.password)
	decoded, err := url.QueryUnescape(q.Encode())
	r.URL.RawQuery = decoded
	r.WithContext(ctx)

	resp, err := client.Do(r)
	if err != nil {
		logger.Error("failed to get token", zap.Error(err))
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("failed to read response body", zap.Error(err))
		return nil, err
	}

	if strings.Contains(string(body), "status") {
		logger.Error("failed to get token", zap.Any("response", string(body)))
		return nil, ie.Internal("failed to get token")
	}

	token := strings.Replace(string(body), "\"", "", -1)

	return &GetTokenReply{
		Token: token,
	}, nil
}