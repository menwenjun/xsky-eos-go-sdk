# AlertInfo

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
**Object** | Pointer to [**ObjectInfo**](ObjectInfo.md) |  | [optional] 
**OspCluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**RelatedResources** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ResolveTime** | Pointer to **time.Time** |  | [optional] 
**ResolveType** | Pointer to **string** |  | [optional] 
**Resolved** | Pointer to **bool** |  | [optional] 
**ResourceId** | Pointer to **int64** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TriggerMode** | Pointer to **string** |  | [optional] 
**TriggerPeriod** | Pointer to **int64** |  | [optional] 
**TriggerValue** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 

## Methods

### NewAlertInfo

`func NewAlertInfo() *AlertInfo`

NewAlertInfo instantiates a new AlertInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertInfoWithDefaults

`func NewAlertInfoWithDefaults() *AlertInfo`

NewAlertInfoWithDefaults instantiates a new AlertInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckTime

`func (o *AlertInfo) GetAckTime() time.Time`

GetAckTime returns the AckTime field if non-nil, zero value otherwise.

### GetAckTimeOk

`func (o *AlertInfo) GetAckTimeOk() (*time.Time, bool)`

GetAckTimeOk returns a tuple with the AckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckTime

`func (o *AlertInfo) SetAckTime(v time.Time)`

SetAckTime sets AckTime field to given value.

### HasAckTime

`func (o *AlertInfo) HasAckTime() bool`

HasAckTime returns a boolean if a field has been set.

### GetAcked

`func (o *AlertInfo) GetAcked() bool`

GetAcked returns the Acked field if non-nil, zero value otherwise.

### GetAckedOk

`func (o *AlertInfo) GetAckedOk() (*bool, bool)`

GetAckedOk returns a tuple with the Acked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcked

`func (o *AlertInfo) SetAcked(v bool)`

SetAcked sets Acked field to given value.

### HasAcked

`func (o *AlertInfo) HasAcked() bool`

HasAcked returns a boolean if a field has been set.

### GetAdvice

`func (o *AlertInfo) GetAdvice() ObjectInfo`

GetAdvice returns the Advice field if non-nil, zero value otherwise.

### GetAdviceOk

`func (o *AlertInfo) GetAdviceOk() (*ObjectInfo, bool)`

GetAdviceOk returns a tuple with the Advice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvice

`func (o *AlertInfo) SetAdvice(v ObjectInfo)`

SetAdvice sets Advice field to given value.

### HasAdvice

`func (o *AlertInfo) HasAdvice() bool`

HasAdvice returns a boolean if a field has been set.

### GetAlarmId

`func (o *AlertInfo) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *AlertInfo) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *AlertInfo) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.

### HasAlarmId

`func (o *AlertInfo) HasAlarmId() bool`

HasAlarmId returns a boolean if a field has been set.

### GetAlertValue

`func (o *AlertInfo) GetAlertValue() string`

GetAlertValue returns the AlertValue field if non-nil, zero value otherwise.

### GetAlertValueOk

`func (o *AlertInfo) GetAlertValueOk() (*string, bool)`

GetAlertValueOk returns a tuple with the AlertValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertValue

`func (o *AlertInfo) SetAlertValue(v string)`

SetAlertValue sets AlertValue field to given value.

### HasAlertValue

`func (o *AlertInfo) HasAlertValue() bool`

HasAlertValue returns a boolean if a field has been set.

### GetCluster

`func (o *AlertInfo) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *AlertInfo) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *AlertInfo) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *AlertInfo) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *AlertInfo) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *AlertInfo) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *AlertInfo) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *AlertInfo) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetData

`func (o *AlertInfo) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlertInfo) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlertInfo) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *AlertInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *AlertInfo) GetDescription() ObjectInfo`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertInfo) GetDescriptionOk() (*ObjectInfo, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertInfo) SetDescription(v ObjectInfo)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AlertInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExportStatus

`func (o *AlertInfo) GetExportStatus() string`

GetExportStatus returns the ExportStatus field if non-nil, zero value otherwise.

### GetExportStatusOk

`func (o *AlertInfo) GetExportStatusOk() (*string, bool)`

GetExportStatusOk returns a tuple with the ExportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportStatus

`func (o *AlertInfo) SetExportStatus(v string)`

SetExportStatus sets ExportStatus field to given value.

### HasExportStatus

`func (o *AlertInfo) HasExportStatus() bool`

HasExportStatus returns a boolean if a field has been set.

### GetExtraData

`func (o *AlertInfo) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *AlertInfo) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *AlertInfo) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *AlertInfo) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetGroup

`func (o *AlertInfo) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AlertInfo) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AlertInfo) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AlertInfo) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHost

`func (o *AlertInfo) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AlertInfo) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AlertInfo) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *AlertInfo) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *AlertInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AlertInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLevel

`func (o *AlertInfo) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AlertInfo) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AlertInfo) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AlertInfo) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetObject

`func (o *AlertInfo) GetObject() ObjectInfo`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AlertInfo) GetObjectOk() (*ObjectInfo, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AlertInfo) SetObject(v ObjectInfo)`

SetObject sets Object field to given value.

### HasObject

`func (o *AlertInfo) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOspCluster

