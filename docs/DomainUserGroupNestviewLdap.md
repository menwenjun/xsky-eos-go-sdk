# DomainUserGroupNestviewLdap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**AdminDn** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**ConnectionTimeout** | Pointer to **int64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsGatewayGroupNum** | Pointer to **int64** |  | [optional] 
**DnsIp** | Pointer to **string** |  | [optional] 
**GroupSuffix** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Ips** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NetbiosName** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**SambaSid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Suffix** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**UserSuffix** | Pointer to **string** |  | [optional] 
**VerifyTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDomainUserGroupNestviewLdap

`func NewDomainUserGroupNestviewLdap() *DomainUserGroupNestviewLdap`

NewDomainUserGroupNestviewLdap instantiates a new DomainUserGroupNestviewLdap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainUserGroupNestviewLdapWithDefaults

`func NewDomainUserGroupNestviewLdapWithDefaults() *DomainUserGroupNestviewLdap`

NewDomainUserGroupNestviewLdapWithDefaults instantiates a new DomainUserGroupNestviewLdap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DomainUserGroupNestviewLdap) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DomainUserGroupNestviewLdap) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DomainUserGroupNestviewLdap) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DomainUserGroupNestviewLdap) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAdminDn

`func (o *DomainUserGroupNestviewLdap) GetAdminDn() string`

GetAdminDn returns the AdminDn field if non-nil, zero value otherwise.

### GetAdminDnOk

`func (o *DomainUserGroupNestviewLdap) GetAdminDnOk() (*string, bool)`

GetAdminDnOk returns a tuple with the AdminDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminDn

`func (o *DomainUserGroupNestviewLdap) SetAdminDn(v string)`

SetAdminDn sets AdminDn field to given value.

### HasAdminDn

`func (o *DomainUserGroupNestviewLdap) HasAdminDn() bool`

HasAdminDn returns a boolean if a field has been set.

### GetCluster

`func (o *DomainUserGroupNestviewLdap) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DomainUserGroupNestviewLdap) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DomainUserGroupNestviewLdap) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DomainUserGroupNestviewLdap) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *DomainUserGroupNestviewLdap) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *DomainUserGroupNestviewLdap) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *DomainUserGroupNestviewLdap) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *DomainUserGroupNestviewLdap) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetCreate

`func (o *DomainUserGroupNestviewLdap) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DomainUserGroupNestviewLdap) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DomainUserGroupNestviewLdap) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DomainUserGroupNestviewLdap) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsGatewayGroupNum

`func (o *DomainUserGroupNestviewLdap) GetDfsGatewayGroupNum() int64`

GetDfsGatewayGroupNum returns the DfsGatewayGroupNum field if non-nil, zero value otherwise.

### GetDfsGatewayGroupNumOk

`func (o *DomainUserGroupNestviewLdap) GetDfsGatewayGroupNumOk() (*int64, bool)`

GetDfsGatewayGroupNumOk returns a tuple with the DfsGatewayGroupNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroupNum

`func (o *DomainUserGroupNestviewLdap) SetDfsGatewayGroupNum(v int64)`

SetDfsGatewayGroupNum sets DfsGatewayGroupNum field to given value.

### HasDfsGatewayGroupNum

`func (o *DomainUserGroupNestviewLdap) HasDfsGatewayGroupNum() bool`

HasDfsGatewayGroupNum returns a boolean if a field has been set.

### GetDnsIp

`func (o *DomainUserGroupNestviewLdap) GetDnsIp() string`

GetDnsIp returns the DnsIp field if non-nil, zero value otherwise.

### GetDnsIpOk

`func (o *DomainUserGroupNestviewLdap) GetDnsIpOk() (*string, bool)`

GetDnsIpOk returns a tuple with the DnsIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIp

`func (o *DomainUserGroupNestviewLdap) SetDnsIp(v string)`

SetDnsIp sets DnsIp field to given value.

### HasDnsIp

`func (o *DomainUserGroupNestviewLdap) HasDnsIp() bool`

HasDnsIp returns a boolean if a field has been set.

### GetGroupSuffix

`func (o *DomainUserGroupNestviewLdap) GetGroupSuffix() string`

GetGroupSuffix returns the GroupSuffix field if non-nil, zero value otherwise.

### GetGroupSuffixOk

`func (o *DomainUserGroupNestviewLdap) GetGroupSuffixOk() (*string, bool)`

GetGroupSuffixOk returns a tuple with the GroupSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSuffix

`func (o *DomainUserGroupNestviewLdap) SetGroupSuffix(v string)`

SetGroupSuffix sets GroupSuffix field to given value.

### HasGroupSuffix

