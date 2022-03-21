// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
)

// Payment is the model entity for the Payment schema.
type Payment struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ApprovedBy holds the value of the "approved_by" field.
	ApprovedBy *string `json:"approved_by,omitempty"`
	// ApprovedAt holds the value of the "approved_at" field.
	ApprovedAt *time.Time `json:"approved_at,omitempty"`
	// MerchantID holds the value of the "merchant_id" field.
	MerchantID int64 `json:"merchant_id,omitempty"`
	// MerchantUserID holds the value of the "merchant_user_id" field.
	MerchantUserID int64 `json:"merchant_user_id,omitempty"`
	// Method holds the value of the "method" field.
	Method int32 `json:"method,omitempty"`
	// Type holds the value of the "type" field.
	Type int32 `json:"type,omitempty"`
	// Status holds the value of the "status" field.
	Status int32 `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PaymentQuery when eager-loading is set.
	Edges PaymentEdges `json:"edges"`
}

// PaymentEdges holds the relations/edges for other nodes in the graph.
type PaymentEdges struct {
	// PaymentBankingDetail holds the value of the payment_banking_detail edge.
	PaymentBankingDetail []*PaymentBankingDetail `json:"payment_banking_detail,omitempty"`
	// PaymentTelcoDetail holds the value of the payment_telco_detail edge.
	PaymentTelcoDetail []*PaymentTelcoDetail `json:"payment_telco_detail,omitempty"`
	// PaymentEWalletDetail holds the value of the payment_e_wallet_detail edge.
	PaymentEWalletDetail []*PaymentEWalletDetail `json:"payment_e_wallet_detail,omitempty"`
	// PaymentCryptoDetail holds the value of the payment_crypto_detail edge.
	PaymentCryptoDetail []*PaymentCryptoDetail `json:"payment_crypto_detail,omitempty"`
	// Revisions holds the value of the revisions edge.
	Revisions []*Revision `json:"revisions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// PaymentBankingDetailOrErr returns the PaymentBankingDetail value or an error if the edge
// was not loaded in eager-loading.
func (e PaymentEdges) PaymentBankingDetailOrErr() ([]*PaymentBankingDetail, error) {
	if e.loadedTypes[0] {
		return e.PaymentBankingDetail, nil
	}
	return nil, &NotLoadedError{edge: "payment_banking_detail"}
}

// PaymentTelcoDetailOrErr returns the PaymentTelcoDetail value or an error if the edge
// was not loaded in eager-loading.
func (e PaymentEdges) PaymentTelcoDetailOrErr() ([]*PaymentTelcoDetail, error) {
	if e.loadedTypes[1] {
		return e.PaymentTelcoDetail, nil
	}
	return nil, &NotLoadedError{edge: "payment_telco_detail"}
}

// PaymentEWalletDetailOrErr returns the PaymentEWalletDetail value or an error if the edge
// was not loaded in eager-loading.
func (e PaymentEdges) PaymentEWalletDetailOrErr() ([]*PaymentEWalletDetail, error) {
	if e.loadedTypes[2] {
		return e.PaymentEWalletDetail, nil
	}
	return nil, &NotLoadedError{edge: "payment_e_wallet_detail"}
}

// PaymentCryptoDetailOrErr returns the PaymentCryptoDetail value or an error if the edge
// was not loaded in eager-loading.
func (e PaymentEdges) PaymentCryptoDetailOrErr() ([]*PaymentCryptoDetail, error) {
	if e.loadedTypes[3] {
		return e.PaymentCryptoDetail, nil
	}
	return nil, &NotLoadedError{edge: "payment_crypto_detail"}
}

