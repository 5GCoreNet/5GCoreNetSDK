package network

import (
	"fmt"
	"github.com/5GCoreNet/5GCoreNetSDK/models/common"
	"github.com/5GCoreNet/5GCoreNetSDK/models/qos"
	"github.com/5GCoreNet/5GCoreNetSDK/models/subscription"
	"regexp"
)

type SubscribedDefaultQos struct {
	FiveQi        *qos.FiveQi              `json:"5qi"`                     // Default 5G QoS identifier.
	Arp           *qos.Arp                 `json:"arp"`                     // Default Allocation and Retention Priority see 3GPP TS 23.501 subclause 5.7.2.7.
	PriorityLevel *qos.FiveQiPriorityLevel `json:"priorityLevel,omitempty"` // Defines the 5QI Priority Level. When present, it contains the 5QI Priority Level value that overrides the standardized or pre-configured value as described in 3GPP TS 23.501.
}

// Validate validates this subscribed default qos.
func (m *SubscribedDefaultQos) Validate() error {
	if err := m.FiveQi.Validate(); err != nil {
		return fmt.Errorf("5qi %s", err.Error())
	}
	if err := m.Arp.Validate(); err != nil {
		return fmt.Errorf("arp %s", err.Error())
	}
	if m.PriorityLevel != nil {
		if err := m.PriorityLevel.Validate(); err != nil {
			return fmt.Errorf("priorityLevel %s", err.Error())
		}
	}
	return nil
}

// Snssai , When Snssai needs to be converted to string (e.g. when used in maps
// as key), the string shall be composed of one to three digits "sst" optionally
// followed by "-" and 6 hexadecimal digits "sd", and shall match the following
// pattern:
// ^([0-9]|[1-9][0-9]|1[0-9][0-9]|2([0-4][0-9]|5[0-5]))(-[A-Fa-f0-9]{6})?$
type Snssai struct {
	// Sst is an unsigned integer, within the range 0 to 255,
	// representing the Slice/Service Type. It indicates the
	// expected Network Slice behaviour in terms of
	// features and services.
	// Values 0 to 127 correspond to the standardized SST
	// range. Values 128 to 255 correspond to the
	// Operator-specific range. See subclause 28.4.2 of
	// 3GPP TS 23.003.
	// Standardized values are defined in
	// subclause 5.15.2.2 of 3GPP TS 23.501
	Sst *common.Uinteger `json:"sst"`
	// Sd is a 3-octet string, representing the Slice Differentiator, in
	// hexadecimal representation. Each character in the
	// string shall take a value of "0" to "9" or "A" to "F" and
	// shall represent 4 bits. The most significant character
	// representing the 4 most significant bits of the SD
	// shall appear first in the string, and the character
	// representing the 4 least significant bit of the SD shall
	// appear last in the string.
	// This is an optional parameter that complements the
	// Slice/Service type(s) to allow to differentiate
	// amongst multiple Network Slices of the same
	// Slice/Service type.
	Sd *string `json:"sd,omitempty"`
}

// Validate validates this snssai.
func (m *Snssai) Validate() error {
	if err := m.Sst.Validate(); err != nil {
		return fmt.Errorf("sst %s", err.Error())
	}
	if *m.Sst < 0 || *m.Sst > 255 {
		return fmt.Errorf("sst must be in the range 0 to 255")
	}
	if m.Sd != nil {
		if !regexp.MustCompile(`^[A-F\d]{6}$`).MatchString(*m.Sd) {
			return fmt.Errorf("sd must match the pattern [0-9A-F]{6}$")
		}
	}
	return nil
}

type PlmnId struct {
	Mcc *Mcc `json:"mcc"` // Mobile Country Code.
	Mnc *Mnc `json:"mnc"` // Mobile Network Code.
}

// Validate validates this plmn id.
func (m *PlmnId) Validate() error {
	if err := m.Mcc.Validate(); err != nil {
		return fmt.Errorf("mcc %s", err.Error())
	}
	if err := m.Mnc.Validate(); err != nil {
		return fmt.Errorf("mnc %s", err.Error())
	}
	return nil
}

type Tai struct {
	PlmnId *PlmnId `json:"plmnId"` // PLMN identity.
	Tac    *Tac    `json:"tac"`    // Tracking Area Code.
}

