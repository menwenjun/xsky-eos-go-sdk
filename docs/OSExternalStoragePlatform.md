# OSExternalStoragePlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** |  | [optional] 
**Authentication** | Pointer to **string** |  | [optional] 
**BluRay** | Pointer to [**OSBuiltinBluRay**](OSBuiltinBluRay.md) |  | [optional] 
**Buckets** | Pointer to [**[]OSExternalStorageBucketInfo**](OSExternalStorageBucketInfo.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**HostStyle** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OsExternalStorageClass** | Pointer to [**OSExternalStorageClass**](OSExternalStorageClass.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int64** |  | [optional] 

## Methods

### NewOSExternalStoragePlatform

`func NewOSExternalStoragePlatform() *OSExternalStoragePlatform`

NewOSExternalStoragePlatform instantiates a new OSExternalStoragePlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSExternalStoragePlatformWithDefaults

`func NewOSExternalStoragePlatformWithDefaults() *OSExternalStoragePlatform`

NewOSExternalStoragePlatformWithDefaults instantiates a new OSExternalStoragePlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *OSExternalStoragePlatform) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *OSExternalStoragePlatform) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *OSExternalStoragePlatform) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *OSExternalStoragePlatform) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAuthentication

`func (o *OSExternalStoragePlatform) GetAuthentication() string`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *OSExternalStoragePlatform) GetAuthenticationOk() (*string, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *OSExternalStoragePlatform) SetAuthentication(v string)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *OSExternalStoragePlatform) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetBluRay

`func (o *OSExternalStoragePlatform) GetBluRay() OSBuiltinBluRay`

GetBluRay returns the BluRay field if non-nil, zero value otherwise.

### GetBluRayOk

`func (o *OSExternalStoragePlatform) GetBluRayOk() (*OSBuiltinBluRay, bool)`

GetBluRayOk returns a tuple with the BluRay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluRay

`func (o *OSExternalStoragePlatform) SetBluRay(v OSBuiltinBluRay)`

SetBluRay sets BluRay field to given value.

### HasBluRay

`func (o *OSExternalStoragePlatform) HasBluRay() bool`

HasBluRay returns a boolean if a field has been set.

### GetBuckets

`func (o *OSExternalStoragePlatform) GetBuckets() []OSExternalStorageBucketInfo`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *OSExternalStoragePlatform) GetBucketsOk() (*[]OSExternalStorageBucketInfo, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *OSExternalStoragePlatform) SetBuckets(v []OSExternalStorageBucketInfo)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *OSExternalStoragePlatform) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetCluster

`func (o *OSExternalStoragePlatform) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSExternalStoragePlatform) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSExternalStoragePlatform) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSExternalStoragePlatform) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConnected

`func (o *OSExternalStoragePlatform) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *OSExternalStoragePlatform) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *OSExternalStoragePlatform) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *OSExternalStoragePlatform) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetCreate

`func (o *OSExternalStoragePlatform) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSExternalStoragePlatform) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSExternalStoragePlatform) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSExternalStoragePlatform) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetEndpoint

`func (o *OSExternalStoragePlatform) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *OSExternalStoragePlatform) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *OSExternalStoragePlatform) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *OSExternalStoragePlatform) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetHostStyle

`func (o *OSExternalStoragePlatform) GetHostStyle() string`

GetHostStyle returns the HostStyle field if non-nil, zero value otherwise.

### GetHostStyleOk

`func (o *OSExternalStoragePlatform) GetHostStyleOk() (*string, bool)`

GetHostStyleOk returns a tuple with the HostStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostStyle

`func (o *OSExternalStoragePlatform) SetHostStyle(v string)`

SetHostStyle sets HostStyle field to given value.

### HasHostStyle

`func (o *OSExternalStoragePlatform) HasHostStyle() bool`

HasHostStyle returns a boolean if a field has been set.

### GetId

`func (o *OSExternalStoragePlatform) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSExternalStoragePlatform) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSExternalStoragePlatform) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OSExternalStoragePlatform) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OSExternalStoragePlatform) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OSExternalStoragePlatform) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OSExternalStoragePlatform) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OSExternalStoragePlatform) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsExternalStorageClass

`func (o *OSExternalStoragePlatform) GetOsExternalStorageClass() OSExternalStorageClass`

GetOsExternalStorageClass returns the OsExternalStorageClass field if non-nil, zero value otherwise.

### GetOsExternalStorageClassOk

`func (o *OSExternalStoragePlatform) GetOsExternalStorageClassOk() (*OSExternalStorageClass, bool)`

GetOsExternalStorageClassOk returns a tuple with the OsExternalStorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsExternalStorageClass

`func (o *OSExternalStoragePlatform) SetOsExternalStorageClass(v OSExternalStorageClass)`

SetOsExternalStorageClass sets OsExternalStorageClass field to given value.

### HasOsExternalStorageClass

`func (o *OSExternalStoragePlatform) HasOsExternalStorageClass() bool`

HasOsExternalStorageClass returns a boolean if a field has been set.

### GetRegion

`func (o *OSExternalStoragePlatform) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OSExternalStoragePlatform) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OSExternalStoragePlatform) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *OSExternalStoragePlatform) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretKey

`func (o *OSExternalStoragePlatform) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *OSExternalStoragePlatform) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *OSExternalStoragePlatform) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *OSExternalStoragePlatform) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetStatus

`func (o *OSExternalStoragePlatform) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSExternalStoragePlatform) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSExternalStoragePlatform) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OSExternalStoragePlatform) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUuid

`func (o *OSExternalStoragePlatform) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OSExternalStoragePlatform) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OSExternalStoragePlatform) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *OSExternalStoragePlatform) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWeight

`func (o *OSExternalStoragePlatform) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *OSExternalStoragePlatform) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *OSExternalStoragePlatform) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *OSExternalStoragePlatform) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


