# DfsQuotaPathValidateRespPathValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistDirQuota** | Pointer to [**DfsQuota**](DfsQuota.md) |  | [optional] 
**IsDirEmpty** | Pointer to **bool** |  | [optional] 
**IsDirExist** | Pointer to **bool** |  | [optional] 
**IsExceedMaxQuota** | Pointer to **bool** |  | [optional] 
**IsQuotaExist** | Pointer to **bool** |  | [optional] 
**UsedFiles** | Pointer to **int64** |  | [optional] 
**UsedKbyte** | Pointer to **int64** |  | [optional] 
**UserGroups** | Pointer to [**[]FSUserGroup**](FSUserGroup.md) |  | [optional] 
**Users** | Pointer to [**[]FSUser**](FSUser.md) |  | [optional] 

## Methods

### NewDfsQuotaPathValidateRespPathValidation

`func NewDfsQuotaPathValidateRespPathValidation() *DfsQuotaPathValidateRespPathValidation`

NewDfsQuotaPathValidateRespPathValidation instantiates a new DfsQuotaPathValidateRespPathValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsQuotaPathValidateRespPathValidationWithDefaults

`func NewDfsQuotaPathValidateRespPathValidationWithDefaults() *DfsQuotaPathValidateRespPathValidation`

NewDfsQuotaPathValidateRespPathValidationWithDefaults instantiates a new DfsQuotaPathValidateRespPathValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistDirQuota

`func (o *DfsQuotaPathValidateRespPathValidation) GetExistDirQuota() DfsQuota`

GetExistDirQuota returns the ExistDirQuota field if non-nil, zero value otherwise.

### GetExistDirQuotaOk

`func (o *DfsQuotaPathValidateRespPathValidation) GetExistDirQuotaOk() (*DfsQuota, bool)`

GetExistDirQuotaOk returns a tuple with the ExistDirQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistDirQuota

`func (o *DfsQuotaPathValidateRespPathValidation) SetExistDirQuota(v DfsQuota)`

SetExistDirQuota sets ExistDirQuota field to given value.

### HasExistDirQuota

`func (o *DfsQuotaPathValidateRespPathValidation) HasExistDirQuota() bool`

HasExistDirQuota returns a boolean if a field has been set.

### GetIsDirEmpty

`func (o *DfsQuotaPathValidateRespPathValidation) GetIsDirEmpty() bool`

GetIsDirEmpty returns the IsDirEmpty field if non-nil, zero value otherwise.

### GetIsDirEmptyOk

`func (o *DfsQuotaPathValidateRespPathValidation) GetIsDirEmptyOk() (*bool, bool)`

GetIsDirEmptyOk returns a tuple with the IsDirEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirEmpty

`func (o *DfsQuotaPathValidateRespPathValidation) SetIsDirEmpty(v bool)`

SetIsDirEmpty sets IsDirEmpty field to given value.

### HasIsDirEmpty

`func (o *DfsQuotaPathValidateRespPathValidation) HasIsDirEmpty() bool`

HasIsDirEmpty returns a boolean if a field has been set.

### GetIsDirExist

`func (o *DfsQuotaPathValidateRespPathValidation) GetIsDirExist() bool`

GetIsDirExist returns the IsDirExist field if non-nil, zero value otherwise.

### GetIsDirExistOk

`func (o *DfsQuotaPathValidateRespPathValidation) GetIsDirExistOk() (*bool, bool)`

GetIsDirExistOk returns a tuple with the IsDirExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirExist

`func (o *DfsQuotaPathValidateRespPathValidation) SetIsDirExist(v bool)`

SetIsDirExist sets IsDirExist field to given value.

### HasIsDirExist

`func (o *DfsQuotaPathValidateRespPathValidation) HasIsDirExist() bool`

HasIsDirExist returns a boolean if a field has been set.

### GetIsExceedMaxQuota

`func (o *DfsQuotaPathValidateRespPathValidation) GetIsExceedMaxQuota() bool`

