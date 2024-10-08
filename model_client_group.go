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

// checks if the ClientGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientGroup{}

// ClientGroup ClientGroup defines the access client group
type ClientGroup struct {
	AccessPathNum *int64 `json:"access_path_num,omitempty"`
	BlockVolumeNum *int64 `json:"block_volume_num,omitempty"`
	Chap *bool `json:"chap,omitempty"`
	ClientNum *int64 `json:"client_num,omitempty"`
	Clients []Client `json:"clients,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Iname *string `json:"iname,omitempty"`
	Isecret *string `json:"isecret,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	Type *string `json:"type,omitempty"`
	Uid *string `json:"uid,omitempty"`
	Update *time.Time `json:"update,omitempty"`
}

// NewClientGroup instantiates a new ClientGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientGroup() *ClientGroup {
	this := ClientGroup{}
	return &this
}

// NewClientGroupWithDefaults instantiates a new ClientGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientGroupWithDefaults() *ClientGroup {
	this := ClientGroup{}
	return &this
}

// GetAccessPathNum returns the AccessPathNum field value if set, zero value otherwise.
func (o *ClientGroup) GetAccessPathNum() int64 {
	if o == nil || IsNil(o.AccessPathNum) {
		var ret int64
		return ret
	}
	return *o.AccessPathNum
}

// GetAccessPathNumOk returns a tuple with the AccessPathNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroup) GetAccessPathNumOk() (*int64, bool) {
	if o == nil || IsNil(o.AccessPathNum) {
		return nil, false
	}
	return o.AccessPathNum, true
}

// HasAccessPathNum returns a boolean if a field has been set.
func (o *ClientGroup) HasAccessPathNum() bool {
	if o != nil && !IsNil(o.AccessPathNum) {
		return true
	}

	return false
}

// SetAccessPathNum gets a reference to the given int64 and assigns it to the AccessPathNum field.
func (o *ClientGroup) SetAccessPathNum(v int64) {
	o.AccessPathNum = &v
}

// GetBlockVolumeNum returns the BlockVolumeNum field value if set, zero value otherwise.
func (o *ClientGroup) GetBlockVolumeNum() int64 {
	if o == nil || IsNil(o.BlockVolumeNum) {
		var ret int64
		return ret
	}
	return *o.BlockVolumeNum
}

// GetBlockVolumeNumOk returns a tuple with the BlockVolumeNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroup) GetBlockVolumeNumOk() (*int64, bool) {
	if o == nil || IsNil(o.BlockVolumeNum) {
		return nil, false
	}
	return o.BlockVolumeNum, true
}

// HasBlockVolumeNum returns a boolean if a field has been set.
func (o *ClientGroup) HasBlockVolumeNum() bool {
	if o != nil && !IsNil(o.BlockVolumeNum) {
		return true
	}

	return false
}

// SetBlockVolumeNum gets a reference to the given int64 and assigns it to the BlockVolumeNum field.
func (o *ClientGroup) SetBlockVolumeNum(v int64) {
	o.BlockVolumeNum = &v
}

// GetChap returns the Chap field value if set, zero value otherwise.
func (o *ClientGroup) GetChap() bool {
	if o == nil || IsNil(o.Chap) {
		var ret bool
		return ret
	}
	return *o.Chap
}

// GetChapOk returns a tuple with the Chap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroup) GetChapOk() (*bool, bool) {
	if o == nil || IsNil(o.Chap) {
		return nil, false
	}
	return o.Chap, true
}

// HasChap returns a boolean if a field has been set.
func (o *ClientGroup) HasChap() bool {
	if o != nil && !IsNil(o.Chap) {
		return true
	}

	return false
}

// SetChap gets a reference to the given bool and assigns it to the Chap field.
func (o *ClientGroup) SetChap(v bool) {
	o.Chap = &v
}

// GetClientNum returns the ClientNum field value if set, zero value otherwise.
func (o *ClientGroup) GetClientNum() int64 {
	if o == nil || IsNil(o.ClientNum) {
		var ret int64
		return ret
	}
	return *o.ClientNum
}

// GetClientNumOk returns a tuple with the ClientNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroup) GetClientNumOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientNum) {
		return nil, false
	}
	return o.ClientNum, true
}

