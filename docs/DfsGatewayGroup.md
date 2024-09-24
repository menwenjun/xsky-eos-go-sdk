# DfsGatewayGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Ad** | Pointer to [**FSActiveDirectoryNestview**](FSActiveDirectoryNestview.md) |  | [optional] 
**BucketNum** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Cpus** | Pointer to **int64** | container resource limit | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Encoding** | Pointer to **string** |  | [optional] 
**GatewayNum** | Pointer to **int64** |  | [optional] 
**HdfsNum** | Pointer to **int64** |  | [optional] 
**HdfsSecurities** | Pointer to **[]string** | hdfs domain auth type | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Ldap** | Pointer to [**FSLdapNestview**](FSLdapNestview.md) |  | [optional] 
**MemoryKbyte** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NfsVersions** | Pointer to **[]string** | NFS attributes | [optional] 
**Port** | Pointer to **int64** | FTP attributes | [optional] 
**ProtocolAuthType** | Pointer to **map[string][]string** |  | [optional] 
**RootfsNum** | Pointer to **int64** |  | [optional] 
**S3Port** | Pointer to **int64** |  | [optional] 
**S3Scheme** | Pointer to **string** |  | [optional] 
**S3UserNum** | Pointer to **int64** |  | [optional] 
**Securities** | Pointer to **[]string** |  | [optional] 
**ShareNums** | Pointer to **map[string]int64** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Smb1Enabled** | Pointer to **bool** |  | [optional] 
**SmbBrowseable** | Pointer to **bool** |  | [optional] 
**SmbHomes** | Pointer to **bool** |  | [optional] 
**SmbPorts** | Pointer to **[]int64** |  | [optional] 
**SmbType** | Pointer to **string** | SMB attributes | [optional] 
**SslCertificate** | Pointer to [**SSLCertificateNestview**](SSLCertificateNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Types** | Pointer to **[]string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**ZoneNum** | Pointer to **int64** |  | [optional] 

## Methods

### NewDfsGatewayGroup

`func NewDfsGatewayGroup() *DfsGatewayGroup`

NewDfsGatewayGroup instantiates a new DfsGatewayGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsGatewayGroupWithDefaults

`func NewDfsGatewayGroupWithDefaults() *DfsGatewayGroup`

NewDfsGatewayGroupWithDefaults instantiates a new DfsGatewayGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DfsGatewayGroup) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsGatewayGroup) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsGatewayGroup) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsGatewayGroup) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAd

`func (o *DfsGatewayGroup) GetAd() FSActiveDirectoryNestview`

GetAd returns the Ad field if non-nil, zero value otherwise.

### GetAdOk

`func (o *DfsGatewayGroup) GetAdOk() (*FSActiveDirectoryNestview, bool)`

GetAdOk returns a tuple with the Ad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAd

`func (o *DfsGatewayGroup) SetAd(v FSActiveDirectoryNestview)`

SetAd sets Ad field to given value.

### HasAd

`func (o *DfsGatewayGroup) HasAd() bool`

HasAd returns a boolean if a field has been set.

### GetBucketNum

`func (o *DfsGatewayGroup) GetBucketNum() int64`

GetBucketNum returns the BucketNum field if non-nil, zero value otherwise.

### GetBucketNumOk

`func (o *DfsGatewayGroup) GetBucketNumOk() (*int64, bool)`

GetBucketNumOk returns a tuple with the BucketNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketNum

`func (o *DfsGatewayGroup) SetBucketNum(v int64)`

SetBucketNum sets BucketNum field to given value.

### HasBucketNum

`func (o *DfsGatewayGroup) HasBucketNum() bool`

HasBucketNum returns a boolean if a field has been set.

### GetCluster

`func (o *DfsGatewayGroup) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsGatewayGroup) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsGatewayGroup) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsGatewayGroup) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCpus

`func (o *DfsGatewayGroup) GetCpus() int64`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *DfsGatewayGroup) GetCpusOk() (*int64, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *DfsGatewayGroup) SetCpus(v int64)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *DfsGatewayGroup) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetCreate

`func (o *DfsGatewayGroup) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsGatewayGroup) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsGatewayGroup) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsGatewayGroup) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *DfsGatewayGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsGatewayGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsGatewayGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsGatewayGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEncoding

