// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: credentials.proto

package credentialspb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Unauthenticated with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Unauthenticated) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Unauthenticated with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UnauthenticatedMultiError, or nil if none found.
func (m *Unauthenticated) ValidateAll() error {
	return m.validate(true)
}

func (m *Unauthenticated) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UnauthenticatedMultiError(errors)
	}
	return nil
}

// UnauthenticatedMultiError is an error wrapping multiple validation errors
// returned by Unauthenticated.ValidateAll() if the designated constraints
// aren't met.
type UnauthenticatedMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnauthenticatedMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnauthenticatedMultiError) AllErrors() []error { return m }

// UnauthenticatedValidationError is the validation error returned by
// Unauthenticated.Validate if the designated constraints aren't met.
type UnauthenticatedValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnauthenticatedValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnauthenticatedValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnauthenticatedValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnauthenticatedValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnauthenticatedValidationError) ErrorName() string { return "UnauthenticatedValidationError" }

// Error satisfies the builtin error interface
func (e UnauthenticatedValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnauthenticated.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnauthenticatedValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnauthenticatedValidationError{}

// Validate checks the field values on CloudEnvironment with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CloudEnvironment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CloudEnvironment with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CloudEnvironmentMultiError, or nil if none found.
func (m *CloudEnvironment) ValidateAll() error {
	return m.validate(true)
}

func (m *CloudEnvironment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CloudEnvironmentMultiError(errors)
	}
	return nil
}

// CloudEnvironmentMultiError is an error wrapping multiple validation errors
// returned by CloudEnvironment.ValidateAll() if the designated constraints
// aren't met.
type CloudEnvironmentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CloudEnvironmentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CloudEnvironmentMultiError) AllErrors() []error { return m }

// CloudEnvironmentValidationError is the validation error returned by
// CloudEnvironment.Validate if the designated constraints aren't met.
type CloudEnvironmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CloudEnvironmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CloudEnvironmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CloudEnvironmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CloudEnvironmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CloudEnvironmentValidationError) ErrorName() string { return "CloudEnvironmentValidationError" }

// Error satisfies the builtin error interface
func (e CloudEnvironmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCloudEnvironment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CloudEnvironmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CloudEnvironmentValidationError{}

// Validate checks the field values on BasicAuth with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BasicAuth) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BasicAuth with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BasicAuthMultiError, or nil
// if none found.
func (m *BasicAuth) ValidateAll() error {
	return m.validate(true)
}

func (m *BasicAuth) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Password

	if len(errors) > 0 {
		return BasicAuthMultiError(errors)
	}
	return nil
}

// BasicAuthMultiError is an error wrapping multiple validation errors returned
// by BasicAuth.ValidateAll() if the designated constraints aren't met.
type BasicAuthMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BasicAuthMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BasicAuthMultiError) AllErrors() []error { return m }

// BasicAuthValidationError is the validation error returned by
// BasicAuth.Validate if the designated constraints aren't met.
type BasicAuthValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BasicAuthValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BasicAuthValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BasicAuthValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BasicAuthValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BasicAuthValidationError) ErrorName() string { return "BasicAuthValidationError" }

// Error satisfies the builtin error interface
func (e BasicAuthValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBasicAuth.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BasicAuthValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BasicAuthValidationError{}

// Validate checks the field values on Header with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Header) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Header with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in HeaderMultiError, or nil if none found.
func (m *Header) ValidateAll() error {
	return m.validate(true)
}

func (m *Header) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Value

	if len(errors) > 0 {
		return HeaderMultiError(errors)
	}
	return nil
}

// HeaderMultiError is an error wrapping multiple validation errors returned by
// Header.ValidateAll() if the designated constraints aren't met.
type HeaderMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HeaderMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HeaderMultiError) AllErrors() []error { return m }

// HeaderValidationError is the validation error returned by Header.Validate if
// the designated constraints aren't met.
type HeaderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeaderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeaderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeaderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeaderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeaderValidationError) ErrorName() string { return "HeaderValidationError" }

// Error satisfies the builtin error interface
func (e HeaderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeader.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeaderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeaderValidationError{}

// Validate checks the field values on ClientCrednetials with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ClientCrednetials) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClientCrednetials with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ClientCrednetialsMultiError, or nil if none found.
func (m *ClientCrednetials) ValidateAll() error {
	return m.validate(true)
}

func (m *ClientCrednetials) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TenantId

	// no validation rules for ClientId

	// no validation rules for ClientSecret

	if len(errors) > 0 {
		return ClientCrednetialsMultiError(errors)
	}
	return nil
}

// ClientCrednetialsMultiError is an error wrapping multiple validation errors
// returned by ClientCrednetials.ValidateAll() if the designated constraints
// aren't met.
type ClientCrednetialsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClientCrednetialsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClientCrednetialsMultiError) AllErrors() []error { return m }

// ClientCrednetialsValidationError is the validation error returned by
// ClientCrednetials.Validate if the designated constraints aren't met.
type ClientCrednetialsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientCrednetialsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientCrednetialsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientCrednetialsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientCrednetialsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientCrednetialsValidationError) ErrorName() string {
	return "ClientCrednetialsValidationError"
}

// Error satisfies the builtin error interface
func (e ClientCrednetialsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientCrednetials.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientCrednetialsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientCrednetialsValidationError{}

// Validate checks the field values on ClientCertificate with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ClientCertificate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClientCertificate with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ClientCertificateMultiError, or nil if none found.
func (m *ClientCertificate) ValidateAll() error {
	return m.validate(true)
}

func (m *ClientCertificate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TenantId

	// no validation rules for ClientId

	// no validation rules for CertificatePath

	// no validation rules for CertificatePassword

	if len(errors) > 0 {
		return ClientCertificateMultiError(errors)
	}
	return nil
}

// ClientCertificateMultiError is an error wrapping multiple validation errors
// returned by ClientCertificate.ValidateAll() if the designated constraints
// aren't met.
type ClientCertificateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClientCertificateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClientCertificateMultiError) AllErrors() []error { return m }

