# DNSGatewayGroupCreateReqGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostIds** | **[]int64** | dns gateway id list | 
**Name** | **string** | dns gateway group name | 
**Origin** | **string** | dns origin | 
**Port** | **int64** | dns service port | 
**Ttl** | Pointer to **int64** | dns ttl with zone | [optional] 
**VipGroups** | Pointer to [**[]DNSGatewayGroupCreateReqGroupVIPGroupsElt**](DNSGatewayGroupCreateReqGroupVIPGroupsElt.md) | dns vip groups | [optional] 

## Methods

### NewDNSGatewayGroupCreateReqGroup

`func NewDNSGatewayGroupCreateReqGroup(hostIds []int64, name string, origin string, port int64, ) *DNSGatewayGroupCreateReqGroup`

NewDNSGatewayGroupCreateReqGroup instantiates a new DNSGatewayGroupCreateReqGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSGatewayGroupCreateReqGroupWithDefaults

`func NewDNSGatewayGroupCreateReqGroupWithDefaults() *DNSGatewayGroupCreateReqGroup`

NewDNSGatewayGroupCreateReqGroupWithDefaults instantiates a new DNSGatewayGroupCreateReqGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostIds

`func (o *DNSGatewayGroupCreateReqGroup) GetHostIds() []int64`

GetHostIds returns the HostIds field if non-nil, zero value otherwise.

### GetHostIdsOk

`func (o *DNSGatewayGroupCreateReqGroup) GetHostIdsOk() (*[]int64, bool)`

GetHostIdsOk returns a tuple with the HostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIds

`func (o *DNSGatewayGroupCreateReqGroup) SetHostIds(v []int64)`

SetHostIds sets HostIds field to given value.


### GetName

`func (o *DNSGatewayGroupCreateReqGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DNSGatewayGroupCreateReqGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DNSGatewayGroupCreateReqGroup) SetName(v string)`

SetName sets Name field to given value.


### GetOrigin

`func (o *DNSGatewayGroupCreateReqGroup) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *DNSGatewayGroupCreateReqGroup) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *DNSGatewayGroupCreateReqGroup) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetPort

`func (o *DNSGatewayGroupCreateReqGroup) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DNSGatewayGroupCreateReqGroup) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DNSGatewayGroupCreateReqGroup) SetPort(v int64)`

SetPort sets Port field to given value.


### GetTtl

`func (o *DNSGatewayGroupCreateReqGroup) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DNSGatewayGroupCreateReqGroup) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DNSGatewayGroupCreateReqGroup) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DNSGatewayGroupCreateReqGroup) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetVipGroups

`func (o *DNSGatewayGroupCreateReqGroup) GetVipGroups() []DNSGatewayGroupCreateReqGroupVIPGroupsElt`

GetVipGroups returns the VipGroups field if non-nil, zero value otherwise.

### GetVipGroupsOk

`func (o *DNSGatewayGroupCreateReqGroup) GetVipGroupsOk() (*[]DNSGatewayGroupCreateReqGroupVIPGroupsElt, bool)`

GetVipGroupsOk returns a tuple with the VipGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipGroups

`func (o *DNSGatewayGroupCreateReqGroup) SetVipGroups(v []DNSGatewayGroupCreateReqGroupVIPGroupsElt)`

SetVipGroups sets VipGroups field to given value.

### HasVipGroups

`func (o *DNSGatewayGroupCreateReqGroup) HasVipGroups() bool`

HasVipGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


