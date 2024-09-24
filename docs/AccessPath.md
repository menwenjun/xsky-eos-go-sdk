# AccessPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**BlockVolumeNum** | Pointer to **int64** |  | [optional] 
**Chap** | Pointer to **bool** |  | [optional] 
**ClientGroupNum** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**HasAluaVolumes** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Iqn** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NvmfType** | Pointer to **string** |  | [optional] 
**ProtectionDomain** | Pointer to [**ProtectionDomainNestview**](ProtectionDomainNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tname** | Pointer to **string** |  | [optional] 
**Tsecret** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**VipGroupNum** | Pointer to **int64** |  | [optional] 

## Methods

### NewAccessPath

`func NewAccessPath() *AccessPath`

NewAccessPath instantiates a new AccessPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPathWithDefaults

`func NewAccessPathWithDefaults() *AccessPath`

NewAccessPathWithDefaults instantiates a new AccessPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *AccessPath) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *AccessPath) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *AccessPath) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *AccessPath) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetBlockVolumeNum

`func (o *AccessPath) GetBlockVolumeNum() int64`

GetBlockVolumeNum returns the BlockVolumeNum field if non-nil, zero value otherwise.

### GetBlockVolumeNumOk

`func (o *AccessPath) GetBlockVolumeNumOk() (*int64, bool)`

GetBlockVolumeNumOk returns a tuple with the BlockVolumeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeNum

`func (o *AccessPath) SetBlockVolumeNum(v int64)`

SetBlockVolumeNum sets BlockVolumeNum field to given value.

### HasBlockVolumeNum

`func (o *AccessPath) HasBlockVolumeNum() bool`

HasBlockVolumeNum returns a boolean if a field has been set.

### GetChap

`func (o *AccessPath) GetChap() bool`

GetChap returns the Chap field if non-nil, zero value otherwise.

### GetChapOk

`func (o *AccessPath) GetChapOk() (*bool, bool)`

GetChapOk returns a tuple with the Chap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChap

`func (o *AccessPath) SetChap(v bool)`

SetChap sets Chap field to given value.

### HasChap

`func (o *AccessPath) HasChap() bool`

HasChap returns a boolean if a field has been set.

### GetClientGroupNum

`func (o *AccessPath) GetClientGroupNum() int64`

GetClientGroupNum returns the ClientGroupNum field if non-nil, zero value otherwise.

### GetClientGroupNumOk

`func (o *AccessPath) GetClientGroupNumOk() (*int64, bool)`

GetClientGroupNumOk returns a tuple with the ClientGroupNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientGroupNum

`func (o *AccessPath) SetClientGroupNum(v int64)`

SetClientGroupNum sets ClientGroupNum field to given value.

### HasClientGroupNum

`func (o *AccessPath) HasClientGroupNum() bool`

HasClientGroupNum returns a boolean if a field has been set.

### GetCluster

`func (o *AccessPath) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *AccessPath) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *AccessPath) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *AccessPath) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *AccessPath) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *AccessPath) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *AccessPath) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *AccessPath) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *AccessPath) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessPath) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessPath) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessPath) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHasAluaVolumes

`func (o *AccessPath) GetHasAluaVolumes() bool`

GetHasAluaVolumes returns the HasAluaVolumes field if non-nil, zero value otherwise.

### GetHasAluaVolumesOk

`func (o *AccessPath) GetHasAluaVolumesOk() (*bool, bool)`

GetHasAluaVolumesOk returns a tuple with the HasAluaVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAluaVolumes

`func (o *AccessPath) SetHasAluaVolumes(v bool)`

SetHasAluaVolumes sets HasAluaVolumes field to given value.

### HasHasAluaVolumes

`func (o *AccessPath) HasHasAluaVolumes() bool`

HasHasAluaVolumes returns a boolean if a field has been set.

### GetHidden

`func (o *AccessPath) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *AccessPath) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *AccessPath) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *AccessPath) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *AccessPath) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessPath) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessPath) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AccessPath) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIqn

