# DpBlockBackupPolicyCreateReqPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DpGatewayId** | **int64** | dp gateway | 
**DpSiteId** | **int64** | dp site | 
**Encrypted** | Pointer to **bool** | encrypted or not | [optional] 
**FullCronExpr** | **string** | cron expression for full backup | 
**IncCronExpr** | Pointer to **string** | cron expression for increment backup | [optional] 
**Name** | **string** | name | 
**RetainedMax** | **int64** | retained max number of backups | 

## Methods

### NewDpBlockBackupPolicyCreateReqPolicy

`func NewDpBlockBackupPolicyCreateReqPolicy(dpGatewayId int64, dpSiteId int64, fullCronExpr string, name string, retainedMax int64, ) *DpBlockBackupPolicyCreateReqPolicy`

NewDpBlockBackupPolicyCreateReqPolicy instantiates a new DpBlockBackupPolicyCreateReqPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockBackupPolicyCreateReqPolicyWithDefaults

`func NewDpBlockBackupPolicyCreateReqPolicyWithDefaults() *DpBlockBackupPolicyCreateReqPolicy`

NewDpBlockBackupPolicyCreateReqPolicyWithDefaults instantiates a new DpBlockBackupPolicyCreateReqPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDpGatewayId

`func (o *DpBlockBackupPolicyCreateReqPolicy) GetDpGatewayId() int64`

GetDpGatewayId returns the DpGatewayId field if non-nil, zero value otherwise.

### GetDpGatewayIdOk

`func (o *DpBlockBackupPolicyCreateReqPolicy) GetDpGatewayIdOk() (*int64, bool)`

GetDpGatewayIdOk returns a tuple with the DpGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpGatewayId

`func (o *DpBlockBackupPolicyCreateReqPolicy) SetDpGatewayId(v int64)`

SetDpGatewayId sets DpGatewayId field to given value.


### GetDpSiteId

`func (o *DpBlockBackupPolicyCreateReqPolicy) GetDpSiteId() int64`

GetDpSiteId returns the DpSiteId field if non-nil, zero value otherwise.

### GetDpSiteIdOk

`func (o *DpBlockBackupPolicyCreateReqPolicy) GetDpSiteIdOk() (*int64, bool)`

GetDpSiteIdOk returns a tuple with the DpSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpSiteId

`func (o *DpBlockBackupPolicyCreateReqPolicy) SetDpSiteId(v int64)`

SetDpSiteId sets DpSiteId field to given value.


### GetEncrypted

`func (o *DpBlockBackupPolicyCreateReqPolicy) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *DpBlockBackupPolicyCreateReqPolicy) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *DpBlockBackupPolicyCreateReqPolicy) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *DpBlockBackupPolicyCreateReqPolicy) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetFullCronExpr

`func (o *DpBlockBackupPolicyCreateReqPolicy) GetFullCronExpr() string`

GetFullCronExpr returns the FullCronExpr field if non-nil, zero value otherwise.

### GetFullCronExprOk

`func (o *DpBlockBackupPolicyCreateReqPolicy) GetFullCronExprOk() (*string, bool)`

GetFullCronExprOk returns a tuple with the FullCronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullCronExpr

`func (o *DpBlockBackupPolicyCreateReqPolicy) SetFullCronExpr(v string)`

SetFullCronExpr sets FullCronExpr field to given value.


### GetIncCronExpr

`func (o *DpBlockBackupPolicyCreateReqPolicy) GetIncCronExpr() string`

GetIncCronExpr returns the IncCronExpr field if non-nil, zero value otherwise.

### GetIncCronExprOk

`func (o *DpBlockBackupPolicyCreateReqPolicy) GetIncCronExprOk() (*string, bool)`

GetIncCronExprOk returns a tuple with the IncCronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncCronExpr

`func (o *DpBlockBackupPolicyCreateReqPolicy) SetIncCronExpr(v string)`

SetIncCronExpr sets IncCronExpr field to given value.

### HasIncCronExpr

`func (o *DpBlockBackupPolicyCreateReqPolicy) HasIncCronExpr() bool`

HasIncCronExpr returns a boolean if a field has been set.

### GetName

`func (o *DpBlockBackupPolicyCreateReqPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpBlockBackupPolicyCreateReqPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpBlockBackupPolicyCreateReqPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetRetainedMax

`func (o *DpBlockBackupPolicyCreateReqPolicy) GetRetainedMax() int64`

GetRetainedMax returns the RetainedMax field if non-nil, zero value otherwise.

### GetRetainedMaxOk

`func (o *DpBlockBackupPolicyCreateReqPolicy) GetRetainedMaxOk() (*int64, bool)`

GetRetainedMaxOk returns a tuple with the RetainedMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedMax

`func (o *DpBlockBackupPolicyCreateReqPolicy) SetRetainedMax(v int64)`

SetRetainedMax sets RetainedMax field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


