# SnmpGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthPassword** | Pointer to **string** | Authentication protocol pass phrase | [optional] 
**AuthProtocol** | Pointer to **string** | Authentication protocol | [optional] 
**ContextEngineId** | Pointer to **string** | Context engine ID | [optional] 
**ContextName** | Pointer to **string** | Context name | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**PrivPassword** | Pointer to **string** | Privacy protocol pass phrase | [optional] 
**PrivProtocol** | Pointer to **string** | Privacy protocol | [optional] 
**ReadCommunity** | Pointer to **string** | V2c specific | [optional] 
**SecurityEngineId** | Pointer to **string** | Security engine ID | [optional] 
**SecurityLevel** | Pointer to **string** | Security level | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**UserName** | Pointer to **string** | V3 specific | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**WriteCommunity** | Pointer to **string** |  | [optional] 

## Methods

### NewSnmpGateway

`func NewSnmpGateway() *SnmpGateway`

NewSnmpGateway instantiates a new SnmpGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpGatewayWithDefaults

`func NewSnmpGatewayWithDefaults() *SnmpGateway`

NewSnmpGatewayWithDefaults instantiates a new SnmpGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthPassword

`func (o *SnmpGateway) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *SnmpGateway) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *SnmpGateway) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *SnmpGateway) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### GetAuthProtocol

`func (o *SnmpGateway) GetAuthProtocol() string`

GetAuthProtocol returns the AuthProtocol field if non-nil, zero value otherwise.

### GetAuthProtocolOk

`func (o *SnmpGateway) GetAuthProtocolOk() (*string, bool)`

GetAuthProtocolOk returns a tuple with the AuthProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProtocol

`func (o *SnmpGateway) SetAuthProtocol(v string)`

SetAuthProtocol sets AuthProtocol field to given value.

### HasAuthProtocol

`func (o *SnmpGateway) HasAuthProtocol() bool`

HasAuthProtocol returns a boolean if a field has been set.

### GetContextEngineId

`func (o *SnmpGateway) GetContextEngineId() string`

GetContextEngineId returns the ContextEngineId field if non-nil, zero value otherwise.

### GetContextEngineIdOk

`func (o *SnmpGateway) GetContextEngineIdOk() (*string, bool)`

GetContextEngineIdOk returns a tuple with the ContextEngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextEngineId

`func (o *SnmpGateway) SetContextEngineId(v string)`

SetContextEngineId sets ContextEngineId field to given value.

### HasContextEngineId

`func (o *SnmpGateway) HasContextEngineId() bool`

HasContextEngineId returns a boolean if a field has been set.

### GetContextName

`func (o *SnmpGateway) GetContextName() string`

GetContextName returns the ContextName field if non-nil, zero value otherwise.

### GetContextNameOk

`func (o *SnmpGateway) GetContextNameOk() (*string, bool)`

GetContextNameOk returns a tuple with the ContextName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextName

`func (o *SnmpGateway) SetContextName(v string)`

SetContextName sets ContextName field to given value.

### HasContextName

`func (o *SnmpGateway) HasContextName() bool`

HasContextName returns a boolean if a field has been set.

### GetCreate

`func (o *SnmpGateway) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *SnmpGateway) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *SnmpGateway) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *SnmpGateway) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetHost

`func (o *SnmpGateway) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SnmpGateway) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SnmpGateway) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *SnmpGateway) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *SnmpGateway) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnmpGateway) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnmpGateway) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SnmpGateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *SnmpGateway) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SnmpGateway) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SnmpGateway) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SnmpGateway) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPort

`func (o *SnmpGateway) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SnmpGateway) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SnmpGateway) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SnmpGateway) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrivPassword

`func (o *SnmpGateway) GetPrivPassword() string`

GetPrivPassword returns the PrivPassword field if non-nil, zero value otherwise.

### GetPrivPasswordOk

`func (o *SnmpGateway) GetPrivPasswordOk() (*string, bool)`

GetPrivPasswordOk returns a tuple with the PrivPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivPassword

`func (o *SnmpGateway) SetPrivPassword(v string)`

SetPrivPassword sets PrivPassword field to given value.

### HasPrivPassword

`func (o *SnmpGateway) HasPrivPassword() bool`

HasPrivPassword returns a boolean if a field has been set.

### GetPrivProtocol

`func (o *SnmpGateway) GetPrivProtocol() string`

GetPrivProtocol returns the PrivProtocol field if non-nil, zero value otherwise.

### GetPrivProtocolOk

`func (o *SnmpGateway) GetPrivProtocolOk() (*string, bool)`

GetPrivProtocolOk returns a tuple with the PrivProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivProtocol

`func (o *SnmpGateway) SetPrivProtocol(v string)`

SetPrivProtocol sets PrivProtocol field to given value.

### HasPrivProtocol

`func (o *SnmpGateway) HasPrivProtocol() bool`

HasPrivProtocol returns a boolean if a field has been set.

### GetReadCommunity

`func (o *SnmpGateway) GetReadCommunity() string`

GetReadCommunity returns the ReadCommunity field if non-nil, zero value otherwise.

### GetReadCommunityOk

`func (o *SnmpGateway) GetReadCommunityOk() (*string, bool)`

GetReadCommunityOk returns a tuple with the ReadCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCommunity

`func (o *SnmpGateway) SetReadCommunity(v string)`

SetReadCommunity sets ReadCommunity field to given value.

### HasReadCommunity

`func (o *SnmpGateway) HasReadCommunity() bool`

HasReadCommunity returns a boolean if a field has been set.

### GetSecurityEngineId

`func (o *SnmpGateway) GetSecurityEngineId() string`

GetSecurityEngineId returns the SecurityEngineId field if non-nil, zero value otherwise.

### GetSecurityEngineIdOk

`func (o *SnmpGateway) GetSecurityEngineIdOk() (*string, bool)`

GetSecurityEngineIdOk returns a tuple with the SecurityEngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityEngineId

`func (o *SnmpGateway) SetSecurityEngineId(v string)`

SetSecurityEngineId sets SecurityEngineId field to given value.

### HasSecurityEngineId

`func (o *SnmpGateway) HasSecurityEngineId() bool`

HasSecurityEngineId returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *SnmpGateway) GetSecurityLevel() string`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *SnmpGateway) GetSecurityLevelOk() (*string, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *SnmpGateway) SetSecurityLevel(v string)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *SnmpGateway) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetStatus

`func (o *SnmpGateway) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SnmpGateway) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SnmpGateway) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SnmpGateway) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *SnmpGateway) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *SnmpGateway) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *SnmpGateway) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *SnmpGateway) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUserName

`func (o *SnmpGateway) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SnmpGateway) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SnmpGateway) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SnmpGateway) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetVersion

`func (o *SnmpGateway) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SnmpGateway) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SnmpGateway) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SnmpGateway) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWriteCommunity

`func (o *SnmpGateway) GetWriteCommunity() string`

GetWriteCommunity returns the WriteCommunity field if non-nil, zero value otherwise.

### GetWriteCommunityOk

`func (o *SnmpGateway) GetWriteCommunityOk() (*string, bool)`

GetWriteCommunityOk returns a tuple with the WriteCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteCommunity

`func (o *SnmpGateway) SetWriteCommunity(v string)`

SetWriteCommunity sets WriteCommunity field to given value.

### HasWriteCommunity

`func (o *SnmpGateway) HasWriteCommunity() bool`

HasWriteCommunity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


