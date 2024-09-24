# ActionLogUserResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to **[]string** |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 

## Methods

### NewActionLogUserResp

`func NewActionLogUserResp() *ActionLogUserResp`

NewActionLogUserResp instantiates a new ActionLogUserResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionLogUserRespWithDefaults

`func NewActionLogUserRespWithDefaults() *ActionLogUserResp`

NewActionLogUserRespWithDefaults instantiates a new ActionLogUserResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *ActionLogUserResp) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ActionLogUserResp) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ActionLogUserResp) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ActionLogUserResp) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetUserId

`func (o *ActionLogUserResp) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ActionLogUserResp) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ActionLogUserResp) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ActionLogUserResp) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


