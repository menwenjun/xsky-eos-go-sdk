# CloudInstanceRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudDatacenter** | Pointer to [**CloudDatacenterNestview**](CloudDatacenterNestview.md) |  | [optional] 
**CloudInstanceId** | Pointer to **string** |  | [optional] 
**CloudPlatform** | Pointer to [**CloudPlatformNestview**](CloudPlatformNestview.md) |  | [optional] 
**CloudVolumeNum** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Cores** | Pointer to **int64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MemoryKbyte** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RootDeviceType** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Samples** | Pointer to [**[]CloudInstanceStat**](CloudInstanceStat.md) |  | [optional] 

## Methods

### NewCloudInstanceRecord

`func NewCloudInstanceRecord() *CloudInstanceRecord`

NewCloudInstanceRecord instantiates a new CloudInstanceRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInstanceRecordWithDefaults

`func NewCloudInstanceRecordWithDefaults() *CloudInstanceRecord`

NewCloudInstanceRecordWithDefaults instantiates a new CloudInstanceRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudDatacenter

`func (o *CloudInstanceRecord) GetCloudDatacenter() CloudDatacenterNestview`

GetCloudDatacenter returns the CloudDatacenter field if non-nil, zero value otherwise.

### GetCloudDatacenterOk

`func (o *CloudInstanceRecord) GetCloudDatacenterOk() (*CloudDatacenterNestview, bool)`

GetCloudDatacenterOk returns a tuple with the CloudDatacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDatacenter

`func (o *CloudInstanceRecord) SetCloudDatacenter(v CloudDatacenterNestview)`

SetCloudDatacenter sets CloudDatacenter field to given value.

### HasCloudDatacenter

`func (o *CloudInstanceRecord) HasCloudDatacenter() bool`

HasCloudDatacenter returns a boolean if a field has been set.

### GetCloudInstanceId

`func (o *CloudInstanceRecord) GetCloudInstanceId() string`

GetCloudInstanceId returns the CloudInstanceId field if non-nil, zero value otherwise.

### GetCloudInstanceIdOk

`func (o *CloudInstanceRecord) GetCloudInstanceIdOk() (*string, bool)`

GetCloudInstanceIdOk returns a tuple with the CloudInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInstanceId

`func (o *CloudInstanceRecord) SetCloudInstanceId(v string)`

SetCloudInstanceId sets CloudInstanceId field to given value.

### HasCloudInstanceId

`func (o *CloudInstanceRecord) HasCloudInstanceId() bool`

HasCloudInstanceId returns a boolean if a field has been set.

### GetCloudPlatform

`func (o *CloudInstanceRecord) GetCloudPlatform() CloudPlatformNestview`

GetCloudPlatform returns the CloudPlatform field if non-nil, zero value otherwise.

### GetCloudPlatformOk

`func (o *CloudInstanceRecord) GetCloudPlatformOk() (*CloudPlatformNestview, bool)`

GetCloudPlatformOk returns a tuple with the CloudPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPlatform

`func (o *CloudInstanceRecord) SetCloudPlatform(v CloudPlatformNestview)`

SetCloudPlatform sets CloudPlatform field to given value.

### HasCloudPlatform

`func (o *CloudInstanceRecord) HasCloudPlatform() bool`

HasCloudPlatform returns a boolean if a field has been set.

### GetCloudVolumeNum

`func (o *CloudInstanceRecord) GetCloudVolumeNum() int64`

GetCloudVolumeNum returns the CloudVolumeNum field if non-nil, zero value otherwise.

### GetCloudVolumeNumOk

`func (o *CloudInstanceRecord) GetCloudVolumeNumOk() (*int64, bool)`

GetCloudVolumeNumOk returns a tuple with the CloudVolumeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudVolumeNum

`func (o *CloudInstanceRecord) SetCloudVolumeNum(v int64)`

SetCloudVolumeNum sets CloudVolumeNum field to given value.

### HasCloudVolumeNum

`func (o *CloudInstanceRecord) HasCloudVolumeNum() bool`

HasCloudVolumeNum returns a boolean if a field has been set.

### GetCluster

`func (o *CloudInstanceRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *CloudInstanceRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *CloudInstanceRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *CloudInstanceRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCores

`func (o *CloudInstanceRecord) GetCores() int64`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *CloudInstanceRecord) GetCoresOk() (*int64, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *CloudInstanceRecord) SetCores(v int64)`

SetCores sets Cores field to given value.

### HasCores

`func (o *CloudInstanceRecord) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetCreate

`func (o *CloudInstanceRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *CloudInstanceRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *CloudInstanceRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *CloudInstanceRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetHostname

`func (o *CloudInstanceRecord) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CloudInstanceRecord) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CloudInstanceRecord) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CloudInstanceRecord) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *CloudInstanceRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudInstanceRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudInstanceRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CloudInstanceRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMemoryKbyte

`func (o *CloudInstanceRecord) GetMemoryKbyte() int64`

GetMemoryKbyte returns the MemoryKbyte field if non-nil, zero value otherwise.

### GetMemoryKbyteOk

`func (o *CloudInstanceRecord) GetMemoryKbyteOk() (*int64, bool)`

GetMemoryKbyteOk returns a tuple with the MemoryKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryKbyte

`func (o *CloudInstanceRecord) SetMemoryKbyte(v int64)`

SetMemoryKbyte sets MemoryKbyte field to given value.

### HasMemoryKbyte

`func (o *CloudInstanceRecord) HasMemoryKbyte() bool`

HasMemoryKbyte returns a boolean if a field has been set.

### GetName

`func (o *CloudInstanceRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudInstanceRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudInstanceRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudInstanceRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRootDeviceType

`func (o *CloudInstanceRecord) GetRootDeviceType() string`

GetRootDeviceType returns the RootDeviceType field if non-nil, zero value otherwise.

### GetRootDeviceTypeOk

`func (o *CloudInstanceRecord) GetRootDeviceTypeOk() (*string, bool)`

GetRootDeviceTypeOk returns a tuple with the RootDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDeviceType

`func (o *CloudInstanceRecord) SetRootDeviceType(v string)`

SetRootDeviceType sets RootDeviceType field to given value.

### HasRootDeviceType

`func (o *CloudInstanceRecord) HasRootDeviceType() bool`

HasRootDeviceType returns a boolean if a field has been set.

### GetUpdate

`func (o *CloudInstanceRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *CloudInstanceRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *CloudInstanceRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *CloudInstanceRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSamples

`func (o *CloudInstanceRecord) GetSamples() []CloudInstanceStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *CloudInstanceRecord) GetSamplesOk() (*[]CloudInstanceStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *CloudInstanceRecord) SetSamples(v []CloudInstanceStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *CloudInstanceRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


