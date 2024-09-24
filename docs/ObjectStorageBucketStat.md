# ObjectStorageBucketStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteObjCategory** | Pointer to [**OSBucketUsageCategory**](OSBucketUsageCategory.md) |  | [optional] 
**GetObjCategory** | Pointer to [**OSBucketUsageCategory**](OSBucketUsageCategory.md) |  | [optional] 
**ListBucketCategory** | Pointer to [**OSBucketUsageCategory**](OSBucketUsageCategory.md) |  | [optional] 
**Objects** | Pointer to **int64** |  | [optional] 
**PostObjCategory** | Pointer to [**OSBucketUsageCategory**](OSBucketUsageCategory.md) |  | [optional] 
**PutObjCategory** | Pointer to [**OSBucketUsageCategory**](OSBucketUsageCategory.md) |  | [optional] 
**UsedCapacityBytes** | Pointer to **map[string]int64** |  | [optional] 
**AllocatedObjects** | Pointer to **int64** |  | [optional] 
**AllocatedSize** | Pointer to **int64** |  | [optional] 
**CacheAllocatedObjects** | Pointer to **int64** |  | [optional] 
**CacheAllocatedSize** | Pointer to **int64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DelOpsPm** | Pointer to **float64** |  | [optional] 
**DownLatencyUs** | Pointer to **float64** |  | [optional] 
**ExternalAllocatedObjects** | Pointer to **int64** |  | [optional] 
**ExternalAllocatedSize** | Pointer to **int64** |  | [optional] 
**LatencyDown** | Pointer to **int64** |  | [optional] 
**LatencyUp** | Pointer to **int64** |  | [optional] 
**ListOpsPm** | Pointer to **float64** |  | [optional] 
**LocalAllocatedObjects** | Pointer to **int64** |  | [optional] 
**LocalAllocatedSize** | Pointer to **int64** |  | [optional] 
**NumDown** | Pointer to **int64** |  | [optional] 
**NumUp** | Pointer to **int64** |  | [optional] 
**RxBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**RxOpsPm** | Pointer to **float64** |  | [optional] 
**StorageClasses** | Pointer to [**map[string]OSStorageClassStat**](OSStorageClassStat.md) |  | [optional] 
**TotalDelOps** | Pointer to **int64** |  | [optional] 
**TotalDelSuccessOps** | Pointer to **int64** |  | [optional] 
**TotalRxBytes** | Pointer to **int64** |  | [optional] 
**TotalRxOps** | Pointer to **int64** |  | [optional] 
**TotalRxSuccessOps** | Pointer to **int64** |  | [optional] 
**TotalTxBytes** | Pointer to **int64** |  | [optional] 
**TotalTxOps** | Pointer to **int64** |  | [optional] 
**TotalTxSuccessOps** | Pointer to **int64** |  | [optional] 
**TxBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**TxOpsPm** | Pointer to **float64** |  | [optional] 
**UpLatencyUs** | Pointer to **float64** |  | [optional] 

## Methods

### NewObjectStorageBucketStat

`func NewObjectStorageBucketStat() *ObjectStorageBucketStat`

NewObjectStorageBucketStat instantiates a new ObjectStorageBucketStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageBucketStatWithDefaults

`func NewObjectStorageBucketStatWithDefaults() *ObjectStorageBucketStat`

NewObjectStorageBucketStatWithDefaults instantiates a new ObjectStorageBucketStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteObjCategory

`func (o *ObjectStorageBucketStat) GetDeleteObjCategory() OSBucketUsageCategory`

GetDeleteObjCategory returns the DeleteObjCategory field if non-nil, zero value otherwise.

### GetDeleteObjCategoryOk

`func (o *ObjectStorageBucketStat) GetDeleteObjCategoryOk() (*OSBucketUsageCategory, bool)`

GetDeleteObjCategoryOk returns a tuple with the DeleteObjCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteObjCategory

