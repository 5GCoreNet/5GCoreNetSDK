package models

type SubscribedDefaultQos struct {
	FiveQi        FiveQi              `json:"5Qi,omitempty"`
	Arp           Arp                 `json:"arp,omitempty"`
	PriorityLevel FiveQiPriorityLevel `json:"priorityLevel,omitempty"`
}

type Snssai struct {
	Sst Uinteger `json:"sst,omitempty"`
	Sd  string   `json:"sd,omitempty"`
}

type PlmnId struct {
	Mcc Mcc `json:"mcc,omitempty"`
	Mnc Mnc `json:"mnc,omitempty"`
}

func (i PlmnId) Validate() error {
	if err := i.Mcc.Validate(); err != nil {
		return err
	}
	if err := i.Mnc.Validate(); err != nil {
		return err
	}
	return nil
}

type Tai struct {
	PlmnId PlmnId `json:"plmnId,omitempty"`
	Tac    Tac    `json:"tac,omitempty"`
}

type Ecgi struct {
	PlmnId      PlmnId      `json:"plmnId,omitempty"`
	EutraCellId EutraCellId `json:"eutraCellId,omitempty"`
}

type Ncgi struct {
	PlmnId   PlmnId   `json:"plmnId,omitempty"`
	NrCellId NrCellId `json:"nrCellId,omitempty"`
}

type UserLocation struct {
	EutraLocation EutraLocation `json:"eutraLocation,omitempty"`
	NrLocation    NrLocation    `json:"nrLocation,omitempty"`
	N3gaLocation  N3gaLocation  `json:"n3gaLocation,omitempty"`
}

type EutraLocation struct {
	Tai                      Tai             `json:"tai,omitempty"`
	Ecgi                     Ecgi            `json:"ecgi,omitempty"`
	AgeOfLocationInformation int64           `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      DateTime        `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation  string          `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string          `json:"geodeticInformation,omitempty"`
	GlobalNgenbId            GlobalRanNodeId `json:"globalNgenbId,omitempty"`
}

type NrLocation struct {
	Tai                      Tai             `json:"tai,omitempty"`
	Ncgi                     Ncgi            `json:"ncgi,omitempty"`
	AgeOfLocationInformation int64           `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      DateTime        `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation  string          `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string          `json:"geodeticInformation,omitempty"`
	GlobalGnbId              GlobalRanNodeId `json:"globalNgenbId,omitempty"`
}

type N3gaLocation struct {
	N3gpppTai  Tai         `json:"n3gppTai,omitempty"`
	N3IwfId    N3IwfId     `json:"n3IwfId,omitempty"`
	UeIpv4Addr Ipv4Address `json:"Ipv4Addr,omitempty"`
	UeIpv6Addr Ipv6Address `json:"Ipv6Addr,omitempty"`
	PortNumber Uinteger    `json:"portNumber,omitempty"`
}

type UpSecurity struct {
	UpIntegr UpIntegrity       `json:"upIntegr,omitempty"`
	UpConfid UpConfidentiality `json:"upConfid,omitempty"`
}

type NgApCause struct {
	Group Uinteger `json:"group,omitempty"`
	Value Uinteger `json:"value,omitempty"`
}

type BackupAmfInfo struct {
	BackupAmf AmfName `json:"backupAmf,omitempty"`
	GuamiList []Guami `json:"guamiList,omitempty"`
}

type RefToBinaryData struct {
	ContentId string `json:"contentId,omitempty"`
}

type RouteToLocation struct {
	Dnai        Dnai             `json:"dnai,omitempty"`
	RouteInfo   RouteInformation `json:"routeInfo,omitempty"`
	RouteProfId string           `json:"routeProfId,omitempty"`
}

type RouteInformation struct {
	Ipv4Addr   Ipv4Address `json:"ipv4Addr,omitempty"`
	Ipv6Addr   Ipv6Address `json:"ipv6Addr,omitempty"`
	PortNumber Uinteger    `json:"portNumber,omitempty"`
}

type Area struct {
	Tacs      []Tac    `json:"tacs,omitempty"`
	AreaCodes AreaCode `json:"areaCodes,omitempty"`
}

type ServiceAreaRestriction struct {
	RestrictionType RestrictionType `json:"restrictionType,omitempty"`
	Areas           []Area          `json:"areas,omitempty"`
	MaxNumOfTAs     Uinteger        `json:"maxNumOfTAs,omitempty"`
}

type PlmnIdRm PlmnId
type EcgiRm Ecgi
type NcgiRm Ncgi
type EutraLocationRm EutraLocation
type NrLocationRm NrLocation
type UpSecurityRm UpSecurity
type RefToBinaryDataRm RefToBinaryData

type PresenceInfo struct {
	PraId               string            `json:"praId,omitempty"`
	PresenceState       PresenceState     `json:"presenceState,omitempty"`
	TrackingAreaList    []Tai             `json:"trackingAreaList,omitempty"`
	EcgiList            []Ecgi            `json:"ecgiList,omitempty"`
	NcgiList            []Ncgi            `json:"ncgiList,omitempty"`
	GlobalRanNodeIdList []GlobalRanNodeId `json:"globalRanNodeIdList,omitempty"`
}

type GlobalRanNodeId struct {
	PlmnId  PlmnId  `json:"plmnId,omitempty"`
	N3IwfId N3IwfId `json:"n3LwfId,omitempty"`
	GNbId   GnbId   `json:"gNbId,omitempty"`
	NgeNbId NgeNbId `json:"ngeNbId,omitempty"`
}

type GnbId struct {
	BitLength Int64  `json:"bitLength,omitempty"`
	GNbValue  string `json:"gNbValue,omitempty"`
}

type PresenceInfoRm PresenceInfo