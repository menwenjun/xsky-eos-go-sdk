/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the OSSearchGatewayRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSSearchGatewayRecord{}

// OSSearchGatewayRecord OSSearchGatewayRecord combine OSSearchGateway and OSSearchGatewayStat to let API comfortable
type OSSearchGatewayRecord struct {
	ActionStatus *string `json:"action_status,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Host *HostNestview `json:"host,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Pool *PoolNestview `json:"pool,omitempty"`
	SearchEngine *OSSearchEngineNestview `json:"search_engine,omitempty"`
	Status *string `json:"status,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	Samples []OSSearchGatewayStat `json:"samples,omitempty"`
}

// NewOSSearchGatewayRecord instantiates a new OSSearchGatewayRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSSearchGatewayRecord() *OSSearchGatewayRecord {
	this := OSSearchGatewayRecord{}
	return &this
}

// NewOSSearchGatewayRecordWithDefaults instantiates a new OSSearchGatewayRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSSearchGatewayRecordWithDefaults() *OSSearchGatewayRecord {
	this := OSSearchGatewayRecord{}
	return &this
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *OSSearchGatewayRecord) GetActionStatus() string {
	if o == nil || IsNil(o.ActionStatus) {
		var ret string
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayRecord) GetActionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *OSSearchGatewayRecord) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given string and assigns it to the ActionStatus field.
func (o *OSSearchGatewayRecord) SetActionStatus(v string) {
	o.ActionStatus = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *OSSearchGatewayRecord) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayRecord) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *OSSearchGatewayRecord) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *OSSearchGatewayRecord) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *OSSearchGatewayRecord) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayRecord) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *OSSearchGatewayRecord) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *OSSearchGatewayRecord) SetCreate(v time.Time) {
	o.Create = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *OSSearchGatewayRecord) GetHost() HostNestview {
	if o == nil || IsNil(o.Host) {
		var ret HostNestview
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayRecord) GetHostOk() (*HostNestview, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *OSSearchGatewayRecord) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given HostNestview and assigns it to the Host field.
func (o *OSSearchGatewayRecord) SetHost(v HostNestview) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OSSearchGatewayRecord) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayRecord) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OSSearchGatewayRecord) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OSSearchGatewayRecord) SetId(v int64) {
	o.Id = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *OSSearchGatewayRecord) GetPool() PoolNestview {
	if o == nil || IsNil(o.Pool) {
		var ret PoolNestview
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayRecord) GetPoolOk() (*PoolNestview, bool) {
	if o == nil || IsNil(o.Pool) {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *OSSearchGatewayRecord) HasPool() bool {
	if o != nil && !IsNil(o.Pool) {
		return true
	}

	return false
}

// SetPool gets a reference to the given PoolNestview and assigns it to the Pool field.
func (o *OSSearchGatewayRecord) SetPool(v PoolNestview) {
	o.Pool = &v
}

// GetSearchEngine returns the SearchEngine field value if set, zero value otherwise.
func (o *OSSearchGatewayRecord) GetSearchEngine() OSSearchEngineNestview {
	if o == nil || IsNil(o.SearchEngine) {
		var ret OSSearchEngineNestview
		return ret
	}
	return *o.SearchEngine
}

// GetSearchEngineOk returns a tuple with the SearchEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayRecord) GetSearchEngineOk() (*OSSearchEngineNestview, bool) {
	if o == nil || IsNil(o.SearchEngine) {
		return nil, false
	}
	return o.SearchEngine, true
}

// HasSearchEngine returns a boolean if a field has been set.
func (o *OSSearchGatewayRecord) HasSearchEngine() bool {
	if o != nil && !IsNil(o.SearchEngine) {
		return true
	}

	return false
}

// SetSearchEngine gets a reference to the given OSSearchEngineNestview and assigns it to the SearchEngine field.
func (o *OSSearchGatewayRecord) SetSearchEngine(v OSSearchEngineNestview) {
	o.SearchEngine = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OSSearchGatewayRecord) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayRecord) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OSSearchGatewayRecord) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OSSearchGatewayRecord) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *OSSearchGatewayRecord) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayRecord) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *OSSearchGatewayRecord) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *OSSearchGatewayRecord) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetSamples returns the Samples field value if set, zero value otherwise.
func (o *OSSearchGatewayRecord) GetSamples() []OSSearchGatewayStat {
	if o == nil || IsNil(o.Samples) {
		var ret []OSSearchGatewayStat
		return ret
	}
	return o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayRecord) GetSamplesOk() ([]OSSearchGatewayStat, bool) {
	if o == nil || IsNil(o.Samples) {
		return nil, false
	}
	return o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *OSSearchGatewayRecord) HasSamples() bool {
	if o != nil && !IsNil(o.Samples) {
		return true
	}

	return false
}

// SetSamples gets a reference to the given []OSSearchGatewayStat and assigns it to the Samples field.
func (o *OSSearchGatewayRecord) SetSamples(v []OSSearchGatewayStat) {
	o.Samples = v
}

func (o OSSearchGatewayRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSSearchGatewayRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionStatus) {
		toSerialize["action_status"] = o.ActionStatus
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Pool) {
		toSerialize["pool"] = o.Pool
	}
	if !IsNil(o.SearchEngine) {
		toSerialize["search_engine"] = o.SearchEngine
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Samples) {
		toSerialize["samples"] = o.Samples
	}
	return toSerialize, nil
}

type NullableOSSearchGatewayRecord struct {
	value *OSSearchGatewayRecord
	isSet bool
}

func (v NullableOSSearchGatewayRecord) Get() *OSSearchGatewayRecord {
	return v.value
}

func (v *NullableOSSearchGatewayRecord) Set(val *OSSearchGatewayRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableOSSearchGatewayRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableOSSearchGatewayRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSSearchGatewayRecord(val *OSSearchGatewayRecord) *NullableOSSearchGatewayRecord {
	return &NullableOSSearchGatewayRecord{value: val, isSet: true}
}

func (v NullableOSSearchGatewayRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSSearchGatewayRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


