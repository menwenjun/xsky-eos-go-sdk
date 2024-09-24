# DomainUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ad** | Pointer to [**FSActiveDirectory**](FSActiveDirectory.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Ldap** | Pointer to [**FSLdap**](FSLdap.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Security** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 

## Methods

### NewDomainUser

`func NewDomainUser() *DomainUser`

NewDomainUser instantiates a new DomainUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainUserWithDefaults

`func NewDomainUserWithDefaults() *DomainUser`

NewDomainUserWithDefaults instantiates a new DomainUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAd

`func (o *DomainUser) GetAd() FSActiveDirectory`

GetAd returns the Ad field if non-nil, zero value otherwise.

### GetAdOk

`func (o *DomainUser) GetAdOk() (*FSActiveDirectory, bool)`

GetAdOk returns a tuple with the Ad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAd

`func (o *DomainUser) SetAd(v FSActiveDirectory)`

SetAd sets Ad field to given value.

### HasAd

`func (o *DomainUser) HasAd() bool`

HasAd returns a boolean if a field has been set.

### GetCreate

`func (o *DomainUser) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DomainUser) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DomainUser) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DomainUser) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *DomainUser) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DomainUser) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DomainUser) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DomainUser) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DomainUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DomainUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLdap

`func (o *DomainUser) GetLdap() FSLdap`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *DomainUser) GetLdapOk() (*FSLdap, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *DomainUser) SetLdap(v FSLdap)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *DomainUser) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetName

`func (o *DomainUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecurity

`func (o *DomainUser) GetSecurity() string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *DomainUser) GetSecurityOk() (*string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *DomainUser) SetSecurity(v string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *DomainUser) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetUpdate

`func (o *DomainUser) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DomainUser) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DomainUser) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DomainUser) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUserId

`func (o *DomainUser) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DomainUser) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DomainUser) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DomainUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


