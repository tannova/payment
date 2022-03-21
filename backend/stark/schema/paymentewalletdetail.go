package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PaymentEWalletDetail holds the schema definition for the PaymentEWalletDetail entity.
type PaymentEWalletDetail struct {
	ent.Schema
}

// Fields of the PaymentEWalletDetail.
func (PaymentEWalletDetail) Fields() []ent.Field {
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
		field.Int32("e_wallet_name").
			NonNegative().
			Optional(),
		field.String("merchant_user_account_phone_number").
			Optional(),
		field.String("merchant_user_account_name").
			Optional(),
		field.String("system_account_phone_number").
			Optional(),
		field.String("system_account_name").
			Optional(),
		field.Uint64("amount").
			Optional(),
		field.Uint64("fee").
			Optional(),
		field.String("note").
			Optional(),
		field.String("image_url").
			Optional(),
		field.String("tx_id").
			Optional(),
	}
}

// Edges of the PaymentEWalletDetail.
func (PaymentEWalletDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("payment", Payment.Type).
			Ref("payment_e_wallet_detail").
			Unique().
			Required(),
	}
}
