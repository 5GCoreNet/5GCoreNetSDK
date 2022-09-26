package network

import (
	"fmt"
	"regexp"
)

// ApplicationId is a string providing an application identifier.
type ApplicationId string

// Validate validates the ApplicationId string.
func (a *ApplicationId) Validate() error {
	if a != nil {
		return nil
	}
	return fmt.Errorf("applicationId mustn't be null")
}

// ApplicationIdRm is defined in the same way as the ApplicationId data type, but with the OpenAPI "nullable: true" property.
type ApplicationIdRm ApplicationId

func (a *ApplicationIdRm) Validate() error {
	if a != nil {
		return (*ApplicationId)(a).Validate()
	}
	return nil
}

// PduSessionId is an unsigned integer identifying a PDU session, within the range 0 to 255, as specified in subclause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007
type PduSessionId uint8

// Validate validates the PduSessionId.
func (p *PduSessionId) Validate() error {
	if p != nil {
		return nil
	}
	return fmt.Errorf("pduSessionId mustn't be null")
}

// Mcc is the Mobile Country Code part of the PLMN, comprising 3 digits, as defined in subclause 9.3.3.5 of 3GPP TS 38.413
type Mcc string

// Validate validates the Mcc string.
func (m *Mcc) Validate() error {
	if m != nil {
		if !regexp.MustCompile("[0-9]{3}$").MatchString(string(*m)) {
			return fmt.Errorf("mcc must follow this pattern '[0-9]{3}$'")
		}
		return nil
	}
	return fmt.Errorf("mcc mustn't be null")
}

// MccRm is defined in the same way as the Mcc data type, but with the OpenAPI "nullable: true" property.
type MccRm Mcc

// Validate validates the MccRm string.
func (m *MccRm) Validate() error {
	if m != nil {
		return (*Mcc)(m).Validate()
	}
	return nil
}

// Mnc is the Mobile Network Code part of the PLMN, comprising 2 or 3 digits, as defined in subclause 9.3.3.5 of 3GPP TS 38.413.
type Mnc string

// Validate validates the Mnc string.
func (m *Mnc) Validate() error {
	if m != nil {
		if !regexp.MustCompile("[0-9]{2,3}$").MatchString(string(*m)) {
			return fmt.Errorf("mnc must follow this pattern '[0-9]{2,3}$'")
		}
		return nil
	}
	return fmt.Errorf("mnc mustn't be null")
}

// MncRm is defined in the same way as the Mnc data type, but with the OpenAPI "nullable: true" property.
type MncRm Mnc

// Validate validates the MncRm string.
func (m *MncRm) Validate() error {
	if m != nil {
		return (*Mnc)(m).Validate()
	}
	return nil
}

// Tac is 2 or 3-octet string identifying a tracking area code as specified in subclause 9.3.3.10 of 3GPP TS 38.413 [11], in hexadecimal representation.
// Each character in the string shall take a value of
// "0" to "9" or "A" to "F" and shall represent 4 bits. The most
// significant character representing the 4 most significant bits of the
// TAC shall appear first in the string, and the character
// representing the 4 least significant bit of the TAC shall appear last
// in the string.
type Tac string

// Validate validates the Tac string.
func (t *Tac) Validate() error {
	if t != nil {
		return nil
	}
	return fmt.Errorf("tac mustn't be null")
}

// TacRm is defined in the same way as the Tac data type, but with the OpenAPI "nullable: true" property.
type TacRm Tac

func (t *TacRm) Validate() error {
	if t != nil {
		return (*Tac)(t).Validate()
	}
	return nil
}

// EutraCellId is a 28-bit string identifying an E-UTRA Cell Id as specified in subclause 9.3.1.9 of 3GPP TS 38.413 [11], in hexadecimal representation.
//  Each character in the string shall take a value of
// "0" to "9" or "A" to "F" and shall represent 4 bits. The most
// significant character representing the 4 most significant bits of the
// Cell Id shall appear first in the string, and the character
// representing the 4 least significant bit of the Cell Id shall appear
// last in the string.
type EutraCellId string

// Validate validates the EutraCellId string.
func (e *EutraCellId) Validate() error {
	if e != nil {
		if !regexp.MustCompile("^[A-Fa-f0-9]{7}$").MatchString(string(*e)) {
			return fmt.Errorf("eutraCellId must follow this pattern '[0-9A-F]{7}$'")
		}
		return nil
	}
	return fmt.Errorf("eutraCellId mustn't be null")
}

// EutraCellIdRm is defined in the same way as the EutraCellId data type, but with the OpenAPI "nullable: true" property.
type EutraCellIdRm EutraCellId

// Validate validates the EutraCellIdRm string.
func (e *EutraCellIdRm) Validate() error {
	if e != nil {
		return (*EutraCellId)(e).Validate()
	}
	return nil
}

