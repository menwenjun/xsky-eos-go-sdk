# CloudInstance

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

## Methods

### NewCloudInstance

`func NewCloudInstance() *CloudInstance`

NewCloudInstance instantiates a new CloudInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInstanceWithDefaults

`func NewCloudInstanceWithDefaults() *CloudInstance`

NewCloudInstanceWithDefaults instantiates a new CloudInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudDatacenter

`func (o *CloudInstance) GetCloudDatacenter() CloudDatacenterNestview`

GetCloudDatacenter returns the CloudDatacenter field if non-nil, zero value otherwise.

### GetCloudDatacenterOk

`func (o *CloudInstance) GetCloudDatacenterOk() (*CloudDatacenterNestview, bool)`

GetCloudDatacenterOk returns a tuple with the CloudDatacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDatacenter

`func (o *CloudInstance) SetCloudDatacenter(v CloudDatacenterNestview)`

SetCloudDatacenter sets CloudDatacenter field to given value.

### HasCloudDatacenter

`func (o *CloudInstance) HasCloudDatacenter() bool`

HasCloudDatacenter returns a boolean if a field has been set.

### GetCloudInstanceId

`func (o *CloudInstance) GetCloudInstanceId() string`

GetCloudInstanceId returns the CloudInstanceId field if non-nil, zero value otherwise.

### GetCloudInstanceIdOk

`func (o *CloudInstance) GetCloudInstanceIdOk() (*string, bool)`

GetCloudInstanceIdOk returns a tuple with the CloudInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInstanceId

`func (o *CloudInstance) SetCloudInstanceId(v string)`

SetCloudInstanceId sets CloudInstanceId field to given value.

### HasCloudInstanceId

`func (o *CloudInstance) HasCloudInstanceId() bool`

HasCloudInstanceId returns a boolean if a field has been set.

### GetCloudPlatform

`func (o *CloudInstance) GetCloudPlatform() CloudPlatformNestview`

GetCloudPlatform returns the CloudPlatform field if non-nil, zero value otherwise.

### GetCloudPlatformOk

`func (o *CloudInstance) GetCloudPlatformOk() (*CloudPlatformNestview, bool)`

GetCloudPlatformOk returns a tuple with the CloudPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPlatform

`func (o *CloudInstance) SetCloudPlatform(v CloudPlatformNestview)`

SetCloudPlatform sets CloudPlatform field to given value.

### HasCloudPlatform

`func (o *CloudInstance) HasCloudPlatform() bool`

HasCloudPlatform returns a boolean if a field has been set.

### GetCloudVolumeNum

`func (o *CloudInstance) GetCloudVolumeNum() int64`

GetCloudVolumeNum returns the CloudVolumeNum field if non-nil, zero value otherwise.

### GetCloudVolumeNumOk

`func (o *CloudInstance) GetCloudVolumeNumOk() (*int64, bool)`

GetCloudVolumeNumOk returns a tuple with the CloudVolumeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudVolumeNum

`func (o *CloudInstance) SetCloudVolumeNum(v int64)`

SetCloudVolumeNum sets CloudVolumeNum field to given value.

### HasCloudVolumeNum

`func (o *CloudInstance) HasCloudVolumeNum() bool`

HasCloudVolumeNum returns a boolean if a field has been set.

### GetCluster

`func (o *CloudInstance) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *CloudInstance) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *CloudInstance) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *CloudInstance) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCores

`func (o *CloudInstance) GetCores() int64`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *CloudInstance) GetCoresOk() (*int64, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *CloudInstance) SetCores(v int64)`

SetCores sets Cores field to given value.

### HasCores

`func (o *CloudInstance) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetCreate

`func (o *CloudInstance) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *CloudInstance) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *CloudInstance) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *CloudInstance) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetHostname

`func (o *CloudInstance) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CloudInstance) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CloudInstance) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CloudInstance) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *CloudInstance) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudInstance) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudInstance) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CloudInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMemoryKbyte

`func (o *CloudInstance) GetMemoryKbyte() int64`

GetMemoryKbyte returns the MemoryKbyte field if non-nil, zero value otherwise.

### GetMemoryKbyteOk

`func (o *CloudInstance) GetMemoryKbyteOk() (*int64, bool)`

GetMemoryKbyteOk returns a tuple with the MemoryKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryKbyte

`func (o *CloudInstance) SetMemoryKbyte(v int64)`

SetMemoryKbyte sets MemoryKbyte field to given value.

### HasMemoryKbyte

`func (o *CloudInstance) HasMemoryKbyte() bool`

HasMemoryKbyte returns a boolean if a field has been set.

### GetName

`func (o *CloudInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRootDeviceType

`func (o *CloudInstance) GetRootDeviceType() string`

GetRootDeviceType returns the RootDeviceType field if non-nil, zero value otherwise.

### GetRootDeviceTypeOk

`func (o *CloudInstance) GetRootDeviceTypeOk() (*string, bool)`

GetRootDeviceTypeOk returns a tuple with the RootDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDeviceType

`func (o *CloudInstance) SetRootDeviceType(v string)`

SetRootDeviceType sets RootDeviceType field to given value.

### HasRootDeviceType

`func (o *CloudInstance) HasRootDeviceType() bool`

HasRootDeviceType returns a boolean if a field has been set.

### GetUpdate

`func (o *CloudInstance) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *CloudInstance) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *CloudInstance) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *CloudInstance) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


