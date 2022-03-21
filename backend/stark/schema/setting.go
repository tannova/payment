package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Setting holds the schema definition for the Setting entity.
type Setting struct {
	ent.Schema
}

// Fields of the Setting.
func (Setting) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").
			NotEmpty(),
		field.String("type").
			NotEmpty(),
		field.String("value").
			NotEmpty(),
	}
}

// Edges of the Setting.
func (Setting) Edges() []ent.Edge {
	return nil
}

func (Setting) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
	}
}
