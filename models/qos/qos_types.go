package qos

import (
	"fmt"
	"regexp"
)

// Qfi is an unsigned integer identifying a QoS flow, within the range 0 to 63.
type Qfi uint8

// Validate validates this qfi.
func (q *Qfi) Validate() error {
	if q == nil {
		return fmt.Errorf("qfi must not be nil")
	}
	if *q < 0 || *q > 63 {
		return fmt.Errorf("qfi %d is out of range [0, 63]", *q)
	}
	return nil
}

// QfiRm data type is defined in the same way as the Qfi data type, but with the OpenAPI "nullable: true" property.
type QfiRm Qfi

// Validate validates this qfi rm.
func (q *QfiRm) Validate() error {
	if q != nil {
		return (*Qfi)(q).Validate()
	}
	return nil
}

// FiveQi (5qi) is an unsigned integer representing a 5G QoS Identifier (see subclause 5.7.2.1 of 3GPP TS 23.501), within the range 0 to 255.
type FiveQi uint8

// Validate validates this 5qi.
func (f *FiveQi) Validate() error {
	if f == nil {
		return fmt.Errorf("5qi must not be nil")
	}
	if *f < 0 || *f > 255 {
		return fmt.Errorf("5qi %d is out of range [0, 255]", *f)
	}
	return nil
}

// FiveQiRm data type is defined in the same way as the FiveQi data type, but with the OpenAPI "nullable: true" property.
type FiveQiRm FiveQi

// Validate validates this 5qi rm.
func (f *FiveQiRm) Validate() error {
	if f != nil {
		return (*FiveQi)(f).Validate()
	}
	return nil
}

// BitRate is a string representing a bit rate.
type BitRate string

// Validate validates this bit rate.
func (b *BitRate) Validate() error {
	if b == nil {
		return fmt.Errorf("bit rate must not be nil")
	}
	if !regexp.MustCompile("^\\d+(\\.\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$").MatchString(string(*b)) {
		return fmt.Errorf("bit rate %s is not valid", *b)
	}
	return nil
}

// BitRateRm data type is defined in the same way as the BitRate data type, but with the OpenAPI "nullable: true" property.
type BitRateRm BitRate

// Validate validates this bit rate rm.
func (b *BitRateRm) Validate() error {
	if b != nil {
		return (*BitRate)(b).Validate()
	}
	return nil
}

// ArpPriorityLevel is an unsignedinteger indicating the ARP Priority Level (see subclause 5.7.2.2 of 3GPP TS 23.501), within the range 1 to 15.
//Values are ordered in decreasing order of priority, i.e. with 1 as the highest priority and 15 as the lowest priority.
type ArpPriorityLevel uint8

// Validate validates this arp priority level.
func (a *ArpPriorityLevel) Validate() error {
	if a == nil {
		return fmt.Errorf("arp priority level must not be nil")
	}
	if *a < 1 || *a > 15 {
		return fmt.Errorf("arp priority level %d is out of range [1, 15]", *a)
	}
	return nil
}

// ArpPriorityLevelRm data type is defined in the same way as the ArpPriorityLevel data type, but with the OpenAPI "nullable: true" property.
type ArpPriorityLevelRm ArpPriorityLevel

// Validate validates this arp priority level rm.
func (a *ArpPriorityLevelRm) Validate() error {
	if a != nil {
		return (*ArpPriorityLevel)(a).Validate()
	}
	return nil
}

// FiveQiPriorityLevel is an unsigned integer indicating the 5QI Priority Level (see subclauses 5.7.3.3 and 5.7.4 of 3GPP TS 23.501), within the range 1 to 127.
//Values are ordered in decreasing order of priority, i.e. with 1 as the highest priority and 127 as the lowest priority.
type FiveQiPriorityLevel uint8

// Validate validates this 5qi priority level.
func (f *FiveQiPriorityLevel) Validate() error {
	if f == nil {
		return fmt.Errorf("5qi priority level must not be nil")
	}
	if *f < 1 || *f > 127 {
		return fmt.Errorf("5qi priority level %d is out of range [1, 127]", *f)
	}
	return nil
}

// FiveQiPriorityLevelRm data type is defined in the same way as the FiveQiPriorityLevel data type, but with the OpenAPI "nullable: true" property.
type FiveQiPriorityLevelRm FiveQiPriorityLevel

// Validate validates this 5qi priority level rm.
func (f *FiveQiPriorityLevelRm) Validate() error {
	if f != nil {
		return (*FiveQiPriorityLevel)(f).Validate()
	}
	return nil
}

