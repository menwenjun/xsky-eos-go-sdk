# OSExternalStoragePlatformInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** |  | [optional] 
**Authentication** | Pointer to **string** |  | [optional] 
**BluRayId** | Pointer to **int64** |  | [optional] 
**Buckets** | Pointer to [**[]OSExternalStorageBucketInfo**](OSExternalStorageBucketInfo.md) |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**HostStyle** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int64** |  | [optional] 

## Methods

### NewOSExternalStoragePlatformInfo

`func NewOSExternalStoragePlatformInfo() *OSExternalStoragePlatformInfo`

NewOSExternalStoragePlatformInfo instantiates a new OSExternalStoragePlatformInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSExternalStoragePlatformInfoWithDefaults

`func NewOSExternalStoragePlatformInfoWithDefaults() *OSExternalStoragePlatformInfo`

NewOSExternalStoragePlatformInfoWithDefaults instantiates a new OSExternalStoragePlatformInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *OSExternalStoragePlatformInfo) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *OSExternalStoragePlatformInfo) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *OSExternalStoragePlatformInfo) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *OSExternalStoragePlatformInfo) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAuthentication

`func (o *OSExternalStoragePlatformInfo) GetAuthentication() string`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *OSExternalStoragePlatformInfo) GetAuthenticationOk() (*string, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *OSExternalStoragePlatformInfo) SetAuthentication(v string)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *OSExternalStoragePlatformInfo) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetBluRayId

`func (o *OSExternalStoragePlatformInfo) GetBluRayId() int64`

GetBluRayId returns the BluRayId field if non-nil, zero value otherwise.

### GetBluRayIdOk

`func (o *OSExternalStoragePlatformInfo) GetBluRayIdOk() (*int64, bool)`

GetBluRayIdOk returns a tuple with the BluRayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluRayId

`func (o *OSExternalStoragePlatformInfo) SetBluRayId(v int64)`

SetBluRayId sets BluRayId field to given value.

### HasBluRayId

`func (o *OSExternalStoragePlatformInfo) HasBluRayId() bool`

HasBluRayId returns a boolean if a field has been set.

### GetBuckets

`func (o *OSExternalStoragePlatformInfo) GetBuckets() []OSExternalStorageBucketInfo`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *OSExternalStoragePlatformInfo) GetBucketsOk() (*[]OSExternalStorageBucketInfo, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *OSExternalStoragePlatformInfo) SetBuckets(v []OSExternalStorageBucketInfo)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *OSExternalStoragePlatformInfo) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetEndpoint

`func (o *OSExternalStoragePlatformInfo) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *OSExternalStoragePlatformInfo) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *OSExternalStoragePlatformInfo) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *OSExternalStoragePlatformInfo) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetHostStyle

`func (o *OSExternalStoragePlatformInfo) GetHostStyle() string`

GetHostStyle returns the HostStyle field if non-nil, zero value otherwise.

### GetHostStyleOk

`func (o *OSExternalStoragePlatformInfo) GetHostStyleOk() (*string, bool)`

GetHostStyleOk returns a tuple with the HostStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostStyle

`func (o *OSExternalStoragePlatformInfo) SetHostStyle(v string)`

SetHostStyle sets HostStyle field to given value.

### HasHostStyle

`func (o *OSExternalStoragePlatformInfo) HasHostStyle() bool`

HasHostStyle returns a boolean if a field has been set.

### GetId

`func (o *OSExternalStoragePlatformInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSExternalStoragePlatformInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSExternalStoragePlatformInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OSExternalStoragePlatformInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OSExternalStoragePlatformInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OSExternalStoragePlatformInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OSExternalStoragePlatformInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OSExternalStoragePlatformInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *OSExternalStoragePlatformInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OSExternalStoragePlatformInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OSExternalStoragePlatformInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *OSExternalStoragePlatformInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretKey

`func (o *OSExternalStoragePlatformInfo) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *OSExternalStoragePlatformInfo) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *OSExternalStoragePlatformInfo) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *OSExternalStoragePlatformInfo) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetWeight

`func (o *OSExternalStoragePlatformInfo) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *OSExternalStoragePlatformInfo) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *OSExternalStoragePlatformInfo) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *OSExternalStoragePlatformInfo) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


