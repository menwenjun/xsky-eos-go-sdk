# DomainUserValidatorCreateReqValidator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayGroupId** | Pointer to **int64** | Gateway Group ID | [optional] 
**Groups** | Pointer to **[]string** | Domain groups | [optional] 
**Type** | **string** | Domain type | 
**Users** | Pointer to **[]string** | Domain users | [optional] 

## Methods

### NewDomainUserValidatorCreateReqValidator

`func NewDomainUserValidatorCreateReqValidator(type_ string, ) *DomainUserValidatorCreateReqValidator`

NewDomainUserValidatorCreateReqValidator instantiates a new DomainUserValidatorCreateReqValidator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainUserValidatorCreateReqValidatorWithDefaults

`func NewDomainUserValidatorCreateReqValidatorWithDefaults() *DomainUserValidatorCreateReqValidator`

NewDomainUserValidatorCreateReqValidatorWithDefaults instantiates a new DomainUserValidatorCreateReqValidator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayGroupId

`func (o *DomainUserValidatorCreateReqValidator) GetGatewayGroupId() int64`

GetGatewayGroupId returns the GatewayGroupId field if non-nil, zero value otherwise.

### GetGatewayGroupIdOk

`func (o *DomainUserValidatorCreateReqValidator) GetGatewayGroupIdOk() (*int64, bool)`

GetGatewayGroupIdOk returns a tuple with the GatewayGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayGroupId

`func (o *DomainUserValidatorCreateReqValidator) SetGatewayGroupId(v int64)`

SetGatewayGroupId sets GatewayGroupId field to given value.

### HasGatewayGroupId

`func (o *DomainUserValidatorCreateReqValidator) HasGatewayGroupId() bool`

HasGatewayGroupId returns a boolean if a field has been set.

### GetGroups

`func (o *DomainUserValidatorCreateReqValidator) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *DomainUserValidatorCreateReqValidator) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *DomainUserValidatorCreateReqValidator) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *DomainUserValidatorCreateReqValidator) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetType

`func (o *DomainUserValidatorCreateReqValidator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DomainUserValidatorCreateReqValidator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DomainUserValidatorCreateReqValidator) SetType(v string)`

SetType sets Type field to given value.


### GetUsers

`func (o *DomainUserValidatorCreateReqValidator) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *DomainUserValidatorCreateReqValidator) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *DomainUserValidatorCreateReqValidator) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *DomainUserValidatorCreateReqValidator) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