// Validate validates this tai.
func (m *Tai) Validate() error {
	if err := m.PlmnId.Validate(); err != nil {
		return fmt.Errorf("plmnId %s", err.Error())
	}
	if err := m.Tac.Validate(); err != nil {
		return fmt.Errorf("tac %s", err.Error())
	}
	return nil
}

type Ecgi struct {
	PlmnId      *PlmnId      `json:"plmnId"`      // PLMN identity.
	EutraCellId *EutraCellId `json:"eutraCellId"` // E-UTRA cell identity.
}

// Validate validates this ecgi.
func (m *Ecgi) Validate() error {
	if err := m.PlmnId.Validate(); err != nil {
		return fmt.Errorf("plmnId %s", err.Error())
	}
	if err := m.EutraCellId.Validate(); err != nil {
		return fmt.Errorf("eutraCellId %s", err.Error())
	}
	return nil
}

type Ncgi struct {
	PlmnId   *PlmnId   `json:"plmnId"`   // PLMN identity.
	NrCellId *NrCellId `json:"nrCellId"` // NR cell identity.
}

// Validate validates this ncgi.
func (m *Ncgi) Validate() error {
	if err := m.PlmnId.Validate(); err != nil {
		return fmt.Errorf("plmnId %s", err.Error())
	}
	if err := m.NrCellId.Validate(); err != nil {
		return fmt.Errorf("nrCellId %s", err.Error())
	}
	return nil
}

type UserLocation struct {
	EutraLocation *EutraLocation `json:"eutraLocation,omitempty"` // E-UTRA user location (see NOTE).
	NrLocation    *NrLocation    `json:"nrLocation,omitempty"`    // NR user location (see NOTE).
	N3gaLocation  *N3gaLocation  `json:"n3gaLocation,omitempty"`  // Non-3GPP access user location (see NOTE).
	// NOTE: At least one of eutraLocation, nrLocation and n3gaLocation shall be present. Several of them may be
	// present.
}

// Validate validates this user location.
func (m *UserLocation) Validate() error {
	if m.EutraLocation == nil && m.NrLocation == nil && m.N3gaLocation == nil {
		return fmt.Errorf("at least one of eutraLocation, nrLocation and n3gaLocation shall be present")
	}
	if err := m.EutraLocation.Validate(); err != nil {
		return fmt.Errorf("eutraLocation %s", err.Error())
	}
	if err := m.NrLocation.Validate(); err != nil {
		return fmt.Errorf("nrLocation %s", err.Error())
	}
	if err := m.N3gaLocation.Validate(); err != nil {
		return fmt.Errorf("n3gaLocation %s", err.Error())
	}
	return nil
}

