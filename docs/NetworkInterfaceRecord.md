# NetworkInterfaceRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BondingMode** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LinkDetected** | Pointer to **bool** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**MasterNetworkInterface** | Pointer to [**NetworkInterfaceNestview**](NetworkInterfaceNestview.md) |  | [optional] 
**Megabits** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Operstate** | Pointer to **string** |  | [optional] 
**OspCluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**Type** | Pointer to **string** | ethernet or bond | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**VfNum** | Pointer to **int64** |  | [optional] 
**Samples** | Pointer to [**[]NetworkInterfaceStat**](NetworkInterfaceStat.md) |  | [optional] 

## Methods

### NewNetworkInterfaceRecord

`func NewNetworkInterfaceRecord() *NetworkInterfaceRecord`

NewNetworkInterfaceRecord instantiates a new NetworkInterfaceRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceRecordWithDefaults

`func NewNetworkInterfaceRecordWithDefaults() *NetworkInterfaceRecord`

NewNetworkInterfaceRecordWithDefaults instantiates a new NetworkInterfaceRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBondingMode

`func (o *NetworkInterfaceRecord) GetBondingMode() string`

GetBondingMode returns the BondingMode field if non-nil, zero value otherwise.

### GetBondingModeOk

`func (o *NetworkInterfaceRecord) GetBondingModeOk() (*string, bool)`

GetBondingModeOk returns a tuple with the BondingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondingMode

`func (o *NetworkInterfaceRecord) SetBondingMode(v string)`

SetBondingMode sets BondingMode field to given value.

### HasBondingMode

`func (o *NetworkInterfaceRecord) HasBondingMode() bool`

HasBondingMode returns a boolean if a field has been set.

### GetCluster

`func (o *NetworkInterfaceRecord) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *NetworkInterfaceRecord) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *NetworkInterfaceRecord) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *NetworkInterfaceRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *NetworkInterfaceRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *NetworkInterfaceRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *NetworkInterfaceRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *NetworkInterfaceRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetHost

`func (o *NetworkInterfaceRecord) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NetworkInterfaceRecord) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NetworkInterfaceRecord) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *NetworkInterfaceRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *NetworkInterfaceRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkInterfaceRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkInterfaceRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkInterfaceRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinkDetected

`func (o *NetworkInterfaceRecord) GetLinkDetected() bool`

GetLinkDetected returns the LinkDetected field if non-nil, zero value otherwise.

### GetLinkDetectedOk

`func (o *NetworkInterfaceRecord) GetLinkDetectedOk() (*bool, bool)`

GetLinkDetectedOk returns a tuple with the LinkDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDetected

`func (o *NetworkInterfaceRecord) SetLinkDetected(v bool)`

SetLinkDetected sets LinkDetected field to given value.

### HasLinkDetected

`func (o *NetworkInterfaceRecord) HasLinkDetected() bool`

HasLinkDetected returns a boolean if a field has been set.

### GetMacAddress

`func (o *NetworkInterfaceRecord) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *NetworkInterfaceRecord) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *NetworkInterfaceRecord) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *NetworkInterfaceRecord) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMasterNetworkInterface

`func (o *NetworkInterfaceRecord) GetMasterNetworkInterface() NetworkInterfaceNestview`

GetMasterNetworkInterface returns the MasterNetworkInterface field if non-nil, zero value otherwise.

### GetMasterNetworkInterfaceOk

`func (o *NetworkInterfaceRecord) GetMasterNetworkInterfaceOk() (*NetworkInterfaceNestview, bool)`

GetMasterNetworkInterfaceOk returns a tuple with the MasterNetworkInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterNetworkInterface

`func (o *NetworkInterfaceRecord) SetMasterNetworkInterface(v NetworkInterfaceNestview)`

SetMasterNetworkInterface sets MasterNetworkInterface field to given value.

### HasMasterNetworkInterface

`func (o *NetworkInterfaceRecord) HasMasterNetworkInterface() bool`

HasMasterNetworkInterface returns a boolean if a field has been set.

### GetMegabits

`func (o *NetworkInterfaceRecord) GetMegabits() int64`

GetMegabits returns the Megabits field if non-nil, zero value otherwise.

### GetMegabitsOk

`func (o *NetworkInterfaceRecord) GetMegabitsOk() (*int64, bool)`

GetMegabitsOk returns a tuple with the Megabits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMegabits

`func (o *NetworkInterfaceRecord) SetMegabits(v int64)`

SetMegabits sets Megabits field to given value.

### HasMegabits

`func (o *NetworkInterfaceRecord) HasMegabits() bool`

HasMegabits returns a boolean if a field has been set.

### GetName

`func (o *NetworkInterfaceRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkInterfaceRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkInterfaceRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkInterfaceRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperstate

`func (o *NetworkInterfaceRecord) GetOperstate() string`

GetOperstate returns the Operstate field if non-nil, zero value otherwise.

### GetOperstateOk

`func (o *NetworkInterfaceRecord) GetOperstateOk() (*string, bool)`

GetOperstateOk returns a tuple with the Operstate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperstate

`func (o *NetworkInterfaceRecord) SetOperstate(v string)`

SetOperstate sets Operstate field to given value.

### HasOperstate

`func (o *NetworkInterfaceRecord) HasOperstate() bool`

HasOperstate returns a boolean if a field has been set.

### GetOspCluster

`func (o *NetworkInterfaceRecord) GetOspCluster() Cluster`

GetOspCluster returns the OspCluster field if non-nil, zero value otherwise.

### GetOspClusterOk

`func (o *NetworkInterfaceRecord) GetOspClusterOk() (*Cluster, bool)`

GetOspClusterOk returns a tuple with the OspCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspCluster

`func (o *NetworkInterfaceRecord) SetOspCluster(v Cluster)`

SetOspCluster sets OspCluster field to given value.

### HasOspCluster

`func (o *NetworkInterfaceRecord) HasOspCluster() bool`

HasOspCluster returns a boolean if a field has been set.

### GetType

`func (o *NetworkInterfaceRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkInterfaceRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkInterfaceRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkInterfaceRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *NetworkInterfaceRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *NetworkInterfaceRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *NetworkInterfaceRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *NetworkInterfaceRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVfNum

`func (o *NetworkInterfaceRecord) GetVfNum() int64`

GetVfNum returns the VfNum field if non-nil, zero value otherwise.

### GetVfNumOk

`func (o *NetworkInterfaceRecord) GetVfNumOk() (*int64, bool)`

GetVfNumOk returns a tuple with the VfNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVfNum

`func (o *NetworkInterfaceRecord) SetVfNum(v int64)`

SetVfNum sets VfNum field to given value.

### HasVfNum

`func (o *NetworkInterfaceRecord) HasVfNum() bool`

HasVfNum returns a boolean if a field has been set.

### GetSamples

`func (o *NetworkInterfaceRecord) GetSamples() []NetworkInterfaceStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *NetworkInterfaceRecord) GetSamplesOk() (*[]NetworkInterfaceStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *NetworkInterfaceRecord) SetSamples(v []NetworkInterfaceStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *NetworkInterfaceRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


