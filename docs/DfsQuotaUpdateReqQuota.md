# DfsQuotaUpdateReqQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesGraceTime** | **int64** | grace time for files soft quota | 
**FilesHardQuota** | **int64** | hard quota of files | 
**FilesRatio** | Pointer to **float32** | ratio of files | [optional] 
**FilesSoftQuota** | **int64** | soft quota of files | 
**FilesSuggestQuota** | **int64** | suggest quota of files | 
**SizeGraceTime** | **int64** | grace time for size soft quota | 
**SizeHardQuota** | **int64** | hard quota of size | 
**SizeRatio** | Pointer to **float32** | ratio of size | [optional] 
**SizeSoftQuota** | **int64** | soft quota of size | 
**SizeSuggestQuota** | **int64** | suggest quota of size | 

## Methods

### NewDfsQuotaUpdateReqQuota

`func NewDfsQuotaUpdateReqQuota(filesGraceTime int64, filesHardQuota int64, filesSoftQuota int64, filesSuggestQuota int64, sizeGraceTime int64, sizeHardQuota int64, sizeSoftQuota int64, sizeSuggestQuota int64, ) *DfsQuotaUpdateReqQuota`

NewDfsQuotaUpdateReqQuota instantiates a new DfsQuotaUpdateReqQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsQuotaUpdateReqQuotaWithDefaults

`func NewDfsQuotaUpdateReqQuotaWithDefaults() *DfsQuotaUpdateReqQuota`

NewDfsQuotaUpdateReqQuotaWithDefaults instantiates a new DfsQuotaUpdateReqQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesGraceTime

`func (o *DfsQuotaUpdateReqQuota) GetFilesGraceTime() int64`

GetFilesGraceTime returns the FilesGraceTime field if non-nil, zero value otherwise.

### GetFilesGraceTimeOk

`func (o *DfsQuotaUpdateReqQuota) GetFilesGraceTimeOk() (*int64, bool)`

GetFilesGraceTimeOk returns a tuple with the FilesGraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesGraceTime

`func (o *DfsQuotaUpdateReqQuota) SetFilesGraceTime(v int64)`

SetFilesGraceTime sets FilesGraceTime field to given value.


### GetFilesHardQuota

`func (o *DfsQuotaUpdateReqQuota) GetFilesHardQuota() int64`

GetFilesHardQuota returns the FilesHardQuota field if non-nil, zero value otherwise.

### GetFilesHardQuotaOk

`func (o *DfsQuotaUpdateReqQuota) GetFilesHardQuotaOk() (*int64, bool)`

GetFilesHardQuotaOk returns a tuple with the FilesHardQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesHardQuota

`func (o *DfsQuotaUpdateReqQuota) SetFilesHardQuota(v int64)`

SetFilesHardQuota sets FilesHardQuota field to given value.


### GetFilesRatio

`func (o *DfsQuotaUpdateReqQuota) GetFilesRatio() float32`

GetFilesRatio returns the FilesRatio field if non-nil, zero value otherwise.

### GetFilesRatioOk

`func (o *DfsQuotaUpdateReqQuota) GetFilesRatioOk() (*float32, bool)`

GetFilesRatioOk returns a tuple with the FilesRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesRatio

`func (o *DfsQuotaUpdateReqQuota) SetFilesRatio(v float32)`

SetFilesRatio sets FilesRatio field to given value.

### HasFilesRatio

`func (o *DfsQuotaUpdateReqQuota) HasFilesRatio() bool`

HasFilesRatio returns a boolean if a field has been set.

### GetFilesSoftQuota

`func (o *DfsQuotaUpdateReqQuota) GetFilesSoftQuota() int64`

GetFilesSoftQuota returns the FilesSoftQuota field if non-nil, zero value otherwise.

### GetFilesSoftQuotaOk

`func (o *DfsQuotaUpdateReqQuota) GetFilesSoftQuotaOk() (*int64, bool)`

GetFilesSoftQuotaOk returns a tuple with the FilesSoftQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesSoftQuota

`func (o *DfsQuotaUpdateReqQuota) SetFilesSoftQuota(v int64)`

