package common

import (
	"fmt"
	"math"
	"net/url"
	"regexp"
)

// Binary is a string with format "binary" as defined in OpenAPI Specification.
type Binary string

// Validate validates the Binary string.
func (b *Binary) Validate() error {
	if b == nil {
		return fmt.Errorf("binary mustn't be null")
	}
	return nil
}

// BinaryRm is is defined in the same way as the "Binary" data type, but with the OpenAPI "nullable: true" property.
type BinaryRm Binary

// Validate validates the BinaryRm string.
func (b *BinaryRm) Validate() error {
	if b == nil {
		return nil
	}
	return (*Binary)(b).Validate()
}

// Bytes is a string with format "byte" as defined in OpenAPI Specification.
type Bytes string

// Validate validates the Bytes string.
func (b *Bytes) Validate() error {
	if b == nil {
		return fmt.Errorf("bytes mustn't be null")
	}
	return nil
}

// BytesRm is defined in the same way as the "Bytes" data type, but with the OpenAPI "nullable: true" property.
type BytesRm Bytes

// Validate validates the BytesRm string.
func (b *BytesRm) Validate() error {
	if b == nil {
		return nil
	}
	return (*Bytes)(b).Validate()
}

// Date is a string with format "date" as defined in OpenAPI Specification.
type Date string

// Validate validates the Date string.
func (d *Date) Validate() error {
	if d == nil {
		return fmt.Errorf("date mustn't be null")
	}
	return nil
}

// DateRm is defined in the same way as the "Date" data type, but with the OpenAPI "nullable: true" property.
type DateRm Date

// Validate validates the DateRm string.
func (d *DateRm) Validate() error {
	if d == nil {
		return nil
	}
	return (*Date)(d).Validate()
}

// DateTime is a string with format "date-time" as defined in OpenAPI Specification.
type DateTime string

// Validate validates the DateTime string.
func (d *DateTime) Validate() error {
	if d == nil {
		return fmt.Errorf("date-time mustn't be null")
	}
	return nil
}

// DateTimeRm is defined in the same way as the "DateTime" data type, but with the OpenAPI "nullable: true" property.
type DateTimeRm DateTime

// Validate validates the DateTimeRm string.
func (d *DateTimeRm) Validate() error {
	if d == nil {
		return nil
	}
	return (*DateTime)(d).Validate()
}

// DiameterIdentity is a string containing a Diameter Identity, according to clause 4.3 of IETF RFC 6733.
// DiameterIdentity must follow this pattern '^([A-Za-z0-9]+(-[A-Za-z0-9]+).)+[a-z]{2,}$'.
type DiameterIdentity string

// Validate validates the DiameterIdentity string
func (d *DiameterIdentity) Validate() error {
	if d == nil {
		return fmt.Errorf("diameter identity mustn't be null")
	}
	if !regexp.MustCompile(`^([A-Za-z0-9]+(-[A-Za-z0-9]+).)+[a-z]{2,}$`).MatchString(string(*d)) {
		return fmt.Errorf("diameter identity must follow this pattern '^([A-Za-z0-9]+(-[A-Za-z0-9]+).)+[a-z]{2,}$'")
	}
	return nil
}

// DiameterIdentityRm is defined in the same way as the "DiameterIdentity" data type, but with the OpenAPI "nullable: true" property.
type DiameterIdentityRm DiameterIdentity

// Validate validates the DiameterIdentityRm string
func (d *DiameterIdentityRm) Validate() error {
	if d == nil {
		return nil
	}
	return (*DiameterIdentity)(d).Validate()
}

// Double is a number with format "double" as defined in OpenAPI Specification.
type Double float64

// Validate validates the Double number.
func (d *Double) Validate() error {
	if d == nil {
		return fmt.Errorf("double mustn't be null")
	}
	return nil
}

// DoubleRm is defined in the same way as the "Double" data type, but with the OpenAPI "nullable: true" property.
type DoubleRm Double

