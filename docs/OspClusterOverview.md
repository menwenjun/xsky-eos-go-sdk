# OspClusterOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to **int64** |  | [optional] 
**BucketsInRecycle** | Pointer to **int64** |  | [optional] 
**BuildingDataBackends** | Pointer to **int64** |  | [optional] 
**BuildingErrorDataBackends** | Pointer to **int64** |  | [optional] 
**BuildingErrorZones** | Pointer to **int64** |  | [optional] 
**BuildingS3LbGroups** | Pointer to **int64** |  | [optional] 
**BuildingZones** | Pointer to **int64** |  | [optional] 
**DataBackends** | Pointer to **int64** |  | [optional] 
**DataBytes** | Pointer to **int64** |  | [optional] 
**DeletingBuckets** | Pointer to **int64** |  | [optional] 
**DeletingDataBackends** | Pointer to **int64** |  | [optional] 
**DeletingErrorBuckets** | Pointer to **int64** |  | [optional] 
**DeletingErrorDataBackends** | Pointer to **int64** |  | [optional] 
**DeletingErrorStoragePolicies** | Pointer to **int64** |  | [optional] 
**DeletingErrorUsers** | Pointer to **int64** |  | [optional] 
**DeletingErrorZones** | Pointer to **int64** |  | [optional] 
**DeletingS3LbGroups** | Pointer to **int64** |  | [optional] 
**DeletingStoragePolicies** | Pointer to **int64** |  | [optional] 
**DeletingUsers** | Pointer to **int64** |  | [optional] 
**DeletingZones** | Pointer to **int64** |  | [optional] 
**ErrorDataBackends** | Pointer to **int64** |  | [optional] 
**ErrorS3LbGroups** | Pointer to **int64** |  | [optional] 
**ErrorZones** | Pointer to **int64** |  | [optional] 
**InstallingHosts** | Pointer to **int64** |  | [optional] 
**NormalHosts** | Pointer to **int64** |  | [optional] 
**ObjectNum** | Pointer to **int64** |  | [optional] 
**OfflineHosts** | Pointer to **int64** |  | [optional] 
**OperationFailedHosts** | Pointer to **int64** |  | [optional] 
**PartitionCreatingDisks** | Pointer to **int64** |  | [optional] 
**S3LbGroups** | Pointer to **int64** |  | [optional] 
**StoragePolicies** | Pointer to **int64** |  | [optional] 
**SystemMetadataClusterStatus** | Pointer to **string** |  | [optional] 
**Users** | Pointer to **int64** |  | [optional] 
**WarningHosts** | Pointer to **int64** |  | [optional] 
**WarningZones** | Pointer to **int64** |  | [optional] 
**Zones** | Pointer to **int64** |  | [optional] 

## Methods

### NewOspClusterOverview

`func NewOspClusterOverview() *OspClusterOverview`

NewOspClusterOverview instantiates a new OspClusterOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspClusterOverviewWithDefaults

`func NewOspClusterOverviewWithDefaults() *OspClusterOverview`

