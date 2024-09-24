# OSExStoragePlatformValidator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**PlatformInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**Results** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewOSExStoragePlatformValidator

`func NewOSExStoragePlatformValidator() *OSExStoragePlatformValidator`

NewOSExStoragePlatformValidator instantiates a new OSExStoragePlatformValidator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSExStoragePlatformValidatorWithDefaults

`func NewOSExStoragePlatformValidatorWithDefaults() *OSExStoragePlatformValidator`

NewOSExStoragePlatformValidatorWithDefaults instantiates a new OSExStoragePlatformValidator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *OSExStoragePlatformValidator) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSExStoragePlatformValidator) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSExStoragePlatformValidator) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSExStoragePlatformValidator) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OSExStoragePlatformValidator) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSExStoragePlatformValidator) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSExStoragePlatformValidator) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSExStoragePlatformValidator) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *OSExStoragePlatformValidator) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSExStoragePlatformValidator) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSExStoragePlatformValidator) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OSExStoragePlatformValidator) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlatformInfo

`func (o *OSExStoragePlatformValidator) GetPlatformInfo() map[string]interface{}`

GetPlatformInfo returns the PlatformInfo field if non-nil, zero value otherwise.

### GetPlatformInfoOk

`func (o *OSExStoragePlatformValidator) GetPlatformInfoOk() (*map[string]interface{}, bool)`

GetPlatformInfoOk returns a tuple with the PlatformInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformInfo

`func (o *OSExStoragePlatformValidator) SetPlatformInfo(v map[string]interface{})`

SetPlatformInfo sets PlatformInfo field to given value.

### HasPlatformInfo

`func (o *OSExStoragePlatformValidator) HasPlatformInfo() bool`

HasPlatformInfo returns a boolean if a field has been set.

### GetResults

`func (o *OSExStoragePlatformValidator) GetResults() []map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *OSExStoragePlatformValidator) GetResultsOk() (*[]map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *OSExStoragePlatformValidator) SetResults(v []map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *OSExStoragePlatformValidator) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStatus

`func (o *OSExStoragePlatformValidator) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSExStoragePlatformValidator) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSExStoragePlatformValidator) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OSExStoragePlatformValidator) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