GetIsExceedMaxQuota returns the IsExceedMaxQuota field if non-nil, zero value otherwise.

### GetIsExceedMaxQuotaOk

`func (o *DfsQuotaPathValidateRespPathValidation) GetIsExceedMaxQuotaOk() (*bool, bool)`

GetIsExceedMaxQuotaOk returns a tuple with the IsExceedMaxQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExceedMaxQuota

`func (o *DfsQuotaPathValidateRespPathValidation) SetIsExceedMaxQuota(v bool)`

SetIsExceedMaxQuota sets IsExceedMaxQuota field to given value.

### HasIsExceedMaxQuota

`func (o *DfsQuotaPathValidateRespPathValidation) HasIsExceedMaxQuota() bool`

HasIsExceedMaxQuota returns a boolean if a field has been set.

### GetIsQuotaExist

`func (o *DfsQuotaPathValidateRespPathValidation) GetIsQuotaExist() bool`

GetIsQuotaExist returns the IsQuotaExist field if non-nil, zero value otherwise.

### GetIsQuotaExistOk

`func (o *DfsQuotaPathValidateRespPathValidation) GetIsQuotaExistOk() (*bool, bool)`

GetIsQuotaExistOk returns a tuple with the IsQuotaExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuotaExist

`func (o *DfsQuotaPathValidateRespPathValidation) SetIsQuotaExist(v bool)`

SetIsQuotaExist sets IsQuotaExist field to given value.

### HasIsQuotaExist

`func (o *DfsQuotaPathValidateRespPathValidation) HasIsQuotaExist() bool`

HasIsQuotaExist returns a boolean if a field has been set.

### GetUsedFiles

`func (o *DfsQuotaPathValidateRespPathValidation) GetUsedFiles() int64`

GetUsedFiles returns the UsedFiles field if non-nil, zero value otherwise.

### GetUsedFilesOk

`func (o *DfsQuotaPathValidateRespPathValidation) GetUsedFilesOk() (*int64, bool)`

GetUsedFilesOk returns a tuple with the UsedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedFiles

`func (o *DfsQuotaPathValidateRespPathValidation) SetUsedFiles(v int64)`

SetUsedFiles sets UsedFiles field to given value.

### HasUsedFiles

`func (o *DfsQuotaPathValidateRespPathValidation) HasUsedFiles() bool`

HasUsedFiles returns a boolean if a field has been set.

### GetUsedKbyte

`func (o *DfsQuotaPathValidateRespPathValidation) GetUsedKbyte() int64`

GetUsedKbyte returns the UsedKbyte field if non-nil, zero value otherwise.

### GetUsedKbyteOk

`func (o *DfsQuotaPathValidateRespPathValidation) GetUsedKbyteOk() (*int64, bool)`

GetUsedKbyteOk returns a tuple with the UsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte

`func (o *DfsQuotaPathValidateRespPathValidation) SetUsedKbyte(v int64)`

SetUsedKbyte sets UsedKbyte field to given value.

### HasUsedKbyte

`func (o *DfsQuotaPathValidateRespPathValidation) HasUsedKbyte() bool`

HasUsedKbyte returns a boolean if a field has been set.

### GetUserGroups

`func (o *DfsQuotaPathValidateRespPathValidation) GetUserGroups() []FSUserGroup`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *DfsQuotaPathValidateRespPathValidation) GetUserGroupsOk() (*[]FSUserGroup, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *DfsQuotaPathValidateRespPathValidation) SetUserGroups(v []FSUserGroup)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *DfsQuotaPathValidateRespPathValidation) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### GetUsers

`func (o *DfsQuotaPathValidateRespPathValidation) GetUsers() []FSUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *DfsQuotaPathValidateRespPathValidation) GetUsersOk() (*[]FSUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *DfsQuotaPathValidateRespPathValidation) SetUsers(v []FSUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *DfsQuotaPathValidateRespPathValidation) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


