package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("created_by").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.String("updated_by").
			NotEmpty(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("approved_by").
			Optional().Nillable(),
		field.Time("approved_at").
			Optional().Nillable(),
		field.Int64("merchant_id").
			NonNegative().
			Default(0),
		field.Int64("merchant_user_id").
			NonNegative().
			Default(0),
		// MethodType: T, I, ...
		field.Int32("method").
			NonNegative().
			Default(0),
		// PaymentType:TOPUP, WITHDRAW
		field.Int32("type").
			NonNegative().
			Default(0),
		field.Int32("status").
			NonNegative().
			Default(0),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("payment_banking_detail", PaymentBankingDetail.Type),
		edge.To("payment_telco_detail", PaymentTelcoDetail.Type),
		edge.To("payment_e_wallet_detail", PaymentEWalletDetail.Type),
		edge.To("payment_crypto_detail", PaymentCryptoDetail.Type),
		edge.To("revisions", Revision.Type),
	}
}
