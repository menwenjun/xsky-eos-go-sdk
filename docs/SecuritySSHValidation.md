# SecuritySSHValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**PermitRootLogin** | Pointer to **bool** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**PubKeyAuthentication** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSecuritySSHValidation

`func NewSecuritySSHValidation() *SecuritySSHValidation`

NewSecuritySSHValidation instantiates a new SecuritySSHValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySSHValidationWithDefaults

`func NewSecuritySSHValidationWithDefaults() *SecuritySSHValidation`

NewSecuritySSHValidationWithDefaults instantiates a new SecuritySSHValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *SecuritySSHValidation) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *SecuritySSHValidation) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *SecuritySSHValidation) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *SecuritySSHValidation) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *SecuritySSHValidation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecuritySSHValidation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecuritySSHValidation) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SecuritySSHValidation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *SecuritySSHValidation) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SecuritySSHValidation) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SecuritySSHValidation) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SecuritySSHValidation) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPermitRootLogin

`func (o *SecuritySSHValidation) GetPermitRootLogin() bool`

GetPermitRootLogin returns the PermitRootLogin field if non-nil, zero value otherwise.

### GetPermitRootLoginOk

`func (o *SecuritySSHValidation) GetPermitRootLoginOk() (*bool, bool)`

GetPermitRootLoginOk returns a tuple with the PermitRootLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitRootLogin

`func (o *SecuritySSHValidation) SetPermitRootLogin(v bool)`

SetPermitRootLogin sets PermitRootLogin field to given value.

### HasPermitRootLogin

`func (o *SecuritySSHValidation) HasPermitRootLogin() bool`

HasPermitRootLogin returns a boolean if a field has been set.

### GetPort

`func (o *SecuritySSHValidation) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SecuritySSHValidation) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SecuritySSHValidation) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SecuritySSHValidation) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPubKeyAuthentication

`func (o *SecuritySSHValidation) GetPubKeyAuthentication() bool`

GetPubKeyAuthentication returns the PubKeyAuthentication field if non-nil, zero value otherwise.

### GetPubKeyAuthenticationOk

`func (o *SecuritySSHValidation) GetPubKeyAuthenticationOk() (*bool, bool)`

GetPubKeyAuthenticationOk returns a tuple with the PubKeyAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKeyAuthentication

`func (o *SecuritySSHValidation) SetPubKeyAuthentication(v bool)`

SetPubKeyAuthentication sets PubKeyAuthentication field to given value.

### HasPubKeyAuthentication

`func (o *SecuritySSHValidation) HasPubKeyAuthentication() bool`

HasPubKeyAuthentication returns a boolean if a field has been set.

### GetStatus

`func (o *SecuritySSHValidation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecuritySSHValidation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecuritySSHValidation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecuritySSHValidation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeout

`func (o *SecuritySSHValidation) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SecuritySSHValidation) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SecuritySSHValidation) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SecuritySSHValidation) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUpdate

`func (o *SecuritySSHValidation) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *SecuritySSHValidation) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *SecuritySSHValidation) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *SecuritySSHValidation) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


