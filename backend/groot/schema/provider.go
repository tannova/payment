package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Telco holds the schema definition for the Telco entity.
type TopUpProvider struct {
	ent.Schema
}

// Fields of the Telco.
func (TopUpProvider) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("name").
			NonNegative().
			Default(0),
		field.Int64("priority").
			NonNegative().
			Default(0),
		field.String("key").
			NotEmpty(),
		field.Bool("enabled").
			Default(false),
		field.Int32("type").
			NonNegative().
			Default(0),
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
