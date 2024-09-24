# FCPort

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

## Methods

### NewFCPort

`func NewFCPort() *FCPort`

NewFCPort instantiates a new FCPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFCPortWithDefaults

`func NewFCPortWithDefaults() *FCPort`

NewFCPortWithDefaults instantiates a new FCPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *FCPort) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *FCPort) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *FCPort) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *FCPort) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *FCPort) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *FCPort) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *FCPort) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *FCPort) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConnOptMode

`func (o *FCPort) GetConnOptMode() string`

GetConnOptMode returns the ConnOptMode field if non-nil, zero value otherwise.

### GetConnOptModeOk

`func (o *FCPort) GetConnOptModeOk() (*string, bool)`

GetConnOptModeOk returns a tuple with the ConnOptMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnOptMode

`func (o *FCPort) SetConnOptMode(v string)`

SetConnOptMode sets ConnOptMode field to given value.

### HasConnOptMode

`func (o *FCPort) HasConnOptMode() bool`

HasConnOptMode returns a boolean if a field has been set.

### GetConnType

`func (o *FCPort) GetConnType() string`

GetConnType returns the ConnType field if non-nil, zero value otherwise.

### GetConnTypeOk

`func (o *FCPort) GetConnTypeOk() (*string, bool)`

GetConnTypeOk returns a tuple with the ConnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnType

`func (o *FCPort) SetConnType(v string)`

SetConnType sets ConnType field to given value.

### HasConnType

`func (o *FCPort) HasConnType() bool`

HasConnType returns a boolean if a field has been set.

### GetCreate

`func (o *FCPort) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *FCPort) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *FCPort) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *FCPort) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDataRateMode

`func (o *FCPort) GetDataRateMode() string`

GetDataRateMode returns the DataRateMode field if non-nil, zero value otherwise.

### GetDataRateModeOk

`func (o *FCPort) GetDataRateModeOk() (*string, bool)`

GetDataRateModeOk returns a tuple with the DataRateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRateMode

`func (o *FCPort) SetDataRateMode(v string)`

SetDataRateMode sets DataRateMode field to given value.

### HasDataRateMode

`func (o *FCPort) HasDataRateMode() bool`

HasDataRateMode returns a boolean if a field has been set.

### GetHealth

`func (o *FCPort) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *FCPort) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *FCPort) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *FCPort) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHost

`func (o *FCPort) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *FCPort) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *FCPort) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *FCPort) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *FCPort) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FCPort) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FCPort) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FCPort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *FCPort) GetLinkSpeed() string`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *FCPort) GetLinkSpeedOk() (*string, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *FCPort) SetLinkSpeed(v string)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *FCPort) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLinkState

`func (o *FCPort) GetLinkState() string`

GetLinkState returns the LinkState field if non-nil, zero value otherwise.

### GetLinkStateOk

`func (o *FCPort) GetLinkStateOk() (*string, bool)`

GetLinkStateOk returns a tuple with the LinkState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkState

`func (o *FCPort) SetLinkState(v string)`

SetLinkState sets LinkState field to given value.

### HasLinkState

`func (o *FCPort) HasLinkState() bool`

HasLinkState returns a boolean if a field has been set.

### GetMaxSpeed

`func (o *FCPort) GetMaxSpeed() string`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *FCPort) GetMaxSpeedOk() (*string, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *FCPort) SetMaxSpeed(v string)`

SetMaxSpeed sets MaxSpeed field to given value.

### HasMaxSpeed

`func (o *FCPort) HasMaxSpeed() bool`

HasMaxSpeed returns a boolean if a field has been set.

### GetPciAddress

`func (o *FCPort) GetPciAddress() string`

GetPciAddress returns the PciAddress field if non-nil, zero value otherwise.

### GetPciAddressOk

`func (o *FCPort) GetPciAddressOk() (*string, bool)`

GetPciAddressOk returns a tuple with the PciAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddress

`func (o *FCPort) SetPciAddress(v string)`

SetPciAddress sets PciAddress field to given value.

### HasPciAddress

`func (o *FCPort) HasPciAddress() bool`

HasPciAddress returns a boolean if a field has been set.

### GetPortId

`func (o *FCPort) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *FCPort) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *FCPort) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *FCPort) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetRecvBytes

