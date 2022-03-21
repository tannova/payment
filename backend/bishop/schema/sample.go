package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Sample struct {
	ent.Schema
}

// Fields of the Sample
func (Sample) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
	}
}