type EutraLocation struct {
	Tai  *Tai  `json:"tai"`  // Tracking Area Identity.
	Ecgi *Ecgi `json:"ecgi"` // E-UTRA Cell Identity.
	// AgeOfLocationInformation represents the elapsed time in minutes since the last network contact of the mobile station.
	// Value "0" indicates that the location information was obtained after a successful paging procedure for
	// Active Location Retrieval when the UE is in idle mode or after a successful NG-RAN location
	// reporting procedure with the eNB when the UE is in connected mode.
	// Any other value than "0" indicates that the location
	// information is the last known one.
	// See 3GPP TS 29.002 [21] subclause 17.7.8.
	AgeOfLocationInformation *common.Int64    `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      *common.DateTime `json:"ueLocationTimestamp,omitempty"` // The value represents the UTC time when the UeLocation information was acquired.
	// GeographicalInformation refers to geographical Information.
	// See 3GPP TS 23.032 subclause 7.3.2. Only the
	// description of an ellipsoid point with uncertainty circle
	// is allowed to be used.
	// Allowed characters are 0-9 and A-F.
	GeographicalInformation *string `json:"geographicalInformation,omitempty"`
	// GeodeticInformation refers to Calling Geodetic Location.
	// See ITU-T Recommendation Q.763 (1999)
	// subclause 3.88.2. Only the description of an ellipsoid
	// point with uncertainty circle is allowed to be used.
	// Allowed characters are 0-9 and A-F.
	GeodeticInformation *string `json:"geodeticInformation,omitempty"`
	// GlobalNgenbId indicates the global identity of the ng-eNodeB in
	// which the UE is currently located.
	// See 3GPP TS 38.413 subclause 9.3.1.8.
	GlobalNgenbId *GlobalRanNodeId `json:"globalNgenbId,omitempty"`
}

// Validate validates this eutra location.
func (m *EutraLocation) Validate() error {
	if err := m.Tai.Validate(); err != nil {
		return fmt.Errorf("tai %s", err.Error())
	}
	if err := m.Ecgi.Validate(); err != nil {
		return fmt.Errorf("ecgi %s", err.Error())
	}
	if m.AgeOfLocationInformation != nil {
		if err := m.AgeOfLocationInformation.Validate(); err != nil {
			return fmt.Errorf("ageOfLocationInformation %s", err.Error())
		}
	}
	if m.UeLocationTimestamp != nil {
		if err := m.UeLocationTimestamp.Validate(); err != nil {
			return fmt.Errorf("ueLocationTimestamp %s", err.Error())
		}
	}
	if m.GeographicalInformation != nil {
		if !regexp.MustCompile(`^[\dA-F]+$`).MatchString(*m.GeographicalInformation) {
			return fmt.Errorf("geographicalInformation must match the pattern `^[0-9A-F]+$`")
		}
	}
	if m.GeodeticInformation != nil {
		if !regexp.MustCompile(`^[\dA-F]+$`).MatchString(*m.GeodeticInformation) {
			return fmt.Errorf("geodeticInformation must match the pattern `^[0-9A-F]+$`")
		}
	}
	if m.GlobalNgenbId != nil {
		if err := m.GlobalNgenbId.Validate(); err != nil {
			return fmt.Errorf("globalNgenbId %s", err.Error())
		}
	}
	return nil
}

type NrLocation struct {
	Tai  *Tai  `json:"tai"`  // Tracking Area Identity.
	Ncgi *Ncgi `json:"ncgi"` // NR Cell Identity.
	// AgeOfLocationInformation represents the elapsed time in minutes since the last network contact of the mobile station.
	// Value "0" indicates that the location information was obtained after a successful paging procedure for
	// Active Location Retrieval when the UE is in idle mode or after a successful NG-RAN location
	// reporting procedure with the gNB when the UE is in connected mode.
	// Any other value than "0" indicates that the location
	// information is the last known one.
	// See 3GPP TS 29.002 [21] subclause 17.7.8.
	AgeOfLocationInformation *common.Int64    `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      *common.DateTime `json:"ueLocationTimestamp,omitempty"` // The value represents the UTC time when the UeLocation information was acquired.
	// GeographicalInformation refers to geographical Information.
	// See 3GPP TS 23.032 subclause 7.3.2. Only the
	// description of an ellipsoid point with uncertainty circle
	// is allowed to be used.
	// Allowed characters are 0-9 and A-F.
	GeographicalInformation *string `json:"geographicalInformation,omitempty"`
	// GeodeticInformation refers to Calling Geodetic Location.
	// See ITU-T Recommendation Q.763 (1999)
	// subclause 3.88.2. Only the description of an ellipsoid
	// point with uncertainty circle is allowed to be used.
	// Allowed characters are 0-9 and A-F.
	GeodeticInformation *string `json:"geodeticInformation,omitempty"`
	// GlobalGnbId indicates the global identity of the gNodeB in
	// which the UE is currently located.
	// See 3GPP TS 38.413 subclause 9.3.1.8.
	GlobalGnbId *GlobalRanNodeId `json:"globalGnbId,omitempty"`
}

// Validate validates this nr location.
func (m *NrLocation) Validate() error {
	if err := m.Tai.Validate(); err != nil {
		return fmt.Errorf("tai %s", err.Error())
	}
	if err := m.Ncgi.Validate(); err != nil {
		return fmt.Errorf("ncgi %s", err.Error())
	}
	if m.AgeOfLocationInformation != nil {
		if err := m.AgeOfLocationInformation.Validate(); err != nil {
			return fmt.Errorf("ageOfLocationInformation %s", err.Error())
		}
	}
	if m.UeLocationTimestamp != nil {
		if err := m.UeLocationTimestamp.Validate(); err != nil {
			return fmt.Errorf("ueLocationTimestamp %s", err.Error())
		}
	}
	if m.GeographicalInformation != nil {
		if !regexp.MustCompile(`^[\dA-F]+$`).MatchString(*m.GeographicalInformation) {
			return fmt.Errorf("geographicalInformation must match the pattern `^[0-9A-F]+$`")
		}
	}
	if m.GeodeticInformation != nil {
		if !regexp.MustCompile(`^[\dA-F]+$`).MatchString(*m.GeodeticInformation) {
			return fmt.Errorf("geodeticInformation must match the pattern `^[0-9A-F]+$`")
		}
	}
	if m.GlobalGnbId != nil {
		if err := m.GlobalGnbId.Validate(); err != nil {
			return fmt.Errorf("globalGnbId %s", err.Error())
		}
	}
	return nil
}

