# DomainUserNestview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ad** | Pointer to [**DomainUserGroupNestviewAd**](DomainUserGroupNestviewAd.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Ldap** | Pointer to [**DomainUserGroupNestviewLdap**](DomainUserGroupNestviewLdap.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainUserNestview

`func NewDomainUserNestview() *DomainUserNestview`

NewDomainUserNestview instantiates a new DomainUserNestview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainUserNestviewWithDefaults

`func NewDomainUserNestviewWithDefaults() *DomainUserNestview`

NewDomainUserNestviewWithDefaults instantiates a new DomainUserNestview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAd

`func (o *DomainUserNestview) GetAd() DomainUserGroupNestviewAd`

GetAd returns the Ad field if non-nil, zero value otherwise.

### GetAdOk

`func (o *DomainUserNestview) GetAdOk() (*DomainUserGroupNestviewAd, bool)`

GetAdOk returns a tuple with the Ad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAd

`func (o *DomainUserNestview) SetAd(v DomainUserGroupNestviewAd)`

SetAd sets Ad field to given value.

### HasAd

`func (o *DomainUserNestview) HasAd() bool`

HasAd returns a boolean if a field has been set.

### GetDescription

`func (o *DomainUserNestview) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DomainUserNestview) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DomainUserNestview) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DomainUserNestview) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DomainUserNestview) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainUserNestview) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainUserNestview) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DomainUserNestview) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLdap

`func (o *DomainUserNestview) GetLdap() DomainUserGroupNestviewLdap`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *DomainUserNestview) GetLdapOk() (*DomainUserGroupNestviewLdap, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *DomainUserNestview) SetLdap(v DomainUserGroupNestviewLdap)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *DomainUserNestview) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetName

`func (o *DomainUserNestview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainUserNestview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainUserNestview) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainUserNestview) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


