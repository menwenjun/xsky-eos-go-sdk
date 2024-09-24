# MetadataClusterCreateReqMetadataCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataServices** | Pointer to [**[]MetadataClusterCreateReqMetadataClusterMetadataServicesElt**](MetadataClusterCreateReqMetadataClusterMetadataServicesElt.md) | MetadataServices is the disk info of metadata cluster | [optional] 
**MinDiskInDc** | Pointer to **int64** | MinDiskInDC is the minimum disk number in data center | [optional] 
**Name** | Pointer to **string** | Name is the name of metadata cluster | [optional] 
**PrimaryPlacementNodeId** | Pointer to **int64** | PrimaryPlacementNodeID is the primary placement node id of metadata cluster | [optional] 

## Methods

### NewMetadataClusterCreateReqMetadataCluster

`func NewMetadataClusterCreateReqMetadataCluster() *MetadataClusterCreateReqMetadataCluster`

NewMetadataClusterCreateReqMetadataCluster instantiates a new MetadataClusterCreateReqMetadataCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataClusterCreateReqMetadataClusterWithDefaults

`func NewMetadataClusterCreateReqMetadataClusterWithDefaults() *MetadataClusterCreateReqMetadataCluster`

NewMetadataClusterCreateReqMetadataClusterWithDefaults instantiates a new MetadataClusterCreateReqMetadataCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataServices

`func (o *MetadataClusterCreateReqMetadataCluster) GetMetadataServices() []MetadataClusterCreateReqMetadataClusterMetadataServicesElt`

GetMetadataServices returns the MetadataServices field if non-nil, zero value otherwise.

### GetMetadataServicesOk

`func (o *MetadataClusterCreateReqMetadataCluster) GetMetadataServicesOk() (*[]MetadataClusterCreateReqMetadataClusterMetadataServicesElt, bool)`

GetMetadataServicesOk returns a tuple with the MetadataServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataServices

`func (o *MetadataClusterCreateReqMetadataCluster) SetMetadataServices(v []MetadataClusterCreateReqMetadataClusterMetadataServicesElt)`

SetMetadataServices sets MetadataServices field to given value.

### HasMetadataServices

`func (o *MetadataClusterCreateReqMetadataCluster) HasMetadataServices() bool`

HasMetadataServices returns a boolean if a field has been set.

### GetMinDiskInDc

`func (o *MetadataClusterCreateReqMetadataCluster) GetMinDiskInDc() int64`

GetMinDiskInDc returns the MinDiskInDc field if non-nil, zero value otherwise.

### GetMinDiskInDcOk

`func (o *MetadataClusterCreateReqMetadataCluster) GetMinDiskInDcOk() (*int64, bool)`

GetMinDiskInDcOk returns a tuple with the MinDiskInDc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDiskInDc

`func (o *MetadataClusterCreateReqMetadataCluster) SetMinDiskInDc(v int64)`

SetMinDiskInDc sets MinDiskInDc field to given value.

### HasMinDiskInDc

`func (o *MetadataClusterCreateReqMetadataCluster) HasMinDiskInDc() bool`

HasMinDiskInDc returns a boolean if a field has been set.

### GetName

`func (o *MetadataClusterCreateReqMetadataCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataClusterCreateReqMetadataCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataClusterCreateReqMetadataCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetadataClusterCreateReqMetadataCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimaryPlacementNodeId

`func (o *MetadataClusterCreateReqMetadataCluster) GetPrimaryPlacementNodeId() int64`

GetPrimaryPlacementNodeId returns the PrimaryPlacementNodeId field if non-nil, zero value otherwise.

### GetPrimaryPlacementNodeIdOk

`func (o *MetadataClusterCreateReqMetadataCluster) GetPrimaryPlacementNodeIdOk() (*int64, bool)`

GetPrimaryPlacementNodeIdOk returns a tuple with the PrimaryPlacementNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPlacementNodeId

`func (o *MetadataClusterCreateReqMetadataCluster) SetPrimaryPlacementNodeId(v int64)`

SetPrimaryPlacementNodeId sets PrimaryPlacementNodeId field to given value.

### HasPrimaryPlacementNodeId

`func (o *MetadataClusterCreateReqMetadataCluster) HasPrimaryPlacementNodeId() bool`

HasPrimaryPlacementNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