type N3gaLocation struct {
	// N3gppTai is the unique non 3GPP TAI used in the PLMN. It shall be present over
	// the 3GPP PLMN internal interfaces, but shall not be present over the N5
	// interface.
	N3gppTai *Tai `json:"n3gppTai,omitempty"`
	// N3IwfId shall contain the N3IWF identifier received over NGAP and shall be
	// encoded as a string of hexadecimal characters. Pattern: '^[A-Fa-f0-9]+$' It
	// shall be present over the 3GPP PLMN internal interfaces, but shall not be
	// present over the N5 interface
	N3IwfId *N3IwfId `json:"n3IwfId,omitempty"`
	// UeIpv4Addr is the UE local IPv4 address (used to reach the N3IWF). The
	// ueIPv4Addr or the ueIPv6Addr shall be present.
	UeIpv4Addr *common.Ipv4Addr `json:"ueIpv4Addr,omitempty"`
	// UeIpv6Addr is the UE local IPv6 address (used to reach the N3IWF). The
	// ueIPv4Addr or the ueIPv6Addr shall be present.
	UeIpv6Addr *common.Ipv6Addr `json:"ueIpv6Addr,omitempty"`
	// PortNumber is the UDP or TCP source port number. It shall be present if NAT is
	// detected.
	PortNumber *common.Uinteger `json:"portNumber,omitempty"`
}

// Validate validates this n3ga location.
func (m *N3gaLocation) Validate() error {
	if m.N3gppTai != nil {
		if err := m.N3gppTai.Validate(); err != nil {
			return fmt.Errorf("n3gppTai %s", err.Error())
		}
	}
	if m.N3IwfId != nil {
		if err := m.N3IwfId.Validate(); err != nil {
			return fmt.Errorf("n3IwfId %s", err.Error())
		}
	}
	if m.UeIpv4Addr == nil && m.UeIpv6Addr == nil {
		return fmt.Errorf("ueIpv4Addr or ueIpv6Addr must be present")
	}
	if m.UeIpv4Addr != nil {
		if err := m.UeIpv4Addr.Validate(); err != nil {
			return fmt.Errorf("ueIpv4Addr %s", err.Error())
		}
	}
	if m.UeIpv6Addr != nil {
		if err := m.UeIpv6Addr.Validate(); err != nil {
			return fmt.Errorf("ueIpv6Addr %s", err.Error())
		}
	}
	if m.PortNumber != nil {
		if err := m.PortNumber.Validate(); err != nil {
			return fmt.Errorf("portNumber %s", err.Error())
		}
	}
	return nil
}

type UpSecurity struct {
	// UpIntegr shall indicate whether UP integrity protection is required, preferred
	// or not needed for all the traffic on the PDU Session.
	UpIntegr *UpIntegrity `json:"upIntegr"`
	// UpConfid This IE shall indicate whether UP confidentiality protection is
	// required, preferred or not needed for all the traffic on the PDU Session.
	UpConfid *UpConfidentiality `json:"upConfid"`
}

// Validate validates this up security.
func (m *UpSecurity) Validate() error {
	if err := m.UpIntegr.Validate(); err != nil {
		return fmt.Errorf("upIntegr %s", err.Error())
	}
	if err := m.UpConfid.Validate(); err != nil {
		return fmt.Errorf("upConfid %s", err.Error())
	}
	return nil
}

type NgApCause struct {
	// Group shall indicate the group of the NGAP cause. The value of this IE shall
	// equal to the ASN.1 value of the specified NGAP cause group. NGAP supports
	// following cause groups defined as separate enumerations, as specified in
	// subclause 9.4.5 of 3GPP TS 38.413 [11], with following values:
	//		0 – radioNetwork
	//		1 – transport
	//		2 – nas
	//		3 – protocol
	//		4 – misc
	Group *common.Uinteger `json:"group"`
	// Value shall carry the NG AP cause value in specific
	// cause group identified by the "group" attribute, as specified in subclause
	// 9.4.5 of 3GPP TS 38.413.
	Value *common.Uinteger `json:"value"`
}

