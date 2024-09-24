# ObjectInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to **string** |  | [optional] 
**UrlInfo** | Pointer to [**[]URLInfo**](URLInfo.md) |  | [optional] 

## Methods

### NewObjectInfo

`func NewObjectInfo() *ObjectInfo`

NewObjectInfo instantiates a new ObjectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectInfoWithDefaults

`func NewObjectInfoWithDefaults() *ObjectInfo`

NewObjectInfoWithDefaults instantiates a new ObjectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *ObjectInfo) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ObjectInfo) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ObjectInfo) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ObjectInfo) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUrlInfo

`func (o *ObjectInfo) GetUrlInfo() []URLInfo`

GetUrlInfo returns the UrlInfo field if non-nil, zero value otherwise.

### GetUrlInfoOk

`func (o *ObjectInfo) GetUrlInfoOk() (*[]URLInfo, bool)`

GetUrlInfoOk returns a tuple with the UrlInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlInfo

`func (o *ObjectInfo) SetUrlInfo(v []URLInfo)`

SetUrlInfo sets UrlInfo field to given value.

### HasUrlInfo

`func (o *ObjectInfo) HasUrlInfo() bool`

HasUrlInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