`func (o *ObjectStorageBucketStat) SetDeleteObjCategory(v OSBucketUsageCategory)`

SetDeleteObjCategory sets DeleteObjCategory field to given value.

### HasDeleteObjCategory

`func (o *ObjectStorageBucketStat) HasDeleteObjCategory() bool`

HasDeleteObjCategory returns a boolean if a field has been set.

### GetGetObjCategory

`func (o *ObjectStorageBucketStat) GetGetObjCategory() OSBucketUsageCategory`

GetGetObjCategory returns the GetObjCategory field if non-nil, zero value otherwise.

### GetGetObjCategoryOk

`func (o *ObjectStorageBucketStat) GetGetObjCategoryOk() (*OSBucketUsageCategory, bool)`

GetGetObjCategoryOk returns a tuple with the GetObjCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetObjCategory

`func (o *ObjectStorageBucketStat) SetGetObjCategory(v OSBucketUsageCategory)`

SetGetObjCategory sets GetObjCategory field to given value.

### HasGetObjCategory

`func (o *ObjectStorageBucketStat) HasGetObjCategory() bool`

HasGetObjCategory returns a boolean if a field has been set.

### GetListBucketCategory

`func (o *ObjectStorageBucketStat) GetListBucketCategory() OSBucketUsageCategory`

GetListBucketCategory returns the ListBucketCategory field if non-nil, zero value otherwise.

### GetListBucketCategoryOk

`func (o *ObjectStorageBucketStat) GetListBucketCategoryOk() (*OSBucketUsageCategory, bool)`

GetListBucketCategoryOk returns a tuple with the ListBucketCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListBucketCategory

`func (o *ObjectStorageBucketStat) SetListBucketCategory(v OSBucketUsageCategory)`

SetListBucketCategory sets ListBucketCategory field to given value.

### HasListBucketCategory

`func (o *ObjectStorageBucketStat) HasListBucketCategory() bool`

HasListBucketCategory returns a boolean if a field has been set.

### GetObjects

`func (o *ObjectStorageBucketStat) GetObjects() int64`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ObjectStorageBucketStat) GetObjectsOk() (*int64, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ObjectStorageBucketStat) SetObjects(v int64)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ObjectStorageBucketStat) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetPostObjCategory

`func (o *ObjectStorageBucketStat) GetPostObjCategory() OSBucketUsageCategory`

GetPostObjCategory returns the PostObjCategory field if non-nil, zero value otherwise.

### GetPostObjCategoryOk

`func (o *ObjectStorageBucketStat) GetPostObjCategoryOk() (*OSBucketUsageCategory, bool)`

GetPostObjCategoryOk returns a tuple with the PostObjCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostObjCategory

`func (o *ObjectStorageBucketStat) SetPostObjCategory(v OSBucketUsageCategory)`

SetPostObjCategory sets PostObjCategory field to given value.

### HasPostObjCategory

`func (o *ObjectStorageBucketStat) HasPostObjCategory() bool`

HasPostObjCategory returns a boolean if a field has been set.

### GetPutObjCategory

`func (o *ObjectStorageBucketStat) GetPutObjCategory() OSBucketUsageCategory`

GetPutObjCategory returns the PutObjCategory field if non-nil, zero value otherwise.

### GetPutObjCategoryOk

`func (o *ObjectStorageBucketStat) GetPutObjCategoryOk() (*OSBucketUsageCategory, bool)`

GetPutObjCategoryOk returns a tuple with the PutObjCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutObjCategory

`func (o *ObjectStorageBucketStat) SetPutObjCategory(v OSBucketUsageCategory)`

SetPutObjCategory sets PutObjCategory field to given value.

### HasPutObjCategory

`func (o *ObjectStorageBucketStat) HasPutObjCategory() bool`

HasPutObjCategory returns a boolean if a field has been set.

### GetUsedCapacityBytes

