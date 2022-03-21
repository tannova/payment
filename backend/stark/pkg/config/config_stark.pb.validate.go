// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: stark/api/config_stark.proto

package stark

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _config_stark_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Config) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetListener() == nil {
		return ConfigValidationError{
			field:  "Listener",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetListener()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigValidationError{
				field:  "Listener",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetLogger() == nil {
		return ConfigValidationError{
			field:  "Logger",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetLogger()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigValidationError{
				field:  "Logger",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDatabase()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigValidationError{
				field:  "Database",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetNatasha()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigValidationError{
				field:  "Natasha",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetBlackwidow()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigValidationError{
				field:  "Blackwidow",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetGroot()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigValidationError{
				field:  "Groot",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDefaultSetting()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigValidationError{
				field:  "DefaultSetting",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ConfigValidationError is the validation error returned by Config.Validate if
// the designated constraints aren't met.
type ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigValidationError) ErrorName() string { return "ConfigValidationError" }

// Error satisfies the builtin error interface
func (e ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigValidationError{}

// Validate checks the field values on DefaultSetting with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DefaultSetting) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPayment()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DefaultSettingValidationError{
				field:  "Payment",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCrypto()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DefaultSettingValidationError{
				field:  "Crypto",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for MapBankName

	// no validation rules for MapEWalletName

	return nil
}

// DefaultSettingValidationError is the validation error returned by
// DefaultSetting.Validate if the designated constraints aren't met.
type DefaultSettingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DefaultSettingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DefaultSettingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DefaultSettingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DefaultSettingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DefaultSettingValidationError) ErrorName() string { return "DefaultSettingValidationError" }

// Error satisfies the builtin error interface
func (e DefaultSettingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDefaultSetting.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DefaultSettingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DefaultSettingValidationError{}

// Validate checks the field values on PaymentCode with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PaymentCode) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Method

	// no validation rules for Provider

	// no validation rules for MerchantId

	// no validation rules for UserPaymentId

	return nil
}

// PaymentCodeValidationError is the validation error returned by
// PaymentCode.Validate if the designated constraints aren't met.
type PaymentCodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaymentCodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaymentCodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaymentCodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaymentCodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaymentCodeValidationError) ErrorName() string { return "PaymentCodeValidationError" }

// Error satisfies the builtin error interface
func (e PaymentCodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaymentCode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaymentCodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaymentCodeValidationError{}

// Validate checks the field values on PaymentCrypto with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PaymentCrypto) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Domain

	// no validation rules for StoreId

	// no validation rules for SecretKey

	return nil
}

// PaymentCryptoValidationError is the validation error returned by
// PaymentCrypto.Validate if the designated constraints aren't met.
type PaymentCryptoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaymentCryptoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaymentCryptoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaymentCryptoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaymentCryptoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaymentCryptoValidationError) ErrorName() string { return "PaymentCryptoValidationError" }

// Error satisfies the builtin error interface
func (e PaymentCryptoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaymentCrypto.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaymentCryptoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaymentCryptoValidationError{}
