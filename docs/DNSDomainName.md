# DNSDomainName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**DnsLoadBalancePolicy** | Pointer to [**DNSLoadBalancePolicyNestview**](DNSLoadBalancePolicyNestview.md) |  | [optional] 
**DnsZone** | Pointer to [**DNSZone**](DNSZone.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ResourceId** | Pointer to **int64** |  | [optional] 
**ResourceName** | Pointer to **string** | resource_name only needs to be displayed on the UI | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDNSDomainName

`func NewDNSDomainName() *DNSDomainName`

NewDNSDomainName instantiates a new DNSDomainName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSDomainNameWithDefaults

`func NewDNSDomainNameWithDefaults() *DNSDomainName`

NewDNSDomainNameWithDefaults instantiates a new DNSDomainName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *DNSDomainName) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DNSDomainName) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DNSDomainName) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DNSDomainName) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDnsLoadBalancePolicy

`func (o *DNSDomainName) GetDnsLoadBalancePolicy() DNSLoadBalancePolicyNestview`

GetDnsLoadBalancePolicy returns the DnsLoadBalancePolicy field if non-nil, zero value otherwise.

### GetDnsLoadBalancePolicyOk

`func (o *DNSDomainName) GetDnsLoadBalancePolicyOk() (*DNSLoadBalancePolicyNestview, bool)`

GetDnsLoadBalancePolicyOk returns a tuple with the DnsLoadBalancePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsLoadBalancePolicy

`func (o *DNSDomainName) SetDnsLoadBalancePolicy(v DNSLoadBalancePolicyNestview)`

SetDnsLoadBalancePolicy sets DnsLoadBalancePolicy field to given value.

### HasDnsLoadBalancePolicy

`func (o *DNSDomainName) HasDnsLoadBalancePolicy() bool`

HasDnsLoadBalancePolicy returns a boolean if a field has been set.

### GetDnsZone

`func (o *DNSDomainName) GetDnsZone() DNSZone`

GetDnsZone returns the DnsZone field if non-nil, zero value otherwise.

### GetDnsZoneOk

`func (o *DNSDomainName) GetDnsZoneOk() (*DNSZone, bool)`

GetDnsZoneOk returns a tuple with the DnsZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZone

`func (o *DNSDomainName) SetDnsZone(v DNSZone)`

SetDnsZone sets DnsZone field to given value.

### HasDnsZone

`func (o *DNSDomainName) HasDnsZone() bool`

HasDnsZone returns a boolean if a field has been set.

### GetId

`func (o *DNSDomainName) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DNSDomainName) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DNSDomainName) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DNSDomainName) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DNSDomainName) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DNSDomainName) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DNSDomainName) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DNSDomainName) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceId

`func (o *DNSDomainName) GetResourceId() int64`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *DNSDomainName) GetResourceIdOk() (*int64, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *DNSDomainName) SetResourceId(v int64)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *DNSDomainName) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceName

`func (o *DNSDomainName) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *DNSDomainName) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *DNSDomainName) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *DNSDomainName) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceType

`func (o *DNSDomainName) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *DNSDomainName) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *DNSDomainName) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *DNSDomainName) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStatus

`func (o *DNSDomainName) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DNSDomainName) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DNSDomainName) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DNSDomainName) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DNSDomainName) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DNSDomainName) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DNSDomainName) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DNSDomainName) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