// PacketDelBudget is an unsigned integer indicating Packet Delay Budget (see subclauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
// Minimum = 1.
type PacketDelBudget uint64

// Validate validates this packet del budget.
func (p *PacketDelBudget) Validate() error {
	if p == nil {
		return fmt.Errorf("packet del budget must not be nil")
	}
	if *p < 1 {
		return fmt.Errorf("packet del budget %d is out of range [1, infinity]", *p)
	}
	return nil
}

// PacketDelBudgetRm data type is defined in the same way as the PacketDelBudget data type, but with the OpenAPI "nullable: true" property.
type PacketDelBudgetRm PacketDelBudget

// Validate validates this packet del budget rm.
func (p *PacketDelBudgetRm) Validate() error {
	if p != nil {
		return (*PacketDelBudget)(p).Validate()
	}
	return nil
}

// PacketErrRate is a string representing Packet Error Rate (see subclause 5.7.3.5 and 5.7.4 of 3GPP TS 23.501), expressed as a "scalar x 10-k" where the scalar and the exponent k are each encoded as one decimal digit.
type PacketErrRate string

// Validate validates this packet err rate.
func (p *PacketErrRate) Validate() error {
	if p == nil {
		return fmt.Errorf("packet err rate must not be nil")
	}
	if !regexp.MustCompile("^([0-9]E-[0-9])$").MatchString(string(*p)) {
		return fmt.Errorf("packet err rate %s is not valid", *p)
	}
	return nil
}

// PacketErrRateRm data type is defined in the same way as the PacketErrRate data type, but with the OpenAPI "nullable: true" property.
type PacketErrRateRm PacketErrRate

// Validate validates this packet err rate rm.
func (p *PacketErrRateRm) Validate() error {
	if p != nil {
		return (*PacketErrRate)(p).Validate()
	}
	return nil
}

// PacketLossRate is an unsigned integer indicating Packet Loss Rate (see subclauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent.
//Minimum = 0. Maximum = 1000.
type PacketLossRate uint16

// Validate validates this packet loss rate.
func (p *PacketLossRate) Validate() error {
	if p == nil {
		return fmt.Errorf("packet loss rate must not be nil")
	}
	if *p < 0 || *p > 1000 {
		return fmt.Errorf("packet loss rate %d is out of range [0, 1000]", *p)
	}
	return nil
}

// PacketLossRateRm data type is defined in the same way as the PacketLossRate data type, but with the OpenAPI "nullable: true" property.
type PacketLossRateRm PacketLossRate

// Validate validates this packet loss rate rm.
func (p *PacketLossRateRm) Validate() error {
	if p != nil {
		return (*PacketLossRate)(p).Validate()
	}
	return nil
}

// AverWindow is an unsigned integer indicating Averaging Window (see subclause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
//Minimum = 1. Maximum = 4095.
type AverWindow uint16

// Validate validates this aver window.
func (a *AverWindow) Validate() error {
	if a == nil {
		return fmt.Errorf("aver window must not be nil")
	}
	if *a < 1 || *a > 4095 {
		return fmt.Errorf("aver window %d is out of range [1, 4095]", *a)
	}
	return nil
}

// AverWindowRm data type is defined in the same way as the AverWindow data type, but with the OpenAPI "nullable: true" property.
type AverWindowRm AverWindow

// Validate validates this aver window rm.
func (a *AverWindowRm) Validate() error {
	if a != nil {
		return (*AverWindow)(a).Validate()
	}
	return nil
}

// MaxDataBurstVol is an unsigned integer indicating Maximum Data Burst Volume (see subclauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.
//Minimum = 1. Maximum = 4095.
type MaxDataBurstVol uint16

// Validate validates this max data burst vol.
func (m *MaxDataBurstVol) Validate() error {
	if m == nil {
		return fmt.Errorf("max data burst vol must not be nil")
	}
	if *m < 1 || *m > 4095 {
		return fmt.Errorf("max data burst vol %d is out of range [1, 4095]", *m)
	}
	return nil
}

// MaxDataBurstVolRm data type is defined in the same way as the MaxDataBurstVol data type, but with the OpenAPI "nullable: true" property.
type MaxDataBurstVolRm MaxDataBurstVol

// Validate validates this max data burst vol rm.
func (m *MaxDataBurstVolRm) Validate() error {
	if m != nil {
		return (*MaxDataBurstVol)(m).Validate()
	}
	return nil
}
