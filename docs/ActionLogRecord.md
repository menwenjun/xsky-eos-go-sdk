# ActionLogRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**ClientIp** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Finish** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**NewData** | Pointer to **map[string]interface{}** |  | [optional] 
**OldData** | Pointer to **map[string]interface{}** |  | [optional] 
**OspCluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Parameter** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to [**ActionLogNestview**](ActionLogNestview.md) |  | [optional] 
**RelatedResources** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ResourceId** | Pointer to **int64** |  | [optional] 
**ResourceKey** | Pointer to **string** | key for resources not managed by us, can be used as resource id | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**UserNestview**](UserNestview.md) |  | [optional] 
**ErrorRecords** | Pointer to [**[]ErrorRecord**](ErrorRecord.md) |  | [optional] 

## Methods

### NewActionLogRecord

`func NewActionLogRecord() *ActionLogRecord`

NewActionLogRecord instantiates a new ActionLogRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionLogRecordWithDefaults

`func NewActionLogRecordWithDefaults() *ActionLogRecord`

NewActionLogRecordWithDefaults instantiates a new ActionLogRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ActionLogRecord) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ActionLogRecord) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ActionLogRecord) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ActionLogRecord) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetClientIp

`func (o *ActionLogRecord) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *ActionLogRecord) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *ActionLogRecord) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *ActionLogRecord) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetCluster

`func (o *ActionLogRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ActionLogRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ActionLogRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ActionLogRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *ActionLogRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ActionLogRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ActionLogRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ActionLogRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetFinish

`func (o *ActionLogRecord) GetFinish() time.Time`

GetFinish returns the Finish field if non-nil, zero value otherwise.

### GetFinishOk

`func (o *ActionLogRecord) GetFinishOk() (*time.Time, bool)`

GetFinishOk returns a tuple with the Finish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinish

`func (o *ActionLogRecord) SetFinish(v time.Time)`

SetFinish sets Finish field to given value.

### HasFinish

`func (o *ActionLogRecord) HasFinish() bool`

HasFinish returns a boolean if a field has been set.

### GetId

`func (o *ActionLogRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionLogRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionLogRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ActionLogRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *ActionLogRecord) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ActionLogRecord) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ActionLogRecord) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ActionLogRecord) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNewData

`func (o *ActionLogRecord) GetNewData() map[string]interface{}`

GetNewData returns the NewData field if non-nil, zero value otherwise.

### GetNewDataOk

`func (o *ActionLogRecord) GetNewDataOk() (*map[string]interface{}, bool)`

GetNewDataOk returns a tuple with the NewData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewData

`func (o *ActionLogRecord) SetNewData(v map[string]interface{})`

SetNewData sets NewData field to given value.

### HasNewData

`func (o *ActionLogRecord) HasNewData() bool`

HasNewData returns a boolean if a field has been set.

### GetOldData

`func (o *ActionLogRecord) GetOldData() map[string]interface{}`

GetOldData returns the OldData field if non-nil, zero value otherwise.

### GetOldDataOk

`func (o *ActionLogRecord) GetOldDataOk() (*map[string]interface{}, bool)`

GetOldDataOk returns a tuple with the OldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldData

`func (o *ActionLogRecord) SetOldData(v map[string]interface{})`

SetOldData sets OldData field to given value.

### HasOldData

`func (o *ActionLogRecord) HasOldData() bool`

HasOldData returns a boolean if a field has been set.

### GetOspCluster

`func (o *ActionLogRecord) GetOspCluster() ClusterNestview`

GetOspCluster returns the OspCluster field if non-nil, zero value otherwise.

### GetOspClusterOk

`func (o *ActionLogRecord) GetOspClusterOk() (*ClusterNestview, bool)`

GetOspClusterOk returns a tuple with the OspCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspCluster

`func (o *ActionLogRecord) SetOspCluster(v ClusterNestview)`

SetOspCluster sets OspCluster field to given value.

### HasOspCluster

`func (o *ActionLogRecord) HasOspCluster() bool`

HasOspCluster returns a boolean if a field has been set.

### GetParameter

`func (o *ActionLogRecord) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ActionLogRecord) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ActionLogRecord) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *ActionLogRecord) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetParent

`func (o *ActionLogRecord) GetParent() ActionLogNestview`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ActionLogRecord) GetParentOk() (*ActionLogNestview, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ActionLogRecord) SetParent(v ActionLogNestview)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ActionLogRecord) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetRelatedResources

`func (o *ActionLogRecord) GetRelatedResources() []map[string]interface{}`

GetRelatedResources returns the RelatedResources field if non-nil, zero value otherwise.

### GetRelatedResourcesOk

`func (o *ActionLogRecord) GetRelatedResourcesOk() (*[]map[string]interface{}, bool)`

GetRelatedResourcesOk returns a tuple with the RelatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResources

`func (o *ActionLogRecord) SetRelatedResources(v []map[string]interface{})`

SetRelatedResources sets RelatedResources field to given value.

### HasRelatedResources

`func (o *ActionLogRecord) HasRelatedResources() bool`

HasRelatedResources returns a boolean if a field has been set.

### GetResourceId

`func (o *ActionLogRecord) GetResourceId() int64`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ActionLogRecord) GetResourceIdOk() (*int64, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ActionLogRecord) SetResourceId(v int64)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ActionLogRecord) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceKey

`func (o *ActionLogRecord) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *ActionLogRecord) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *ActionLogRecord) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.

### HasResourceKey

`func (o *ActionLogRecord) HasResourceKey() bool`

HasResourceKey returns a boolean if a field has been set.

### GetResourceType

`func (o *ActionLogRecord) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ActionLogRecord) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ActionLogRecord) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ActionLogRecord) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStart

`func (o *ActionLogRecord) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ActionLogRecord) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ActionLogRecord) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *ActionLogRecord) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStatus

`func (o *ActionLogRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionLogRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionLogRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActionLogRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUser

`func (o *ActionLogRecord) GetUser() UserNestview`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ActionLogRecord) GetUserOk() (*UserNestview, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ActionLogRecord) SetUser(v UserNestview)`

SetUser sets User field to given value.

### HasUser

`func (o *ActionLogRecord) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetErrorRecords

`func (o *ActionLogRecord) GetErrorRecords() []ErrorRecord`

GetErrorRecords returns the ErrorRecords field if non-nil, zero value otherwise.

### GetErrorRecordsOk

`func (o *ActionLogRecord) GetErrorRecordsOk() (*[]ErrorRecord, bool)`

GetErrorRecordsOk returns a tuple with the ErrorRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRecords

`func (o *ActionLogRecord) SetErrorRecords(v []ErrorRecord)`

SetErrorRecords sets ErrorRecords field to given value.

### HasErrorRecords

`func (o *ActionLogRecord) HasErrorRecords() bool`

HasErrorRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


