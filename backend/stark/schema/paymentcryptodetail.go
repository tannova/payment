package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PaymentCryptoDetail holds the schema definition for the PaymentCryptoDetail entity.
type PaymentCryptoDetail struct {
	ent.Schema
}

// Fields of the PaymentCryptoDetail.
func (PaymentCryptoDetail) Fields() []ent.Field {
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
		field.Int32("crypto_type").
			NonNegative().
			Default(0),
		field.Int32("crypto_network_type").
			NonNegative().
			Default(0),
		field.Int32("crypto_wallet_name").
			NonNegative().
			Default(0),
		field.String("receiver_address").
			NotEmpty(),
		field.String("sender_address").
			Optional(),
		field.Float("amount").
			Optional(),
		field.Float("received_amount").
			Optional(),
		field.String("tx_hash").
			Optional(),
		field.String("tx_id").
			Optional(),
		field.Float("fee").
			Optional(),
		field.String("image_url").
			Optional(),
	}
}

func (PaymentCryptoDetail) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tx_id"),
	}
}

// Edges of the PaymentCryptoDetail.
func (PaymentCryptoDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("payment", Payment.Type).
			Ref("payment_crypto_detail").
			Unique().
			Required(),
	}
}
