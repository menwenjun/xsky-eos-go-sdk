# DfsNFSShareCreateReqShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description of share | [optional] 
**DfsGatewayGroupId** | Pointer to **int64** | gateway group id | [optional] 
**DfsNfsShareAcls** | Pointer to [**[]DfsNFSShareACLReq**](DfsNFSShareACLReq.md) | access control array | [optional] 
**DfsRootfsId** | **int64** | rootfs id | 
**Name** | **string** | name of share | 
**NfsVersions** | Pointer to **[]string** | nfs versions of nfs supported | [optional] 
**Path** | **string** | directory path | 

## Methods

### NewDfsNFSShareCreateReqShare

`func NewDfsNFSShareCreateReqShare(dfsRootfsId int64, name string, path string, ) *DfsNFSShareCreateReqShare`

NewDfsNFSShareCreateReqShare instantiates a new DfsNFSShareCreateReqShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsNFSShareCreateReqShareWithDefaults

`func NewDfsNFSShareCreateReqShareWithDefaults() *DfsNFSShareCreateReqShare`

NewDfsNFSShareCreateReqShareWithDefaults instantiates a new DfsNFSShareCreateReqShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DfsNFSShareCreateReqShare) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsNFSShareCreateReqShare) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsNFSShareCreateReqShare) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsNFSShareCreateReqShare) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsGatewayGroupId

`func (o *DfsNFSShareCreateReqShare) GetDfsGatewayGroupId() int64`

GetDfsGatewayGroupId returns the DfsGatewayGroupId field if non-nil, zero value otherwise.

### GetDfsGatewayGroupIdOk

`func (o *DfsNFSShareCreateReqShare) GetDfsGatewayGroupIdOk() (*int64, bool)`

GetDfsGatewayGroupIdOk returns a tuple with the DfsGatewayGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroupId

`func (o *DfsNFSShareCreateReqShare) SetDfsGatewayGroupId(v int64)`

SetDfsGatewayGroupId sets DfsGatewayGroupId field to given value.

### HasDfsGatewayGroupId

`func (o *DfsNFSShareCreateReqShare) HasDfsGatewayGroupId() bool`

HasDfsGatewayGroupId returns a boolean if a field has been set.

### GetDfsNfsShareAcls

`func (o *DfsNFSShareCreateReqShare) GetDfsNfsShareAcls() []DfsNFSShareACLReq`

GetDfsNfsShareAcls returns the DfsNfsShareAcls field if non-nil, zero value otherwise.

### GetDfsNfsShareAclsOk

`func (o *DfsNFSShareCreateReqShare) GetDfsNfsShareAclsOk() (*[]DfsNFSShareACLReq, bool)`

GetDfsNfsShareAclsOk returns a tuple with the DfsNfsShareAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsNfsShareAcls

`func (o *DfsNFSShareCreateReqShare) SetDfsNfsShareAcls(v []DfsNFSShareACLReq)`

SetDfsNfsShareAcls sets DfsNfsShareAcls field to given value.

### HasDfsNfsShareAcls

`func (o *DfsNFSShareCreateReqShare) HasDfsNfsShareAcls() bool`

HasDfsNfsShareAcls returns a boolean if a field has been set.

### GetDfsRootfsId

`func (o *DfsNFSShareCreateReqShare) GetDfsRootfsId() int64`

GetDfsRootfsId returns the DfsRootfsId field if non-nil, zero value otherwise.

### GetDfsRootfsIdOk

`func (o *DfsNFSShareCreateReqShare) GetDfsRootfsIdOk() (*int64, bool)`

GetDfsRootfsIdOk returns a tuple with the DfsRootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfsId

`func (o *DfsNFSShareCreateReqShare) SetDfsRootfsId(v int64)`

SetDfsRootfsId sets DfsRootfsId field to given value.


### GetName

`func (o *DfsNFSShareCreateReqShare) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsNFSShareCreateReqShare) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsNFSShareCreateReqShare) SetName(v string)`

SetName sets Name field to given value.


### GetNfsVersions

`func (o *DfsNFSShareCreateReqShare) GetNfsVersions() []string`

GetNfsVersions returns the NfsVersions field if non-nil, zero value otherwise.

### GetNfsVersionsOk

`func (o *DfsNFSShareCreateReqShare) GetNfsVersionsOk() (*[]string, bool)`

GetNfsVersionsOk returns a tuple with the NfsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsVersions

`func (o *DfsNFSShareCreateReqShare) SetNfsVersions(v []string)`

SetNfsVersions sets NfsVersions field to given value.

### HasNfsVersions

`func (o *DfsNFSShareCreateReqShare) HasNfsVersions() bool`

HasNfsVersions returns a boolean if a field has been set.

### GetPath

`func (o *DfsNFSShareCreateReqShare) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsNFSShareCreateReqShare) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsNFSShareCreateReqShare) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


