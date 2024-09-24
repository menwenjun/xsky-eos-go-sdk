# AlertRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckTime** | Pointer to **time.Time** |  | [optional] 
**Acked** | Pointer to **bool** |  | [optional] 
**Advice** | Pointer to [**ObjectInfo**](ObjectInfo.md) |  | [optional] 
**AlarmId** | Pointer to **string** |  | [optional] 
**AlertValue** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to [**ObjectInfo**](ObjectInfo.md) |  | [optional] 
**ExportStatus** | Pointer to **string** |  | [optional] 
**ExtraData** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**OspCluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**RelatedResources** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ResolveTime** | Pointer to **time.Time** |  | [optional] 
**ResolveType** | Pointer to **string** |  | [optional] 
**Resolved** | Pointer to **bool** |  | [optional] 
**ResourceId** | Pointer to **int64** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** | resource name | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TriggerMode** | Pointer to **string** |  | [optional] 
**TriggerPeriod** | Pointer to **int64** |  | [optional] 
**TriggerValue** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**ErrorRecords** | Pointer to [**[]ErrorRecord**](ErrorRecord.md) |  | [optional] 

## Methods

### NewAlertRecord

`func NewAlertRecord() *AlertRecord`

NewAlertRecord instantiates a new AlertRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRecordWithDefaults

`func NewAlertRecordWithDefaults() *AlertRecord`

NewAlertRecordWithDefaults instantiates a new AlertRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckTime

`func (o *AlertRecord) GetAckTime() time.Time`

GetAckTime returns the AckTime field if non-nil, zero value otherwise.

### GetAckTimeOk

`func (o *AlertRecord) GetAckTimeOk() (*time.Time, bool)`

GetAckTimeOk returns a tuple with the AckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckTime

`func (o *AlertRecord) SetAckTime(v time.Time)`

SetAckTime sets AckTime field to given value.

### HasAckTime

`func (o *AlertRecord) HasAckTime() bool`

HasAckTime returns a boolean if a field has been set.

### GetAcked

`func (o *AlertRecord) GetAcked() bool`

GetAcked returns the Acked field if non-nil, zero value otherwise.

### GetAckedOk

`func (o *AlertRecord) GetAckedOk() (*bool, bool)`

GetAckedOk returns a tuple with the Acked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcked

`func (o *AlertRecord) SetAcked(v bool)`

SetAcked sets Acked field to given value.

### HasAcked

`func (o *AlertRecord) HasAcked() bool`

HasAcked returns a boolean if a field has been set.

### GetAdvice

`func (o *AlertRecord) GetAdvice() ObjectInfo`

GetAdvice returns the Advice field if non-nil, zero value otherwise.

### GetAdviceOk

`func (o *AlertRecord) GetAdviceOk() (*ObjectInfo, bool)`

GetAdviceOk returns a tuple with the Advice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvice

`func (o *AlertRecord) SetAdvice(v ObjectInfo)`

SetAdvice sets Advice field to given value.

### HasAdvice

`func (o *AlertRecord) HasAdvice() bool`

HasAdvice returns a boolean if a field has been set.

### GetAlarmId

`func (o *AlertRecord) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *AlertRecord) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *AlertRecord) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.

### HasAlarmId

`func (o *AlertRecord) HasAlarmId() bool`

HasAlarmId returns a boolean if a field has been set.

### GetAlertValue

`func (o *AlertRecord) GetAlertValue() string`

GetAlertValue returns the AlertValue field if non-nil, zero value otherwise.

### GetAlertValueOk

`func (o *AlertRecord) GetAlertValueOk() (*string, bool)`

GetAlertValueOk returns a tuple with the AlertValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertValue

`func (o *AlertRecord) SetAlertValue(v string)`

SetAlertValue sets AlertValue field to given value.

### HasAlertValue

`func (o *AlertRecord) HasAlertValue() bool`

HasAlertValue returns a boolean if a field has been set.

### GetCluster

`func (o *AlertRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *AlertRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *AlertRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *AlertRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *AlertRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *AlertRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *AlertRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *AlertRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetData

`func (o *AlertRecord) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlertRecord) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlertRecord) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *AlertRecord) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *AlertRecord) GetDescription() ObjectInfo`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertRecord) GetDescriptionOk() (*ObjectInfo, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertRecord) SetDescription(v ObjectInfo)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AlertRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExportStatus

`func (o *AlertRecord) GetExportStatus() string`

GetExportStatus returns the ExportStatus field if non-nil, zero value otherwise.

### GetExportStatusOk

`func (o *AlertRecord) GetExportStatusOk() (*string, bool)`

GetExportStatusOk returns a tuple with the ExportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportStatus

`func (o *AlertRecord) SetExportStatus(v string)`

SetExportStatus sets ExportStatus field to given value.

### HasExportStatus

`func (o *AlertRecord) HasExportStatus() bool`

HasExportStatus returns a boolean if a field has been set.

### GetExtraData

`func (o *AlertRecord) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *AlertRecord) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *AlertRecord) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *AlertRecord) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetGroup

`func (o *AlertRecord) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AlertRecord) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AlertRecord) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AlertRecord) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHost

`func (o *AlertRecord) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AlertRecord) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AlertRecord) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *AlertRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *AlertRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AlertRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLevel

`func (o *AlertRecord) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AlertRecord) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AlertRecord) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AlertRecord) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetOspCluster

`func (o *AlertRecord) GetOspCluster() ClusterNestview`

GetOspCluster returns the OspCluster field if non-nil, zero value otherwise.

### GetOspClusterOk

