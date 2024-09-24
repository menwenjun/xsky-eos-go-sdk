# DNSGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DnsGatewayGroup** | Pointer to [**DNSGatewayGroupNestview**](DNSGatewayGroupNestview.md) |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDNSGateway

`func NewDNSGateway() *DNSGateway`

NewDNSGateway instantiates a new DNSGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSGatewayWithDefaults

`func NewDNSGatewayWithDefaults() *DNSGateway`

NewDNSGatewayWithDefaults instantiates a new DNSGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DNSGateway) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DNSGateway) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DNSGateway) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DNSGateway) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *DNSGateway) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DNSGateway) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DNSGateway) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DNSGateway) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DNSGateway) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DNSGateway) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DNSGateway) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DNSGateway) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDnsGatewayGroup

`func (o *DNSGateway) GetDnsGatewayGroup() DNSGatewayGroupNestview`

GetDnsGatewayGroup returns the DnsGatewayGroup field if non-nil, zero value otherwise.

### GetDnsGatewayGroupOk

`func (o *DNSGateway) GetDnsGatewayGroupOk() (*DNSGatewayGroupNestview, bool)`

GetDnsGatewayGroupOk returns a tuple with the DnsGatewayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsGatewayGroup

`func (o *DNSGateway) SetDnsGatewayGroup(v DNSGatewayGroupNestview)`

SetDnsGatewayGroup sets DnsGatewayGroup field to given value.

### HasDnsGatewayGroup

`func (o *DNSGateway) HasDnsGatewayGroup() bool`

HasDnsGatewayGroup returns a boolean if a field has been set.

### GetHost

`func (o *DNSGateway) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DNSGateway) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DNSGateway) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *DNSGateway) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *DNSGateway) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DNSGateway) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DNSGateway) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DNSGateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *DNSGateway) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DNSGateway) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DNSGateway) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DNSGateway) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DNSGateway) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DNSGateway) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DNSGateway) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DNSGateway) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


