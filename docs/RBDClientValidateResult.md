# RBDClientValidateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminIp** | Pointer to **string** | admin ip of rbd client | [optional] 
**Hostname** | Pointer to **string** | hostname of rbd client, only appears when name_exist is true | [optional] 
**IpExist** | Pointer to **bool** | ip of rbd client already exist | [optional] 
**NameExist** | Pointer to **bool** | name of rbd client already exist | [optional] 
**PublicIp** | Pointer to **string** | public ip of rbd client | [optional] 
**TokenInvalid** | Pointer to **bool** | token is invalid | [optional] 
**Unreachable** | Pointer to **bool** | rbd client is unreachable or not | [optional] 

## Methods

### NewRBDClientValidateResult

`func NewRBDClientValidateResult() *RBDClientValidateResult`

NewRBDClientValidateResult instantiates a new RBDClientValidateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRBDClientValidateResultWithDefaults

`func NewRBDClientValidateResultWithDefaults() *RBDClientValidateResult`

NewRBDClientValidateResultWithDefaults instantiates a new RBDClientValidateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminIp

`func (o *RBDClientValidateResult) GetAdminIp() string`

GetAdminIp returns the AdminIp field if non-nil, zero value otherwise.

### GetAdminIpOk

`func (o *RBDClientValidateResult) GetAdminIpOk() (*string, bool)`

GetAdminIpOk returns a tuple with the AdminIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminIp

`func (o *RBDClientValidateResult) SetAdminIp(v string)`

SetAdminIp sets AdminIp field to given value.

### HasAdminIp

`func (o *RBDClientValidateResult) HasAdminIp() bool`

HasAdminIp returns a boolean if a field has been set.

### GetHostname

`func (o *RBDClientValidateResult) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *RBDClientValidateResult) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *RBDClientValidateResult) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *RBDClientValidateResult) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpExist

`func (o *RBDClientValidateResult) GetIpExist() bool`

GetIpExist returns the IpExist field if non-nil, zero value otherwise.

### GetIpExistOk

`func (o *RBDClientValidateResult) GetIpExistOk() (*bool, bool)`

GetIpExistOk returns a tuple with the IpExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpExist

`func (o *RBDClientValidateResult) SetIpExist(v bool)`

SetIpExist sets IpExist field to given value.

### HasIpExist

`func (o *RBDClientValidateResult) HasIpExist() bool`

HasIpExist returns a boolean if a field has been set.

### GetNameExist

`func (o *RBDClientValidateResult) GetNameExist() bool`

GetNameExist returns the NameExist field if non-nil, zero value otherwise.

### GetNameExistOk

`func (o *RBDClientValidateResult) GetNameExistOk() (*bool, bool)`

GetNameExistOk returns a tuple with the NameExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameExist

`func (o *RBDClientValidateResult) SetNameExist(v bool)`

SetNameExist sets NameExist field to given value.

### HasNameExist

`func (o *RBDClientValidateResult) HasNameExist() bool`

HasNameExist returns a boolean if a field has been set.

### GetPublicIp

`func (o *RBDClientValidateResult) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *RBDClientValidateResult) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *RBDClientValidateResult) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *RBDClientValidateResult) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetTokenInvalid

`func (o *RBDClientValidateResult) GetTokenInvalid() bool`

GetTokenInvalid returns the TokenInvalid field if non-nil, zero value otherwise.

### GetTokenInvalidOk

`func (o *RBDClientValidateResult) GetTokenInvalidOk() (*bool, bool)`

GetTokenInvalidOk returns a tuple with the TokenInvalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenInvalid

`func (o *RBDClientValidateResult) SetTokenInvalid(v bool)`

SetTokenInvalid sets TokenInvalid field to given value.

### HasTokenInvalid

`func (o *RBDClientValidateResult) HasTokenInvalid() bool`

HasTokenInvalid returns a boolean if a field has been set.

### GetUnreachable

`func (o *RBDClientValidateResult) GetUnreachable() bool`

GetUnreachable returns the Unreachable field if non-nil, zero value otherwise.

### GetUnreachableOk

`func (o *RBDClientValidateResult) GetUnreachableOk() (*bool, bool)`

GetUnreachableOk returns a tuple with the Unreachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreachable

`func (o *RBDClientValidateResult) SetUnreachable(v bool)`

SetUnreachable sets Unreachable field to given value.

### HasUnreachable

`func (o *RBDClientValidateResult) HasUnreachable() bool`

HasUnreachable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


