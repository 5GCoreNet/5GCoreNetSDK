package network

import "fmt"

type AccessType string

// Validate validates the AccessType string.
func (t *AccessType) Validate() error {
	if t == nil {
		return fmt.Errorf("accessType must not be nil")
	}
	switch *t {
	case AccessType3GPPAccess:
	case AccessTypeNon3GPPAccess:
	default:
		return fmt.Errorf("invalid access type: %s", *t)
	}
	return nil
}

const (
	AccessType3GPPAccess    AccessType = "3GPP_ACCESS"      // 3GPP access
	AccessTypeNon3GPPAccess AccessType = "NON_3_GPP_ACCESS" // Non-3GPP access
)

type RatType string

// Validate validates the RatType string.
func (t *RatType) Validate() error {
	if t == nil {
		return fmt.Errorf("ratType must not be nil")
	}
	switch *t {
	case RatTypeNR:
	case RatTypeEUTRA:
	case RatTypeWLAN:
	case RatTypeVIRTUAL:
	default:
		return fmt.Errorf("invalid rat type: %s", t)
	}
	return nil
}

const (
	RatTypeNR      RatType = "NR"      // New Radio
	RatTypeEUTRA   RatType = "EUTRA"   // (WB) Evolved Universal Terrestrial Radio Access
	RatTypeWLAN    RatType = "WLAN"    // Wireless LAN
	RatTypeVIRTUAL RatType = "VIRTUAL" // Virtual. Virtual shall be used if the N3IWF does not know the access technology used for an untrusted non-3GPP access.
)

// PduSessionType indicates the type of PDU session.
type PduSessionType string

// Validate validates the PduSessionType string.
func (t *PduSessionType) Validate() error {
	if t == nil {
		return fmt.Errorf("pduSessionType must not be nil")
	}
	switch *t {
	case PduSessionTypeIpv4:
	case PduSessionTypeIpv6:
	case PduSessionTypeIpv4v6:
	case PduSessionTypeUnstructured:
	case PduSessionTypeEthernet:
	default:
		return fmt.Errorf("invalid pdu session type: %s", *t)
	}
	return nil
}

const (
	PduSessionTypeIpv4         PduSessionType = "IPV4"         // IPv4
	PduSessionTypeIpv6         PduSessionType = "IPV6"         // IPv6
	PduSessionTypeIpv4v6       PduSessionType = "IPV4V6"       // IPv4v6
	PduSessionTypeUnstructured PduSessionType = "UNSTRUCTURED" // Unstructured
	PduSessionTypeEthernet     PduSessionType = "ETHERNET"     // Ethernet
)

// UpIntegrity indicates whether UP integrity protection is required, preferred or not needed for all the traffic on the PDU Session.
type UpIntegrity string

// Validate validates the UpIntegrity string.
func (t *UpIntegrity) Validate() error {
	if t == nil {
		return fmt.Errorf("upIntegrity must not be nil")
	}
	switch *t {
	case UpIntegrityRequired:
	case UpIntegrityPreferred:
	case UpIntegrityNotNeeded:
	default:
		return fmt.Errorf("invalid up integrity: %s", *t)
	}
	return nil
}

const (
	UpIntegrityRequired  UpIntegrity = "REQUIRED"   // UP integrity protection shall apply for all the traffic on the PDU Session.
	UpIntegrityPreferred UpIntegrity = "PREFERRED"  // UP integrity protection should apply for all the traffic on the PDU Session.
	UpIntegrityNotNeeded UpIntegrity = "NOT_NEEDED" // UP integrity protection shall not apply on the PDU Session.
)

// UpConfidentiality indicates whether UP confidentiality protection is required, preferred or not needed for all the traffic on the PDU Session.
type UpConfidentiality string

// Validate validates the UpConfidentiality string.
func (t *UpConfidentiality) Validate() error {
	if t == nil {
		return fmt.Errorf("upConfidentiality must not be nil")
	}
	switch *t {
	case UpConfidentialityRequired:
	case UpConfidentialityPreferred:
	case UpConfidentialityNotNeeded:
	default:
		return fmt.Errorf("invalid up confidentiality: %s", *t)
	}
	return nil
}

