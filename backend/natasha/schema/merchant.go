package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Merchant ...
type Merchant struct {
	ent.Schema
}

// Fields ...
func (Merchant) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name"),
		field.String("email_contact").
			NotEmpty().
			Unique(),
		// the domain might change for each environments
		// so we only store path in our storage (S3)
		field.String("logo_path"),
		field.String("webhook_url").
			NotEmpty(),
		field.String("slack_webhook_url"),
		field.Uint64("safety_limit").
			Default(uint64(50000000)),
		field.String("created_by").
			MaxLen(255).
			NotEmpty(),
		field.String("updated_by").
			MaxLen(255).
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}
