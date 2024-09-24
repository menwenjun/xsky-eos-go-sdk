# NetworkDiagnosisItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**DstHost** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**DstInterface** | Pointer to **string** |  | [optional] 
**DstIp** | Pointer to **string** |  | [optional] 
**DstMegabits** | Pointer to **int64** |  | [optional] 
**DstType** | Pointer to **string** |  | [optional] 
**Execute** | Pointer to **time.Time** |  | [optional] 
**Finish** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**NetworkDiagnosis** | Pointer to [**NetworkDiagnosisNestview**](NetworkDiagnosisNestview.md) |  | [optional] 
**Networks** | Pointer to **[]string** |  | [optional] 
**SrcHost** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**SrcInterface** | Pointer to **string** |  | [optional] 
**SrcIp** | Pointer to **string** |  | [optional] 
**SrcMegabits** | Pointer to **int64** |  | [optional] 
**SrcType** | Pointer to **string** |  | [optional] 
**Stat** | Pointer to [**NetworkDiagnosisStat**](NetworkDiagnosisStat.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewNetworkDiagnosisItem

`func NewNetworkDiagnosisItem() *NetworkDiagnosisItem`

NewNetworkDiagnosisItem instantiates a new NetworkDiagnosisItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDiagnosisItemWithDefaults

`func NewNetworkDiagnosisItemWithDefaults() *NetworkDiagnosisItem`

NewNetworkDiagnosisItemWithDefaults instantiates a new NetworkDiagnosisItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *NetworkDiagnosisItem) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *NetworkDiagnosisItem) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *NetworkDiagnosisItem) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *NetworkDiagnosisItem) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDstHost

`func (o *NetworkDiagnosisItem) GetDstHost() HostNestview`

GetDstHost returns the DstHost field if non-nil, zero value otherwise.

### GetDstHostOk

`func (o *NetworkDiagnosisItem) GetDstHostOk() (*HostNestview, bool)`

GetDstHostOk returns a tuple with the DstHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstHost

`func (o *NetworkDiagnosisItem) SetDstHost(v HostNestview)`

SetDstHost sets DstHost field to given value.

### HasDstHost

`func (o *NetworkDiagnosisItem) HasDstHost() bool`

HasDstHost returns a boolean if a field has been set.

### GetDstInterface

`func (o *NetworkDiagnosisItem) GetDstInterface() string`

GetDstInterface returns the DstInterface field if non-nil, zero value otherwise.

### GetDstInterfaceOk

`func (o *NetworkDiagnosisItem) GetDstInterfaceOk() (*string, bool)`

GetDstInterfaceOk returns a tuple with the DstInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstInterface

`func (o *NetworkDiagnosisItem) SetDstInterface(v string)`

SetDstInterface sets DstInterface field to given value.

### HasDstInterface

`func (o *NetworkDiagnosisItem) HasDstInterface() bool`

HasDstInterface returns a boolean if a field has been set.

### GetDstIp

`func (o *NetworkDiagnosisItem) GetDstIp() string`

GetDstIp returns the DstIp field if non-nil, zero value otherwise.

### GetDstIpOk

`func (o *NetworkDiagnosisItem) GetDstIpOk() (*string, bool)`

GetDstIpOk returns a tuple with the DstIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstIp

`func (o *NetworkDiagnosisItem) SetDstIp(v string)`

SetDstIp sets DstIp field to given value.

### HasDstIp

`func (o *NetworkDiagnosisItem) HasDstIp() bool`

HasDstIp returns a boolean if a field has been set.

### GetDstMegabits

`func (o *NetworkDiagnosisItem) GetDstMegabits() int64`

GetDstMegabits returns the DstMegabits field if non-nil, zero value otherwise.

### GetDstMegabitsOk

`func (o *NetworkDiagnosisItem) GetDstMegabitsOk() (*int64, bool)`

GetDstMegabitsOk returns a tuple with the DstMegabits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstMegabits

`func (o *NetworkDiagnosisItem) SetDstMegabits(v int64)`

SetDstMegabits sets DstMegabits field to given value.

### HasDstMegabits

`func (o *NetworkDiagnosisItem) HasDstMegabits() bool`

HasDstMegabits returns a boolean if a field has been set.

### GetDstType

`func (o *NetworkDiagnosisItem) GetDstType() string`

GetDstType returns the DstType field if non-nil, zero value otherwise.

### GetDstTypeOk

`func (o *NetworkDiagnosisItem) GetDstTypeOk() (*string, bool)`

GetDstTypeOk returns a tuple with the DstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstType

`func (o *NetworkDiagnosisItem) SetDstType(v string)`

SetDstType sets DstType field to given value.

### HasDstType

`func (o *NetworkDiagnosisItem) HasDstType() bool`

HasDstType returns a boolean if a field has been set.

### GetExecute

`func (o *NetworkDiagnosisItem) GetExecute() time.Time`

GetExecute returns the Execute field if non-nil, zero value otherwise.

### GetExecuteOk

`func (o *NetworkDiagnosisItem) GetExecuteOk() (*time.Time, bool)`

GetExecuteOk returns a tuple with the Execute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecute

`func (o *NetworkDiagnosisItem) SetExecute(v time.Time)`

SetExecute sets Execute field to given value.

### HasExecute

`func (o *NetworkDiagnosisItem) HasExecute() bool`

HasExecute returns a boolean if a field has been set.

### GetFinish

`func (o *NetworkDiagnosisItem) GetFinish() time.Time`

GetFinish returns the Finish field if non-nil, zero value otherwise.

### GetFinishOk

`func (o *NetworkDiagnosisItem) GetFinishOk() (*time.Time, bool)`

GetFinishOk returns a tuple with the Finish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinish

