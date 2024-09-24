# FSUserUpdateSecondaryGroupsReqUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description of file storage user | [optional] 
**PrimaryGroupId** | Pointer to **int64** | primary group of user | [optional] 
**UserGroupIds** | Pointer to **[]int64** | secondary group of user | [optional] 

## Methods

### NewFSUserUpdateSecondaryGroupsReqUser

`func NewFSUserUpdateSecondaryGroupsReqUser() *FSUserUpdateSecondaryGroupsReqUser`

NewFSUserUpdateSecondaryGroupsReqUser instantiates a new FSUserUpdateSecondaryGroupsReqUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFSUserUpdateSecondaryGroupsReqUserWithDefaults

`func NewFSUserUpdateSecondaryGroupsReqUserWithDefaults() *FSUserUpdateSecondaryGroupsReqUser`

NewFSUserUpdateSecondaryGroupsReqUserWithDefaults instantiates a new FSUserUpdateSecondaryGroupsReqUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FSUserUpdateSecondaryGroupsReqUser) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FSUserUpdateSecondaryGroupsReqUser) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FSUserUpdateSecondaryGroupsReqUser) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FSUserUpdateSecondaryGroupsReqUser) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrimaryGroupId

`func (o *FSUserUpdateSecondaryGroupsReqUser) GetPrimaryGroupId() int64`

GetPrimaryGroupId returns the PrimaryGroupId field if non-nil, zero value otherwise.

### GetPrimaryGroupIdOk

`func (o *FSUserUpdateSecondaryGroupsReqUser) GetPrimaryGroupIdOk() (*int64, bool)`

GetPrimaryGroupIdOk returns a tuple with the PrimaryGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGroupId

`func (o *FSUserUpdateSecondaryGroupsReqUser) SetPrimaryGroupId(v int64)`

SetPrimaryGroupId sets PrimaryGroupId field to given value.

### HasPrimaryGroupId

`func (o *FSUserUpdateSecondaryGroupsReqUser) HasPrimaryGroupId() bool`

HasPrimaryGroupId returns a boolean if a field has been set.

### GetUserGroupIds

`func (o *FSUserUpdateSecondaryGroupsReqUser) GetUserGroupIds() []int64`

GetUserGroupIds returns the UserGroupIds field if non-nil, zero value otherwise.

### GetUserGroupIdsOk

`func (o *FSUserUpdateSecondaryGroupsReqUser) GetUserGroupIdsOk() (*[]int64, bool)`

GetUserGroupIdsOk returns a tuple with the UserGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupIds

`func (o *FSUserUpdateSecondaryGroupsReqUser) SetUserGroupIds(v []int64)`

SetUserGroupIds sets UserGroupIds field to given value.

### HasUserGroupIds

`func (o *FSUserUpdateSecondaryGroupsReqUser) HasUserGroupIds() bool`

HasUserGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