`func (o *DfsGatewayGroup) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *DfsGatewayGroup) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *DfsGatewayGroup) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *DfsGatewayGroup) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetGatewayNum

`func (o *DfsGatewayGroup) GetGatewayNum() int64`

GetGatewayNum returns the GatewayNum field if non-nil, zero value otherwise.

### GetGatewayNumOk

`func (o *DfsGatewayGroup) GetGatewayNumOk() (*int64, bool)`

GetGatewayNumOk returns a tuple with the GatewayNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayNum

`func (o *DfsGatewayGroup) SetGatewayNum(v int64)`

SetGatewayNum sets GatewayNum field to given value.

### HasGatewayNum

`func (o *DfsGatewayGroup) HasGatewayNum() bool`

HasGatewayNum returns a boolean if a field has been set.

### GetHdfsNum

`func (o *DfsGatewayGroup) GetHdfsNum() int64`

GetHdfsNum returns the HdfsNum field if non-nil, zero value otherwise.

### GetHdfsNumOk

`func (o *DfsGatewayGroup) GetHdfsNumOk() (*int64, bool)`

GetHdfsNumOk returns a tuple with the HdfsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsNum

`func (o *DfsGatewayGroup) SetHdfsNum(v int64)`

SetHdfsNum sets HdfsNum field to given value.

### HasHdfsNum

`func (o *DfsGatewayGroup) HasHdfsNum() bool`

HasHdfsNum returns a boolean if a field has been set.

### GetHdfsSecurities

`func (o *DfsGatewayGroup) GetHdfsSecurities() []string`

GetHdfsSecurities returns the HdfsSecurities field if non-nil, zero value otherwise.

### GetHdfsSecuritiesOk

`func (o *DfsGatewayGroup) GetHdfsSecuritiesOk() (*[]string, bool)`

GetHdfsSecuritiesOk returns a tuple with the HdfsSecurities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsSecurities

`func (o *DfsGatewayGroup) SetHdfsSecurities(v []string)`

SetHdfsSecurities sets HdfsSecurities field to given value.

### HasHdfsSecurities

`func (o *DfsGatewayGroup) HasHdfsSecurities() bool`

HasHdfsSecurities returns a boolean if a field has been set.

### GetId

`func (o *DfsGatewayGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsGatewayGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsGatewayGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsGatewayGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLdap

`func (o *DfsGatewayGroup) GetLdap() FSLdapNestview`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *DfsGatewayGroup) GetLdapOk() (*FSLdapNestview, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *DfsGatewayGroup) SetLdap(v FSLdapNestview)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *DfsGatewayGroup) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetMemoryKbyte

`func (o *DfsGatewayGroup) GetMemoryKbyte() int64`

GetMemoryKbyte returns the MemoryKbyte field if non-nil, zero value otherwise.

### GetMemoryKbyteOk

`func (o *DfsGatewayGroup) GetMemoryKbyteOk() (*int64, bool)`

GetMemoryKbyteOk returns a tuple with the MemoryKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryKbyte

`func (o *DfsGatewayGroup) SetMemoryKbyte(v int64)`

SetMemoryKbyte sets MemoryKbyte field to given value.

### HasMemoryKbyte

`func (o *DfsGatewayGroup) HasMemoryKbyte() bool`

HasMemoryKbyte returns a boolean if a field has been set.

### GetName

`func (o *DfsGatewayGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsGatewayGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsGatewayGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsGatewayGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNfsVersions

`func (o *DfsGatewayGroup) GetNfsVersions() []string`

GetNfsVersions returns the NfsVersions field if non-nil, zero value otherwise.

### GetNfsVersionsOk

`func (o *DfsGatewayGroup) GetNfsVersionsOk() (*[]string, bool)`

GetNfsVersionsOk returns a tuple with the NfsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsVersions

`func (o *DfsGatewayGroup) SetNfsVersions(v []string)`

SetNfsVersions sets NfsVersions field to given value.

### HasNfsVersions

`func (o *DfsGatewayGroup) HasNfsVersions() bool`

HasNfsVersions returns a boolean if a field has been set.

### GetPort

`func (o *DfsGatewayGroup) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DfsGatewayGroup) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DfsGatewayGroup) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DfsGatewayGroup) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocolAuthType

`func (o *DfsGatewayGroup) GetProtocolAuthType() map[string][]string`

GetProtocolAuthType returns the ProtocolAuthType field if non-nil, zero value otherwise.

### GetProtocolAuthTypeOk

`func (o *DfsGatewayGroup) GetProtocolAuthTypeOk() (*map[string][]string, bool)`

GetProtocolAuthTypeOk returns a tuple with the ProtocolAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolAuthType

`func (o *DfsGatewayGroup) SetProtocolAuthType(v map[string][]string)`

SetProtocolAuthType sets ProtocolAuthType field to given value.

### HasProtocolAuthType

`func (o *DfsGatewayGroup) HasProtocolAuthType() bool`

HasProtocolAuthType returns a boolean if a field has been set.

### GetRootfsNum

`func (o *DfsGatewayGroup) GetRootfsNum() int64`

GetRootfsNum returns the RootfsNum field if non-nil, zero value otherwise.

### GetRootfsNumOk

`func (o *DfsGatewayGroup) GetRootfsNumOk() (*int64, bool)`

GetRootfsNumOk returns a tuple with the RootfsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfsNum

`func (o *DfsGatewayGroup) SetRootfsNum(v int64)`

SetRootfsNum sets RootfsNum field to given value.

### HasRootfsNum

`func (o *DfsGatewayGroup) HasRootfsNum() bool`

HasRootfsNum returns a boolean if a field has been set.

### GetS3Port

`func (o *DfsGatewayGroup) GetS3Port() int64`

GetS3Port returns the S3Port field if non-nil, zero value otherwise.

### GetS3PortOk

`func (o *DfsGatewayGroup) GetS3PortOk() (*int64, bool)`

GetS3PortOk returns a tuple with the S3Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Port

`func (o *DfsGatewayGroup) SetS3Port(v int64)`

SetS3Port sets S3Port field to given value.

### HasS3Port

`func (o *DfsGatewayGroup) HasS3Port() bool`

HasS3Port returns a boolean if a field has been set.

### GetS3Scheme

`func (o *DfsGatewayGroup) GetS3Scheme() string`

GetS3Scheme returns the S3Scheme field if non-nil, zero value otherwise.

### GetS3SchemeOk

`func (o *DfsGatewayGroup) GetS3SchemeOk() (*string, bool)`

GetS3SchemeOk returns a tuple with the S3Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Scheme

`func (o *DfsGatewayGroup) SetS3Scheme(v string)`

SetS3Scheme sets S3Scheme field to given value.

### HasS3Scheme

`func (o *DfsGatewayGroup) HasS3Scheme() bool`

HasS3Scheme returns a boolean if a field has been set.

### GetS3UserNum

`func (o *DfsGatewayGroup) GetS3UserNum() int64`

GetS3UserNum returns the S3UserNum field if non-nil, zero value otherwise.

### GetS3UserNumOk

`func (o *DfsGatewayGroup) GetS3UserNumOk() (*int64, bool)`

GetS3UserNumOk returns a tuple with the S3UserNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3UserNum

`func (o *DfsGatewayGroup) SetS3UserNum(v int64)`

SetS3UserNum sets S3UserNum field to given value.

### HasS3UserNum

`func (o *DfsGatewayGroup) HasS3UserNum() bool`

HasS3UserNum returns a boolean if a field has been set.

### GetSecurities

`func (o *DfsGatewayGroup) GetSecurities() []string`

GetSecurities returns the Securities field if non-nil, zero value otherwise.

### GetSecuritiesOk

`func (o *DfsGatewayGroup) GetSecuritiesOk() (*[]string, bool)`

GetSecuritiesOk returns a tuple with the Securities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurities

`func (o *DfsGatewayGroup) SetSecurities(v []string)`

SetSecurities sets Securities field to given value.

### HasSecurities

`func (o *DfsGatewayGroup) HasSecurities() bool`

HasSecurities returns a boolean if a field has been set.

### GetShareNums

`func (o *DfsGatewayGroup) GetShareNums() map[string]int64`

GetShareNums returns the ShareNums field if non-nil, zero value otherwise.

### GetShareNumsOk

`func (o *DfsGatewayGroup) GetShareNumsOk() (*map[string]int64, bool)`

GetShareNumsOk returns a tuple with the ShareNums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareNums

`func (o *DfsGatewayGroup) SetShareNums(v map[string]int64)`

SetShareNums sets ShareNums field to given value.

### HasShareNums

`func (o *DfsGatewayGroup) HasShareNums() bool`

HasShareNums returns a boolean if a field has been set.

### GetShared