// Validate validates the DoubleRm number.
func (d *DoubleRm) Validate() error {
	if d == nil {
		return nil
	}
	return (*Double)(d).Validate()
}

// DurationSec is an unsigned integer identifying a period of time in units of seconds
type DurationSec uint64

func (d *DurationSec) Validate() error {
	if d == nil {
		return fmt.Errorf("duration sec mustn't be null")
	}
	return nil
}

// DurationSecRm is defined in the same way as the DurationSec data type, but with the OpenAPI "nullable: true" property.
type DurationSecRm DurationSec

// Validate validates the DurationSecRm.
func (d *DurationSecRm) Validate() error {
	if d == nil {
		return nil
	}
	return (*DurationSec)(d).Validate()
}

// Float is a number with format "float" as defined in OpenAPI Specification.
type Float float64

func (f *Float) Validate() error {
	if f == nil {
		return fmt.Errorf("float mustn't be null")
	}
	return nil
}

// FloatRm is defined in the same way as the Float data type, but with the OpenAPI "nullable: true" property.
type FloatRm Float

// Validate validates the FloatRm.
func (f *FloatRm) Validate() error {
	if f == nil {
		return nil
	}
	return (*Float)(f).Validate()
}

// Uint16 is an unsigned 6-bit integers, i.e. only value between 0 and 65535 are permissible.
type Uint16 uint16

// Validate validates the Uint16.
func (u *Uint16) Validate() error {
	if u == nil {
		return fmt.Errorf("uint16 mustn't be null")
	}
	if *u < 0 || *u > 65535 {
		return fmt.Errorf("uint16 must be between 0 and 65535")
	}
	return nil
}

// Uint16Rm is defined in the same way as the Uint16 data type, but with the OpenAPI "nullable: true" property.
type Uint16Rm Uint16

// Validate validates the Uint16Rm.
func (u *Uint16Rm) Validate() error {
	if u == nil {
		return nil
	}
	return (*Uint16)(u).Validate()
}

// Int32 is an integer with format "int32" as defined in OpenAPI Specification.
type Int32 int32

// Validate validates the Int32.
func (i *Int32) Validate() error {
	if i == nil {
		return fmt.Errorf("int32 mustn't be null")
	}
	return nil
}

// Int32Rm is defined in the same way as the Int32 data type, but with the OpenAPI "nullable: true" property.
type Int32Rm Int32

// Validate validates the Int32Rm.
func (i *Int32Rm) Validate() error {
	if i == nil {
		return nil
	}
	return (*Int32)(i).Validate()
}

// Int64 is an integer with format "int64" as defined in OpenAPI Specification.
type Int64 int64

// Validate validates the Int64 number.
func (i *Int64) Validate() error {
	if i == nil {
		return fmt.Errorf("int64 mustn't be null")
	}
	return nil
}

// Int64Rm is defined in the same way as the Int64 data type, but with the OpenAPI "nullable: true" property.
type Int64Rm Int64

// Validate validates the Int64Rm number.
func (i *Int64Rm) Validate() error {
	if i == nil {
		return nil
	}
	return (*Int64)(i).Validate()
}

// Ipv4Addr is a string dentifying a IPv4 address formatted in the "dotted decimal" notation as defined in in IETF RFC 1166.
// Must follow this pattern '^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$'.
type Ipv4Addr string

