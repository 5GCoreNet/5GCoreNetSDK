package barring

import "fmt"

// RoamingOdb defines the Barring of Roaming as. See 3GPP TS 23.015 for further description.
type RoamingOdb string

// Validate validates the RoamingOdb string.
func (r *RoamingOdb) Validate() error {
	if r == nil {
		return fmt.Errorf("roamingOdb must not be nil")
	}
	switch *r {
	case RoamingOdbOutsideHomePlmn:
	case RoamingOdbOutsideHomePlmnCountry:
	default:
		return fmt.Errorf("invalid roaming odb: %s", r)
	}
	return nil
}

const (
	RoamingOdbOutsideHomePlmn        RoamingOdb = "OUTSIDE_HOME_PLMN"         // Barring of roaming outside the home PLMN.
	RoamingOdbOutsideHomePlmnCountry RoamingOdb = "OUTSIDE_HOME_PLMN_COUNTRY" // Barring of roaming outside the home PLMN country.
)

// OdbPacketServices The enumeration OdbPacketServices defines the Barring of Packet Oriented Services. See 3GPP TS 23.015 for further description.
type OdbPacketServices string

// Validate validates the OdbPacketServices string.
func (o *OdbPacketServices) Validate() error {
	if o == nil {
		return fmt.Errorf("odbPacketServices must not be nil")
	}
	switch *o {
	case OdbPacketServicesAllPacketServices:
	case OdbPacketServicesRoamerAccessHplmnAp:
	case OdbPacketServicesRoamerAccessVplmnAp:
	default:
		return fmt.Errorf("invalid odb packet services: %s", o)
	}
	return nil

}

const (
	OdbPacketServicesAllPacketServices   OdbPacketServices = "ALL_PACKET_SERVICES"    // Barring of all Packet Oriented Services.
	OdbPacketServicesRoamerAccessHplmnAp OdbPacketServices = "ROAMER_ACCESS_HPLMN_AP" // Barring of Packet Oriented Services from access points that are within the HPLMN whilst the subscriber is roaming in a VPLMN.
	OdbPacketServicesRoamerAccessVplmnAp OdbPacketServices = "ROAMER_ACCESS_VPLMN_AP" // Barring of Packet Oriented Services from access points that are within the roamed to VPLMN.
)
