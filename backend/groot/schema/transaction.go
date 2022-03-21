package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

type Transaction struct {
	ent.Schema
}

// Fields of the Telco.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("tx_id"),
		field.String("provider"),
		field.Text("request"),
		field.Text("response"),
		field.Text("callback_response"),
		field.Text("code"),
		field.Text("serial"),
		field.Uint64("amount").
			Default(0).
			Min(0),
		field.String("status"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