// RevisionsOrErr returns the Revisions value or an error if the edge
// was not loaded in eager-loading.
func (e PaymentEdges) RevisionsOrErr() ([]*Revision, error) {
	if e.loadedTypes[4] {
		return e.Revisions, nil
	}
	return nil, &NotLoadedError{edge: "revisions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Payment) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case payment.FieldID, payment.FieldMerchantID, payment.FieldMerchantUserID, payment.FieldMethod, payment.FieldType, payment.FieldStatus:
			values[i] = new(sql.NullInt64)
		case payment.FieldCreatedBy, payment.FieldUpdatedBy, payment.FieldApprovedBy:
			values[i] = new(sql.NullString)
		case payment.FieldCreatedAt, payment.FieldUpdatedAt, payment.FieldApprovedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Payment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Payment fields.
func (pa *Payment) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case payment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int64(value.Int64)
		case payment.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pa.CreatedBy = value.String
			}
		case payment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pa.CreatedAt = value.Time
			}
		case payment.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pa.UpdatedBy = value.String
			}
		case payment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pa.UpdatedAt = value.Time
			}
		case payment.FieldApprovedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field approved_by", values[i])
			} else if value.Valid {
				pa.ApprovedBy = new(string)
				*pa.ApprovedBy = value.String
			}
		case payment.FieldApprovedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field approved_at", values[i])
			} else if value.Valid {
				pa.ApprovedAt = new(time.Time)
				*pa.ApprovedAt = value.Time
			}
		case payment.FieldMerchantID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field merchant_id", values[i])
			} else if value.Valid {
				pa.MerchantID = value.Int64
			}
		case payment.FieldMerchantUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field merchant_user_id", values[i])
			} else if value.Valid {
				pa.MerchantUserID = value.Int64
			}
		case payment.FieldMethod:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				pa.Method = int32(value.Int64)
			}
		case payment.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				pa.Type = int32(value.Int64)
			}
		case payment.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pa.Status = int32(value.Int64)
			}
		}
	}
	return nil
}

// QueryPaymentBankingDetail queries the "payment_banking_detail" edge of the Payment entity.
func (pa *Payment) QueryPaymentBankingDetail() *PaymentBankingDetailQuery {
	return (&PaymentClient{config: pa.config}).QueryPaymentBankingDetail(pa)
}

// QueryPaymentTelcoDetail queries the "payment_telco_detail" edge of the Payment entity.
func (pa *Payment) QueryPaymentTelcoDetail() *PaymentTelcoDetailQuery {
	return (&PaymentClient{config: pa.config}).QueryPaymentTelcoDetail(pa)
}

// QueryPaymentEWalletDetail queries the "payment_e_wallet_detail" edge of the Payment entity.
func (pa *Payment) QueryPaymentEWalletDetail() *PaymentEWalletDetailQuery {
	return (&PaymentClient{config: pa.config}).QueryPaymentEWalletDetail(pa)
}

// QueryPaymentCryptoDetail queries the "payment_crypto_detail" edge of the Payment entity.
func (pa *Payment) QueryPaymentCryptoDetail() *PaymentCryptoDetailQuery {
	return (&PaymentClient{config: pa.config}).QueryPaymentCryptoDetail(pa)
}

// QueryRevisions queries the "revisions" edge of the Payment entity.
func (pa *Payment) QueryRevisions() *RevisionQuery {
	return (&PaymentClient{config: pa.config}).QueryRevisions(pa)
}

// Update returns a builder for updating this Payment.
// Note that you need to call Payment.Unwrap() before calling this method if this Payment
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Payment) Update() *PaymentUpdateOne {
	return (&PaymentClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the Payment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Payment) Unwrap() *Payment {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Payment is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Payment) String() string {
	var builder strings.Builder
	builder.WriteString("Payment(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", created_by=")
	builder.WriteString(pa.CreatedBy)
	builder.WriteString(", created_at=")
	builder.WriteString(pa.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_by=")
	builder.WriteString(pa.UpdatedBy)
	builder.WriteString(", updated_at=")
	builder.WriteString(pa.UpdatedAt.Format(time.ANSIC))
	if v := pa.ApprovedBy; v != nil {
		builder.WriteString(", approved_by=")
		builder.WriteString(*v)
	}
	if v := pa.ApprovedAt; v != nil {
		builder.WriteString(", approved_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", merchant_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.MerchantID))
	builder.WriteString(", merchant_user_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.MerchantUserID))
	builder.WriteString(", method=")
	builder.WriteString(fmt.Sprintf("%v", pa.Method))
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", pa.Type))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", pa.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Payments is a parsable slice of Payment.
type Payments []*Payment

func (pa Payments) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}