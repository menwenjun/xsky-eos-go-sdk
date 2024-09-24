# VMNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**Mask** | Pointer to **int64** |  | [optional] 
**Master** | Pointer to [**VMNetworkInterfaceNestview**](VMNetworkInterfaceNestview.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NatConfig** | Pointer to [**[]PortMapPair**](PortMapPair.md) |  | [optional] 
**NetworkType** | Pointer to **string** |  | [optional] 
**PciAddress** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Vm** | Pointer to [**VirtualMachineNestview**](VirtualMachineNestview.md) |  | [optional] 

## Methods

### NewVMNetworkInterface

`func NewVMNetworkInterface() *VMNetworkInterface`

NewVMNetworkInterface instantiates a new VMNetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMNetworkInterfaceWithDefaults

`func NewVMNetworkInterfaceWithDefaults() *VMNetworkInterface`

NewVMNetworkInterfaceWithDefaults instantiates a new VMNetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *VMNetworkInterface) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *VMNetworkInterface) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *VMNetworkInterface) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *VMNetworkInterface) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetGateway

`func (o *VMNetworkInterface) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *VMNetworkInterface) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *VMNetworkInterface) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *VMNetworkInterface) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetId

`func (o *VMNetworkInterface) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMNetworkInterface) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMNetworkInterface) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VMNetworkInterface) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *VMNetworkInterface) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VMNetworkInterface) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VMNetworkInterface) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VMNetworkInterface) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *VMNetworkInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VMNetworkInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VMNetworkInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VMNetworkInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMask

`func (o *VMNetworkInterface) GetMask() int64`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *VMNetworkInterface) GetMaskOk() (*int64, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *VMNetworkInterface) SetMask(v int64)`

SetMask sets Mask field to given value.

### HasMask

`func (o *VMNetworkInterface) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetMaster

`func (o *VMNetworkInterface) GetMaster() VMNetworkInterfaceNestview`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *VMNetworkInterface) GetMasterOk() (*VMNetworkInterfaceNestview, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *VMNetworkInterface) SetMaster(v VMNetworkInterfaceNestview)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *VMNetworkInterface) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetName

`func (o *VMNetworkInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VMNetworkInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VMNetworkInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VMNetworkInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNatConfig

`func (o *VMNetworkInterface) GetNatConfig() []PortMapPair`

GetNatConfig returns the NatConfig field if non-nil, zero value otherwise.

### GetNatConfigOk

`func (o *VMNetworkInterface) GetNatConfigOk() (*[]PortMapPair, bool)`

GetNatConfigOk returns a tuple with the NatConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatConfig

`func (o *VMNetworkInterface) SetNatConfig(v []PortMapPair)`

SetNatConfig sets NatConfig field to given value.

### HasNatConfig

`func (o *VMNetworkInterface) HasNatConfig() bool`

HasNatConfig returns a boolean if a field has been set.

### GetNetworkType

`func (o *VMNetworkInterface) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *VMNetworkInterface) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *VMNetworkInterface) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *VMNetworkInterface) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetPciAddress

`func (o *VMNetworkInterface) GetPciAddress() string`

GetPciAddress returns the PciAddress field if non-nil, zero value otherwise.

### GetPciAddressOk

`func (o *VMNetworkInterface) GetPciAddressOk() (*string, bool)`

GetPciAddressOk returns a tuple with the PciAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddress

`func (o *VMNetworkInterface) SetPciAddress(v string)`

SetPciAddress sets PciAddress field to given value.

### HasPciAddress

`func (o *VMNetworkInterface) HasPciAddress() bool`

HasPciAddress returns a boolean if a field has been set.

### GetType

`func (o *VMNetworkInterface) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VMNetworkInterface) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VMNetworkInterface) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VMNetworkInterface) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *VMNetworkInterface) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *VMNetworkInterface) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *VMNetworkInterface) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *VMNetworkInterface) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVm

`func (o *VMNetworkInterface) GetVm() VirtualMachineNestview`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *VMNetworkInterface) GetVmOk() (*VirtualMachineNestview, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *VMNetworkInterface) SetVm(v VirtualMachineNestview)`

SetVm sets Vm field to given value.

### HasVm

`func (o *VMNetworkInterface) HasVm() bool`

HasVm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


