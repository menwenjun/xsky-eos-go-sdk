# DfsCSIDirectoryReqQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesGraceTime** | Pointer to **int64** | grace time for files soft quota | [optional] 
**FilesHardQuota** | Pointer to **int64** | hard quota of files | [optional] 
**FilesRatio** | Pointer to **float32** | ratio of files | [optional] 
**FilesSoftQuota** | Pointer to **int64** | soft quota of files | [optional] 
**FilesSuggestQuota** | Pointer to **int64** | suggest quota of files | [optional] 
**SizeGraceTime** | Pointer to **int64** | grace time for size soft quota | [optional] 
**SizeHardQuota** | Pointer to **int64** | hard quota of size | [optional] 
**SizeRatio** | Pointer to **float32** | ratio of size | [optional] 
**SizeSoftQuota** | Pointer to **int64** | soft quota of size | [optional] 
**SizeSuggestQuota** | Pointer to **int64** | suggest quota of size | [optional] 

## Methods

### NewDfsCSIDirectoryReqQuota

`func NewDfsCSIDirectoryReqQuota() *DfsCSIDirectoryReqQuota`

NewDfsCSIDirectoryReqQuota instantiates a new DfsCSIDirectoryReqQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsCSIDirectoryReqQuotaWithDefaults

`func NewDfsCSIDirectoryReqQuotaWithDefaults() *DfsCSIDirectoryReqQuota`

NewDfsCSIDirectoryReqQuotaWithDefaults instantiates a new DfsCSIDirectoryReqQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesGraceTime

`func (o *DfsCSIDirectoryReqQuota) GetFilesGraceTime() int64`

GetFilesGraceTime returns the FilesGraceTime field if non-nil, zero value otherwise.

### GetFilesGraceTimeOk

`func (o *DfsCSIDirectoryReqQuota) GetFilesGraceTimeOk() (*int64, bool)`

GetFilesGraceTimeOk returns a tuple with the FilesGraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesGraceTime

`func (o *DfsCSIDirectoryReqQuota) SetFilesGraceTime(v int64)`

SetFilesGraceTime sets FilesGraceTime field to given value.

### HasFilesGraceTime

`func (o *DfsCSIDirectoryReqQuota) HasFilesGraceTime() bool`

HasFilesGraceTime returns a boolean if a field has been set.

### GetFilesHardQuota

`func (o *DfsCSIDirectoryReqQuota) GetFilesHardQuota() int64`

GetFilesHardQuota returns the FilesHardQuota field if non-nil, zero value otherwise.

### GetFilesHardQuotaOk

`func (o *DfsCSIDirectoryReqQuota) GetFilesHardQuotaOk() (*int64, bool)`

GetFilesHardQuotaOk returns a tuple with the FilesHardQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesHardQuota

`func (o *DfsCSIDirectoryReqQuota) SetFilesHardQuota(v int64)`

SetFilesHardQuota sets FilesHardQuota field to given value.

### HasFilesHardQuota

`func (o *DfsCSIDirectoryReqQuota) HasFilesHardQuota() bool`

HasFilesHardQuota returns a boolean if a field has been set.

### GetFilesRatio

`func (o *DfsCSIDirectoryReqQuota) GetFilesRatio() float32`

GetFilesRatio returns the FilesRatio field if non-nil, zero value otherwise.

### GetFilesRatioOk

`func (o *DfsCSIDirectoryReqQuota) GetFilesRatioOk() (*float32, bool)`

GetFilesRatioOk returns a tuple with the FilesRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesRatio

`func (o *DfsCSIDirectoryReqQuota) SetFilesRatio(v float32)`

SetFilesRatio sets FilesRatio field to given value.

### HasFilesRatio

`func (o *DfsCSIDirectoryReqQuota) HasFilesRatio() bool`

HasFilesRatio returns a boolean if a field has been set.

### GetFilesSoftQuota

`func (o *DfsCSIDirectoryReqQuota) GetFilesSoftQuota() int64`

GetFilesSoftQuota returns the FilesSoftQuota field if non-nil, zero value otherwise.

### GetFilesSoftQuotaOk

`func (o *DfsCSIDirectoryReqQuota) GetFilesSoftQuotaOk() (*int64, bool)`

GetFilesSoftQuotaOk returns a tuple with the FilesSoftQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesSoftQuota

`func (o *DfsCSIDirectoryReqQuota) SetFilesSoftQuota(v int64)`

SetFilesSoftQuota sets FilesSoftQuota field to given value.

### HasFilesSoftQuota

`func (o *DfsCSIDirectoryReqQuota) HasFilesSoftQuota() bool`

HasFilesSoftQuota returns a boolean if a field has been set.

### GetFilesSuggestQuota

`func (o *DfsCSIDirectoryReqQuota) GetFilesSuggestQuota() int64`

