package models

import "fmt"

// TraceDepth defines how detailed information should be recorded in the trace. See 3GPP TS 32.422 for further description of the values.
type TraceDepth string

// Validate validates the TraceDepth string.
func (t *TraceDepth) Validate() error {
	if t == nil {
		return fmt.Errorf("traceDepth must not be nil")
	}
	switch *t {
	case TraceDepthMinimum:
	case TraceDepthMedium:
	case TraceDepthMaximum:
	case TraceDepthMinimumWoVendorExtension:
	case TraceDepthMediumWoVendorExtension:
	case TraceDepthMaximumWoVendorExtension:
	default:
		return fmt.Errorf("invalid trace depth: %s", t)
	}
	return nil
}

const (
	TraceDepthMinimum                  TraceDepth = "MINIMUM"                     // Minimum trace depth.
	TraceDepthMedium                   TraceDepth = "MEDIUM"                      // Medium trace depth.
	TraceDepthMaximum                  TraceDepth = "MAXIMUM"                     // Maximum trace depth.
	TraceDepthMinimumWoVendorExtension TraceDepth = "MINIMUM_WO_VENDOR_EXTENSION" // Minimum without vendor specific extension.
	TraceDepthMediumWoVendorExtension  TraceDepth = "MEDIUM_WO_VENDOR_EXTENSION"  // Medium without vendor specific extension.
	TraceDepthMaximumWoVendorExtension TraceDepth = "MAXIMUM_WO_VENDOR_EXTENSION" // Maximum without vendor specific extension.
)

// TraceDepthRm the same way as the TraceDepth enumeration, but with the OpenAPI "nullable: true" property.
type TraceDepthRm TraceDepth

// Validate validates the TraceDepthRm string.
func (t *TraceDepthRm) Validate() error {
	if t == nil {
		return nil
	}
	return (*TraceDepth)(t).Validate()
}
