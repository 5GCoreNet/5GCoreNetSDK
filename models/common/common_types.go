package common

import (
	"fmt"
	"net/url"
	"regexp"
)

type OpenAPINullString string

const (
	NullString OpenAPINullString = "null"
)

// Binary is a string with format "binary" as defined in OpenAPI Specification.
type Binary string

// Validate validates the Binary string.
// Binary string mustn't be "null".
func (b Binary) Validate() error {
	if b != Binary(NullString) {
		return nil
	}
	return fmt.Errorf("binary mustn't be null, use binaryRm instead")
}

// BinaryRm is is defined in the same way as the "Binary" data type, but with the OpenAPI "nullable: true" property.
type BinaryRm struct {
	Binary Binary `json:"binary,omitempty"`
	Null   bool   `json:"null,omitempty"`
}

// Validate validates the BinaryRm string.
// As there is no validation for BinaryRm strings, this function always returns nil.
func (b BinaryRm) Validate() error {
	return nil
}

// Bytes is a string with format "byte" as defined in OpenAPI Specification.
type Bytes string

// Validate validates the Bytes string.
// As there is no validation for Bytes strings, this function always returns nil.
func (b Bytes) Validate() error {
	if b != Bytes(NullString) {
		return nil
	}
	return fmt.Errorf("bytes mustn't be nulln, use bytesRm instead")
}

// BytesRm is defined in the same way as the "Bytes" data type, but with the OpenAPI "nullable: true" property.
type BytesRm Bytes

// Validate validates the BytesRm string.
// As there is no validation for BytesRm strings, this function always returns nil.
func (b BytesRm) Validate() error {
	return nil
}

// Date is a string with format "date" as defined in OpenAPI Specification.
type Date string

// Validate validates the Date string.
// As there is no validation for Date strings, this function always returns nil.
func (d Date) Validate() error {
	if d != Date(NullString) {
		return nil
	}
	return fmt.Errorf("date mustn't be null, use dateRm instead")
}

// DateRm is defined in the same way as the "Date" data type, but with the OpenAPI "nullable: true" property.
type DateRm Date

// Validate validates the DateRm string.
// As there is no validation for DateRm strings, this function always returns nil.
func (d DateRm) Validate() error {
	return nil
}

// DateTime is a string with format "date-time" as defined in OpenAPI Specification.
type DateTime string

// Validate validates the DateTime string.
// As there is no validation for DateTime strings, this function always returns nil.
func (d *DateTime) Validate() error {
	if d != nil {
		return nil
	}
	return fmt.Errorf("date-time mustn't be null, use date-timeRm instead")
}

// DateTimeRm is defined in the same way as the "DateTime" data type, but with the OpenAPI "nullable: true" property.
type DateTimeRm DateTime

// Validate validates the DateTimeRm string.
// As there is no validation for DateTimeRm strings, this function always returns nil.
func (d DateTimeRm) Validate() error {
	return nil
}

// DiameterIdentity is a string containing a Diameter Identity, according to clause 4.3 of IETF RFC 6733.
// DiameterIdentity must follow this pattern '^([A-Za-z0-9]+(-[A-Za-z0-9]+).)+[a-z]{2,}$'.
type DiameterIdentity string

// Validate validates the DiameterIdentity string
func (d DiameterIdentity) Validate() error {
	if d != DiameterIdentity(NullString) {
		if matched, _ := regexp.MatchString("^([A-Za-z0-9]+(-[A-Za-z0-9]+).)+[a-z]{2,}$", string(d)); !matched {
			return fmt.Errorf("diameter identity must follow this pattern '^([A-Za-z0-9]+(-[A-Za-z0-9]+).)+[a-z]{2,}$'")
		}
		return nil
	}
	return fmt.Errorf("diameter identity mustn't be null")
}

// DiameterIdentityRm is defined in the same way as the "DiameterIdentity" data type, but with the OpenAPI "nullable: true" property.
type DiameterIdentityRm DiameterIdentity

// Validate validates the DiameterIdentityRm string
func (d DiameterIdentityRm) Validate() error {
	if d != DiameterIdentityRm(NullString) {
		return DiameterIdentity(d).Validate()
	}
	return nil
}

// Double is a number with format "double" as defined in OpenAPI Specification.
type Double float64

// Validate validates the Double number.
// As there is no validation for Double number, this function always returns nil.
func (d Double) Validate() error {

	return nil
}

// DoubleRm is defined in the same way as the "Double" data type, but with the OpenAPI "nullable: true" property.
type DoubleRm Double

// Validate validates the DoubleRm number.
// As there is no validation for DoubleRm number, this function always returns nil.
func (d DoubleRm) Validate() error {
	return nil
}

// DurationSec is an unsigned integer identifying a period of time in units of seconds
type DurationSec uint64

func (d DurationSec) Validate() error {
	// TODO: Add validation function
	return nil
}

type DurationSecRm DurationSec

func (d DurationSecRm) Validate() error {
	// TODO: Add validation function
	if d != 0 {
		return DurationSec(d).Validate()
	}
	return nil
}

type Float float64

func (f Float) Validate() error {
	//TODO: Add validation function
	return nil
}

type FloatRm Float

func (f FloatRm) Validate() error {
	// TODO: Add validation function
	if f != 0 {
		return Float(f).Validate()
	}
	return nil
}

type Uint16 uint16

func (u Uint16) Validate() error {
	// TODO: Add validation function
	return nil
}

type Uint16Rm Uint16

func (u Uint16Rm) Validate() error {
	// TODO: Add validation function
	if u != 0 {
		return Uint16(u).Validate()
	}
	return nil
}

type Int32 int32
type Int32Rm Int32
type Int64 int64

// Validate validates the Int64 number.
func (i *Int64) Validate() error {
	if i != nil {
		return nil
	}
	return fmt.Errorf("int64 mustn't be null")
}

type Int64Rm Int64
type Ipv4Addr string

// Validate validates the Ipv4Addr string.
func (a *Ipv4Addr) Validate() error {
	if a != nil {
		return nil
	}
	return fmt.Errorf("ipv4 address mustn't be null")
}

type Ipv4AddrRm Ipv4Addr
type Ipv6Addr string

// Validate validates the Ipv6Addr string.
func (a *Ipv6Addr) Validate() error {
	if a != nil {
		return nil
	}
	return fmt.Errorf("ipv6 address mustn't be null")
}

type Ipv6AddrRm Ipv6Addr
type Ipv6Prefix string
type Ipv6PrefixRm Ipv6Prefix
type MacAddr48 string
type MacAddr48Rm MacAddr48
type SupportedFeatures string
type Uinteger uint64

// Validate validates the Uinteger number.
func (u *Uinteger) Validate() error {
	if u != nil {
		return nil
	}
	return fmt.Errorf("uinteger mustn't be null")
}

type UintegerRm Uinteger
type Uint32 uint32
type Uint32Rm Uint32
type Uint64 uint64
type Uint64Rm Uint64

// Uri is a string providing an URI formatted according to IETF RFC 3986.
type Uri string

// Validate validates the Uri string
// Uri musn't be null and must follow the IETF RFC 3986 format.
func (u Uri) Validate() error {
	if u != Uri(NullString) {
		_, err := url.ParseRequestURI(string(u))
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("uri mustn't be null")
}

type UriRm Uri
type VarUeId string
type VarUeIdRm VarUeId
type TimeZone string
type TimeZoneRm TimeZone