SetFilesSoftQuota sets FilesSoftQuota field to given value.


### GetFilesSuggestQuota

`func (o *DfsQuotaUpdateReqQuota) GetFilesSuggestQuota() int64`

GetFilesSuggestQuota returns the FilesSuggestQuota field if non-nil, zero value otherwise.

### GetFilesSuggestQuotaOk

`func (o *DfsQuotaUpdateReqQuota) GetFilesSuggestQuotaOk() (*int64, bool)`

GetFilesSuggestQuotaOk returns a tuple with the FilesSuggestQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesSuggestQuota

`func (o *DfsQuotaUpdateReqQuota) SetFilesSuggestQuota(v int64)`

SetFilesSuggestQuota sets FilesSuggestQuota field to given value.


### GetSizeGraceTime

`func (o *DfsQuotaUpdateReqQuota) GetSizeGraceTime() int64`

GetSizeGraceTime returns the SizeGraceTime field if non-nil, zero value otherwise.

### GetSizeGraceTimeOk

`func (o *DfsQuotaUpdateReqQuota) GetSizeGraceTimeOk() (*int64, bool)`

GetSizeGraceTimeOk returns a tuple with the SizeGraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGraceTime

`func (o *DfsQuotaUpdateReqQuota) SetSizeGraceTime(v int64)`

SetSizeGraceTime sets SizeGraceTime field to given value.


### GetSizeHardQuota

`func (o *DfsQuotaUpdateReqQuota) GetSizeHardQuota() int64`

GetSizeHardQuota returns the SizeHardQuota field if non-nil, zero value otherwise.

### GetSizeHardQuotaOk

`func (o *DfsQuotaUpdateReqQuota) GetSizeHardQuotaOk() (*int64, bool)`

GetSizeHardQuotaOk returns a tuple with the SizeHardQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeHardQuota

`func (o *DfsQuotaUpdateReqQuota) SetSizeHardQuota(v int64)`

SetSizeHardQuota sets SizeHardQuota field to given value.


### GetSizeRatio

`func (o *DfsQuotaUpdateReqQuota) GetSizeRatio() float32`

GetSizeRatio returns the SizeRatio field if non-nil, zero value otherwise.

### GetSizeRatioOk

`func (o *DfsQuotaUpdateReqQuota) GetSizeRatioOk() (*float32, bool)`

GetSizeRatioOk returns a tuple with the SizeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeRatio

`func (o *DfsQuotaUpdateReqQuota) SetSizeRatio(v float32)`

SetSizeRatio sets SizeRatio field to given value.

### HasSizeRatio

`func (o *DfsQuotaUpdateReqQuota) HasSizeRatio() bool`

HasSizeRatio returns a boolean if a field has been set.

### GetSizeSoftQuota

`func (o *DfsQuotaUpdateReqQuota) GetSizeSoftQuota() int64`

GetSizeSoftQuota returns the SizeSoftQuota field if non-nil, zero value otherwise.

### GetSizeSoftQuotaOk

`func (o *DfsQuotaUpdateReqQuota) GetSizeSoftQuotaOk() (*int64, bool)`

GetSizeSoftQuotaOk returns a tuple with the SizeSoftQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSoftQuota

`func (o *DfsQuotaUpdateReqQuota) SetSizeSoftQuota(v int64)`

SetSizeSoftQuota sets SizeSoftQuota field to given value.


### GetSizeSuggestQuota

`func (o *DfsQuotaUpdateReqQuota) GetSizeSuggestQuota() int64`

GetSizeSuggestQuota returns the SizeSuggestQuota field if non-nil, zero value otherwise.

### GetSizeSuggestQuotaOk

`func (o *DfsQuotaUpdateReqQuota) GetSizeSuggestQuotaOk() (*int64, bool)`

GetSizeSuggestQuotaOk returns a tuple with the SizeSuggestQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSuggestQuota

`func (o *DfsQuotaUpdateReqQuota) SetSizeSuggestQuota(v int64)`

SetSizeSuggestQuota sets SizeSuggestQuota field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


