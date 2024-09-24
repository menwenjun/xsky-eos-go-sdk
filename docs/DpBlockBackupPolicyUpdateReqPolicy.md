# DpBlockBackupPolicyUpdateReqPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Encrypted** | Pointer to **bool** | encrypted or not | [optional] 
**FullCronExpr** | Pointer to **string** | cron expression for full backup | [optional] 
**IncCronExpr** | Pointer to **string** | cron expression for increment backup | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**RetainedMax** | Pointer to **int64** | retained max number of backups | [optional] 

## Methods

### NewDpBlockBackupPolicyUpdateReqPolicy

`func NewDpBlockBackupPolicyUpdateReqPolicy() *DpBlockBackupPolicyUpdateReqPolicy`

NewDpBlockBackupPolicyUpdateReqPolicy instantiates a new DpBlockBackupPolicyUpdateReqPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockBackupPolicyUpdateReqPolicyWithDefaults

`func NewDpBlockBackupPolicyUpdateReqPolicyWithDefaults() *DpBlockBackupPolicyUpdateReqPolicy`

NewDpBlockBackupPolicyUpdateReqPolicyWithDefaults instantiates a new DpBlockBackupPolicyUpdateReqPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncrypted

`func (o *DpBlockBackupPolicyUpdateReqPolicy) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *DpBlockBackupPolicyUpdateReqPolicy) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *DpBlockBackupPolicyUpdateReqPolicy) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *DpBlockBackupPolicyUpdateReqPolicy) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetFullCronExpr

`func (o *DpBlockBackupPolicyUpdateReqPolicy) GetFullCronExpr() string`

GetFullCronExpr returns the FullCronExpr field if non-nil, zero value otherwise.

### GetFullCronExprOk

`func (o *DpBlockBackupPolicyUpdateReqPolicy) GetFullCronExprOk() (*string, bool)`

GetFullCronExprOk returns a tuple with the FullCronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullCronExpr

`func (o *DpBlockBackupPolicyUpdateReqPolicy) SetFullCronExpr(v string)`

SetFullCronExpr sets FullCronExpr field to given value.

### HasFullCronExpr

`func (o *DpBlockBackupPolicyUpdateReqPolicy) HasFullCronExpr() bool`

HasFullCronExpr returns a boolean if a field has been set.

### GetIncCronExpr

`func (o *DpBlockBackupPolicyUpdateReqPolicy) GetIncCronExpr() string`

GetIncCronExpr returns the IncCronExpr field if non-nil, zero value otherwise.

### GetIncCronExprOk

`func (o *DpBlockBackupPolicyUpdateReqPolicy) GetIncCronExprOk() (*string, bool)`

GetIncCronExprOk returns a tuple with the IncCronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncCronExpr

`func (o *DpBlockBackupPolicyUpdateReqPolicy) SetIncCronExpr(v string)`

SetIncCronExpr sets IncCronExpr field to given value.

### HasIncCronExpr

`func (o *DpBlockBackupPolicyUpdateReqPolicy) HasIncCronExpr() bool`

HasIncCronExpr returns a boolean if a field has been set.

### GetName

`func (o *DpBlockBackupPolicyUpdateReqPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpBlockBackupPolicyUpdateReqPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpBlockBackupPolicyUpdateReqPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DpBlockBackupPolicyUpdateReqPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetainedMax

`func (o *DpBlockBackupPolicyUpdateReqPolicy) GetRetainedMax() int64`

GetRetainedMax returns the RetainedMax field if non-nil, zero value otherwise.

### GetRetainedMaxOk

`func (o *DpBlockBackupPolicyUpdateReqPolicy) GetRetainedMaxOk() (*int64, bool)`

GetRetainedMaxOk returns a tuple with the RetainedMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedMax

`func (o *DpBlockBackupPolicyUpdateReqPolicy) SetRetainedMax(v int64)`

SetRetainedMax sets RetainedMax field to given value.

### HasRetainedMax

`func (o *DpBlockBackupPolicyUpdateReqPolicy) HasRetainedMax() bool`

HasRetainedMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


