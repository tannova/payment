package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PaymentTelcoDetail
type PaymentTelcoDetail struct {
	ent.Schema
}

// Fields of the Telco.
func (PaymentTelcoDetail) Fields() []ent.Field {
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
		field.Int32("telco_name").
			NonNegative().
			Default(0),
		field.String("serial_number").
			Default(""),
		field.String("card_id").
			Default(""),
		field.Uint64("charged_amount").
			Default(0),
		field.Uint64("amount").
			Default(0),
	}
}

// Edges of the Telco.
func (PaymentTelcoDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("payment", Payment.Type).
			Ref("payment_telco_detail").
			Unique(),
	}
}
