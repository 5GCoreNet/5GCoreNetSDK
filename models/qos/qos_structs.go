package qos

import "fmt"

type Arp struct {
	PriorityLevel *ArpPriorityLevel        `json:"priorityLevel"` // Defines the relative importance of a resource request.
	PreemptCap    *PreemptionCapability    `json:"preemptCap"`    // Defines whether a service data flow may get resources that were already assigned to another service data flow with a lower priority level.
	PreemptVuln   *PreemptionVulnerability `json:"preemptVuln"`   // Defines whether a service data flow may lose the resources assigned to it in order to admit a service data flow with higher priority level.
}

// Validate validates this arp.
func (a *Arp) Validate() error {
	if a.PreemptVuln == nil {
		return fmt.Errorf("preemptionVulnerability must not be nil")
	}
	if err := a.PreemptVuln.Validate(); err != nil {
		return fmt.Errorf("preemptionVulnerability %s", err.Error())
	}
	if a.PreemptCap == nil {
		return fmt.Errorf("preemptionCapability must not be nil")
	}
	if err := a.PreemptCap.Validate(); err != nil {
		return fmt.Errorf("preemptionCapability %s", err.Error())
	}
	if a.PriorityLevel == nil {
		return fmt.Errorf("priorityLevel must not be nil")
	}
	if err := a.PriorityLevel.Validate(); err != nil {
		return fmt.Errorf("priorityLevel %s", err.Error())
	}
	return nil
}

type Ambr struct {
	Uplink   *BitRate `json:"uplink"`   // AMBR for uplink
	Downlink *BitRate `json:"downlink"` // AMBR for downlink
}

// Validate validates this ambr.
func (a *Ambr) Validate() error {
	if a.Downlink == nil {
		return fmt.Errorf("downlink must not be nil")
	}
	if err := a.Downlink.Validate(); err != nil {
		return fmt.Errorf("downlink %s", err.Error())
	}
	if a.Uplink == nil {
		return fmt.Errorf("uplink must not be nil")
	}
	if err := a.Uplink.Validate(); err != nil {
		return fmt.Errorf("uplink %s", err.Error())
	}
	return nil
}

type Dynamic5Qi struct {
	ResourceType      *QosResourceType     `json:"resourceType"`              // Defines the 5QI resource type.
	PriorityLevel     *FiveQiPriorityLevel `json:"priorityLevel"`             // Defines the 5QI Priority Level.
	PacketDelayBudget *PacketDelBudget     `json:"packetDelayBudget"`         // Defines the packet delay budget.
	PacketErrRate     *PacketErrRate       `json:"packetErrRate"`             // Defines the packet error rate.
	AverWindow        *AverWindow          `json:"averWindow,omitempty"`      // Defines the averaging window. This IE shall be present only for a GBR QoS flow or a Delay Critical GBR QoS flow.
	MaxDataBurstVol   *MaxDataBurstVol     `json:"maxDataBurstVol,omitempty"` // Defines the maximum data burst volume. This IE shall be present for a Delay Critical GBR QoS flow.
}

// Validate validates this dynamic5Qi.
func (d *Dynamic5Qi) Validate() error {
	if d.ResourceType == nil {
		return fmt.Errorf("resourceType must not be nil")
	}
	if err := d.ResourceType.Validate(); err != nil {
		return fmt.Errorf("resourceType %s", err.Error())
	}
	if d.PriorityLevel == nil {
		return fmt.Errorf("priorityLevel must not be nil")
	}
	if err := d.PriorityLevel.Validate(); err != nil {
		return fmt.Errorf("priorityLevel %s", err.Error())
	}
	if d.PacketDelayBudget == nil {
		return fmt.Errorf("packetDelayBudget must not be nil")
	}
	if err := d.PacketDelayBudget.Validate(); err != nil {
		return fmt.Errorf("packetDelayBudget %s", err.Error())
	}
	if d.PacketErrRate == nil {
		return fmt.Errorf("packetErrRate must not be nil")
	}
	if err := d.PacketErrRate.Validate(); err != nil {
		return fmt.Errorf("packetErrRate %s", err.Error())
	}
	if d.AverWindow != nil {
		if err := d.AverWindow.Validate(); err != nil {
			return fmt.Errorf("averWindow %s", err.Error())
		}
	}
	if d.MaxDataBurstVol != nil {
		if err := d.MaxDataBurstVol.Validate(); err != nil {
			return fmt.Errorf("maxDataBurstVol %s", err.Error())
		}
	}
	return nil
}

type NonDynamic5Qi struct {
	PriorityLevel   *FiveQiPriorityLevel `json:"priorityLevel,omitempty"`   // Defines the 5QI Priority Level. When present, it contains the 5QI Priority Level value that overrides the standardized or pre-configured value.
	AverWindow      *AverWindow          `json:"averWindow,omitempty"`      // Defines the averaging window. This IE may be present for a GBR QoS flow or a Delay Critical GBR QoS flow. When present, it contains the Averaging Window that overrides the standardized or pre-configured value
	MaxDataBurstVol *MaxDataBurstVol     `json:"maxDataBurstVol,omitempty"` // Defines the maximum data burst volume. This IE may be present for a Delay Critical GBR QoS flow. When present, it contains the Maximum Data Burst Volume value that overrides the standardized or pre-configured value.
}

// Validate validates this nonDynamic5Qi.
func (n *NonDynamic5Qi) Validate() error {
	if n.PriorityLevel != nil {
		if err := n.PriorityLevel.Validate(); err != nil {
			return fmt.Errorf("priorityLevel %s", err.Error())
		}
	}
	if n.AverWindow != nil {
		if err := n.AverWindow.Validate(); err != nil {
			return fmt.Errorf("averWindow %s", err.Error())
		}
	}
	if n.MaxDataBurstVol != nil {
		if err := n.MaxDataBurstVol.Validate(); err != nil {
			return fmt.Errorf("maxDataBurstVol %s", err.Error())
		}
	}
	return nil
}

// ArpRm is defined in the same way as the Arp data type, but with the OpenAPI "nullable: true" property.
type ArpRm Arp

// Validate validates this arpRm.
func (a *ArpRm) Validate() error {
	if a == nil {
		return nil
	}
	if err := (*Arp)(a).Validate(); err != nil {
		return err
	}
	return nil
}

// AmbrRm is defined in the same way as the Ambr data type, but with the OpenAPI "nullable: true" property.
type AmbrRm Ambr

// Validate validates this ambrRm.
func (a *AmbrRm) Validate() error {
	if a == nil {
		return nil
	}
	if err := (*Ambr)(a).Validate(); err != nil {
		return err
	}
	return nil
}
