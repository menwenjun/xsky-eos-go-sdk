# MetadataClusterStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableDataKbyte** | Pointer to **int64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DeleteLatencyUs** | Pointer to **float64** |  | [optional] 
**DeleteOps** | Pointer to **float64** |  | [optional] 
**GcDataKbyte** | Pointer to **int64** |  | [optional] 
**ListLatencyUs** | Pointer to **float64** |  | [optional] 
**ListOps** | Pointer to **float64** |  | [optional] 
**NumAvailableInodes** | Pointer to **int64** |  | [optional] 
**NumGcInodes** | Pointer to **int64** |  | [optional] 
**NumInodes** | Pointer to **int64** |  | [optional] 
**NumUsedInodes** | Pointer to **int64** |  | [optional] 
**ReadLatencyUs** | Pointer to **float64** |  | [optional] 
**ReadOps** | Pointer to **float64** | primary metadata service stat | [optional] 
**RecordedDataKbyte** | Pointer to **int64** | data kbyte used by user, different from UsedDataKbyte | [optional] 
**RecoveryDone** | Pointer to **int64** |  | [optional] 
**RecoveryLeftSecond** | Pointer to **float64** |  | [optional] 
**RecoveryOps** | Pointer to **float64** |  | [optional] 
**RecoveryTotal** | Pointer to **int64** |  | [optional] 
**TotalDataKbyte** | Pointer to **int64** |  | [optional] 
**TotalKbyte** | Pointer to **int64** |  | [optional] 
**TrashFiles** | Pointer to **int64** |  | [optional] 
**TrashKbyte** | Pointer to **int64** |  | [optional] 
**UsedDataKbyte** | Pointer to **int64** |  | [optional] 
**UsedKbyte** | Pointer to **int64** |  | [optional] 
**WriteLatencyUs** | Pointer to **float64** |  | [optional] 
**WriteOps** | Pointer to **float64** |  | [optional] 

## Methods

### NewMetadataClusterStat

`func NewMetadataClusterStat() *MetadataClusterStat`

NewMetadataClusterStat instantiates a new MetadataClusterStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataClusterStatWithDefaults

`func NewMetadataClusterStatWithDefaults() *MetadataClusterStat`

NewMetadataClusterStatWithDefaults instantiates a new MetadataClusterStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableDataKbyte

`func (o *MetadataClusterStat) GetAvailableDataKbyte() int64`

GetAvailableDataKbyte returns the AvailableDataKbyte field if non-nil, zero value otherwise.

### GetAvailableDataKbyteOk

`func (o *MetadataClusterStat) GetAvailableDataKbyteOk() (*int64, bool)`

GetAvailableDataKbyteOk returns a tuple with the AvailableDataKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableDataKbyte

`func (o *MetadataClusterStat) SetAvailableDataKbyte(v int64)`

SetAvailableDataKbyte sets AvailableDataKbyte field to given value.

### HasAvailableDataKbyte

`func (o *MetadataClusterStat) HasAvailableDataKbyte() bool`

HasAvailableDataKbyte returns a boolean if a field has been set.

### GetCreate

