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

// checks if the DNSDomainName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DNSDomainName{}

// DNSDomainName DNSDomainName model. +X:model:generate;check_get +X:benchmark:
type DNSDomainName struct {
	Create *time.Time `json:"create,omitempty"`
	DnsLoadBalancePolicy *DNSLoadBalancePolicyNestview `json:"dns_load_balance_policy,omitempty"`
	DnsZone *DNSZone `json:"dns_zone,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ResourceId *int64 `json:"resource_id,omitempty"`
	// resource_name only needs to be displayed on the UI
	ResourceName *string `json:"resource_name,omitempty"`
	ResourceType *string `json:"resource_type,omitempty"`
	Status *string `json:"status,omitempty"`
	Update *time.Time `json:"update,omitempty"`
}

// NewDNSDomainName instantiates a new DNSDomainName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDNSDomainName() *DNSDomainName {
	this := DNSDomainName{}
	return &this
}

// NewDNSDomainNameWithDefaults instantiates a new DNSDomainName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDNSDomainNameWithDefaults() *DNSDomainName {
	this := DNSDomainName{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DNSDomainName) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSDomainName) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DNSDomainName) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DNSDomainName) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDnsLoadBalancePolicy returns the DnsLoadBalancePolicy field value if set, zero value otherwise.
func (o *DNSDomainName) GetDnsLoadBalancePolicy() DNSLoadBalancePolicyNestview {
	if o == nil || IsNil(o.DnsLoadBalancePolicy) {
		var ret DNSLoadBalancePolicyNestview
		return ret
	}
	return *o.DnsLoadBalancePolicy
}

// GetDnsLoadBalancePolicyOk returns a tuple with the DnsLoadBalancePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSDomainName) GetDnsLoadBalancePolicyOk() (*DNSLoadBalancePolicyNestview, bool) {
	if o == nil || IsNil(o.DnsLoadBalancePolicy) {
		return nil, false
	}
	return o.DnsLoadBalancePolicy, true
}

// HasDnsLoadBalancePolicy returns a boolean if a field has been set.
func (o *DNSDomainName) HasDnsLoadBalancePolicy() bool {
	if o != nil && !IsNil(o.DnsLoadBalancePolicy) {
		return true
	}

	return false
}

// SetDnsLoadBalancePolicy gets a reference to the given DNSLoadBalancePolicyNestview and assigns it to the DnsLoadBalancePolicy field.
func (o *DNSDomainName) SetDnsLoadBalancePolicy(v DNSLoadBalancePolicyNestview) {
	o.DnsLoadBalancePolicy = &v
}

// GetDnsZone returns the DnsZone field value if set, zero value otherwise.
func (o *DNSDomainName) GetDnsZone() DNSZone {
	if o == nil || IsNil(o.DnsZone) {
		var ret DNSZone
		return ret
	}
	return *o.DnsZone
}

// GetDnsZoneOk returns a tuple with the DnsZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSDomainName) GetDnsZoneOk() (*DNSZone, bool) {
	if o == nil || IsNil(o.DnsZone) {
		return nil, false
	}
	return o.DnsZone, true
}

// HasDnsZone returns a boolean if a field has been set.
func (o *DNSDomainName) HasDnsZone() bool {
	if o != nil && !IsNil(o.DnsZone) {
		return true
	}

	return false
}

// SetDnsZone gets a reference to the given DNSZone and assigns it to the DnsZone field.
func (o *DNSDomainName) SetDnsZone(v DNSZone) {
	o.DnsZone = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DNSDomainName) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSDomainName) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DNSDomainName) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DNSDomainName) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DNSDomainName) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSDomainName) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DNSDomainName) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DNSDomainName) SetName(v string) {
	o.Name = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *DNSDomainName) GetResourceId() int64 {
	if o == nil || IsNil(o.ResourceId) {
		var ret int64
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSDomainName) GetResourceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *DNSDomainName) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given int64 and assigns it to the ResourceId field.
func (o *DNSDomainName) SetResourceId(v int64) {
	o.ResourceId = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *DNSDomainName) GetResourceName() string {
	if o == nil || IsNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSDomainName) GetResourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceName) {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *DNSDomainName) HasResourceName() bool {
	if o != nil && !IsNil(o.ResourceName) {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *DNSDomainName) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *DNSDomainName) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSDomainName) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *DNSDomainName) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *DNSDomainName) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DNSDomainName) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSDomainName) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DNSDomainName) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DNSDomainName) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DNSDomainName) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSDomainName) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DNSDomainName) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *DNSDomainName) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o DNSDomainName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DNSDomainName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.DnsLoadBalancePolicy) {
		toSerialize["dns_load_balance_policy"] = o.DnsLoadBalancePolicy
	}
	if !IsNil(o.DnsZone) {
		toSerialize["dns_zone"] = o.DnsZone
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ResourceId) {
		toSerialize["resource_id"] = o.ResourceId
	}
	if !IsNil(o.ResourceName) {
		toSerialize["resource_name"] = o.ResourceName
	}
	if !IsNil(o.ResourceType) {
		toSerialize["resource_type"] = o.ResourceType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	return toSerialize, nil
}

type NullableDNSDomainName struct {
	value *DNSDomainName
	isSet bool
}

func (v NullableDNSDomainName) Get() *DNSDomainName {
	return v.value
}

func (v *NullableDNSDomainName) Set(val *DNSDomainName) {
	v.value = val
	v.isSet = true
}

func (v NullableDNSDomainName) IsSet() bool {
	return v.isSet
}

func (v *NullableDNSDomainName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDNSDomainName(val *DNSDomainName) *NullableDNSDomainName {
	return &NullableDNSDomainName{value: val, isSet: true}
}

func (v NullableDNSDomainName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDNSDomainName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


