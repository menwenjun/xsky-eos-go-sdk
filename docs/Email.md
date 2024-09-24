# Email

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** | plain or html. multipart/mixed is unsupported | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**To** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEmail

`func NewEmail() *Email`

NewEmail instantiates a new Email object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailWithDefaults

`func NewEmailWithDefaults() *Email`

NewEmailWithDefaults instantiates a new Email object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *Email) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Email) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Email) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *Email) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetContentType

`func (o *Email) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *Email) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *Email) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *Email) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetFrom

`func (o *Email) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Email) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Email) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Email) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetSubject

`func (o *Email) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Email) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Email) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Email) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTo

`func (o *Email) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Email) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Email) SetTo(v []string)`

SetTo sets To field to given value.

### HasTo

`func (o *Email) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


