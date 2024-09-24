# DNSLoadBalancePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDNSLoadBalancePolicy

`func NewDNSLoadBalancePolicy() *DNSLoadBalancePolicy`

NewDNSLoadBalancePolicy instantiates a new DNSLoadBalancePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSLoadBalancePolicyWithDefaults

`func NewDNSLoadBalancePolicyWithDefaults() *DNSLoadBalancePolicy`

NewDNSLoadBalancePolicyWithDefaults instantiates a new DNSLoadBalancePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *DNSLoadBalancePolicy) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DNSLoadBalancePolicy) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DNSLoadBalancePolicy) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DNSLoadBalancePolicy) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *DNSLoadBalancePolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DNSLoadBalancePolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DNSLoadBalancePolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DNSLoadBalancePolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DNSLoadBalancePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DNSLoadBalancePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DNSLoadBalancePolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DNSLoadBalancePolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceType

`func (o *DNSLoadBalancePolicy) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *DNSLoadBalancePolicy) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *DNSLoadBalancePolicy) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *DNSLoadBalancePolicy) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetUpdate

`func (o *DNSLoadBalancePolicy) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DNSLoadBalancePolicy) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DNSLoadBalancePolicy) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DNSLoadBalancePolicy) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


