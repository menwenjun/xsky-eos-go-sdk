# OsdOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveNum** | **int64** | status | 
**ArchiveNum** | **int64** |  | 
**CompoundNum** | **int64** |  | 
**CriticalUsageNum** | **int64** |  | 
**DataNum** | **int64** | role | 
**DoingNum** | **int64** |  | 
**ErrorNum** | **int64** |  | 
**HddTypeNum** | **int64** | disk type | 
**HybridTypeNum** | **int64** |  | 
**IndexNum** | **int64** |  | 
**OfflineNum** | **int64** |  | 
**SsdTypeNum** | **int64** |  | 
**TierCacheNum** | **int64** |  | 
**TierDataNum** | **int64** |  | 
**WarnNum** | **int64** |  | 
**WarnUsageNum** | **int64** | usage | 

## Methods

### NewOsdOverview

`func NewOsdOverview(activeNum int64, archiveNum int64, compoundNum int64, criticalUsageNum int64, dataNum int64, doingNum int64, errorNum int64, hddTypeNum int64, hybridTypeNum int64, indexNum int64, offlineNum int64, ssdTypeNum int64, tierCacheNum int64, tierDataNum int64, warnNum int64, warnUsageNum int64, ) *OsdOverview`

NewOsdOverview instantiates a new OsdOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsdOverviewWithDefaults

`func NewOsdOverviewWithDefaults() *OsdOverview`

NewOsdOverviewWithDefaults instantiates a new OsdOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveNum

`func (o *OsdOverview) GetActiveNum() int64`

GetActiveNum returns the ActiveNum field if non-nil, zero value otherwise.

### GetActiveNumOk

`func (o *OsdOverview) GetActiveNumOk() (*int64, bool)`

GetActiveNumOk returns a tuple with the ActiveNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveNum

`func (o *OsdOverview) SetActiveNum(v int64)`

SetActiveNum sets ActiveNum field to given value.


### GetArchiveNum

`func (o *OsdOverview) GetArchiveNum() int64`

GetArchiveNum returns the ArchiveNum field if non-nil, zero value otherwise.

### GetArchiveNumOk

`func (o *OsdOverview) GetArchiveNumOk() (*int64, bool)`

GetArchiveNumOk returns a tuple with the ArchiveNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveNum

`func (o *OsdOverview) SetArchiveNum(v int64)`

SetArchiveNum sets ArchiveNum field to given value.


### GetCompoundNum

`func (o *OsdOverview) GetCompoundNum() int64`

GetCompoundNum returns the CompoundNum field if non-nil, zero value otherwise.

### GetCompoundNumOk

`func (o *OsdOverview) GetCompoundNumOk() (*int64, bool)`

GetCompoundNumOk returns a tuple with the CompoundNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompoundNum

`func (o *OsdOverview) SetCompoundNum(v int64)`

SetCompoundNum sets CompoundNum field to given value.


### GetCriticalUsageNum

`func (o *OsdOverview) GetCriticalUsageNum() int64`

GetCriticalUsageNum returns the CriticalUsageNum field if non-nil, zero value otherwise.

### GetCriticalUsageNumOk

`func (o *OsdOverview) GetCriticalUsageNumOk() (*int64, bool)`

GetCriticalUsageNumOk returns a tuple with the CriticalUsageNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalUsageNum

`func (o *OsdOverview) SetCriticalUsageNum(v int64)`

SetCriticalUsageNum sets CriticalUsageNum field to given value.


### GetDataNum

`func (o *OsdOverview) GetDataNum() int64`

GetDataNum returns the DataNum field if non-nil, zero value otherwise.

### GetDataNumOk

`func (o *OsdOverview) GetDataNumOk() (*int64, bool)`

GetDataNumOk returns a tuple with the DataNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataNum

`func (o *OsdOverview) SetDataNum(v int64)`

SetDataNum sets DataNum field to given value.


### GetDoingNum

`func (o *OsdOverview) GetDoingNum() int64`

GetDoingNum returns the DoingNum field if non-nil, zero value otherwise.

### GetDoingNumOk

`func (o *OsdOverview) GetDoingNumOk() (*int64, bool)`

GetDoingNumOk returns a tuple with the DoingNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoingNum

`func (o *OsdOverview) SetDoingNum(v int64)`

SetDoingNum sets DoingNum field to given value.


### GetErrorNum

`func (o *OsdOverview) GetErrorNum() int64`

GetErrorNum returns the ErrorNum field if non-nil, zero value otherwise.

### GetErrorNumOk

`func (o *OsdOverview) GetErrorNumOk() (*int64, bool)`

GetErrorNumOk returns a tuple with the ErrorNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNum

`func (o *OsdOverview) SetErrorNum(v int64)`

SetErrorNum sets ErrorNum field to given value.


### GetHddTypeNum

`func (o *OsdOverview) GetHddTypeNum() int64`

GetHddTypeNum returns the HddTypeNum field if non-nil, zero value otherwise.

### GetHddTypeNumOk

`func (o *OsdOverview) GetHddTypeNumOk() (*int64, bool)`

GetHddTypeNumOk returns a tuple with the HddTypeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHddTypeNum

`func (o *OsdOverview) SetHddTypeNum(v int64)`

SetHddTypeNum sets HddTypeNum field to given value.


### GetHybridTypeNum

`func (o *OsdOverview) GetHybridTypeNum() int64`

GetHybridTypeNum returns the HybridTypeNum field if non-nil, zero value otherwise.

### GetHybridTypeNumOk

`func (o *OsdOverview) GetHybridTypeNumOk() (*int64, bool)`

GetHybridTypeNumOk returns a tuple with the HybridTypeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybridTypeNum

`func (o *OsdOverview) SetHybridTypeNum(v int64)`

SetHybridTypeNum sets HybridTypeNum field to given value.


### GetIndexNum

`func (o *OsdOverview) GetIndexNum() int64`

GetIndexNum returns the IndexNum field if non-nil, zero value otherwise.

### GetIndexNumOk

`func (o *OsdOverview) GetIndexNumOk() (*int64, bool)`

GetIndexNumOk returns a tuple with the IndexNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNum

`func (o *OsdOverview) SetIndexNum(v int64)`

SetIndexNum sets IndexNum field to given value.


### GetOfflineNum

`func (o *OsdOverview) GetOfflineNum() int64`

GetOfflineNum returns the OfflineNum field if non-nil, zero value otherwise.

### GetOfflineNumOk

`func (o *OsdOverview) GetOfflineNumOk() (*int64, bool)`

GetOfflineNumOk returns a tuple with the OfflineNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineNum

`func (o *OsdOverview) SetOfflineNum(v int64)`

SetOfflineNum sets OfflineNum field to given value.


### GetSsdTypeNum

`func (o *OsdOverview) GetSsdTypeNum() int64`

GetSsdTypeNum returns the SsdTypeNum field if non-nil, zero value otherwise.

### GetSsdTypeNumOk

`func (o *OsdOverview) GetSsdTypeNumOk() (*int64, bool)`

GetSsdTypeNumOk returns a tuple with the SsdTypeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdTypeNum

`func (o *OsdOverview) SetSsdTypeNum(v int64)`

SetSsdTypeNum sets SsdTypeNum field to given value.


### GetTierCacheNum

`func (o *OsdOverview) GetTierCacheNum() int64`

GetTierCacheNum returns the TierCacheNum field if non-nil, zero value otherwise.

### GetTierCacheNumOk

`func (o *OsdOverview) GetTierCacheNumOk() (*int64, bool)`

GetTierCacheNumOk returns a tuple with the TierCacheNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierCacheNum

`func (o *OsdOverview) SetTierCacheNum(v int64)`

SetTierCacheNum sets TierCacheNum field to given value.


### GetTierDataNum

`func (o *OsdOverview) GetTierDataNum() int64`

GetTierDataNum returns the TierDataNum field if non-nil, zero value otherwise.

### GetTierDataNumOk

`func (o *OsdOverview) GetTierDataNumOk() (*int64, bool)`

GetTierDataNumOk returns a tuple with the TierDataNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierDataNum

`func (o *OsdOverview) SetTierDataNum(v int64)`

SetTierDataNum sets TierDataNum field to given value.


### GetWarnNum

`func (o *OsdOverview) GetWarnNum() int64`

GetWarnNum returns the WarnNum field if non-nil, zero value otherwise.

### GetWarnNumOk

`func (o *OsdOverview) GetWarnNumOk() (*int64, bool)`

GetWarnNumOk returns a tuple with the WarnNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnNum

`func (o *OsdOverview) SetWarnNum(v int64)`

SetWarnNum sets WarnNum field to given value.


### GetWarnUsageNum

`func (o *OsdOverview) GetWarnUsageNum() int64`

GetWarnUsageNum returns the WarnUsageNum field if non-nil, zero value otherwise.

### GetWarnUsageNumOk

`func (o *OsdOverview) GetWarnUsageNumOk() (*int64, bool)`

GetWarnUsageNumOk returns a tuple with the WarnUsageNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnUsageNum

`func (o *OsdOverview) SetWarnUsageNum(v int64)`

SetWarnUsageNum sets WarnUsageNum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


