// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systemewallet"
)

// SystemEWallet is the model entity for the SystemEWallet schema.
type SystemEWallet struct {
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
	// EWalletName holds the value of the "e_wallet_name" field.
	EWalletName int32 `json:"e_wallet_name,omitempty"`
	// Status holds the value of the "status" field.
	Status int32 `json:"status,omitempty"`
	// MerchantID holds the value of the "merchant_id" field.
	MerchantID int64 `json:"merchant_id,omitempty"`
	// AccountPhoneNumber holds the value of the "account_phone_number" field.
	AccountPhoneNumber string `json:"account_phone_number,omitempty"`
	// AccountName holds the value of the "account_name" field.
	AccountName string `json:"account_name,omitempty"`
	// Balance holds the value of the "balance" field.
	Balance uint64 `json:"balance,omitempty"`
	// DailyBalance holds the value of the "daily_balance" field.
	DailyBalance uint64 `json:"daily_balance,omitempty"`
	// DailyBalanceLimit holds the value of the "daily_balance_limit" field.
	DailyBalanceLimit uint64 `json:"daily_balance_limit,omitempty"`
	// DailyUsedAmount holds the value of the "daily_used_amount" field.
	DailyUsedAmount int64 `json:"daily_used_amount,omitempty"`
	// LastUpdatedBalance holds the value of the "last_updated_balance" field.
	LastUpdatedBalance *time.Time `json:"last_updated_balance,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SystemEWallet) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case systemewallet.FieldID, systemewallet.FieldEWalletName, systemewallet.FieldStatus, systemewallet.FieldMerchantID, systemewallet.FieldBalance, systemewallet.FieldDailyBalance, systemewallet.FieldDailyBalanceLimit, systemewallet.FieldDailyUsedAmount:
			values[i] = new(sql.NullInt64)
		case systemewallet.FieldCreatedBy, systemewallet.FieldUpdatedBy, systemewallet.FieldAccountPhoneNumber, systemewallet.FieldAccountName:
			values[i] = new(sql.NullString)
		case systemewallet.FieldCreatedAt, systemewallet.FieldUpdatedAt, systemewallet.FieldLastUpdatedBalance:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SystemEWallet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SystemEWallet fields.
func (se *SystemEWallet) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case systemewallet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			se.ID = int64(value.Int64)
		case systemewallet.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				se.CreatedAt = value.Time
			}
		case systemewallet.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				se.UpdatedAt = value.Time
			}
		case systemewallet.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				se.CreatedBy = value.String
			}
		case systemewallet.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				se.UpdatedBy = value.String
			}
		case systemewallet.FieldEWalletName:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field e_wallet_name", values[i])
			} else if value.Valid {
				se.EWalletName = int32(value.Int64)
			}
		case systemewallet.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				se.Status = int32(value.Int64)
			}
		case systemewallet.FieldMerchantID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field merchant_id", values[i])
			} else if value.Valid {
				se.MerchantID = value.Int64
			}
		case systemewallet.FieldAccountPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account_phone_number", values[i])
			} else if value.Valid {
				se.AccountPhoneNumber = value.String
			}
		case systemewallet.FieldAccountName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account_name", values[i])
			} else if value.Valid {
				se.AccountName = value.String
			}
		case systemewallet.FieldBalance:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field balance", values[i])
			} else if value.Valid {
				se.Balance = uint64(value.Int64)
			}
		case systemewallet.FieldDailyBalance:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field daily_balance", values[i])
			} else if value.Valid {
				se.DailyBalance = uint64(value.Int64)
			}
		case systemewallet.FieldDailyBalanceLimit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field daily_balance_limit", values[i])
			} else if value.Valid {
				se.DailyBalanceLimit = uint64(value.Int64)
			}
		case systemewallet.FieldDailyUsedAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field daily_used_amount", values[i])
			} else if value.Valid {
				se.DailyUsedAmount = value.Int64
			}
		case systemewallet.FieldLastUpdatedBalance:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_updated_balance", values[i])
			} else if value.Valid {
				se.LastUpdatedBalance = new(time.Time)
				*se.LastUpdatedBalance = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SystemEWallet.
// Note that you need to call SystemEWallet.Unwrap() before calling this method if this SystemEWallet
// was returned from a transaction, and the transaction was committed or rolled back.
func (se *SystemEWallet) Update() *SystemEWalletUpdateOne {
	return (&SystemEWalletClient{config: se.config}).UpdateOne(se)
}

// Unwrap unwraps the SystemEWallet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (se *SystemEWallet) Unwrap() *SystemEWallet {
	tx, ok := se.config.driver.(*txDriver)
	if !ok {
		panic("ent: SystemEWallet is not a transactional entity")
	}
	se.config.driver = tx.drv
	return se
}

// String implements the fmt.Stringer.
func (se *SystemEWallet) String() string {
	var builder strings.Builder
	builder.WriteString("SystemEWallet(")
	builder.WriteString(fmt.Sprintf("id=%v", se.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(se.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(se.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", created_by=")
	builder.WriteString(se.CreatedBy)
	builder.WriteString(", updated_by=")
	builder.WriteString(se.UpdatedBy)
	builder.WriteString(", e_wallet_name=")
	builder.WriteString(fmt.Sprintf("%v", se.EWalletName))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", se.Status))
	builder.WriteString(", merchant_id=")
	builder.WriteString(fmt.Sprintf("%v", se.MerchantID))
	builder.WriteString(", account_phone_number=")
	builder.WriteString(se.AccountPhoneNumber)
	builder.WriteString(", account_name=")
	builder.WriteString(se.AccountName)
	builder.WriteString(", balance=")
	builder.WriteString(fmt.Sprintf("%v", se.Balance))
	builder.WriteString(", daily_balance=")
	builder.WriteString(fmt.Sprintf("%v", se.DailyBalance))
	builder.WriteString(", daily_balance_limit=")
	builder.WriteString(fmt.Sprintf("%v", se.DailyBalanceLimit))
	builder.WriteString(", daily_used_amount=")
	builder.WriteString(fmt.Sprintf("%v", se.DailyUsedAmount))
	if v := se.LastUpdatedBalance; v != nil {
		builder.WriteString(", last_updated_balance=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// SystemEWallets is a parsable slice of SystemEWallet.
type SystemEWallets []*SystemEWallet

func (se SystemEWallets) config(cfg config) {
	for _i := range se {
		se[_i].config = cfg
	}
}
