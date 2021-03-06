// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CryptoWalletsColumns holds the columns for the "crypto_wallets" table.
	CryptoWalletsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_by", Type: field.TypeString},
		{Name: "merchant_id", Type: field.TypeInt64, Default: 0},
		{Name: "merchant_user_id", Type: field.TypeInt64, Default: 0},
		{Name: "crypto_type", Type: field.TypeInt32, Default: 0},
		{Name: "crypto_network_type", Type: field.TypeInt32, Default: 0},
		{Name: "address", Type: field.TypeString},
		{Name: "status", Type: field.TypeInt32, Default: 0},
	}
	// CryptoWalletsTable holds the schema information for the "crypto_wallets" table.
	CryptoWalletsTable = &schema.Table{
		Name:       "crypto_wallets",
		Columns:    CryptoWalletsColumns,
		PrimaryKey: []*schema.Column{CryptoWalletsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "cryptowallet_crypto_type_crypto_network_type_address",
				Unique:  true,
				Columns: []*schema.Column{CryptoWalletsColumns[7], CryptoWalletsColumns[8], CryptoWalletsColumns[9]},
			},
			{
				Name:    "cryptowallet_merchant_id",
				Unique:  false,
				Columns: []*schema.Column{CryptoWalletsColumns[5]},
			},
			{
				Name:    "cryptowallet_crypto_type_crypto_network_type_status",
				Unique:  false,
				Columns: []*schema.Column{CryptoWalletsColumns[7], CryptoWalletsColumns[8], CryptoWalletsColumns[10]},
			},
		},
	}
	// MerchantUserBankAccountsColumns holds the columns for the "merchant_user_bank_accounts" table.
	MerchantUserBankAccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_by", Type: field.TypeString},
		{Name: "bank_name", Type: field.TypeInt32, Default: 0},
		{Name: "account_number", Type: field.TypeString},
		{Name: "account_name", Type: field.TypeString},
	}
	// MerchantUserBankAccountsTable holds the schema information for the "merchant_user_bank_accounts" table.
	MerchantUserBankAccountsTable = &schema.Table{
		Name:       "merchant_user_bank_accounts",
		Columns:    MerchantUserBankAccountsColumns,
		PrimaryKey: []*schema.Column{MerchantUserBankAccountsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "merchantuserbankaccount_bank_name_account_number_account_name",
				Unique:  true,
				Columns: []*schema.Column{MerchantUserBankAccountsColumns[5], MerchantUserBankAccountsColumns[6], MerchantUserBankAccountsColumns[7]},
			},
		},
	}
	// PaymentsColumns holds the columns for the "payments" table.
	PaymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_by", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "approved_by", Type: field.TypeString, Nullable: true},
		{Name: "approved_at", Type: field.TypeTime, Nullable: true},
		{Name: "merchant_id", Type: field.TypeInt64, Default: 0},
		{Name: "merchant_user_id", Type: field.TypeInt64, Default: 0},
		{Name: "method", Type: field.TypeInt32, Default: 0},
		{Name: "type", Type: field.TypeInt32, Default: 0},
		{Name: "status", Type: field.TypeInt32, Default: 0},
	}
	// PaymentsTable holds the schema information for the "payments" table.
	PaymentsTable = &schema.Table{
		Name:       "payments",
		Columns:    PaymentsColumns,
		PrimaryKey: []*schema.Column{PaymentsColumns[0]},
	}
	// PaymentBankingDetailsColumns holds the columns for the "payment_banking_details" table.
	PaymentBankingDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_by", Type: field.TypeString},
		{Name: "payment_code", Type: field.TypeString, Nullable: true},
		{Name: "merchant_user_id", Type: field.TypeInt64, Nullable: true},
		{Name: "merchant_user_bank_name", Type: field.TypeInt32, Nullable: true},
		{Name: "merchant_user_account_number", Type: field.TypeString, Nullable: true},
		{Name: "merchant_user_account_name", Type: field.TypeString, Nullable: true},
		{Name: "system_account_bank_name", Type: field.TypeInt32, Nullable: true},
		{Name: "system_account_number", Type: field.TypeString, Nullable: true},
		{Name: "system_account_name", Type: field.TypeString, Nullable: true},
		{Name: "image_url", Type: field.TypeString, Nullable: true},
		{Name: "tx_id", Type: field.TypeString, Nullable: true},
		{Name: "amount", Type: field.TypeUint64, Nullable: true},
		{Name: "fee", Type: field.TypeUint64, Nullable: true},
		{Name: "payment_payment_banking_detail", Type: field.TypeInt64, Nullable: true},
	}
	// PaymentBankingDetailsTable holds the schema information for the "payment_banking_details" table.
	PaymentBankingDetailsTable = &schema.Table{
		Name:       "payment_banking_details",
		Columns:    PaymentBankingDetailsColumns,
		PrimaryKey: []*schema.Column{PaymentBankingDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payment_banking_details_payments_payment_banking_detail",
				Columns:    []*schema.Column{PaymentBankingDetailsColumns[17]},
				RefColumns: []*schema.Column{PaymentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PaymentCryptoDetailsColumns holds the columns for the "payment_crypto_details" table.
	PaymentCryptoDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_by", Type: field.TypeString},
		{Name: "crypto_type", Type: field.TypeInt32, Default: 0},
		{Name: "crypto_network_type", Type: field.TypeInt32, Default: 0},
		{Name: "crypto_wallet_name", Type: field.TypeInt32, Default: 0},
		{Name: "receiver_address", Type: field.TypeString},
		{Name: "sender_address", Type: field.TypeString, Nullable: true},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true},
		{Name: "received_amount", Type: field.TypeFloat64, Nullable: true},
		{Name: "tx_hash", Type: field.TypeString, Nullable: true},
		{Name: "tx_id", Type: field.TypeString, Nullable: true},
		{Name: "fee", Type: field.TypeFloat64, Nullable: true},
		{Name: "image_url", Type: field.TypeString, Nullable: true},
		{Name: "payment_payment_crypto_detail", Type: field.TypeInt64, Nullable: true},
	}
	// PaymentCryptoDetailsTable holds the schema information for the "payment_crypto_details" table.
	PaymentCryptoDetailsTable = &schema.Table{
		Name:       "payment_crypto_details",
		Columns:    PaymentCryptoDetailsColumns,
		PrimaryKey: []*schema.Column{PaymentCryptoDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payment_crypto_details_payments_payment_crypto_detail",
				Columns:    []*schema.Column{PaymentCryptoDetailsColumns[16]},
				RefColumns: []*schema.Column{PaymentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PaymentEwalletDetailsColumns holds the columns for the "payment_ewallet_details" table.
	PaymentEwalletDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_by", Type: field.TypeString},
		{Name: "payment_code", Type: field.TypeString, Nullable: true},
		{Name: "merchant_user_id", Type: field.TypeInt64, Nullable: true},
		{Name: "e_wallet_name", Type: field.TypeInt32, Nullable: true},
		{Name: "merchant_user_account_phone_number", Type: field.TypeString, Nullable: true},
		{Name: "merchant_user_account_name", Type: field.TypeString, Nullable: true},
		{Name: "system_account_phone_number", Type: field.TypeString, Nullable: true},
		{Name: "system_account_name", Type: field.TypeString, Nullable: true},
		{Name: "amount", Type: field.TypeUint64, Nullable: true},
		{Name: "fee", Type: field.TypeUint64, Nullable: true},
		{Name: "note", Type: field.TypeString, Nullable: true},
		{Name: "image_url", Type: field.TypeString, Nullable: true},
		{Name: "tx_id", Type: field.TypeString, Nullable: true},
		{Name: "payment_payment_e_wallet_detail", Type: field.TypeInt64, Nullable: true},
	}
	// PaymentEwalletDetailsTable holds the schema information for the "payment_ewallet_details" table.
	PaymentEwalletDetailsTable = &schema.Table{
		Name:       "payment_ewallet_details",
		Columns:    PaymentEwalletDetailsColumns,
		PrimaryKey: []*schema.Column{PaymentEwalletDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payment_ewallet_details_payments_payment_e_wallet_detail",
				Columns:    []*schema.Column{PaymentEwalletDetailsColumns[17]},
				RefColumns: []*schema.Column{PaymentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PaymentTelcoDetailsColumns holds the columns for the "payment_telco_details" table.
	PaymentTelcoDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_by", Type: field.TypeString},
		{Name: "telco_name", Type: field.TypeInt32, Default: 0},
		{Name: "serial_number", Type: field.TypeString, Default: ""},
		{Name: "card_id", Type: field.TypeString, Default: ""},
		{Name: "charged_amount", Type: field.TypeUint64, Default: 0},
		{Name: "amount", Type: field.TypeUint64, Default: 0},
		{Name: "payment_payment_telco_detail", Type: field.TypeInt64, Nullable: true},
	}
	// PaymentTelcoDetailsTable holds the schema information for the "payment_telco_details" table.
	PaymentTelcoDetailsTable = &schema.Table{
		Name:       "payment_telco_details",
		Columns:    PaymentTelcoDetailsColumns,
		PrimaryKey: []*schema.Column{PaymentTelcoDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payment_telco_details_payments_payment_telco_detail",
				Columns:    []*schema.Column{PaymentTelcoDetailsColumns[10]},
				RefColumns: []*schema.Column{PaymentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RevisionsColumns holds the columns for the "revisions" table.
	RevisionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_by", Type: field.TypeString},
		{Name: "status", Type: field.TypeInt32, Default: 0},
		{Name: "note", Type: field.TypeString},
		{Name: "payment_revisions", Type: field.TypeInt64, Nullable: true},
	}
	// RevisionsTable holds the schema information for the "revisions" table.
	RevisionsTable = &schema.Table{
		Name:       "revisions",
		Columns:    RevisionsColumns,
		PrimaryKey: []*schema.Column{RevisionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "revisions_payments_revisions",
				Columns:    []*schema.Column{RevisionsColumns[7]},
				RefColumns: []*schema.Column{PaymentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SettingsColumns holds the columns for the "settings" table.
	SettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
	}
	// SettingsTable holds the schema information for the "settings" table.
	SettingsTable = &schema.Table{
		Name:       "settings",
		Columns:    SettingsColumns,
		PrimaryKey: []*schema.Column{SettingsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "setting_name",
				Unique:  true,
				Columns: []*schema.Column{SettingsColumns[1]},
			},
		},
	}
	// SystemBankAccountsColumns holds the columns for the "system_bank_accounts" table.
	SystemBankAccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_by", Type: field.TypeString},
		{Name: "merchant_id", Type: field.TypeInt64, Default: 0},
		{Name: "status", Type: field.TypeInt32, Default: 0},
		{Name: "bank_name", Type: field.TypeInt32, Default: 0},
		{Name: "account_number", Type: field.TypeString},
		{Name: "account_name", Type: field.TypeString},
		{Name: "branch", Type: field.TypeString},
		{Name: "balance", Type: field.TypeUint64, Default: 0},
		{Name: "last_updated_balance", Type: field.TypeTime, Nullable: true},
		{Name: "daily_balance_limit", Type: field.TypeUint64, Default: 0},
		{Name: "daily_used_amount", Type: field.TypeInt64, Default: 0},
		{Name: "daily_balance", Type: field.TypeUint64, Default: 0},
	}
	// SystemBankAccountsTable holds the schema information for the "system_bank_accounts" table.
	SystemBankAccountsTable = &schema.Table{
		Name:       "system_bank_accounts",
		Columns:    SystemBankAccountsColumns,
		PrimaryKey: []*schema.Column{SystemBankAccountsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "systembankaccount_bank_name_account_number_account_name",
				Unique:  true,
				Columns: []*schema.Column{SystemBankAccountsColumns[7], SystemBankAccountsColumns[8], SystemBankAccountsColumns[9]},
			},
		},
	}
	// SystemCryptoHotWalletsColumns holds the columns for the "system_crypto_hot_wallets" table.
	SystemCryptoHotWalletsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_by", Type: field.TypeString},
		{Name: "merchant_id", Type: field.TypeInt64, Default: 0},
		{Name: "crypto_type", Type: field.TypeInt32, Default: 0},
		{Name: "crypto_network_type", Type: field.TypeInt32, Default: 0},
		{Name: "address", Type: field.TypeString},
		{Name: "total_balance", Type: field.TypeFloat64, Default: 0},
		{Name: "balance", Type: field.TypeFloat64, Default: 0},
		{Name: "status", Type: field.TypeInt32, Default: 0},
	}
	// SystemCryptoHotWalletsTable holds the schema information for the "system_crypto_hot_wallets" table.
	SystemCryptoHotWalletsTable = &schema.Table{
		Name:       "system_crypto_hot_wallets",
		Columns:    SystemCryptoHotWalletsColumns,
		PrimaryKey: []*schema.Column{SystemCryptoHotWalletsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "systemcryptohotwallet_crypto_type_crypto_network_type_address",
				Unique:  true,
				Columns: []*schema.Column{SystemCryptoHotWalletsColumns[6], SystemCryptoHotWalletsColumns[7], SystemCryptoHotWalletsColumns[8]},
			},
			{
				Name:    "systemcryptohotwallet_merchant_id",
				Unique:  false,
				Columns: []*schema.Column{SystemCryptoHotWalletsColumns[5]},
			},
			{
				Name:    "systemcryptohotwallet_crypto_type_crypto_network_type_status",
				Unique:  false,
				Columns: []*schema.Column{SystemCryptoHotWalletsColumns[6], SystemCryptoHotWalletsColumns[7], SystemCryptoHotWalletsColumns[11]},
			},
		},
	}
	// SystemEwalletsColumns holds the columns for the "system_ewallets" table.
	SystemEwalletsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_by", Type: field.TypeString},
		{Name: "e_wallet_name", Type: field.TypeInt32, Default: 0},
		{Name: "status", Type: field.TypeInt32, Default: 0},
		{Name: "merchant_id", Type: field.TypeInt64, Default: 0},
		{Name: "account_phone_number", Type: field.TypeString},
		{Name: "account_name", Type: field.TypeString},
		{Name: "balance", Type: field.TypeUint64, Default: 0},
		{Name: "daily_balance", Type: field.TypeUint64, Default: 0},
		{Name: "daily_balance_limit", Type: field.TypeUint64, Default: 0},
		{Name: "daily_used_amount", Type: field.TypeInt64, Default: 0},
		{Name: "last_updated_balance", Type: field.TypeTime, Nullable: true},
	}
	// SystemEwalletsTable holds the schema information for the "system_ewallets" table.
	SystemEwalletsTable = &schema.Table{
		Name:       "system_ewallets",
		Columns:    SystemEwalletsColumns,
		PrimaryKey: []*schema.Column{SystemEwalletsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "systemewallet_merchant_id_e_wallet_name_account_phone_number",
				Unique:  true,
				Columns: []*schema.Column{SystemEwalletsColumns[7], SystemEwalletsColumns[5], SystemEwalletsColumns[8]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CryptoWalletsTable,
		MerchantUserBankAccountsTable,
		PaymentsTable,
		PaymentBankingDetailsTable,
		PaymentCryptoDetailsTable,
		PaymentEwalletDetailsTable,
		PaymentTelcoDetailsTable,
		RevisionsTable,
		SettingsTable,
		SystemBankAccountsTable,
		SystemCryptoHotWalletsTable,
		SystemEwalletsTable,
	}
)

func init() {
	PaymentBankingDetailsTable.ForeignKeys[0].RefTable = PaymentsTable
	PaymentCryptoDetailsTable.ForeignKeys[0].RefTable = PaymentsTable
	PaymentEwalletDetailsTable.ForeignKeys[0].RefTable = PaymentsTable
	PaymentTelcoDetailsTable.ForeignKeys[0].RefTable = PaymentsTable
	RevisionsTable.ForeignKeys[0].RefTable = PaymentsTable
}