`func (o *NetworkDiagnosisItem) SetFinish(v time.Time)`

SetFinish sets Finish field to given value.

### HasFinish

`func (o *NetworkDiagnosisItem) HasFinish() bool`

HasFinish returns a boolean if a field has been set.

### GetId

`func (o *NetworkDiagnosisItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkDiagnosisItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkDiagnosisItem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkDiagnosisItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *NetworkDiagnosisItem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NetworkDiagnosisItem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NetworkDiagnosisItem) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NetworkDiagnosisItem) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNetworkDiagnosis

`func (o *NetworkDiagnosisItem) GetNetworkDiagnosis() NetworkDiagnosisNestview`

GetNetworkDiagnosis returns the NetworkDiagnosis field if non-nil, zero value otherwise.

### GetNetworkDiagnosisOk

`func (o *NetworkDiagnosisItem) GetNetworkDiagnosisOk() (*NetworkDiagnosisNestview, bool)`

GetNetworkDiagnosisOk returns a tuple with the NetworkDiagnosis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDiagnosis

`func (o *NetworkDiagnosisItem) SetNetworkDiagnosis(v NetworkDiagnosisNestview)`

SetNetworkDiagnosis sets NetworkDiagnosis field to given value.

### HasNetworkDiagnosis

`func (o *NetworkDiagnosisItem) HasNetworkDiagnosis() bool`

HasNetworkDiagnosis returns a boolean if a field has been set.

### GetNetworks

`func (o *NetworkDiagnosisItem) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *NetworkDiagnosisItem) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *NetworkDiagnosisItem) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *NetworkDiagnosisItem) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetSrcHost

`func (o *NetworkDiagnosisItem) GetSrcHost() HostNestview`

GetSrcHost returns the SrcHost field if non-nil, zero value otherwise.

### GetSrcHostOk

`func (o *NetworkDiagnosisItem) GetSrcHostOk() (*HostNestview, bool)`

GetSrcHostOk returns a tuple with the SrcHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcHost

`func (o *NetworkDiagnosisItem) SetSrcHost(v HostNestview)`

SetSrcHost sets SrcHost field to given value.

### HasSrcHost

`func (o *NetworkDiagnosisItem) HasSrcHost() bool`

HasSrcHost returns a boolean if a field has been set.

### GetSrcInterface

`func (o *NetworkDiagnosisItem) GetSrcInterface() string`

GetSrcInterface returns the SrcInterface field if non-nil, zero value otherwise.

### GetSrcInterfaceOk

`func (o *NetworkDiagnosisItem) GetSrcInterfaceOk() (*string, bool)`

GetSrcInterfaceOk returns a tuple with the SrcInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcInterface

`func (o *NetworkDiagnosisItem) SetSrcInterface(v string)`

SetSrcInterface sets SrcInterface field to given value.

### HasSrcInterface

`func (o *NetworkDiagnosisItem) HasSrcInterface() bool`

HasSrcInterface returns a boolean if a field has been set.

### GetSrcIp

`func (o *NetworkDiagnosisItem) GetSrcIp() string`

GetSrcIp returns the SrcIp field if non-nil, zero value otherwise.

### GetSrcIpOk

`func (o *NetworkDiagnosisItem) GetSrcIpOk() (*string, bool)`

GetSrcIpOk returns a tuple with the SrcIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIp

`func (o *NetworkDiagnosisItem) SetSrcIp(v string)`

SetSrcIp sets SrcIp field to given value.

### HasSrcIp

`func (o *NetworkDiagnosisItem) HasSrcIp() bool`

HasSrcIp returns a boolean if a field has been set.

### GetSrcMegabits

`func (o *NetworkDiagnosisItem) GetSrcMegabits() int64`

GetSrcMegabits returns the SrcMegabits field if non-nil, zero value otherwise.

### GetSrcMegabitsOk

`func (o *NetworkDiagnosisItem) GetSrcMegabitsOk() (*int64, bool)`

GetSrcMegabitsOk returns a tuple with the SrcMegabits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcMegabits

`func (o *NetworkDiagnosisItem) SetSrcMegabits(v int64)`

SetSrcMegabits sets SrcMegabits field to given value.

### HasSrcMegabits

`func (o *NetworkDiagnosisItem) HasSrcMegabits() bool`

HasSrcMegabits returns a boolean if a field has been set.

### GetSrcType

`func (o *NetworkDiagnosisItem) GetSrcType() string`

GetSrcType returns the SrcType field if non-nil, zero value otherwise.

### GetSrcTypeOk

`func (o *NetworkDiagnosisItem) GetSrcTypeOk() (*string, bool)`

GetSrcTypeOk returns a tuple with the SrcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcType

`func (o *NetworkDiagnosisItem) SetSrcType(v string)`

SetSrcType sets SrcType field to given value.

### HasSrcType

`func (o *NetworkDiagnosisItem) HasSrcType() bool`

HasSrcType returns a boolean if a field has been set.

### GetStat

`func (o *NetworkDiagnosisItem) GetStat() NetworkDiagnosisStat`

GetStat returns the Stat field if non-nil, zero value otherwise.

### GetStatOk

`func (o *NetworkDiagnosisItem) GetStatOk() (*NetworkDiagnosisStat, bool)`

GetStatOk returns a tuple with the Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStat

`func (o *NetworkDiagnosisItem) SetStat(v NetworkDiagnosisStat)`

SetStat sets Stat field to given value.

### HasStat

`func (o *NetworkDiagnosisItem) HasStat() bool`

HasStat returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkDiagnosisItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkDiagnosisItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkDiagnosisItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkDiagnosisItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *NetworkDiagnosisItem) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *NetworkDiagnosisItem) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *NetworkDiagnosisItem) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *NetworkDiagnosisItem) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


