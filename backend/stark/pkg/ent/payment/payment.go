// Code generated by entc, DO NOT EDIT.

package payment

import (
	"time"
)

const (
	// Label holds the string label denoting the payment type in the database.
	Label = "payment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldApprovedBy holds the string denoting the approved_by field in the database.
	FieldApprovedBy = "approved_by"
	// FieldApprovedAt holds the string denoting the approved_at field in the database.
	FieldApprovedAt = "approved_at"
	// FieldMerchantID holds the string denoting the merchant_id field in the database.
	FieldMerchantID = "merchant_id"
	// FieldMerchantUserID holds the string denoting the merchant_user_id field in the database.
	FieldMerchantUserID = "merchant_user_id"
	// FieldMethod holds the string denoting the method field in the database.
	FieldMethod = "method"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgePaymentBankingDetail holds the string denoting the payment_banking_detail edge name in mutations.
	EdgePaymentBankingDetail = "payment_banking_detail"
	// EdgePaymentTelcoDetail holds the string denoting the payment_telco_detail edge name in mutations.
	EdgePaymentTelcoDetail = "payment_telco_detail"
	// EdgePaymentEWalletDetail holds the string denoting the payment_e_wallet_detail edge name in mutations.
	EdgePaymentEWalletDetail = "payment_e_wallet_detail"
	// EdgePaymentCryptoDetail holds the string denoting the payment_crypto_detail edge name in mutations.
	EdgePaymentCryptoDetail = "payment_crypto_detail"
	// EdgeRevisions holds the string denoting the revisions edge name in mutations.
	EdgeRevisions = "revisions"
	// Table holds the table name of the payment in the database.
	Table = "payments"
	// PaymentBankingDetailTable is the table that holds the payment_banking_detail relation/edge.
	PaymentBankingDetailTable = "payment_banking_details"
	// PaymentBankingDetailInverseTable is the table name for the PaymentBankingDetail entity.
	// It exists in this package in order to avoid circular dependency with the "paymentbankingdetail" package.
	PaymentBankingDetailInverseTable = "payment_banking_details"
	// PaymentBankingDetailColumn is the table column denoting the payment_banking_detail relation/edge.
	PaymentBankingDetailColumn = "payment_payment_banking_detail"
	// PaymentTelcoDetailTable is the table that holds the payment_telco_detail relation/edge.
	PaymentTelcoDetailTable = "payment_telco_details"
	// PaymentTelcoDetailInverseTable is the table name for the PaymentTelcoDetail entity.
	// It exists in this package in order to avoid circular dependency with the "paymenttelcodetail" package.
	PaymentTelcoDetailInverseTable = "payment_telco_details"
	// PaymentTelcoDetailColumn is the table column denoting the payment_telco_detail relation/edge.
	PaymentTelcoDetailColumn = "payment_payment_telco_detail"
	// PaymentEWalletDetailTable is the table that holds the payment_e_wallet_detail relation/edge.
	PaymentEWalletDetailTable = "payment_ewallet_details"
	// PaymentEWalletDetailInverseTable is the table name for the PaymentEWalletDetail entity.
	// It exists in this package in order to avoid circular dependency with the "paymentewalletdetail" package.
	PaymentEWalletDetailInverseTable = "payment_ewallet_details"
	// PaymentEWalletDetailColumn is the table column denoting the payment_e_wallet_detail relation/edge.
	PaymentEWalletDetailColumn = "payment_payment_e_wallet_detail"
	// PaymentCryptoDetailTable is the table that holds the payment_crypto_detail relation/edge.
	PaymentCryptoDetailTable = "payment_crypto_details"
	// PaymentCryptoDetailInverseTable is the table name for the PaymentCryptoDetail entity.
	// It exists in this package in order to avoid circular dependency with the "paymentcryptodetail" package.
	PaymentCryptoDetailInverseTable = "payment_crypto_details"
	// PaymentCryptoDetailColumn is the table column denoting the payment_crypto_detail relation/edge.
	PaymentCryptoDetailColumn = "payment_payment_crypto_detail"
	// RevisionsTable is the table that holds the revisions relation/edge.
	RevisionsTable = "revisions"
	// RevisionsInverseTable is the table name for the Revision entity.
	// It exists in this package in order to avoid circular dependency with the "revision" package.
	RevisionsInverseTable = "revisions"
	// RevisionsColumn is the table column denoting the revisions relation/edge.
	RevisionsColumn = "payment_revisions"
)

// Columns holds all SQL columns for payment fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldApprovedBy,
	FieldApprovedAt,
	FieldMerchantID,
	FieldMerchantUserID,
	FieldMethod,
	FieldType,
	FieldStatus,
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
	// CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	CreatedByValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// UpdatedByValidator is a validator for the "updated_by" field. It is called by the builders before save.
	UpdatedByValidator func(string) error
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultMerchantID holds the default value on creation for the "merchant_id" field.
	DefaultMerchantID int64
	// MerchantIDValidator is a validator for the "merchant_id" field. It is called by the builders before save.
	MerchantIDValidator func(int64) error
	// DefaultMerchantUserID holds the default value on creation for the "merchant_user_id" field.
	DefaultMerchantUserID int64
	// MerchantUserIDValidator is a validator for the "merchant_user_id" field. It is called by the builders before save.
	MerchantUserIDValidator func(int64) error
	// DefaultMethod holds the default value on creation for the "method" field.
	DefaultMethod int32
	// MethodValidator is a validator for the "method" field. It is called by the builders before save.
	MethodValidator func(int32) error
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType int32
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(int32) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int32
	// StatusValidator is a validator for the "status" field. It is called by the builders before save.
	StatusValidator func(int32) error
)
