package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Telco holds the schema definition for the Telco entity.
type Telco struct {
	ent.Schema
}

// Fields of the Telco.
func (Telco) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("telco_name").
			NonNegative().
			Default(0),
		field.String("code").
			NotEmpty(),
		field.String("serial").
			NotEmpty(),
		field.Uint64("amount").
			Default(0).
			Min(0),
		field.Bool("used").
			Default(false),
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