`func (o *AlertRecord) GetOspClusterOk() (*ClusterNestview, bool)`

GetOspClusterOk returns a tuple with the OspCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspCluster

`func (o *AlertRecord) SetOspCluster(v ClusterNestview)`

SetOspCluster sets OspCluster field to given value.

### HasOspCluster

`func (o *AlertRecord) HasOspCluster() bool`

HasOspCluster returns a boolean if a field has been set.

### GetReason

`func (o *AlertRecord) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AlertRecord) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AlertRecord) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AlertRecord) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRelatedResources

`func (o *AlertRecord) GetRelatedResources() []map[string]interface{}`

GetRelatedResources returns the RelatedResources field if non-nil, zero value otherwise.

### GetRelatedResourcesOk

`func (o *AlertRecord) GetRelatedResourcesOk() (*[]map[string]interface{}, bool)`

GetRelatedResourcesOk returns a tuple with the RelatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResources

`func (o *AlertRecord) SetRelatedResources(v []map[string]interface{})`

SetRelatedResources sets RelatedResources field to given value.

### HasRelatedResources

`func (o *AlertRecord) HasRelatedResources() bool`

HasRelatedResources returns a boolean if a field has been set.

### GetResolveTime

`func (o *AlertRecord) GetResolveTime() time.Time`

GetResolveTime returns the ResolveTime field if non-nil, zero value otherwise.

### GetResolveTimeOk

`func (o *AlertRecord) GetResolveTimeOk() (*time.Time, bool)`

GetResolveTimeOk returns a tuple with the ResolveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveTime

`func (o *AlertRecord) SetResolveTime(v time.Time)`

SetResolveTime sets ResolveTime field to given value.

### HasResolveTime

`func (o *AlertRecord) HasResolveTime() bool`

HasResolveTime returns a boolean if a field has been set.

### GetResolveType

`func (o *AlertRecord) GetResolveType() string`

GetResolveType returns the ResolveType field if non-nil, zero value otherwise.

### GetResolveTypeOk

`func (o *AlertRecord) GetResolveTypeOk() (*string, bool)`

GetResolveTypeOk returns a tuple with the ResolveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveType

`func (o *AlertRecord) SetResolveType(v string)`

SetResolveType sets ResolveType field to given value.

### HasResolveType

`func (o *AlertRecord) HasResolveType() bool`

HasResolveType returns a boolean if a field has been set.

### GetResolved

`func (o *AlertRecord) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *AlertRecord) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *AlertRecord) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *AlertRecord) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetResourceId

`func (o *AlertRecord) GetResourceId() int64`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AlertRecord) GetResourceIdOk() (*int64, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AlertRecord) SetResourceId(v int64)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AlertRecord) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceName

`func (o *AlertRecord) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AlertRecord) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AlertRecord) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AlertRecord) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceType

`func (o *AlertRecord) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AlertRecord) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AlertRecord) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AlertRecord) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStatus

`func (o *AlertRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTriggerMode

`func (o *AlertRecord) GetTriggerMode() string`

GetTriggerMode returns the TriggerMode field if non-nil, zero value otherwise.

### GetTriggerModeOk

`func (o *AlertRecord) GetTriggerModeOk() (*string, bool)`

GetTriggerModeOk returns a tuple with the TriggerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMode

`func (o *AlertRecord) SetTriggerMode(v string)`

SetTriggerMode sets TriggerMode field to given value.

### HasTriggerMode

`func (o *AlertRecord) HasTriggerMode() bool`

HasTriggerMode returns a boolean if a field has been set.

### GetTriggerPeriod

`func (o *AlertRecord) GetTriggerPeriod() int64`

GetTriggerPeriod returns the TriggerPeriod field if non-nil, zero value otherwise.

### GetTriggerPeriodOk

`func (o *AlertRecord) GetTriggerPeriodOk() (*int64, bool)`

GetTriggerPeriodOk returns a tuple with the TriggerPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPeriod

`func (o *AlertRecord) SetTriggerPeriod(v int64)`

SetTriggerPeriod sets TriggerPeriod field to given value.

### HasTriggerPeriod

`func (o *AlertRecord) HasTriggerPeriod() bool`

HasTriggerPeriod returns a boolean if a field has been set.

### GetTriggerValue

`func (o *AlertRecord) GetTriggerValue() string`

GetTriggerValue returns the TriggerValue field if non-nil, zero value otherwise.

### GetTriggerValueOk

`func (o *AlertRecord) GetTriggerValueOk() (*string, bool)`

GetTriggerValueOk returns a tuple with the TriggerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerValue

`func (o *AlertRecord) SetTriggerValue(v string)`

SetTriggerValue sets TriggerValue field to given value.

### HasTriggerValue

`func (o *AlertRecord) HasTriggerValue() bool`

HasTriggerValue returns a boolean if a field has been set.

### GetType

`func (o *AlertRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AlertRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWeight

`func (o *AlertRecord) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *AlertRecord) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *AlertRecord) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *AlertRecord) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetErrorRecords

`func (o *AlertRecord) GetErrorRecords() []ErrorRecord`

GetErrorRecords returns the ErrorRecords field if non-nil, zero value otherwise.

### GetErrorRecordsOk

`func (o *AlertRecord) GetErrorRecordsOk() (*[]ErrorRecord, bool)`

GetErrorRecordsOk returns a tuple with the ErrorRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRecords

`func (o *AlertRecord) SetErrorRecords(v []ErrorRecord)`

SetErrorRecords sets ErrorRecords field to given value.

### HasErrorRecords

`func (o *AlertRecord) HasErrorRecords() bool`

HasErrorRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


