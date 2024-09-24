# DfsQosRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsPath** | Pointer to [**DfsPath**](DfsPath.md) |  | [optional] 
**DfsQosPolicy** | Pointer to [**DfsQosPolicy**](DfsQosPolicy.md) |  | [optional] 
**DfsRootfs** | Pointer to [**DfsRootfs**](DfsRootfs.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IoStatus** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**HdfsNum** | Pointer to **int64** |  | [optional] 
**IsBucket** | Pointer to **bool** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Samples** | Pointer to [**[]DfsQosStat**](DfsQosStat.md) |  | [optional] 
**Shares** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDfsQosRecord

`func NewDfsQosRecord() *DfsQosRecord`

NewDfsQosRecord instantiates a new DfsQosRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsQosRecordWithDefaults

`func NewDfsQosRecordWithDefaults() *DfsQosRecord`

NewDfsQosRecordWithDefaults instantiates a new DfsQosRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DfsQosRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsQosRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsQosRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsQosRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsQosRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsQosRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsQosRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsQosRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsPath

`func (o *DfsQosRecord) GetDfsPath() DfsPath`

GetDfsPath returns the DfsPath field if non-nil, zero value otherwise.

### GetDfsPathOk

`func (o *DfsQosRecord) GetDfsPathOk() (*DfsPath, bool)`

GetDfsPathOk returns a tuple with the DfsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPath

`func (o *DfsQosRecord) SetDfsPath(v DfsPath)`

SetDfsPath sets DfsPath field to given value.

### HasDfsPath

`func (o *DfsQosRecord) HasDfsPath() bool`

HasDfsPath returns a boolean if a field has been set.

### GetDfsQosPolicy

`func (o *DfsQosRecord) GetDfsQosPolicy() DfsQosPolicy`

GetDfsQosPolicy returns the DfsQosPolicy field if non-nil, zero value otherwise.

### GetDfsQosPolicyOk

`func (o *DfsQosRecord) GetDfsQosPolicyOk() (*DfsQosPolicy, bool)`

GetDfsQosPolicyOk returns a tuple with the DfsQosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsQosPolicy

`func (o *DfsQosRecord) SetDfsQosPolicy(v DfsQosPolicy)`

SetDfsQosPolicy sets DfsQosPolicy field to given value.

### HasDfsQosPolicy

`func (o *DfsQosRecord) HasDfsQosPolicy() bool`

HasDfsQosPolicy returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *DfsQosRecord) GetDfsRootfs() DfsRootfs`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsQosRecord) GetDfsRootfsOk() (*DfsRootfs, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsQosRecord) SetDfsRootfs(v DfsRootfs)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *DfsQosRecord) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetId

`func (o *DfsQosRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsQosRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsQosRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsQosRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIoStatus

`func (o *DfsQosRecord) GetIoStatus() string`

GetIoStatus returns the IoStatus field if non-nil, zero value otherwise.

### GetIoStatusOk

`func (o *DfsQosRecord) GetIoStatusOk() (*string, bool)`

GetIoStatusOk returns a tuple with the IoStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoStatus

`func (o *DfsQosRecord) SetIoStatus(v string)`

SetIoStatus sets IoStatus field to given value.

### HasIoStatus

`func (o *DfsQosRecord) HasIoStatus() bool`

HasIoStatus returns a boolean if a field has been set.

### GetStatus

`func (o *DfsQosRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsQosRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsQosRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsQosRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsQosRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsQosRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsQosRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsQosRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetHdfsNum

`func (o *DfsQosRecord) GetHdfsNum() int64`

GetHdfsNum returns the HdfsNum field if non-nil, zero value otherwise.

### GetHdfsNumOk

`func (o *DfsQosRecord) GetHdfsNumOk() (*int64, bool)`

GetHdfsNumOk returns a tuple with the HdfsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsNum

`func (o *DfsQosRecord) SetHdfsNum(v int64)`

SetHdfsNum sets HdfsNum field to given value.

### HasHdfsNum

`func (o *DfsQosRecord) HasHdfsNum() bool`

HasHdfsNum returns a boolean if a field has been set.

### GetIsBucket

`func (o *DfsQosRecord) GetIsBucket() bool`

GetIsBucket returns the IsBucket field if non-nil, zero value otherwise.

### GetIsBucketOk

`func (o *DfsQosRecord) GetIsBucketOk() (*bool, bool)`

GetIsBucketOk returns a tuple with the IsBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBucket

`func (o *DfsQosRecord) SetIsBucket(v bool)`

SetIsBucket sets IsBucket field to given value.

### HasIsBucket

`func (o *DfsQosRecord) HasIsBucket() bool`

HasIsBucket returns a boolean if a field has been set.

### GetPath

`func (o *DfsQosRecord) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsQosRecord) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsQosRecord) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DfsQosRecord) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSamples

`func (o *DfsQosRecord) GetSamples() []DfsQosStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *DfsQosRecord) GetSamplesOk() (*[]DfsQosStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *DfsQosRecord) SetSamples(v []DfsQosStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *DfsQosRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetShares

`func (o *DfsQosRecord) GetShares() []string`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *DfsQosRecord) GetSharesOk() (*[]string, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *DfsQosRecord) SetShares(v []string)`

SetShares sets Shares field to given value.

### HasShares

`func (o *DfsQosRecord) HasShares() bool`

HasShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


