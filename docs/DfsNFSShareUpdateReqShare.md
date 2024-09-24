# DfsNFSShareUpdateReqShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description of share | [optional] 
**DfsGatewayGroupId** | Pointer to **int64** | new dfs gateway group id | [optional] 
**Name** | Pointer to **string** | name of share | [optional] 
**NfsVersions** | Pointer to **[]string** | nfs versions of nfs supported | [optional] 

## Methods

### NewDfsNFSShareUpdateReqShare

`func NewDfsNFSShareUpdateReqShare() *DfsNFSShareUpdateReqShare`

NewDfsNFSShareUpdateReqShare instantiates a new DfsNFSShareUpdateReqShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsNFSShareUpdateReqShareWithDefaults

`func NewDfsNFSShareUpdateReqShareWithDefaults() *DfsNFSShareUpdateReqShare`

NewDfsNFSShareUpdateReqShareWithDefaults instantiates a new DfsNFSShareUpdateReqShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DfsNFSShareUpdateReqShare) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsNFSShareUpdateReqShare) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsNFSShareUpdateReqShare) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsNFSShareUpdateReqShare) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsGatewayGroupId

`func (o *DfsNFSShareUpdateReqShare) GetDfsGatewayGroupId() int64`

GetDfsGatewayGroupId returns the DfsGatewayGroupId field if non-nil, zero value otherwise.

### GetDfsGatewayGroupIdOk

`func (o *DfsNFSShareUpdateReqShare) GetDfsGatewayGroupIdOk() (*int64, bool)`

GetDfsGatewayGroupIdOk returns a tuple with the DfsGatewayGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroupId

`func (o *DfsNFSShareUpdateReqShare) SetDfsGatewayGroupId(v int64)`

SetDfsGatewayGroupId sets DfsGatewayGroupId field to given value.

### HasDfsGatewayGroupId

`func (o *DfsNFSShareUpdateReqShare) HasDfsGatewayGroupId() bool`

HasDfsGatewayGroupId returns a boolean if a field has been set.

### GetName

`func (o *DfsNFSShareUpdateReqShare) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsNFSShareUpdateReqShare) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsNFSShareUpdateReqShare) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsNFSShareUpdateReqShare) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNfsVersions

`func (o *DfsNFSShareUpdateReqShare) GetNfsVersions() []string`

GetNfsVersions returns the NfsVersions field if non-nil, zero value otherwise.

### GetNfsVersionsOk

`func (o *DfsNFSShareUpdateReqShare) GetNfsVersionsOk() (*[]string, bool)`

GetNfsVersionsOk returns a tuple with the NfsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsVersions

`func (o *DfsNFSShareUpdateReqShare) SetNfsVersions(v []string)`

SetNfsVersions sets NfsVersions field to given value.

### HasNfsVersions

`func (o *DfsNFSShareUpdateReqShare) HasNfsVersions() bool`

HasNfsVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


