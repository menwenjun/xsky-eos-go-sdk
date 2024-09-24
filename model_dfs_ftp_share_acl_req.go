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

// checks if the DfsFTPShareACLReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsFTPShareACLReq{}

// DfsFTPShareACLReq struct for DfsFTPShareACLReq
type DfsFTPShareACLReq struct {
	// enable admin user permission
	AdminEnabled *bool `json:"admin_enabled,omitempty"`
	// enable creating files
	CreateEnabled *bool `json:"create_enabled,omitempty"`
	// enable deleting files
	DeleteEnabled *bool `json:"delete_enabled,omitempty"`
	// max bandwidth of downloading
	DownloadBandwidth *int64 `json:"download_bandwidth,omitempty"`
	// enable downloading files
	DownloadEnabled *bool `json:"download_enabled,omitempty"`
	// id of user
	FsUserId *int64 `json:"fs_user_id,omitempty"`
	// id of user group
	Id *int64 `json:"id,omitempty"`
	// enable listing files
	ListEnabled *bool `json:"list_enabled,omitempty"`
	// enable renaming files
	RenameEnabled *bool `json:"rename_enabled,omitempty"`
	// type of share acl
	Type *string `json:"type,omitempty"`
	// max bandwidth of uploading
	UploadBandwidth *int64 `json:"upload_bandwidth,omitempty"`
	// enable uploading files
	UploadEnabled *bool `json:"upload_enabled,omitempty"`
}

// NewDfsFTPShareACLReq instantiates a new DfsFTPShareACLReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsFTPShareACLReq() *DfsFTPShareACLReq {
	this := DfsFTPShareACLReq{}
	return &this
}

// NewDfsFTPShareACLReqWithDefaults instantiates a new DfsFTPShareACLReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsFTPShareACLReqWithDefaults() *DfsFTPShareACLReq {
	this := DfsFTPShareACLReq{}
	return &this
}

// GetAdminEnabled returns the AdminEnabled field value if set, zero value otherwise.
func (o *DfsFTPShareACLReq) GetAdminEnabled() bool {
	if o == nil || IsNil(o.AdminEnabled) {
		var ret bool
		return ret
	}
	return *o.AdminEnabled
}

// GetAdminEnabledOk returns a tuple with the AdminEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFTPShareACLReq) GetAdminEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AdminEnabled) {
		return nil, false
	}
	return o.AdminEnabled, true
}

// HasAdminEnabled returns a boolean if a field has been set.
func (o *DfsFTPShareACLReq) HasAdminEnabled() bool {
	if o != nil && !IsNil(o.AdminEnabled) {
		return true
	}

	return false
}

// SetAdminEnabled gets a reference to the given bool and assigns it to the AdminEnabled field.
func (o *DfsFTPShareACLReq) SetAdminEnabled(v bool) {
	o.AdminEnabled = &v
}

// GetCreateEnabled returns the CreateEnabled field value if set, zero value otherwise.
func (o *DfsFTPShareACLReq) GetCreateEnabled() bool {
	if o == nil || IsNil(o.CreateEnabled) {
		var ret bool
		return ret
	}
	return *o.CreateEnabled
}

// GetCreateEnabledOk returns a tuple with the CreateEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFTPShareACLReq) GetCreateEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateEnabled) {
		return nil, false
	}
	return o.CreateEnabled, true
}

// HasCreateEnabled returns a boolean if a field has been set.
func (o *DfsFTPShareACLReq) HasCreateEnabled() bool {
	if o != nil && !IsNil(o.CreateEnabled) {
		return true
	}

	return false
}

// SetCreateEnabled gets a reference to the given bool and assigns it to the CreateEnabled field.
func (o *DfsFTPShareACLReq) SetCreateEnabled(v bool) {
	o.CreateEnabled = &v
}

// GetDeleteEnabled returns the DeleteEnabled field value if set, zero value otherwise.
func (o *DfsFTPShareACLReq) GetDeleteEnabled() bool {
	if o == nil || IsNil(o.DeleteEnabled) {
		var ret bool
		return ret
	}
	return *o.DeleteEnabled
}

// GetDeleteEnabledOk returns a tuple with the DeleteEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFTPShareACLReq) GetDeleteEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteEnabled) {
		return nil, false
	}
	return o.DeleteEnabled, true
}

// HasDeleteEnabled returns a boolean if a field has been set.
func (o *DfsFTPShareACLReq) HasDeleteEnabled() bool {
	if o != nil && !IsNil(o.DeleteEnabled) {
		return true
	}

	return false
}