const (
	UpConfidentialityRequired  UpConfidentiality = "REQUIRED"   // UP confidentiality protection shall apply for all the traffic on the PDU Session.
	UpConfidentialityPreferred UpConfidentiality = "PREFERRED"  // UP confidentiality protection should apply for all the traffic on the PDU Session.
	UpConfidentialityNotNeeded UpConfidentiality = "NOT_NEEDED" // UP confidentiality protection shall not apply on the PDU Session.
)

// SscMode represents the service and session continuity mode.
type SscMode string

// Validate validates the SscMode string.
func (t *SscMode) Validate() error {
	if t == nil {
		return fmt.Errorf("sscMode must not be nil")
	}
	switch *t {
	case SscMode1:
	case SscMode2:
	case SscMode3:
	default:
		return fmt.Errorf("invalid ssc mode: %s", *t)
	}
	return nil
}

const (
	SscMode1 SscMode = "SSC_MODE_1" // see 3GPP TS 23.501.
	SscMode2 SscMode = "SSC_MODE_2" // see 3GPP TS 23.501.
	SscMode3 SscMode = "SSC_MODE_3" // see 3GPP TS 23.501.
)

// DnaiChangeType represents the type of a DNAI change. A NF service consumer may subscribe to
// "EARLY", "LATE" or "EARLY_LATE" types of DNAI change. The types of observed DNAI change the SMF may
// notify are "EARLY" or "LATE".
type DnaiChangeType string

// Validate validates the DnaiChangeType string.
func (t *DnaiChangeType) Validate() error {
	if t == nil {
		return fmt.Errorf("dnaiChangeType must not be nil")
	}
	switch *t {
	case DnaiChangeTypeEarly:
	case DnaiChangeTypeLate:
	case DnaiChangeTypeEarlyLate:
	default:
		return fmt.Errorf("invalid dnai change type: %s", *t)
	}
	return nil
}

const (
	DnaiChangeTypeEarly     DnaiChangeType = "EARLY"      // Early notification of UP path reconfiguration.
	DnaiChangeTypeEarlyLate DnaiChangeType = "EARLY_LATE" // Early and late notification of UP path reconfiguration. This value shall only be present in the subscription to the DNAI change event.
	DnaiChangeTypeLate      DnaiChangeType = "LATE"       // Late notification of UP path reconfiguration.
)

type RestrictionType string

// Validate validates the RestrictionType string.
func (t *RestrictionType) Validate() error {
	if t == nil {
		return fmt.Errorf("restrictionType must not be nil")
	}
	switch *t {
	case RestrictionTypeAllowed:
	case RestrictionTypeNotAllowed:
	default:
		return fmt.Errorf("invalid restriction type: %s", *t)
	}
	return nil
}

const (
	RestrictionTypeAllowed    RestrictionType = "ALLOWED_AREAS"     // This value indicates that areas are allowed.
	RestrictionTypeNotAllowed RestrictionType = "NOT_ALLOWED_AREAS" // This value indicates that areas are not allowed.
)

type CoreNetworkType string

// Validate validates the CoreNetworkType string.
func (t *CoreNetworkType) Validate() error {
	if t == nil {
		return fmt.Errorf("coreNetworkType must not be nil")
	}
	switch *t {
	case CoreNetworkType5GC:
	case CoreNetworkTypeEPC:
	default:
		return fmt.Errorf("invalid core network type: %s", *t)
	}
	return nil
}

const (
	CoreNetworkType5GC CoreNetworkType = "5GC" // 5G Core.
	CoreNetworkTypeEPC CoreNetworkType = "EPC" // Evolved Packet Core.
)

// AccessTypeRm is defined in the same way as the AccessType enumeration, but with the OpenAPI "nullable: true" property.
type AccessTypeRm AccessType

// Validate validates the AccessTypeRm string.
func (t *AccessTypeRm) Validate() error {
	if t == nil {
		return nil
	}
	return (*AccessType)(t).Validate()
}

