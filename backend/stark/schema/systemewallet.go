package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// SystemEWallet holds the schema definition for the SystemEWallet entity.
type SystemEWallet struct {
	ent.Schema
}

// Fields of the SystemEWallet.
func (SystemEWallet) Fields() []ent.Field {
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
		field.Int32("e_wallet_name").
			NonNegative().
			Default(0),
		field.Int32("status").
			NonNegative().
			Default(0),
		field.Int64("merchant_id").
			NonNegative().
			Default(0),
		field.String("account_phone_number").
			NotEmpty(),
		field.String("account_name").
			NotEmpty(),
		field.Uint64("balance").
			Default(0),
		field.Uint64("daily_balance").
			Default(0),
		field.Uint64("daily_balance_limit").
			Default(0),
		field.Int64("daily_used_amount").
			Default(0),
		field.Time("last_updated_balance").
			Optional().
			Nillable(),
	}
}

// Edges of the SystemEWallet.
func (SystemEWallet) Edges() []ent.Edge {
	return nil
}

func (SystemEWallet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("merchant_id", "e_wallet_name", "account_phone_number").Unique(),
	}
}