// Validate validates this ng ap cause.
func (m *NgApCause) Validate() error {
	if *m.Group < 0 || *m.Group > 4 {
		return fmt.Errorf("group must be in the range [0, 4]")
	}
	if err := m.Group.Validate(); err != nil {
		return fmt.Errorf("group %s", err.Error())
	}
	if err := m.Value.Validate(); err != nil {
		return fmt.Errorf("value %s", err.Error())
	}
	return nil
}

type BackupAmfInfo struct {
	// BackupAmf shall contain the AMF name of the backup
	// AMF related to the specific GUAMI(s) (see
	// subclause 5.21.2.3 of 3GPP TS 23.501). If no
	// GUAMI is included in BackupAmfinfo, the AMF name
	// of the backup AMF is related to all the GUAMI(s)
	// supported by the AMF.
	BackupAmf *AmfName `json:"backupAmf"`
	// If present, this IE shall contain the GUAMI(s).
	GuamiList []*subscription.Guami `json:"guamiList,omitempty"`
}

// Validate validates this backup amf info.
func (m *BackupAmfInfo) Validate() error {
	if err := m.BackupAmf.Validate(); err != nil {
		return fmt.Errorf("backupAmf %s", err.Error())
	}
	for i, v := range m.GuamiList {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("guamiList[%d] %s", i, err.Error())
		}
	}
	return nil
}

type RefToBinaryData struct {
	ContentId *string `json:"contentId"` // This IE shall contain the value of the Content-ID header of the referenced binary body part.
}

// Validate validates this ref to binary data.
func (m *RefToBinaryData) Validate() error {
	if m.ContentId != nil {
		return nil
	}
	return fmt.Errorf("contentId must be present")
}

type RouteToLocation struct {
	Dnai        *Dnai             `json:"dnai"`                  // Identifies the location of the application.
	RouteInfo   *RouteInformation `json:"routeInfo,omitempty"`   // Includes the traffic routing information.
	RouteProfId *string           `json:"routeProfId,omitempty"` // Identifies the routing profile Id.
	// NOTE: Either the "routeInfo" attribute or the "routeProfId" attribute shall be included in the
	// "RouteToLocation" data type.
}

// Validate validates this route to location.
func (m *RouteToLocation) Validate() error {
	if err := m.Dnai.Validate(); err != nil {
		return fmt.Errorf("dnai %s", err.Error())
	}
	if m.RouteInfo == nil && m.RouteProfId == nil {
		return fmt.Errorf("routeInfo or routeProfId must be present")
	}
	if m.RouteInfo != nil && m.RouteProfId != nil {
		return fmt.Errorf("routeInfo and routeProfId must not be present at the same time")
	}
	if m.RouteInfo != nil {
		if err := m.RouteInfo.Validate(); err != nil {
			return fmt.Errorf("routeInfo %s", err.Error())
		}
	}
	if m.RouteProfId != nil {
		return nil
	}
	return fmt.Errorf("routeProfId must be present")
}

type RouteInformation struct {
	Ipv4Addr   *common.Ipv4Addr `json:"ipv4Addr,omitempty"` // Ipv4address of the tunnel end point in the data network.
	Ipv6Addr   *common.Ipv6Addr `json:"ipv6Addr,omitempty"` // Ipv6 address of the tunnel end point in the data network.
	PortNumber *common.Uinteger `json:"portNumber"`         // UDP port number of the tunnel end point in the data network.
	// NOTE: At least one of the "ipv4Addr" attribute and the "ipv6Addr" attribute shall be included in the
	// "RouteInformation" data type
}

// Validate validates this route information.
func (m *RouteInformation) Validate() error {
	if m.Ipv4Addr == nil && m.Ipv6Addr == nil {
		return fmt.Errorf("ipv4Addr or ipv6Addr must be present")
	}
	if m.Ipv4Addr != nil {
		if err := m.Ipv4Addr.Validate(); err != nil {
			return fmt.Errorf("ipv4Addr %s", err.Error())
		}
	}
	if m.Ipv6Addr != nil {
		if err := m.Ipv6Addr.Validate(); err != nil {
			return fmt.Errorf("ipv6Addr %s", err.Error())
		}
	}
	if err := m.PortNumber.Validate(); err != nil {
		return fmt.Errorf("portNumber %s", err.Error())
	}
	return nil
}