`func (o *DfsGatewayGroup) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *DfsGatewayGroup) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *DfsGatewayGroup) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *DfsGatewayGroup) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSmb1Enabled

`func (o *DfsGatewayGroup) GetSmb1Enabled() bool`

GetSmb1Enabled returns the Smb1Enabled field if non-nil, zero value otherwise.

### GetSmb1EnabledOk

`func (o *DfsGatewayGroup) GetSmb1EnabledOk() (*bool, bool)`

GetSmb1EnabledOk returns a tuple with the Smb1Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmb1Enabled

`func (o *DfsGatewayGroup) SetSmb1Enabled(v bool)`

SetSmb1Enabled sets Smb1Enabled field to given value.

### HasSmb1Enabled

`func (o *DfsGatewayGroup) HasSmb1Enabled() bool`

HasSmb1Enabled returns a boolean if a field has been set.

### GetSmbBrowseable

`func (o *DfsGatewayGroup) GetSmbBrowseable() bool`

GetSmbBrowseable returns the SmbBrowseable field if non-nil, zero value otherwise.

### GetSmbBrowseableOk

`func (o *DfsGatewayGroup) GetSmbBrowseableOk() (*bool, bool)`

GetSmbBrowseableOk returns a tuple with the SmbBrowseable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbBrowseable

`func (o *DfsGatewayGroup) SetSmbBrowseable(v bool)`

SetSmbBrowseable sets SmbBrowseable field to given value.

### HasSmbBrowseable

`func (o *DfsGatewayGroup) HasSmbBrowseable() bool`

HasSmbBrowseable returns a boolean if a field has been set.

### GetSmbHomes

`func (o *DfsGatewayGroup) GetSmbHomes() bool`

GetSmbHomes returns the SmbHomes field if non-nil, zero value otherwise.

### GetSmbHomesOk

`func (o *DfsGatewayGroup) GetSmbHomesOk() (*bool, bool)`

GetSmbHomesOk returns a tuple with the SmbHomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbHomes

`func (o *DfsGatewayGroup) SetSmbHomes(v bool)`

SetSmbHomes sets SmbHomes field to given value.

### HasSmbHomes

`func (o *DfsGatewayGroup) HasSmbHomes() bool`

HasSmbHomes returns a boolean if a field has been set.

### GetSmbPorts

`func (o *DfsGatewayGroup) GetSmbPorts() []int64`

GetSmbPorts returns the SmbPorts field if non-nil, zero value otherwise.

### GetSmbPortsOk

`func (o *DfsGatewayGroup) GetSmbPortsOk() (*[]int64, bool)`

GetSmbPortsOk returns a tuple with the SmbPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbPorts

`func (o *DfsGatewayGroup) SetSmbPorts(v []int64)`

SetSmbPorts sets SmbPorts field to given value.

### HasSmbPorts

`func (o *DfsGatewayGroup) HasSmbPorts() bool`

HasSmbPorts returns a boolean if a field has been set.

### GetSmbType

`func (o *DfsGatewayGroup) GetSmbType() string`

GetSmbType returns the SmbType field if non-nil, zero value otherwise.

### GetSmbTypeOk

`func (o *DfsGatewayGroup) GetSmbTypeOk() (*string, bool)`

GetSmbTypeOk returns a tuple with the SmbType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbType

`func (o *DfsGatewayGroup) SetSmbType(v string)`

SetSmbType sets SmbType field to given value.

### HasSmbType

`func (o *DfsGatewayGroup) HasSmbType() bool`

HasSmbType returns a boolean if a field has been set.

### GetSslCertificate

`func (o *DfsGatewayGroup) GetSslCertificate() SSLCertificateNestview`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *DfsGatewayGroup) GetSslCertificateOk() (*SSLCertificateNestview, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *DfsGatewayGroup) SetSslCertificate(v SSLCertificateNestview)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *DfsGatewayGroup) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetStatus

`func (o *DfsGatewayGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsGatewayGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsGatewayGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsGatewayGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTypes

`func (o *DfsGatewayGroup) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *DfsGatewayGroup) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *DfsGatewayGroup) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *DfsGatewayGroup) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsGatewayGroup) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsGatewayGroup) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsGatewayGroup) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsGatewayGroup) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetZoneNum

`func (o *DfsGatewayGroup) GetZoneNum() int64`

GetZoneNum returns the ZoneNum field if non-nil, zero value otherwise.

### GetZoneNumOk

`func (o *DfsGatewayGroup) GetZoneNumOk() (*int64, bool)`

GetZoneNumOk returns a tuple with the ZoneNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNum

`func (o *DfsGatewayGroup) SetZoneNum(v int64)`

SetZoneNum sets ZoneNum field to given value.

### HasZoneNum

`func (o *DfsGatewayGroup) HasZoneNum() bool`

HasZoneNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


