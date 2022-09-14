package models

type Guami struct {
	PlmnId PlmnId `json:"plmnId,omitempty"`
	AmfId  AmfId  `json:"amfId,omitempty"`
}

type NetworkId struct {
	Mcc Mcc `json:"mcc,omitempty"`
	Mnc Mnc `json:"mnc,omitempty"`
}