`func (o *MetadataClusterStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *MetadataClusterStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *MetadataClusterStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *MetadataClusterStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDeleteLatencyUs

`func (o *MetadataClusterStat) GetDeleteLatencyUs() float64`

GetDeleteLatencyUs returns the DeleteLatencyUs field if non-nil, zero value otherwise.

### GetDeleteLatencyUsOk

`func (o *MetadataClusterStat) GetDeleteLatencyUsOk() (*float64, bool)`

GetDeleteLatencyUsOk returns a tuple with the DeleteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteLatencyUs

`func (o *MetadataClusterStat) SetDeleteLatencyUs(v float64)`

SetDeleteLatencyUs sets DeleteLatencyUs field to given value.

### HasDeleteLatencyUs

`func (o *MetadataClusterStat) HasDeleteLatencyUs() bool`

HasDeleteLatencyUs returns a boolean if a field has been set.

### GetDeleteOps

`func (o *MetadataClusterStat) GetDeleteOps() float64`

GetDeleteOps returns the DeleteOps field if non-nil, zero value otherwise.

### GetDeleteOpsOk

`func (o *MetadataClusterStat) GetDeleteOpsOk() (*float64, bool)`

GetDeleteOpsOk returns a tuple with the DeleteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOps

`func (o *MetadataClusterStat) SetDeleteOps(v float64)`

SetDeleteOps sets DeleteOps field to given value.

### HasDeleteOps

`func (o *MetadataClusterStat) HasDeleteOps() bool`

HasDeleteOps returns a boolean if a field has been set.

### GetGcDataKbyte

`func (o *MetadataClusterStat) GetGcDataKbyte() int64`

GetGcDataKbyte returns the GcDataKbyte field if non-nil, zero value otherwise.

### GetGcDataKbyteOk

`func (o *MetadataClusterStat) GetGcDataKbyteOk() (*int64, bool)`

GetGcDataKbyteOk returns a tuple with the GcDataKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcDataKbyte

`func (o *MetadataClusterStat) SetGcDataKbyte(v int64)`

SetGcDataKbyte sets GcDataKbyte field to given value.

### HasGcDataKbyte

`func (o *MetadataClusterStat) HasGcDataKbyte() bool`

HasGcDataKbyte returns a boolean if a field has been set.

### GetListLatencyUs

`func (o *MetadataClusterStat) GetListLatencyUs() float64`

GetListLatencyUs returns the ListLatencyUs field if non-nil, zero value otherwise.

### GetListLatencyUsOk

`func (o *MetadataClusterStat) GetListLatencyUsOk() (*float64, bool)`

GetListLatencyUsOk returns a tuple with the ListLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListLatencyUs

`func (o *MetadataClusterStat) SetListLatencyUs(v float64)`

SetListLatencyUs sets ListLatencyUs field to given value.

### HasListLatencyUs

`func (o *MetadataClusterStat) HasListLatencyUs() bool`

HasListLatencyUs returns a boolean if a field has been set.

### GetListOps

`func (o *MetadataClusterStat) GetListOps() float64`

GetListOps returns the ListOps field if non-nil, zero value otherwise.

### GetListOpsOk

`func (o *MetadataClusterStat) GetListOpsOk() (*float64, bool)`

GetListOpsOk returns a tuple with the ListOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOps

`func (o *MetadataClusterStat) SetListOps(v float64)`

SetListOps sets ListOps field to given value.

### HasListOps

`func (o *MetadataClusterStat) HasListOps() bool`

HasListOps returns a boolean if a field has been set.

### GetNumAvailableInodes

`func (o *MetadataClusterStat) GetNumAvailableInodes() int64`

GetNumAvailableInodes returns the NumAvailableInodes field if non-nil, zero value otherwise.

### GetNumAvailableInodesOk

`func (o *MetadataClusterStat) GetNumAvailableInodesOk() (*int64, bool)`

GetNumAvailableInodesOk returns a tuple with the NumAvailableInodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAvailableInodes

`func (o *MetadataClusterStat) SetNumAvailableInodes(v int64)`

SetNumAvailableInodes sets NumAvailableInodes field to given value.

### HasNumAvailableInodes

`func (o *MetadataClusterStat) HasNumAvailableInodes() bool`

HasNumAvailableInodes returns a boolean if a field has been set.

### GetNumGcInodes

`func (o *MetadataClusterStat) GetNumGcInodes() int64`

GetNumGcInodes returns the NumGcInodes field if non-nil, zero value otherwise.

### GetNumGcInodesOk

`func (o *MetadataClusterStat) GetNumGcInodesOk() (*int64, bool)`

GetNumGcInodesOk returns a tuple with the NumGcInodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumGcInodes

`func (o *MetadataClusterStat) SetNumGcInodes(v int64)`

SetNumGcInodes sets NumGcInodes field to given value.

### HasNumGcInodes

`func (o *MetadataClusterStat) HasNumGcInodes() bool`

HasNumGcInodes returns a boolean if a field has been set.

### GetNumInodes

`func (o *MetadataClusterStat) GetNumInodes() int64`

GetNumInodes returns the NumInodes field if non-nil, zero value otherwise.

### GetNumInodesOk

`func (o *MetadataClusterStat) GetNumInodesOk() (*int64, bool)`

GetNumInodesOk returns a tuple with the NumInodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInodes

`func (o *MetadataClusterStat) SetNumInodes(v int64)`

SetNumInodes sets NumInodes field to given value.

### HasNumInodes

`func (o *MetadataClusterStat) HasNumInodes() bool`

HasNumInodes returns a boolean if a field has been set.

### GetNumUsedInodes

`func (o *MetadataClusterStat) GetNumUsedInodes() int64`

GetNumUsedInodes returns the NumUsedInodes field if non-nil, zero value otherwise.

### GetNumUsedInodesOk

`func (o *MetadataClusterStat) GetNumUsedInodesOk() (*int64, bool)`

GetNumUsedInodesOk returns a tuple with the NumUsedInodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUsedInodes

`func (o *MetadataClusterStat) SetNumUsedInodes(v int64)`

SetNumUsedInodes sets NumUsedInodes field to given value.

### HasNumUsedInodes

`func (o *MetadataClusterStat) HasNumUsedInodes() bool`

HasNumUsedInodes returns a boolean if a field has been set.

### GetReadLatencyUs

`func (o *MetadataClusterStat) GetReadLatencyUs() float64`

GetReadLatencyUs returns the ReadLatencyUs field if non-nil, zero value otherwise.

### GetReadLatencyUsOk

`func (o *MetadataClusterStat) GetReadLatencyUsOk() (*float64, bool)`

GetReadLatencyUsOk returns a tuple with the ReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadLatencyUs

`func (o *MetadataClusterStat) SetReadLatencyUs(v float64)`

SetReadLatencyUs sets ReadLatencyUs field to given value.

### HasReadLatencyUs

`func (o *MetadataClusterStat) HasReadLatencyUs() bool`

HasReadLatencyUs returns a boolean if a field has been set.

### GetReadOps

`func (o *MetadataClusterStat) GetReadOps() float64`

GetReadOps returns the ReadOps field if non-nil, zero value otherwise.

### GetReadOpsOk

`func (o *MetadataClusterStat) GetReadOpsOk() (*float64, bool)`

GetReadOpsOk returns a tuple with the ReadOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOps

`func (o *MetadataClusterStat) SetReadOps(v float64)`

SetReadOps sets ReadOps field to given value.

### HasReadOps

`func (o *MetadataClusterStat) HasReadOps() bool`

HasReadOps returns a boolean if a field has been set.

### GetRecordedDataKbyte

`func (o *MetadataClusterStat) GetRecordedDataKbyte() int64`

GetRecordedDataKbyte returns the RecordedDataKbyte field if non-nil, zero value otherwise.

### GetRecordedDataKbyteOk

`func (o *MetadataClusterStat) GetRecordedDataKbyteOk() (*int64, bool)`

GetRecordedDataKbyteOk returns a tuple with the RecordedDataKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedDataKbyte

`func (o *MetadataClusterStat) SetRecordedDataKbyte(v int64)`

SetRecordedDataKbyte sets RecordedDataKbyte field to given value.

### HasRecordedDataKbyte

`func (o *MetadataClusterStat) HasRecordedDataKbyte() bool`

HasRecordedDataKbyte returns a boolean if a field has been set.

### GetRecoveryDone

`func (o *MetadataClusterStat) GetRecoveryDone() int64`

GetRecoveryDone returns the RecoveryDone field if non-nil, zero value otherwise.

### GetRecoveryDoneOk

`func (o *MetadataClusterStat) GetRecoveryDoneOk() (*int64, bool)`

GetRecoveryDoneOk returns a tuple with the RecoveryDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryDone

`func (o *MetadataClusterStat) SetRecoveryDone(v int64)`

SetRecoveryDone sets RecoveryDone field to given value.

### HasRecoveryDone

`func (o *MetadataClusterStat) HasRecoveryDone() bool`

HasRecoveryDone returns a boolean if a field has been set.

### GetRecoveryLeftSecond

`func (o *MetadataClusterStat) GetRecoveryLeftSecond() float64`

GetRecoveryLeftSecond returns the RecoveryLeftSecond field if non-nil, zero value otherwise.

### GetRecoveryLeftSecondOk

`func (o *MetadataClusterStat) GetRecoveryLeftSecondOk() (*float64, bool)`

GetRecoveryLeftSecondOk returns a tuple with the RecoveryLeftSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryLeftSecond

`func (o *MetadataClusterStat) SetRecoveryLeftSecond(v float64)`

SetRecoveryLeftSecond sets RecoveryLeftSecond field to given value.

### HasRecoveryLeftSecond

`func (o *MetadataClusterStat) HasRecoveryLeftSecond() bool`

HasRecoveryLeftSecond returns a boolean if a field has been set.

### GetRecoveryOps

`func (o *MetadataClusterStat) GetRecoveryOps() float64`

GetRecoveryOps returns the RecoveryOps field if non-nil, zero value otherwise.

### GetRecoveryOpsOk

`func (o *MetadataClusterStat) GetRecoveryOpsOk() (*float64, bool)`

GetRecoveryOpsOk returns a tuple with the RecoveryOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryOps

`func (o *MetadataClusterStat) SetRecoveryOps(v float64)`

SetRecoveryOps sets RecoveryOps field to given value.

### HasRecoveryOps

`func (o *MetadataClusterStat) HasRecoveryOps() bool`

HasRecoveryOps returns a boolean if a field has been set.

### GetRecoveryTotal

`func (o *MetadataClusterStat) GetRecoveryTotal() int64`

GetRecoveryTotal returns the RecoveryTotal field if non-nil, zero value otherwise.

### GetRecoveryTotalOk

`func (o *MetadataClusterStat) GetRecoveryTotalOk() (*int64, bool)`

GetRecoveryTotalOk returns a tuple with the RecoveryTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTotal

`func (o *MetadataClusterStat) SetRecoveryTotal(v int64)`

SetRecoveryTotal sets RecoveryTotal field to given value.

### HasRecoveryTotal

`func (o *MetadataClusterStat) HasRecoveryTotal() bool`

HasRecoveryTotal returns a boolean if a field has been set.

### GetTotalDataKbyte

`func (o *MetadataClusterStat) GetTotalDataKbyte() int64`

GetTotalDataKbyte returns the TotalDataKbyte field if non-nil, zero value otherwise.

### GetTotalDataKbyteOk

`func (o *MetadataClusterStat) GetTotalDataKbyteOk() (*int64, bool)`

GetTotalDataKbyteOk returns a tuple with the TotalDataKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDataKbyte

`func (o *MetadataClusterStat) SetTotalDataKbyte(v int64)`

SetTotalDataKbyte sets TotalDataKbyte field to given value.

### HasTotalDataKbyte

`func (o *MetadataClusterStat) HasTotalDataKbyte() bool`

HasTotalDataKbyte returns a boolean if a field has been set.

### GetTotalKbyte

`func (o *MetadataClusterStat) GetTotalKbyte() int64`

GetTotalKbyte returns the TotalKbyte field if non-nil, zero value otherwise.

### GetTotalKbyteOk

`func (o *MetadataClusterStat) GetTotalKbyteOk() (*int64, bool)`

GetTotalKbyteOk returns a tuple with the TotalKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalKbyte

`func (o *MetadataClusterStat) SetTotalKbyte(v int64)`

SetTotalKbyte sets TotalKbyte field to given value.

### HasTotalKbyte

`func (o *MetadataClusterStat) HasTotalKbyte() bool`

HasTotalKbyte returns a boolean if a field has been set.

### GetTrashFiles

`func (o *MetadataClusterStat) GetTrashFiles() int64`

GetTrashFiles returns the TrashFiles field if non-nil, zero value otherwise.

### GetTrashFilesOk

`func (o *MetadataClusterStat) GetTrashFilesOk() (*int64, bool)`

GetTrashFilesOk returns a tuple with the TrashFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashFiles

`func (o *MetadataClusterStat) SetTrashFiles(v int64)`

SetTrashFiles sets TrashFiles field to given value.

### HasTrashFiles

`func (o *MetadataClusterStat) HasTrashFiles() bool`

HasTrashFiles returns a boolean if a field has been set.

### GetTrashKbyte

`func (o *MetadataClusterStat) GetTrashKbyte() int64`

GetTrashKbyte returns the TrashKbyte field if non-nil, zero value otherwise.

### GetTrashKbyteOk

`func (o *MetadataClusterStat) GetTrashKbyteOk() (*int64, bool)`

GetTrashKbyteOk returns a tuple with the TrashKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashKbyte

`func (o *MetadataClusterStat) SetTrashKbyte(v int64)`

SetTrashKbyte sets TrashKbyte field to given value.

### HasTrashKbyte

`func (o *MetadataClusterStat) HasTrashKbyte() bool`

HasTrashKbyte returns a boolean if a field has been set.

### GetUsedDataKbyte

`func (o *MetadataClusterStat) GetUsedDataKbyte() int64`

GetUsedDataKbyte returns the UsedDataKbyte field if non-nil, zero value otherwise.

### GetUsedDataKbyteOk

`func (o *MetadataClusterStat) GetUsedDataKbyteOk() (*int64, bool)`

GetUsedDataKbyteOk returns a tuple with the UsedDataKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedDataKbyte

`func (o *MetadataClusterStat) SetUsedDataKbyte(v int64)`

SetUsedDataKbyte sets UsedDataKbyte field to given value.

### HasUsedDataKbyte

`func (o *MetadataClusterStat) HasUsedDataKbyte() bool`

HasUsedDataKbyte returns a boolean if a field has been set.

### GetUsedKbyte

`func (o *MetadataClusterStat) GetUsedKbyte() int64`

GetUsedKbyte returns the UsedKbyte field if non-nil, zero value otherwise.

### GetUsedKbyteOk

`func (o *MetadataClusterStat) GetUsedKbyteOk() (*int64, bool)`

GetUsedKbyteOk returns a tuple with the UsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte

`func (o *MetadataClusterStat) SetUsedKbyte(v int64)`

SetUsedKbyte sets UsedKbyte field to given value.

### HasUsedKbyte

`func (o *MetadataClusterStat) HasUsedKbyte() bool`

HasUsedKbyte returns a boolean if a field has been set.

### GetWriteLatencyUs

`func (o *MetadataClusterStat) GetWriteLatencyUs() float64`

GetWriteLatencyUs returns the WriteLatencyUs field if non-nil, zero value otherwise.

### GetWriteLatencyUsOk

`func (o *MetadataClusterStat) GetWriteLatencyUsOk() (*float64, bool)`

GetWriteLatencyUsOk returns a tuple with the WriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLatencyUs

`func (o *MetadataClusterStat) SetWriteLatencyUs(v float64)`

SetWriteLatencyUs sets WriteLatencyUs field to given value.

### HasWriteLatencyUs

`func (o *MetadataClusterStat) HasWriteLatencyUs() bool`

HasWriteLatencyUs returns a boolean if a field has been set.

### GetWriteOps

`func (o *MetadataClusterStat) GetWriteOps() float64`

GetWriteOps returns the WriteOps field if non-nil, zero value otherwise.

### GetWriteOpsOk

`func (o *MetadataClusterStat) GetWriteOpsOk() (*float64, bool)`

GetWriteOpsOk returns a tuple with the WriteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteOps

`func (o *MetadataClusterStat) SetWriteOps(v float64)`

SetWriteOps sets WriteOps field to given value.

### HasWriteOps

`func (o *MetadataClusterStat) HasWriteOps() bool`

HasWriteOps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


