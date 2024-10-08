/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DpSiteConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpSiteConfig{}

// DpSiteConfig DpSiteConfig contains config of different sites in dpg
type DpSiteConfig struct {
	BlockAsyncReplication *AsyncReplicationConfig `json:"block_async_replication,omitempty"`
	BlockSnapshotBackup *SnapshotBackupConfig `json:"block_snapshot_backup,omitempty"`
	BlockVolumeReplication *BlockReplicationConfig `json:"block_volume_replication,omitempty"`
}

// NewDpSiteConfig instantiates a new DpSiteConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpSiteConfig() *DpSiteConfig {
	this := DpSiteConfig{}
	return &this
}

// NewDpSiteConfigWithDefaults instantiates a new DpSiteConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpSiteConfigWithDefaults() *DpSiteConfig {
	this := DpSiteConfig{}
	return &this
}

// GetBlockAsyncReplication returns the BlockAsyncReplication field value if set, zero value otherwise.
func (o *DpSiteConfig) GetBlockAsyncReplication() AsyncReplicationConfig {
	if o == nil || IsNil(o.BlockAsyncReplication) {
		var ret AsyncReplicationConfig
		return ret
	}
	return *o.BlockAsyncReplication
}

// GetBlockAsyncReplicationOk returns a tuple with the BlockAsyncReplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpSiteConfig) GetBlockAsyncReplicationOk() (*AsyncReplicationConfig, bool) {
	if o == nil || IsNil(o.BlockAsyncReplication) {
		return nil, false
	}
	return o.BlockAsyncReplication, true
}

// HasBlockAsyncReplication returns a boolean if a field has been set.
func (o *DpSiteConfig) HasBlockAsyncReplication() bool {
	if o != nil && !IsNil(o.BlockAsyncReplication) {
		return true
	}

	return false
}

// SetBlockAsyncReplication gets a reference to the given AsyncReplicationConfig and assigns it to the BlockAsyncReplication field.
func (o *DpSiteConfig) SetBlockAsyncReplication(v AsyncReplicationConfig) {
	o.BlockAsyncReplication = &v
}

// GetBlockSnapshotBackup returns the BlockSnapshotBackup field value if set, zero value otherwise.
func (o *DpSiteConfig) GetBlockSnapshotBackup() SnapshotBackupConfig {
	if o == nil || IsNil(o.BlockSnapshotBackup) {
		var ret SnapshotBackupConfig
		return ret
	}
	return *o.BlockSnapshotBackup
}

// GetBlockSnapshotBackupOk returns a tuple with the BlockSnapshotBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpSiteConfig) GetBlockSnapshotBackupOk() (*SnapshotBackupConfig, bool) {
	if o == nil || IsNil(o.BlockSnapshotBackup) {
		return nil, false
	}
	return o.BlockSnapshotBackup, true
}

// HasBlockSnapshotBackup returns a boolean if a field has been set.
func (o *DpSiteConfig) HasBlockSnapshotBackup() bool {
	if o != nil && !IsNil(o.BlockSnapshotBackup) {
		return true
	}

	return false
}

// SetBlockSnapshotBackup gets a reference to the given SnapshotBackupConfig and assigns it to the BlockSnapshotBackup field.
func (o *DpSiteConfig) SetBlockSnapshotBackup(v SnapshotBackupConfig) {
	o.BlockSnapshotBackup = &v
}

// GetBlockVolumeReplication returns the BlockVolumeReplication field value if set, zero value otherwise.
func (o *DpSiteConfig) GetBlockVolumeReplication() BlockReplicationConfig {
	if o == nil || IsNil(o.BlockVolumeReplication) {
		var ret BlockReplicationConfig
		return ret
	}
	return *o.BlockVolumeReplication
}

// GetBlockVolumeReplicationOk returns a tuple with the BlockVolumeReplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpSiteConfig) GetBlockVolumeReplicationOk() (*BlockReplicationConfig, bool) {
	if o == nil || IsNil(o.BlockVolumeReplication) {
		return nil, false
	}
	return o.BlockVolumeReplication, true
}

// HasBlockVolumeReplication returns a boolean if a field has been set.
func (o *DpSiteConfig) HasBlockVolumeReplication() bool {
	if o != nil && !IsNil(o.BlockVolumeReplication) {
		return true
	}

	return false
}

// SetBlockVolumeReplication gets a reference to the given BlockReplicationConfig and assigns it to the BlockVolumeReplication field.
func (o *DpSiteConfig) SetBlockVolumeReplication(v BlockReplicationConfig) {
	o.BlockVolumeReplication = &v
}

func (o DpSiteConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpSiteConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockAsyncReplication) {
		toSerialize["block_async_replication"] = o.BlockAsyncReplication
	}
	if !IsNil(o.BlockSnapshotBackup) {
		toSerialize["block_snapshot_backup"] = o.BlockSnapshotBackup
	}
	if !IsNil(o.BlockVolumeReplication) {
		toSerialize["block_volume_replication"] = o.BlockVolumeReplication
	}
	return toSerialize, nil
}

type NullableDpSiteConfig struct {
	value *DpSiteConfig
	isSet bool
}

func (v NullableDpSiteConfig) Get() *DpSiteConfig {
	return v.value
}

func (v *NullableDpSiteConfig) Set(val *DpSiteConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDpSiteConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDpSiteConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpSiteConfig(val *DpSiteConfig) *NullableDpSiteConfig {
	return &NullableDpSiteConfig{value: val, isSet: true}
}

func (v NullableDpSiteConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpSiteConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


