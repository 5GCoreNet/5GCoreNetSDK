package subscription

import (
	"fmt"
	"regexp"
)

// TODO: Add description regarding the ETSI documentation
// TODO: Add validation functions that follow the ETSI documentation

// Dnn reprsents  a Data Network as defined in subclause 9A of
// 3GPP TS 23.003. It shall be formatted as string in which the
// labels are separated by dots (e.g. "Label1.Label2.Label3").
type Dnn string

// Validate validates the Dnn string.
func (d *Dnn) Validate() error {
	if d == nil {
		return fmt.Errorf("dnn mustn't be null")
	}
	return nil
}

// DnnRm is defined in the same way as the Dnn data type, but with the OpenAPI "nullable: true" property.
type DnnRm Dnn

// Validate validates the DnnRm string.
func (d *DnnRm) Validate() error {
	if d == nil {
		return nil
	}
	return (*Dnn)(d).Validate()
}

// Gpsi dentifying a Gpsi shall contain either an External Id or an MSISDN. It
// shall be formatted as follows:
//	-External Identifier: "extid-<extid>, where <extid> shall be formatted
//	according to subclause 19.7.2 of 3GPP TS 23.003 that describes an External
//	Identifier. -MSISDN: "msisdn-<msisdn>, where <msisdn> shall be formatted
//	according to subclause 3.3 of 3GPP TS 23.003 that describes an MSISDN.
type Gpsi string

// Validate validates the Gpsi string.
func (g *Gpsi) Validate() error {
	if g == nil {
		return fmt.Errorf("gpsi mustn't be null")
	}
	if !regexp.MustCompile(`^(msisdn-\d{5,15}|extid-.+@.+|.+)$`).MatchString(string(*g)) {
		return fmt.Errorf("gpsi must follow this pattern '^(msisdn-[0-9]{5,15}|extid-.+@.+|.+)$'")
	}
	return nil
}

// GpsiRm is defined in the same way as the Gpsi data type, but with the OpenAPI "nullable: true" property.
type GpsiRm Gpsi

// Validate validates the GpsiRm string.
func (g *GpsiRm) Validate() error {
	if g == nil {
		return nil
	}
	return (*Gpsi)(g).Validate()
}

// GroupId identifying a group of devices network internal globally
// unique ID which identifies a set of IMSIs, as specified in
// subclause 19.9 of 3GPP TS 23.003.
type GroupId string

// Validate validates the GroupId string.
func (g *GroupId) Validate() error {
	if g == nil {
		return fmt.Errorf("group id mustn't be null")
	}
	if !regexp.MustCompile(`^[A-Fa-f\d]{8}-\d{3}-\d{2,3}-([A-Fa-f\d][A-Fa-f\d]){1,10}$`).MatchString(string(*g)) {
		return fmt.Errorf("group id must follow this pattern '^imsi-[0-9]{5,15}$'")
	}
	return nil
}

// GroupIdRm is defined in the same way as the GroupId data type, but with the OpenAPI "nullable: true" property.
type GroupIdRm GroupId

// Validate validates the GroupIdRm string.
func (g *GroupIdRm) Validate() error {
	if g == nil {
		return nil
	}
	return (*GroupId)(g).Validate()
}

// Pei representing a Permanent Equipment Identifier, if it contains an IMEI or
// IMEISV it is defined as specified in subclause 6.2 of 3GPP TS 23.003.
type Pei string

// Validate validates the Pei string.
func (p *Pei) Validate() error {
	if p == nil {
		return fmt.Errorf("pei mustn't be null")
	}
	if !regexp.MustCompile(`^(imei-\d{15}|imeisv-\d{16}|.+)$`).MatchString(string(*p)) {
		return fmt.Errorf("pei must follow this pattern '^imei-[0-9]{5,15}$'")
	}
	return nil
}

// PeiRm is defined in the same way as the Pei data type, but with the OpenAPI "nullable: true" property.
type PeiRm Pei

// Validate validates the PeiRm string.
func (p *PeiRm) Validate() error {
	if p == nil {
		return nil
	}
	return (*Pei)(p).Validate()
}

// Supi identifying a Supi shall contain either an IMSI or an NAI. It
// shall be formatted as follows for:
//	- IMSI "imsi-<imsi>, <imsi> shall be formatted according to subclause 2.2 of
//	3GPP TS 23.003 that describes an IMSI. - NAI "nai-<nai>, <nai> shall be
//	formatted according to subclause 28.6.2 of 3GPP TS 23.003 that describes
//	an NAI.
// 	To enable that the value is used as part of an URI, the string shall
// 	only contain characters allowed according to the "lower-with-
// 	hyphen" naming convention defined in 3GPP TS 29.501.
type Supi string

