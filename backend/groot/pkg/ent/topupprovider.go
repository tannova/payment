// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/mcuc/monorepo/backend/groot/pkg/ent/topupprovider"
)

// TopUpProvider is the model entity for the TopUpProvider schema.
type TopUpProvider struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name int64 `json:"name,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority int64 `json:"priority,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Enabled holds the value of the "enabled" field.
	Enabled bool `json:"enabled,omitempty"`
	// Type holds the value of the "type" field.
	Type int32 `json:"type,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TopUpProvider) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case topupprovider.FieldEnabled:
			values[i] = new(sql.NullBool)
		case topupprovider.FieldID, topupprovider.FieldName, topupprovider.FieldPriority, topupprovider.FieldType:
			values[i] = new(sql.NullInt64)
		case topupprovider.FieldKey, topupprovider.FieldCreatedBy, topupprovider.FieldUpdatedBy:
			values[i] = new(sql.NullString)
		case topupprovider.FieldCreatedAt, topupprovider.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TopUpProvider", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TopUpProvider fields.
func (tup *TopUpProvider) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case topupprovider.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tup.ID = int64(value.Int64)
		case topupprovider.FieldName:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tup.Name = value.Int64
			}
		case topupprovider.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				tup.Priority = value.Int64
			}
		case topupprovider.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				tup.Key = value.String
			}
		case topupprovider.FieldEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enabled", values[i])
			} else if value.Valid {
				tup.Enabled = value.Bool
			}
		case topupprovider.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				tup.Type = int32(value.Int64)
			}
		case topupprovider.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tup.CreatedAt = value.Time
			}
		case topupprovider.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tup.UpdatedAt = value.Time
			}
		case topupprovider.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				tup.CreatedBy = value.String
			}
		case topupprovider.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				tup.UpdatedBy = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TopUpProvider.
// Note that you need to call TopUpProvider.Unwrap() before calling this method if this TopUpProvider
// was returned from a transaction, and the transaction was committed or rolled back.
func (tup *TopUpProvider) Update() *TopUpProviderUpdateOne {
	return (&TopUpProviderClient{config: tup.config}).UpdateOne(tup)
}

// Unwrap unwraps the TopUpProvider entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tup *TopUpProvider) Unwrap() *TopUpProvider {
	tx, ok := tup.config.driver.(*txDriver)
	if !ok {
		panic("ent: TopUpProvider is not a transactional entity")
	}
	tup.config.driver = tx.drv
	return tup
}

// String implements the fmt.Stringer.
func (tup *TopUpProvider) String() string {
	var builder strings.Builder
	builder.WriteString("TopUpProvider(")
	builder.WriteString(fmt.Sprintf("id=%v", tup.ID))
	builder.WriteString(", name=")
	builder.WriteString(fmt.Sprintf("%v", tup.Name))
	builder.WriteString(", priority=")
	builder.WriteString(fmt.Sprintf("%v", tup.Priority))
	builder.WriteString(", key=")
	builder.WriteString(tup.Key)
	builder.WriteString(", enabled=")
	builder.WriteString(fmt.Sprintf("%v", tup.Enabled))
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", tup.Type))
	builder.WriteString(", created_at=")
	builder.WriteString(tup.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(tup.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", created_by=")
	builder.WriteString(tup.CreatedBy)
	builder.WriteString(", updated_by=")
	builder.WriteString(tup.UpdatedBy)
	builder.WriteByte(')')
	return builder.String()
}

// TopUpProviders is a parsable slice of TopUpProvider.
type TopUpProviders []*TopUpProvider

func (tup TopUpProviders) config(cfg config) {
	for _i := range tup {
		tup[_i].config = cfg
	}
}
