package transformer

import (
	groot "gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/groot/pkg/ent"
)

func GetTelco(telco *ent.Telco) *groot.Card {
	return &groot.Card{
		Code:   telco.Code,
		Serial: telco.Serial,
		TelcoName: groot.TelcoName(telco.TelcoName),
		Amount: telco.Amount,
	}
}
