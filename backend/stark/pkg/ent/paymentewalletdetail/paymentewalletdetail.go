// Code generated by entc, DO NOT EDIT.

package paymentewalletdetail

import (
	"time"
)

const (
	// Label holds the string label denoting the paymentewalletdetail type in the database.
	Label = "payment_ewallet_detail"
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
	// FieldPaymentCode holds the string denoting the payment_code field in the database.
	FieldPaymentCode = "payment_code"
	// FieldMerchantUserID holds the string denoting the merchant_user_id field in the database.
	FieldMerchantUserID = "merchant_user_id"
	// FieldEWalletName holds the string denoting the e_wallet_name field in the database.
	FieldEWalletName = "e_wallet_name"
	// FieldMerchantUserAccountPhoneNumber holds the string denoting the merchant_user_account_phone_number field in the database.
	FieldMerchantUserAccountPhoneNumber = "merchant_user_account_phone_number"
	// FieldMerchantUserAccountName holds the string denoting the merchant_user_account_name field in the database.
	FieldMerchantUserAccountName = "merchant_user_account_name"
	// FieldSystemAccountPhoneNumber holds the string denoting the system_account_phone_number field in the database.
	FieldSystemAccountPhoneNumber = "system_account_phone_number"
	// FieldSystemAccountName holds the string denoting the system_account_name field in the database.
	FieldSystemAccountName = "system_account_name"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldFee holds the string denoting the fee field in the database.
	FieldFee = "fee"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// FieldTxID holds the string denoting the tx_id field in the database.
	FieldTxID = "tx_id"
	// EdgePayment holds the string denoting the payment edge name in mutations.
	EdgePayment = "payment"
	// Table holds the table name of the paymentewalletdetail in the database.
	Table = "payment_ewallet_details"
	// PaymentTable is the table that holds the payment relation/edge.
	PaymentTable = "payment_ewallet_details"
	// PaymentInverseTable is the table name for the Payment entity.
	// It exists in this package in order to avoid circular dependency with the "payment" package.
	PaymentInverseTable = "payments"
	// PaymentColumn is the table column denoting the payment relation/edge.
	PaymentColumn = "payment_payment_e_wallet_detail"
)

// Columns holds all SQL columns for paymentewalletdetail fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldPaymentCode,
	FieldMerchantUserID,
	FieldEWalletName,
	FieldMerchantUserAccountPhoneNumber,
	FieldMerchantUserAccountName,
	FieldSystemAccountPhoneNumber,
	FieldSystemAccountName,
	FieldAmount,
	FieldFee,
	FieldNote,
	FieldImageURL,
	FieldTxID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "payment_ewallet_details"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"payment_payment_e_wallet_detail",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// EWalletNameValidator is a validator for the "e_wallet_name" field. It is called by the builders before save.
	EWalletNameValidator func(int32) error
)
