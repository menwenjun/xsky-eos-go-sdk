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

// checks if the DfsS3BucketCreateReqBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsS3BucketCreateReqBucket{}

// DfsS3BucketCreateReqBucket struct for DfsS3BucketCreateReqBucket
type DfsS3BucketCreateReqBucket struct {
	AllUserPermission *string `json:"all_user_permission,omitempty"`
	AuthUserPermission *string `json:"auth_user_permission,omitempty"`
	DataVerify *bool `json:"data_verify,omitempty"`
	// description of bucket
	Description *string `json:"description,omitempty"`
	EnableEtag *bool `json:"enable_etag,omitempty"`
	Name string `json:"name"`
	OwnerId int64 `json:"owner_id"`
	OwnerPermission *string `json:"owner_permission,omitempty"`
	Path string `json:"path"`
	RootfsId int64 `json:"rootfs_id"`
}

type _DfsS3BucketCreateReqBucket DfsS3BucketCreateReqBucket

// NewDfsS3BucketCreateReqBucket instantiates a new DfsS3BucketCreateReqBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsS3BucketCreateReqBucket(name string, ownerId int64, path string, rootfsId int64) *DfsS3BucketCreateReqBucket {
	this := DfsS3BucketCreateReqBucket{}
	this.Name = name
	this.OwnerId = ownerId
	this.Path = path
	this.RootfsId = rootfsId
	return &this
}

// NewDfsS3BucketCreateReqBucketWithDefaults instantiates a new DfsS3BucketCreateReqBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsS3BucketCreateReqBucketWithDefaults() *DfsS3BucketCreateReqBucket {
	this := DfsS3BucketCreateReqBucket{}
	return &this
}

// GetAllUserPermission returns the AllUserPermission field value if set, zero value otherwise.
func (o *DfsS3BucketCreateReqBucket) GetAllUserPermission() string {
	if o == nil || IsNil(o.AllUserPermission) {
		var ret string
		return ret
	}
	return *o.AllUserPermission
}

// GetAllUserPermissionOk returns a tuple with the AllUserPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketCreateReqBucket) GetAllUserPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.AllUserPermission) {
		return nil, false
	}
	return o.AllUserPermission, true
}

// HasAllUserPermission returns a boolean if a field has been set.
func (o *DfsS3BucketCreateReqBucket) HasAllUserPermission() bool {
	if o != nil && !IsNil(o.AllUserPermission) {
		return true
	}

	return false
}

// SetAllUserPermission gets a reference to the given string and assigns it to the AllUserPermission field.
func (o *DfsS3BucketCreateReqBucket) SetAllUserPermission(v string) {
	o.AllUserPermission = &v
}

// GetAuthUserPermission returns the AuthUserPermission field value if set, zero value otherwise.
func (o *DfsS3BucketCreateReqBucket) GetAuthUserPermission() string {
	if o == nil || IsNil(o.AuthUserPermission) {
		var ret string
		return ret
	}
	return *o.AuthUserPermission
}

// GetAuthUserPermissionOk returns a tuple with the AuthUserPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketCreateReqBucket) GetAuthUserPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.AuthUserPermission) {
		return nil, false
	}
	return o.AuthUserPermission, true
}

// HasAuthUserPermission returns a boolean if a field has been set.
func (o *DfsS3BucketCreateReqBucket) HasAuthUserPermission() bool {
	if o != nil && !IsNil(o.AuthUserPermission) {
		return true
	}

	return false
}

// SetAuthUserPermission gets a reference to the given string and assigns it to the AuthUserPermission field.
func (o *DfsS3BucketCreateReqBucket) SetAuthUserPermission(v string) {
	o.AuthUserPermission = &v
}

// GetDataVerify returns the DataVerify field value if set, zero value otherwise.
func (o *DfsS3BucketCreateReqBucket) GetDataVerify() bool {
	if o == nil || IsNil(o.DataVerify) {
		var ret bool
		return ret
	}
	return *o.DataVerify
}

// GetDataVerifyOk returns a tuple with the DataVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketCreateReqBucket) GetDataVerifyOk() (*bool, bool) {
	if o == nil || IsNil(o.DataVerify) {
		return nil, false
	}
	return o.DataVerify, true
}

// HasDataVerify returns a boolean if a field has been set.
func (o *DfsS3BucketCreateReqBucket) HasDataVerify() bool {
	if o != nil && !IsNil(o.DataVerify) {
		return true
	}

	return false
}

// SetDataVerify gets a reference to the given bool and assigns it to the DataVerify field.
func (o *DfsS3BucketCreateReqBucket) SetDataVerify(v bool) {
	o.DataVerify = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DfsS3BucketCreateReqBucket) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketCreateReqBucket) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DfsS3BucketCreateReqBucket) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DfsS3BucketCreateReqBucket) SetDescription(v string) {
	o.Description = &v
}

