# SnmpGatewayReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int64** |  | [optional] 
**AuthPassword** | Pointer to **string** |  | [optional] 
**AuthProtocol** | Pointer to **string** |  | [optional] 
**ContextEngineId** | Pointer to **string** |  | [optional] 
**ContextName** | Pointer to **string** |  | [optional] 
**HostId** | Pointer to **int64** |  | [optional] 
**PrivPassword** | Pointer to **string** |  | [optional] 
**PrivProtocol** | Pointer to **string** |  | [optional] 
**ReadCommunity** | Pointer to **string** |  | [optional] 
**SecurityEngineId** | Pointer to **string** |  | [optional] 
**SecurityLevel** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**WriteCommunity** | Pointer to **string** |  | [optional] 

## Methods

### NewSnmpGatewayReq

`func NewSnmpGatewayReq() *SnmpGatewayReq`

NewSnmpGatewayReq instantiates a new SnmpGatewayReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpGatewayReqWithDefaults

`func NewSnmpGatewayReqWithDefaults() *SnmpGatewayReq`

NewSnmpGatewayReqWithDefaults instantiates a new SnmpGatewayReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *SnmpGatewayReq) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SnmpGatewayReq) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SnmpGatewayReq) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SnmpGatewayReq) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAuthPassword

`func (o *SnmpGatewayReq) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *SnmpGatewayReq) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *SnmpGatewayReq) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *SnmpGatewayReq) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### GetAuthProtocol

`func (o *SnmpGatewayReq) GetAuthProtocol() string`

GetAuthProtocol returns the AuthProtocol field if non-nil, zero value otherwise.

### GetAuthProtocolOk

`func (o *SnmpGatewayReq) GetAuthProtocolOk() (*string, bool)`

GetAuthProtocolOk returns a tuple with the AuthProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProtocol

`func (o *SnmpGatewayReq) SetAuthProtocol(v string)`

SetAuthProtocol sets AuthProtocol field to given value.

### HasAuthProtocol

`func (o *SnmpGatewayReq) HasAuthProtocol() bool`

HasAuthProtocol returns a boolean if a field has been set.

### GetContextEngineId

`func (o *SnmpGatewayReq) GetContextEngineId() string`

GetContextEngineId returns the ContextEngineId field if non-nil, zero value otherwise.

### GetContextEngineIdOk

`func (o *SnmpGatewayReq) GetContextEngineIdOk() (*string, bool)`

GetContextEngineIdOk returns a tuple with the ContextEngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextEngineId

`func (o *SnmpGatewayReq) SetContextEngineId(v string)`

SetContextEngineId sets ContextEngineId field to given value.

### HasContextEngineId

`func (o *SnmpGatewayReq) HasContextEngineId() bool`

HasContextEngineId returns a boolean if a field has been set.

### GetContextName

`func (o *SnmpGatewayReq) GetContextName() string`

GetContextName returns the ContextName field if non-nil, zero value otherwise.

### GetContextNameOk

`func (o *SnmpGatewayReq) GetContextNameOk() (*string, bool)`

GetContextNameOk returns a tuple with the ContextName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextName

`func (o *SnmpGatewayReq) SetContextName(v string)`

SetContextName sets ContextName field to given value.

### HasContextName

`func (o *SnmpGatewayReq) HasContextName() bool`

HasContextName returns a boolean if a field has been set.

### GetHostId

`func (o *SnmpGatewayReq) GetHostId() int64`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *SnmpGatewayReq) GetHostIdOk() (*int64, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *SnmpGatewayReq) SetHostId(v int64)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *SnmpGatewayReq) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetPrivPassword

`func (o *SnmpGatewayReq) GetPrivPassword() string`

GetPrivPassword returns the PrivPassword field if non-nil, zero value otherwise.

### GetPrivPasswordOk

`func (o *SnmpGatewayReq) GetPrivPasswordOk() (*string, bool)`

GetPrivPasswordOk returns a tuple with the PrivPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivPassword

`func (o *SnmpGatewayReq) SetPrivPassword(v string)`

SetPrivPassword sets PrivPassword field to given value.

### HasPrivPassword

`func (o *SnmpGatewayReq) HasPrivPassword() bool`

HasPrivPassword returns a boolean if a field has been set.

### GetPrivProtocol

`func (o *SnmpGatewayReq) GetPrivProtocol() string`

GetPrivProtocol returns the PrivProtocol field if non-nil, zero value otherwise.

### GetPrivProtocolOk

`func (o *SnmpGatewayReq) GetPrivProtocolOk() (*string, bool)`

GetPrivProtocolOk returns a tuple with the PrivProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivProtocol

`func (o *SnmpGatewayReq) SetPrivProtocol(v string)`

SetPrivProtocol sets PrivProtocol field to given value.

### HasPrivProtocol

`func (o *SnmpGatewayReq) HasPrivProtocol() bool`

HasPrivProtocol returns a boolean if a field has been set.

### GetReadCommunity

`func (o *SnmpGatewayReq) GetReadCommunity() string`

GetReadCommunity returns the ReadCommunity field if non-nil, zero value otherwise.

### GetReadCommunityOk

`func (o *SnmpGatewayReq) GetReadCommunityOk() (*string, bool)`

GetReadCommunityOk returns a tuple with the ReadCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCommunity

`func (o *SnmpGatewayReq) SetReadCommunity(v string)`

SetReadCommunity sets ReadCommunity field to given value.

### HasReadCommunity

`func (o *SnmpGatewayReq) HasReadCommunity() bool`

HasReadCommunity returns a boolean if a field has been set.

### GetSecurityEngineId

`func (o *SnmpGatewayReq) GetSecurityEngineId() string`

GetSecurityEngineId returns the SecurityEngineId field if non-nil, zero value otherwise.

### GetSecurityEngineIdOk

`func (o *SnmpGatewayReq) GetSecurityEngineIdOk() (*string, bool)`

GetSecurityEngineIdOk returns a tuple with the SecurityEngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityEngineId

`func (o *SnmpGatewayReq) SetSecurityEngineId(v string)`

SetSecurityEngineId sets SecurityEngineId field to given value.

### HasSecurityEngineId

`func (o *SnmpGatewayReq) HasSecurityEngineId() bool`

HasSecurityEngineId returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *SnmpGatewayReq) GetSecurityLevel() string`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *SnmpGatewayReq) GetSecurityLevelOk() (*string, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *SnmpGatewayReq) SetSecurityLevel(v string)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *SnmpGatewayReq) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetUserName

`func (o *SnmpGatewayReq) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SnmpGatewayReq) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SnmpGatewayReq) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SnmpGatewayReq) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetVersion

`func (o *SnmpGatewayReq) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SnmpGatewayReq) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SnmpGatewayReq) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SnmpGatewayReq) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWriteCommunity

`func (o *SnmpGatewayReq) GetWriteCommunity() string`

GetWriteCommunity returns the WriteCommunity field if non-nil, zero value otherwise.

### GetWriteCommunityOk

`func (o *SnmpGatewayReq) GetWriteCommunityOk() (*string, bool)`

GetWriteCommunityOk returns a tuple with the WriteCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteCommunity

`func (o *SnmpGatewayReq) SetWriteCommunity(v string)`

SetWriteCommunity sets WriteCommunity field to given value.

### HasWriteCommunity

`func (o *SnmpGatewayReq) HasWriteCommunity() bool`

HasWriteCommunity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


