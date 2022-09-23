package models

import "fmt"

// TODO: Add description regarding the ETSI documentation
// TODO: Add validation functions that follow the ETSI documentation

// Dnn reprsents  a Data Network as defined in subclause 9A of
// 3GPP TS 23.003. It shall be formatted as string in which the
// labels are separated by dots (e.g. "Label1.Label2.Label3").
type Dnn string

// Validate validates the Dnn string.
func (d *Dnn) Validate() error {
	if d == nil {
		return fmt.Errorf("dnn mustn't be null")
	}
	return nil
}

// DnnRm is defined in the same way as the Dnn data type, but with the OpenAPI "nullable: true" property.
type DnnRm Dnn

// Validate validates the DnnRm string.
func (d *DnnRm) Validate() error {
	if d == nil {
		return nil
	}
	return (*Dnn)(d).Validate()
}

type Gpsi string
type GpsiRm Gpsi
type GroupId string
type GroupIdRm GroupId
type Pei string
type PeiRm Pei
type Supi string
type SupiRm Supi
type NfInstanceId string
type AmfId string

func (i AmfId) Validate() error {
	return nil
}

type AmfRegionId string
type AmfSetId string
type RfspIndex int8
type RfspIndexRm RfspIndex
type NfGroupId string
