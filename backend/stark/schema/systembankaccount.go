package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// SystemBankAccount holds the schema definition for the SystemBankAccount entity.
type SystemBankAccount struct {
	ent.Schema
}

// Fields of the SystemBankAccount.
func (SystemBankAccount) Fields() []ent.Field {
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
		field.Int64("merchant_id").
			NonNegative().
			Default(0),
		field.Int32("status").
			NonNegative().
			Default(0),
		field.Int32("bank_name").
			NonNegative().
			Default(0),
		field.String("account_number").
			NotEmpty(),
		field.String("account_name").
			NotEmpty(),
		field.String("branch"),
		field.Uint64("balance").
			Default(0),
		field.Time("last_updated_balance").
			Optional().
			Nillable(),
		field.Uint64("daily_balance_limit").
			Default(0),
		field.Int64("daily_used_amount").
			Default(0),
		field.Uint64("daily_balance").
			Default(0),
	}
}

// Edges of the SystemBankAccount.
func (SystemBankAccount) Edges() []ent.Edge {
	return nil
}

func (SystemBankAccount) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("bank_name", "account_number", "account_name").Unique(),
	}
}
