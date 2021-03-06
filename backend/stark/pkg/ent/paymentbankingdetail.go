// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentbankingdetail"
)

// PaymentBankingDetail is the model entity for the PaymentBankingDetail schema.
type PaymentBankingDetail struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// PaymentCode holds the value of the "payment_code" field.
	PaymentCode string `json:"payment_code,omitempty"`
	// MerchantUserID holds the value of the "merchant_user_id" field.
	MerchantUserID int64 `json:"merchant_user_id,omitempty"`
	// MerchantUserBankName holds the value of the "merchant_user_bank_name" field.
	MerchantUserBankName int32 `json:"merchant_user_bank_name,omitempty"`
	// MerchantUserAccountNumber holds the value of the "merchant_user_account_number" field.
	MerchantUserAccountNumber string `json:"merchant_user_account_number,omitempty"`
	// MerchantUserAccountName holds the value of the "merchant_user_account_name" field.
	MerchantUserAccountName string `json:"merchant_user_account_name,omitempty"`
	// SystemAccountBankName holds the value of the "system_account_bank_name" field.
	SystemAccountBankName int32 `json:"system_account_bank_name,omitempty"`
	// SystemAccountNumber holds the value of the "system_account_number" field.
	SystemAccountNumber string `json:"system_account_number,omitempty"`
	// SystemAccountName holds the value of the "system_account_name" field.
	SystemAccountName string `json:"system_account_name,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
	// TxID holds the value of the "tx_id" field.
	TxID string `json:"tx_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount uint64 `json:"amount,omitempty"`
	// Fee holds the value of the "fee" field.
	Fee uint64 `json:"fee,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PaymentBankingDetailQuery when eager-loading is set.
	Edges                          PaymentBankingDetailEdges `json:"edges"`
	payment_payment_banking_detail *int64
}

// PaymentBankingDetailEdges holds the relations/edges for other nodes in the graph.
type PaymentBankingDetailEdges struct {
	// Payment holds the value of the payment edge.
	Payment *Payment `json:"payment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PaymentOrErr returns the Payment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentBankingDetailEdges) PaymentOrErr() (*Payment, error) {
	if e.loadedTypes[0] {
		if e.Payment == nil {
			// The edge payment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: payment.Label}
		}
		return e.Payment, nil
	}
	return nil, &NotLoadedError{edge: "payment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PaymentBankingDetail) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case paymentbankingdetail.FieldID, paymentbankingdetail.FieldMerchantUserID, paymentbankingdetail.FieldMerchantUserBankName, paymentbankingdetail.FieldSystemAccountBankName, paymentbankingdetail.FieldAmount, paymentbankingdetail.FieldFee:
			values[i] = new(sql.NullInt64)
		case paymentbankingdetail.FieldCreatedBy, paymentbankingdetail.FieldUpdatedBy, paymentbankingdetail.FieldPaymentCode, paymentbankingdetail.FieldMerchantUserAccountNumber, paymentbankingdetail.FieldMerchantUserAccountName, paymentbankingdetail.FieldSystemAccountNumber, paymentbankingdetail.FieldSystemAccountName, paymentbankingdetail.FieldImageURL, paymentbankingdetail.FieldTxID:
			values[i] = new(sql.NullString)
		case paymentbankingdetail.FieldCreatedAt, paymentbankingdetail.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case paymentbankingdetail.ForeignKeys[0]: // payment_payment_banking_detail
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PaymentBankingDetail", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PaymentBankingDetail fields.
func (pbd *PaymentBankingDetail) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case paymentbankingdetail.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pbd.ID = int64(value.Int64)
		case paymentbankingdetail.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pbd.CreatedAt = value.Time
			}
		case paymentbankingdetail.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pbd.UpdatedAt = value.Time
			}
		case paymentbankingdetail.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pbd.CreatedBy = value.String
			}
		case paymentbankingdetail.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pbd.UpdatedBy = value.String
			}
		case paymentbankingdetail.FieldPaymentCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payment_code", values[i])
			} else if value.Valid {
				pbd.PaymentCode = value.String
			}
		case paymentbankingdetail.FieldMerchantUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field merchant_user_id", values[i])
			} else if value.Valid {
				pbd.MerchantUserID = value.Int64
			}
		case paymentbankingdetail.FieldMerchantUserBankName:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field merchant_user_bank_name", values[i])
			} else if value.Valid {
				pbd.MerchantUserBankName = int32(value.Int64)
			}
		case paymentbankingdetail.FieldMerchantUserAccountNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field merchant_user_account_number", values[i])
			} else if value.Valid {
				pbd.MerchantUserAccountNumber = value.String
			}
		case paymentbankingdetail.FieldMerchantUserAccountName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field merchant_user_account_name", values[i])
			} else if value.Valid {
				pbd.MerchantUserAccountName = value.String
			}
		case paymentbankingdetail.FieldSystemAccountBankName:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field system_account_bank_name", values[i])
			} else if value.Valid {
				pbd.SystemAccountBankName = int32(value.Int64)
			}
		case paymentbankingdetail.FieldSystemAccountNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field system_account_number", values[i])
			} else if value.Valid {
				pbd.SystemAccountNumber = value.String
			}
		case paymentbankingdetail.FieldSystemAccountName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field system_account_name", values[i])
			} else if value.Valid {
				pbd.SystemAccountName = value.String
			}
		case paymentbankingdetail.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				pbd.ImageURL = value.String
			}
		case paymentbankingdetail.FieldTxID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tx_id", values[i])
			} else if value.Valid {
				pbd.TxID = value.String
			}
		case paymentbankingdetail.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				pbd.Amount = uint64(value.Int64)
			}
		case paymentbankingdetail.FieldFee:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field fee", values[i])
			} else if value.Valid {
				pbd.Fee = uint64(value.Int64)
			}
		case paymentbankingdetail.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field payment_payment_banking_detail", value)
			} else if value.Valid {
				pbd.payment_payment_banking_detail = new(int64)
				*pbd.payment_payment_banking_detail = int64(value.Int64)
			}
		}
	}
	return nil
}

