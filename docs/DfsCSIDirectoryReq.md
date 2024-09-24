# DfsCSIDirectoryReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfsDirectory** | [**DfsCSIDirectoryReqDirectory**](DfsCSIDirectoryReqDirectory.md) |  | 
**DfsQuota** | Pointer to [**DfsCSIDirectoryReqQuota**](DfsCSIDirectoryReqQuota.md) |  | [optional] 

## Methods

### NewDfsCSIDirectoryReq

`func NewDfsCSIDirectoryReq(dfsDirectory DfsCSIDirectoryReqDirectory, ) *DfsCSIDirectoryReq`

NewDfsCSIDirectoryReq instantiates a new DfsCSIDirectoryReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsCSIDirectoryReqWithDefaults

`func NewDfsCSIDirectoryReqWithDefaults() *DfsCSIDirectoryReq`

NewDfsCSIDirectoryReqWithDefaults instantiates a new DfsCSIDirectoryReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfsDirectory

`func (o *DfsCSIDirectoryReq) GetDfsDirectory() DfsCSIDirectoryReqDirectory`

GetDfsDirectory returns the DfsDirectory field if non-nil, zero value otherwise.

### GetDfsDirectoryOk

`func (o *DfsCSIDirectoryReq) GetDfsDirectoryOk() (*DfsCSIDirectoryReqDirectory, bool)`

GetDfsDirectoryOk returns a tuple with the DfsDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsDirectory

`func (o *DfsCSIDirectoryReq) SetDfsDirectory(v DfsCSIDirectoryReqDirectory)`

SetDfsDirectory sets DfsDirectory field to given value.


### GetDfsQuota

`func (o *DfsCSIDirectoryReq) GetDfsQuota() DfsCSIDirectoryReqQuota`

GetDfsQuota returns the DfsQuota field if non-nil, zero value otherwise.

### GetDfsQuotaOk

`func (o *DfsCSIDirectoryReq) GetDfsQuotaOk() (*DfsCSIDirectoryReqQuota, bool)`

GetDfsQuotaOk returns a tuple with the DfsQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsQuota

`func (o *DfsCSIDirectoryReq) SetDfsQuota(v DfsCSIDirectoryReqQuota)`

SetDfsQuota sets DfsQuota field to given value.

### HasDfsQuota

`func (o *DfsCSIDirectoryReq) HasDfsQuota() bool`

HasDfsQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