// Validate validates the Ipv4Addr string.
func (a *Ipv4Addr) Validate() error {
	if a == nil {
		return fmt.Errorf("ipv4 address mustn't be null")
	}
	if !regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$`).MatchString(string(*a)) {
		return fmt.Errorf("ipv4 address must follow this pattern '^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$'")
	}
	return nil
}

// Ipv4AddrRm is defined in the same way as the Ipv4Addr data type, but with the OpenAPI "nullable: true" property.
type Ipv4AddrRm Ipv4Addr

// Validate validates the Ipv4AddrRm string.
func (a *Ipv4AddrRm) Validate() error {
	if a == nil {
		return nil
	}
	return (*Ipv4Addr)(a).Validate()
}

// Ipv6Addr is a string identifying an IPv6 address formatted according to clause 4
//of IETF RFC 5952. The mixed IPv4 IPv6 notation according to clause 5 of
//IETF RFC 5952 shall not be used.
// Must follow this pattern '^((:|(0?|([1-9a-f][0-9a-f]{0,3}))):)((0?|([1-9a-f][0-9a-f]{0,3})):){0,6}(:|(0?|([1-9a-f][0-9a-f]{0,3})))$' or '((([^:]+:){7}([^:]+))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?))$'.
type Ipv6Addr string

// Validate validates the Ipv6Addr string.
func (a *Ipv6Addr) Validate() error {
	if a == nil {
		return fmt.Errorf("ipv6 address mustn't be null")
	}
	if !regexp.MustCompile(`^((:|(0?|([1-9a-f][0-9a-f]{0,3}))):)((0?|([1-9a-f][0-9a-f]{0,3})):){0,6}(:|(0?|([1-9a-f][0-9a-f]{0,3})))$`).MatchString(string(*a)) || !regexp.MustCompile(`((([^:]+:){7}([^:]+))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?))$`).MatchString(string(*a)) {
		return fmt.Errorf("ipv6 address must follow this pattern '^((:|(0?|([1-9a-f][0-9a-f]{0,3}))):)((0?|([1-9a-f][0-9a-f]{0,3})):){0,6}(:|(0?|([1-9a-f][0-9a-f]{0,3})))$' or '((([^:]+:){7}([^:]+))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?))$'")
	}
	return nil
}

// Ipv6AddrRm is defined in the same way as the Ipv6Addr data type, but with the OpenAPI "nullable: true" property.
type Ipv6AddrRm Ipv6Addr

// Validate validates the Ipv6AddrRm string.
func (a *Ipv6AddrRm) Validate() error {
	if a == nil {
		return nil
	}
	return (*Ipv6Addr)(a).Validate()
}

// Ipv6Prefix is a string identifying an IPv6 prefix formatted according to clause 4 of IETF RFC 5952.
// Must follow this pattern: ^((:|(0?|([1-9a-f][0-9a-f]{0,3}))):)((0?|([1-9a-f][0-9a-f]{0,3})):){0,6}(:|(0?|([1-9a-f][0-9a-f]{0,3})))(\/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))$' or '^((([^:]+:){7}([^:]+))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?))(\/.+)$'.
type Ipv6Prefix string

// Validate validates the Ipv6Prefix string.
func (a *Ipv6Prefix) Validate() error {
	if a == nil {
		return fmt.Errorf("ipv6 prefix mustn't be null")
	}
	if !regexp.MustCompile(`^((:|(0?|([1-9a-f][0-9a-f]{0,3}))):)((0?|([1-9a-f][0-9a-f]{0,3})):){0,6}(:|(0?|([1-9a-f][0-9a-f]{0,3})))(\\/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))$`).MatchString(string(*a)) || !regexp.MustCompile(`^((([^:]+:){7}([^:]+))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?))(\\/.)$`).MatchString(string(*a)) {
		return fmt.Errorf("ipv6 prefix must follow this pattern: '^((:|(0?|([1-9a-f][0-9a-f]{0,3}))):)((0?|([1-9a-f][0-9a-f]{0,3})):){0,6}(:|(0?|([1-9a-f][0-9a-f]{0,3})))(\\/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))$' or '^((([^:]+:){7}([^:]+))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?))(\\/.)$'")
	}
	return nil
}

// Ipv6PrefixRm is defined in the same way as the Ipv6Prefix data type, but with the OpenAPI "nullable: true" property.
type Ipv6PrefixRm Ipv6Prefix

// Validate validates the Ipv6PrefixRm string.
func (a *Ipv6PrefixRm) Validate() error {
	if a == nil {
		return nil
	}
	return (*Ipv6Prefix)(a).Validate()
}

// MacAddr48 is a string identifying a MAC address formatted in the hexadecimal notation according to subclause 1.1 and subclause 2.1 of IETF RFC 7042.
// Must follow this pattern: '^([0-9a-fA-F]{2})((-[0-9a-fA-F]{2}){5})$'.
type MacAddr48 string

// Validate validates the MacAddr48 string.
func (a *MacAddr48) Validate() error {
	if a == nil {
		return fmt.Errorf("mac address mustn't be null")
	}
	if !regexp.MustCompile(`^([0-9a-fA-F]{2})((-[0-9a-fA-F]{2}){5})$`).MatchString(string(*a)) {
		return fmt.Errorf("mac address must follow this pattern: '^([0-9a-fA-F]{2})((-[0-9a-fA-F]{2}){5})$'")
	}
	return nil
}

// MacAddr48Rm is defined in the same way as the MacAddr48 data type, but with the OpenAPI "nullable: true" property.
type MacAddr48Rm MacAddr48

// Validate validates the MacAddr48Rm string.
func (a *MacAddr48Rm) Validate() error {
	if a == nil {
		return nil
	}
	return (*MacAddr48)(a).Validate()
}

// SupportedFeatures is a string used to indicate the features supported by an
// API that is used as defined in subclause 6.6 in 3GPP TS 29.500. The string
// shall contain a bitmask indicating supported features in hexadecimal
// representation:
//		Each character in the string shall take a value of "0" to "9" or "A"
//		to "F" and shall represent the support of 4 features as described
//		in table 5.2.2-3. The most significant character representing the
//		highest-numbered features shall appear first in the string, and the
//		character representing features 1 to 4 shall appear last in the
//		string. The list of features and their numbering (starting with 1)
//		are defined separately for each API. If the string contains a lower
//		number of characters than there are defined features for an API,
//		all features that would be represented by characters that are not
//		present in the string are not supported.
type SupportedFeatures string

// Validate validates the SupportedFeatures string.
func (a *SupportedFeatures) Validate() error {
	if a == nil {
		return fmt.Errorf("supported features mustn't be null")
	}
	if !regexp.MustCompile(`^([0-9a-fA-F]{1,8})$`).MatchString(string(*a)) {
		return fmt.Errorf("supported features must follow this pattern: '^([0-9a-fA-F]{1,8})$'")
	}
	return nil
}

// Uinteger is an unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
type Uinteger uint64

// Validate validates the Uinteger number.
func (u *Uinteger) Validate() error {
	if u == nil {
		return fmt.Errorf("uinteger64 mustn't be null")
	}
	return nil
}

// UintegerRm is defined in the same way as the Uinteger data type, but with the OpenAPI "nullable: true" property.
type UintegerRm Uinteger

// Validate validates the UintegerRm number.
func (u *UintegerRm) Validate() error {
	if u == nil {
		return nil
	}
	return (*Uinteger)(u).Validate()
}

// Uint32 is an unsigned 32-bit integers, i.e. only value 0 and 32-bit integers above 0 are permissible.
type Uint32 uint32

// Validate validates the Uint32 number.
func (u *Uint32) Validate() error {
	if u == nil {
		return fmt.Errorf("uint32 mustn't be null")
	}
	if *u < 0 || *u > math.MaxUint32 {
		return fmt.Errorf("uint32 must be between 0 and %d", math.MaxUint32)
	}
	return nil
}

// Uint32Rm is defined in the same way as the Uint32 data type, but with the OpenAPI "nullable: true" property.
type Uint32Rm Uint32

// Validate validates the Uint32Rm number.
func (u *Uint32Rm) Validate() error {
	if u == nil {
		return nil
	}
	return (*Uint32)(u).Validate()
}

// Uint64 is an unsigned 64-bit integers, i.e. only value 0 and 64-bit integers above 0 are permissible.
type Uint64 uint64

// Validate validates the Uint64 number.
func (u *Uint64) Validate() error {
	if u == nil {
		return fmt.Errorf("uint64 mustn't be null")
	}
	if *u < 0 || *u > math.MaxUint64 {
		return fmt.Errorf("uint64 must be between 0 and %d", math.MaxUint64)
	}
	return nil
}

// Uint64Rm is defined in the same way as the Uint64 data type, but with the OpenAPI "nullable: true" property.
type Uint64Rm Uint64

// Validate validates the Uint64Rm number.
func (u *Uint64Rm) Validate() error {
	if u == nil {
		return nil
	}
	return (*Uint64)(u).Validate()
}

// Uri is a string providing an URI formatted according to IETF RFC 3986.
type Uri string

// Validate validates the Uri string
// Uri musn't be null and must follow the IETF RFC 3986 format.
func (u *Uri) Validate() error {
	if u == nil {
		return fmt.Errorf("uri mustn't be null")
	}
	if _, err := url.ParseRequestURI(string(*u)); err != nil {
		return fmt.Errorf("uri must follow the IETF RFC 3986 format")
	}
	return nil
}

// UriRm is defined in the same way as the Uri data type, but with the OpenAPI "nullable: true" property.
type UriRm Uri

// Validate validates the UriRm string.
func (u *UriRm) Validate() error {
	if u == nil {
		return nil
	}
	return (*Uri)(u).Validate()
}

// VarUeId is a string representing the SUPI or GPSI.
// Must follow the pattern: '^(imsi-[0-9]{5,15}|nai-.+|msisdn-[0-9]{5,15}|extid-[^@]+@[^@]+|.+)$'.
type VarUeId string

// Validate validates the VarUeId string.
func (v *VarUeId) Validate() error {
	if v == nil {
		return fmt.Errorf("var ue id mustn't be null")
	}
	if !regexp.MustCompile(`^(imsi-[0-9]{5,15}|nai-.+|msisdn-[0-9]{5,15}|extid-[^@]+@[^@]+|.+)$`).MatchString(string(*v)) {
		return fmt.Errorf("var ue id must follow the pattern: '^(imsi-[0-9]{5,15}|nai-.+|msisdn-[0-9]{5,15}|extid-[^@]+@[^@]+|.+)$'")
	}
	return nil
}

// VarUeIdRm is defined in the same way as the VarUeId data type, but with the OpenAPI "nullable: true" property.
type VarUeIdRm VarUeId

// Validate validates the VarUeIdRm string.
func (v *VarUeIdRm) Validate() error {
	if v == nil {
		return nil
	}
	return (*VarUeId)(v).Validate()
}

// TimeZone is a string with format "<time-numoffset>" optionally appended by "<daylightSavingTime>", where:
//		- <time-numoffset> shall represent the time zone adjusted for daylight saving time and be encoded as time-numoffset as
//		defined in subclause 5.6 of IETF RFC 3339 [10];
//		- <daylightSavingTime> shall represent the adjustment that has been made and shall be encoded as "+1" or
//		"+2" for a +1 or +2 hours adjustment.
//
//Example: "-08:00+1" (for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time).
type TimeZone string

// Validate validates the TimeZone string.
func (t *TimeZone) Validate() error {
	if t == nil {
		return fmt.Errorf("timezone mustn't be null")
	}
	if !regexp.MustCompile(`^([+-][0-9]{2}:[0-9]{2})([+-][0-9])?$`).MatchString(string(*t)) {
		return fmt.Errorf("timezone must follow the pattern: '^([+-][0-9]{2}:[0-9]{2})([+-][0-9])?$'")
	}
	return nil
}

// TimeZoneRm is defined in the same way as the TimeZone data type, but with the OpenAPI "nullable: true" property.
type TimeZoneRm TimeZone

// Validate validates the TimeZoneRm string.
func (t *TimeZoneRm) Validate() error {
	if t == nil {
		return nil
	}
	return (*TimeZone)(t).Validate()
}
