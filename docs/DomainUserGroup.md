# DomainUserGroup

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
**UserGroupId** | Pointer to **int64** |  | [optional] 

## Methods

### NewDomainUserGroup

`func NewDomainUserGroup() *DomainUserGroup`

NewDomainUserGroup instantiates a new DomainUserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainUserGroupWithDefaults

`func NewDomainUserGroupWithDefaults() *DomainUserGroup`

NewDomainUserGroupWithDefaults instantiates a new DomainUserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAd

`func (o *DomainUserGroup) GetAd() FSActiveDirectory`

GetAd returns the Ad field if non-nil, zero value otherwise.

### GetAdOk

`func (o *DomainUserGroup) GetAdOk() (*FSActiveDirectory, bool)`

GetAdOk returns a tuple with the Ad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAd

`func (o *DomainUserGroup) SetAd(v FSActiveDirectory)`

SetAd sets Ad field to given value.

### HasAd

`func (o *DomainUserGroup) HasAd() bool`

HasAd returns a boolean if a field has been set.

### GetCreate

`func (o *DomainUserGroup) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DomainUserGroup) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DomainUserGroup) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DomainUserGroup) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *DomainUserGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DomainUserGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DomainUserGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DomainUserGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DomainUserGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainUserGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainUserGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DomainUserGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLdap

`func (o *DomainUserGroup) GetLdap() FSLdap`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *DomainUserGroup) GetLdapOk() (*FSLdap, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *DomainUserGroup) SetLdap(v FSLdap)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *DomainUserGroup) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetName

`func (o *DomainUserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainUserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainUserGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainUserGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecurity

`func (o *DomainUserGroup) GetSecurity() string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *DomainUserGroup) GetSecurityOk() (*string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *DomainUserGroup) SetSecurity(v string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *DomainUserGroup) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetUpdate

`func (o *DomainUserGroup) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DomainUserGroup) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DomainUserGroup) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DomainUserGroup) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUserGroupId

`func (o *DomainUserGroup) GetUserGroupId() int64`

GetUserGroupId returns the UserGroupId field if non-nil, zero value otherwise.

### GetUserGroupIdOk

`func (o *DomainUserGroup) GetUserGroupIdOk() (*int64, bool)`

GetUserGroupIdOk returns a tuple with the UserGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupId

`func (o *DomainUserGroup) SetUserGroupId(v int64)`

SetUserGroupId sets UserGroupId field to given value.

### HasUserGroupId

`func (o *DomainUserGroup) HasUserGroupId() bool`

HasUserGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


