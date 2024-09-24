/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BlockVolumeMigrationJobsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockVolumeMigrationJobsResp{}

// BlockVolumeMigrationJobsResp struct for BlockVolumeMigrationJobsResp
type BlockVolumeMigrationJobsResp struct {
	// volume migration jobs
	BlockVolumeMigrationJobs []VolumeMigrationJob `json:"block_volume_migration_jobs"`
}

type _BlockVolumeMigrationJobsResp BlockVolumeMigrationJobsResp

// NewBlockVolumeMigrationJobsResp instantiates a new BlockVolumeMigrationJobsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockVolumeMigrationJobsResp(blockVolumeMigrationJobs []VolumeMigrationJob) *BlockVolumeMigrationJobsResp {
	this := BlockVolumeMigrationJobsResp{}
	this.BlockVolumeMigrationJobs = blockVolumeMigrationJobs
	return &this
}

// NewBlockVolumeMigrationJobsRespWithDefaults instantiates a new BlockVolumeMigrationJobsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockVolumeMigrationJobsRespWithDefaults() *BlockVolumeMigrationJobsResp {
	this := BlockVolumeMigrationJobsResp{}
	return &this
}

// GetBlockVolumeMigrationJobs returns the BlockVolumeMigrationJobs field value
func (o *BlockVolumeMigrationJobsResp) GetBlockVolumeMigrationJobs() []VolumeMigrationJob {
	if o == nil {
		var ret []VolumeMigrationJob
		return ret
	}

	return o.BlockVolumeMigrationJobs
}

// GetBlockVolumeMigrationJobsOk returns a tuple with the BlockVolumeMigrationJobs field value
// and a boolean to check if the value has been set.
func (o *BlockVolumeMigrationJobsResp) GetBlockVolumeMigrationJobsOk() ([]VolumeMigrationJob, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockVolumeMigrationJobs, true
}

// SetBlockVolumeMigrationJobs sets field value
func (o *BlockVolumeMigrationJobsResp) SetBlockVolumeMigrationJobs(v []VolumeMigrationJob) {
	o.BlockVolumeMigrationJobs = v
}

func (o BlockVolumeMigrationJobsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockVolumeMigrationJobsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["block_volume_migration_jobs"] = o.BlockVolumeMigrationJobs
	return toSerialize, nil
}

func (o *BlockVolumeMigrationJobsResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"block_volume_migration_jobs",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBlockVolumeMigrationJobsResp := _BlockVolumeMigrationJobsResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockVolumeMigrationJobsResp)

	if err != nil {
		return err
	}

	*o = BlockVolumeMigrationJobsResp(varBlockVolumeMigrationJobsResp)

	return err
}

type NullableBlockVolumeMigrationJobsResp struct {
	value *BlockVolumeMigrationJobsResp
	isSet bool
}

func (v NullableBlockVolumeMigrationJobsResp) Get() *BlockVolumeMigrationJobsResp {
	return v.value
}

func (v *NullableBlockVolumeMigrationJobsResp) Set(val *BlockVolumeMigrationJobsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockVolumeMigrationJobsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockVolumeMigrationJobsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockVolumeMigrationJobsResp(val *BlockVolumeMigrationJobsResp) *NullableBlockVolumeMigrationJobsResp {
	return &NullableBlockVolumeMigrationJobsResp{value: val, isSet: true}
}

func (v NullableBlockVolumeMigrationJobsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockVolumeMigrationJobsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