// Validate validates the Supi string.
func (s *Supi) Validate() error {
	if s == nil {
		return fmt.Errorf("supi mustn't be null")
	}
	if !regexp.MustCompile(`^(imsi-\d{5,15}|nai-.+|.+)$`).MatchString(string(*s)) {
		return fmt.Errorf("supi must follow this pattern '^(imsi-[0-9]{5,15}|nai-.+@.+|.+)$'")
	}
	return nil
}

// SupiRm is defined in the same way as the Supi data type, but with the OpenAPI "nullable: true" property.
type SupiRm Supi

// Validate validates the SupiRm string.
func (s *SupiRm) Validate() error {
	if s == nil {
		return nil
	}
	return (*Supi)(s).Validate()
}

// NfInstanceId uniquely identifying a NF instance. The format of the NF Instance
// ID shall be a Universally Unique Identifier (UUID) version 4, as described in
// IETF RFC 4122
type NfInstanceId string

// Validate validates the NfInstanceId string.
func (n *NfInstanceId) Validate() error {
	if n == nil {
		return fmt.Errorf("nf instance id mustn't be null")
	}
	if !regexp.MustCompile(`^[\da-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-f]{4}-[\da-f]{12}$`).MatchString(string(*n)) {
		return fmt.Errorf("nf instance id must follow this pattern (UUID v4) '^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$'")
	}
	return nil
}

// AmfId identifying the AMF ID composed of AMF Region ID (8
// bits), AMF Set ID (10 bits) and AMF Pointer (6 bits) as specified in subclause
// 2.10.1 of 3GPP TS 23.003. It is encoded as a string of 6 hexadecimal
// characters (i.e., 24 bits).
type AmfId string

// Validate validates the AmfId string.
func (i *AmfId) Validate() error {
	if i == nil {
		return fmt.Errorf("amf id mustn't be null")
	}
	if !regexp.MustCompile(`^[\dA-Faa-f]{6}$`).MatchString(string(*i)) {
		return fmt.Errorf("amf id must follow this pattern '^[0-9A-Faa-f]{6}$'")
	}
	return nil
}

// AmfRegionId identifying the AMF Region ID (8 bits), as specified in subclause
// 2.10.1 of 3GPP TS 23.003. It is encoded as a string of 2 hexadecimal
// characters (i.e. 8 bits).
type AmfRegionId string

// Validate validates the AmfRegionId string.
func (i *AmfRegionId) Validate() error {
	if i == nil {
		return fmt.Errorf("amf region id mustn't be null")
	}
	if !regexp.MustCompile(`^[\dA-Faa-f]{2}$`).MatchString(string(*i)) {
		return fmt.Errorf("amf region id must follow this pattern '^[0-9A-Faa-f]{2}$'")
	}
	return nil
}

// AmfSetId identifying the AMF Set ID (10 bits) as specified in
// subclause 2.10.1 of 3GPP TS 23.003. It is encoded as a string of 3
// hexadecimal characters where the first character is limited to values 0 to 3
// (i.e. 10 bits).
type AmfSetId string

// Validate validates the AmfSetId string.
func (i *AmfSetId) Validate() error {
	if i == nil {
		return fmt.Errorf("amf set id mustn't be null")
	}
	if !regexp.MustCompile(`^[0-3][\dA-Faa-f]{2}$`).MatchString(string(*i)) {
		return fmt.Errorf("amf set id must follow this pattern '^[0-3][0-9A-Faa-f]{2}$'")
	}
	return nil
}

// RfspIndex representing the "Subscriber Profile ID for
// RAT/Frequency Priority" as specified in 3GPP TS 36.413.
// Minimum = 1. Maximum = 256.
type RfspIndex uint16

// Validate validates the RfspIndex.
func (i *RfspIndex) Validate() error {
	if i == nil {
		return fmt.Errorf("rfsp index mustn't be null")
	}
	if *i < 1 || *i > 256 {
		return fmt.Errorf("rfsp index must be between 1 and 256")
	}
	return nil
}

// RfspIndexRm is defined in the same way as the RfspIndex data type, but with the OpenAPI "nullable: true" property.
type RfspIndexRm RfspIndex

// Validate validates the RfspIndexRm.
func (i *RfspIndexRm) Validate() error {
	if i == nil {
		return nil
	}
	return (*RfspIndex)(i).Validate()
}

// NfGroupId is the identifier of a group of NFs.
type NfGroupId string

// Validate validates the NfGroupId string.
func (n *NfGroupId) Validate() error {
	if n == nil {
		return fmt.Errorf("nf group id mustn't be null")
	}
	return nil
}