GetFilesSuggestQuota returns the FilesSuggestQuota field if non-nil, zero value otherwise.

### GetFilesSuggestQuotaOk

`func (o *DfsCSIDirectoryReqQuota) GetFilesSuggestQuotaOk() (*int64, bool)`

GetFilesSuggestQuotaOk returns a tuple with the FilesSuggestQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesSuggestQuota

`func (o *DfsCSIDirectoryReqQuota) SetFilesSuggestQuota(v int64)`

SetFilesSuggestQuota sets FilesSuggestQuota field to given value.

### HasFilesSuggestQuota

`func (o *DfsCSIDirectoryReqQuota) HasFilesSuggestQuota() bool`

HasFilesSuggestQuota returns a boolean if a field has been set.

### GetSizeGraceTime

`func (o *DfsCSIDirectoryReqQuota) GetSizeGraceTime() int64`

GetSizeGraceTime returns the SizeGraceTime field if non-nil, zero value otherwise.

### GetSizeGraceTimeOk

`func (o *DfsCSIDirectoryReqQuota) GetSizeGraceTimeOk() (*int64, bool)`

GetSizeGraceTimeOk returns a tuple with the SizeGraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGraceTime

`func (o *DfsCSIDirectoryReqQuota) SetSizeGraceTime(v int64)`

SetSizeGraceTime sets SizeGraceTime field to given value.

### HasSizeGraceTime

`func (o *DfsCSIDirectoryReqQuota) HasSizeGraceTime() bool`

HasSizeGraceTime returns a boolean if a field has been set.

### GetSizeHardQuota

`func (o *DfsCSIDirectoryReqQuota) GetSizeHardQuota() int64`

GetSizeHardQuota returns the SizeHardQuota field if non-nil, zero value otherwise.

### GetSizeHardQuotaOk

`func (o *DfsCSIDirectoryReqQuota) GetSizeHardQuotaOk() (*int64, bool)`

GetSizeHardQuotaOk returns a tuple with the SizeHardQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeHardQuota

`func (o *DfsCSIDirectoryReqQuota) SetSizeHardQuota(v int64)`

SetSizeHardQuota sets SizeHardQuota field to given value.

### HasSizeHardQuota

`func (o *DfsCSIDirectoryReqQuota) HasSizeHardQuota() bool`

HasSizeHardQuota returns a boolean if a field has been set.

### GetSizeRatio

`func (o *DfsCSIDirectoryReqQuota) GetSizeRatio() float32`

GetSizeRatio returns the SizeRatio field if non-nil, zero value otherwise.

### GetSizeRatioOk

`func (o *DfsCSIDirectoryReqQuota) GetSizeRatioOk() (*float32, bool)`

GetSizeRatioOk returns a tuple with the SizeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeRatio

`func (o *DfsCSIDirectoryReqQuota) SetSizeRatio(v float32)`

SetSizeRatio sets SizeRatio field to given value.

### HasSizeRatio

`func (o *DfsCSIDirectoryReqQuota) HasSizeRatio() bool`

HasSizeRatio returns a boolean if a field has been set.

### GetSizeSoftQuota

`func (o *DfsCSIDirectoryReqQuota) GetSizeSoftQuota() int64`

GetSizeSoftQuota returns the SizeSoftQuota field if non-nil, zero value otherwise.

### GetSizeSoftQuotaOk

`func (o *DfsCSIDirectoryReqQuota) GetSizeSoftQuotaOk() (*int64, bool)`

GetSizeSoftQuotaOk returns a tuple with the SizeSoftQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSoftQuota

`func (o *DfsCSIDirectoryReqQuota) SetSizeSoftQuota(v int64)`

SetSizeSoftQuota sets SizeSoftQuota field to given value.

### HasSizeSoftQuota

`func (o *DfsCSIDirectoryReqQuota) HasSizeSoftQuota() bool`

HasSizeSoftQuota returns a boolean if a field has been set.

### GetSizeSuggestQuota

`func (o *DfsCSIDirectoryReqQuota) GetSizeSuggestQuota() int64`

GetSizeSuggestQuota returns the SizeSuggestQuota field if non-nil, zero value otherwise.

### GetSizeSuggestQuotaOk

`func (o *DfsCSIDirectoryReqQuota) GetSizeSuggestQuotaOk() (*int64, bool)`

GetSizeSuggestQuotaOk returns a tuple with the SizeSuggestQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSuggestQuota

`func (o *DfsCSIDirectoryReqQuota) SetSizeSuggestQuota(v int64)`

SetSizeSuggestQuota sets SizeSuggestQuota field to given value.

### HasSizeSuggestQuota

`func (o *DfsCSIDirectoryReqQuota) HasSizeSuggestQuota() bool`

HasSizeSuggestQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


