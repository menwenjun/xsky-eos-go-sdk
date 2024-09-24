# AccessPathCreateReqAccessPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chap** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**HostIds** | Pointer to **[]int64** |  | [optional] 
**MappingGroups** | Pointer to [**[]MappingGroupReq**](MappingGroupReq.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NvmfType** | Pointer to **string** |  | [optional] 
**ProtectionDomainId** | Pointer to **int64** | Deprecated | [optional] 
**Targets** | Pointer to [**[]TargetReq**](TargetReq.md) |  | [optional] 
**Tname** | Pointer to **string** |  | [optional] 
**Tsecret** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessPathCreateReqAccessPath

`func NewAccessPathCreateReqAccessPath() *AccessPathCreateReqAccessPath`

NewAccessPathCreateReqAccessPath instantiates a new AccessPathCreateReqAccessPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPathCreateReqAccessPathWithDefaults

`func NewAccessPathCreateReqAccessPathWithDefaults() *AccessPathCreateReqAccessPath`

NewAccessPathCreateReqAccessPathWithDefaults instantiates a new AccessPathCreateReqAccessPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChap

`func (o *AccessPathCreateReqAccessPath) GetChap() bool`

GetChap returns the Chap field if non-nil, zero value otherwise.

### GetChapOk

`func (o *AccessPathCreateReqAccessPath) GetChapOk() (*bool, bool)`

GetChapOk returns a tuple with the Chap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChap

`func (o *AccessPathCreateReqAccessPath) SetChap(v bool)`

SetChap sets Chap field to given value.

### HasChap

`func (o *AccessPathCreateReqAccessPath) HasChap() bool`

HasChap returns a boolean if a field has been set.

### GetDescription

`func (o *AccessPathCreateReqAccessPath) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessPathCreateReqAccessPath) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessPathCreateReqAccessPath) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessPathCreateReqAccessPath) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHostIds

`func (o *AccessPathCreateReqAccessPath) GetHostIds() []int64`

GetHostIds returns the HostIds field if non-nil, zero value otherwise.

### GetHostIdsOk

`func (o *AccessPathCreateReqAccessPath) GetHostIdsOk() (*[]int64, bool)`

GetHostIdsOk returns a tuple with the HostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIds

`func (o *AccessPathCreateReqAccessPath) SetHostIds(v []int64)`

SetHostIds sets HostIds field to given value.

### HasHostIds

`func (o *AccessPathCreateReqAccessPath) HasHostIds() bool`

HasHostIds returns a boolean if a field has been set.

### GetMappingGroups

`func (o *AccessPathCreateReqAccessPath) GetMappingGroups() []MappingGroupReq`

GetMappingGroups returns the MappingGroups field if non-nil, zero value otherwise.

### GetMappingGroupsOk

`func (o *AccessPathCreateReqAccessPath) GetMappingGroupsOk() (*[]MappingGroupReq, bool)`

GetMappingGroupsOk returns a tuple with the MappingGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingGroups

`func (o *AccessPathCreateReqAccessPath) SetMappingGroups(v []MappingGroupReq)`

SetMappingGroups sets MappingGroups field to given value.

### HasMappingGroups

`func (o *AccessPathCreateReqAccessPath) HasMappingGroups() bool`

HasMappingGroups returns a boolean if a field has been set.

### GetName

`func (o *AccessPathCreateReqAccessPath) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessPathCreateReqAccessPath) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessPathCreateReqAccessPath) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessPathCreateReqAccessPath) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNvmfType

`func (o *AccessPathCreateReqAccessPath) GetNvmfType() string`

GetNvmfType returns the NvmfType field if non-nil, zero value otherwise.

### GetNvmfTypeOk

`func (o *AccessPathCreateReqAccessPath) GetNvmfTypeOk() (*string, bool)`

GetNvmfTypeOk returns a tuple with the NvmfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvmfType

`func (o *AccessPathCreateReqAccessPath) SetNvmfType(v string)`

SetNvmfType sets NvmfType field to given value.

### HasNvmfType

`func (o *AccessPathCreateReqAccessPath) HasNvmfType() bool`

HasNvmfType returns a boolean if a field has been set.

### GetProtectionDomainId

`func (o *AccessPathCreateReqAccessPath) GetProtectionDomainId() int64`

GetProtectionDomainId returns the ProtectionDomainId field if non-nil, zero value otherwise.

### GetProtectionDomainIdOk

`func (o *AccessPathCreateReqAccessPath) GetProtectionDomainIdOk() (*int64, bool)`

GetProtectionDomainIdOk returns a tuple with the ProtectionDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionDomainId

`func (o *AccessPathCreateReqAccessPath) SetProtectionDomainId(v int64)`

SetProtectionDomainId sets ProtectionDomainId field to given value.

### HasProtectionDomainId

`func (o *AccessPathCreateReqAccessPath) HasProtectionDomainId() bool`

HasProtectionDomainId returns a boolean if a field has been set.

### GetTargets

`func (o *AccessPathCreateReqAccessPath) GetTargets() []TargetReq`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *AccessPathCreateReqAccessPath) GetTargetsOk() (*[]TargetReq, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *AccessPathCreateReqAccessPath) SetTargets(v []TargetReq)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *AccessPathCreateReqAccessPath) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetTname

`func (o *AccessPathCreateReqAccessPath) GetTname() string`

GetTname returns the Tname field if non-nil, zero value otherwise.

### GetTnameOk

`func (o *AccessPathCreateReqAccessPath) GetTnameOk() (*string, bool)`

GetTnameOk returns a tuple with the Tname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTname

`func (o *AccessPathCreateReqAccessPath) SetTname(v string)`

SetTname sets Tname field to given value.

### HasTname

`func (o *AccessPathCreateReqAccessPath) HasTname() bool`

HasTname returns a boolean if a field has been set.

### GetTsecret

`func (o *AccessPathCreateReqAccessPath) GetTsecret() string`

GetTsecret returns the Tsecret field if non-nil, zero value otherwise.

### GetTsecretOk

`func (o *AccessPathCreateReqAccessPath) GetTsecretOk() (*string, bool)`

GetTsecretOk returns a tuple with the Tsecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsecret

`func (o *AccessPathCreateReqAccessPath) SetTsecret(v string)`

SetTsecret sets Tsecret field to given value.

### HasTsecret

`func (o *AccessPathCreateReqAccessPath) HasTsecret() bool`

HasTsecret returns a boolean if a field has been set.

### GetType

`func (o *AccessPathCreateReqAccessPath) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessPathCreateReqAccessPath) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessPathCreateReqAccessPath) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccessPathCreateReqAccessPath) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUid

`func (o *AccessPathCreateReqAccessPath) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AccessPathCreateReqAccessPath) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AccessPathCreateReqAccessPath) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AccessPathCreateReqAccessPath) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


