# Alert

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

## Methods

### NewAlert

`func NewAlert() *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckTime

`func (o *Alert) GetAckTime() time.Time`

GetAckTime returns the AckTime field if non-nil, zero value otherwise.

### GetAckTimeOk

`func (o *Alert) GetAckTimeOk() (*time.Time, bool)`

GetAckTimeOk returns a tuple with the AckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckTime

`func (o *Alert) SetAckTime(v time.Time)`

SetAckTime sets AckTime field to given value.

### HasAckTime

`func (o *Alert) HasAckTime() bool`

HasAckTime returns a boolean if a field has been set.

### GetAcked

`func (o *Alert) GetAcked() bool`

GetAcked returns the Acked field if non-nil, zero value otherwise.

### GetAckedOk

`func (o *Alert) GetAckedOk() (*bool, bool)`

GetAckedOk returns a tuple with the Acked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcked

`func (o *Alert) SetAcked(v bool)`

SetAcked sets Acked field to given value.

### HasAcked

`func (o *Alert) HasAcked() bool`

HasAcked returns a boolean if a field has been set.

### GetAdvice

`func (o *Alert) GetAdvice() ObjectInfo`

GetAdvice returns the Advice field if non-nil, zero value otherwise.

### GetAdviceOk

`func (o *Alert) GetAdviceOk() (*ObjectInfo, bool)`

GetAdviceOk returns a tuple with the Advice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvice

`func (o *Alert) SetAdvice(v ObjectInfo)`

SetAdvice sets Advice field to given value.

### HasAdvice

`func (o *Alert) HasAdvice() bool`

HasAdvice returns a boolean if a field has been set.

### GetAlarmId

`func (o *Alert) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *Alert) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *Alert) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.

### HasAlarmId

`func (o *Alert) HasAlarmId() bool`

HasAlarmId returns a boolean if a field has been set.

### GetAlertValue

`func (o *Alert) GetAlertValue() string`

GetAlertValue returns the AlertValue field if non-nil, zero value otherwise.

### GetAlertValueOk

`func (o *Alert) GetAlertValueOk() (*string, bool)`

GetAlertValueOk returns a tuple with the AlertValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertValue

`func (o *Alert) SetAlertValue(v string)`

SetAlertValue sets AlertValue field to given value.

### HasAlertValue

`func (o *Alert) HasAlertValue() bool`

HasAlertValue returns a boolean if a field has been set.

### GetCluster

`func (o *Alert) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Alert) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Alert) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Alert) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *Alert) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *Alert) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *Alert) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *Alert) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetData

`func (o *Alert) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Alert) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Alert) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *Alert) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *Alert) GetDescription() ObjectInfo`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Alert) GetDescriptionOk() (*ObjectInfo, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Alert) SetDescription(v ObjectInfo)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Alert) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExportStatus

`func (o *Alert) GetExportStatus() string`

GetExportStatus returns the ExportStatus field if non-nil, zero value otherwise.

### GetExportStatusOk

`func (o *Alert) GetExportStatusOk() (*string, bool)`

GetExportStatusOk returns a tuple with the ExportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportStatus

`func (o *Alert) SetExportStatus(v string)`

SetExportStatus sets ExportStatus field to given value.

### HasExportStatus

`func (o *Alert) HasExportStatus() bool`

HasExportStatus returns a boolean if a field has been set.

### GetExtraData

`func (o *Alert) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *Alert) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *Alert) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *Alert) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetGroup

`func (o *Alert) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Alert) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Alert) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Alert) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHost

`func (o *Alert) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Alert) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Alert) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *Alert) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *Alert) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Alert) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Alert) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Alert) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLevel

`func (o *Alert) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Alert) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Alert) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Alert) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetOspCluster

`func (o *Alert) GetOspCluster() ClusterNestview`

GetOspCluster returns the OspCluster field if non-nil, zero value otherwise.

### GetOspClusterOk

`func (o *Alert) GetOspClusterOk() (*ClusterNestview, bool)`

GetOspClusterOk returns a tuple with the OspCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspCluster

`func (o *Alert) SetOspCluster(v ClusterNestview)`

SetOspCluster sets OspCluster field to given value.

### HasOspCluster

`func (o *Alert) HasOspCluster() bool`

HasOspCluster returns a boolean if a field has been set.

### GetReason

`func (o *Alert) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Alert) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Alert) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Alert) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRelatedResources

