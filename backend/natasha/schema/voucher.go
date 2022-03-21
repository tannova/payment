package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Voucher holds the schema definition for the Voucher entity.
type Voucher struct {
	ent.Schema
}

// Fields of the Voucher.
func (Voucher) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("merchant_id").
			NonNegative().
			Default(0),
		field.Int64("payment_id").
			Optional().
			NonNegative().
			Default(0),
		// amount here can be negative if the type is withdraw
		field.Int64("amount"),
		field.Int32("type").
			NonNegative().
			Default(0),
		field.Int32("status").
			Optional().
			NonNegative().
			Default(0),
		field.Int32("payee_provider").
			Optional().
			NonNegative().
			Default(0),
		field.String("payee_account").
			Optional(),
		field.String("payee_name").
			Optional(),
		field.Int32("payer_provider").
			Optional().
			NonNegative().
			Default(0),
		field.String("payer_account").
			Optional(),
		field.String("payer_name").
			Optional(),
		field.String("tx_id").
			Optional(),
		field.String("image_url").
			Optional(),
		field.String("creator_note").
			Optional(),
		field.String("handled_by").
			Optional(),
		field.String("handler_note").
			Optional(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("created_by").
			NotEmpty(),
		field.String("updated_by").
			NotEmpty(),
	}
}