`func (o *DomainUserGroupNestviewLdap) HasGroupSuffix() bool`

HasGroupSuffix returns a boolean if a field has been set.

### GetId

`func (o *DomainUserGroupNestviewLdap) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainUserGroupNestviewLdap) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainUserGroupNestviewLdap) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DomainUserGroupNestviewLdap) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *DomainUserGroupNestviewLdap) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DomainUserGroupNestviewLdap) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DomainUserGroupNestviewLdap) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DomainUserGroupNestviewLdap) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIps

`func (o *DomainUserGroupNestviewLdap) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *DomainUserGroupNestviewLdap) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *DomainUserGroupNestviewLdap) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *DomainUserGroupNestviewLdap) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetName

`func (o *DomainUserGroupNestviewLdap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainUserGroupNestviewLdap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainUserGroupNestviewLdap) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainUserGroupNestviewLdap) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetbiosName

`func (o *DomainUserGroupNestviewLdap) GetNetbiosName() string`

GetNetbiosName returns the NetbiosName field if non-nil, zero value otherwise.

### GetNetbiosNameOk

`func (o *DomainUserGroupNestviewLdap) GetNetbiosNameOk() (*string, bool)`

GetNetbiosNameOk returns a tuple with the NetbiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosName

`func (o *DomainUserGroupNestviewLdap) SetNetbiosName(v string)`

SetNetbiosName sets NetbiosName field to given value.

### HasNetbiosName

`func (o *DomainUserGroupNestviewLdap) HasNetbiosName() bool`

HasNetbiosName returns a boolean if a field has been set.

### GetPort

`func (o *DomainUserGroupNestviewLdap) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DomainUserGroupNestviewLdap) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DomainUserGroupNestviewLdap) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DomainUserGroupNestviewLdap) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSambaSid

`func (o *DomainUserGroupNestviewLdap) GetSambaSid() string`

GetSambaSid returns the SambaSid field if non-nil, zero value otherwise.

### GetSambaSidOk

`func (o *DomainUserGroupNestviewLdap) GetSambaSidOk() (*string, bool)`

GetSambaSidOk returns a tuple with the SambaSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSambaSid

`func (o *DomainUserGroupNestviewLdap) SetSambaSid(v string)`

SetSambaSid sets SambaSid field to given value.

### HasSambaSid

`func (o *DomainUserGroupNestviewLdap) HasSambaSid() bool`

HasSambaSid returns a boolean if a field has been set.

### GetStatus

`func (o *DomainUserGroupNestviewLdap) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainUserGroupNestviewLdap) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainUserGroupNestviewLdap) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DomainUserGroupNestviewLdap) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuffix

`func (o *DomainUserGroupNestviewLdap) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *DomainUserGroupNestviewLdap) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *DomainUserGroupNestviewLdap) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *DomainUserGroupNestviewLdap) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### GetTimeout

`func (o *DomainUserGroupNestviewLdap) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DomainUserGroupNestviewLdap) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DomainUserGroupNestviewLdap) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DomainUserGroupNestviewLdap) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUpdate

`func (o *DomainUserGroupNestviewLdap) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DomainUserGroupNestviewLdap) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DomainUserGroupNestviewLdap) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DomainUserGroupNestviewLdap) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUserSuffix

`func (o *DomainUserGroupNestviewLdap) GetUserSuffix() string`

GetUserSuffix returns the UserSuffix field if non-nil, zero value otherwise.

### GetUserSuffixOk

`func (o *DomainUserGroupNestviewLdap) GetUserSuffixOk() (*string, bool)`

GetUserSuffixOk returns a tuple with the UserSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSuffix

`func (o *DomainUserGroupNestviewLdap) SetUserSuffix(v string)`

SetUserSuffix sets UserSuffix field to given value.

### HasUserSuffix

`func (o *DomainUserGroupNestviewLdap) HasUserSuffix() bool`

HasUserSuffix returns a boolean if a field has been set.

### GetVerifyTime

`func (o *DomainUserGroupNestviewLdap) GetVerifyTime() time.Time`

GetVerifyTime returns the VerifyTime field if non-nil, zero value otherwise.

### GetVerifyTimeOk

`func (o *DomainUserGroupNestviewLdap) GetVerifyTimeOk() (*time.Time, bool)`

GetVerifyTimeOk returns a tuple with the VerifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyTime

`func (o *DomainUserGroupNestviewLdap) SetVerifyTime(v time.Time)`

SetVerifyTime sets VerifyTime field to given value.

### HasVerifyTime

`func (o *DomainUserGroupNestviewLdap) HasVerifyTime() bool`

HasVerifyTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


