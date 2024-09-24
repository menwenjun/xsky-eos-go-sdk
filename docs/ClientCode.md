# ClientCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewClientCode

`func NewClientCode() *ClientCode`

NewClientCode instantiates a new ClientCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCodeWithDefaults

`func NewClientCodeWithDefaults() *ClientCode`

NewClientCodeWithDefaults instantiates a new ClientCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ClientCode) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ClientCode) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ClientCode) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ClientCode) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreate

`func (o *ClientCode) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ClientCode) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ClientCode) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ClientCode) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *ClientCode) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientCode) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientCode) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClientCode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ClientCode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientCode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientCode) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClientCode) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


