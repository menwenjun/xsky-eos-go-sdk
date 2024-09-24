# VMDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Disk** | Pointer to [**DiskNestview**](DiskNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Vm** | Pointer to [**VirtualMachineNestview**](VirtualMachineNestview.md) |  | [optional] 
**Volume** | Pointer to [**VolumeNestview**](VolumeNestview.md) |  | [optional] 

## Methods

### NewVMDisk

`func NewVMDisk() *VMDisk`

NewVMDisk instantiates a new VMDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMDiskWithDefaults

`func NewVMDiskWithDefaults() *VMDisk`

NewVMDiskWithDefaults instantiates a new VMDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *VMDisk) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VMDisk) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VMDisk) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VMDisk) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDisk

`func (o *VMDisk) GetDisk() DiskNestview`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *VMDisk) GetDiskOk() (*DiskNestview, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *VMDisk) SetDisk(v DiskNestview)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *VMDisk) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetId

`func (o *VMDisk) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMDisk) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMDisk) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VMDisk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *VMDisk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VMDisk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VMDisk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VMDisk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVm

`func (o *VMDisk) GetVm() VirtualMachineNestview`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *VMDisk) GetVmOk() (*VirtualMachineNestview, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *VMDisk) SetVm(v VirtualMachineNestview)`

SetVm sets Vm field to given value.

### HasVm

`func (o *VMDisk) HasVm() bool`

HasVm returns a boolean if a field has been set.

### GetVolume

`func (o *VMDisk) GetVolume() VolumeNestview`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *VMDisk) GetVolumeOk() (*VolumeNestview, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *VMDisk) SetVolume(v VolumeNestview)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *VMDisk) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


