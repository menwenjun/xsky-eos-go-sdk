# EmailGroupRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**Emails** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEmailGroupRecord

`func NewEmailGroupRecord() *EmailGroupRecord`

NewEmailGroupRecord instantiates a new EmailGroupRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailGroupRecordWithDefaults

`func NewEmailGroupRecordWithDefaults() *EmailGroupRecord`

NewEmailGroupRecordWithDefaults instantiates a new EmailGroupRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *EmailGroupRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *EmailGroupRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *EmailGroupRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *EmailGroupRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetEmails

`func (o *EmailGroupRecord) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *EmailGroupRecord) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *EmailGroupRecord) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *EmailGroupRecord) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetId

`func (o *EmailGroupRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailGroupRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailGroupRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EmailGroupRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EmailGroupRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailGroupRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailGroupRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmailGroupRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdate

`func (o *EmailGroupRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *EmailGroupRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *EmailGroupRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *EmailGroupRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


