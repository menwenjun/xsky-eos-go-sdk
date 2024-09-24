# S3LoadBalancerKeepalived

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InterfaceName** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Rid** | Pointer to [**S3LoadBalancerKeepalivedRID**](S3LoadBalancerKeepalivedRID.md) |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**S3LoadBalancer** | Pointer to [**S3LoadBalancerNestview**](S3LoadBalancerNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Vip** | Pointer to **string** |  | [optional] 
**VipMask** | Pointer to **int64** |  | [optional] 

## Methods

### NewS3LoadBalancerKeepalived

`func NewS3LoadBalancerKeepalived() *S3LoadBalancerKeepalived`

NewS3LoadBalancerKeepalived instantiates a new S3LoadBalancerKeepalived object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LoadBalancerKeepalivedWithDefaults

`func NewS3LoadBalancerKeepalivedWithDefaults() *S3LoadBalancerKeepalived`

NewS3LoadBalancerKeepalivedWithDefaults instantiates a new S3LoadBalancerKeepalived object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *S3LoadBalancerKeepalived) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *S3LoadBalancerKeepalived) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *S3LoadBalancerKeepalived) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *S3LoadBalancerKeepalived) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *S3LoadBalancerKeepalived) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *S3LoadBalancerKeepalived) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *S3LoadBalancerKeepalived) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *S3LoadBalancerKeepalived) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetHost

`func (o *S3LoadBalancerKeepalived) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *S3LoadBalancerKeepalived) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *S3LoadBalancerKeepalived) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *S3LoadBalancerKeepalived) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *S3LoadBalancerKeepalived) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3LoadBalancerKeepalived) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3LoadBalancerKeepalived) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *S3LoadBalancerKeepalived) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceName

`func (o *S3LoadBalancerKeepalived) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *S3LoadBalancerKeepalived) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *S3LoadBalancerKeepalived) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *S3LoadBalancerKeepalived) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetIp

`func (o *S3LoadBalancerKeepalived) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *S3LoadBalancerKeepalived) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *S3LoadBalancerKeepalived) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *S3LoadBalancerKeepalived) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetRid

`func (o *S3LoadBalancerKeepalived) GetRid() S3LoadBalancerKeepalivedRID`

GetRid returns the Rid field if non-nil, zero value otherwise.

### GetRidOk

`func (o *S3LoadBalancerKeepalived) GetRidOk() (*S3LoadBalancerKeepalivedRID, bool)`

GetRidOk returns a tuple with the Rid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRid

`func (o *S3LoadBalancerKeepalived) SetRid(v S3LoadBalancerKeepalivedRID)`

SetRid sets Rid field to given value.

### HasRid

`func (o *S3LoadBalancerKeepalived) HasRid() bool`

HasRid returns a boolean if a field has been set.

### GetRole

`func (o *S3LoadBalancerKeepalived) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *S3LoadBalancerKeepalived) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *S3LoadBalancerKeepalived) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *S3LoadBalancerKeepalived) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetS3LoadBalancer

`func (o *S3LoadBalancerKeepalived) GetS3LoadBalancer() S3LoadBalancerNestview`

GetS3LoadBalancer returns the S3LoadBalancer field if non-nil, zero value otherwise.

### GetS3LoadBalancerOk

`func (o *S3LoadBalancerKeepalived) GetS3LoadBalancerOk() (*S3LoadBalancerNestview, bool)`

GetS3LoadBalancerOk returns a tuple with the S3LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3LoadBalancer

`func (o *S3LoadBalancerKeepalived) SetS3LoadBalancer(v S3LoadBalancerNestview)`

SetS3LoadBalancer sets S3LoadBalancer field to given value.

### HasS3LoadBalancer

`func (o *S3LoadBalancerKeepalived) HasS3LoadBalancer() bool`

HasS3LoadBalancer returns a boolean if a field has been set.

### GetStatus

`func (o *S3LoadBalancerKeepalived) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *S3LoadBalancerKeepalived) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *S3LoadBalancerKeepalived) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *S3LoadBalancerKeepalived) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *S3LoadBalancerKeepalived) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *S3LoadBalancerKeepalived) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *S3LoadBalancerKeepalived) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *S3LoadBalancerKeepalived) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVip

`func (o *S3LoadBalancerKeepalived) GetVip() string`

GetVip returns the Vip field if non-nil, zero value otherwise.

### GetVipOk

`func (o *S3LoadBalancerKeepalived) GetVipOk() (*string, bool)`

GetVipOk returns a tuple with the Vip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVip

`func (o *S3LoadBalancerKeepalived) SetVip(v string)`

SetVip sets Vip field to given value.

### HasVip

`func (o *S3LoadBalancerKeepalived) HasVip() bool`

HasVip returns a boolean if a field has been set.

### GetVipMask

`func (o *S3LoadBalancerKeepalived) GetVipMask() int64`

GetVipMask returns the VipMask field if non-nil, zero value otherwise.

### GetVipMaskOk

`func (o *S3LoadBalancerKeepalived) GetVipMaskOk() (*int64, bool)`

GetVipMaskOk returns a tuple with the VipMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipMask

`func (o *S3LoadBalancerKeepalived) SetVipMask(v int64)`

SetVipMask sets VipMask field to given value.

### HasVipMask

`func (o *S3LoadBalancerKeepalived) HasVipMask() bool`

HasVipMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


