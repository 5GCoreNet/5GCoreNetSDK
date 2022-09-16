package models

import "fmt"

// PreemptionCapability indicates the pre-emption capability of a request on other QoS flows. See subclause 5.7.2.2 of 3GPP TS 23.501.
type PreemptionCapability string

// Validate validates the PreemptionCapability string.
func (p *PreemptionCapability) Validate() error {
	if p == nil {
		return fmt.Errorf("preemptionCapability must not be nil")
	}
	switch *p {
	case PreemptionCapabilityNotPreempt:
	case PreemptionCapabilityMayPreempt:
	default:
		return fmt.Errorf("invalid preemption capability: %s", p)
	}
	return nil
}

const (
	PreemptionCapabilityNotPreempt PreemptionCapability = "NOT_PREEMPT" // Shall not trigger pre-emption.
	PreemptionCapabilityMayPreempt PreemptionCapability = "MAY_PREEMPT" // May trigger pre-emption.
)

// PreemptionVulnerability indicates the pre-emption vulnerability of the QoS flow to pre-emption from other QoS flows. See subclause 5.7.2.2 of 3GPP TS 23.501.
type PreemptionVulnerability string

// Validate validates the PreemptionVulnerability string.
func (p *PreemptionVulnerability) Validate() error {
	if p == nil {
		return fmt.Errorf("preemptionVulnerability must not be nil")
	}
	switch *p {
	case PreemptionVulnerabilityNotPreemptable:
	case PreemptionVulnerabilityPreemptable:
	default:
		return fmt.Errorf("invalid preemption vulnerability: %s", p)
	}
	return nil
}

const (
	PreemptionVulnerabilityNotPreemptable PreemptionVulnerability = "NOT_PREEMPTABLE" // Shall not be pre-empted.
	PreemptionVulnerabilityPreemptable    PreemptionVulnerability = "PREEMPTABLE"     // May be pre-empted.
)

// ReflectiveQosAttribute indicates whether certain traffic of the QoS flow may be subject to Reflective QoS (see subclause 5.7.2.3 of 3GPP TS 23.501).
type ReflectiveQosAttribute string

// Validate validates the ReflectiveQosAttribute string.
func (r *ReflectiveQosAttribute) Validate() error {
	if r == nil {
		return fmt.Errorf("reflectiveQosAttribute must not be nil")
	}
	switch *r {
	case ReflectiveQosAttributeRQoS:
	case ReflectiveQosAttributeNoRQoS:
	default:
		return fmt.Errorf("invalid reflective qos attribute: %s", r)
	}
	return nil
}

const (
	ReflectiveQosAttributeRQoS   ReflectiveQosAttribute = "RQOS"    // Certain traffic of the Qos flow may be subject to Reflective QoS.
	ReflectiveQosAttributeNoRQoS ReflectiveQosAttribute = "NO_RQOS" // Traffic of the Qos flow is not subject to Reflective QoS.
)

// NotificationControl indicates whether notifications are requested from the RAN when the GFBR can no longer (or again) be fulfilled for a QoS Flow during the lifetime of the QoS Flow (see subclause 5.7.2.4 of 3GPP TS 23.501).
type NotificationControl string

// Validate validates the NotificationControl string.
func (n *NotificationControl) Validate() error {
	if n == nil {
		return fmt.Errorf("notificationControl must not be nil")
	}
	switch *n {
	case NotificationControlRequested:
	case NotificationControlNotRequested:
	default:
		return fmt.Errorf("invalid notification control: %s", n)
	}
	return nil
}

const (
	NotificationControlRequested    NotificationControl = "REQUESTED"     // Notifications are requested from the RAN.
	NotificationControlNotRequested NotificationControl = "NOT_REQUESTED" // Notifications are not requested from the RAN.
)

// QosResourceType indicates whether a QoS Flow is non-GBR, delay critical GBR, or non-delay critical GBR (see subclauses 5.7.3.4 and 5.7.3.5 of 3GPP TS 23.501).
type QosResourceType string

// Validate validates the QosResourceType string.
func (q *QosResourceType) Validate() error {
	if q == nil {
		return fmt.Errorf("qosResourceType must not be nil")
	}
	switch *q {
	case QosResourceTypeNonGBR:
	case QosResourceTypeNonCriticalGBR:
	case QosResourceTypeCriticalGBR:
	default:
		return fmt.Errorf("invalid qos resource type: %s", q)
	}
	return nil
}

const (
	QosResourceTypeNonGBR         QosResourceType = "NON_GBR"          // Non-GBR QoS Flow.
	QosResourceTypeNonCriticalGBR QosResourceType = "NON_CRITICAL_GBR" // Non-delay critical GBR QoS Flow.
	QosResourceTypeCriticalGBR    QosResourceType = "CRITICAL_GBR"     // Delay critical GBR QoS Flow.
)

// PreemptionCapabilityRm is defined in the same way as the "PreemptionCapability" enumeration, but with the OpenAPI "nullable: true" property.
type PreemptionCapabilityRm PreemptionCapability

// Validate validates the PreemptionCapabilityRm string.
func (p *PreemptionCapabilityRm) Validate() error {
	if p != nil {
		return (*PreemptionCapability)(p).Validate()
	}
	return nil
}

// PreemptionVulnerabilityRm is defined in the same way as the "PreemptionVulnerability" enumeration, but with the OpenAPI "nullable: true" property.
type PreemptionVulnerabilityRm PreemptionVulnerability

// Validate validates the PreemptionVulnerabilityRm string.
func (p *PreemptionVulnerabilityRm) Validate() error {
	if p != nil {
		return (*PreemptionVulnerability)(p).Validate()
	}
	return nil
}

// ReflectiveQosAttributeRm is defined in the same way as the "ReflectiveQosAttribute" enumeration, but with the OpenAPI "nullable: true" property.
type ReflectiveQosAttributeRm ReflectiveQosAttribute

// Validate validates the ReflectiveQosAttributeRm string.
func (r *ReflectiveQosAttributeRm) Validate() error {
	if r != nil {
		return (*ReflectiveQosAttribute)(r).Validate()
	}
	return nil
}

// NotificationControlRm is defined in the same way as the "NotificationControl" enumeration, but with the OpenAPI "nullable: true" property.
type NotificationControlRm NotificationControl

// Validate validates the NotificationControlRm string.
func (n *NotificationControlRm) Validate() error {
	if n != nil {
		return (*NotificationControl)(n).Validate()
	}
	return nil
}

// QosResourceTypeRm is defined in the same way as the "QosResourceType" enumeration, but with the OpenAPI "nullable: true" property.
type QosResourceTypeRm QosResourceType

// Validate validates the QosResourceTypeRm string.
func (q *QosResourceTypeRm) Validate() error {
	if q != nil {
		return (*QosResourceType)(q).Validate()
	}
	return nil
}

// AdditionalQosFlowInfo provides additional QoS flow information (see subclause 9.3.1.12 3GPP TS 38.413).
type AdditionalQosFlowInfo string

// Validate validates the AdditionalQosFlowInfo string.
func (a *AdditionalQosFlowInfo) Validate() error {
	if a == nil {
		return fmt.Errorf("additionalQosFlowInfo must not be nil")
	}
	switch *a {
	case AdditionalQosFlowInfoMoreLikely:
	default:
		return fmt.Errorf("invalid additional qos flow info: %s", a)
	}
	return nil
}

const (
	AdditionalQosFlowInfoMoreLikely AdditionalQosFlowInfo = "MORE_LIKELY" // Traffic for the QoS flow is likely to appear more often than traffic for other flows established for the PDU session.
)
