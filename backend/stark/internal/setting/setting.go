package setting

import (
	"context"

	"go.uber.org/zap"

	conf "gitlab.com/mcuc/monorepo/backend/stark/pkg/config"
)

type Setting interface {
	GetConfigPayment(ctx context.Context) *conf.PaymentCode
	GetCryptoConfig() *conf.PaymentCrypto
	GetMapBankName() map[string]string
	GetMapEWalletName() map[string]string
}

type settingImpl struct {
	logger         *zap.Logger
	defaultSetting *conf.DefaultSetting
}

func NewSettingImpl(logger *zap.Logger, config *conf.Config) Setting {
	return &settingImpl{logger: logger,
		defaultSetting: config.DefaultSetting}
}

func (impl *settingImpl) GetConfigPayment(ctx context.Context) *conf.PaymentCode {
	return impl.defaultSetting.Payment
}

func (impl *settingImpl) GetCryptoConfig() *conf.PaymentCrypto {
	return impl.defaultSetting.Crypto
}

func (impl *settingImpl) GetMapBankName() map[string]string {
	return impl.defaultSetting.MapBankName
}

func (impl *settingImpl) GetMapEWalletName() map[string]string {
	return impl.defaultSetting.MapEWalletName
}
