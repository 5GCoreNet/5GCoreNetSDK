package trace

import (
	"fmt"
	"github.com/5GCoreNet/5GCoreNetSDK/models/common"
	"regexp"
)

type TraceData struct {
	// TraceRef is the trace reference (see 3GPP TS 32.422).
	// TraceRef shall be encoded as the concatenation of MCC, MNC and Trace ID as follows: <MCC><MNC>-<Trace ID>
	// The Trace ID shall be encoded as a 3 octet string in hexadecimal representation. Each character in the Trace ID string shall take a value of "0" to "9" or "A"
	// to "F" and shall represent 4 bits. The most significant character representing the 4 most significant bits of the Trace ID shall appear first in the string, and the
	// character representing the 4 least significant bit of the Trace ID shall appear last in the string.
	TraceRef string `json:"traceRef"`
	// TraceDepth (see 3GPP TS 32.422 [19]).
	TraceDepth *TraceDepth `json:"traceDepth"`
	// NeTypeList is the list of NE Types (see 3GPP TS 32.422).
	//It shall be encoded as an octet string in hexadecimal representation. Each character in the string shall take a value of "0" to "9" or "A" to "F" and shall represent 4 bits.
	//The most significant character representing the 4 most significant bits shall appear first in the string, and the character representing the
	//4 least significant bit shall appear last in the string. Octets shall be coded according to 3GPP TS 32.422.
	NeTypeList string `json:"neTypeList"`
	// EventList is the list of Triggering events (see 3GPP TS 32.422). It shall be encoded as an octet string in hexadecimal representation. Each character in the string shall
	//take a value of "0" to "9" or "A" to "F" and shall represent 4 bits. The most significant character representing the 4 most significant bits shall appear
	//first in the string, and the character representing the 4 least significant bit shall appear last in the string. Octets shall be coded according to 3GPP TS 32.422.
	EventList string `json:"eventList"`
	// CollectionEntityIpv4Addr is the IPv4 Address of the Trace Collection Entity (see 3GPP TS 32.422).
	//At least one of the collectionEntityIpv4Addr or collectionEntityIpv6Addr attributes shall be present.
	CollectionEntityIpv4Addr *common.Ipv4Addr `json:"collectionEntityIpv4Addr,omitempty"`
	// CollectionEntityIpv6Addr is the IPv6 Address of the Trace Collection Entity (see 3GPP TS 32.422).
	//At least one of the collectionEntityIpv4Addr or collectionEntityIpv6Addr attributes shall be present.
	CollectionEntityIpv6Addr *common.Ipv6Addr `json:"collectionEntityIpv6Addr,omitempty"`
	// InterfaceList is the list of interfaces (see 3GPP TS 32.422). It shall be encoded as an octet string in hexadecimal representation. Each character in the string shall
	//take a value of "0" to "9" or "A" to "F" and shall represent 4 bits. The most significant character representing the 4 most significant bits shall appear first in the string, and the character representing the
	//4 least significant bit shall appear last in the string. Octets shall be coded according to 3GPP TS 32.422.
	//If this attribute is not present, all the interfaces applicable to the list of NE types indicated in the neTypeList attribute should be traced.
	InterfaceList string `json:"interfaceList,omitempty"`
}

// Validate validates this trace data
func (m *TraceData) Validate() error {
	if !regexp.MustCompile("[0-9]{3}[0-9]{2,3}-[A-Fa-f0-9]{6}$'").MatchString(m.TraceRef) {
		return fmt.Errorf("invalid trace reference: %s", m.TraceRef)
	}
	if err := m.TraceDepth.Validate(); err != nil {
		return err
	}
	if !regexp.MustCompile("^[A-Fa-f0-9]+$").MatchString(m.NeTypeList) {
		return fmt.Errorf("invalid NE type list: %s", m.NeTypeList)
	}
	if !regexp.MustCompile("^[A-Fa-f0-9]+$").MatchString(m.EventList) {
		return fmt.Errorf("invalid event list: %s", m.EventList)
	}
	if m.CollectionEntityIpv4Addr != nil && m.CollectionEntityIpv6Addr != nil {
		return fmt.Errorf("collectionEntityIpv4Addr and collectionEntityIpv6Addr must not be set at the same time")
	}
	if m.CollectionEntityIpv4Addr == nil && m.CollectionEntityIpv6Addr == nil {
		return fmt.Errorf("collectionEntityIpv4Addr or collectionEntityIpv6Addr must be set")
	}
	if m.CollectionEntityIpv4Addr != nil {
		if err := m.CollectionEntityIpv4Addr.Validate(); err != nil {
			return err
		}
	}
	if m.CollectionEntityIpv6Addr != nil {
		if err := m.CollectionEntityIpv6Addr.Validate(); err != nil {
			return err
		}
	}
	if !regexp.MustCompile("^[A-Fa-f0-9]+$").MatchString(m.InterfaceList) {
		return fmt.Errorf("invalid interface list: %s", m.InterfaceList)
	}
	return nil
}
