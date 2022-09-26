package barring

type OdbData struct {
	RoamingOdb *RoamingOdb `json:"roamingOdb,omitempty"` // Barring of Roaming (see 3GPP TS 23.015).
}

// Validate validates this odb data
func (m *OdbData) Validate() error {
	if m.RoamingOdb != nil {
		if err := m.RoamingOdb.Validate(); err != nil {
			return err
		}
	}
	return nil
}