NewOspClusterOverviewWithDefaults instantiates a new OspClusterOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *OspClusterOverview) GetBuckets() int64`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *OspClusterOverview) GetBucketsOk() (*int64, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *OspClusterOverview) SetBuckets(v int64)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *OspClusterOverview) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetBucketsInRecycle

`func (o *OspClusterOverview) GetBucketsInRecycle() int64`

GetBucketsInRecycle returns the BucketsInRecycle field if non-nil, zero value otherwise.

### GetBucketsInRecycleOk

`func (o *OspClusterOverview) GetBucketsInRecycleOk() (*int64, bool)`

GetBucketsInRecycleOk returns a tuple with the BucketsInRecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketsInRecycle

`func (o *OspClusterOverview) SetBucketsInRecycle(v int64)`

SetBucketsInRecycle sets BucketsInRecycle field to given value.

### HasBucketsInRecycle

`func (o *OspClusterOverview) HasBucketsInRecycle() bool`

HasBucketsInRecycle returns a boolean if a field has been set.

### GetBuildingDataBackends

`func (o *OspClusterOverview) GetBuildingDataBackends() int64`

GetBuildingDataBackends returns the BuildingDataBackends field if non-nil, zero value otherwise.

### GetBuildingDataBackendsOk

`func (o *OspClusterOverview) GetBuildingDataBackendsOk() (*int64, bool)`

GetBuildingDataBackendsOk returns a tuple with the BuildingDataBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingDataBackends

`func (o *OspClusterOverview) SetBuildingDataBackends(v int64)`

SetBuildingDataBackends sets BuildingDataBackends field to given value.

### HasBuildingDataBackends

`func (o *OspClusterOverview) HasBuildingDataBackends() bool`

HasBuildingDataBackends returns a boolean if a field has been set.

### GetBuildingErrorDataBackends

`func (o *OspClusterOverview) GetBuildingErrorDataBackends() int64`

GetBuildingErrorDataBackends returns the BuildingErrorDataBackends field if non-nil, zero value otherwise.

### GetBuildingErrorDataBackendsOk

`func (o *OspClusterOverview) GetBuildingErrorDataBackendsOk() (*int64, bool)`

GetBuildingErrorDataBackendsOk returns a tuple with the BuildingErrorDataBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingErrorDataBackends

`func (o *OspClusterOverview) SetBuildingErrorDataBackends(v int64)`

SetBuildingErrorDataBackends sets BuildingErrorDataBackends field to given value.

### HasBuildingErrorDataBackends

`func (o *OspClusterOverview) HasBuildingErrorDataBackends() bool`

HasBuildingErrorDataBackends returns a boolean if a field has been set.

### GetBuildingErrorZones

`func (o *OspClusterOverview) GetBuildingErrorZones() int64`

GetBuildingErrorZones returns the BuildingErrorZones field if non-nil, zero value otherwise.

### GetBuildingErrorZonesOk

`func (o *OspClusterOverview) GetBuildingErrorZonesOk() (*int64, bool)`

GetBuildingErrorZonesOk returns a tuple with the BuildingErrorZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingErrorZones

`func (o *OspClusterOverview) SetBuildingErrorZones(v int64)`

SetBuildingErrorZones sets BuildingErrorZones field to given value.

### HasBuildingErrorZones

`func (o *OspClusterOverview) HasBuildingErrorZones() bool`

HasBuildingErrorZones returns a boolean if a field has been set.

### GetBuildingS3LbGroups

`func (o *OspClusterOverview) GetBuildingS3LbGroups() int64`

GetBuildingS3LbGroups returns the BuildingS3LbGroups field if non-nil, zero value otherwise.

### GetBuildingS3LbGroupsOk

`func (o *OspClusterOverview) GetBuildingS3LbGroupsOk() (*int64, bool)`

GetBuildingS3LbGroupsOk returns a tuple with the BuildingS3LbGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingS3LbGroups

`func (o *OspClusterOverview) SetBuildingS3LbGroups(v int64)`

SetBuildingS3LbGroups sets BuildingS3LbGroups field to given value.

### HasBuildingS3LbGroups

`func (o *OspClusterOverview) HasBuildingS3LbGroups() bool`

HasBuildingS3LbGroups returns a boolean if a field has been set.

### GetBuildingZones

`func (o *OspClusterOverview) GetBuildingZones() int64`

GetBuildingZones returns the BuildingZones field if non-nil, zero value otherwise.

### GetBuildingZonesOk

`func (o *OspClusterOverview) GetBuildingZonesOk() (*int64, bool)`

GetBuildingZonesOk returns a tuple with the BuildingZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingZones

`func (o *OspClusterOverview) SetBuildingZones(v int64)`

SetBuildingZones sets BuildingZones field to given value.

### HasBuildingZones

`func (o *OspClusterOverview) HasBuildingZones() bool`

HasBuildingZones returns a boolean if a field has been set.

### GetDataBackends

`func (o *OspClusterOverview) GetDataBackends() int64`

GetDataBackends returns the DataBackends field if non-nil, zero value otherwise.

### GetDataBackendsOk

`func (o *OspClusterOverview) GetDataBackendsOk() (*int64, bool)`

GetDataBackendsOk returns a tuple with the DataBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBackends

`func (o *OspClusterOverview) SetDataBackends(v int64)`

SetDataBackends sets DataBackends field to given value.

### HasDataBackends

`func (o *OspClusterOverview) HasDataBackends() bool`

HasDataBackends returns a boolean if a field has been set.

### GetDataBytes

`func (o *OspClusterOverview) GetDataBytes() int64`

GetDataBytes returns the DataBytes field if non-nil, zero value otherwise.

### GetDataBytesOk

`func (o *OspClusterOverview) GetDataBytesOk() (*int64, bool)`

GetDataBytesOk returns a tuple with the DataBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBytes

`func (o *OspClusterOverview) SetDataBytes(v int64)`

SetDataBytes sets DataBytes field to given value.

### HasDataBytes

`func (o *OspClusterOverview) HasDataBytes() bool`

HasDataBytes returns a boolean if a field has been set.

### GetDeletingBuckets

`func (o *OspClusterOverview) GetDeletingBuckets() int64`

GetDeletingBuckets returns the DeletingBuckets field if non-nil, zero value otherwise.

### GetDeletingBucketsOk

`func (o *OspClusterOverview) GetDeletingBucketsOk() (*int64, bool)`

GetDeletingBucketsOk returns a tuple with the DeletingBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletingBuckets

`func (o *OspClusterOverview) SetDeletingBuckets(v int64)`

SetDeletingBuckets sets DeletingBuckets field to given value.

### HasDeletingBuckets

`func (o *OspClusterOverview) HasDeletingBuckets() bool`

HasDeletingBuckets returns a boolean if a field has been set.

### GetDeletingDataBackends

`func (o *OspClusterOverview) GetDeletingDataBackends() int64`

GetDeletingDataBackends returns the DeletingDataBackends field if non-nil, zero value otherwise.

### GetDeletingDataBackendsOk

`func (o *OspClusterOverview) GetDeletingDataBackendsOk() (*int64, bool)`

GetDeletingDataBackendsOk returns a tuple with the DeletingDataBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletingDataBackends

`func (o *OspClusterOverview) SetDeletingDataBackends(v int64)`

SetDeletingDataBackends sets DeletingDataBackends field to given value.

### HasDeletingDataBackends

`func (o *OspClusterOverview) HasDeletingDataBackends() bool`

HasDeletingDataBackends returns a boolean if a field has been set.

### GetDeletingErrorBuckets

`func (o *OspClusterOverview) GetDeletingErrorBuckets() int64`

GetDeletingErrorBuckets returns the DeletingErrorBuckets field if non-nil, zero value otherwise.

### GetDeletingErrorBucketsOk

`func (o *OspClusterOverview) GetDeletingErrorBucketsOk() (*int64, bool)`

GetDeletingErrorBucketsOk returns a tuple with the DeletingErrorBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletingErrorBuckets

`func (o *OspClusterOverview) SetDeletingErrorBuckets(v int64)`

SetDeletingErrorBuckets sets DeletingErrorBuckets field to given value.

### HasDeletingErrorBuckets

`func (o *OspClusterOverview) HasDeletingErrorBuckets() bool`

HasDeletingErrorBuckets returns a boolean if a field has been set.

### GetDeletingErrorDataBackends

`func (o *OspClusterOverview) GetDeletingErrorDataBackends() int64`

GetDeletingErrorDataBackends returns the DeletingErrorDataBackends field if non-nil, zero value otherwise.

### GetDeletingErrorDataBackendsOk

`func (o *OspClusterOverview) GetDeletingErrorDataBackendsOk() (*int64, bool)`

GetDeletingErrorDataBackendsOk returns a tuple with the DeletingErrorDataBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletingErrorDataBackends

`func (o *OspClusterOverview) SetDeletingErrorDataBackends(v int64)`

SetDeletingErrorDataBackends sets DeletingErrorDataBackends field to given value.

### HasDeletingErrorDataBackends

`func (o *OspClusterOverview) HasDeletingErrorDataBackends() bool`

HasDeletingErrorDataBackends returns a boolean if a field has been set.

### GetDeletingErrorStoragePolicies

`func (o *OspClusterOverview) GetDeletingErrorStoragePolicies() int64`

GetDeletingErrorStoragePolicies returns the DeletingErrorStoragePolicies field if non-nil, zero value otherwise.

### GetDeletingErrorStoragePoliciesOk

`func (o *OspClusterOverview) GetDeletingErrorStoragePoliciesOk() (*int64, bool)`

GetDeletingErrorStoragePoliciesOk returns a tuple with the DeletingErrorStoragePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletingErrorStoragePolicies

`func (o *OspClusterOverview) SetDeletingErrorStoragePolicies(v int64)`

SetDeletingErrorStoragePolicies sets DeletingErrorStoragePolicies field to given value.

### HasDeletingErrorStoragePolicies

`func (o *OspClusterOverview) HasDeletingErrorStoragePolicies() bool`

HasDeletingErrorStoragePolicies returns a boolean if a field has been set.

### GetDeletingErrorUsers

`func (o *OspClusterOverview) GetDeletingErrorUsers() int64`

GetDeletingErrorUsers returns the DeletingErrorUsers field if non-nil, zero value otherwise.

### GetDeletingErrorUsersOk

`func (o *OspClusterOverview) GetDeletingErrorUsersOk() (*int64, bool)`

GetDeletingErrorUsersOk returns a tuple with the DeletingErrorUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletingErrorUsers

`func (o *OspClusterOverview) SetDeletingErrorUsers(v int64)`

SetDeletingErrorUsers sets DeletingErrorUsers field to given value.

### HasDeletingErrorUsers

`func (o *OspClusterOverview) HasDeletingErrorUsers() bool`

HasDeletingErrorUsers returns a boolean if a field has been set.

### GetDeletingErrorZones

`func (o *OspClusterOverview) GetDeletingErrorZones() int64`

GetDeletingErrorZones returns the DeletingErrorZones field if non-nil, zero value otherwise.

### GetDeletingErrorZonesOk

`func (o *OspClusterOverview) GetDeletingErrorZonesOk() (*int64, bool)`

GetDeletingErrorZonesOk returns a tuple with the DeletingErrorZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletingErrorZones

`func (o *OspClusterOverview) SetDeletingErrorZones(v int64)`

SetDeletingErrorZones sets DeletingErrorZones field to given value.

### HasDeletingErrorZones

`func (o *OspClusterOverview) HasDeletingErrorZones() bool`

HasDeletingErrorZones returns a boolean if a field has been set.

### GetDeletingS3LbGroups

`func (o *OspClusterOverview) GetDeletingS3LbGroups() int64`

GetDeletingS3LbGroups returns the DeletingS3LbGroups field if non-nil, zero value otherwise.

### GetDeletingS3LbGroupsOk

`func (o *OspClusterOverview) GetDeletingS3LbGroupsOk() (*int64, bool)`

GetDeletingS3LbGroupsOk returns a tuple with the DeletingS3LbGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletingS3LbGroups

`func (o *OspClusterOverview) SetDeletingS3LbGroups(v int64)`

SetDeletingS3LbGroups sets DeletingS3LbGroups field to given value.

### HasDeletingS3LbGroups

`func (o *OspClusterOverview) HasDeletingS3LbGroups() bool`

HasDeletingS3LbGroups returns a boolean if a field has been set.

### GetDeletingStoragePolicies

`func (o *OspClusterOverview) GetDeletingStoragePolicies() int64`

GetDeletingStoragePolicies returns the DeletingStoragePolicies field if non-nil, zero value otherwise.

### GetDeletingStoragePoliciesOk

`func (o *OspClusterOverview) GetDeletingStoragePoliciesOk() (*int64, bool)`

GetDeletingStoragePoliciesOk returns a tuple with the DeletingStoragePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletingStoragePolicies

`func (o *OspClusterOverview) SetDeletingStoragePolicies(v int64)`

SetDeletingStoragePolicies sets DeletingStoragePolicies field to given value.

### HasDeletingStoragePolicies

`func (o *OspClusterOverview) HasDeletingStoragePolicies() bool`

HasDeletingStoragePolicies returns a boolean if a field has been set.

### GetDeletingUsers

`func (o *OspClusterOverview) GetDeletingUsers() int64`

GetDeletingUsers returns the DeletingUsers field if non-nil, zero value otherwise.

### GetDeletingUsersOk

`func (o *OspClusterOverview) GetDeletingUsersOk() (*int64, bool)`

GetDeletingUsersOk returns a tuple with the DeletingUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletingUsers

`func (o *OspClusterOverview) SetDeletingUsers(v int64)`

SetDeletingUsers sets DeletingUsers field to given value.

### HasDeletingUsers

`func (o *OspClusterOverview) HasDeletingUsers() bool`

HasDeletingUsers returns a boolean if a field has been set.

### GetDeletingZones

`func (o *OspClusterOverview) GetDeletingZones() int64`

GetDeletingZones returns the DeletingZones field if non-nil, zero value otherwise.

### GetDeletingZonesOk

`func (o *OspClusterOverview) GetDeletingZonesOk() (*int64, bool)`

GetDeletingZonesOk returns a tuple with the DeletingZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletingZones

`func (o *OspClusterOverview) SetDeletingZones(v int64)`

SetDeletingZones sets DeletingZones field to given value.

### HasDeletingZones

`func (o *OspClusterOverview) HasDeletingZones() bool`

HasDeletingZones returns a boolean if a field has been set.

### GetErrorDataBackends

`func (o *OspClusterOverview) GetErrorDataBackends() int64`

GetErrorDataBackends returns the ErrorDataBackends field if non-nil, zero value otherwise.

### GetErrorDataBackendsOk

`func (o *OspClusterOverview) GetErrorDataBackendsOk() (*int64, bool)`

GetErrorDataBackendsOk returns a tuple with the ErrorDataBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDataBackends

`func (o *OspClusterOverview) SetErrorDataBackends(v int64)`

SetErrorDataBackends sets ErrorDataBackends field to given value.

### HasErrorDataBackends

`func (o *OspClusterOverview) HasErrorDataBackends() bool`

HasErrorDataBackends returns a boolean if a field has been set.

### GetErrorS3LbGroups

`func (o *OspClusterOverview) GetErrorS3LbGroups() int64`

GetErrorS3LbGroups returns the ErrorS3LbGroups field if non-nil, zero value otherwise.

### GetErrorS3LbGroupsOk

`func (o *OspClusterOverview) GetErrorS3LbGroupsOk() (*int64, bool)`

GetErrorS3LbGroupsOk returns a tuple with the ErrorS3LbGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorS3LbGroups

`func (o *OspClusterOverview) SetErrorS3LbGroups(v int64)`

SetErrorS3LbGroups sets ErrorS3LbGroups field to given value.

### HasErrorS3LbGroups

`func (o *OspClusterOverview) HasErrorS3LbGroups() bool`

HasErrorS3LbGroups returns a boolean if a field has been set.

### GetErrorZones

`func (o *OspClusterOverview) GetErrorZones() int64`

GetErrorZones returns the ErrorZones field if non-nil, zero value otherwise.

### GetErrorZonesOk

`func (o *OspClusterOverview) GetErrorZonesOk() (*int64, bool)`

GetErrorZonesOk returns a tuple with the ErrorZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorZones

`func (o *OspClusterOverview) SetErrorZones(v int64)`

SetErrorZones sets ErrorZones field to given value.

### HasErrorZones

`func (o *OspClusterOverview) HasErrorZones() bool`

HasErrorZones returns a boolean if a field has been set.

### GetInstallingHosts

`func (o *OspClusterOverview) GetInstallingHosts() int64`

GetInstallingHosts returns the InstallingHosts field if non-nil, zero value otherwise.

### GetInstallingHostsOk

`func (o *OspClusterOverview) GetInstallingHostsOk() (*int64, bool)`

GetInstallingHostsOk returns a tuple with the InstallingHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallingHosts

`func (o *OspClusterOverview) SetInstallingHosts(v int64)`

SetInstallingHosts sets InstallingHosts field to given value.

### HasInstallingHosts

`func (o *OspClusterOverview) HasInstallingHosts() bool`

HasInstallingHosts returns a boolean if a field has been set.

### GetNormalHosts

`func (o *OspClusterOverview) GetNormalHosts() int64`

GetNormalHosts returns the NormalHosts field if non-nil, zero value otherwise.

### GetNormalHostsOk

`func (o *OspClusterOverview) GetNormalHostsOk() (*int64, bool)`

GetNormalHostsOk returns a tuple with the NormalHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalHosts

`func (o *OspClusterOverview) SetNormalHosts(v int64)`

SetNormalHosts sets NormalHosts field to given value.

### HasNormalHosts

`func (o *OspClusterOverview) HasNormalHosts() bool`

HasNormalHosts returns a boolean if a field has been set.

### GetObjectNum

`func (o *OspClusterOverview) GetObjectNum() int64`

GetObjectNum returns the ObjectNum field if non-nil, zero value otherwise.

### GetObjectNumOk

`func (o *OspClusterOverview) GetObjectNumOk() (*int64, bool)`

GetObjectNumOk returns a tuple with the ObjectNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectNum

`func (o *OspClusterOverview) SetObjectNum(v int64)`

SetObjectNum sets ObjectNum field to given value.

### HasObjectNum

`func (o *OspClusterOverview) HasObjectNum() bool`

HasObjectNum returns a boolean if a field has been set.

### GetOfflineHosts

`func (o *OspClusterOverview) GetOfflineHosts() int64`

GetOfflineHosts returns the OfflineHosts field if non-nil, zero value otherwise.

### GetOfflineHostsOk

`func (o *OspClusterOverview) GetOfflineHostsOk() (*int64, bool)`

GetOfflineHostsOk returns a tuple with the OfflineHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineHosts

`func (o *OspClusterOverview) SetOfflineHosts(v int64)`

SetOfflineHosts sets OfflineHosts field to given value.

### HasOfflineHosts

`func (o *OspClusterOverview) HasOfflineHosts() bool`

HasOfflineHosts returns a boolean if a field has been set.

### GetOperationFailedHosts

`func (o *OspClusterOverview) GetOperationFailedHosts() int64`

GetOperationFailedHosts returns the OperationFailedHosts field if non-nil, zero value otherwise.

### GetOperationFailedHostsOk

`func (o *OspClusterOverview) GetOperationFailedHostsOk() (*int64, bool)`

GetOperationFailedHostsOk returns a tuple with the OperationFailedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationFailedHosts

`func (o *OspClusterOverview) SetOperationFailedHosts(v int64)`

SetOperationFailedHosts sets OperationFailedHosts field to given value.

### HasOperationFailedHosts

`func (o *OspClusterOverview) HasOperationFailedHosts() bool`

HasOperationFailedHosts returns a boolean if a field has been set.

### GetPartitionCreatingDisks

`func (o *OspClusterOverview) GetPartitionCreatingDisks() int64`

GetPartitionCreatingDisks returns the PartitionCreatingDisks field if non-nil, zero value otherwise.

### GetPartitionCreatingDisksOk

`func (o *OspClusterOverview) GetPartitionCreatingDisksOk() (*int64, bool)`

GetPartitionCreatingDisksOk returns a tuple with the PartitionCreatingDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionCreatingDisks

`func (o *OspClusterOverview) SetPartitionCreatingDisks(v int64)`

SetPartitionCreatingDisks sets PartitionCreatingDisks field to given value.

### HasPartitionCreatingDisks

`func (o *OspClusterOverview) HasPartitionCreatingDisks() bool`

HasPartitionCreatingDisks returns a boolean if a field has been set.

### GetS3LbGroups

`func (o *OspClusterOverview) GetS3LbGroups() int64`

GetS3LbGroups returns the S3LbGroups field if non-nil, zero value otherwise.

### GetS3LbGroupsOk

`func (o *OspClusterOverview) GetS3LbGroupsOk() (*int64, bool)`

GetS3LbGroupsOk returns a tuple with the S3LbGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3LbGroups

`func (o *OspClusterOverview) SetS3LbGroups(v int64)`

SetS3LbGroups sets S3LbGroups field to given value.

### HasS3LbGroups

`func (o *OspClusterOverview) HasS3LbGroups() bool`

HasS3LbGroups returns a boolean if a field has been set.

### GetStoragePolicies

`func (o *OspClusterOverview) GetStoragePolicies() int64`

GetStoragePolicies returns the StoragePolicies field if non-nil, zero value otherwise.

### GetStoragePoliciesOk

`func (o *OspClusterOverview) GetStoragePoliciesOk() (*int64, bool)`

GetStoragePoliciesOk returns a tuple with the StoragePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicies

`func (o *OspClusterOverview) SetStoragePolicies(v int64)`

SetStoragePolicies sets StoragePolicies field to given value.

### HasStoragePolicies

`func (o *OspClusterOverview) HasStoragePolicies() bool`

HasStoragePolicies returns a boolean if a field has been set.

### GetSystemMetadataClusterStatus

`func (o *OspClusterOverview) GetSystemMetadataClusterStatus() string`

GetSystemMetadataClusterStatus returns the SystemMetadataClusterStatus field if non-nil, zero value otherwise.

### GetSystemMetadataClusterStatusOk

`func (o *OspClusterOverview) GetSystemMetadataClusterStatusOk() (*string, bool)`

GetSystemMetadataClusterStatusOk returns a tuple with the SystemMetadataClusterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMetadataClusterStatus

`func (o *OspClusterOverview) SetSystemMetadataClusterStatus(v string)`

SetSystemMetadataClusterStatus sets SystemMetadataClusterStatus field to given value.

### HasSystemMetadataClusterStatus

`func (o *OspClusterOverview) HasSystemMetadataClusterStatus() bool`

HasSystemMetadataClusterStatus returns a boolean if a field has been set.

### GetUsers

`func (o *OspClusterOverview) GetUsers() int64`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *OspClusterOverview) GetUsersOk() (*int64, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *OspClusterOverview) SetUsers(v int64)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *OspClusterOverview) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetWarningHosts

`func (o *OspClusterOverview) GetWarningHosts() int64`

GetWarningHosts returns the WarningHosts field if non-nil, zero value otherwise.

### GetWarningHostsOk

`func (o *OspClusterOverview) GetWarningHostsOk() (*int64, bool)`

GetWarningHostsOk returns a tuple with the WarningHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningHosts

`func (o *OspClusterOverview) SetWarningHosts(v int64)`

SetWarningHosts sets WarningHosts field to given value.

### HasWarningHosts

`func (o *OspClusterOverview) HasWarningHosts() bool`

HasWarningHosts returns a boolean if a field has been set.

### GetWarningZones

`func (o *OspClusterOverview) GetWarningZones() int64`

GetWarningZones returns the WarningZones field if non-nil, zero value otherwise.

### GetWarningZonesOk

`func (o *OspClusterOverview) GetWarningZonesOk() (*int64, bool)`

GetWarningZonesOk returns a tuple with the WarningZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningZones

`func (o *OspClusterOverview) SetWarningZones(v int64)`

SetWarningZones sets WarningZones field to given value.

### HasWarningZones

`func (o *OspClusterOverview) HasWarningZones() bool`

HasWarningZones returns a boolean if a field has been set.

### GetZones

`func (o *OspClusterOverview) GetZones() int64`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *OspClusterOverview) GetZonesOk() (*int64, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *OspClusterOverview) SetZones(v int64)`

SetZones sets Zones field to given value.

### HasZones

`func (o *OspClusterOverview) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


