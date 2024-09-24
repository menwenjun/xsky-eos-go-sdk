# HostEncSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**Enclosures** | Pointer to [**[]Enclosure**](Enclosure.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 

## Methods

### NewHostEncSpec

`func NewHostEncSpec() *HostEncSpec`

NewHostEncSpec instantiates a new HostEncSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostEncSpecWithDefaults

`func NewHostEncSpecWithDefaults() *HostEncSpec`

NewHostEncSpecWithDefaults instantiates a new HostEncSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *HostEncSpec) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *HostEncSpec) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *HostEncSpec) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *HostEncSpec) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetEnclosures

`func (o *HostEncSpec) GetEnclosures() []Enclosure`

GetEnclosures returns the Enclosures field if non-nil, zero value otherwise.

### GetEnclosuresOk

`func (o *HostEncSpec) GetEnclosuresOk() (*[]Enclosure, bool)`

GetEnclosuresOk returns a tuple with the Enclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosures

`func (o *HostEncSpec) SetEnclosures(v []Enclosure)`

SetEnclosures sets Enclosures field to given value.

### HasEnclosures

`func (o *HostEncSpec) HasEnclosures() bool`

HasEnclosures returns a boolean if a field has been set.

### GetId

`func (o *HostEncSpec) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostEncSpec) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostEncSpec) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *HostEncSpec) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *HostEncSpec) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *HostEncSpec) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *HostEncSpec) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *HostEncSpec) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetUpdate

`func (o *HostEncSpec) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *HostEncSpec) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *HostEncSpec) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *HostEncSpec) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVendor

`func (o *HostEncSpec) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HostEncSpec) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HostEncSpec) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *HostEncSpec) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