// SetDeleteEnabled gets a reference to the given bool and assigns it to the DeleteEnabled field.
func (o *DfsFTPShareACLReq) SetDeleteEnabled(v bool) {
	o.DeleteEnabled = &v
}

// GetDownloadBandwidth returns the DownloadBandwidth field value if set, zero value otherwise.
func (o *DfsFTPShareACLReq) GetDownloadBandwidth() int64 {
	if o == nil || IsNil(o.DownloadBandwidth) {
		var ret int64
		return ret
	}
	return *o.DownloadBandwidth
}

// GetDownloadBandwidthOk returns a tuple with the DownloadBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFTPShareACLReq) GetDownloadBandwidthOk() (*int64, bool) {
	if o == nil || IsNil(o.DownloadBandwidth) {
		return nil, false
	}
	return o.DownloadBandwidth, true
}

// HasDownloadBandwidth returns a boolean if a field has been set.
func (o *DfsFTPShareACLReq) HasDownloadBandwidth() bool {
	if o != nil && !IsNil(o.DownloadBandwidth) {
		return true
	}

	return false
}

// SetDownloadBandwidth gets a reference to the given int64 and assigns it to the DownloadBandwidth field.
func (o *DfsFTPShareACLReq) SetDownloadBandwidth(v int64) {
	o.DownloadBandwidth = &v
}

// GetDownloadEnabled returns the DownloadEnabled field value if set, zero value otherwise.
func (o *DfsFTPShareACLReq) GetDownloadEnabled() bool {
	if o == nil || IsNil(o.DownloadEnabled) {
		var ret bool
		return ret
	}
	return *o.DownloadEnabled
}

// GetDownloadEnabledOk returns a tuple with the DownloadEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFTPShareACLReq) GetDownloadEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DownloadEnabled) {
		return nil, false
	}
	return o.DownloadEnabled, true
}

// HasDownloadEnabled returns a boolean if a field has been set.
func (o *DfsFTPShareACLReq) HasDownloadEnabled() bool {
	if o != nil && !IsNil(o.DownloadEnabled) {
		return true
	}

	return false
}

// SetDownloadEnabled gets a reference to the given bool and assigns it to the DownloadEnabled field.
func (o *DfsFTPShareACLReq) SetDownloadEnabled(v bool) {
	o.DownloadEnabled = &v
}

// GetFsUserId returns the FsUserId field value if set, zero value otherwise.
func (o *DfsFTPShareACLReq) GetFsUserId() int64 {
	if o == nil || IsNil(o.FsUserId) {
		var ret int64
		return ret
	}
	return *o.FsUserId
}

// GetFsUserIdOk returns a tuple with the FsUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFTPShareACLReq) GetFsUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.FsUserId) {
		return nil, false
	}
	return o.FsUserId, true
}

// HasFsUserId returns a boolean if a field has been set.
func (o *DfsFTPShareACLReq) HasFsUserId() bool {
	if o != nil && !IsNil(o.FsUserId) {
		return true
	}

	return false
}

// SetFsUserId gets a reference to the given int64 and assigns it to the FsUserId field.
func (o *DfsFTPShareACLReq) SetFsUserId(v int64) {
	o.FsUserId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DfsFTPShareACLReq) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFTPShareACLReq) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DfsFTPShareACLReq) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DfsFTPShareACLReq) SetId(v int64) {
	o.Id = &v
}

// GetListEnabled returns the ListEnabled field value if set, zero value otherwise.
func (o *DfsFTPShareACLReq) GetListEnabled() bool {
	if o == nil || IsNil(o.ListEnabled) {
		var ret bool
		return ret
	}
	return *o.ListEnabled
}

// GetListEnabledOk returns a tuple with the ListEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFTPShareACLReq) GetListEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ListEnabled) {
		return nil, false
	}
	return o.ListEnabled, true
}

// HasListEnabled returns a boolean if a field has been set.
func (o *DfsFTPShareACLReq) HasListEnabled() bool {
	if o != nil && !IsNil(o.ListEnabled) {
		return true
	}

	return false
}

// SetListEnabled gets a reference to the given bool and assigns it to the ListEnabled field.
func (o *DfsFTPShareACLReq) SetListEnabled(v bool) {
	o.ListEnabled = &v
}

// GetRenameEnabled returns the RenameEnabled field value if set, zero value otherwise.
func (o *DfsFTPShareACLReq) GetRenameEnabled() bool {
	if o == nil || IsNil(o.RenameEnabled) {
		var ret bool
		return ret
	}
	return *o.RenameEnabled
}

// GetRenameEnabledOk returns a tuple with the RenameEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFTPShareACLReq) GetRenameEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RenameEnabled) {
		return nil, false
	}
	return o.RenameEnabled, true
}