`func (o *FCPort) GetRecvBytes() int64`

GetRecvBytes returns the RecvBytes field if non-nil, zero value otherwise.

### GetRecvBytesOk

`func (o *FCPort) GetRecvBytesOk() (*int64, bool)`

GetRecvBytesOk returns a tuple with the RecvBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvBytes

`func (o *FCPort) SetRecvBytes(v int64)`

SetRecvBytes sets RecvBytes field to given value.

### HasRecvBytes

`func (o *FCPort) HasRecvBytes() bool`

HasRecvBytes returns a boolean if a field has been set.

### GetRecvFrames

`func (o *FCPort) GetRecvFrames() int64`

GetRecvFrames returns the RecvFrames field if non-nil, zero value otherwise.

### GetRecvFramesOk

`func (o *FCPort) GetRecvFramesOk() (*int64, bool)`

GetRecvFramesOk returns a tuple with the RecvFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvFrames

`func (o *FCPort) SetRecvFrames(v int64)`

SetRecvFrames sets RecvFrames field to given value.

### HasRecvFrames

`func (o *FCPort) HasRecvFrames() bool`

HasRecvFrames returns a boolean if a field has been set.

### GetRgHost

`func (o *FCPort) GetRgHost() int64`

GetRgHost returns the RgHost field if non-nil, zero value otherwise.

### GetRgHostOk

`func (o *FCPort) GetRgHostOk() (*int64, bool)`

GetRgHostOk returns a tuple with the RgHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRgHost

`func (o *FCPort) SetRgHost(v int64)`

SetRgHost sets RgHost field to given value.

### HasRgHost

`func (o *FCPort) HasRgHost() bool`

HasRgHost returns a boolean if a field has been set.

### GetSendBytes

`func (o *FCPort) GetSendBytes() int64`

GetSendBytes returns the SendBytes field if non-nil, zero value otherwise.

### GetSendBytesOk

`func (o *FCPort) GetSendBytesOk() (*int64, bool)`

GetSendBytesOk returns a tuple with the SendBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBytes

`func (o *FCPort) SetSendBytes(v int64)`

SetSendBytes sets SendBytes field to given value.

### HasSendBytes

`func (o *FCPort) HasSendBytes() bool`

HasSendBytes returns a boolean if a field has been set.

### GetSendFrames

`func (o *FCPort) GetSendFrames() int64`

GetSendFrames returns the SendFrames field if non-nil, zero value otherwise.

### GetSendFramesOk

`func (o *FCPort) GetSendFramesOk() (*int64, bool)`

GetSendFramesOk returns a tuple with the SendFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendFrames

`func (o *FCPort) SetSendFrames(v int64)`

SetSendFrames sets SendFrames field to given value.

### HasSendFrames

`func (o *FCPort) HasSendFrames() bool`

HasSendFrames returns a boolean if a field has been set.

### GetStart

`func (o *FCPort) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *FCPort) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *FCPort) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *FCPort) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSupportedSpeeds

`func (o *FCPort) GetSupportedSpeeds() []string`

GetSupportedSpeeds returns the SupportedSpeeds field if non-nil, zero value otherwise.

### GetSupportedSpeedsOk

`func (o *FCPort) GetSupportedSpeedsOk() (*[]string, bool)`

GetSupportedSpeedsOk returns a tuple with the SupportedSpeeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSpeeds

`func (o *FCPort) SetSupportedSpeeds(v []string)`

SetSupportedSpeeds sets SupportedSpeeds field to given value.

### HasSupportedSpeeds

`func (o *FCPort) HasSupportedSpeeds() bool`

HasSupportedSpeeds returns a boolean if a field has been set.

### GetUpdate

`func (o *FCPort) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *FCPort) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *FCPort) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *FCPort) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetWwpn

`func (o *FCPort) GetWwpn() string`

GetWwpn returns the Wwpn field if non-nil, zero value otherwise.

### GetWwpnOk

`func (o *FCPort) GetWwpnOk() (*string, bool)`

GetWwpnOk returns a tuple with the Wwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpn

`func (o *FCPort) SetWwpn(v string)`

SetWwpn sets Wwpn field to given value.

### HasWwpn

`func (o *FCPort) HasWwpn() bool`

HasWwpn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


