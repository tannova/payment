// Code generated by entc, DO NOT EDIT.

package merchantuserbankaccount

import (
	"time"
)

const (
	// Label holds the string label denoting the merchantuserbankaccount type in the database.
	Label = "merchant_user_bank_account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldBankName holds the string denoting the bank_name field in the database.
	FieldBankName = "bank_name"
	// FieldAccountNumber holds the string denoting the account_number field in the database.
	FieldAccountNumber = "account_number"
	// FieldAccountName holds the string denoting the account_name field in the database.
	FieldAccountName = "account_name"
	// Table holds the table name of the merchantuserbankaccount in the database.
	Table = "merchant_user_bank_accounts"
)

// Columns holds all SQL columns for merchantuserbankaccount fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldBankName,
	FieldAccountNumber,
	FieldAccountName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	CreatedByValidator func(string) error
	// UpdatedByValidator is a validator for the "updated_by" field. It is called by the builders before save.
	UpdatedByValidator func(string) error
	// DefaultBankName holds the default value on creation for the "bank_name" field.
	DefaultBankName int32
	// BankNameValidator is a validator for the "bank_name" field. It is called by the builders before save.
	BankNameValidator func(int32) error
	// AccountNumberValidator is a validator for the "account_number" field. It is called by the builders before save.
	AccountNumberValidator func(string) error
	// AccountNameValidator is a validator for the "account_name" field. It is called by the builders before save.
	AccountNameValidator func(string) error
)
