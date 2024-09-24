# FCPortRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**ConnOptMode** | Pointer to **string** |  | [optional] 
**ConnType** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DataRateMode** | Pointer to **string** |  | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LinkSpeed** | Pointer to **string** |  | [optional] 
**LinkState** | Pointer to **string** |  | [optional] 
**MaxSpeed** | Pointer to **string** |  | [optional] 
**PciAddress** | Pointer to **string** |  | [optional] 
**PortId** | Pointer to **string** |  | [optional] 
**RecvBytes** | Pointer to **int64** |  | [optional] 
**RecvFrames** | Pointer to **int64** |  | [optional] 
**RgHost** | Pointer to **int64** |  | [optional] 
**SendBytes** | Pointer to **int64** |  | [optional] 
**SendFrames** | Pointer to **int64** |  | [optional] 
**Start** | Pointer to **time.Time** |  | [optional] 
**SupportedSpeeds** | Pointer to **[]string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Wwpn** | Pointer to **string** |  | [optional] 
**FcPortErrcode** | Pointer to [**FCPortErrCode**](FCPortErrCode.md) |  | [optional] 

## Methods

### NewFCPortRecord

`func NewFCPortRecord() *FCPortRecord`

NewFCPortRecord instantiates a new FCPortRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFCPortRecordWithDefaults

`func NewFCPortRecordWithDefaults() *FCPortRecord`

NewFCPortRecordWithDefaults instantiates a new FCPortRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *FCPortRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *FCPortRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *FCPortRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *FCPortRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *FCPortRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *FCPortRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *FCPortRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *FCPortRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConnOptMode

`func (o *FCPortRecord) GetConnOptMode() string`

GetConnOptMode returns the ConnOptMode field if non-nil, zero value otherwise.

### GetConnOptModeOk

`func (o *FCPortRecord) GetConnOptModeOk() (*string, bool)`

GetConnOptModeOk returns a tuple with the ConnOptMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnOptMode

`func (o *FCPortRecord) SetConnOptMode(v string)`

SetConnOptMode sets ConnOptMode field to given value.

### HasConnOptMode

`func (o *FCPortRecord) HasConnOptMode() bool`

HasConnOptMode returns a boolean if a field has been set.

### GetConnType

`func (o *FCPortRecord) GetConnType() string`

GetConnType returns the ConnType field if non-nil, zero value otherwise.

### GetConnTypeOk

`func (o *FCPortRecord) GetConnTypeOk() (*string, bool)`

GetConnTypeOk returns a tuple with the ConnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnType

`func (o *FCPortRecord) SetConnType(v string)`

SetConnType sets ConnType field to given value.

### HasConnType

`func (o *FCPortRecord) HasConnType() bool`

HasConnType returns a boolean if a field has been set.

### GetCreate

`func (o *FCPortRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *FCPortRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *FCPortRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *FCPortRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDataRateMode

`func (o *FCPortRecord) GetDataRateMode() string`

GetDataRateMode returns the DataRateMode field if non-nil, zero value otherwise.

### GetDataRateModeOk

`func (o *FCPortRecord) GetDataRateModeOk() (*string, bool)`

GetDataRateModeOk returns a tuple with the DataRateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRateMode

`func (o *FCPortRecord) SetDataRateMode(v string)`

SetDataRateMode sets DataRateMode field to given value.

### HasDataRateMode

`func (o *FCPortRecord) HasDataRateMode() bool`

HasDataRateMode returns a boolean if a field has been set.

### GetHealth

`func (o *FCPortRecord) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *FCPortRecord) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *FCPortRecord) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *FCPortRecord) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHost

`func (o *FCPortRecord) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *FCPortRecord) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *FCPortRecord) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *FCPortRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *FCPortRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FCPortRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FCPortRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FCPortRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *FCPortRecord) GetLinkSpeed() string`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *FCPortRecord) GetLinkSpeedOk() (*string, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *FCPortRecord) SetLinkSpeed(v string)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *FCPortRecord) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLinkState

`func (o *FCPortRecord) GetLinkState() string`

GetLinkState returns the LinkState field if non-nil, zero value otherwise.

### GetLinkStateOk

`func (o *FCPortRecord) GetLinkStateOk() (*string, bool)`

GetLinkStateOk returns a tuple with the LinkState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkState

`func (o *FCPortRecord) SetLinkState(v string)`

SetLinkState sets LinkState field to given value.

### HasLinkState

`func (o *FCPortRecord) HasLinkState() bool`

HasLinkState returns a boolean if a field has been set.

### GetMaxSpeed

`func (o *FCPortRecord) GetMaxSpeed() string`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *FCPortRecord) GetMaxSpeedOk() (*string, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *FCPortRecord) SetMaxSpeed(v string)`

SetMaxSpeed sets MaxSpeed field to given value.

### HasMaxSpeed

`func (o *FCPortRecord) HasMaxSpeed() bool`

HasMaxSpeed returns a boolean if a field has been set.

### GetPciAddress

`func (o *FCPortRecord) GetPciAddress() string`

GetPciAddress returns the PciAddress field if non-nil, zero value otherwise.

### GetPciAddressOk

`func (o *FCPortRecord) GetPciAddressOk() (*string, bool)`

GetPciAddressOk returns a tuple with the PciAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddress

`func (o *FCPortRecord) SetPciAddress(v string)`

SetPciAddress sets PciAddress field to given value.

### HasPciAddress

`func (o *FCPortRecord) HasPciAddress() bool`

HasPciAddress returns a boolean if a field has been set.

### GetPortId

`func (o *FCPortRecord) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *FCPortRecord) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *FCPortRecord) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *FCPortRecord) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetRecvBytes

