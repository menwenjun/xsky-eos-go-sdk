# IdentityPlatformExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**InsecureSkipVerify** | Pointer to **bool** |  | [optional] 
**LogoutRedirectUrl** | Pointer to **string** |  | [optional] 
**LogoutUrl** | Pointer to **string** |  | [optional] 
**UserKeys** | Pointer to [**IdentityPlatformExtraUserKeys**](IdentityPlatformExtraUserKeys.md) |  | [optional] 
**UserinfoUrl** | Pointer to **string** |  | [optional] 
**VerifiedUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewIdentityPlatformExtra

`func NewIdentityPlatformExtra() *IdentityPlatformExtra`

NewIdentityPlatformExtra instantiates a new IdentityPlatformExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityPlatformExtraWithDefaults

`func NewIdentityPlatformExtraWithDefaults() *IdentityPlatformExtra`

NewIdentityPlatformExtraWithDefaults instantiates a new IdentityPlatformExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *IdentityPlatformExtra) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *IdentityPlatformExtra) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *IdentityPlatformExtra) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *IdentityPlatformExtra) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetClientId

`func (o *IdentityPlatformExtra) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IdentityPlatformExtra) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IdentityPlatformExtra) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IdentityPlatformExtra) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *IdentityPlatformExtra) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IdentityPlatformExtra) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IdentityPlatformExtra) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *IdentityPlatformExtra) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetInsecureSkipVerify

`func (o *IdentityPlatformExtra) GetInsecureSkipVerify() bool`

GetInsecureSkipVerify returns the InsecureSkipVerify field if non-nil, zero value otherwise.

### GetInsecureSkipVerifyOk

`func (o *IdentityPlatformExtra) GetInsecureSkipVerifyOk() (*bool, bool)`

GetInsecureSkipVerifyOk returns a tuple with the InsecureSkipVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSkipVerify

`func (o *IdentityPlatformExtra) SetInsecureSkipVerify(v bool)`

SetInsecureSkipVerify sets InsecureSkipVerify field to given value.

### HasInsecureSkipVerify

`func (o *IdentityPlatformExtra) HasInsecureSkipVerify() bool`

HasInsecureSkipVerify returns a boolean if a field has been set.

### GetLogoutRedirectUrl

`func (o *IdentityPlatformExtra) GetLogoutRedirectUrl() string`

GetLogoutRedirectUrl returns the LogoutRedirectUrl field if non-nil, zero value otherwise.

### GetLogoutRedirectUrlOk

`func (o *IdentityPlatformExtra) GetLogoutRedirectUrlOk() (*string, bool)`

GetLogoutRedirectUrlOk returns a tuple with the LogoutRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutRedirectUrl

`func (o *IdentityPlatformExtra) SetLogoutRedirectUrl(v string)`

SetLogoutRedirectUrl sets LogoutRedirectUrl field to given value.

### HasLogoutRedirectUrl

`func (o *IdentityPlatformExtra) HasLogoutRedirectUrl() bool`

HasLogoutRedirectUrl returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *IdentityPlatformExtra) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *IdentityPlatformExtra) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *IdentityPlatformExtra) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *IdentityPlatformExtra) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetUserKeys

`func (o *IdentityPlatformExtra) GetUserKeys() IdentityPlatformExtraUserKeys`

GetUserKeys returns the UserKeys field if non-nil, zero value otherwise.

### GetUserKeysOk

`func (o *IdentityPlatformExtra) GetUserKeysOk() (*IdentityPlatformExtraUserKeys, bool)`

GetUserKeysOk returns a tuple with the UserKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKeys

`func (o *IdentityPlatformExtra) SetUserKeys(v IdentityPlatformExtraUserKeys)`

SetUserKeys sets UserKeys field to given value.

### HasUserKeys

`func (o *IdentityPlatformExtra) HasUserKeys() bool`

HasUserKeys returns a boolean if a field has been set.

### GetUserinfoUrl

`func (o *IdentityPlatformExtra) GetUserinfoUrl() string`

GetUserinfoUrl returns the UserinfoUrl field if non-nil, zero value otherwise.

### GetUserinfoUrlOk

`func (o *IdentityPlatformExtra) GetUserinfoUrlOk() (*string, bool)`

GetUserinfoUrlOk returns a tuple with the UserinfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserinfoUrl

`func (o *IdentityPlatformExtra) SetUserinfoUrl(v string)`

SetUserinfoUrl sets UserinfoUrl field to given value.

### HasUserinfoUrl

`func (o *IdentityPlatformExtra) HasUserinfoUrl() bool`

HasUserinfoUrl returns a boolean if a field has been set.

### GetVerifiedUrl

`func (o *IdentityPlatformExtra) GetVerifiedUrl() string`

GetVerifiedUrl returns the VerifiedUrl field if non-nil, zero value otherwise.

### GetVerifiedUrlOk

`func (o *IdentityPlatformExtra) GetVerifiedUrlOk() (*string, bool)`

GetVerifiedUrlOk returns a tuple with the VerifiedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedUrl

`func (o *IdentityPlatformExtra) SetVerifiedUrl(v string)`

SetVerifiedUrl sets VerifiedUrl field to given value.

### HasVerifiedUrl

`func (o *IdentityPlatformExtra) HasVerifiedUrl() bool`

HasVerifiedUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


