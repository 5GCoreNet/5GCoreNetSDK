package models

// TODO: Add description regarding the ETSI documentation
// TODO: Add validation functions that follow the ETSI documentation
type Dnn string
type DnnRm string
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
type RfspIndexRm int8
type NfGroupId string