type Area struct {
	Tacs      []*Tac    `json:"tacs,omitempty"`      // List of TACs; shall be present if and only if areaCode is absent.
	AreaCodes *AreaCode `json:"areaCodes,omitempty"` // Area Code; shall be present if and only if tacs is absent.
}

// Validate validates this area.
func (m *Area) Validate() error {
	if m.Tacs == nil && m.AreaCodes == nil {
		return fmt.Errorf("tacs or areaCodes must be present")
	}
	if m.Tacs != nil && m.AreaCodes != nil {
		return fmt.Errorf("tacs and areaCodes must not be present at the same time")
	}
	for i, v := range m.Tacs {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("tacs[%d] %s", i, err.Error())
		}
	}
	if m.AreaCodes != nil {
		if err := m.AreaCodes.Validate(); err != nil {
			return fmt.Errorf("areaCodes %s", err.Error())
		}
	}
	return nil
}

type ServiceAreaRestriction struct {
	// RestrictionType is a string. "ALLOWED_AREAS" or "NOT_ALLOWED_AREAS" shall be
	// present if and only if the areas attribute is present
	RestrictionType *RestrictionType `json:"restrictionType,omitempty"`
	// Areas is a list of Area.
	// These areas are:
	//     - allowed areas if RestrictionType is
	//           "ALLOWED_AREAS"
	//     - not allowed areas if RestrictionType is
	//           "NOT_ALLOWED_AREAS"
	Areas []*Area `json:"areas,omitempty"`
	// MaxNumOfTAs is the Maximum number of allowed tracking areas.
	// This attribute shall be absent when attribute "restrictionType" takes the
	// value "NOT_ALLOWED_AREAS".
	MaxNumOfTAs *common.Uinteger `json:"maxNumOfTAs,omitempty"`
	// NOTE: The empty Area array is used when service is allowed/restricted nowhere.
}

// Validate validates this service area restriction.
func (m *ServiceAreaRestriction) Validate() error {
	if m.RestrictionType == nil && m.Areas == nil && m.MaxNumOfTAs == nil {
		return nil
	}
	if m.Areas == nil {
		return nil
	}
	if m.Areas != nil {
		if m.RestrictionType == nil {
			return fmt.Errorf("restrictionType must be present")
		}
		if *m.RestrictionType == RestrictionTypeNotAllowed && m.MaxNumOfTAs != nil {
			return fmt.Errorf("maxNumOfTAs must be absent")
		}
	}
	if m.RestrictionType != nil {
		if err := m.RestrictionType.Validate(); err != nil {
			return fmt.Errorf("restrictionType %s", err.Error())
		}
	}
	if m.MaxNumOfTAs != nil {
		if err := m.MaxNumOfTAs.Validate(); err != nil {
			return fmt.Errorf("maxNumOfTAs %s", err.Error())
		}
	}
	for i, v := range m.Areas {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("areas[%d] %s", i, err.Error())
		}
	}
	return nil
}

// PlmnIdRm is defined in the same way as the "PlmnId" data type, but with the OpenAPI "nullable: true" property.
type PlmnIdRm PlmnId

// Validate validates this plmn id rm.
func (m *PlmnIdRm) Validate() error {
	if m == nil {
		return nil
	}
	if err := (*PlmnId)(m).Validate(); err != nil {
		return fmt.Errorf("plmnId %s", err.Error())
	}
	return nil
}

// TaiRm is defined in the same way as the "Tai" data type, but with the OpenAPI "nullable: true" property.
type TaiRm Tai

// Validate validates this tai rm.
func (m *TaiRm) Validate() error {
	if m == nil {
		return nil
	}
	if err := (*Tai)(m).Validate(); err != nil {
		return fmt.Errorf("tai %s", err.Error())
	}
	return nil
}

// EcgiRm is defined in the same way as the "Ecgi" data type, but with the OpenAPI "nullable: true" property.
type EcgiRm Ecgi

// Validate validates this ecgi rm.
func (m *EcgiRm) Validate() error {
	if m == nil {
		return nil
	}
	if err := (*Ecgi)(m).Validate(); err != nil {
		return fmt.Errorf("ecgi %s", err.Error())
	}
	return nil
}

// NcgiRm is defined in the same way as the "Ncgi" data type, but with the OpenAPI "nullable: true" property.
type NcgiRm Ncgi