`func (o *Alert) GetRelatedResources() []map[string]interface{}`

GetRelatedResources returns the RelatedResources field if non-nil, zero value otherwise.

### GetRelatedResourcesOk

`func (o *Alert) GetRelatedResourcesOk() (*[]map[string]interface{}, bool)`

GetRelatedResourcesOk returns a tuple with the RelatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResources

`func (o *Alert) SetRelatedResources(v []map[string]interface{})`

SetRelatedResources sets RelatedResources field to given value.

### HasRelatedResources

`func (o *Alert) HasRelatedResources() bool`

HasRelatedResources returns a boolean if a field has been set.

### GetResolveTime

`func (o *Alert) GetResolveTime() time.Time`

GetResolveTime returns the ResolveTime field if non-nil, zero value otherwise.

### GetResolveTimeOk

`func (o *Alert) GetResolveTimeOk() (*time.Time, bool)`

GetResolveTimeOk returns a tuple with the ResolveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveTime

`func (o *Alert) SetResolveTime(v time.Time)`

SetResolveTime sets ResolveTime field to given value.

### HasResolveTime

`func (o *Alert) HasResolveTime() bool`

HasResolveTime returns a boolean if a field has been set.

### GetResolveType

`func (o *Alert) GetResolveType() string`

GetResolveType returns the ResolveType field if non-nil, zero value otherwise.

### GetResolveTypeOk

`func (o *Alert) GetResolveTypeOk() (*string, bool)`

GetResolveTypeOk returns a tuple with the ResolveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveType

`func (o *Alert) SetResolveType(v string)`

SetResolveType sets ResolveType field to given value.

### HasResolveType

`func (o *Alert) HasResolveType() bool`

HasResolveType returns a boolean if a field has been set.

### GetResolved

`func (o *Alert) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *Alert) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *Alert) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *Alert) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetResourceId

`func (o *Alert) GetResourceId() int64`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Alert) GetResourceIdOk() (*int64, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Alert) SetResourceId(v int64)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *Alert) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceName

`func (o *Alert) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *Alert) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *Alert) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *Alert) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceType

`func (o *Alert) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Alert) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Alert) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *Alert) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStatus

`func (o *Alert) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Alert) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Alert) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Alert) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTriggerMode

`func (o *Alert) GetTriggerMode() string`

GetTriggerMode returns the TriggerMode field if non-nil, zero value otherwise.

### GetTriggerModeOk

`func (o *Alert) GetTriggerModeOk() (*string, bool)`

GetTriggerModeOk returns a tuple with the TriggerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMode

`func (o *Alert) SetTriggerMode(v string)`

SetTriggerMode sets TriggerMode field to given value.

### HasTriggerMode

`func (o *Alert) HasTriggerMode() bool`

HasTriggerMode returns a boolean if a field has been set.

### GetTriggerPeriod

`func (o *Alert) GetTriggerPeriod() int64`

GetTriggerPeriod returns the TriggerPeriod field if non-nil, zero value otherwise.

### GetTriggerPeriodOk

`func (o *Alert) GetTriggerPeriodOk() (*int64, bool)`

GetTriggerPeriodOk returns a tuple with the TriggerPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPeriod

`func (o *Alert) SetTriggerPeriod(v int64)`

SetTriggerPeriod sets TriggerPeriod field to given value.

### HasTriggerPeriod

`func (o *Alert) HasTriggerPeriod() bool`

HasTriggerPeriod returns a boolean if a field has been set.

### GetTriggerValue

`func (o *Alert) GetTriggerValue() string`

GetTriggerValue returns the TriggerValue field if non-nil, zero value otherwise.

### GetTriggerValueOk

`func (o *Alert) GetTriggerValueOk() (*string, bool)`

GetTriggerValueOk returns a tuple with the TriggerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerValue

`func (o *Alert) SetTriggerValue(v string)`

SetTriggerValue sets TriggerValue field to given value.

### HasTriggerValue

`func (o *Alert) HasTriggerValue() bool`

HasTriggerValue returns a boolean if a field has been set.

### GetType

`func (o *Alert) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Alert) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Alert) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Alert) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWeight

`func (o *Alert) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Alert) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Alert) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *Alert) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


