# DfsGatewayZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**DfsGatewayGroup** | Pointer to [**DfsGatewayGroupNestview**](DfsGatewayGroupNestview.md) |  | [optional] 
**DfsHdfsNum** | Pointer to **int64** |  | [optional] 
**DfsHdfses** | Pointer to [**[]DfsHdfs**](DfsHdfs.md) |  | [optional] 
**DnsDomainNames** | Pointer to [**[]DNSDomainName**](DNSDomainName.md) |  | [optional] 
**GatewayNum** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**VipBalanceStatus** | Pointer to **string** |  | [optional] 
**VipGateways** | Pointer to **[]string** | gateways for VIP addresses in policy routing | [optional] 
**VipNum** | Pointer to **int64** |  | [optional] 

## Methods

### NewDfsGatewayZone

`func NewDfsGatewayZone() *DfsGatewayZone`

NewDfsGatewayZone instantiates a new DfsGatewayZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsGatewayZoneWithDefaults

`func NewDfsGatewayZoneWithDefaults() *DfsGatewayZone`

NewDfsGatewayZoneWithDefaults instantiates a new DfsGatewayZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DfsGatewayZone) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsGatewayZone) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsGatewayZone) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsGatewayZone) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *DfsGatewayZone) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsGatewayZone) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsGatewayZone) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsGatewayZone) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsGatewayZone) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsGatewayZone) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsGatewayZone) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsGatewayZone) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDefault

`func (o *DfsGatewayZone) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DfsGatewayZone) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *DfsGatewayZone) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *DfsGatewayZone) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDfsGatewayGroup

`func (o *DfsGatewayZone) GetDfsGatewayGroup() DfsGatewayGroupNestview`

GetDfsGatewayGroup returns the DfsGatewayGroup field if non-nil, zero value otherwise.

### GetDfsGatewayGroupOk

`func (o *DfsGatewayZone) GetDfsGatewayGroupOk() (*DfsGatewayGroupNestview, bool)`

GetDfsGatewayGroupOk returns a tuple with the DfsGatewayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroup

`func (o *DfsGatewayZone) SetDfsGatewayGroup(v DfsGatewayGroupNestview)`

SetDfsGatewayGroup sets DfsGatewayGroup field to given value.

### HasDfsGatewayGroup

`func (o *DfsGatewayZone) HasDfsGatewayGroup() bool`

HasDfsGatewayGroup returns a boolean if a field has been set.

### GetDfsHdfsNum

`func (o *DfsGatewayZone) GetDfsHdfsNum() int64`

GetDfsHdfsNum returns the DfsHdfsNum field if non-nil, zero value otherwise.

### GetDfsHdfsNumOk

`func (o *DfsGatewayZone) GetDfsHdfsNumOk() (*int64, bool)`

GetDfsHdfsNumOk returns a tuple with the DfsHdfsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfsNum

`func (o *DfsGatewayZone) SetDfsHdfsNum(v int64)`

SetDfsHdfsNum sets DfsHdfsNum field to given value.

### HasDfsHdfsNum

`func (o *DfsGatewayZone) HasDfsHdfsNum() bool`

HasDfsHdfsNum returns a boolean if a field has been set.

### GetDfsHdfses

`func (o *DfsGatewayZone) GetDfsHdfses() []DfsHdfs`

GetDfsHdfses returns the DfsHdfses field if non-nil, zero value otherwise.

### GetDfsHdfsesOk

`func (o *DfsGatewayZone) GetDfsHdfsesOk() (*[]DfsHdfs, bool)`

GetDfsHdfsesOk returns a tuple with the DfsHdfses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfses

`func (o *DfsGatewayZone) SetDfsHdfses(v []DfsHdfs)`

SetDfsHdfses sets DfsHdfses field to given value.

### HasDfsHdfses

`func (o *DfsGatewayZone) HasDfsHdfses() bool`

HasDfsHdfses returns a boolean if a field has been set.

### GetDnsDomainNames

`func (o *DfsGatewayZone) GetDnsDomainNames() []DNSDomainName`

GetDnsDomainNames returns the DnsDomainNames field if non-nil, zero value otherwise.

### GetDnsDomainNamesOk

`func (o *DfsGatewayZone) GetDnsDomainNamesOk() (*[]DNSDomainName, bool)`

GetDnsDomainNamesOk returns a tuple with the DnsDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomainNames

`func (o *DfsGatewayZone) SetDnsDomainNames(v []DNSDomainName)`

SetDnsDomainNames sets DnsDomainNames field to given value.

### HasDnsDomainNames

`func (o *DfsGatewayZone) HasDnsDomainNames() bool`

HasDnsDomainNames returns a boolean if a field has been set.

### GetGatewayNum

`func (o *DfsGatewayZone) GetGatewayNum() int64`

GetGatewayNum returns the GatewayNum field if non-nil, zero value otherwise.

### GetGatewayNumOk

`func (o *DfsGatewayZone) GetGatewayNumOk() (*int64, bool)`

GetGatewayNumOk returns a tuple with the GatewayNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayNum

`func (o *DfsGatewayZone) SetGatewayNum(v int64)`

SetGatewayNum sets GatewayNum field to given value.

### HasGatewayNum

`func (o *DfsGatewayZone) HasGatewayNum() bool`

HasGatewayNum returns a boolean if a field has been set.

### GetId

`func (o *DfsGatewayZone) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsGatewayZone) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsGatewayZone) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsGatewayZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DfsGatewayZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsGatewayZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsGatewayZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsGatewayZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DfsGatewayZone) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsGatewayZone) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsGatewayZone) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsGatewayZone) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsGatewayZone) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsGatewayZone) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsGatewayZone) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsGatewayZone) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVipBalanceStatus

`func (o *DfsGatewayZone) GetVipBalanceStatus() string`

GetVipBalanceStatus returns the VipBalanceStatus field if non-nil, zero value otherwise.

### GetVipBalanceStatusOk

`func (o *DfsGatewayZone) GetVipBalanceStatusOk() (*string, bool)`

GetVipBalanceStatusOk returns a tuple with the VipBalanceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipBalanceStatus

`func (o *DfsGatewayZone) SetVipBalanceStatus(v string)`

SetVipBalanceStatus sets VipBalanceStatus field to given value.

### HasVipBalanceStatus

`func (o *DfsGatewayZone) HasVipBalanceStatus() bool`

HasVipBalanceStatus returns a boolean if a field has been set.

### GetVipGateways

`func (o *DfsGatewayZone) GetVipGateways() []string`

GetVipGateways returns the VipGateways field if non-nil, zero value otherwise.

### GetVipGatewaysOk

`func (o *DfsGatewayZone) GetVipGatewaysOk() (*[]string, bool)`

GetVipGatewaysOk returns a tuple with the VipGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipGateways

`func (o *DfsGatewayZone) SetVipGateways(v []string)`

SetVipGateways sets VipGateways field to given value.

### HasVipGateways

`func (o *DfsGatewayZone) HasVipGateways() bool`

HasVipGateways returns a boolean if a field has been set.

### GetVipNum

`func (o *DfsGatewayZone) GetVipNum() int64`

GetVipNum returns the VipNum field if non-nil, zero value otherwise.

### GetVipNumOk

`func (o *DfsGatewayZone) GetVipNumOk() (*int64, bool)`

GetVipNumOk returns a tuple with the VipNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipNum

`func (o *DfsGatewayZone) SetVipNum(v int64)`

SetVipNum sets VipNum field to given value.

### HasVipNum

`func (o *DfsGatewayZone) HasVipNum() bool`

HasVipNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


