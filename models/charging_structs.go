package models

import "fmt"

type SecondaryRatUsageReport struct {
	SecondaryRatType  *RatType              `json:"secondaryRatType"`  // Secondary RAT type.
	QosFlowsUsageData []*QosFlowUsageReport `json:"qosFlowsUsageData"` // QoS flows usage data.
}

// Validate validates this secondary rat usage report.
func (m *SecondaryRatUsageReport) Validate() error {
	if err := m.SecondaryRatType.Validate(); err != nil {
		return fmt.Errorf("secondaryRatType: %s", err)
	}
	if len(m.QosFlowsUsageData) < 1 {
		return fmt.Errorf("qosFlowsUsageData must contain at least one element")
	}
	for i, item := range m.QosFlowsUsageData {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("qosFlowsUsageData[%d]: %s", i, err)
		}
	}
	return nil
}

type QosFlowUsageReport struct {
	Qfi            *Qfi      `json:"qfi"`            // QoS Flow Indicator.
	StartTimestamp *DateTime `json:"startTimeStamp"` // UTC time indicating the start time of the collection period of the included usage data for DL and UL.
	EndTimestamp   *DateTime `json:"endTimeStamp"`   // UTC time indicating the end time of the collection period of the included usage data for DL and UL.
	DownlinkVolume *Int64    `json:"downlinkVolume"` // Data usage for DL, encoding a number of octets.
	UplinkVolume   *Int64    `json:"uplinkVolume"`   // Data usage for UL, encoding a number of octets.
}

// Validate validates this qos flow usage report.
func (m *QosFlowUsageReport) Validate() error {
	if err := m.Qfi.Validate(); err != nil {
		return fmt.Errorf("qfi: %s", err)
	}
	if err := m.StartTimestamp.Validate(); err != nil {
		return fmt.Errorf("startTimestamp: %s", err)
	}
	if err := m.EndTimestamp.Validate(); err != nil {
		return fmt.Errorf("endTimestamp: %s", err)
	}
	if err := m.DownlinkVolume.Validate(); err != nil {
		return fmt.Errorf("downlinkVolume: %s", err)
	}
	if err := m.UplinkVolume.Validate(); err != nil {
		return fmt.Errorf("uplinkVolume: %s", err)
	}
	return nil
}
