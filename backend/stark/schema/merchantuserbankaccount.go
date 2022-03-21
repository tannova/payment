package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// MerchantUserBankAccount holds the schema definition for the MerchantUserBankAccount entity.
type MerchantUserBankAccount struct {
	ent.Schema
}

// Fields of the MerchantUserBankAccount.
func (MerchantUserBankAccount) Fields() []ent.Field {
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
		field.Int32("bank_name").
			NonNegative().
			Default(0),
		field.String("account_number").
			NotEmpty(),
		field.String("account_name").
			NotEmpty(),
	}
}

// Edges of the MerchantUserBankAccount.
func (MerchantUserBankAccount) Edges() []ent.Edge {
	return nil
}

func (MerchantUserBankAccount) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("bank_name", "account_number", "account_name").Unique(),
	}
}
