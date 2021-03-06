// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentcryptodetail"
)

// PaymentCryptoDetail is the model entity for the PaymentCryptoDetail schema.
type PaymentCryptoDetail struct {
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
	// CryptoType holds the value of the "crypto_type" field.
	CryptoType int32 `json:"crypto_type,omitempty"`
	// CryptoNetworkType holds the value of the "crypto_network_type" field.
	CryptoNetworkType int32 `json:"crypto_network_type,omitempty"`
	// CryptoWalletName holds the value of the "crypto_wallet_name" field.
	CryptoWalletName int32 `json:"crypto_wallet_name,omitempty"`
	// ReceiverAddress holds the value of the "receiver_address" field.
	ReceiverAddress string `json:"receiver_address,omitempty"`
	// SenderAddress holds the value of the "sender_address" field.
	SenderAddress string `json:"sender_address,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount float64 `json:"amount,omitempty"`
	// ReceivedAmount holds the value of the "received_amount" field.
	ReceivedAmount float64 `json:"received_amount,omitempty"`
	// TxHash holds the value of the "tx_hash" field.
	TxHash string `json:"tx_hash,omitempty"`
	// TxID holds the value of the "tx_id" field.
	TxID string `json:"tx_id,omitempty"`
	// Fee holds the value of the "fee" field.
	Fee float64 `json:"fee,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PaymentCryptoDetailQuery when eager-loading is set.
	Edges                         PaymentCryptoDetailEdges `json:"edges"`
	payment_payment_crypto_detail *int64
}

