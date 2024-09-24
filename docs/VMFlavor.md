# VMFlavor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **int64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MemoryKbyte** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RootDiskSize** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewVMFlavor

`func NewVMFlavor() *VMFlavor`

NewVMFlavor instantiates a new VMFlavor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMFlavorWithDefaults

`func NewVMFlavorWithDefaults() *VMFlavor`

NewVMFlavorWithDefaults instantiates a new VMFlavor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *VMFlavor) GetCpu() int64`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *VMFlavor) GetCpuOk() (*int64, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *VMFlavor) SetCpu(v int64)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *VMFlavor) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetCreate

`func (o *VMFlavor) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *VMFlavor) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *VMFlavor) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *VMFlavor) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *VMFlavor) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMFlavor) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMFlavor) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VMFlavor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMemoryKbyte

`func (o *VMFlavor) GetMemoryKbyte() int64`

GetMemoryKbyte returns the MemoryKbyte field if non-nil, zero value otherwise.

### GetMemoryKbyteOk

`func (o *VMFlavor) GetMemoryKbyteOk() (*int64, bool)`

GetMemoryKbyteOk returns a tuple with the MemoryKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryKbyte

`func (o *VMFlavor) SetMemoryKbyte(v int64)`

SetMemoryKbyte sets MemoryKbyte field to given value.

### HasMemoryKbyte

`func (o *VMFlavor) HasMemoryKbyte() bool`

HasMemoryKbyte returns a boolean if a field has been set.

### GetName

`func (o *VMFlavor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VMFlavor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VMFlavor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VMFlavor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRootDiskSize

`func (o *VMFlavor) GetRootDiskSize() int64`

GetRootDiskSize returns the RootDiskSize field if non-nil, zero value otherwise.

### GetRootDiskSizeOk

`func (o *VMFlavor) GetRootDiskSizeOk() (*int64, bool)`

GetRootDiskSizeOk returns a tuple with the RootDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSize

`func (o *VMFlavor) SetRootDiskSize(v int64)`

SetRootDiskSize sets RootDiskSize field to given value.

### HasRootDiskSize

`func (o *VMFlavor) HasRootDiskSize() bool`

HasRootDiskSize returns a boolean if a field has been set.

### GetUpdate

`func (o *VMFlavor) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *VMFlavor) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *VMFlavor) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *VMFlavor) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


