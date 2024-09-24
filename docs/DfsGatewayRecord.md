# DfsGatewayRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**ConnNum** | Pointer to **int64** |  | [optional] 
**Cpus** | Pointer to **int64** | container resource limit | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**CtdbServiceStatus** | Pointer to **string** |  | [optional] 
**DfsGatewayGroup** | Pointer to [**DfsGatewayGroupNestview**](DfsGatewayGroupNestview.md) |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MemoryKbyte** | Pointer to **int64** |  | [optional] 
**S3Port** | Pointer to **int64** |  | [optional] 
**S3Scheme** | Pointer to **string** |  | [optional] 
**SslCertificate** | Pointer to [**SSLCertificateNestview**](SSLCertificateNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Samples** | Pointer to [**[]DfsGatewayStat**](DfsGatewayStat.md) |  | [optional] 

## Methods

### NewDfsGatewayRecord

`func NewDfsGatewayRecord() *DfsGatewayRecord`

NewDfsGatewayRecord instantiates a new DfsGatewayRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsGatewayRecordWithDefaults

`func NewDfsGatewayRecordWithDefaults() *DfsGatewayRecord`

NewDfsGatewayRecordWithDefaults instantiates a new DfsGatewayRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DfsGatewayRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsGatewayRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsGatewayRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsGatewayRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConnNum

`func (o *DfsGatewayRecord) GetConnNum() int64`

GetConnNum returns the ConnNum field if non-nil, zero value otherwise.

### GetConnNumOk

`func (o *DfsGatewayRecord) GetConnNumOk() (*int64, bool)`

GetConnNumOk returns a tuple with the ConnNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnNum

`func (o *DfsGatewayRecord) SetConnNum(v int64)`

SetConnNum sets ConnNum field to given value.

### HasConnNum

`func (o *DfsGatewayRecord) HasConnNum() bool`

HasConnNum returns a boolean if a field has been set.

### GetCpus

`func (o *DfsGatewayRecord) GetCpus() int64`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *DfsGatewayRecord) GetCpusOk() (*int64, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *DfsGatewayRecord) SetCpus(v int64)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *DfsGatewayRecord) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetCreate

`func (o *DfsGatewayRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsGatewayRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsGatewayRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsGatewayRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCtdbServiceStatus

`func (o *DfsGatewayRecord) GetCtdbServiceStatus() string`

GetCtdbServiceStatus returns the CtdbServiceStatus field if non-nil, zero value otherwise.

### GetCtdbServiceStatusOk

`func (o *DfsGatewayRecord) GetCtdbServiceStatusOk() (*string, bool)`

GetCtdbServiceStatusOk returns a tuple with the CtdbServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtdbServiceStatus

`func (o *DfsGatewayRecord) SetCtdbServiceStatus(v string)`

SetCtdbServiceStatus sets CtdbServiceStatus field to given value.

### HasCtdbServiceStatus

`func (o *DfsGatewayRecord) HasCtdbServiceStatus() bool`

HasCtdbServiceStatus returns a boolean if a field has been set.

### GetDfsGatewayGroup

`func (o *DfsGatewayRecord) GetDfsGatewayGroup() DfsGatewayGroupNestview`

GetDfsGatewayGroup returns the DfsGatewayGroup field if non-nil, zero value otherwise.

### GetDfsGatewayGroupOk

`func (o *DfsGatewayRecord) GetDfsGatewayGroupOk() (*DfsGatewayGroupNestview, bool)`

GetDfsGatewayGroupOk returns a tuple with the DfsGatewayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroup

`func (o *DfsGatewayRecord) SetDfsGatewayGroup(v DfsGatewayGroupNestview)`

SetDfsGatewayGroup sets DfsGatewayGroup field to given value.

### HasDfsGatewayGroup

`func (o *DfsGatewayRecord) HasDfsGatewayGroup() bool`

HasDfsGatewayGroup returns a boolean if a field has been set.

### GetHost

`func (o *DfsGatewayRecord) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DfsGatewayRecord) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DfsGatewayRecord) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *DfsGatewayRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *DfsGatewayRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsGatewayRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsGatewayRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsGatewayRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMemoryKbyte

`func (o *DfsGatewayRecord) GetMemoryKbyte() int64`

GetMemoryKbyte returns the MemoryKbyte field if non-nil, zero value otherwise.

### GetMemoryKbyteOk

`func (o *DfsGatewayRecord) GetMemoryKbyteOk() (*int64, bool)`

GetMemoryKbyteOk returns a tuple with the MemoryKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryKbyte

`func (o *DfsGatewayRecord) SetMemoryKbyte(v int64)`

SetMemoryKbyte sets MemoryKbyte field to given value.

### HasMemoryKbyte

`func (o *DfsGatewayRecord) HasMemoryKbyte() bool`

HasMemoryKbyte returns a boolean if a field has been set.

### GetS3Port

`func (o *DfsGatewayRecord) GetS3Port() int64`

GetS3Port returns the S3Port field if non-nil, zero value otherwise.

### GetS3PortOk

`func (o *DfsGatewayRecord) GetS3PortOk() (*int64, bool)`

GetS3PortOk returns a tuple with the S3Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Port

`func (o *DfsGatewayRecord) SetS3Port(v int64)`

SetS3Port sets S3Port field to given value.

### HasS3Port

`func (o *DfsGatewayRecord) HasS3Port() bool`

HasS3Port returns a boolean if a field has been set.

### GetS3Scheme

`func (o *DfsGatewayRecord) GetS3Scheme() string`

GetS3Scheme returns the S3Scheme field if non-nil, zero value otherwise.

### GetS3SchemeOk

`func (o *DfsGatewayRecord) GetS3SchemeOk() (*string, bool)`

GetS3SchemeOk returns a tuple with the S3Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Scheme

`func (o *DfsGatewayRecord) SetS3Scheme(v string)`

SetS3Scheme sets S3Scheme field to given value.

### HasS3Scheme

`func (o *DfsGatewayRecord) HasS3Scheme() bool`

HasS3Scheme returns a boolean if a field has been set.

### GetSslCertificate

`func (o *DfsGatewayRecord) GetSslCertificate() SSLCertificateNestview`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *DfsGatewayRecord) GetSslCertificateOk() (*SSLCertificateNestview, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *DfsGatewayRecord) SetSslCertificate(v SSLCertificateNestview)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *DfsGatewayRecord) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetStatus

`func (o *DfsGatewayRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsGatewayRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsGatewayRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsGatewayRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsGatewayRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsGatewayRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsGatewayRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsGatewayRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSamples

`func (o *DfsGatewayRecord) GetSamples() []DfsGatewayStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *DfsGatewayRecord) GetSamplesOk() (*[]DfsGatewayStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *DfsGatewayRecord) SetSamples(v []DfsGatewayStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *DfsGatewayRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


