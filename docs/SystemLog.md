# SystemLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalog** | Pointer to **string** |  | [optional] 
**Host** | Pointer to [**Host**](Host.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 

## Methods

### NewSystemLog

`func NewSystemLog() *SystemLog`

NewSystemLog instantiates a new SystemLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemLogWithDefaults

`func NewSystemLogWithDefaults() *SystemLog`

NewSystemLogWithDefaults instantiates a new SystemLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalog

`func (o *SystemLog) GetCatalog() string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SystemLog) GetCatalogOk() (*string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SystemLog) SetCatalog(v string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SystemLog) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetHost

`func (o *SystemLog) GetHost() Host`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SystemLog) GetHostOk() (*Host, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SystemLog) SetHost(v Host)`

SetHost sets Host field to given value.

### HasHost

`func (o *SystemLog) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *SystemLog) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemLog) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemLog) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SystemLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *SystemLog) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *SystemLog) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *SystemLog) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *SystemLog) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetName

`func (o *SystemLog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemLog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemLog) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemLog) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *SystemLog) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SystemLog) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SystemLog) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *SystemLog) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