// QueryPayment queries the "payment" edge of the PaymentBankingDetail entity.
func (pbd *PaymentBankingDetail) QueryPayment() *PaymentQuery {
	return (&PaymentBankingDetailClient{config: pbd.config}).QueryPayment(pbd)
}

// Update returns a builder for updating this PaymentBankingDetail.
// Note that you need to call PaymentBankingDetail.Unwrap() before calling this method if this PaymentBankingDetail
// was returned from a transaction, and the transaction was committed or rolled back.
func (pbd *PaymentBankingDetail) Update() *PaymentBankingDetailUpdateOne {
	return (&PaymentBankingDetailClient{config: pbd.config}).UpdateOne(pbd)
}

// Unwrap unwraps the PaymentBankingDetail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pbd *PaymentBankingDetail) Unwrap() *PaymentBankingDetail {
	tx, ok := pbd.config.driver.(*txDriver)
	if !ok {
		panic("ent: PaymentBankingDetail is not a transactional entity")
	}
	pbd.config.driver = tx.drv
	return pbd
}

// String implements the fmt.Stringer.
func (pbd *PaymentBankingDetail) String() string {
	var builder strings.Builder
	builder.WriteString("PaymentBankingDetail(")
	builder.WriteString(fmt.Sprintf("id=%v", pbd.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(pbd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pbd.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", created_by=")
	builder.WriteString(pbd.CreatedBy)
	builder.WriteString(", updated_by=")
	builder.WriteString(pbd.UpdatedBy)
	builder.WriteString(", payment_code=")
	builder.WriteString(pbd.PaymentCode)
	builder.WriteString(", merchant_user_id=")
	builder.WriteString(fmt.Sprintf("%v", pbd.MerchantUserID))
	builder.WriteString(", merchant_user_bank_name=")
	builder.WriteString(fmt.Sprintf("%v", pbd.MerchantUserBankName))
	builder.WriteString(", merchant_user_account_number=")
	builder.WriteString(pbd.MerchantUserAccountNumber)
	builder.WriteString(", merchant_user_account_name=")
	builder.WriteString(pbd.MerchantUserAccountName)
	builder.WriteString(", system_account_bank_name=")
	builder.WriteString(fmt.Sprintf("%v", pbd.SystemAccountBankName))
	builder.WriteString(", system_account_number=")
	builder.WriteString(pbd.SystemAccountNumber)
	builder.WriteString(", system_account_name=")
	builder.WriteString(pbd.SystemAccountName)
	builder.WriteString(", image_url=")
	builder.WriteString(pbd.ImageURL)
	builder.WriteString(", tx_id=")
	builder.WriteString(pbd.TxID)
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", pbd.Amount))
	builder.WriteString(", fee=")
	builder.WriteString(fmt.Sprintf("%v", pbd.Fee))
	builder.WriteByte(')')
	return builder.String()
}

// PaymentBankingDetails is a parsable slice of PaymentBankingDetail.
type PaymentBankingDetails []*PaymentBankingDetail

func (pbd PaymentBankingDetails) config(cfg config) {
	for _i := range pbd {
		pbd[_i].config = cfg
	}
}