`func (o *AccessPath) GetIqn() string`

GetIqn returns the Iqn field if non-nil, zero value otherwise.

### GetIqnOk

`func (o *AccessPath) GetIqnOk() (*string, bool)`

GetIqnOk returns a tuple with the Iqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqn

`func (o *AccessPath) SetIqn(v string)`

SetIqn sets Iqn field to given value.

### HasIqn

`func (o *AccessPath) HasIqn() bool`

HasIqn returns a boolean if a field has been set.

### GetName

`func (o *AccessPath) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessPath) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessPath) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessPath) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNvmfType

`func (o *AccessPath) GetNvmfType() string`

GetNvmfType returns the NvmfType field if non-nil, zero value otherwise.

### GetNvmfTypeOk

`func (o *AccessPath) GetNvmfTypeOk() (*string, bool)`

GetNvmfTypeOk returns a tuple with the NvmfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvmfType

`func (o *AccessPath) SetNvmfType(v string)`

SetNvmfType sets NvmfType field to given value.

### HasNvmfType

`func (o *AccessPath) HasNvmfType() bool`

HasNvmfType returns a boolean if a field has been set.

### GetProtectionDomain

`func (o *AccessPath) GetProtectionDomain() ProtectionDomainNestview`

GetProtectionDomain returns the ProtectionDomain field if non-nil, zero value otherwise.

### GetProtectionDomainOk

`func (o *AccessPath) GetProtectionDomainOk() (*ProtectionDomainNestview, bool)`

GetProtectionDomainOk returns a tuple with the ProtectionDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionDomain

`func (o *AccessPath) SetProtectionDomain(v ProtectionDomainNestview)`

SetProtectionDomain sets ProtectionDomain field to given value.

### HasProtectionDomain

`func (o *AccessPath) HasProtectionDomain() bool`

HasProtectionDomain returns a boolean if a field has been set.

### GetStatus

`func (o *AccessPath) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessPath) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccessPath) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccessPath) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTname

`func (o *AccessPath) GetTname() string`

GetTname returns the Tname field if non-nil, zero value otherwise.

### GetTnameOk

`func (o *AccessPath) GetTnameOk() (*string, bool)`

GetTnameOk returns a tuple with the Tname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTname

`func (o *AccessPath) SetTname(v string)`

SetTname sets Tname field to given value.

### HasTname

`func (o *AccessPath) HasTname() bool`

HasTname returns a boolean if a field has been set.

### GetTsecret

`func (o *AccessPath) GetTsecret() string`

GetTsecret returns the Tsecret field if non-nil, zero value otherwise.

### GetTsecretOk

`func (o *AccessPath) GetTsecretOk() (*string, bool)`

GetTsecretOk returns a tuple with the Tsecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsecret

`func (o *AccessPath) SetTsecret(v string)`

SetTsecret sets Tsecret field to given value.

### HasTsecret

`func (o *AccessPath) HasTsecret() bool`

HasTsecret returns a boolean if a field has been set.

### GetType

`func (o *AccessPath) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessPath) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessPath) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccessPath) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUid

`func (o *AccessPath) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AccessPath) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AccessPath) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AccessPath) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUpdate

`func (o *AccessPath) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *AccessPath) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *AccessPath) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *AccessPath) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVipGroupNum

`func (o *AccessPath) GetVipGroupNum() int64`

GetVipGroupNum returns the VipGroupNum field if non-nil, zero value otherwise.

### GetVipGroupNumOk

`func (o *AccessPath) GetVipGroupNumOk() (*int64, bool)`

GetVipGroupNumOk returns a tuple with the VipGroupNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipGroupNum

`func (o *AccessPath) SetVipGroupNum(v int64)`

SetVipGroupNum sets VipGroupNum field to given value.

### HasVipGroupNum

`func (o *AccessPath) HasVipGroupNum() bool`

HasVipGroupNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


