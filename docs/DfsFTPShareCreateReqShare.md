# DfsFTPShareCreateReqShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description of share | [optional] 
**DfsFtpShareAcls** | Pointer to [**[]DfsFTPShareACLReq**](DfsFTPShareACLReq.md) | access control array | [optional] 
**DfsGatewayGroupId** | Pointer to **int64** | gateway group id | [optional] 
**DfsRootfsId** | **int64** | rootfs id | 
**Name** | **string** | name of share | 
**Path** | **string** | diectrory path | 

## Methods

### NewDfsFTPShareCreateReqShare

`func NewDfsFTPShareCreateReqShare(dfsRootfsId int64, name string, path string, ) *DfsFTPShareCreateReqShare`

NewDfsFTPShareCreateReqShare instantiates a new DfsFTPShareCreateReqShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsFTPShareCreateReqShareWithDefaults

`func NewDfsFTPShareCreateReqShareWithDefaults() *DfsFTPShareCreateReqShare`

NewDfsFTPShareCreateReqShareWithDefaults instantiates a new DfsFTPShareCreateReqShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DfsFTPShareCreateReqShare) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsFTPShareCreateReqShare) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsFTPShareCreateReqShare) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsFTPShareCreateReqShare) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsFtpShareAcls

`func (o *DfsFTPShareCreateReqShare) GetDfsFtpShareAcls() []DfsFTPShareACLReq`

GetDfsFtpShareAcls returns the DfsFtpShareAcls field if non-nil, zero value otherwise.

### GetDfsFtpShareAclsOk

`func (o *DfsFTPShareCreateReqShare) GetDfsFtpShareAclsOk() (*[]DfsFTPShareACLReq, bool)`

GetDfsFtpShareAclsOk returns a tuple with the DfsFtpShareAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsFtpShareAcls

`func (o *DfsFTPShareCreateReqShare) SetDfsFtpShareAcls(v []DfsFTPShareACLReq)`

SetDfsFtpShareAcls sets DfsFtpShareAcls field to given value.

### HasDfsFtpShareAcls

`func (o *DfsFTPShareCreateReqShare) HasDfsFtpShareAcls() bool`

HasDfsFtpShareAcls returns a boolean if a field has been set.

### GetDfsGatewayGroupId

`func (o *DfsFTPShareCreateReqShare) GetDfsGatewayGroupId() int64`

GetDfsGatewayGroupId returns the DfsGatewayGroupId field if non-nil, zero value otherwise.

### GetDfsGatewayGroupIdOk

`func (o *DfsFTPShareCreateReqShare) GetDfsGatewayGroupIdOk() (*int64, bool)`

GetDfsGatewayGroupIdOk returns a tuple with the DfsGatewayGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroupId

`func (o *DfsFTPShareCreateReqShare) SetDfsGatewayGroupId(v int64)`

SetDfsGatewayGroupId sets DfsGatewayGroupId field to given value.

### HasDfsGatewayGroupId

`func (o *DfsFTPShareCreateReqShare) HasDfsGatewayGroupId() bool`

HasDfsGatewayGroupId returns a boolean if a field has been set.

### GetDfsRootfsId

`func (o *DfsFTPShareCreateReqShare) GetDfsRootfsId() int64`

GetDfsRootfsId returns the DfsRootfsId field if non-nil, zero value otherwise.

### GetDfsRootfsIdOk

`func (o *DfsFTPShareCreateReqShare) GetDfsRootfsIdOk() (*int64, bool)`

GetDfsRootfsIdOk returns a tuple with the DfsRootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfsId

`func (o *DfsFTPShareCreateReqShare) SetDfsRootfsId(v int64)`

SetDfsRootfsId sets DfsRootfsId field to given value.


### GetName

`func (o *DfsFTPShareCreateReqShare) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsFTPShareCreateReqShare) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsFTPShareCreateReqShare) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *DfsFTPShareCreateReqShare) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsFTPShareCreateReqShare) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsFTPShareCreateReqShare) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


