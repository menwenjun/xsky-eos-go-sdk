# DfsFileQos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfsQos** | Pointer to [**DfsQos**](DfsQos.md) |  | [optional] 
**State** | Pointer to **string** | qos state of file | [optional] 

## Methods

### NewDfsFileQos

`func NewDfsFileQos() *DfsFileQos`

NewDfsFileQos instantiates a new DfsFileQos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsFileQosWithDefaults

`func NewDfsFileQosWithDefaults() *DfsFileQos`

NewDfsFileQosWithDefaults instantiates a new DfsFileQos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfsQos

`func (o *DfsFileQos) GetDfsQos() DfsQos`

GetDfsQos returns the DfsQos field if non-nil, zero value otherwise.

### GetDfsQosOk

`func (o *DfsFileQos) GetDfsQosOk() (*DfsQos, bool)`

GetDfsQosOk returns a tuple with the DfsQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsQos

`func (o *DfsFileQos) SetDfsQos(v DfsQos)`

SetDfsQos sets DfsQos field to given value.

### HasDfsQos

`func (o *DfsFileQos) HasDfsQos() bool`

HasDfsQos returns a boolean if a field has been set.

### GetState

`func (o *DfsFileQos) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DfsFileQos) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DfsFileQos) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DfsFileQos) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


