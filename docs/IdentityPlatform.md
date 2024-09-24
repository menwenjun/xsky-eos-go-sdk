# IdentityPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** | time of creating identity platform | [optional] 
**Enabled** | Pointer to **bool** | enable the identity platform or not | [optional] 
**Extra** | Pointer to **map[string]interface{}** |  | [optional] 
**FailureNum** | Pointer to **int64** | num of failed authorization | [optional] 
**Id** | Pointer to **int64** | id of identity platform | [optional] 
**Key** | Pointer to **string** | secret key of identity platform | [optional] 
**LoginPageEnabled** | Pointer to **bool** |  | [optional] 
**LogoutUrl** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | name of identity platform | [optional] 
**SuccessNum** | Pointer to **int64** | num of successful authorization | [optional] 
**Type** | Pointer to **string** | type of identity platform | [optional] 
**Update** | Pointer to **time.Time** | time of updating identity platform | [optional] 
**Url** | Pointer to **string** | url of identity platform | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** | vendor for type | [optional] 

## Methods

### NewIdentityPlatform

`func NewIdentityPlatform() *IdentityPlatform`

NewIdentityPlatform instantiates a new IdentityPlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityPlatformWithDefaults

`func NewIdentityPlatformWithDefaults() *IdentityPlatform`

NewIdentityPlatformWithDefaults instantiates a new IdentityPlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *IdentityPlatform) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *IdentityPlatform) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *IdentityPlatform) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *IdentityPlatform) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetEnabled

`func (o *IdentityPlatform) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentityPlatform) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentityPlatform) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IdentityPlatform) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExtra

`func (o *IdentityPlatform) GetExtra() map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *IdentityPlatform) GetExtraOk() (*map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *IdentityPlatform) SetExtra(v map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *IdentityPlatform) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetFailureNum

`func (o *IdentityPlatform) GetFailureNum() int64`

GetFailureNum returns the FailureNum field if non-nil, zero value otherwise.

### GetFailureNumOk

`func (o *IdentityPlatform) GetFailureNumOk() (*int64, bool)`

GetFailureNumOk returns a tuple with the FailureNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureNum

`func (o *IdentityPlatform) SetFailureNum(v int64)`

SetFailureNum sets FailureNum field to given value.

### HasFailureNum

`func (o *IdentityPlatform) HasFailureNum() bool`

HasFailureNum returns a boolean if a field has been set.

### GetId

`func (o *IdentityPlatform) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityPlatform) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityPlatform) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityPlatform) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *IdentityPlatform) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IdentityPlatform) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IdentityPlatform) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *IdentityPlatform) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLoginPageEnabled

`func (o *IdentityPlatform) GetLoginPageEnabled() bool`

GetLoginPageEnabled returns the LoginPageEnabled field if non-nil, zero value otherwise.

### GetLoginPageEnabledOk

`func (o *IdentityPlatform) GetLoginPageEnabledOk() (*bool, bool)`

GetLoginPageEnabledOk returns a tuple with the LoginPageEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginPageEnabled

`func (o *IdentityPlatform) SetLoginPageEnabled(v bool)`

SetLoginPageEnabled sets LoginPageEnabled field to given value.

### HasLoginPageEnabled

`func (o *IdentityPlatform) HasLoginPageEnabled() bool`

HasLoginPageEnabled returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *IdentityPlatform) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *IdentityPlatform) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *IdentityPlatform) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *IdentityPlatform) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetName

`func (o *IdentityPlatform) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityPlatform) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityPlatform) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityPlatform) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSuccessNum

`func (o *IdentityPlatform) GetSuccessNum() int64`

GetSuccessNum returns the SuccessNum field if non-nil, zero value otherwise.

### GetSuccessNumOk

`func (o *IdentityPlatform) GetSuccessNumOk() (*int64, bool)`

GetSuccessNumOk returns a tuple with the SuccessNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessNum

`func (o *IdentityPlatform) SetSuccessNum(v int64)`

SetSuccessNum sets SuccessNum field to given value.

### HasSuccessNum

`func (o *IdentityPlatform) HasSuccessNum() bool`

HasSuccessNum returns a boolean if a field has been set.

### GetType

`func (o *IdentityPlatform) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityPlatform) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityPlatform) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IdentityPlatform) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *IdentityPlatform) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *IdentityPlatform) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *IdentityPlatform) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *IdentityPlatform) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUrl

`func (o *IdentityPlatform) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IdentityPlatform) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IdentityPlatform) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IdentityPlatform) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUuid

`func (o *IdentityPlatform) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *IdentityPlatform) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *IdentityPlatform) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *IdentityPlatform) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendor

`func (o *IdentityPlatform) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *IdentityPlatform) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *IdentityPlatform) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *IdentityPlatform) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


