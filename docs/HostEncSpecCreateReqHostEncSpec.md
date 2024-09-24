# HostEncSpecCreateReqHostEncSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enclosures** | [**[]Enclosure**](Enclosure.md) | host enclosures | 
**Model** | **string** | host model | 
**Vendor** | **string** | host vendor | 

## Methods

### NewHostEncSpecCreateReqHostEncSpec

`func NewHostEncSpecCreateReqHostEncSpec(enclosures []Enclosure, model string, vendor string, ) *HostEncSpecCreateReqHostEncSpec`

NewHostEncSpecCreateReqHostEncSpec instantiates a new HostEncSpecCreateReqHostEncSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostEncSpecCreateReqHostEncSpecWithDefaults

`func NewHostEncSpecCreateReqHostEncSpecWithDefaults() *HostEncSpecCreateReqHostEncSpec`

NewHostEncSpecCreateReqHostEncSpecWithDefaults instantiates a new HostEncSpecCreateReqHostEncSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnclosures

`func (o *HostEncSpecCreateReqHostEncSpec) GetEnclosures() []Enclosure`

GetEnclosures returns the Enclosures field if non-nil, zero value otherwise.

### GetEnclosuresOk

`func (o *HostEncSpecCreateReqHostEncSpec) GetEnclosuresOk() (*[]Enclosure, bool)`

GetEnclosuresOk returns a tuple with the Enclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosures

`func (o *HostEncSpecCreateReqHostEncSpec) SetEnclosures(v []Enclosure)`

SetEnclosures sets Enclosures field to given value.


### GetModel

`func (o *HostEncSpecCreateReqHostEncSpec) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *HostEncSpecCreateReqHostEncSpec) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *HostEncSpecCreateReqHostEncSpec) SetModel(v string)`

SetModel sets Model field to given value.


### GetVendor

`func (o *HostEncSpecCreateReqHostEncSpec) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HostEncSpecCreateReqHostEncSpec) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HostEncSpecCreateReqHostEncSpec) SetVendor(v string)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


