# DfsGatewayGroupUpdateTypesReqGatewayGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdId** | Pointer to **int64** | active directory id | [optional] 
**Encoding** | Pointer to **string** | ftp encoding format, default is utf8 | [optional] 
**HdfsSecurities** | Pointer to **[]string** | security type for hdfs (local, ldap) | [optional] 
**LdapId** | Pointer to **int64** | ldap id | [optional] 
**NfsVersions** | Pointer to **[]string** | nfs versions supported | [optional] 
**Securities** | Pointer to **[]string** | security type for smb/quota (local, ad, ldap) | [optional] 
**Smb1Enabled** | Pointer to **bool** | smb version 1.0 enabled | [optional] 
**SmbBrowseable** | Pointer to **bool** | smb browseable enable | [optional] 
**SmbHomes** | Pointer to **bool** | smb Homes share enable | [optional] 
**SmbPorts** | Pointer to **[]int64** | smb ports | [optional] 
**SmbType** | Pointer to **string** | smb service type (smb, xsmb) | [optional] 
**Types** | Pointer to **[]string** | types of supported | [optional] 

## Methods

### NewDfsGatewayGroupUpdateTypesReqGatewayGroup

`func NewDfsGatewayGroupUpdateTypesReqGatewayGroup() *DfsGatewayGroupUpdateTypesReqGatewayGroup`

NewDfsGatewayGroupUpdateTypesReqGatewayGroup instantiates a new DfsGatewayGroupUpdateTypesReqGatewayGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsGatewayGroupUpdateTypesReqGatewayGroupWithDefaults

`func NewDfsGatewayGroupUpdateTypesReqGatewayGroupWithDefaults() *DfsGatewayGroupUpdateTypesReqGatewayGroup`

NewDfsGatewayGroupUpdateTypesReqGatewayGroupWithDefaults instantiates a new DfsGatewayGroupUpdateTypesReqGatewayGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdId

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetAdId() int64`

GetAdId returns the AdId field if non-nil, zero value otherwise.

### GetAdIdOk

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetAdIdOk() (*int64, bool)`

GetAdIdOk returns a tuple with the AdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdId

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) SetAdId(v int64)`

SetAdId sets AdId field to given value.

### HasAdId

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) HasAdId() bool`

HasAdId returns a boolean if a field has been set.

### GetEncoding

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetHdfsSecurities

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetHdfsSecurities() []string`

GetHdfsSecurities returns the HdfsSecurities field if non-nil, zero value otherwise.

### GetHdfsSecuritiesOk

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetHdfsSecuritiesOk() (*[]string, bool)`

GetHdfsSecuritiesOk returns a tuple with the HdfsSecurities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsSecurities

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) SetHdfsSecurities(v []string)`

SetHdfsSecurities sets HdfsSecurities field to given value.

### HasHdfsSecurities

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) HasHdfsSecurities() bool`

HasHdfsSecurities returns a boolean if a field has been set.

### GetLdapId

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetLdapId() int64`

GetLdapId returns the LdapId field if non-nil, zero value otherwise.

### GetLdapIdOk

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetLdapIdOk() (*int64, bool)`

GetLdapIdOk returns a tuple with the LdapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapId

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) SetLdapId(v int64)`

SetLdapId sets LdapId field to given value.

### HasLdapId

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) HasLdapId() bool`

HasLdapId returns a boolean if a field has been set.

### GetNfsVersions

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetNfsVersions() []string`

GetNfsVersions returns the NfsVersions field if non-nil, zero value otherwise.

### GetNfsVersionsOk

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetNfsVersionsOk() (*[]string, bool)`

GetNfsVersionsOk returns a tuple with the NfsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsVersions

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) SetNfsVersions(v []string)`

SetNfsVersions sets NfsVersions field to given value.

### HasNfsVersions

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) HasNfsVersions() bool`

HasNfsVersions returns a boolean if a field has been set.

### GetSecurities

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetSecurities() []string`

GetSecurities returns the Securities field if non-nil, zero value otherwise.

### GetSecuritiesOk

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetSecuritiesOk() (*[]string, bool)`

GetSecuritiesOk returns a tuple with the Securities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurities

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) SetSecurities(v []string)`

SetSecurities sets Securities field to given value.

### HasSecurities

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) HasSecurities() bool`

HasSecurities returns a boolean if a field has been set.

### GetSmb1Enabled

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetSmb1Enabled() bool`

GetSmb1Enabled returns the Smb1Enabled field if non-nil, zero value otherwise.

### GetSmb1EnabledOk

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetSmb1EnabledOk() (*bool, bool)`

GetSmb1EnabledOk returns a tuple with the Smb1Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmb1Enabled

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) SetSmb1Enabled(v bool)`

SetSmb1Enabled sets Smb1Enabled field to given value.

### HasSmb1Enabled

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) HasSmb1Enabled() bool`

HasSmb1Enabled returns a boolean if a field has been set.

### GetSmbBrowseable

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetSmbBrowseable() bool`

GetSmbBrowseable returns the SmbBrowseable field if non-nil, zero value otherwise.

### GetSmbBrowseableOk

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetSmbBrowseableOk() (*bool, bool)`

GetSmbBrowseableOk returns a tuple with the SmbBrowseable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbBrowseable

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) SetSmbBrowseable(v bool)`

SetSmbBrowseable sets SmbBrowseable field to given value.

### HasSmbBrowseable

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) HasSmbBrowseable() bool`

HasSmbBrowseable returns a boolean if a field has been set.

### GetSmbHomes

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetSmbHomes() bool`

GetSmbHomes returns the SmbHomes field if non-nil, zero value otherwise.

### GetSmbHomesOk

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetSmbHomesOk() (*bool, bool)`

GetSmbHomesOk returns a tuple with the SmbHomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbHomes

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) SetSmbHomes(v bool)`

SetSmbHomes sets SmbHomes field to given value.

### HasSmbHomes

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) HasSmbHomes() bool`

HasSmbHomes returns a boolean if a field has been set.

### GetSmbPorts

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetSmbPorts() []int64`

GetSmbPorts returns the SmbPorts field if non-nil, zero value otherwise.

### GetSmbPortsOk

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetSmbPortsOk() (*[]int64, bool)`

GetSmbPortsOk returns a tuple with the SmbPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbPorts

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) SetSmbPorts(v []int64)`

SetSmbPorts sets SmbPorts field to given value.

### HasSmbPorts

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) HasSmbPorts() bool`

HasSmbPorts returns a boolean if a field has been set.

### GetSmbType

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetSmbType() string`

GetSmbType returns the SmbType field if non-nil, zero value otherwise.

### GetSmbTypeOk

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetSmbTypeOk() (*string, bool)`

GetSmbTypeOk returns a tuple with the SmbType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbType

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) SetSmbType(v string)`

SetSmbType sets SmbType field to given value.

### HasSmbType

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) HasSmbType() bool`

HasSmbType returns a boolean if a field has been set.

### GetTypes

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *DfsGatewayGroupUpdateTypesReqGatewayGroup) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


