# SecuritySSHValidationReqValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermitRootLogin** | Pointer to **bool** |  | [optional] 
**Port** | **int64** |  | 
**PubKeyAuthentication** | Pointer to **bool** |  | [optional] 
**TimeOut** | Pointer to **int64** |  | [optional] 

## Methods

### NewSecuritySSHValidationReqValidation

`func NewSecuritySSHValidationReqValidation(port int64, ) *SecuritySSHValidationReqValidation`

NewSecuritySSHValidationReqValidation instantiates a new SecuritySSHValidationReqValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySSHValidationReqValidationWithDefaults

`func NewSecuritySSHValidationReqValidationWithDefaults() *SecuritySSHValidationReqValidation`

NewSecuritySSHValidationReqValidationWithDefaults instantiates a new SecuritySSHValidationReqValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermitRootLogin

`func (o *SecuritySSHValidationReqValidation) GetPermitRootLogin() bool`

GetPermitRootLogin returns the PermitRootLogin field if non-nil, zero value otherwise.

### GetPermitRootLoginOk

`func (o *SecuritySSHValidationReqValidation) GetPermitRootLoginOk() (*bool, bool)`

GetPermitRootLoginOk returns a tuple with the PermitRootLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitRootLogin

`func (o *SecuritySSHValidationReqValidation) SetPermitRootLogin(v bool)`

SetPermitRootLogin sets PermitRootLogin field to given value.

### HasPermitRootLogin

`func (o *SecuritySSHValidationReqValidation) HasPermitRootLogin() bool`

HasPermitRootLogin returns a boolean if a field has been set.

### GetPort

`func (o *SecuritySSHValidationReqValidation) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SecuritySSHValidationReqValidation) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SecuritySSHValidationReqValidation) SetPort(v int64)`

SetPort sets Port field to given value.


### GetPubKeyAuthentication

`func (o *SecuritySSHValidationReqValidation) GetPubKeyAuthentication() bool`

GetPubKeyAuthentication returns the PubKeyAuthentication field if non-nil, zero value otherwise.

### GetPubKeyAuthenticationOk

`func (o *SecuritySSHValidationReqValidation) GetPubKeyAuthenticationOk() (*bool, bool)`

GetPubKeyAuthenticationOk returns a tuple with the PubKeyAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKeyAuthentication

`func (o *SecuritySSHValidationReqValidation) SetPubKeyAuthentication(v bool)`

SetPubKeyAuthentication sets PubKeyAuthentication field to given value.

### HasPubKeyAuthentication

`func (o *SecuritySSHValidationReqValidation) HasPubKeyAuthentication() bool`

HasPubKeyAuthentication returns a boolean if a field has been set.

### GetTimeOut

`func (o *SecuritySSHValidationReqValidation) GetTimeOut() int64`

GetTimeOut returns the TimeOut field if non-nil, zero value otherwise.

### GetTimeOutOk

`func (o *SecuritySSHValidationReqValidation) GetTimeOutOk() (*int64, bool)`

GetTimeOutOk returns a tuple with the TimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOut

`func (o *SecuritySSHValidationReqValidation) SetTimeOut(v int64)`

SetTimeOut sets TimeOut field to given value.

### HasTimeOut

`func (o *SecuritySSHValidationReqValidation) HasTimeOut() bool`

HasTimeOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


