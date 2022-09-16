package models

import "fmt"

// ChargingId is the charging identifier allowing correlation of charging information
type ChargingId uint32

// Validate validates the ChargingId.
func (c *ChargingId) Validate() error {
	if c == nil {
		return fmt.Errorf("chargingId must not be nil")
	}
	return nil
}

// RatingGroupId is the identifier of a Rating Group
type RatingGroupId uint32

// Validate validates the ServiceId.
func (r *RatingGroupId) Validate() error {
	if r == nil {
		return fmt.Errorf("ratingGroupId must not be nil")
	}
	return nil
}

// ServiceId is the identifier of a Service
type ServiceId uint32

// Validate validates the ServiceId.
func (s *ServiceId) Validate() error {
	if s == nil {
		return fmt.Errorf("serviceId must not be nil")
	}
	return nil
}
