# SmartAttr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttrId** | Pointer to **int64** | sata only | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**Flag** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | base | [optional] 
**RawValue** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Thresh** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**WhenFailed** | Pointer to **string** |  | [optional] 
**Worst** | Pointer to **string** |  | [optional] 

## Methods

### NewSmartAttr

`func NewSmartAttr() *SmartAttr`

NewSmartAttr instantiates a new SmartAttr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartAttrWithDefaults

`func NewSmartAttrWithDefaults() *SmartAttr`

NewSmartAttrWithDefaults instantiates a new SmartAttr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttrId

`func (o *SmartAttr) GetAttrId() int64`

GetAttrId returns the AttrId field if non-nil, zero value otherwise.

### GetAttrIdOk

`func (o *SmartAttr) GetAttrIdOk() (*int64, bool)`

GetAttrIdOk returns a tuple with the AttrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrId

`func (o *SmartAttr) SetAttrId(v int64)`

SetAttrId sets AttrId field to given value.

### HasAttrId

`func (o *SmartAttr) HasAttrId() bool`

HasAttrId returns a boolean if a field has been set.

### GetCreate

`func (o *SmartAttr) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *SmartAttr) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *SmartAttr) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *SmartAttr) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDeviceId

`func (o *SmartAttr) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SmartAttr) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SmartAttr) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SmartAttr) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetFlag

`func (o *SmartAttr) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *SmartAttr) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *SmartAttr) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *SmartAttr) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetId

`func (o *SmartAttr) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmartAttr) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmartAttr) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SmartAttr) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SmartAttr) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmartAttr) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmartAttr) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SmartAttr) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRawValue

`func (o *SmartAttr) GetRawValue() string`

GetRawValue returns the RawValue field if non-nil, zero value otherwise.

### GetRawValueOk

`func (o *SmartAttr) GetRawValueOk() (*string, bool)`

GetRawValueOk returns a tuple with the RawValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValue

`func (o *SmartAttr) SetRawValue(v string)`

SetRawValue sets RawValue field to given value.

### HasRawValue

`func (o *SmartAttr) HasRawValue() bool`

HasRawValue returns a boolean if a field has been set.

### GetStatus

`func (o *SmartAttr) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SmartAttr) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SmartAttr) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SmartAttr) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThresh

`func (o *SmartAttr) GetThresh() string`

GetThresh returns the Thresh field if non-nil, zero value otherwise.

### GetThreshOk

`func (o *SmartAttr) GetThreshOk() (*string, bool)`

GetThreshOk returns a tuple with the Thresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresh

`func (o *SmartAttr) SetThresh(v string)`

SetThresh sets Thresh field to given value.

### HasThresh

`func (o *SmartAttr) HasThresh() bool`

HasThresh returns a boolean if a field has been set.

### GetType

`func (o *SmartAttr) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SmartAttr) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SmartAttr) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SmartAttr) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *SmartAttr) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SmartAttr) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SmartAttr) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SmartAttr) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetWhenFailed

`func (o *SmartAttr) GetWhenFailed() string`

GetWhenFailed returns the WhenFailed field if non-nil, zero value otherwise.

### GetWhenFailedOk

`func (o *SmartAttr) GetWhenFailedOk() (*string, bool)`

GetWhenFailedOk returns a tuple with the WhenFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenFailed

`func (o *SmartAttr) SetWhenFailed(v string)`

SetWhenFailed sets WhenFailed field to given value.

### HasWhenFailed

`func (o *SmartAttr) HasWhenFailed() bool`

HasWhenFailed returns a boolean if a field has been set.

### GetWorst

`func (o *SmartAttr) GetWorst() string`

GetWorst returns the Worst field if non-nil, zero value otherwise.

### GetWorstOk

`func (o *SmartAttr) GetWorstOk() (*string, bool)`

GetWorstOk returns a tuple with the Worst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorst

`func (o *SmartAttr) SetWorst(v string)`

SetWorst sets Worst field to given value.

### HasWorst

`func (o *SmartAttr) HasWorst() bool`

HasWorst returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


