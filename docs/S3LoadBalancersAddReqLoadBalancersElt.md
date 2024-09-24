# S3LoadBalancersAddReqLoadBalancersElt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | **int64** | host of load balancer | 
**InterfaceName** | Pointer to **string** | vip will be bound to interface, exclusive to ip | [optional] 
**Ip** | Pointer to **string** | vip will be bound to interface of the gateway ip, exclusive to interface_name | [optional] 
**Vip** | **string** | vip of load balancer | 

## Methods

### NewS3LoadBalancersAddReqLoadBalancersElt

`func NewS3LoadBalancersAddReqLoadBalancersElt(hostId int64, vip string, ) *S3LoadBalancersAddReqLoadBalancersElt`

NewS3LoadBalancersAddReqLoadBalancersElt instantiates a new S3LoadBalancersAddReqLoadBalancersElt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LoadBalancersAddReqLoadBalancersEltWithDefaults

`func NewS3LoadBalancersAddReqLoadBalancersEltWithDefaults() *S3LoadBalancersAddReqLoadBalancersElt`

NewS3LoadBalancersAddReqLoadBalancersEltWithDefaults instantiates a new S3LoadBalancersAddReqLoadBalancersElt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *S3LoadBalancersAddReqLoadBalancersElt) GetHostId() int64`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *S3LoadBalancersAddReqLoadBalancersElt) GetHostIdOk() (*int64, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *S3LoadBalancersAddReqLoadBalancersElt) SetHostId(v int64)`

SetHostId sets HostId field to given value.


### GetInterfaceName

`func (o *S3LoadBalancersAddReqLoadBalancersElt) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *S3LoadBalancersAddReqLoadBalancersElt) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *S3LoadBalancersAddReqLoadBalancersElt) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *S3LoadBalancersAddReqLoadBalancersElt) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetIp

`func (o *S3LoadBalancersAddReqLoadBalancersElt) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *S3LoadBalancersAddReqLoadBalancersElt) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *S3LoadBalancersAddReqLoadBalancersElt) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *S3LoadBalancersAddReqLoadBalancersElt) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetVip

`func (o *S3LoadBalancersAddReqLoadBalancersElt) GetVip() string`

GetVip returns the Vip field if non-nil, zero value otherwise.

### GetVipOk

`func (o *S3LoadBalancersAddReqLoadBalancersElt) GetVipOk() (*string, bool)`

GetVipOk returns a tuple with the Vip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVip

`func (o *S3LoadBalancersAddReqLoadBalancersElt) SetVip(v string)`

SetVip sets Vip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