// ClientCertificateValidationError is the validation error returned by
// ClientCertificate.Validate if the designated constraints aren't met.
type ClientCertificateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientCertificateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientCertificateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientCertificateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientCertificateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientCertificateValidationError) ErrorName() string {
	return "ClientCertificateValidationError"
}

// Error satisfies the builtin error interface
func (e ClientCertificateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientCertificate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientCertificateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientCertificateValidationError{}

// Validate checks the field values on Oauth2 with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Oauth2) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Oauth2 with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in Oauth2MultiError, or nil if none found.
func (m *Oauth2) ValidateAll() error {
	return m.validate(true)
}

func (m *Oauth2) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RefreshToken

	// no validation rules for ClientId

	// no validation rules for ClientSecret

	if len(errors) > 0 {
		return Oauth2MultiError(errors)
	}
	return nil
}

// Oauth2MultiError is an error wrapping multiple validation errors returned by
// Oauth2.ValidateAll() if the designated constraints aren't met.
type Oauth2MultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Oauth2MultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Oauth2MultiError) AllErrors() []error { return m }

// Oauth2ValidationError is the validation error returned by Oauth2.Validate if
// the designated constraints aren't met.
type Oauth2ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Oauth2ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Oauth2ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Oauth2ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Oauth2ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Oauth2ValidationError) ErrorName() string { return "Oauth2ValidationError" }

// Error satisfies the builtin error interface
func (e Oauth2ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOauth2.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Oauth2ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Oauth2ValidationError{}

// Validate checks the field values on KeySecret with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *KeySecret) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on KeySecret with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in KeySecretMultiError, or nil
// if none found.
func (m *KeySecret) ValidateAll() error {
	return m.validate(true)
}

func (m *KeySecret) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Secret

	if len(errors) > 0 {
		return KeySecretMultiError(errors)
	}
	return nil
}

// KeySecretMultiError is an error wrapping multiple validation errors returned
// by KeySecret.ValidateAll() if the designated constraints aren't met.
type KeySecretMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m KeySecretMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m KeySecretMultiError) AllErrors() []error { return m }

// KeySecretValidationError is the validation error returned by
// KeySecret.Validate if the designated constraints aren't met.
type KeySecretValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KeySecretValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KeySecretValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KeySecretValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KeySecretValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KeySecretValidationError) ErrorName() string { return "KeySecretValidationError" }

// Error satisfies the builtin error interface
func (e KeySecretValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKeySecret.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KeySecretValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KeySecretValidationError{}

// Validate checks the field values on AWS with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *AWS) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AWS with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AWSMultiError, or nil if none found.
func (m *AWS) ValidateAll() error {
	return m.validate(true)
}

func (m *AWS) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Secret

	// no validation rules for Region

	if len(errors) > 0 {
		return AWSMultiError(errors)
	}
	return nil
}

// AWSMultiError is an error wrapping multiple validation errors returned by
// AWS.ValidateAll() if the designated constraints aren't met.
type AWSMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AWSMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AWSMultiError) AllErrors() []error { return m }

// AWSValidationError is the validation error returned by AWS.Validate if the
// designated constraints aren't met.
type AWSValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AWSValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AWSValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AWSValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AWSValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AWSValidationError) ErrorName() string { return "AWSValidationError" }

// Error satisfies the builtin error interface
func (e AWSValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAWS.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AWSValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AWSValidationError{}

// Validate checks the field values on SES with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *SES) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SES with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SESMultiError, or nil if none found.
func (m *SES) ValidateAll() error {
	return m.validate(true)
}

func (m *SES) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCreds()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SESValidationError{
					field:  "Creds",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SESValidationError{
					field:  "Creds",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreds()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SESValidationError{
				field:  "Creds",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Sender

	if len(errors) > 0 {
		return SESMultiError(errors)
	}
	return nil
}

// SESMultiError is an error wrapping multiple validation errors returned by
// SES.ValidateAll() if the designated constraints aren't met.
type SESMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SESMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SESMultiError) AllErrors() []error { return m }

// SESValidationError is the validation error returned by SES.Validate if the
// designated constraints aren't met.
type SESValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SESValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SESValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SESValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SESValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SESValidationError) ErrorName() string { return "SESValidationError" }

// Error satisfies the builtin error interface
func (e SESValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSES.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SESValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SESValidationError{}

// Validate checks the field values on GitHubApp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GitHubApp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GitHubApp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GitHubAppMultiError, or nil
// if none found.
func (m *GitHubApp) ValidateAll() error {
	return m.validate(true)
}

func (m *GitHubApp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PrivateKey

	// no validation rules for InstallationId

	// no validation rules for AppId

	if len(errors) > 0 {
		return GitHubAppMultiError(errors)
	}
	return nil
}

// GitHubAppMultiError is an error wrapping multiple validation errors returned
// by GitHubApp.ValidateAll() if the designated constraints aren't met.
type GitHubAppMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GitHubAppMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GitHubAppMultiError) AllErrors() []error { return m }

// GitHubAppValidationError is the validation error returned by
// GitHubApp.Validate if the designated constraints aren't met.
type GitHubAppValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GitHubAppValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GitHubAppValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GitHubAppValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GitHubAppValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GitHubAppValidationError) ErrorName() string { return "GitHubAppValidationError" }

// Error satisfies the builtin error interface
func (e GitHubAppValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGitHubApp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GitHubAppValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GitHubAppValidationError{}
