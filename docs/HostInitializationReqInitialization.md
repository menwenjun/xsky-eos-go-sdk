# HostInitializationReqInitialization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminIps** | **[]string** |  | 
**DisableFirewalld** | Pointer to **bool** |  | [optional] 
**HostnamePrefix** | Pointer to **string** |  | [optional] 
**HostnameSuffixLen** | Pointer to **int64** |  | [optional] 
**RsaKeyId** | Pointer to **string** |  | [optional] 
**SetChrony** | Pointer to **bool** |  | [optional] 
**SetHostname** | Pointer to **bool** |  | [optional] 
**SshPassword** | Pointer to **string** |  | [optional] 
**SshUser** | **string** |  | 

## Methods

### NewHostInitializationReqInitialization

`func NewHostInitializationReqInitialization(adminIps []string, sshUser string, ) *HostInitializationReqInitialization`

NewHostInitializationReqInitialization instantiates a new HostInitializationReqInitialization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostInitializationReqInitializationWithDefaults

`func NewHostInitializationReqInitializationWithDefaults() *HostInitializationReqInitialization`

NewHostInitializationReqInitializationWithDefaults instantiates a new HostInitializationReqInitialization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminIps

`func (o *HostInitializationReqInitialization) GetAdminIps() []string`

GetAdminIps returns the AdminIps field if non-nil, zero value otherwise.

### GetAdminIpsOk

`func (o *HostInitializationReqInitialization) GetAdminIpsOk() (*[]string, bool)`

GetAdminIpsOk returns a tuple with the AdminIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminIps

`func (o *HostInitializationReqInitialization) SetAdminIps(v []string)`

SetAdminIps sets AdminIps field to given value.


### GetDisableFirewalld

`func (o *HostInitializationReqInitialization) GetDisableFirewalld() bool`

GetDisableFirewalld returns the DisableFirewalld field if non-nil, zero value otherwise.

### GetDisableFirewalldOk

`func (o *HostInitializationReqInitialization) GetDisableFirewalldOk() (*bool, bool)`

GetDisableFirewalldOk returns a tuple with the DisableFirewalld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFirewalld

`func (o *HostInitializationReqInitialization) SetDisableFirewalld(v bool)`

SetDisableFirewalld sets DisableFirewalld field to given value.

### HasDisableFirewalld

`func (o *HostInitializationReqInitialization) HasDisableFirewalld() bool`

HasDisableFirewalld returns a boolean if a field has been set.

### GetHostnamePrefix

`func (o *HostInitializationReqInitialization) GetHostnamePrefix() string`

GetHostnamePrefix returns the HostnamePrefix field if non-nil, zero value otherwise.

### GetHostnamePrefixOk

`func (o *HostInitializationReqInitialization) GetHostnamePrefixOk() (*string, bool)`

GetHostnamePrefixOk returns a tuple with the HostnamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnamePrefix

`func (o *HostInitializationReqInitialization) SetHostnamePrefix(v string)`

SetHostnamePrefix sets HostnamePrefix field to given value.

### HasHostnamePrefix

`func (o *HostInitializationReqInitialization) HasHostnamePrefix() bool`

HasHostnamePrefix returns a boolean if a field has been set.

### GetHostnameSuffixLen

`func (o *HostInitializationReqInitialization) GetHostnameSuffixLen() int64`

GetHostnameSuffixLen returns the HostnameSuffixLen field if non-nil, zero value otherwise.

### GetHostnameSuffixLenOk

`func (o *HostInitializationReqInitialization) GetHostnameSuffixLenOk() (*int64, bool)`

GetHostnameSuffixLenOk returns a tuple with the HostnameSuffixLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameSuffixLen

`func (o *HostInitializationReqInitialization) SetHostnameSuffixLen(v int64)`

SetHostnameSuffixLen sets HostnameSuffixLen field to given value.

### HasHostnameSuffixLen

`func (o *HostInitializationReqInitialization) HasHostnameSuffixLen() bool`

HasHostnameSuffixLen returns a boolean if a field has been set.

### GetRsaKeyId

`func (o *HostInitializationReqInitialization) GetRsaKeyId() string`

GetRsaKeyId returns the RsaKeyId field if non-nil, zero value otherwise.

### GetRsaKeyIdOk

`func (o *HostInitializationReqInitialization) GetRsaKeyIdOk() (*string, bool)`

GetRsaKeyIdOk returns a tuple with the RsaKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaKeyId

`func (o *HostInitializationReqInitialization) SetRsaKeyId(v string)`

SetRsaKeyId sets RsaKeyId field to given value.

### HasRsaKeyId

`func (o *HostInitializationReqInitialization) HasRsaKeyId() bool`

HasRsaKeyId returns a boolean if a field has been set.

### GetSetChrony

`func (o *HostInitializationReqInitialization) GetSetChrony() bool`

GetSetChrony returns the SetChrony field if non-nil, zero value otherwise.

### GetSetChronyOk

`func (o *HostInitializationReqInitialization) GetSetChronyOk() (*bool, bool)`

GetSetChronyOk returns a tuple with the SetChrony field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetChrony

`func (o *HostInitializationReqInitialization) SetSetChrony(v bool)`

SetSetChrony sets SetChrony field to given value.

### HasSetChrony

`func (o *HostInitializationReqInitialization) HasSetChrony() bool`

HasSetChrony returns a boolean if a field has been set.

### GetSetHostname

`func (o *HostInitializationReqInitialization) GetSetHostname() bool`

GetSetHostname returns the SetHostname field if non-nil, zero value otherwise.

### GetSetHostnameOk

`func (o *HostInitializationReqInitialization) GetSetHostnameOk() (*bool, bool)`

GetSetHostnameOk returns a tuple with the SetHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetHostname

`func (o *HostInitializationReqInitialization) SetSetHostname(v bool)`

SetSetHostname sets SetHostname field to given value.

### HasSetHostname

`func (o *HostInitializationReqInitialization) HasSetHostname() bool`

HasSetHostname returns a boolean if a field has been set.

### GetSshPassword

`func (o *HostInitializationReqInitialization) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *HostInitializationReqInitialization) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *HostInitializationReqInitialization) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *HostInitializationReqInitialization) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshUser

`func (o *HostInitializationReqInitialization) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *HostInitializationReqInitialization) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *HostInitializationReqInitialization) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


