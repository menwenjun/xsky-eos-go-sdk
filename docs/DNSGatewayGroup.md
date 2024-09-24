# DNSGatewayGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDNSGatewayGroup

`func NewDNSGatewayGroup() *DNSGatewayGroup`

NewDNSGatewayGroup instantiates a new DNSGatewayGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSGatewayGroupWithDefaults

`func NewDNSGatewayGroupWithDefaults() *DNSGatewayGroup`

NewDNSGatewayGroupWithDefaults instantiates a new DNSGatewayGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DNSGatewayGroup) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DNSGatewayGroup) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DNSGatewayGroup) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DNSGatewayGroup) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *DNSGatewayGroup) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DNSGatewayGroup) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DNSGatewayGroup) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DNSGatewayGroup) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DNSGatewayGroup) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DNSGatewayGroup) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DNSGatewayGroup) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DNSGatewayGroup) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *DNSGatewayGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DNSGatewayGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DNSGatewayGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DNSGatewayGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DNSGatewayGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DNSGatewayGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DNSGatewayGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DNSGatewayGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *DNSGatewayGroup) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DNSGatewayGroup) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DNSGatewayGroup) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DNSGatewayGroup) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *DNSGatewayGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DNSGatewayGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DNSGatewayGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DNSGatewayGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DNSGatewayGroup) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DNSGatewayGroup) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DNSGatewayGroup) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DNSGatewayGroup) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


