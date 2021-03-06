// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: stark/api/stark.proto

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
var _stark_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Payment with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Payment) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PaymentValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PaymentValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for CreatedBy

	// no validation rules for UpdatedBy

	// no validation rules for ApprovedBy

	// no validation rules for MerchantId

	// no validation rules for Method

	// no validation rules for Type

	// no validation rules for Status

	return nil
}

// PaymentValidationError is the validation error returned by Payment.Validate
// if the designated constraints aren't met.
type PaymentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaymentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaymentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaymentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaymentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaymentValidationError) ErrorName() string { return "PaymentValidationError" }

// Error satisfies the builtin error interface
func (e PaymentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaymentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaymentValidationError{}

// Validate checks the field values on PaymentDetail with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PaymentDetail) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Payload.(type) {

	case *PaymentDetail_Banking:

		if v, ok := interface{}(m.GetBanking()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PaymentDetailValidationError{
					field:  "Banking",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *PaymentDetail_EWallet:

		if v, ok := interface{}(m.GetEWallet()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PaymentDetailValidationError{
					field:  "EWallet",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *PaymentDetail_Telco:

		if v, ok := interface{}(m.GetTelco()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PaymentDetailValidationError{
					field:  "Telco",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *PaymentDetail_Crypto:

		if v, ok := interface{}(m.GetCrypto()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PaymentDetailValidationError{
					field:  "Crypto",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PaymentDetailValidationError is the validation error returned by
// PaymentDetail.Validate if the designated constraints aren't met.
type PaymentDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaymentDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaymentDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaymentDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaymentDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaymentDetailValidationError) ErrorName() string { return "PaymentDetailValidationError" }

// Error satisfies the builtin error interface
func (e PaymentDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaymentDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaymentDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaymentDetailValidationError{}

// Validate checks the field values on PaymentWithDetail with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PaymentWithDetail) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPayment()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PaymentWithDetailValidationError{
				field:  "Payment",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetPaymentDetail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PaymentWithDetailValidationError{
				field:  "PaymentDetail",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PaymentWithDetailValidationError is the validation error returned by
// PaymentWithDetail.Validate if the designated constraints aren't met.
type PaymentWithDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaymentWithDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaymentWithDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaymentWithDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaymentWithDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaymentWithDetailValidationError) ErrorName() string {
	return "PaymentWithDetailValidationError"
}

// Error satisfies the builtin error interface
func (e PaymentWithDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaymentWithDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaymentWithDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaymentWithDetailValidationError{}

// Validate checks the field values on Revision with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Revision) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RevisionValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for CreatedBy

	// no validation rules for PaymentId

	// no validation rules for Status

	// no validation rules for Note

	return nil
}

// RevisionValidationError is the validation error returned by
// Revision.Validate if the designated constraints aren't met.
type RevisionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RevisionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RevisionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RevisionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RevisionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RevisionValidationError) ErrorName() string { return "RevisionValidationError" }

// Error satisfies the builtin error interface
func (e RevisionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRevision.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RevisionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RevisionValidationError{}

// Validate checks the field values on CompletePaymentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CompletePaymentRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PaymentId

	// no validation rules for PaymentMethod

	// no validation rules for PaymentType

	// no validation rules for PaymentStatus

	if v, ok := interface{}(m.GetPaymentDetail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CompletePaymentRequestValidationError{
				field:  "PaymentDetail",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for MexCurrentBalance

	return nil
}

// CompletePaymentRequestValidationError is the validation error returned by
// CompletePaymentRequest.Validate if the designated constraints aren't met.
type CompletePaymentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompletePaymentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompletePaymentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompletePaymentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompletePaymentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompletePaymentRequestValidationError) ErrorName() string {
	return "CompletePaymentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CompletePaymentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompletePaymentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompletePaymentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompletePaymentRequestValidationError{}

// Validate checks the field values on CompletePaymentReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CompletePaymentReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CompletePaymentReplyValidationError is the validation error returned by
// CompletePaymentReply.Validate if the designated constraints aren't met.
type CompletePaymentReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompletePaymentReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompletePaymentReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompletePaymentReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompletePaymentReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompletePaymentReplyValidationError) ErrorName() string {
	return "CompletePaymentReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CompletePaymentReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompletePaymentReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompletePaymentReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompletePaymentReplyValidationError{}
