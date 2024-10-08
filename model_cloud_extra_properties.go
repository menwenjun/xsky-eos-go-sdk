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

// checks if the CloudExtraProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudExtraProperties{}

// CloudExtraProperties CloudExtraProperties contains extra properties
type CloudExtraProperties struct {
	// csi driver versions
	CsiDriverVersions []string `json:"csi_driver_versions,omitempty"`
	// domain id for openstack
	DomainId *string `json:"domain_id,omitempty"`
	// domain name for openstack
	DomainName *string `json:"domain_name,omitempty"`
	// kubernetes version
	K8sVersion *string `json:"k8s_version,omitempty"`
	// regions for openstack
	Regions []string `json:"regions,omitempty"`
	// tenant id for openstack
	TenantId *string `json:"tenant_id,omitempty"`
	// tenant name for openstack
	TenantName *string `json:"tenant_name,omitempty"`
}

// NewCloudExtraProperties instantiates a new CloudExtraProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudExtraProperties() *CloudExtraProperties {
	this := CloudExtraProperties{}
	return &this
}

// NewCloudExtraPropertiesWithDefaults instantiates a new CloudExtraProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudExtraPropertiesWithDefaults() *CloudExtraProperties {
	this := CloudExtraProperties{}
	return &this
}

// GetCsiDriverVersions returns the CsiDriverVersions field value if set, zero value otherwise.
func (o *CloudExtraProperties) GetCsiDriverVersions() []string {
	if o == nil || IsNil(o.CsiDriverVersions) {
		var ret []string
		return ret
	}
	return o.CsiDriverVersions
}

// GetCsiDriverVersionsOk returns a tuple with the CsiDriverVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudExtraProperties) GetCsiDriverVersionsOk() ([]string, bool) {
	if o == nil || IsNil(o.CsiDriverVersions) {
		return nil, false
	}
	return o.CsiDriverVersions, true
}

// HasCsiDriverVersions returns a boolean if a field has been set.
func (o *CloudExtraProperties) HasCsiDriverVersions() bool {
	if o != nil && !IsNil(o.CsiDriverVersions) {
		return true
	}

	return false
}

// SetCsiDriverVersions gets a reference to the given []string and assigns it to the CsiDriverVersions field.
func (o *CloudExtraProperties) SetCsiDriverVersions(v []string) {
	o.CsiDriverVersions = v
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *CloudExtraProperties) GetDomainId() string {
	if o == nil || IsNil(o.DomainId) {
		var ret string
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudExtraProperties) GetDomainIdOk() (*string, bool) {
	if o == nil || IsNil(o.DomainId) {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *CloudExtraProperties) HasDomainId() bool {
	if o != nil && !IsNil(o.DomainId) {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given string and assigns it to the DomainId field.
func (o *CloudExtraProperties) SetDomainId(v string) {
	o.DomainId = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *CloudExtraProperties) GetDomainName() string {
	if o == nil || IsNil(o.DomainName) {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudExtraProperties) GetDomainNameOk() (*string, bool) {
	if o == nil || IsNil(o.DomainName) {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *CloudExtraProperties) HasDomainName() bool {
	if o != nil && !IsNil(o.DomainName) {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *CloudExtraProperties) SetDomainName(v string) {
	o.DomainName = &v
}

// GetK8sVersion returns the K8sVersion field value if set, zero value otherwise.
func (o *CloudExtraProperties) GetK8sVersion() string {
	if o == nil || IsNil(o.K8sVersion) {
		var ret string
		return ret
	}
	return *o.K8sVersion
}

// GetK8sVersionOk returns a tuple with the K8sVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudExtraProperties) GetK8sVersionOk() (*string, bool) {
	if o == nil || IsNil(o.K8sVersion) {
		return nil, false
	}
	return o.K8sVersion, true
}

// HasK8sVersion returns a boolean if a field has been set.
func (o *CloudExtraProperties) HasK8sVersion() bool {
	if o != nil && !IsNil(o.K8sVersion) {
		return true
	}

	return false
}

// SetK8sVersion gets a reference to the given string and assigns it to the K8sVersion field.
func (o *CloudExtraProperties) SetK8sVersion(v string) {
	o.K8sVersion = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *CloudExtraProperties) GetRegions() []string {
	if o == nil || IsNil(o.Regions) {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudExtraProperties) GetRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *CloudExtraProperties) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *CloudExtraProperties) SetRegions(v []string) {
	o.Regions = v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CloudExtraProperties) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudExtraProperties) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CloudExtraProperties) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CloudExtraProperties) SetTenantId(v string) {
	o.TenantId = &v
}

// GetTenantName returns the TenantName field value if set, zero value otherwise.
func (o *CloudExtraProperties) GetTenantName() string {
	if o == nil || IsNil(o.TenantName) {
		var ret string
		return ret
	}
	return *o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudExtraProperties) GetTenantNameOk() (*string, bool) {
	if o == nil || IsNil(o.TenantName) {
		return nil, false
	}
	return o.TenantName, true
}

// HasTenantName returns a boolean if a field has been set.
func (o *CloudExtraProperties) HasTenantName() bool {
	if o != nil && !IsNil(o.TenantName) {
		return true
	}

	return false
}

// SetTenantName gets a reference to the given string and assigns it to the TenantName field.
func (o *CloudExtraProperties) SetTenantName(v string) {
	o.TenantName = &v
}

func (o CloudExtraProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudExtraProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CsiDriverVersions) {
		toSerialize["csi_driver_versions"] = o.CsiDriverVersions
	}
	if !IsNil(o.DomainId) {
		toSerialize["domain_id"] = o.DomainId
	}
	if !IsNil(o.DomainName) {
		toSerialize["domain_name"] = o.DomainName
	}
	if !IsNil(o.K8sVersion) {
		toSerialize["k8s_version"] = o.K8sVersion
	}
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenant_id"] = o.TenantId
	}
	if !IsNil(o.TenantName) {
		toSerialize["tenant_name"] = o.TenantName
	}
	return toSerialize, nil
}

type NullableCloudExtraProperties struct {
	value *CloudExtraProperties
	isSet bool
}

func (v NullableCloudExtraProperties) Get() *CloudExtraProperties {
	return v.value
}

func (v *NullableCloudExtraProperties) Set(val *CloudExtraProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudExtraProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudExtraProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudExtraProperties(val *CloudExtraProperties) *NullableCloudExtraProperties {
	return &NullableCloudExtraProperties{value: val, isSet: true}
}

func (v NullableCloudExtraProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudExtraProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


