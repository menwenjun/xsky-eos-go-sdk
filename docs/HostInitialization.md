# HostInitialization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminIps** | Pointer to **[]string** |  | [optional] 
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DisableFirewalld** | Pointer to **bool** |  | [optional] 
**HostnamePrefix** | Pointer to **string** |  | [optional] 
**HostnameSuffixLen** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**SetChrony** | Pointer to **bool** |  | [optional] 
**SetHostname** | Pointer to **bool** |  | [optional] 
**SshUser** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewHostInitialization

`func NewHostInitialization() *HostInitialization`

NewHostInitialization instantiates a new HostInitialization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostInitializationWithDefaults

`func NewHostInitializationWithDefaults() *HostInitialization`

NewHostInitializationWithDefaults instantiates a new HostInitialization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminIps

`func (o *HostInitialization) GetAdminIps() []string`

GetAdminIps returns the AdminIps field if non-nil, zero value otherwise.

### GetAdminIpsOk

`func (o *HostInitialization) GetAdminIpsOk() (*[]string, bool)`

GetAdminIpsOk returns a tuple with the AdminIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminIps

`func (o *HostInitialization) SetAdminIps(v []string)`

SetAdminIps sets AdminIps field to given value.

### HasAdminIps

`func (o *HostInitialization) HasAdminIps() bool`

HasAdminIps returns a boolean if a field has been set.

### GetCluster

`func (o *HostInitialization) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HostInitialization) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HostInitialization) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HostInitialization) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *HostInitialization) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *HostInitialization) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *HostInitialization) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *HostInitialization) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDisableFirewalld

`func (o *HostInitialization) GetDisableFirewalld() bool`

GetDisableFirewalld returns the DisableFirewalld field if non-nil, zero value otherwise.

### GetDisableFirewalldOk

`func (o *HostInitialization) GetDisableFirewalldOk() (*bool, bool)`

GetDisableFirewalldOk returns a tuple with the DisableFirewalld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFirewalld

`func (o *HostInitialization) SetDisableFirewalld(v bool)`

SetDisableFirewalld sets DisableFirewalld field to given value.

### HasDisableFirewalld

`func (o *HostInitialization) HasDisableFirewalld() bool`

HasDisableFirewalld returns a boolean if a field has been set.

### GetHostnamePrefix

`func (o *HostInitialization) GetHostnamePrefix() string`

GetHostnamePrefix returns the HostnamePrefix field if non-nil, zero value otherwise.

### GetHostnamePrefixOk

`func (o *HostInitialization) GetHostnamePrefixOk() (*string, bool)`

GetHostnamePrefixOk returns a tuple with the HostnamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnamePrefix

`func (o *HostInitialization) SetHostnamePrefix(v string)`

SetHostnamePrefix sets HostnamePrefix field to given value.

### HasHostnamePrefix

`func (o *HostInitialization) HasHostnamePrefix() bool`

HasHostnamePrefix returns a boolean if a field has been set.

### GetHostnameSuffixLen

`func (o *HostInitialization) GetHostnameSuffixLen() int64`

GetHostnameSuffixLen returns the HostnameSuffixLen field if non-nil, zero value otherwise.

### GetHostnameSuffixLenOk

`func (o *HostInitialization) GetHostnameSuffixLenOk() (*int64, bool)`

GetHostnameSuffixLenOk returns a tuple with the HostnameSuffixLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameSuffixLen

`func (o *HostInitialization) SetHostnameSuffixLen(v int64)`

SetHostnameSuffixLen sets HostnameSuffixLen field to given value.

### HasHostnameSuffixLen

`func (o *HostInitialization) HasHostnameSuffixLen() bool`

HasHostnameSuffixLen returns a boolean if a field has been set.

### GetId

`func (o *HostInitialization) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostInitialization) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostInitialization) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *HostInitialization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *HostInitialization) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HostInitialization) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HostInitialization) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HostInitialization) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSetChrony

`func (o *HostInitialization) GetSetChrony() bool`

GetSetChrony returns the SetChrony field if non-nil, zero value otherwise.

### GetSetChronyOk

`func (o *HostInitialization) GetSetChronyOk() (*bool, bool)`

GetSetChronyOk returns a tuple with the SetChrony field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetChrony

`func (o *HostInitialization) SetSetChrony(v bool)`

SetSetChrony sets SetChrony field to given value.

### HasSetChrony

`func (o *HostInitialization) HasSetChrony() bool`

HasSetChrony returns a boolean if a field has been set.

### GetSetHostname

`func (o *HostInitialization) GetSetHostname() bool`

GetSetHostname returns the SetHostname field if non-nil, zero value otherwise.

### GetSetHostnameOk

`func (o *HostInitialization) GetSetHostnameOk() (*bool, bool)`

GetSetHostnameOk returns a tuple with the SetHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetHostname

`func (o *HostInitialization) SetSetHostname(v bool)`

SetSetHostname sets SetHostname field to given value.

### HasSetHostname

`func (o *HostInitialization) HasSetHostname() bool`

HasSetHostname returns a boolean if a field has been set.

### GetSshUser

`func (o *HostInitialization) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *HostInitialization) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *HostInitialization) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *HostInitialization) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.

### GetStatus

`func (o *HostInitialization) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostInitialization) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostInitialization) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HostInitialization) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