// Validate validates this ncgi rm.
func (m *NcgiRm) Validate() error {
	if m == nil {
		return nil
	}
	if err := (*Ncgi)(m).Validate(); err != nil {
		return fmt.Errorf("ncgi %s", err.Error())
	}
	return nil
}

// EutraLocationRm is defined in the same way as the "EutraLocation" data type, but with the OpenAPI "nullable: true" property.
type EutraLocationRm EutraLocation

// Validate validates this eutra location rm.
func (m *EutraLocationRm) Validate() error {
	if m == nil {
		return nil
	}
	if err := (*EutraLocation)(m).Validate(); err != nil {
		return fmt.Errorf("eutraLocation %s", err.Error())
	}
	return nil
}

// NrLocationRm is defined in the same way as the "NrLocation" data type, but with the OpenAPI "nullable: true" property.
type NrLocationRm NrLocation

// Validate validates this nr location rm.
func (m *NrLocationRm) Validate() error {
	if m == nil {
		return nil
	}
	if err := (*NrLocation)(m).Validate(); err != nil {
		return fmt.Errorf("nrLocation %s", err.Error())
	}
	return nil
}

// UpSecurityRm is defined in the same way as the "UpSecurity" data type, but with the OpenAPI "nullable: true" property.
type UpSecurityRm UpSecurity

// Validate validates this up security rm.
func (m *UpSecurityRm) Validate() error {
	if m == nil {
		return nil
	}
	if err := (*UpSecurity)(m).Validate(); err != nil {
		return fmt.Errorf("upSecurity %s", err.Error())
	}
	return nil
}

// RefToBinaryDataRm is defined in the same way as the "RefToBinaryData" data type, but with the OpenAPI "nullable: true" property.
type RefToBinaryDataRm RefToBinaryData

// Validate validates this ref to binary data rm.
func (m *RefToBinaryDataRm) Validate() error {
	if m == nil {
		return nil
	}
	if err := (*RefToBinaryData)(m).Validate(); err != nil {
		return fmt.Errorf("refToBinaryData %s", err.Error())
	}
	return nil
}

type PresenceInfo struct {
	// PraId represents an identifier to the specified area. This
	// IE shall be present if the Area of Interest subscribed
	// or reported is a Presence Reporting Area.
	PraId *string `json:"praId,omitempty"`
	// PresenceState indicates whether the UE is inside or outside of the
	// area of interest (e.g. presence reporting area or the
	// LADN area), or if the presence reporting area is
	// inactive in the serving node.
	PresenceState *PresenceState `json:"presenceState,omitempty"`
	// TrackingAreaList represents the list of tracking areas that constitutes
	// the area. This IE shall be present if the subscription
	// or the event report is for tracking UE presence in the
	// tracking areas. For non 3GPP access the TAI shall
	// be the N3GPP TAI.
	TrackingAreaList []*Tai `json:"trackingAreaList,omitempty"`
	// EcgiList represents the list of EUTRAN cell Ids that
	// constitutes the area. This IE shall be present if the
	// Area of Interest subscribed is a list of EUTRAN cell
	// Ids.
	EcgiList []*Ecgi `json:"ecgiList,omitempty"`
	// NcgiList the list of NR cell Ids that constitutes the
	// area. This IE shall be present if the Area of Interest
	// subscribed is a list of NR cell Ids.
	NcgiList []*Ncgi `json:"ncgiList,omitempty"`
	// GlobalRanNodeIdList represents the list of NG RAN node identifiers that
	//constitutes the area. This IE shall be present if the
	//Area of Interest subscribed is a list of NG RAN node identifiers.
	GlobalRanNodeIdList []*GlobalRanNodeId `json:"globalRanNodeIdList,omitempty"`
}

// Validate validates this presence info.
func (m *PresenceInfo) Validate() error {
	if m == nil {
		return nil
	}
	if m.PresenceState != nil {
		if err := m.PresenceState.Validate(); err != nil {
			return fmt.Errorf("presenceState %s", err.Error())
		}
	}
	if m.TrackingAreaList != nil {
		for i, v := range m.TrackingAreaList {
			if err := v.Validate(); err != nil {
				return fmt.Errorf("trackingAreaList[%d] %s", i, err.Error())
			}
		}
	}
	if m.EcgiList != nil {
		for i, v := range m.EcgiList {
			if err := v.Validate(); err != nil {
				return fmt.Errorf("ecgiList[%d] %s", i, err.Error())
			}
		}
	}
	if m.NcgiList != nil {
		for i, v := range m.NcgiList {
			if err := v.Validate(); err != nil {
				return fmt.Errorf("ncgiList[%d] %s", i, err.Error())
			}
		}
	}
	if m.GlobalRanNodeIdList != nil {
		for i, v := range m.GlobalRanNodeIdList {
			if err := v.Validate(); err != nil {
				return fmt.Errorf("globalRanNodeIdList[%d] %s", i, err.Error())
			}
		}
	}
	return nil
}