// HasRenameEnabled returns a boolean if a field has been set.
func (o *DfsFTPShareACLReq) HasRenameEnabled() bool {
	if o != nil && !IsNil(o.RenameEnabled) {
		return true
	}

	return false
}

// SetRenameEnabled gets a reference to the given bool and assigns it to the RenameEnabled field.
func (o *DfsFTPShareACLReq) SetRenameEnabled(v bool) {
	o.RenameEnabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DfsFTPShareACLReq) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFTPShareACLReq) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DfsFTPShareACLReq) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DfsFTPShareACLReq) SetType(v string) {
	o.Type = &v
}

// GetUploadBandwidth returns the UploadBandwidth field value if set, zero value otherwise.
func (o *DfsFTPShareACLReq) GetUploadBandwidth() int64 {
	if o == nil || IsNil(o.UploadBandwidth) {
		var ret int64
		return ret
	}
	return *o.UploadBandwidth
}

// GetUploadBandwidthOk returns a tuple with the UploadBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFTPShareACLReq) GetUploadBandwidthOk() (*int64, bool) {
	if o == nil || IsNil(o.UploadBandwidth) {
		return nil, false
	}
	return o.UploadBandwidth, true
}

// HasUploadBandwidth returns a boolean if a field has been set.
func (o *DfsFTPShareACLReq) HasUploadBandwidth() bool {
	if o != nil && !IsNil(o.UploadBandwidth) {
		return true
	}

	return false
}

// SetUploadBandwidth gets a reference to the given int64 and assigns it to the UploadBandwidth field.
func (o *DfsFTPShareACLReq) SetUploadBandwidth(v int64) {
	o.UploadBandwidth = &v
}

// GetUploadEnabled returns the UploadEnabled field value if set, zero value otherwise.
func (o *DfsFTPShareACLReq) GetUploadEnabled() bool {
	if o == nil || IsNil(o.UploadEnabled) {
		var ret bool
		return ret
	}
	return *o.UploadEnabled
}

// GetUploadEnabledOk returns a tuple with the UploadEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFTPShareACLReq) GetUploadEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.UploadEnabled) {
		return nil, false
	}
	return o.UploadEnabled, true
}

// HasUploadEnabled returns a boolean if a field has been set.
func (o *DfsFTPShareACLReq) HasUploadEnabled() bool {
	if o != nil && !IsNil(o.UploadEnabled) {
		return true
	}

	return false
}

// SetUploadEnabled gets a reference to the given bool and assigns it to the UploadEnabled field.
func (o *DfsFTPShareACLReq) SetUploadEnabled(v bool) {
	o.UploadEnabled = &v
}

func (o DfsFTPShareACLReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsFTPShareACLReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminEnabled) {
		toSerialize["admin_enabled"] = o.AdminEnabled
	}
	if !IsNil(o.CreateEnabled) {
		toSerialize["create_enabled"] = o.CreateEnabled
	}
	if !IsNil(o.DeleteEnabled) {
		toSerialize["delete_enabled"] = o.DeleteEnabled
	}
	if !IsNil(o.DownloadBandwidth) {
		toSerialize["download_bandwidth"] = o.DownloadBandwidth
	}
	if !IsNil(o.DownloadEnabled) {
		toSerialize["download_enabled"] = o.DownloadEnabled
	}
	if !IsNil(o.FsUserId) {
		toSerialize["fs_user_id"] = o.FsUserId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ListEnabled) {
		toSerialize["list_enabled"] = o.ListEnabled
	}
	if !IsNil(o.RenameEnabled) {
		toSerialize["rename_enabled"] = o.RenameEnabled
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UploadBandwidth) {
		toSerialize["upload_bandwidth"] = o.UploadBandwidth
	}
	if !IsNil(o.UploadEnabled) {
		toSerialize["upload_enabled"] = o.UploadEnabled
	}
	return toSerialize, nil
}

type NullableDfsFTPShareACLReq struct {
	value *DfsFTPShareACLReq
	isSet bool
}

func (v NullableDfsFTPShareACLReq) Get() *DfsFTPShareACLReq {
	return v.value
}

func (v *NullableDfsFTPShareACLReq) Set(val *DfsFTPShareACLReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsFTPShareACLReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsFTPShareACLReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsFTPShareACLReq(val *DfsFTPShareACLReq) *NullableDfsFTPShareACLReq {
	return &NullableDfsFTPShareACLReq{value: val, isSet: true}
}

func (v NullableDfsFTPShareACLReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsFTPShareACLReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


