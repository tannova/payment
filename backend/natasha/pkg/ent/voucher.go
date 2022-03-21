// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/voucher"
)

// Voucher is the model entity for the Voucher schema.
type Voucher struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// MerchantID holds the value of the "merchant_id" field.
	MerchantID int64 `json:"merchant_id,omitempty"`
	// PaymentID holds the value of the "payment_id" field.
	PaymentID int64 `json:"payment_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount int64 `json:"amount,omitempty"`
	// Type holds the value of the "type" field.
	Type int32 `json:"type,omitempty"`
	// Status holds the value of the "status" field.
	Status int32 `json:"status,omitempty"`
	// PayeeProvider holds the value of the "payee_provider" field.
	PayeeProvider int32 `json:"payee_provider,omitempty"`
	// PayeeAccount holds the value of the "payee_account" field.
	PayeeAccount string `json:"payee_account,omitempty"`
	// PayeeName holds the value of the "payee_name" field.
	PayeeName string `json:"payee_name,omitempty"`
	// PayerProvider holds the value of the "payer_provider" field.
	PayerProvider int32 `json:"payer_provider,omitempty"`
	// PayerAccount holds the value of the "payer_account" field.
	PayerAccount string `json:"payer_account,omitempty"`
	// PayerName holds the value of the "payer_name" field.
	PayerName string `json:"payer_name,omitempty"`
	// TxID holds the value of the "tx_id" field.
	TxID string `json:"tx_id,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
	// CreatorNote holds the value of the "creator_note" field.
	CreatorNote string `json:"creator_note,omitempty"`
	// HandledBy holds the value of the "handled_by" field.
	HandledBy string `json:"handled_by,omitempty"`
	// HandlerNote holds the value of the "handler_note" field.
	HandlerNote string `json:"handler_note,omitempty"`
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
func (*Voucher) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case voucher.FieldID, voucher.FieldMerchantID, voucher.FieldPaymentID, voucher.FieldAmount, voucher.FieldType, voucher.FieldStatus, voucher.FieldPayeeProvider, voucher.FieldPayerProvider:
			values[i] = &sql.NullInt64{}
		case voucher.FieldPayeeAccount, voucher.FieldPayeeName, voucher.FieldPayerAccount, voucher.FieldPayerName, voucher.FieldTxID, voucher.FieldImageURL, voucher.FieldCreatorNote, voucher.FieldHandledBy, voucher.FieldHandlerNote, voucher.FieldCreatedBy, voucher.FieldUpdatedBy:
			values[i] = &sql.NullString{}
		case voucher.FieldCreatedAt, voucher.FieldUpdatedAt:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Voucher", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Voucher fields.
func (v *Voucher) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case voucher.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int64(value.Int64)
		case voucher.FieldMerchantID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field merchant_id", values[i])
			} else if value.Valid {
				v.MerchantID = value.Int64
			}
		case voucher.FieldPaymentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field payment_id", values[i])
			} else if value.Valid {
				v.PaymentID = value.Int64
			}
		case voucher.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				v.Amount = value.Int64
			}
		case voucher.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				v.Type = int32(value.Int64)
			}
		case voucher.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				v.Status = int32(value.Int64)
			}
		case voucher.FieldPayeeProvider:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field payee_provider", values[i])
			} else if value.Valid {
				v.PayeeProvider = int32(value.Int64)
			}
		case voucher.FieldPayeeAccount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payee_account", values[i])
			} else if value.Valid {
				v.PayeeAccount = value.String
			}
		case voucher.FieldPayeeName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payee_name", values[i])
			} else if value.Valid {
				v.PayeeName = value.String
			}
		case voucher.FieldPayerProvider:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field payer_provider", values[i])
			} else if value.Valid {
				v.PayerProvider = int32(value.Int64)
			}
		case voucher.FieldPayerAccount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payer_account", values[i])
			} else if value.Valid {
				v.PayerAccount = value.String
			}
		case voucher.FieldPayerName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payer_name", values[i])
			} else if value.Valid {
				v.PayerName = value.String
			}
		case voucher.FieldTxID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tx_id", values[i])
			} else if value.Valid {
				v.TxID = value.String
			}
		case voucher.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				v.ImageURL = value.String
			}
		case voucher.FieldCreatorNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator_note", values[i])
			} else if value.Valid {
				v.CreatorNote = value.String
			}
		case voucher.FieldHandledBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field handled_by", values[i])
			} else if value.Valid {
				v.HandledBy = value.String
			}
		case voucher.FieldHandlerNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field handler_note", values[i])
			} else if value.Valid {
				v.HandlerNote = value.String
			}
		case voucher.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				v.CreatedAt = value.Time
			}
		case voucher.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				v.UpdatedAt = value.Time
			}
		case voucher.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				v.CreatedBy = value.String
			}
		case voucher.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				v.UpdatedBy = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Voucher.
// Note that you need to call Voucher.Unwrap() before calling this method if this Voucher
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Voucher) Update() *VoucherUpdateOne {
	return (&VoucherClient{config: v.config}).UpdateOne(v)
}

// Unwrap unwraps the Voucher entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Voucher) Unwrap() *Voucher {
	tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Voucher is not a transactional entity")
	}
	v.config.driver = tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Voucher) String() string {
	var builder strings.Builder
	builder.WriteString("Voucher(")
	builder.WriteString(fmt.Sprintf("id=%v", v.ID))
	builder.WriteString(", merchant_id=")
	builder.WriteString(fmt.Sprintf("%v", v.MerchantID))
	builder.WriteString(", payment_id=")
	builder.WriteString(fmt.Sprintf("%v", v.PaymentID))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", v.Amount))
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", v.Type))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", v.Status))
	builder.WriteString(", payee_provider=")
	builder.WriteString(fmt.Sprintf("%v", v.PayeeProvider))
	builder.WriteString(", payee_account=")
	builder.WriteString(v.PayeeAccount)
	builder.WriteString(", payee_name=")
	builder.WriteString(v.PayeeName)
	builder.WriteString(", payer_provider=")
	builder.WriteString(fmt.Sprintf("%v", v.PayerProvider))
	builder.WriteString(", payer_account=")
	builder.WriteString(v.PayerAccount)
	builder.WriteString(", payer_name=")
	builder.WriteString(v.PayerName)
	builder.WriteString(", tx_id=")
	builder.WriteString(v.TxID)
	builder.WriteString(", image_url=")
	builder.WriteString(v.ImageURL)
	builder.WriteString(", creator_note=")
	builder.WriteString(v.CreatorNote)
	builder.WriteString(", handled_by=")
	builder.WriteString(v.HandledBy)
	builder.WriteString(", handler_note=")
	builder.WriteString(v.HandlerNote)
	builder.WriteString(", created_at=")
	builder.WriteString(v.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(v.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", created_by=")
	builder.WriteString(v.CreatedBy)
	builder.WriteString(", updated_by=")
	builder.WriteString(v.UpdatedBy)
	builder.WriteByte(')')
	return builder.String()
}

// Vouchers is a parsable slice of Voucher.
type Vouchers []*Voucher

func (v Vouchers) config(cfg config) {
	for _i := range v {
		v[_i].config = cfg
	}
}
