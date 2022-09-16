package models

import "fmt"

// TODO: Add description regarding the ETSI documentation

type AccessType string

const (
	AccessTypeNonThreeGPPAccess AccessType = "NON_3_GPP_ACCESS"
	AccessTypeThreeGPPAccess    AccessType = "3GPP_ACCESS"
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
	RatTypeNR      RatType = "NR"
	RatTypeEUTRA   RatType = "EUTRA"
	RatTypeWLAN    RatType = "WLAN"
	RatTypeVIRTUAL RatType = "VIRTUAL"
)

type PduSessionType string

const (
	PduSessionTypeIpv4         PduSessionType = "IPV4"
	PduSessionTypeIpv6         PduSessionType = "IPV6"
	PduSessionTypeIpv4v6       PduSessionType = "IPV4V6"
	PduSessionTypeUnstructured PduSessionType = "UNSTRUCTURED"
	PduSessionTypeEthernet     PduSessionType = "ETHERNET"
)

type UpIntegrity string

const (
	UpIntegrityRequired  UpIntegrity = "REQUIRED"
	UpIntegrityPreferred UpIntegrity = "PREFERRED"
	UpIntegrityNotNeeded UpIntegrity = "NOT_NEEDED"
)

type UpConfidentiality string

const (
	UpConfidentialityRequired  UpConfidentiality = "REQUIRED"
	UpConfidentialityPreferred UpConfidentiality = "PREFERRED"
	UpConfidentialityNotNeeded UpConfidentiality = "NOT_NEEDED"
)

type SscMode string

const (
	SscMode1 SscMode = "SSC_MODE_1"
	SscMode2 SscMode = "SSC_MODE_2"
	SscMode3 SscMode = "SSC_MODE_3"
)

type DnaiChangeType string

const (
	DnaiChangeTypeEarly     DnaiChangeType = "EARLY"
	DnaiChangeTypeEarlyLate DnaiChangeType = "EARLY_LATE"
	DnaiChangeTypeLate      DnaiChangeType = "LATE"
)

type RestrictionType string

const (
	RestrictionTypeAllowed    RestrictionType = "ALLOWED_AREAS"
	RestrictionTypeNotAllowed RestrictionType = "NOT_ALLOWED_AREAS"
)

type CoreNetworkType string

const (
	CoreNetworkType5GC CoreNetworkType = "5GC"
	CoreNetworkTypeEPC CoreNetworkType = "EPC"
)

// TODO: Implement nullable enums for all enum in this file

type PresenceState string

const (
	PresenceStateInArea    PresenceState = "IN_AREA"
	PresenceStateOutOfArea PresenceState = "OUT_OF_AREA"
	PresenceStateUnknown   PresenceState = "UNKNOWN"
	PresenceStateInactive  PresenceState = "INACTIVE"
)