// RatTypeRm is defined in the same way as the RatType enumeration, but with the OpenAPI "nullable: true" property.
type RatTypeRm RatType

// Validate validates the RatTypeRm string.
func (t *RatTypeRm) Validate() error {
	if t == nil {
		return nil
	}
	return (*RatType)(t).Validate()
}

// PduSessionTypeRm is defined in the same way as the PduSessionType enumeration, but with the OpenAPI "nullable: true" property.
type PduSessionTypeRm PduSessionType

// Validate validates the PduSessionTypeRm string.
func (t *PduSessionTypeRm) Validate() error {
	if t == nil {
		return nil
	}
	return (*PduSessionType)(t).Validate()
}

// UpIntegrityRm is defined in the same way as the UpIntegrity enumeration, but with the OpenAPI "nullable: true" property.
type UpIntegrityRm UpIntegrity

// Validate validates the UpIntegrityRm string.
func (t *UpIntegrityRm) Validate() error {
	if t == nil {
		return nil
	}
	return (*UpIntegrity)(t).Validate()
}

// UpConfidentialityRm is defined in the same way as the UpConfidentiality enumeration, but with the OpenAPI "nullable: true" property.
type UpConfidentialityRm UpConfidentiality

// Validate validates the UpConfidentialityRm string.
func (t *UpConfidentialityRm) Validate() error {
	if t == nil {
		return nil
	}
	return (*UpConfidentiality)(t).Validate()
}

// SscModeRm is defined in the same way as the SscMode enumeration, but with the OpenAPI "nullable: true" property.
type SscModeRm SscMode

// Validate validates the SscModeRm string.
func (t *SscModeRm) Validate() error {
	if t == nil {
		return nil
	}
	return (*SscMode)(t).Validate()
}

// DnaiChangeTypeRm is defined in the same way as the DnaiChangeType enumeration, but with the OpenAPI "nullable: true" property.
type DnaiChangeTypeRm DnaiChangeType

// Validate validates the DnaiChangeTypeRm string.
func (t *DnaiChangeTypeRm) Validate() error {
	if t == nil {
		return nil
	}
	return (*DnaiChangeType)(t).Validate()
}

// RestrictionTypeRm is defined in the same way as the RestrictionType enumeration, but with the OpenAPI "nullable: true" property.
type RestrictionTypeRm RestrictionType

// Validate validates the RestrictionTypeRm string.
func (t *RestrictionTypeRm) Validate() error {
	if t == nil {
		return nil
	}
	return (*RestrictionType)(t).Validate()
}

// CoreNetworkTypeRm is defined in the same way as the CoreNetworkType enumeration, but with the OpenAPI "nullable: true" property.
type CoreNetworkTypeRm CoreNetworkType

// Validate validates the CoreNetworkTypeRm string.
func (t *CoreNetworkTypeRm) Validate() error {
	if t == nil {
		return nil
	}
	return (*CoreNetworkType)(t).Validate()
}

type PresenceState string

// Validate validates the PresenceState string.
func (t *PresenceState) Validate() error {
	if t == nil {
		return fmt.Errorf("presenceState must not be nil")
	}
	switch *t {
	case PresenceStateInArea:
	case PresenceStateOutOfArea:
	case PresenceStateUnknown:
	case PresenceStateInactive:
	default:
		return fmt.Errorf("invalid presence state: %s", *t)
	}
	return nil
}

const (
	PresenceStateInArea    PresenceState = "IN_AREA"     // Indicates that the UE is inside or enters the presence reporting area.
	PresenceStateOutOfArea PresenceState = "OUT_OF_AREA" // Indicates that the UE is outside or leaves the presence reporting area.
	PresenceStateUnknown   PresenceState = "UNKNOWN"     // Indicates it is unknown whether the UE is in the presence reporting area or not.
	PresenceStateInactive  PresenceState = "INACTIVE"    // Indicates that the presence reporting area is inactive in the serving node.
)
