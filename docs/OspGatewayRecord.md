# OspGatewayRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Etag** | Pointer to **string** |  | [optional] 
**GatewayId** | Pointer to **int64** |  | [optional] 
**GatewayIp** | Pointer to **string** |  | [optional] 
**GatewayName** | Pointer to **string** |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PlatformIp** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**S3Port** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**XmsPort** | Pointer to **int64** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**Samples** | Pointer to [**[]OspGatewayStat**](OspGatewayStat.md) |  | [optional] 

## Methods

### NewOspGatewayRecord

`func NewOspGatewayRecord() *OspGatewayRecord`

NewOspGatewayRecord instantiates a new OspGatewayRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspGatewayRecordWithDefaults

`func NewOspGatewayRecordWithDefaults() *OspGatewayRecord`

NewOspGatewayRecordWithDefaults instantiates a new OspGatewayRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *OspGatewayRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *OspGatewayRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *OspGatewayRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *OspGatewayRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *OspGatewayRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OspGatewayRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OspGatewayRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OspGatewayRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OspGatewayRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OspGatewayRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OspGatewayRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OspGatewayRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *OspGatewayRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OspGatewayRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OspGatewayRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OspGatewayRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEtag

`func (o *OspGatewayRecord) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *OspGatewayRecord) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *OspGatewayRecord) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *OspGatewayRecord) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetGatewayId

`func (o *OspGatewayRecord) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *OspGatewayRecord) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *OspGatewayRecord) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *OspGatewayRecord) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayIp

`func (o *OspGatewayRecord) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *OspGatewayRecord) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *OspGatewayRecord) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *OspGatewayRecord) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetGatewayName

`func (o *OspGatewayRecord) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *OspGatewayRecord) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *OspGatewayRecord) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.

### HasGatewayName

`func (o *OspGatewayRecord) HasGatewayName() bool`

HasGatewayName returns a boolean if a field has been set.

### GetHost

`func (o *OspGatewayRecord) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *OspGatewayRecord) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *OspGatewayRecord) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *OspGatewayRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *OspGatewayRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OspGatewayRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OspGatewayRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OspGatewayRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OspGatewayRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OspGatewayRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OspGatewayRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OspGatewayRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatformIp

`func (o *OspGatewayRecord) GetPlatformIp() string`

GetPlatformIp returns the PlatformIp field if non-nil, zero value otherwise.

### GetPlatformIpOk

`func (o *OspGatewayRecord) GetPlatformIpOk() (*string, bool)`

GetPlatformIpOk returns a tuple with the PlatformIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIp

`func (o *OspGatewayRecord) SetPlatformIp(v string)`

SetPlatformIp sets PlatformIp field to given value.

### HasPlatformIp

`func (o *OspGatewayRecord) HasPlatformIp() bool`

HasPlatformIp returns a boolean if a field has been set.

### GetRole

`func (o *OspGatewayRecord) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OspGatewayRecord) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OspGatewayRecord) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OspGatewayRecord) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetS3Port

`func (o *OspGatewayRecord) GetS3Port() int64`

GetS3Port returns the S3Port field if non-nil, zero value otherwise.

### GetS3PortOk

`func (o *OspGatewayRecord) GetS3PortOk() (*int64, bool)`

GetS3PortOk returns a tuple with the S3Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Port

`func (o *OspGatewayRecord) SetS3Port(v int64)`

SetS3Port sets S3Port field to given value.

### HasS3Port

`func (o *OspGatewayRecord) HasS3Port() bool`

HasS3Port returns a boolean if a field has been set.

### GetStatus

`func (o *OspGatewayRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OspGatewayRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OspGatewayRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OspGatewayRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetXmsPort

`func (o *OspGatewayRecord) GetXmsPort() int64`

GetXmsPort returns the XmsPort field if non-nil, zero value otherwise.

### GetXmsPortOk

`func (o *OspGatewayRecord) GetXmsPortOk() (*int64, bool)`

GetXmsPortOk returns a tuple with the XmsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmsPort

`func (o *OspGatewayRecord) SetXmsPort(v int64)`

SetXmsPort sets XmsPort field to given value.

### HasXmsPort

`func (o *OspGatewayRecord) HasXmsPort() bool`

HasXmsPort returns a boolean if a field has been set.

### GetZoneId

`func (o *OspGatewayRecord) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *OspGatewayRecord) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *OspGatewayRecord) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *OspGatewayRecord) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetSamples

`func (o *OspGatewayRecord) GetSamples() []OspGatewayStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *OspGatewayRecord) GetSamplesOk() (*[]OspGatewayStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *OspGatewayRecord) SetSamples(v []OspGatewayStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *OspGatewayRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