`func (o *ObjectStorageBucketStat) GetUsedCapacityBytes() map[string]int64`

GetUsedCapacityBytes returns the UsedCapacityBytes field if non-nil, zero value otherwise.

### GetUsedCapacityBytesOk

`func (o *ObjectStorageBucketStat) GetUsedCapacityBytesOk() (*map[string]int64, bool)`

GetUsedCapacityBytesOk returns a tuple with the UsedCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCapacityBytes

`func (o *ObjectStorageBucketStat) SetUsedCapacityBytes(v map[string]int64)`

SetUsedCapacityBytes sets UsedCapacityBytes field to given value.

### HasUsedCapacityBytes

`func (o *ObjectStorageBucketStat) HasUsedCapacityBytes() bool`

HasUsedCapacityBytes returns a boolean if a field has been set.

### GetAllocatedObjects

`func (o *ObjectStorageBucketStat) GetAllocatedObjects() int64`

GetAllocatedObjects returns the AllocatedObjects field if non-nil, zero value otherwise.

### GetAllocatedObjectsOk

`func (o *ObjectStorageBucketStat) GetAllocatedObjectsOk() (*int64, bool)`

GetAllocatedObjectsOk returns a tuple with the AllocatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedObjects

`func (o *ObjectStorageBucketStat) SetAllocatedObjects(v int64)`

SetAllocatedObjects sets AllocatedObjects field to given value.

### HasAllocatedObjects

`func (o *ObjectStorageBucketStat) HasAllocatedObjects() bool`

HasAllocatedObjects returns a boolean if a field has been set.

### GetAllocatedSize

`func (o *ObjectStorageBucketStat) GetAllocatedSize() int64`

GetAllocatedSize returns the AllocatedSize field if non-nil, zero value otherwise.

### GetAllocatedSizeOk

`func (o *ObjectStorageBucketStat) GetAllocatedSizeOk() (*int64, bool)`

GetAllocatedSizeOk returns a tuple with the AllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedSize

`func (o *ObjectStorageBucketStat) SetAllocatedSize(v int64)`

SetAllocatedSize sets AllocatedSize field to given value.

### HasAllocatedSize

`func (o *ObjectStorageBucketStat) HasAllocatedSize() bool`

HasAllocatedSize returns a boolean if a field has been set.

### GetCacheAllocatedObjects

`func (o *ObjectStorageBucketStat) GetCacheAllocatedObjects() int64`

GetCacheAllocatedObjects returns the CacheAllocatedObjects field if non-nil, zero value otherwise.

### GetCacheAllocatedObjectsOk

`func (o *ObjectStorageBucketStat) GetCacheAllocatedObjectsOk() (*int64, bool)`

GetCacheAllocatedObjectsOk returns a tuple with the CacheAllocatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAllocatedObjects

`func (o *ObjectStorageBucketStat) SetCacheAllocatedObjects(v int64)`

SetCacheAllocatedObjects sets CacheAllocatedObjects field to given value.

### HasCacheAllocatedObjects

`func (o *ObjectStorageBucketStat) HasCacheAllocatedObjects() bool`

HasCacheAllocatedObjects returns a boolean if a field has been set.

### GetCacheAllocatedSize

`func (o *ObjectStorageBucketStat) GetCacheAllocatedSize() int64`

GetCacheAllocatedSize returns the CacheAllocatedSize field if non-nil, zero value otherwise.

### GetCacheAllocatedSizeOk

`func (o *ObjectStorageBucketStat) GetCacheAllocatedSizeOk() (*int64, bool)`

GetCacheAllocatedSizeOk returns a tuple with the CacheAllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAllocatedSize

`func (o *ObjectStorageBucketStat) SetCacheAllocatedSize(v int64)`

SetCacheAllocatedSize sets CacheAllocatedSize field to given value.

### HasCacheAllocatedSize

`func (o *ObjectStorageBucketStat) HasCacheAllocatedSize() bool`

