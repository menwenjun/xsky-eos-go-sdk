# DfsStoragePolicyResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfsStoragePolicy** | [**DfsStoragePolicy**](DfsStoragePolicy.md) |  | 
**LinkPathInfo** | Pointer to [**[]PolicyLinkPathInfo**](PolicyLinkPathInfo.md) | link path info | [optional] 

## Methods

### NewDfsStoragePolicyResp

`func NewDfsStoragePolicyResp(dfsStoragePolicy DfsStoragePolicy, ) *DfsStoragePolicyResp`

NewDfsStoragePolicyResp instantiates a new DfsStoragePolicyResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsStoragePolicyRespWithDefaults

`func NewDfsStoragePolicyRespWithDefaults() *DfsStoragePolicyResp`

NewDfsStoragePolicyRespWithDefaults instantiates a new DfsStoragePolicyResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfsStoragePolicy

`func (o *DfsStoragePolicyResp) GetDfsStoragePolicy() DfsStoragePolicy`

GetDfsStoragePolicy returns the DfsStoragePolicy field if non-nil, zero value otherwise.

### GetDfsStoragePolicyOk

`func (o *DfsStoragePolicyResp) GetDfsStoragePolicyOk() (*DfsStoragePolicy, bool)`

GetDfsStoragePolicyOk returns a tuple with the DfsStoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsStoragePolicy

`func (o *DfsStoragePolicyResp) SetDfsStoragePolicy(v DfsStoragePolicy)`

SetDfsStoragePolicy sets DfsStoragePolicy field to given value.


### GetLinkPathInfo

`func (o *DfsStoragePolicyResp) GetLinkPathInfo() []PolicyLinkPathInfo`

GetLinkPathInfo returns the LinkPathInfo field if non-nil, zero value otherwise.

### GetLinkPathInfoOk

`func (o *DfsStoragePolicyResp) GetLinkPathInfoOk() (*[]PolicyLinkPathInfo, bool)`

GetLinkPathInfoOk returns a tuple with the LinkPathInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPathInfo

`func (o *DfsStoragePolicyResp) SetLinkPathInfo(v []PolicyLinkPathInfo)`

SetLinkPathInfo sets LinkPathInfo field to given value.

### HasLinkPathInfo

`func (o *DfsStoragePolicyResp) HasLinkPathInfo() bool`

HasLinkPathInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


