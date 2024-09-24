# DfsSMBShareUpdateACLReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description of user or user group | [optional] 
**Id** | Pointer to **int64** | id of user group | [optional] 
**Permission** | Pointer to **string** | readonly or readwrite access | [optional] 

## Methods

### NewDfsSMBShareUpdateACLReq

`func NewDfsSMBShareUpdateACLReq() *DfsSMBShareUpdateACLReq`

NewDfsSMBShareUpdateACLReq instantiates a new DfsSMBShareUpdateACLReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBShareUpdateACLReqWithDefaults

`func NewDfsSMBShareUpdateACLReqWithDefaults() *DfsSMBShareUpdateACLReq`

NewDfsSMBShareUpdateACLReqWithDefaults instantiates a new DfsSMBShareUpdateACLReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DfsSMBShareUpdateACLReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsSMBShareUpdateACLReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsSMBShareUpdateACLReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsSMBShareUpdateACLReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DfsSMBShareUpdateACLReq) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsSMBShareUpdateACLReq) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsSMBShareUpdateACLReq) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsSMBShareUpdateACLReq) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPermission

`func (o *DfsSMBShareUpdateACLReq) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *DfsSMBShareUpdateACLReq) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *DfsSMBShareUpdateACLReq) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *DfsSMBShareUpdateACLReq) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


