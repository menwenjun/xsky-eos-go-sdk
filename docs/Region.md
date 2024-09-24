# Region

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**AccessUrl** | Pointer to **string** |  | [optional] 
**AdminNetwork** | Pointer to **string** | TODO(siming): removed in 602 | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IsLocal** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StatsReservedDays** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewRegion

`func NewRegion() *Region`

NewRegion instantiates a new Region object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionWithDefaults

`func NewRegionWithDefaults() *Region`

NewRegionWithDefaults instantiates a new Region object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *Region) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *Region) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *Region) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *Region) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAccessUrl

`func (o *Region) GetAccessUrl() string`

GetAccessUrl returns the AccessUrl field if non-nil, zero value otherwise.

### GetAccessUrlOk

`func (o *Region) GetAccessUrlOk() (*string, bool)`

GetAccessUrlOk returns a tuple with the AccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessUrl

`func (o *Region) SetAccessUrl(v string)`

SetAccessUrl sets AccessUrl field to given value.

### HasAccessUrl

`func (o *Region) HasAccessUrl() bool`

HasAccessUrl returns a boolean if a field has been set.

### GetAdminNetwork

`func (o *Region) GetAdminNetwork() string`

GetAdminNetwork returns the AdminNetwork field if non-nil, zero value otherwise.

### GetAdminNetworkOk

`func (o *Region) GetAdminNetworkOk() (*string, bool)`

GetAdminNetworkOk returns a tuple with the AdminNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminNetwork

`func (o *Region) SetAdminNetwork(v string)`

SetAdminNetwork sets AdminNetwork field to given value.

### HasAdminNetwork

`func (o *Region) HasAdminNetwork() bool`

HasAdminNetwork returns a boolean if a field has been set.

### GetCreate

`func (o *Region) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *Region) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *Region) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *Region) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *Region) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Region) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Region) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Region) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsLocal

`func (o *Region) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *Region) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *Region) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *Region) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetName

`func (o *Region) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Region) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Region) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Region) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatsReservedDays

`func (o *Region) GetStatsReservedDays() int64`

GetStatsReservedDays returns the StatsReservedDays field if non-nil, zero value otherwise.

### GetStatsReservedDaysOk

`func (o *Region) GetStatsReservedDaysOk() (*int64, bool)`

GetStatsReservedDaysOk returns a tuple with the StatsReservedDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsReservedDays

`func (o *Region) SetStatsReservedDays(v int64)`

SetStatsReservedDays sets StatsReservedDays field to given value.

### HasStatsReservedDays

`func (o *Region) HasStatsReservedDays() bool`

HasStatsReservedDays returns a boolean if a field has been set.

### GetStatus

`func (o *Region) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Region) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Region) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Region) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *Region) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *Region) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *Region) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *Region) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUuid

`func (o *Region) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Region) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Region) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Region) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