`func (o *AlertInfo) GetOspCluster() ClusterNestview`

GetOspCluster returns the OspCluster field if non-nil, zero value otherwise.

### GetOspClusterOk

`func (o *AlertInfo) GetOspClusterOk() (*ClusterNestview, bool)`

GetOspClusterOk returns a tuple with the OspCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspCluster

`func (o *AlertInfo) SetOspCluster(v ClusterNestview)`

SetOspCluster sets OspCluster field to given value.

### HasOspCluster

`func (o *AlertInfo) HasOspCluster() bool`

HasOspCluster returns a boolean if a field has been set.

### GetReason

`func (o *AlertInfo) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AlertInfo) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AlertInfo) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AlertInfo) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRelatedResources

`func (o *AlertInfo) GetRelatedResources() []map[string]interface{}`

GetRelatedResources returns the RelatedResources field if non-nil, zero value otherwise.

### GetRelatedResourcesOk

`func (o *AlertInfo) GetRelatedResourcesOk() (*[]map[string]interface{}, bool)`

GetRelatedResourcesOk returns a tuple with the RelatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResources

`func (o *AlertInfo) SetRelatedResources(v []map[string]interface{})`

SetRelatedResources sets RelatedResources field to given value.

### HasRelatedResources

`func (o *AlertInfo) HasRelatedResources() bool`

HasRelatedResources returns a boolean if a field has been set.

### GetResolveTime

`func (o *AlertInfo) GetResolveTime() time.Time`

GetResolveTime returns the ResolveTime field if non-nil, zero value otherwise.

### GetResolveTimeOk

`func (o *AlertInfo) GetResolveTimeOk() (*time.Time, bool)`

GetResolveTimeOk returns a tuple with the ResolveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveTime

`func (o *AlertInfo) SetResolveTime(v time.Time)`

SetResolveTime sets ResolveTime field to given value.

### HasResolveTime

`func (o *AlertInfo) HasResolveTime() bool`

HasResolveTime returns a boolean if a field has been set.

### GetResolveType

`func (o *AlertInfo) GetResolveType() string`

GetResolveType returns the ResolveType field if non-nil, zero value otherwise.

### GetResolveTypeOk

`func (o *AlertInfo) GetResolveTypeOk() (*string, bool)`

GetResolveTypeOk returns a tuple with the ResolveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveType

`func (o *AlertInfo) SetResolveType(v string)`

SetResolveType sets ResolveType field to given value.

### HasResolveType

`func (o *AlertInfo) HasResolveType() bool`

HasResolveType returns a boolean if a field has been set.

### GetResolved

`func (o *AlertInfo) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *AlertInfo) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *AlertInfo) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *AlertInfo) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetResourceId

`func (o *AlertInfo) GetResourceId() int64`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AlertInfo) GetResourceIdOk() (*int64, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AlertInfo) SetResourceId(v int64)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AlertInfo) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceName

`func (o *AlertInfo) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AlertInfo) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AlertInfo) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AlertInfo) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceType

`func (o *AlertInfo) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AlertInfo) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AlertInfo) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AlertInfo) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStatus

`func (o *AlertInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTriggerMode

`func (o *AlertInfo) GetTriggerMode() string`

GetTriggerMode returns the TriggerMode field if non-nil, zero value otherwise.

### GetTriggerModeOk

`func (o *AlertInfo) GetTriggerModeOk() (*string, bool)`

GetTriggerModeOk returns a tuple with the TriggerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMode

`func (o *AlertInfo) SetTriggerMode(v string)`

SetTriggerMode sets TriggerMode field to given value.

### HasTriggerMode

`func (o *AlertInfo) HasTriggerMode() bool`

HasTriggerMode returns a boolean if a field has been set.

### GetTriggerPeriod

`func (o *AlertInfo) GetTriggerPeriod() int64`

GetTriggerPeriod returns the TriggerPeriod field if non-nil, zero value otherwise.

### GetTriggerPeriodOk

`func (o *AlertInfo) GetTriggerPeriodOk() (*int64, bool)`

GetTriggerPeriodOk returns a tuple with the TriggerPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPeriod

`func (o *AlertInfo) SetTriggerPeriod(v int64)`

SetTriggerPeriod sets TriggerPeriod field to given value.

### HasTriggerPeriod

`func (o *AlertInfo) HasTriggerPeriod() bool`

HasTriggerPeriod returns a boolean if a field has been set.

### GetTriggerValue

`func (o *AlertInfo) GetTriggerValue() string`

GetTriggerValue returns the TriggerValue field if non-nil, zero value otherwise.

### GetTriggerValueOk

`func (o *AlertInfo) GetTriggerValueOk() (*string, bool)`

GetTriggerValueOk returns a tuple with the TriggerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerValue

`func (o *AlertInfo) SetTriggerValue(v string)`

SetTriggerValue sets TriggerValue field to given value.

### HasTriggerValue

`func (o *AlertInfo) HasTriggerValue() bool`

HasTriggerValue returns a boolean if a field has been set.

### GetType

`func (o *AlertInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AlertInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWeight

`func (o *AlertInfo) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *AlertInfo) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *AlertInfo) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *AlertInfo) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


