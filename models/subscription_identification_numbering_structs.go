package models

import "fmt"

type Guami struct {
	PlmnId *PlmnId `json:"plmnId"` // PLMN Identity.
	AmfId  *AmfId  `json:"amfId"`  // AMF Identity.
}

// Validate validates the Guami.
func (g *Guami) Validate() error {
	if g.PlmnId == nil || g.AmfId == nil {
		return fmt.Errorf("plmnId and amfId mustn't be null")
	}
	if err := g.PlmnId.Validate(); err != nil {
		return err
	}
	if err := g.AmfId.Validate(); err != nil {
		return err
	}
	return nil
}

type NetworkId struct {
	Mcc *Mcc `json:"mcc,omitempty"` // Mobile Country Code.
	Mnc *Mnc `json:"mnc,omitempty"` // Mobile Network Code.
	// NOTE: At least one MNC or MCC shall be included.
}

// Validate validates the NetworkId.
func (n *NetworkId) Validate() error {
	if n.Mcc == nil && n.Mnc == nil {
		return fmt.Errorf("mcc and mnc mustn't be null")
	}
	if n.Mcc != nil {
		if err := n.Mcc.Validate(); err != nil {
			return err
		}
	}
	if n.Mnc != nil {
		if err := n.Mnc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// GuamiRm is defined in the same way as the Guami data type, but with the OpenAPI "nullable: true" property.
type GuamiRm Guami

// Validate validates the GuamiRm.
func (g *GuamiRm) Validate() error {
	if g == nil {
		return nil
	}
	return (*Guami)(g).Validate()
}
