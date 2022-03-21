package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PaymentBankingDetail
type PaymentBankingDetail struct {
	ent.Schema
}

// Fields of the PaymentBankingDetail.
func (PaymentBankingDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("created_by").
			NotEmpty(),
		field.String("updated_by").
			NotEmpty(),
		field.String("payment_code").
			Optional(),
		field.Int64("merchant_user_id").
			Optional(),
		field.Int32("merchant_user_bank_name").
			NonNegative().
			Optional(),
		field.String("merchant_user_account_number").
			Optional(),
		field.String("merchant_user_account_name").
			Optional(),
		field.Int32("system_account_bank_name").
			NonNegative().
			Optional(),
		field.String("system_account_number").
			Optional(),
		field.String("system_account_name").
			Optional(),
		field.String("image_url").
			Optional(),
		field.String("tx_id").
			Optional(),
		field.Uint64("amount").
			Optional(),
		field.Uint64("fee").
			Optional(),
	}
}

// Edges of the Banking.
func (PaymentBankingDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("payment", Payment.Type).
			Ref("payment_banking_detail").
			Unique().
			Required(),
	}
}
