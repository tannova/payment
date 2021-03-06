// Code generated by entc, DO NOT EDIT.

package voucher

import (
	"time"
)

const (
	// Label holds the string label denoting the voucher type in the database.
	Label = "voucher"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMerchantID holds the string denoting the merchant_id field in the database.
	FieldMerchantID = "merchant_id"
	// FieldPaymentID holds the string denoting the payment_id field in the database.
	FieldPaymentID = "payment_id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPayeeProvider holds the string denoting the payee_provider field in the database.
	FieldPayeeProvider = "payee_provider"
	// FieldPayeeAccount holds the string denoting the payee_account field in the database.
	FieldPayeeAccount = "payee_account"
	// FieldPayeeName holds the string denoting the payee_name field in the database.
	FieldPayeeName = "payee_name"
	// FieldPayerProvider holds the string denoting the payer_provider field in the database.
	FieldPayerProvider = "payer_provider"
	// FieldPayerAccount holds the string denoting the payer_account field in the database.
	FieldPayerAccount = "payer_account"
	// FieldPayerName holds the string denoting the payer_name field in the database.
	FieldPayerName = "payer_name"
	// FieldTxID holds the string denoting the tx_id field in the database.
	FieldTxID = "tx_id"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// FieldCreatorNote holds the string denoting the creator_note field in the database.
	FieldCreatorNote = "creator_note"
	// FieldHandledBy holds the string denoting the handled_by field in the database.
	FieldHandledBy = "handled_by"
	// FieldHandlerNote holds the string denoting the handler_note field in the database.
	FieldHandlerNote = "handler_note"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// Table holds the table name of the voucher in the database.
	Table = "vouchers"
)

// Columns holds all SQL columns for voucher fields.
var Columns = []string{
	FieldID,
	FieldMerchantID,
	FieldPaymentID,
	FieldAmount,
	FieldType,
	FieldStatus,
	FieldPayeeProvider,
	FieldPayeeAccount,
	FieldPayeeName,
	FieldPayerProvider,
	FieldPayerAccount,
	FieldPayerName,
	FieldTxID,
	FieldImageURL,
	FieldCreatorNote,
	FieldHandledBy,
	FieldHandlerNote,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
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
	// DefaultMerchantID holds the default value on creation for the "merchant_id" field.
	DefaultMerchantID int64
	// MerchantIDValidator is a validator for the "merchant_id" field. It is called by the builders before save.
	MerchantIDValidator func(int64) error
	// DefaultPaymentID holds the default value on creation for the "payment_id" field.
	DefaultPaymentID int64
	// PaymentIDValidator is a validator for the "payment_id" field. It is called by the builders before save.
	PaymentIDValidator func(int64) error
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType int32
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(int32) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int32
	// StatusValidator is a validator for the "status" field. It is called by the builders before save.
	StatusValidator func(int32) error
	// DefaultPayeeProvider holds the default value on creation for the "payee_provider" field.
	DefaultPayeeProvider int32
	// PayeeProviderValidator is a validator for the "payee_provider" field. It is called by the builders before save.
	PayeeProviderValidator func(int32) error
	// DefaultPayerProvider holds the default value on creation for the "payer_provider" field.
	DefaultPayerProvider int32
	// PayerProviderValidator is a validator for the "payer_provider" field. It is called by the builders before save.
	PayerProviderValidator func(int32) error
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
)
