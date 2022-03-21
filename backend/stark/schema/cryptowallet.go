package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// CryptoWallet holds the schema definition for the CryptoWallet entity.
type CryptoWallet struct {
	ent.Schema
}

// Fields of the CryptoWallet.
func (CryptoWallet) Fields() []ent.Field {
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
		field.Int64("merchant_user_id").
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
		field.Int32("status").
			NonNegative().
			Default(0),
	}
}

// Edges of the CryptoWallet.
func (CryptoWallet) Edges() []ent.Edge {
	return nil
}

func (CryptoWallet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("crypto_type", "crypto_network_type", "address").Unique(),
		index.Fields("merchant_id"),
		index.Fields("crypto_type", "crypto_network_type", "status"),
	}
}
