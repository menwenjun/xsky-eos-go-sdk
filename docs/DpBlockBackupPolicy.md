# DpBlockBackupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockVolumeNum** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DpGateway** | Pointer to [**DpGatewayNestview**](DpGatewayNestview.md) |  | [optional] 
**DpSite** | Pointer to [**DpSiteNestview**](DpSiteNestview.md) |  | [optional] 
**Encrypted** | Pointer to **bool** |  | [optional] 
**FullCronExpr** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IncCronExpr** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RetainedMax** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDpBlockBackupPolicy

`func NewDpBlockBackupPolicy() *DpBlockBackupPolicy`

NewDpBlockBackupPolicy instantiates a new DpBlockBackupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockBackupPolicyWithDefaults

`func NewDpBlockBackupPolicyWithDefaults() *DpBlockBackupPolicy`

NewDpBlockBackupPolicyWithDefaults instantiates a new DpBlockBackupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockVolumeNum

`func (o *DpBlockBackupPolicy) GetBlockVolumeNum() int64`

GetBlockVolumeNum returns the BlockVolumeNum field if non-nil, zero value otherwise.

### GetBlockVolumeNumOk

`func (o *DpBlockBackupPolicy) GetBlockVolumeNumOk() (*int64, bool)`

GetBlockVolumeNumOk returns a tuple with the BlockVolumeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeNum

`func (o *DpBlockBackupPolicy) SetBlockVolumeNum(v int64)`

SetBlockVolumeNum sets BlockVolumeNum field to given value.

### HasBlockVolumeNum

`func (o *DpBlockBackupPolicy) HasBlockVolumeNum() bool`

HasBlockVolumeNum returns a boolean if a field has been set.

### GetCluster

`func (o *DpBlockBackupPolicy) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DpBlockBackupPolicy) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DpBlockBackupPolicy) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DpBlockBackupPolicy) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DpBlockBackupPolicy) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DpBlockBackupPolicy) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DpBlockBackupPolicy) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DpBlockBackupPolicy) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDpGateway

`func (o *DpBlockBackupPolicy) GetDpGateway() DpGatewayNestview`

GetDpGateway returns the DpGateway field if non-nil, zero value otherwise.

### GetDpGatewayOk

`func (o *DpBlockBackupPolicy) GetDpGatewayOk() (*DpGatewayNestview, bool)`

GetDpGatewayOk returns a tuple with the DpGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpGateway

`func (o *DpBlockBackupPolicy) SetDpGateway(v DpGatewayNestview)`

SetDpGateway sets DpGateway field to given value.

### HasDpGateway

`func (o *DpBlockBackupPolicy) HasDpGateway() bool`

HasDpGateway returns a boolean if a field has been set.

### GetDpSite

`func (o *DpBlockBackupPolicy) GetDpSite() DpSiteNestview`

GetDpSite returns the DpSite field if non-nil, zero value otherwise.

### GetDpSiteOk

`func (o *DpBlockBackupPolicy) GetDpSiteOk() (*DpSiteNestview, bool)`

GetDpSiteOk returns a tuple with the DpSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpSite

`func (o *DpBlockBackupPolicy) SetDpSite(v DpSiteNestview)`

SetDpSite sets DpSite field to given value.

### HasDpSite

`func (o *DpBlockBackupPolicy) HasDpSite() bool`

HasDpSite returns a boolean if a field has been set.

### GetEncrypted

`func (o *DpBlockBackupPolicy) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *DpBlockBackupPolicy) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *DpBlockBackupPolicy) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *DpBlockBackupPolicy) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetFullCronExpr

`func (o *DpBlockBackupPolicy) GetFullCronExpr() string`

GetFullCronExpr returns the FullCronExpr field if non-nil, zero value otherwise.

### GetFullCronExprOk

`func (o *DpBlockBackupPolicy) GetFullCronExprOk() (*string, bool)`

GetFullCronExprOk returns a tuple with the FullCronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullCronExpr

`func (o *DpBlockBackupPolicy) SetFullCronExpr(v string)`

SetFullCronExpr sets FullCronExpr field to given value.

### HasFullCronExpr

`func (o *DpBlockBackupPolicy) HasFullCronExpr() bool`

HasFullCronExpr returns a boolean if a field has been set.

### GetId

`func (o *DpBlockBackupPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpBlockBackupPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpBlockBackupPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DpBlockBackupPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncCronExpr

`func (o *DpBlockBackupPolicy) GetIncCronExpr() string`

GetIncCronExpr returns the IncCronExpr field if non-nil, zero value otherwise.

### GetIncCronExprOk

`func (o *DpBlockBackupPolicy) GetIncCronExprOk() (*string, bool)`

GetIncCronExprOk returns a tuple with the IncCronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncCronExpr

`func (o *DpBlockBackupPolicy) SetIncCronExpr(v string)`

SetIncCronExpr sets IncCronExpr field to given value.

### HasIncCronExpr

`func (o *DpBlockBackupPolicy) HasIncCronExpr() bool`

HasIncCronExpr returns a boolean if a field has been set.

### GetName

`func (o *DpBlockBackupPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpBlockBackupPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpBlockBackupPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DpBlockBackupPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetainedMax

`func (o *DpBlockBackupPolicy) GetRetainedMax() int64`

GetRetainedMax returns the RetainedMax field if non-nil, zero value otherwise.

### GetRetainedMaxOk

`func (o *DpBlockBackupPolicy) GetRetainedMaxOk() (*int64, bool)`

GetRetainedMaxOk returns a tuple with the RetainedMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedMax

`func (o *DpBlockBackupPolicy) SetRetainedMax(v int64)`

SetRetainedMax sets RetainedMax field to given value.

### HasRetainedMax

`func (o *DpBlockBackupPolicy) HasRetainedMax() bool`

HasRetainedMax returns a boolean if a field has been set.

### GetStatus

`func (o *DpBlockBackupPolicy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpBlockBackupPolicy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpBlockBackupPolicy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DpBlockBackupPolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DpBlockBackupPolicy) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DpBlockBackupPolicy) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DpBlockBackupPolicy) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DpBlockBackupPolicy) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


