# AccessTokenCreateReqAccessToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description of access token | [optional] 
**Name** | **string** | name of access token | 
**Roles** | **[]string** | roles of access token | 

## Methods

### NewAccessTokenCreateReqAccessToken

`func NewAccessTokenCreateReqAccessToken(name string, roles []string, ) *AccessTokenCreateReqAccessToken`

NewAccessTokenCreateReqAccessToken instantiates a new AccessTokenCreateReqAccessToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenCreateReqAccessTokenWithDefaults

`func NewAccessTokenCreateReqAccessTokenWithDefaults() *AccessTokenCreateReqAccessToken`

NewAccessTokenCreateReqAccessTokenWithDefaults instantiates a new AccessTokenCreateReqAccessToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AccessTokenCreateReqAccessToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessTokenCreateReqAccessToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessTokenCreateReqAccessToken) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessTokenCreateReqAccessToken) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AccessTokenCreateReqAccessToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessTokenCreateReqAccessToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessTokenCreateReqAccessToken) SetName(v string)`

SetName sets Name field to given value.


### GetRoles

`func (o *AccessTokenCreateReqAccessToken) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AccessTokenCreateReqAccessToken) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AccessTokenCreateReqAccessToken) SetRoles(v []string)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