HasCacheAllocatedSize returns a boolean if a field has been set.

### GetCreate

`func (o *ObjectStorageBucketStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ObjectStorageBucketStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ObjectStorageBucketStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ObjectStorageBucketStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDelOpsPm

`func (o *ObjectStorageBucketStat) GetDelOpsPm() float64`

GetDelOpsPm returns the DelOpsPm field if non-nil, zero value otherwise.

### GetDelOpsPmOk

`func (o *ObjectStorageBucketStat) GetDelOpsPmOk() (*float64, bool)`

GetDelOpsPmOk returns a tuple with the DelOpsPm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelOpsPm

`func (o *ObjectStorageBucketStat) SetDelOpsPm(v float64)`

SetDelOpsPm sets DelOpsPm field to given value.

### HasDelOpsPm

`func (o *ObjectStorageBucketStat) HasDelOpsPm() bool`

HasDelOpsPm returns a boolean if a field has been set.

### GetDownLatencyUs

`func (o *ObjectStorageBucketStat) GetDownLatencyUs() float64`

GetDownLatencyUs returns the DownLatencyUs field if non-nil, zero value otherwise.

### GetDownLatencyUsOk

`func (o *ObjectStorageBucketStat) GetDownLatencyUsOk() (*float64, bool)`

GetDownLatencyUsOk returns a tuple with the DownLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLatencyUs

`func (o *ObjectStorageBucketStat) SetDownLatencyUs(v float64)`

SetDownLatencyUs sets DownLatencyUs field to given value.

### HasDownLatencyUs

`func (o *ObjectStorageBucketStat) HasDownLatencyUs() bool`

HasDownLatencyUs returns a boolean if a field has been set.

### GetExternalAllocatedObjects

`func (o *ObjectStorageBucketStat) GetExternalAllocatedObjects() int64`

GetExternalAllocatedObjects returns the ExternalAllocatedObjects field if non-nil, zero value otherwise.

### GetExternalAllocatedObjectsOk

`func (o *ObjectStorageBucketStat) GetExternalAllocatedObjectsOk() (*int64, bool)`

GetExternalAllocatedObjectsOk returns a tuple with the ExternalAllocatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAllocatedObjects

`func (o *ObjectStorageBucketStat) SetExternalAllocatedObjects(v int64)`

SetExternalAllocatedObjects sets ExternalAllocatedObjects field to given value.

### HasExternalAllocatedObjects

`func (o *ObjectStorageBucketStat) HasExternalAllocatedObjects() bool`

HasExternalAllocatedObjects returns a boolean if a field has been set.

### GetExternalAllocatedSize

`func (o *ObjectStorageBucketStat) GetExternalAllocatedSize() int64`

GetExternalAllocatedSize returns the ExternalAllocatedSize field if non-nil, zero value otherwise.

### GetExternalAllocatedSizeOk

`func (o *ObjectStorageBucketStat) GetExternalAllocatedSizeOk() (*int64, bool)`

GetExternalAllocatedSizeOk returns a tuple with the ExternalAllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAllocatedSize

`func (o *ObjectStorageBucketStat) SetExternalAllocatedSize(v int64)`

SetExternalAllocatedSize sets ExternalAllocatedSize field to given value.

### HasExternalAllocatedSize

`func (o *ObjectStorageBucketStat) HasExternalAllocatedSize() bool`

HasExternalAllocatedSize returns a boolean if a field has been set.

### GetLatencyDown

`func (o *ObjectStorageBucketStat) GetLatencyDown() int64`

GetLatencyDown returns the LatencyDown field if non-nil, zero value otherwise.

### GetLatencyDownOk

`func (o *ObjectStorageBucketStat) GetLatencyDownOk() (*int64, bool)`

GetLatencyDownOk returns a tuple with the LatencyDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyDown

`func (o *ObjectStorageBucketStat) SetLatencyDown(v int64)`

SetLatencyDown sets LatencyDown field to given value.

### HasLatencyDown

`func (o *ObjectStorageBucketStat) HasLatencyDown() bool`

HasLatencyDown returns a boolean if a field has been set.

### GetLatencyUp

`func (o *ObjectStorageBucketStat) GetLatencyUp() int64`

GetLatencyUp returns the LatencyUp field if non-nil, zero value otherwise.

### GetLatencyUpOk

`func (o *ObjectStorageBucketStat) GetLatencyUpOk() (*int64, bool)`

GetLatencyUpOk returns a tuple with the LatencyUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyUp

`func (o *ObjectStorageBucketStat) SetLatencyUp(v int64)`

SetLatencyUp sets LatencyUp field to given value.

### HasLatencyUp

`func (o *ObjectStorageBucketStat) HasLatencyUp() bool`

HasLatencyUp returns a boolean if a field has been set.

### GetListOpsPm

`func (o *ObjectStorageBucketStat) GetListOpsPm() float64`

GetListOpsPm returns the ListOpsPm field if non-nil, zero value otherwise.

### GetListOpsPmOk

`func (o *ObjectStorageBucketStat) GetListOpsPmOk() (*float64, bool)`

GetListOpsPmOk returns a tuple with the ListOpsPm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOpsPm

`func (o *ObjectStorageBucketStat) SetListOpsPm(v float64)`

SetListOpsPm sets ListOpsPm field to given value.

### HasListOpsPm

`func (o *ObjectStorageBucketStat) HasListOpsPm() bool`

HasListOpsPm returns a boolean if a field has been set.

### GetLocalAllocatedObjects

`func (o *ObjectStorageBucketStat) GetLocalAllocatedObjects() int64`

GetLocalAllocatedObjects returns the LocalAllocatedObjects field if non-nil, zero value otherwise.

### GetLocalAllocatedObjectsOk

`func (o *ObjectStorageBucketStat) GetLocalAllocatedObjectsOk() (*int64, bool)`

GetLocalAllocatedObjectsOk returns a tuple with the LocalAllocatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAllocatedObjects

`func (o *ObjectStorageBucketStat) SetLocalAllocatedObjects(v int64)`

SetLocalAllocatedObjects sets LocalAllocatedObjects field to given value.

### HasLocalAllocatedObjects

`func (o *ObjectStorageBucketStat) HasLocalAllocatedObjects() bool`

HasLocalAllocatedObjects returns a boolean if a field has been set.

### GetLocalAllocatedSize

`func (o *ObjectStorageBucketStat) GetLocalAllocatedSize() int64`

GetLocalAllocatedSize returns the LocalAllocatedSize field if non-nil, zero value otherwise.

### GetLocalAllocatedSizeOk

`func (o *ObjectStorageBucketStat) GetLocalAllocatedSizeOk() (*int64, bool)`

GetLocalAllocatedSizeOk returns a tuple with the LocalAllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAllocatedSize

`func (o *ObjectStorageBucketStat) SetLocalAllocatedSize(v int64)`

SetLocalAllocatedSize sets LocalAllocatedSize field to given value.

### HasLocalAllocatedSize

`func (o *ObjectStorageBucketStat) HasLocalAllocatedSize() bool`

HasLocalAllocatedSize returns a boolean if a field has been set.

### GetNumDown

`func (o *ObjectStorageBucketStat) GetNumDown() int64`

GetNumDown returns the NumDown field if non-nil, zero value otherwise.

### GetNumDownOk

`func (o *ObjectStorageBucketStat) GetNumDownOk() (*int64, bool)`

GetNumDownOk returns a tuple with the NumDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDown

`func (o *ObjectStorageBucketStat) SetNumDown(v int64)`

SetNumDown sets NumDown field to given value.

### HasNumDown

`func (o *ObjectStorageBucketStat) HasNumDown() bool`

HasNumDown returns a boolean if a field has been set.

### GetNumUp

`func (o *ObjectStorageBucketStat) GetNumUp() int64`

GetNumUp returns the NumUp field if non-nil, zero value otherwise.

### GetNumUpOk

`func (o *ObjectStorageBucketStat) GetNumUpOk() (*int64, bool)`

GetNumUpOk returns a tuple with the NumUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUp

`func (o *ObjectStorageBucketStat) SetNumUp(v int64)`

SetNumUp sets NumUp field to given value.

### HasNumUp

`func (o *ObjectStorageBucketStat) HasNumUp() bool`

HasNumUp returns a boolean if a field has been set.

### GetRxBandwidthKbyte

`func (o *ObjectStorageBucketStat) GetRxBandwidthKbyte() float64`

GetRxBandwidthKbyte returns the RxBandwidthKbyte field if non-nil, zero value otherwise.

### GetRxBandwidthKbyteOk

`func (o *ObjectStorageBucketStat) GetRxBandwidthKbyteOk() (*float64, bool)`

GetRxBandwidthKbyteOk returns a tuple with the RxBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBandwidthKbyte

`func (o *ObjectStorageBucketStat) SetRxBandwidthKbyte(v float64)`

SetRxBandwidthKbyte sets RxBandwidthKbyte field to given value.

### HasRxBandwidthKbyte

`func (o *ObjectStorageBucketStat) HasRxBandwidthKbyte() bool`

HasRxBandwidthKbyte returns a boolean if a field has been set.

### GetRxOpsPm

`func (o *ObjectStorageBucketStat) GetRxOpsPm() float64`

GetRxOpsPm returns the RxOpsPm field if non-nil, zero value otherwise.

### GetRxOpsPmOk

`func (o *ObjectStorageBucketStat) GetRxOpsPmOk() (*float64, bool)`

GetRxOpsPmOk returns a tuple with the RxOpsPm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxOpsPm

`func (o *ObjectStorageBucketStat) SetRxOpsPm(v float64)`

SetRxOpsPm sets RxOpsPm field to given value.

### HasRxOpsPm

`func (o *ObjectStorageBucketStat) HasRxOpsPm() bool`

HasRxOpsPm returns a boolean if a field has been set.

### GetStorageClasses

`func (o *ObjectStorageBucketStat) GetStorageClasses() map[string]OSStorageClassStat`

GetStorageClasses returns the StorageClasses field if non-nil, zero value otherwise.

### GetStorageClassesOk

`func (o *ObjectStorageBucketStat) GetStorageClassesOk() (*map[string]OSStorageClassStat, bool)`

GetStorageClassesOk returns a tuple with the StorageClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClasses

`func (o *ObjectStorageBucketStat) SetStorageClasses(v map[string]OSStorageClassStat)`

SetStorageClasses sets StorageClasses field to given value.

### HasStorageClasses

`func (o *ObjectStorageBucketStat) HasStorageClasses() bool`

HasStorageClasses returns a boolean if a field has been set.

### GetTotalDelOps

`func (o *ObjectStorageBucketStat) GetTotalDelOps() int64`

GetTotalDelOps returns the TotalDelOps field if non-nil, zero value otherwise.

### GetTotalDelOpsOk

`func (o *ObjectStorageBucketStat) GetTotalDelOpsOk() (*int64, bool)`

GetTotalDelOpsOk returns a tuple with the TotalDelOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDelOps

`func (o *ObjectStorageBucketStat) SetTotalDelOps(v int64)`

SetTotalDelOps sets TotalDelOps field to given value.

### HasTotalDelOps

`func (o *ObjectStorageBucketStat) HasTotalDelOps() bool`

HasTotalDelOps returns a boolean if a field has been set.

### GetTotalDelSuccessOps

`func (o *ObjectStorageBucketStat) GetTotalDelSuccessOps() int64`

GetTotalDelSuccessOps returns the TotalDelSuccessOps field if non-nil, zero value otherwise.

### GetTotalDelSuccessOpsOk

`func (o *ObjectStorageBucketStat) GetTotalDelSuccessOpsOk() (*int64, bool)`

GetTotalDelSuccessOpsOk returns a tuple with the TotalDelSuccessOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDelSuccessOps

`func (o *ObjectStorageBucketStat) SetTotalDelSuccessOps(v int64)`

SetTotalDelSuccessOps sets TotalDelSuccessOps field to given value.

### HasTotalDelSuccessOps

`func (o *ObjectStorageBucketStat) HasTotalDelSuccessOps() bool`

HasTotalDelSuccessOps returns a boolean if a field has been set.

### GetTotalRxBytes

`func (o *ObjectStorageBucketStat) GetTotalRxBytes() int64`

GetTotalRxBytes returns the TotalRxBytes field if non-nil, zero value otherwise.

### GetTotalRxBytesOk

`func (o *ObjectStorageBucketStat) GetTotalRxBytesOk() (*int64, bool)`

GetTotalRxBytesOk returns a tuple with the TotalRxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRxBytes

`func (o *ObjectStorageBucketStat) SetTotalRxBytes(v int64)`

SetTotalRxBytes sets TotalRxBytes field to given value.

### HasTotalRxBytes

`func (o *ObjectStorageBucketStat) HasTotalRxBytes() bool`

HasTotalRxBytes returns a boolean if a field has been set.

### GetTotalRxOps

`func (o *ObjectStorageBucketStat) GetTotalRxOps() int64`

GetTotalRxOps returns the TotalRxOps field if non-nil, zero value otherwise.

### GetTotalRxOpsOk

`func (o *ObjectStorageBucketStat) GetTotalRxOpsOk() (*int64, bool)`

GetTotalRxOpsOk returns a tuple with the TotalRxOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRxOps

`func (o *ObjectStorageBucketStat) SetTotalRxOps(v int64)`

SetTotalRxOps sets TotalRxOps field to given value.

### HasTotalRxOps

`func (o *ObjectStorageBucketStat) HasTotalRxOps() bool`

HasTotalRxOps returns a boolean if a field has been set.

### GetTotalRxSuccessOps

`func (o *ObjectStorageBucketStat) GetTotalRxSuccessOps() int64`

GetTotalRxSuccessOps returns the TotalRxSuccessOps field if non-nil, zero value otherwise.

### GetTotalRxSuccessOpsOk

`func (o *ObjectStorageBucketStat) GetTotalRxSuccessOpsOk() (*int64, bool)`

GetTotalRxSuccessOpsOk returns a tuple with the TotalRxSuccessOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRxSuccessOps

`func (o *ObjectStorageBucketStat) SetTotalRxSuccessOps(v int64)`

SetTotalRxSuccessOps sets TotalRxSuccessOps field to given value.

### HasTotalRxSuccessOps

`func (o *ObjectStorageBucketStat) HasTotalRxSuccessOps() bool`

HasTotalRxSuccessOps returns a boolean if a field has been set.

### GetTotalTxBytes

`func (o *ObjectStorageBucketStat) GetTotalTxBytes() int64`

GetTotalTxBytes returns the TotalTxBytes field if non-nil, zero value otherwise.

### GetTotalTxBytesOk

`func (o *ObjectStorageBucketStat) GetTotalTxBytesOk() (*int64, bool)`

GetTotalTxBytesOk returns a tuple with the TotalTxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTxBytes

`func (o *ObjectStorageBucketStat) SetTotalTxBytes(v int64)`

SetTotalTxBytes sets TotalTxBytes field to given value.

### HasTotalTxBytes

`func (o *ObjectStorageBucketStat) HasTotalTxBytes() bool`

HasTotalTxBytes returns a boolean if a field has been set.

### GetTotalTxOps

`func (o *ObjectStorageBucketStat) GetTotalTxOps() int64`

GetTotalTxOps returns the TotalTxOps field if non-nil, zero value otherwise.

### GetTotalTxOpsOk

`func (o *ObjectStorageBucketStat) GetTotalTxOpsOk() (*int64, bool)`

GetTotalTxOpsOk returns a tuple with the TotalTxOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTxOps

`func (o *ObjectStorageBucketStat) SetTotalTxOps(v int64)`

SetTotalTxOps sets TotalTxOps field to given value.

### HasTotalTxOps

`func (o *ObjectStorageBucketStat) HasTotalTxOps() bool`

HasTotalTxOps returns a boolean if a field has been set.

### GetTotalTxSuccessOps

`func (o *ObjectStorageBucketStat) GetTotalTxSuccessOps() int64`

GetTotalTxSuccessOps returns the TotalTxSuccessOps field if non-nil, zero value otherwise.

### GetTotalTxSuccessOpsOk

`func (o *ObjectStorageBucketStat) GetTotalTxSuccessOpsOk() (*int64, bool)`

GetTotalTxSuccessOpsOk returns a tuple with the TotalTxSuccessOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTxSuccessOps

`func (o *ObjectStorageBucketStat) SetTotalTxSuccessOps(v int64)`

SetTotalTxSuccessOps sets TotalTxSuccessOps field to given value.

### HasTotalTxSuccessOps

`func (o *ObjectStorageBucketStat) HasTotalTxSuccessOps() bool`

HasTotalTxSuccessOps returns a boolean if a field has been set.

### GetTxBandwidthKbyte

`func (o *ObjectStorageBucketStat) GetTxBandwidthKbyte() float64`

GetTxBandwidthKbyte returns the TxBandwidthKbyte field if non-nil, zero value otherwise.

### GetTxBandwidthKbyteOk

`func (o *ObjectStorageBucketStat) GetTxBandwidthKbyteOk() (*float64, bool)`

GetTxBandwidthKbyteOk returns a tuple with the TxBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBandwidthKbyte

`func (o *ObjectStorageBucketStat) SetTxBandwidthKbyte(v float64)`

SetTxBandwidthKbyte sets TxBandwidthKbyte field to given value.

### HasTxBandwidthKbyte

`func (o *ObjectStorageBucketStat) HasTxBandwidthKbyte() bool`

HasTxBandwidthKbyte returns a boolean if a field has been set.

### GetTxOpsPm

`func (o *ObjectStorageBucketStat) GetTxOpsPm() float64`

GetTxOpsPm returns the TxOpsPm field if non-nil, zero value otherwise.

### GetTxOpsPmOk

`func (o *ObjectStorageBucketStat) GetTxOpsPmOk() (*float64, bool)`

GetTxOpsPmOk returns a tuple with the TxOpsPm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxOpsPm

`func (o *ObjectStorageBucketStat) SetTxOpsPm(v float64)`

SetTxOpsPm sets TxOpsPm field to given value.

### HasTxOpsPm

`func (o *ObjectStorageBucketStat) HasTxOpsPm() bool`

HasTxOpsPm returns a boolean if a field has been set.

### GetUpLatencyUs

`func (o *ObjectStorageBucketStat) GetUpLatencyUs() float64`

GetUpLatencyUs returns the UpLatencyUs field if non-nil, zero value otherwise.

### GetUpLatencyUsOk

`func (o *ObjectStorageBucketStat) GetUpLatencyUsOk() (*float64, bool)`

GetUpLatencyUsOk returns a tuple with the UpLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLatencyUs

`func (o *ObjectStorageBucketStat) SetUpLatencyUs(v float64)`

SetUpLatencyUs sets UpLatencyUs field to given value.

### HasUpLatencyUs

`func (o *ObjectStorageBucketStat) HasUpLatencyUs() bool`

HasUpLatencyUs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


