# VirtualMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**CustomCpu** | Pointer to **int64** |  | [optional] 
**Disks** | Pointer to [**[]VMDisk**](VMDisk.md) |  | [optional] 
**Flavor** | Pointer to [**VMFlavorNestview**](VMFlavorNestview.md) |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Image** | Pointer to [**VMImageNestview**](VMImageNestview.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Nics** | Pointer to [**[]VMNetworkInterface**](VMNetworkInterface.md) |  | [optional] 
**RootVolume** | Pointer to [**VolumeNestview**](VolumeNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**VmStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewVirtualMachine

`func NewVirtualMachine() *VirtualMachine`

NewVirtualMachine instantiates a new VirtualMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMachineWithDefaults

`func NewVirtualMachineWithDefaults() *VirtualMachine`

NewVirtualMachineWithDefaults instantiates a new VirtualMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *VirtualMachine) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *VirtualMachine) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *VirtualMachine) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *VirtualMachine) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualMachine) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualMachine) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualMachine) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualMachine) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *VirtualMachine) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *VirtualMachine) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *VirtualMachine) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *VirtualMachine) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCustomCpu

`func (o *VirtualMachine) GetCustomCpu() int64`

GetCustomCpu returns the CustomCpu field if non-nil, zero value otherwise.

### GetCustomCpuOk

`func (o *VirtualMachine) GetCustomCpuOk() (*int64, bool)`

GetCustomCpuOk returns a tuple with the CustomCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCpu

`func (o *VirtualMachine) SetCustomCpu(v int64)`

SetCustomCpu sets CustomCpu field to given value.

### HasCustomCpu

`func (o *VirtualMachine) HasCustomCpu() bool`

HasCustomCpu returns a boolean if a field has been set.

### GetDisks

`func (o *VirtualMachine) GetDisks() []VMDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VirtualMachine) GetDisksOk() (*[]VMDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VirtualMachine) SetDisks(v []VMDisk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *VirtualMachine) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetFlavor

`func (o *VirtualMachine) GetFlavor() VMFlavorNestview`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *VirtualMachine) GetFlavorOk() (*VMFlavorNestview, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *VirtualMachine) SetFlavor(v VMFlavorNestview)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *VirtualMachine) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetHost

`func (o *VirtualMachine) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualMachine) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualMachine) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualMachine) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *VirtualMachine) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualMachine) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualMachine) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualMachine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *VirtualMachine) GetImage() VMImageNestview`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *VirtualMachine) GetImageOk() (*VMImageNestview, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *VirtualMachine) SetImage(v VMImageNestview)`

SetImage sets Image field to given value.

### HasImage

`func (o *VirtualMachine) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *VirtualMachine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualMachine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualMachine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualMachine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNics

`func (o *VirtualMachine) GetNics() []VMNetworkInterface`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *VirtualMachine) GetNicsOk() (*[]VMNetworkInterface, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *VirtualMachine) SetNics(v []VMNetworkInterface)`

SetNics sets Nics field to given value.

### HasNics

`func (o *VirtualMachine) HasNics() bool`

HasNics returns a boolean if a field has been set.

### GetRootVolume

`func (o *VirtualMachine) GetRootVolume() VolumeNestview`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *VirtualMachine) GetRootVolumeOk() (*VolumeNestview, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *VirtualMachine) SetRootVolume(v VolumeNestview)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *VirtualMachine) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualMachine) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualMachine) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualMachine) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualMachine) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *VirtualMachine) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *VirtualMachine) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *VirtualMachine) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *VirtualMachine) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUuid

`func (o *VirtualMachine) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualMachine) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualMachine) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualMachine) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVmStatus

`func (o *VirtualMachine) GetVmStatus() string`

GetVmStatus returns the VmStatus field if non-nil, zero value otherwise.

### GetVmStatusOk

`func (o *VirtualMachine) GetVmStatusOk() (*string, bool)`

GetVmStatusOk returns a tuple with the VmStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmStatus

`func (o *VirtualMachine) SetVmStatus(v string)`

SetVmStatus sets VmStatus field to given value.

### HasVmStatus

`func (o *VirtualMachine) HasVmStatus() bool`

HasVmStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