// NrCellId is a 36-bit string identifying an NR Cell Id as specified in subclause 9.3.1.7 of 3GPP TS 38.413 [11], in hexadecimal representation.
// Each character in the string shall take a value of "0" to "9" or "A"
// to "F" and shall represent 4 bits. The most significant character
// representing the 4 most significant bits of the Cell Id shall appear
// first in the string, and the character representing the 4 least
// significant bit of the Cell Id shall appear last in the string.
type NrCellId string

// Validate validates the NrCellId string.
func (n *NrCellId) Validate() error {
	if n != nil {
		if !regexp.MustCompile("^[A-Fa-f0-9]{9}$").MatchString(string(*n)) {
			return fmt.Errorf("nrCellId must follow this pattern '[0-9A-F]{9}$'")
		}
		return nil
	}
	return fmt.Errorf("nrCellId mustn't be null")
}

// NrCellIdRm is defined in the same way as the NrCellId data type, but with the OpenAPI "nullable: true" property.
type NrCellIdRm NrCellId

// Validate validates the NrCellIdRm string.
func (n *NrCellIdRm) Validate() error {
	if n != nil {
		return (*NrCellId)(n).Validate()
	}
	return nil
}

// Dnai is DNAI (Data network access identifier), see subclause 5.6.7 of 3GPP TS 23.501.
type Dnai string

// Validate validates the Dnai string.
func (d *Dnai) Validate() error {
	if d != nil {
		return nil
	}
	return fmt.Errorf("dnai mustn't be null")
}

// DnaiRm is defined in the same way as the Dnai data type, but with the OpenAPI "nullable: true" property.
type DnaiRm Dnai

// Validate validates the DnaiRm string.
func (d *DnaiRm) Validate() error {
	if d != nil {
		return (*Dnai)(d).Validate()
	}
	return nil
}

// FiveGMmCause represents the 5GMM cause code values as specified in 3GPP TS 24.501.
type FiveGMmCause uint64

// Validate validates the FiveGMmCause value.
func (f *FiveGMmCause) Validate() error {
	if f != nil {
		return nil
	}
	return fmt.Errorf("fiveGMmCause mustn't be null")
}

// AreaCode is operator specific.
type AreaCode string

// Validate validates the AreaCode string.
func (a *AreaCode) Validate() error {
	if a != nil {
		return nil
	}
	return fmt.Errorf("areaCode mustn't be null")
}

// AreaCodeRm is defined in the same way as the AreaCode data type, but with the OpenAPI "nullable: true" property.
type AreaCodeRm AreaCode

// Validate validates the AreaCodeRm string.
func (a *AreaCodeRm) Validate() error {
	if a != nil {
		return (*AreaCode)(a).Validate()
	}
	return nil
}

// AmfName is the FQDN (Fully Qualified Domain Name) of the AMF as defined in subclause 28.3.2.5 of 3GPP TS 23.003.
type AmfName string

// Validate validates the AmfName string.
func (a *AmfName) Validate() error {
	if a != nil {
		return nil
	}
	return fmt.Errorf("amfName mustn't be null")
}

// N3IwfId represents the identifier of the N3IWF ID as specified in subclause 9.3.1.57 of 3GPP TS 38.413 [11]
type N3IwfId string

// Validate validates the NThreeIwfId string.
func (n *N3IwfId) Validate() error {
	if n != nil {
		if !regexp.MustCompile("^[A-Fa-f0-9]+$").MatchString(string(*n)) {
			return fmt.Errorf("nThreeIwfId must follow this pattern '[0-9A-F]{8}$'")
		}
		return nil
	}
	return fmt.Errorf("nThreeIwfId mustn't be null")
}

// NgeNbId represents the identifier of the ng-eNB ID as specified in subclause 9.3.1.8 of 3GPP TS 38.413.
// The string shall be formatted with following pattern:
// Pattern: '^('MacroNGeNB-[A-Fa-f0-9]{5}|
//			  LMacroNGeNB-[A-Fa-f0-9]{6}|
//            SMacroNGeNB-[A-Fa-f0-9]{5})$'
// The value of the ng-eNB ID shall be encoded in hexadecimal
// representation. Each character in the string shall take a value of
// "0" to "9" or "A" to "F" and shall represent 4 bits. The padding 0
// shall be added to make multiple nibbles, so the most significant
// character representing the padding 0 if required together with the
// 4 most significant bits of the ng-eNB ID shall appear first in the
// string, and the character representing the 4 least significant bit of
// the ng-eNB ID (to form a nibble) shall appear last in the string.
type NgeNbId string

// Validate validates the NgeNbId string.
func (n *NgeNbId) Validate() error {
	if n != nil {
		if !regexp.MustCompile("^(MacroNGeNB-[A-Fa-f0-9]{5}|LMacroNGeNB-[A-Fa-f0-9]{6}|SMacroNGeNB-[A-Fa-f0-9]{5})$").MatchString(string(*n)) {
			return fmt.Errorf("ngeNbId must follow this pattern '^(MacroNGeNB-[0-9A-F]{5}|LMacroNGeNB-[0-9A-F]{6}|SMacroNGeNB-[0-9A-F]{5})$'")
		}
		return nil
	}
	return fmt.Errorf("ngeNbId mustn't be null")
}