// HasClientNum returns a boolean if a field has been set.
func (o *ClientGroup) HasClientNum() bool {
	if o != nil && !IsNil(o.ClientNum) {
		return true
	}

	return false
}

// SetClientNum gets a reference to the given int64 and assigns it to the ClientNum field.
func (o *ClientGroup) SetClientNum(v int64) {
	o.ClientNum = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *ClientGroup) GetClients() []Client {
	if o == nil || IsNil(o.Clients) {
		var ret []Client
		return ret
	}
	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroup) GetClientsOk() ([]Client, bool) {
	if o == nil || IsNil(o.Clients) {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *ClientGroup) HasClients() bool {
	if o != nil && !IsNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given []Client and assigns it to the Clients field.
func (o *ClientGroup) SetClients(v []Client) {
	o.Clients = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *ClientGroup) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroup) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ClientGroup) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *ClientGroup) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *ClientGroup) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroup) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ClientGroup) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *ClientGroup) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClientGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClientGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClientGroup) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClientGroup) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroup) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClientGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ClientGroup) SetId(v int64) {
	o.Id = &v
}

// GetIname returns the Iname field value if set, zero value otherwise.
func (o *ClientGroup) GetIname() string {
	if o == nil || IsNil(o.Iname) {
		var ret string
		return ret
	}
	return *o.Iname
}

// GetInameOk returns a tuple with the Iname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroup) GetInameOk() (*string, bool) {
	if o == nil || IsNil(o.Iname) {
		return nil, false
	}
	return o.Iname, true
}

// HasIname returns a boolean if a field has been set.
func (o *ClientGroup) HasIname() bool {
	if o != nil && !IsNil(o.Iname) {
		return true
	}

	return false
}

// SetIname gets a reference to the given string and assigns it to the Iname field.
func (o *ClientGroup) SetIname(v string) {
	o.Iname = &v
}

// GetIsecret returns the Isecret field value if set, zero value otherwise.
func (o *ClientGroup) GetIsecret() string {
	if o == nil || IsNil(o.Isecret) {
		var ret string
		return ret
	}
	return *o.Isecret
}

// GetIsecretOk returns a tuple with the Isecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroup) GetIsecretOk() (*string, bool) {
	if o == nil || IsNil(o.Isecret) {
		return nil, false
	}
	return o.Isecret, true
}

// HasIsecret returns a boolean if a field has been set.
func (o *ClientGroup) HasIsecret() bool {
	if o != nil && !IsNil(o.Isecret) {
		return true
	}

	return false
}

// SetIsecret gets a reference to the given string and assigns it to the Isecret field.
func (o *ClientGroup) SetIsecret(v string) {
	o.Isecret = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClientGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClientGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClientGroup) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClientGroup) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroup) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClientGroup) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ClientGroup) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClientGroup) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroup) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClientGroup) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClientGroup) SetType(v string) {
	o.Type = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *ClientGroup) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroup) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *ClientGroup) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *ClientGroup) SetUid(v string) {
	o.Uid = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *ClientGroup) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroup) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *ClientGroup) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *ClientGroup) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o ClientGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessPathNum) {
		toSerialize["access_path_num"] = o.AccessPathNum
	}
	if !IsNil(o.BlockVolumeNum) {
		toSerialize["block_volume_num"] = o.BlockVolumeNum
	}
	if !IsNil(o.Chap) {
		toSerialize["chap"] = o.Chap
	}
	if !IsNil(o.ClientNum) {
		toSerialize["client_num"] = o.ClientNum
	}
	if !IsNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Iname) {
		toSerialize["iname"] = o.Iname
	}
	if !IsNil(o.Isecret) {
		toSerialize["isecret"] = o.Isecret
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	return toSerialize, nil
}

type NullableClientGroup struct {
	value *ClientGroup
	isSet bool
}

func (v NullableClientGroup) Get() *ClientGroup {
	return v.value
}

func (v *NullableClientGroup) Set(val *ClientGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableClientGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableClientGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientGroup(val *ClientGroup) *NullableClientGroup {
	return &NullableClientGroup{value: val, isSet: true}
}

func (v NullableClientGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


