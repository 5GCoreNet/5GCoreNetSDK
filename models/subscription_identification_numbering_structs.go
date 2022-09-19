package models

// TODO: Add description regarding the ETSI documentation
// TODO: Add validation functions that follow the ETSI documentation

type Guami struct {
	PlmnId *PlmnId `json:"plmnId,omitempty"`
	AmfId  *AmfId  `json:"amfId,omitempty"`
}

func (g *Guami) Validate() error {
	// TODO: Add validation functions that follow the ETSI documentation
	return nil
}

type NetworkId struct {
	Mcc Mcc `json:"mcc,omitempty"`
	Mnc Mnc `json:"mnc,omitempty"`
}

type GuamiRm Guami

func (g *GuamiRm) Validate() error {
	if g.PlmnId != nil {
		if err := g.PlmnId.Validate(); err != nil {
			return err
		}
	}
	if g.AmfId != nil {
		if err := g.AmfId.Validate(); err != nil {
			return err
		}
	}
	return nil
}