// GetEnableEtag returns the EnableEtag field value if set, zero value otherwise.
func (o *DfsS3BucketCreateReqBucket) GetEnableEtag() bool {
	if o == nil || IsNil(o.EnableEtag) {
		var ret bool
		return ret
	}
	return *o.EnableEtag
}

// GetEnableEtagOk returns a tuple with the EnableEtag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketCreateReqBucket) GetEnableEtagOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableEtag) {
		return nil, false
	}
	return o.EnableEtag, true
}

// HasEnableEtag returns a boolean if a field has been set.
func (o *DfsS3BucketCreateReqBucket) HasEnableEtag() bool {
	if o != nil && !IsNil(o.EnableEtag) {
		return true
	}

	return false
}

// SetEnableEtag gets a reference to the given bool and assigns it to the EnableEtag field.
func (o *DfsS3BucketCreateReqBucket) SetEnableEtag(v bool) {
	o.EnableEtag = &v
}

// GetName returns the Name field value
func (o *DfsS3BucketCreateReqBucket) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DfsS3BucketCreateReqBucket) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DfsS3BucketCreateReqBucket) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value
func (o *DfsS3BucketCreateReqBucket) GetOwnerId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *DfsS3BucketCreateReqBucket) GetOwnerIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *DfsS3BucketCreateReqBucket) SetOwnerId(v int64) {
	o.OwnerId = v
}

// GetOwnerPermission returns the OwnerPermission field value if set, zero value otherwise.
func (o *DfsS3BucketCreateReqBucket) GetOwnerPermission() string {
	if o == nil || IsNil(o.OwnerPermission) {
		var ret string
		return ret
	}
	return *o.OwnerPermission
}

// GetOwnerPermissionOk returns a tuple with the OwnerPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketCreateReqBucket) GetOwnerPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerPermission) {
		return nil, false
	}
	return o.OwnerPermission, true
}

// HasOwnerPermission returns a boolean if a field has been set.
func (o *DfsS3BucketCreateReqBucket) HasOwnerPermission() bool {
	if o != nil && !IsNil(o.OwnerPermission) {
		return true
	}

	return false
}

// SetOwnerPermission gets a reference to the given string and assigns it to the OwnerPermission field.
func (o *DfsS3BucketCreateReqBucket) SetOwnerPermission(v string) {
	o.OwnerPermission = &v
}

// GetPath returns the Path field value
func (o *DfsS3BucketCreateReqBucket) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DfsS3BucketCreateReqBucket) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DfsS3BucketCreateReqBucket) SetPath(v string) {
	o.Path = v
}

// GetRootfsId returns the RootfsId field value
func (o *DfsS3BucketCreateReqBucket) GetRootfsId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RootfsId
}

// GetRootfsIdOk returns a tuple with the RootfsId field value
// and a boolean to check if the value has been set.
func (o *DfsS3BucketCreateReqBucket) GetRootfsIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootfsId, true
}

// SetRootfsId sets field value
func (o *DfsS3BucketCreateReqBucket) SetRootfsId(v int64) {
	o.RootfsId = v
}

func (o DfsS3BucketCreateReqBucket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsS3BucketCreateReqBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllUserPermission) {
		toSerialize["all_user_permission"] = o.AllUserPermission
	}
	if !IsNil(o.AuthUserPermission) {
		toSerialize["auth_user_permission"] = o.AuthUserPermission
	}
	if !IsNil(o.DataVerify) {
		toSerialize["data_verify"] = o.DataVerify
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EnableEtag) {
		toSerialize["enable_etag"] = o.EnableEtag
	}
	toSerialize["name"] = o.Name
	toSerialize["owner_id"] = o.OwnerId
	if !IsNil(o.OwnerPermission) {
		toSerialize["owner_permission"] = o.OwnerPermission
	}
	toSerialize["path"] = o.Path
	toSerialize["rootfs_id"] = o.RootfsId
	return toSerialize, nil
}

func (o *DfsS3BucketCreateReqBucket) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"owner_id",
		"path",
		"rootfs_id",
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

	varDfsS3BucketCreateReqBucket := _DfsS3BucketCreateReqBucket{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsS3BucketCreateReqBucket)

	if err != nil {
		return err
	}

	*o = DfsS3BucketCreateReqBucket(varDfsS3BucketCreateReqBucket)

	return err
}

type NullableDfsS3BucketCreateReqBucket struct {
	value *DfsS3BucketCreateReqBucket
	isSet bool
}

func (v NullableDfsS3BucketCreateReqBucket) Get() *DfsS3BucketCreateReqBucket {
	return v.value
}

func (v *NullableDfsS3BucketCreateReqBucket) Set(val *DfsS3BucketCreateReqBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsS3BucketCreateReqBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsS3BucketCreateReqBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsS3BucketCreateReqBucket(val *DfsS3BucketCreateReqBucket) *NullableDfsS3BucketCreateReqBucket {
	return &NullableDfsS3BucketCreateReqBucket{value: val, isSet: true}
}

func (v NullableDfsS3BucketCreateReqBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsS3BucketCreateReqBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


