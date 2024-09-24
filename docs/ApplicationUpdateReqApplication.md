# ApplicationUpdateReqApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | enable or disable the application | [optional] 
**Name** | **string** | name of application | 
**Url** | **string** | url of application | 

## Methods

### NewApplicationUpdateReqApplication

`func NewApplicationUpdateReqApplication(name string, url string, ) *ApplicationUpdateReqApplication`

NewApplicationUpdateReqApplication instantiates a new ApplicationUpdateReqApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationUpdateReqApplicationWithDefaults

`func NewApplicationUpdateReqApplicationWithDefaults() *ApplicationUpdateReqApplication`

NewApplicationUpdateReqApplicationWithDefaults instantiates a new ApplicationUpdateReqApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApplicationUpdateReqApplication) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationUpdateReqApplication) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationUpdateReqApplication) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApplicationUpdateReqApplication) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *ApplicationUpdateReqApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationUpdateReqApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationUpdateReqApplication) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *ApplicationUpdateReqApplication) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApplicationUpdateReqApplication) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApplicationUpdateReqApplication) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


