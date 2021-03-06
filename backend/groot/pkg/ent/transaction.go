// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/mcuc/monorepo/backend/groot/pkg/ent/transaction"
)

// Transaction is the model entity for the Transaction schema.
type Transaction struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// TxID holds the value of the "tx_id" field.
	TxID string `json:"tx_id,omitempty"`
	// Provider holds the value of the "provider" field.
	Provider string `json:"provider,omitempty"`
	// Request holds the value of the "request" field.
	Request string `json:"request,omitempty"`
	// Response holds the value of the "response" field.
	Response string `json:"response,omitempty"`
	// CallbackResponse holds the value of the "callback_response" field.
	CallbackResponse string `json:"callback_response,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Serial holds the value of the "serial" field.
	Serial string `json:"serial,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount uint64 `json:"amount,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transaction) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case transaction.FieldID, transaction.FieldAmount:
			values[i] = new(sql.NullInt64)
		case transaction.FieldTxID, transaction.FieldProvider, transaction.FieldRequest, transaction.FieldResponse, transaction.FieldCallbackResponse, transaction.FieldCode, transaction.FieldSerial, transaction.FieldStatus:
			values[i] = new(sql.NullString)
		case transaction.FieldCreatedAt, transaction.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Transaction", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transaction fields.
func (t *Transaction) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transaction.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int64(value.Int64)
		case transaction.FieldTxID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tx_id", values[i])
			} else if value.Valid {
				t.TxID = value.String
			}
		case transaction.FieldProvider:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider", values[i])
			} else if value.Valid {
				t.Provider = value.String
			}
		case transaction.FieldRequest:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field request", values[i])
			} else if value.Valid {
				t.Request = value.String
			}
		case transaction.FieldResponse:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field response", values[i])
			} else if value.Valid {
				t.Response = value.String
			}
		case transaction.FieldCallbackResponse:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field callback_response", values[i])
			} else if value.Valid {
				t.CallbackResponse = value.String
			}
		case transaction.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				t.Code = value.String
			}
		case transaction.FieldSerial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial", values[i])
			} else if value.Valid {
				t.Serial = value.String
			}
		case transaction.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				t.Amount = uint64(value.Int64)
			}
		case transaction.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = value.String
			}
		case transaction.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case transaction.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Transaction.
// Note that you need to call Transaction.Unwrap() before calling this method if this Transaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transaction) Update() *TransactionUpdateOne {
	return (&TransactionClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Transaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transaction) Unwrap() *Transaction {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transaction is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transaction) String() string {
	var builder strings.Builder
	builder.WriteString("Transaction(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", tx_id=")
	builder.WriteString(t.TxID)
	builder.WriteString(", provider=")
	builder.WriteString(t.Provider)
	builder.WriteString(", request=")
	builder.WriteString(t.Request)
	builder.WriteString(", response=")
	builder.WriteString(t.Response)
	builder.WriteString(", callback_response=")
	builder.WriteString(t.CallbackResponse)
	builder.WriteString(", code=")
	builder.WriteString(t.Code)
	builder.WriteString(", serial=")
	builder.WriteString(t.Serial)
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", t.Amount))
	builder.WriteString(", status=")
	builder.WriteString(t.Status)
	builder.WriteString(", created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Transactions is a parsable slice of Transaction.
type Transactions []*Transaction

func (t Transactions) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
