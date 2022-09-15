package models

// TODO: Add description regarding the ETSI documentation
// TODO: Add validation functions that follow the ETSI documentation
// TODO: Uncomment the following lines when Types are implemented
/*
type Guami struct {
	PlmnId PlmnId `json:"plmnId,omitempty"`
	AmfId  AmfId  `json:"amfId,omitempty"`
}
*/
type NetworkId struct {
	Mcc Mcc `json:"mcc,omitempty"`
	Mnc Mnc `json:"mnc,omitempty"`
}

/*
type GuamiRm struct {
	PlmnIdRm PlmnIdRm `json:"plmnId,omitempty"`
	AmfIdRm  AmfIdRm  `json:"amfId,omitempty"`
}
*/
