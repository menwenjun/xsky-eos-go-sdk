# FSLdapUpdateReqInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminDn** | Pointer to **string** | ldap admin dn | [optional] 
**ConnectionTimeout** | Pointer to **int64** | timeout for connection | [optional] 
**DnsIp** | Pointer to **string** | dns server ip | [optional] 
**GroupSuffix** | Pointer to **string** | group suffix | [optional] 
**Ip** | **string** | ip of server | 
**Ips** | Pointer to **[]string** | ips of standby servers | [optional] 
**Name** | **string** | name of ldap | 
**Password** | Pointer to **string** | bind password | [optional] 
**Port** | **int64** | ldap service port | 
**Suffix** | **string** | ldap suffix | 
**Timeout** | Pointer to **int64** | timeout for searching | [optional] 
**UserSuffix** | Pointer to **string** | user suffix | [optional] 

## Methods

### NewFSLdapUpdateReqInfo

`func NewFSLdapUpdateReqInfo(ip string, name string, port int64, suffix string, ) *FSLdapUpdateReqInfo`

NewFSLdapUpdateReqInfo instantiates a new FSLdapUpdateReqInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFSLdapUpdateReqInfoWithDefaults

`func NewFSLdapUpdateReqInfoWithDefaults() *FSLdapUpdateReqInfo`

NewFSLdapUpdateReqInfoWithDefaults instantiates a new FSLdapUpdateReqInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminDn

`func (o *FSLdapUpdateReqInfo) GetAdminDn() string`

GetAdminDn returns the AdminDn field if non-nil, zero value otherwise.

### GetAdminDnOk

`func (o *FSLdapUpdateReqInfo) GetAdminDnOk() (*string, bool)`

GetAdminDnOk returns a tuple with the AdminDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminDn

`func (o *FSLdapUpdateReqInfo) SetAdminDn(v string)`

SetAdminDn sets AdminDn field to given value.

### HasAdminDn

`func (o *FSLdapUpdateReqInfo) HasAdminDn() bool`

HasAdminDn returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *FSLdapUpdateReqInfo) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *FSLdapUpdateReqInfo) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *FSLdapUpdateReqInfo) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *FSLdapUpdateReqInfo) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetDnsIp

`func (o *FSLdapUpdateReqInfo) GetDnsIp() string`

GetDnsIp returns the DnsIp field if non-nil, zero value otherwise.

### GetDnsIpOk

`func (o *FSLdapUpdateReqInfo) GetDnsIpOk() (*string, bool)`

GetDnsIpOk returns a tuple with the DnsIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIp

`func (o *FSLdapUpdateReqInfo) SetDnsIp(v string)`

SetDnsIp sets DnsIp field to given value.

### HasDnsIp

`func (o *FSLdapUpdateReqInfo) HasDnsIp() bool`

HasDnsIp returns a boolean if a field has been set.

### GetGroupSuffix

`func (o *FSLdapUpdateReqInfo) GetGroupSuffix() string`

GetGroupSuffix returns the GroupSuffix field if non-nil, zero value otherwise.

### GetGroupSuffixOk

`func (o *FSLdapUpdateReqInfo) GetGroupSuffixOk() (*string, bool)`

GetGroupSuffixOk returns a tuple with the GroupSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSuffix

`func (o *FSLdapUpdateReqInfo) SetGroupSuffix(v string)`

SetGroupSuffix sets GroupSuffix field to given value.

### HasGroupSuffix

`func (o *FSLdapUpdateReqInfo) HasGroupSuffix() bool`

HasGroupSuffix returns a boolean if a field has been set.

### GetIp

`func (o *FSLdapUpdateReqInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *FSLdapUpdateReqInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *FSLdapUpdateReqInfo) SetIp(v string)`

SetIp sets Ip field to given value.


### GetIps

`func (o *FSLdapUpdateReqInfo) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *FSLdapUpdateReqInfo) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *FSLdapUpdateReqInfo) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *FSLdapUpdateReqInfo) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetName

`func (o *FSLdapUpdateReqInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FSLdapUpdateReqInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FSLdapUpdateReqInfo) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *FSLdapUpdateReqInfo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FSLdapUpdateReqInfo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FSLdapUpdateReqInfo) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *FSLdapUpdateReqInfo) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *FSLdapUpdateReqInfo) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FSLdapUpdateReqInfo) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FSLdapUpdateReqInfo) SetPort(v int64)`

SetPort sets Port field to given value.


### GetSuffix

`func (o *FSLdapUpdateReqInfo) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *FSLdapUpdateReqInfo) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *FSLdapUpdateReqInfo) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.


### GetTimeout

`func (o *FSLdapUpdateReqInfo) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *FSLdapUpdateReqInfo) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *FSLdapUpdateReqInfo) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *FSLdapUpdateReqInfo) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUserSuffix

`func (o *FSLdapUpdateReqInfo) GetUserSuffix() string`

GetUserSuffix returns the UserSuffix field if non-nil, zero value otherwise.

### GetUserSuffixOk

`func (o *FSLdapUpdateReqInfo) GetUserSuffixOk() (*string, bool)`

GetUserSuffixOk returns a tuple with the UserSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSuffix

`func (o *FSLdapUpdateReqInfo) SetUserSuffix(v string)`

SetUserSuffix sets UserSuffix field to given value.

### HasUserSuffix

`func (o *FSLdapUpdateReqInfo) HasUserSuffix() bool`

HasUserSuffix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


