# SecuritySSHConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**PermitRootLogin** | Pointer to **bool** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**PubKeyAuthentication** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewSecuritySSHConfig

`func NewSecuritySSHConfig() *SecuritySSHConfig`

NewSecuritySSHConfig instantiates a new SecuritySSHConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySSHConfigWithDefaults

`func NewSecuritySSHConfigWithDefaults() *SecuritySSHConfig`

NewSecuritySSHConfigWithDefaults instantiates a new SecuritySSHConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *SecuritySSHConfig) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *SecuritySSHConfig) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *SecuritySSHConfig) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *SecuritySSHConfig) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCreate

`func (o *SecuritySSHConfig) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *SecuritySSHConfig) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *SecuritySSHConfig) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *SecuritySSHConfig) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *SecuritySSHConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecuritySSHConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecuritySSHConfig) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SecuritySSHConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *SecuritySSHConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SecuritySSHConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SecuritySSHConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SecuritySSHConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPermitRootLogin

`func (o *SecuritySSHConfig) GetPermitRootLogin() bool`

GetPermitRootLogin returns the PermitRootLogin field if non-nil, zero value otherwise.

### GetPermitRootLoginOk

`func (o *SecuritySSHConfig) GetPermitRootLoginOk() (*bool, bool)`

GetPermitRootLoginOk returns a tuple with the PermitRootLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitRootLogin

`func (o *SecuritySSHConfig) SetPermitRootLogin(v bool)`

SetPermitRootLogin sets PermitRootLogin field to given value.

### HasPermitRootLogin

`func (o *SecuritySSHConfig) HasPermitRootLogin() bool`

HasPermitRootLogin returns a boolean if a field has been set.

### GetPort

`func (o *SecuritySSHConfig) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SecuritySSHConfig) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SecuritySSHConfig) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SecuritySSHConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPubKeyAuthentication

`func (o *SecuritySSHConfig) GetPubKeyAuthentication() bool`

GetPubKeyAuthentication returns the PubKeyAuthentication field if non-nil, zero value otherwise.

### GetPubKeyAuthenticationOk

`func (o *SecuritySSHConfig) GetPubKeyAuthenticationOk() (*bool, bool)`

GetPubKeyAuthenticationOk returns a tuple with the PubKeyAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKeyAuthentication

`func (o *SecuritySSHConfig) SetPubKeyAuthentication(v bool)`

SetPubKeyAuthentication sets PubKeyAuthentication field to given value.

### HasPubKeyAuthentication

`func (o *SecuritySSHConfig) HasPubKeyAuthentication() bool`

HasPubKeyAuthentication returns a boolean if a field has been set.

### GetStatus

`func (o *SecuritySSHConfig) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecuritySSHConfig) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecuritySSHConfig) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecuritySSHConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeout

`func (o *SecuritySSHConfig) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SecuritySSHConfig) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SecuritySSHConfig) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SecuritySSHConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUpdate

`func (o *SecuritySSHConfig) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *SecuritySSHConfig) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *SecuritySSHConfig) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *SecuritySSHConfig) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUsername

`func (o *SecuritySSHConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SecuritySSHConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SecuritySSHConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SecuritySSHConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


