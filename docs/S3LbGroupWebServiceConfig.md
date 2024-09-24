# S3LbGroupWebServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | **string** |  | 
**Credential** | Pointer to [**S3LbGroupWebServiceConfigCredential**](S3LbGroupWebServiceConfigCredential.md) |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Protocol** | **string** |  | 
**SharePath** | **string** |  | 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewS3LbGroupWebServiceConfig

`func NewS3LbGroupWebServiceConfig(authentication string, protocol string, sharePath string, ) *S3LbGroupWebServiceConfig`

NewS3LbGroupWebServiceConfig instantiates a new S3LbGroupWebServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LbGroupWebServiceConfigWithDefaults

`func NewS3LbGroupWebServiceConfigWithDefaults() *S3LbGroupWebServiceConfig`

NewS3LbGroupWebServiceConfigWithDefaults instantiates a new S3LbGroupWebServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthentication

`func (o *S3LbGroupWebServiceConfig) GetAuthentication() string`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *S3LbGroupWebServiceConfig) GetAuthenticationOk() (*string, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *S3LbGroupWebServiceConfig) SetAuthentication(v string)`

SetAuthentication sets Authentication field to given value.


### GetCredential

`func (o *S3LbGroupWebServiceConfig) GetCredential() S3LbGroupWebServiceConfigCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *S3LbGroupWebServiceConfig) GetCredentialOk() (*S3LbGroupWebServiceConfigCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *S3LbGroupWebServiceConfig) SetCredential(v S3LbGroupWebServiceConfigCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *S3LbGroupWebServiceConfig) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetPassword

`func (o *S3LbGroupWebServiceConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *S3LbGroupWebServiceConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *S3LbGroupWebServiceConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *S3LbGroupWebServiceConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProtocol

`func (o *S3LbGroupWebServiceConfig) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *S3LbGroupWebServiceConfig) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *S3LbGroupWebServiceConfig) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetSharePath

`func (o *S3LbGroupWebServiceConfig) GetSharePath() string`

GetSharePath returns the SharePath field if non-nil, zero value otherwise.

### GetSharePathOk

`func (o *S3LbGroupWebServiceConfig) GetSharePathOk() (*string, bool)`

GetSharePathOk returns a tuple with the SharePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePath

`func (o *S3LbGroupWebServiceConfig) SetSharePath(v string)`

SetSharePath sets SharePath field to given value.


### GetUserName

`func (o *S3LbGroupWebServiceConfig) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *S3LbGroupWebServiceConfig) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *S3LbGroupWebServiceConfig) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *S3LbGroupWebServiceConfig) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


