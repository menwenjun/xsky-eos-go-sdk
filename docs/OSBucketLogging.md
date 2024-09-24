# OSBucketLogging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**SourceBucketName** | Pointer to **string** |  | [optional] 
**TargetBucketName** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOSBucketLogging

`func NewOSBucketLogging() *OSBucketLogging`

NewOSBucketLogging instantiates a new OSBucketLogging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSBucketLoggingWithDefaults

`func NewOSBucketLoggingWithDefaults() *OSBucketLogging`

NewOSBucketLoggingWithDefaults instantiates a new OSBucketLogging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *OSBucketLogging) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSBucketLogging) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSBucketLogging) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSBucketLogging) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OSBucketLogging) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSBucketLogging) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSBucketLogging) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSBucketLogging) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *OSBucketLogging) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSBucketLogging) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSBucketLogging) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OSBucketLogging) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrefix

`func (o *OSBucketLogging) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *OSBucketLogging) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *OSBucketLogging) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *OSBucketLogging) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetSourceBucketName

`func (o *OSBucketLogging) GetSourceBucketName() string`

GetSourceBucketName returns the SourceBucketName field if non-nil, zero value otherwise.

### GetSourceBucketNameOk

`func (o *OSBucketLogging) GetSourceBucketNameOk() (*string, bool)`

GetSourceBucketNameOk returns a tuple with the SourceBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBucketName

`func (o *OSBucketLogging) SetSourceBucketName(v string)`

SetSourceBucketName sets SourceBucketName field to given value.

### HasSourceBucketName

`func (o *OSBucketLogging) HasSourceBucketName() bool`

HasSourceBucketName returns a boolean if a field has been set.

### GetTargetBucketName

`func (o *OSBucketLogging) GetTargetBucketName() string`

GetTargetBucketName returns the TargetBucketName field if non-nil, zero value otherwise.

### GetTargetBucketNameOk

`func (o *OSBucketLogging) GetTargetBucketNameOk() (*string, bool)`

GetTargetBucketNameOk returns a tuple with the TargetBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBucketName

`func (o *OSBucketLogging) SetTargetBucketName(v string)`

SetTargetBucketName sets TargetBucketName field to given value.

### HasTargetBucketName

`func (o *OSBucketLogging) HasTargetBucketName() bool`

HasTargetBucketName returns a boolean if a field has been set.

### GetUpdate

`func (o *OSBucketLogging) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OSBucketLogging) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OSBucketLogging) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OSBucketLogging) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


