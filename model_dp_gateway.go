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

// checks if the DpGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpGateway{}

// DpGateway DpGateway is a process to execute data protection jobs
type DpGateway struct {
	AdminPort *int64 `json:"admin_port,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	GatewayPort *int64 `json:"gateway_port,omitempty"`
	Host *HostNestview `json:"host,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	PolicyNum *int64 `json:"policy_num,omitempty"`
	Status *string `json:"status,omitempty"`
	Update *time.Time `json:"update,omitempty"`
}

// NewDpGateway instantiates a new DpGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpGateway() *DpGateway {
	this := DpGateway{}
	return &this
}

// NewDpGatewayWithDefaults instantiates a new DpGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpGatewayWithDefaults() *DpGateway {
	this := DpGateway{}
	return &this
}

// GetAdminPort returns the AdminPort field value if set, zero value otherwise.
func (o *DpGateway) GetAdminPort() int64 {
	if o == nil || IsNil(o.AdminPort) {
		var ret int64
		return ret
	}
	return *o.AdminPort
}

// GetAdminPortOk returns a tuple with the AdminPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpGateway) GetAdminPortOk() (*int64, bool) {
	if o == nil || IsNil(o.AdminPort) {
		return nil, false
	}
	return o.AdminPort, true
}

// HasAdminPort returns a boolean if a field has been set.
func (o *DpGateway) HasAdminPort() bool {
	if o != nil && !IsNil(o.AdminPort) {
		return true
	}

	return false
}

// SetAdminPort gets a reference to the given int64 and assigns it to the AdminPort field.
func (o *DpGateway) SetAdminPort(v int64) {
	o.AdminPort = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DpGateway) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpGateway) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DpGateway) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *DpGateway) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DpGateway) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpGateway) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DpGateway) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DpGateway) SetCreate(v time.Time) {
	o.Create = &v
}

// GetGatewayPort returns the GatewayPort field value if set, zero value otherwise.
func (o *DpGateway) GetGatewayPort() int64 {
	if o == nil || IsNil(o.GatewayPort) {
		var ret int64
		return ret
	}
	return *o.GatewayPort
}

// GetGatewayPortOk returns a tuple with the GatewayPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpGateway) GetGatewayPortOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayPort) {
		return nil, false
	}
	return o.GatewayPort, true
}

// HasGatewayPort returns a boolean if a field has been set.
func (o *DpGateway) HasGatewayPort() bool {
	if o != nil && !IsNil(o.GatewayPort) {
		return true
	}

	return false
}

// SetGatewayPort gets a reference to the given int64 and assigns it to the GatewayPort field.
func (o *DpGateway) SetGatewayPort(v int64) {
	o.GatewayPort = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *DpGateway) GetHost() HostNestview {
	if o == nil || IsNil(o.Host) {
		var ret HostNestview
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpGateway) GetHostOk() (*HostNestview, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *DpGateway) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given HostNestview and assigns it to the Host field.
func (o *DpGateway) SetHost(v HostNestview) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DpGateway) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpGateway) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DpGateway) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DpGateway) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DpGateway) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpGateway) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DpGateway) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DpGateway) SetName(v string) {
	o.Name = &v
}

// GetPolicyNum returns the PolicyNum field value if set, zero value otherwise.
func (o *DpGateway) GetPolicyNum() int64 {
	if o == nil || IsNil(o.PolicyNum) {
		var ret int64
		return ret
	}
	return *o.PolicyNum
}

// GetPolicyNumOk returns a tuple with the PolicyNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpGateway) GetPolicyNumOk() (*int64, bool) {
	if o == nil || IsNil(o.PolicyNum) {
		return nil, false
	}
	return o.PolicyNum, true
}

// HasPolicyNum returns a boolean if a field has been set.
func (o *DpGateway) HasPolicyNum() bool {
	if o != nil && !IsNil(o.PolicyNum) {
		return true
	}

	return false
}

// SetPolicyNum gets a reference to the given int64 and assigns it to the PolicyNum field.
func (o *DpGateway) SetPolicyNum(v int64) {
	o.PolicyNum = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DpGateway) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpGateway) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DpGateway) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DpGateway) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DpGateway) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpGateway) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DpGateway) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *DpGateway) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o DpGateway) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminPort) {
		toSerialize["admin_port"] = o.AdminPort
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.GatewayPort) {
		toSerialize["gateway_port"] = o.GatewayPort
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PolicyNum) {
		toSerialize["policy_num"] = o.PolicyNum
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	return toSerialize, nil
}

type NullableDpGateway struct {
	value *DpGateway
	isSet bool
}

func (v NullableDpGateway) Get() *DpGateway {
	return v.value
}

func (v *NullableDpGateway) Set(val *DpGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableDpGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableDpGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpGateway(val *DpGateway) *NullableDpGateway {
	return &NullableDpGateway{value: val, isSet: true}
}

func (v NullableDpGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