`func (o *FCPortRecord) GetRecvBytes() int64`

GetRecvBytes returns the RecvBytes field if non-nil, zero value otherwise.

### GetRecvBytesOk

`func (o *FCPortRecord) GetRecvBytesOk() (*int64, bool)`

GetRecvBytesOk returns a tuple with the RecvBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvBytes

`func (o *FCPortRecord) SetRecvBytes(v int64)`

SetRecvBytes sets RecvBytes field to given value.

### HasRecvBytes

`func (o *FCPortRecord) HasRecvBytes() bool`

HasRecvBytes returns a boolean if a field has been set.

### GetRecvFrames

`func (o *FCPortRecord) GetRecvFrames() int64`

GetRecvFrames returns the RecvFrames field if non-nil, zero value otherwise.

### GetRecvFramesOk

`func (o *FCPortRecord) GetRecvFramesOk() (*int64, bool)`

GetRecvFramesOk returns a tuple with the RecvFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvFrames

`func (o *FCPortRecord) SetRecvFrames(v int64)`

SetRecvFrames sets RecvFrames field to given value.

### HasRecvFrames

`func (o *FCPortRecord) HasRecvFrames() bool`

HasRecvFrames returns a boolean if a field has been set.

### GetRgHost

`func (o *FCPortRecord) GetRgHost() int64`

GetRgHost returns the RgHost field if non-nil, zero value otherwise.

### GetRgHostOk

`func (o *FCPortRecord) GetRgHostOk() (*int64, bool)`

GetRgHostOk returns a tuple with the RgHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRgHost

`func (o *FCPortRecord) SetRgHost(v int64)`

SetRgHost sets RgHost field to given value.

### HasRgHost

`func (o *FCPortRecord) HasRgHost() bool`

HasRgHost returns a boolean if a field has been set.

### GetSendBytes

`func (o *FCPortRecord) GetSendBytes() int64`

GetSendBytes returns the SendBytes field if non-nil, zero value otherwise.

### GetSendBytesOk

`func (o *FCPortRecord) GetSendBytesOk() (*int64, bool)`

GetSendBytesOk returns a tuple with the SendBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBytes

`func (o *FCPortRecord) SetSendBytes(v int64)`

SetSendBytes sets SendBytes field to given value.

### HasSendBytes

`func (o *FCPortRecord) HasSendBytes() bool`

HasSendBytes returns a boolean if a field has been set.

### GetSendFrames

`func (o *FCPortRecord) GetSendFrames() int64`

GetSendFrames returns the SendFrames field if non-nil, zero value otherwise.

### GetSendFramesOk

`func (o *FCPortRecord) GetSendFramesOk() (*int64, bool)`

GetSendFramesOk returns a tuple with the SendFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendFrames

`func (o *FCPortRecord) SetSendFrames(v int64)`

SetSendFrames sets SendFrames field to given value.

### HasSendFrames

`func (o *FCPortRecord) HasSendFrames() bool`

HasSendFrames returns a boolean if a field has been set.

### GetStart

`func (o *FCPortRecord) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *FCPortRecord) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *FCPortRecord) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *FCPortRecord) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSupportedSpeeds

`func (o *FCPortRecord) GetSupportedSpeeds() []string`

GetSupportedSpeeds returns the SupportedSpeeds field if non-nil, zero value otherwise.

### GetSupportedSpeedsOk

`func (o *FCPortRecord) GetSupportedSpeedsOk() (*[]string, bool)`

GetSupportedSpeedsOk returns a tuple with the SupportedSpeeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSpeeds

`func (o *FCPortRecord) SetSupportedSpeeds(v []string)`

SetSupportedSpeeds sets SupportedSpeeds field to given value.

### HasSupportedSpeeds

`func (o *FCPortRecord) HasSupportedSpeeds() bool`

HasSupportedSpeeds returns a boolean if a field has been set.

### GetUpdate

`func (o *FCPortRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *FCPortRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *FCPortRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *FCPortRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetWwpn

`func (o *FCPortRecord) GetWwpn() string`

GetWwpn returns the Wwpn field if non-nil, zero value otherwise.

### GetWwpnOk

`func (o *FCPortRecord) GetWwpnOk() (*string, bool)`

GetWwpnOk returns a tuple with the Wwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpn

`func (o *FCPortRecord) SetWwpn(v string)`

SetWwpn sets Wwpn field to given value.

### HasWwpn

`func (o *FCPortRecord) HasWwpn() bool`

HasWwpn returns a boolean if a field has been set.

### GetFcPortErrcode

`func (o *FCPortRecord) GetFcPortErrcode() FCPortErrCode`

GetFcPortErrcode returns the FcPortErrcode field if non-nil, zero value otherwise.

### GetFcPortErrcodeOk

`func (o *FCPortRecord) GetFcPortErrcodeOk() (*FCPortErrCode, bool)`

GetFcPortErrcodeOk returns a tuple with the FcPortErrcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPortErrcode

`func (o *FCPortRecord) SetFcPortErrcode(v FCPortErrCode)`

SetFcPortErrcode sets FcPortErrcode field to given value.

### HasFcPortErrcode

`func (o *FCPortRecord) HasFcPortErrcode() bool`

HasFcPortErrcode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


