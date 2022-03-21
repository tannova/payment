package telcosetting

import (
	"context"
	"encoding/json"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
)

type Setting interface {
	GetSettings(ctx context.Context, request *stark.GetSettingsRequest) (*stark.GetSettingsReply, error)
}

type setting struct {
	entClient *ent.Client
}

func New(
	entClient *ent.Client,
) Setting {
	return &setting{
		entClient: entClient,
	}
}

func (s setting) GetSettings(ctx context.Context, request *stark.GetSettingsRequest) (*stark.GetSettingsReply, error) {
	logger := logging.Logger(ctx)

	settings, err := s.entClient.Setting.Query().All(ctx)
	if err != nil {
		logger.Error("failed to get setting", zap.Error(err))
		return nil, status.Internal(err.Error())
	}

	m := make(map[string]*ent.Setting)

	for _, s := range settings {
		m[s.Name] = s
	}

	topUpAutoApproval := false
	if topUpAutoApprovalSetting, ok := m["top_up_auto_approval"]; ok {
		topUpAutoApproval = topUpAutoApprovalSetting.Value == "true"
	}
	enableThirdParty := false
	if enableThirdPartySetting, ok := m["enable_third_party"]; ok {
		enableThirdParty = enableThirdPartySetting.Value == "true"
	}

	chargeCardThirdPartyPriorities := []*stark.Provider{
		{
			Name:     "thenhanh365",
			Enable:   true,
			Priority: 1,
		},
		{
			Name:     "naptudong",
			Enable:   true,
			Priority: 2,
		},
	}
	if chargeCardThirdPartyPrioritiesSetting, ok := m["charge_card_third_party_priorities"]; ok {
		type providers struct {
			Name     string `json:"name"`
			Enable   bool   `json:"enable"`
			Priority int64  `json:"priority"`
		}
		var priorityProviders []providers
		json.Unmarshal([]byte(chargeCardThirdPartyPrioritiesSetting.Value), &priorityProviders)
		chargeCardThirdPartyPriorities = make([]*stark.Provider, 0)
		for _, p := range priorityProviders {
			chargeCardThirdPartyPriorities = append(chargeCardThirdPartyPriorities, &stark.Provider{
				Name:     p.Name,
				Priority: p.Priority,
				Enable:   p.Enable,
			})
		}
	}
	getCardThirdPartyPriorities := []*stark.Provider{
		{
			Name:     "banthe247",
			Enable:   true,
			Priority: 1,
		},
	}
	if getCardThirdPartyPrioritiesSetting, ok := m["get_card_third_party_priorities"]; ok {
		type providers struct {
			Name     string `json:"name"`
			Enable   bool   `json:"enable"`
			Priority int64  `json:"priority"`
		}
		var priorityProviders []providers
		json.Unmarshal([]byte(getCardThirdPartyPrioritiesSetting.Value), &priorityProviders)
		getCardThirdPartyPriorities = make([]*stark.Provider, 0)
		for _, p := range priorityProviders {
			getCardThirdPartyPriorities = append(getCardThirdPartyPriorities, &stark.Provider{
				Name:     p.Name,
				Priority: p.Priority,
				Enable:   p.Enable,
			})
		}
	}

	return &stark.GetSettingsReply{
		TopUpAutoApproval:   topUpAutoApproval,
		EnableThirdParty:    enableThirdParty,
		ChargeCardProviders: chargeCardThirdPartyPriorities,
		GetCardProviders:    getCardThirdPartyPriorities,
	}, nil
}