// PaymentCryptoDetailEdges holds the relations/edges for other nodes in the graph.
type PaymentCryptoDetailEdges struct {
	// Payment holds the value of the payment edge.
	Payment *Payment `json:"payment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PaymentOrErr returns the Payment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentCryptoDetailEdges) PaymentOrErr() (*Payment, error) {
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
func (*PaymentCryptoDetail) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case paymentcryptodetail.FieldAmount, paymentcryptodetail.FieldReceivedAmount, paymentcryptodetail.FieldFee:
			values[i] = new(sql.NullFloat64)
		case paymentcryptodetail.FieldID, paymentcryptodetail.FieldCryptoType, paymentcryptodetail.FieldCryptoNetworkType, paymentcryptodetail.FieldCryptoWalletName:
			values[i] = new(sql.NullInt64)
		case paymentcryptodetail.FieldCreatedBy, paymentcryptodetail.FieldUpdatedBy, paymentcryptodetail.FieldReceiverAddress, paymentcryptodetail.FieldSenderAddress, paymentcryptodetail.FieldTxHash, paymentcryptodetail.FieldTxID, paymentcryptodetail.FieldImageURL:
			values[i] = new(sql.NullString)
		case paymentcryptodetail.FieldCreatedAt, paymentcryptodetail.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case paymentcryptodetail.ForeignKeys[0]: // payment_payment_crypto_detail
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PaymentCryptoDetail", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PaymentCryptoDetail fields.
func (pcd *PaymentCryptoDetail) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case paymentcryptodetail.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pcd.ID = int64(value.Int64)
		case paymentcryptodetail.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pcd.CreatedAt = value.Time
			}
		case paymentcryptodetail.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pcd.UpdatedAt = value.Time
			}
		case paymentcryptodetail.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pcd.CreatedBy = value.String
			}
		case paymentcryptodetail.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pcd.UpdatedBy = value.String
			}
		case paymentcryptodetail.FieldCryptoType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field crypto_type", values[i])
			} else if value.Valid {
				pcd.CryptoType = int32(value.Int64)
			}
		case paymentcryptodetail.FieldCryptoNetworkType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field crypto_network_type", values[i])
			} else if value.Valid {
				pcd.CryptoNetworkType = int32(value.Int64)
			}
		case paymentcryptodetail.FieldCryptoWalletName:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field crypto_wallet_name", values[i])
			} else if value.Valid {
				pcd.CryptoWalletName = int32(value.Int64)
			}
		case paymentcryptodetail.FieldReceiverAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field receiver_address", values[i])
			} else if value.Valid {
				pcd.ReceiverAddress = value.String
			}
		case paymentcryptodetail.FieldSenderAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sender_address", values[i])
			} else if value.Valid {
				pcd.SenderAddress = value.String
			}
		case paymentcryptodetail.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				pcd.Amount = value.Float64
			}
		case paymentcryptodetail.FieldReceivedAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field received_amount", values[i])
			} else if value.Valid {
				pcd.ReceivedAmount = value.Float64
			}
		case paymentcryptodetail.FieldTxHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tx_hash", values[i])
			} else if value.Valid {
				pcd.TxHash = value.String
			}
		case paymentcryptodetail.FieldTxID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tx_id", values[i])
			} else if value.Valid {
				pcd.TxID = value.String
			}
		case paymentcryptodetail.FieldFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field fee", values[i])
			} else if value.Valid {
				pcd.Fee = value.Float64
			}
		case paymentcryptodetail.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				pcd.ImageURL = value.String
			}
		case paymentcryptodetail.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field payment_payment_crypto_detail", value)
			} else if value.Valid {
				pcd.payment_payment_crypto_detail = new(int64)
				*pcd.payment_payment_crypto_detail = int64(value.Int64)
			}
		}
	}
	return nil
}

// QueryPayment queries the "payment" edge of the PaymentCryptoDetail entity.
func (pcd *PaymentCryptoDetail) QueryPayment() *PaymentQuery {
	return (&PaymentCryptoDetailClient{config: pcd.config}).QueryPayment(pcd)
}

// Update returns a builder for updating this PaymentCryptoDetail.
// Note that you need to call PaymentCryptoDetail.Unwrap() before calling this method if this PaymentCryptoDetail
// was returned from a transaction, and the transaction was committed or rolled back.
func (pcd *PaymentCryptoDetail) Update() *PaymentCryptoDetailUpdateOne {
	return (&PaymentCryptoDetailClient{config: pcd.config}).UpdateOne(pcd)
}

// Unwrap unwraps the PaymentCryptoDetail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pcd *PaymentCryptoDetail) Unwrap() *PaymentCryptoDetail {
	tx, ok := pcd.config.driver.(*txDriver)
	if !ok {
		panic("ent: PaymentCryptoDetail is not a transactional entity")
	}
	pcd.config.driver = tx.drv
	return pcd
}

// String implements the fmt.Stringer.
func (pcd *PaymentCryptoDetail) String() string {
	var builder strings.Builder
	builder.WriteString("PaymentCryptoDetail(")
	builder.WriteString(fmt.Sprintf("id=%v", pcd.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(pcd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pcd.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", created_by=")
	builder.WriteString(pcd.CreatedBy)
	builder.WriteString(", updated_by=")
	builder.WriteString(pcd.UpdatedBy)
	builder.WriteString(", crypto_type=")
	builder.WriteString(fmt.Sprintf("%v", pcd.CryptoType))
	builder.WriteString(", crypto_network_type=")
	builder.WriteString(fmt.Sprintf("%v", pcd.CryptoNetworkType))
	builder.WriteString(", crypto_wallet_name=")
	builder.WriteString(fmt.Sprintf("%v", pcd.CryptoWalletName))
	builder.WriteString(", receiver_address=")
	builder.WriteString(pcd.ReceiverAddress)
	builder.WriteString(", sender_address=")
	builder.WriteString(pcd.SenderAddress)
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", pcd.Amount))
	builder.WriteString(", received_amount=")
	builder.WriteString(fmt.Sprintf("%v", pcd.ReceivedAmount))
	builder.WriteString(", tx_hash=")
	builder.WriteString(pcd.TxHash)
	builder.WriteString(", tx_id=")
	builder.WriteString(pcd.TxID)
	builder.WriteString(", fee=")
	builder.WriteString(fmt.Sprintf("%v", pcd.Fee))
	builder.WriteString(", image_url=")
	builder.WriteString(pcd.ImageURL)
	builder.WriteByte(')')
	return builder.String()
}

// PaymentCryptoDetails is a parsable slice of PaymentCryptoDetail.
type PaymentCryptoDetails []*PaymentCryptoDetail

func (pcd PaymentCryptoDetails) config(cfg config) {
	for _i := range pcd {
		pcd[_i].config = cfg
	}
}