type GlobalRanNodeId struct {
	// PlmnId indicates the identity of the PLMN that the RAN node
	// belongs to.
	PlmnId *PlmnId `json:"plmnId"`
	// This IE shall be included if the RAN node belongs to
	// non 3GPP access (i.e a N3IWF).
	// (NOTE).
	N3IwfId *N3IwfId `json:"n3IwfId,omitempty"`
	// This IE shall be included if the RAN Node Id
	// represents a gNB. When present, this IE shall
	// contain the identifier of the gNB. (NOTE).
	GNbId *GnbId `json:"gNbId,omitempty"`
	// This IE shall be included if the RAN Node Id
	// represents a NG-eNB. When present, this IE shall
	// contain the identifier of an NG-eNB. (NOTE).
	NgeNbId *NgeNbId `json:"ngeNbId,omitempty"`
	// NOTE : At most one of the three attributes n3IwfId, gNbIdm, ngeNbId shall be present.
}

// Validate validates this global ran node id.
func (m *GlobalRanNodeId) Validate() error {
	if err := m.PlmnId.Validate(); err != nil {
		return fmt.Errorf("plmnId %s", err.Error())
	}
	if m.GNbId != nil && m.NgeNbId != nil || m.GNbId != nil && m.N3IwfId != nil || m.NgeNbId != nil && m.N3IwfId != nil {
		return fmt.Errorf("at most one of the three attributes n3IwfId, gNbIdm, ngeNbId shall be present")
	}
	if m.N3IwfId != nil {
		if err := m.N3IwfId.Validate(); err != nil {
			return fmt.Errorf("n3IwfId %s", err.Error())
		}
	}
	if m.GNbId != nil {
		if err := m.GNbId.Validate(); err != nil {
			return fmt.Errorf("gNbId %s", err.Error())
		}
	}
	if m.NgeNbId != nil {
		if err := m.NgeNbId.Validate(); err != nil {
			return fmt.Errorf("ngeNbId %s", err.Error())
		}
	}
	return nil
}

type GnbId struct {
	// BitLength is an unsigned integer representing the bit length of the gNB ID as
	// defined in subclause 9.3.1.6 of 3GPP TS 38.413, within the range 22 to 32
	BitLength *common.Int64 `json:"bitLength"`
	// GNbValue represents the identifier of the gNB.
	// The value of the gNB ID shall be encoded in hexadecimal
	// representation. Each character in the string shall take a value of "0" to "9"
	// or "A" to "F" and shall represent 4 bits. The padding 0 shall be added to make
	// multiple nibbles, the most significant character representing the padding 0 if
	// required together with the 4 most significant bits of the gNB ID shall appear
	// first in the string, and the character representing the 4 least significant
	// bit of the gNB ID shall appear last in the string.
	// Examples: A 30 bit value "382A3F47" indicates a gNB ID with value 0x382A3F47 A 22 bit value "2A3F47"
	// indicates a gNB ID with value 0x2A3F47
	GNbValue *string `json:"gNbValue"`
}

// Validate validates this gnb id.
func (m *GnbId) Validate() error {
	if err := m.BitLength.Validate(); err != nil {
		return fmt.Errorf("bitLength %s", err.Error())
	}
	if !regexp.MustCompile(`^[A-Fa-f\d]{6,8}$`).MatchString(*m.GNbValue) {
		return fmt.Errorf("gNbValue must match the pattern ^[A-Fa-f0-9]{6,8}$")
	}
	return nil
}

// PresenceInfoRm is defined in the same way as the "PresenceInfo" data type, but with the OpenAPI "nullable: true" property.
type PresenceInfoRm PresenceInfo

// Validate validates this presence info rm.
func (m *PresenceInfoRm) Validate() error {
	if m == nil {
		return nil
	}
	if err := (*PresenceInfo)(m).Validate(); err != nil {
		return fmt.Errorf("presenceInfo %s", err.Error())
	}
	return nil
}
