# EmailSendReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to [**EmailSendReqEmail**](EmailSendReqEmail.md) |  | [optional] 
**EmailConfig** | Pointer to [**EmailConfig**](EmailConfig.md) |  | [optional] 

## Methods

### NewEmailSendReq

`func NewEmailSendReq() *EmailSendReq`

NewEmailSendReq instantiates a new EmailSendReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSendReqWithDefaults

`func NewEmailSendReqWithDefaults() *EmailSendReq`

NewEmailSendReqWithDefaults instantiates a new EmailSendReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EmailSendReq) GetEmail() EmailSendReqEmail`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailSendReq) GetEmailOk() (*EmailSendReqEmail, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailSendReq) SetEmail(v EmailSendReqEmail)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmailSendReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailConfig

`func (o *EmailSendReq) GetEmailConfig() EmailConfig`

GetEmailConfig returns the EmailConfig field if non-nil, zero value otherwise.

### GetEmailConfigOk

`func (o *EmailSendReq) GetEmailConfigOk() (*EmailConfig, bool)`

GetEmailConfigOk returns a tuple with the EmailConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfig

`func (o *EmailSendReq) SetEmailConfig(v EmailConfig)`

SetEmailConfig sets EmailConfig field to given value.

### HasEmailConfig

`func (o *EmailSendReq) HasEmailConfig() bool`

HasEmailConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


