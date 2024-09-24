# EmailSendReqEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** |  | 
**ContentType** | Pointer to **string** | content type, &#39;text/plain&#39; or &#39;text/html&#39; | [optional] 
**To** | **[]string** |  | 

## Methods

### NewEmailSendReqEmail

`func NewEmailSendReqEmail(body string, to []string, ) *EmailSendReqEmail`

NewEmailSendReqEmail instantiates a new EmailSendReqEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSendReqEmailWithDefaults

`func NewEmailSendReqEmailWithDefaults() *EmailSendReqEmail`

NewEmailSendReqEmailWithDefaults instantiates a new EmailSendReqEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *EmailSendReqEmail) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EmailSendReqEmail) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EmailSendReqEmail) SetBody(v string)`

SetBody sets Body field to given value.


### GetContentType

`func (o *EmailSendReqEmail) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *EmailSendReqEmail) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *EmailSendReqEmail) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *EmailSendReqEmail) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetTo

`func (o *EmailSendReqEmail) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailSendReqEmail) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailSendReqEmail) SetTo(v []string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


