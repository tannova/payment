package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// SystemCryptoHotWallet holds the schema definition for the SystemCryptoHotWallet entity.
type SystemCryptoHotWallet struct {
	ent.Schema
}

// Fields of the SystemCryptoHotWallet.
func (SystemCryptoHotWallet) Fields() []ent.Field {
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
		field.Int32("crypto_type").
			NonNegative().
			Default(0),
		field.Int32("crypto_network_type").
			NonNegative().
			Default(0),
		field.String("address").
			NotEmpty(),
		field.Float("total_balance").
			Default(0),
		field.Float("balance").
			Default(0),
		field.Int32("status").
			NonNegative().
			Default(0),
	}
}

// Edges of the SystemCryptoHotWallet.
func (SystemCryptoHotWallet) Edges() []ent.Edge {
	return nil
}

func (SystemCryptoHotWallet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("crypto_type", "crypto_network_type", "address").Unique(),
		index.Fields("merchant_id"),
		index.Fields("crypto_type", "crypto_network_type", "status"),
	}
}
