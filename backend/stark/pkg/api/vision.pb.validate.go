// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: stark/api/vision.proto

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
var _vision_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetPaymentDetailRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetPaymentDetailRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return GetPaymentDetailRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// GetPaymentDetailRequestValidationError is the validation error returned by
// GetPaymentDetailRequest.Validate if the designated constraints aren't met.
type GetPaymentDetailRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPaymentDetailRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPaymentDetailRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPaymentDetailRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPaymentDetailRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPaymentDetailRequestValidationError) ErrorName() string {
	return "GetPaymentDetailRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetPaymentDetailRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPaymentDetailRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPaymentDetailRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPaymentDetailRequestValidationError{}

// Validate checks the field values on GetPaymentDetailReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetPaymentDetailReply) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPayment()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetPaymentDetailReplyValidationError{
				field:  "Payment",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetPaymentDetail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetPaymentDetailReplyValidationError{
				field:  "PaymentDetail",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetRevisions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetPaymentDetailReplyValidationError{
					field:  fmt.Sprintf("Revisions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetPaymentDetailReplyValidationError is the validation error returned by
// GetPaymentDetailReply.Validate if the designated constraints aren't met.
type GetPaymentDetailReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPaymentDetailReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPaymentDetailReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPaymentDetailReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPaymentDetailReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPaymentDetailReplyValidationError) ErrorName() string {
	return "GetPaymentDetailReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetPaymentDetailReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPaymentDetailReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPaymentDetailReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPaymentDetailReplyValidationError{}

// Validate checks the field values on ListPaymentsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListPaymentsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Page

	// no validation rules for Size

	if v, ok := interface{}(m.GetFrom()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListPaymentsRequestValidationError{
				field:  "From",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListPaymentsRequestValidationError{
				field:  "To",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OrderAscCreatedAt

	// no validation rules for OrderAscUpdatedAt

	return nil
}

// ListPaymentsRequestValidationError is the validation error returned by
// ListPaymentsRequest.Validate if the designated constraints aren't met.
type ListPaymentsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPaymentsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPaymentsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPaymentsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPaymentsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPaymentsRequestValidationError) ErrorName() string {
	return "ListPaymentsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListPaymentsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPaymentsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPaymentsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPaymentsRequestValidationError{}

// Validate checks the field values on ListPaymentsReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListPaymentsReply) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRecords() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListPaymentsReplyValidationError{
					field:  fmt.Sprintf("Records[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	// no validation rules for CurrentPage

	return nil
}

// ListPaymentsReplyValidationError is the validation error returned by
// ListPaymentsReply.Validate if the designated constraints aren't met.
type ListPaymentsReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPaymentsReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPaymentsReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPaymentsReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPaymentsReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPaymentsReplyValidationError) ErrorName() string {
	return "ListPaymentsReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListPaymentsReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPaymentsReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPaymentsReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPaymentsReplyValidationError{}

// Validate checks the field values on GetPaymentInfoByPaymentCodeRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *GetPaymentInfoByPaymentCodeRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	return nil
}

// GetPaymentInfoByPaymentCodeRequestValidationError is the validation error
// returned by GetPaymentInfoByPaymentCodeRequest.Validate if the designated
// constraints aren't met.
type GetPaymentInfoByPaymentCodeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPaymentInfoByPaymentCodeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPaymentInfoByPaymentCodeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPaymentInfoByPaymentCodeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPaymentInfoByPaymentCodeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPaymentInfoByPaymentCodeRequestValidationError) ErrorName() string {
	return "GetPaymentInfoByPaymentCodeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetPaymentInfoByPaymentCodeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPaymentInfoByPaymentCodeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPaymentInfoByPaymentCodeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPaymentInfoByPaymentCodeRequestValidationError{}

// Validate checks the field values on GetPaymentInfoByPaymentCodeReply with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *GetPaymentInfoByPaymentCodeReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MerchantId

	// no validation rules for MerchantUserId

	// no validation rules for PaymentMethod

	switch m.Payload.(type) {

	case *GetPaymentInfoByPaymentCodeReply_BankName:
		// no validation rules for BankName

	case *GetPaymentInfoByPaymentCodeReply_EWalletName:
		// no validation rules for EWalletName

	}

	return nil
}

// GetPaymentInfoByPaymentCodeReplyValidationError is the validation error
// returned by GetPaymentInfoByPaymentCodeReply.Validate if the designated
// constraints aren't met.
type GetPaymentInfoByPaymentCodeReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPaymentInfoByPaymentCodeReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPaymentInfoByPaymentCodeReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPaymentInfoByPaymentCodeReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPaymentInfoByPaymentCodeReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPaymentInfoByPaymentCodeReplyValidationError) ErrorName() string {
	return "GetPaymentInfoByPaymentCodeReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetPaymentInfoByPaymentCodeReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPaymentInfoByPaymentCodeReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPaymentInfoByPaymentCodeReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPaymentInfoByPaymentCodeReplyValidationError{}
