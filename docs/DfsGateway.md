# DfsGateway

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

## Methods

### NewDfsGateway

`func NewDfsGateway() *DfsGateway`

NewDfsGateway instantiates a new DfsGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsGatewayWithDefaults

`func NewDfsGatewayWithDefaults() *DfsGateway`

NewDfsGatewayWithDefaults instantiates a new DfsGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DfsGateway) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsGateway) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsGateway) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsGateway) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConnNum

`func (o *DfsGateway) GetConnNum() int64`

GetConnNum returns the ConnNum field if non-nil, zero value otherwise.

### GetConnNumOk

`func (o *DfsGateway) GetConnNumOk() (*int64, bool)`

GetConnNumOk returns a tuple with the ConnNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnNum

`func (o *DfsGateway) SetConnNum(v int64)`

SetConnNum sets ConnNum field to given value.

### HasConnNum

`func (o *DfsGateway) HasConnNum() bool`

HasConnNum returns a boolean if a field has been set.

### GetCpus

`func (o *DfsGateway) GetCpus() int64`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *DfsGateway) GetCpusOk() (*int64, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *DfsGateway) SetCpus(v int64)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *DfsGateway) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetCreate

`func (o *DfsGateway) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsGateway) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsGateway) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsGateway) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCtdbServiceStatus

`func (o *DfsGateway) GetCtdbServiceStatus() string`

GetCtdbServiceStatus returns the CtdbServiceStatus field if non-nil, zero value otherwise.

### GetCtdbServiceStatusOk

`func (o *DfsGateway) GetCtdbServiceStatusOk() (*string, bool)`

GetCtdbServiceStatusOk returns a tuple with the CtdbServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtdbServiceStatus

`func (o *DfsGateway) SetCtdbServiceStatus(v string)`

SetCtdbServiceStatus sets CtdbServiceStatus field to given value.

### HasCtdbServiceStatus

`func (o *DfsGateway) HasCtdbServiceStatus() bool`

HasCtdbServiceStatus returns a boolean if a field has been set.

### GetDfsGatewayGroup

`func (o *DfsGateway) GetDfsGatewayGroup() DfsGatewayGroupNestview`

GetDfsGatewayGroup returns the DfsGatewayGroup field if non-nil, zero value otherwise.

### GetDfsGatewayGroupOk

`func (o *DfsGateway) GetDfsGatewayGroupOk() (*DfsGatewayGroupNestview, bool)`

GetDfsGatewayGroupOk returns a tuple with the DfsGatewayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroup

`func (o *DfsGateway) SetDfsGatewayGroup(v DfsGatewayGroupNestview)`

SetDfsGatewayGroup sets DfsGatewayGroup field to given value.

### HasDfsGatewayGroup

`func (o *DfsGateway) HasDfsGatewayGroup() bool`

HasDfsGatewayGroup returns a boolean if a field has been set.

### GetHost

`func (o *DfsGateway) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DfsGateway) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DfsGateway) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *DfsGateway) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *DfsGateway) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsGateway) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsGateway) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsGateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMemoryKbyte

`func (o *DfsGateway) GetMemoryKbyte() int64`

GetMemoryKbyte returns the MemoryKbyte field if non-nil, zero value otherwise.

### GetMemoryKbyteOk

`func (o *DfsGateway) GetMemoryKbyteOk() (*int64, bool)`

GetMemoryKbyteOk returns a tuple with the MemoryKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryKbyte

`func (o *DfsGateway) SetMemoryKbyte(v int64)`

SetMemoryKbyte sets MemoryKbyte field to given value.

### HasMemoryKbyte

`func (o *DfsGateway) HasMemoryKbyte() bool`

HasMemoryKbyte returns a boolean if a field has been set.

### GetS3Port

`func (o *DfsGateway) GetS3Port() int64`

GetS3Port returns the S3Port field if non-nil, zero value otherwise.

### GetS3PortOk

`func (o *DfsGateway) GetS3PortOk() (*int64, bool)`

GetS3PortOk returns a tuple with the S3Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Port

`func (o *DfsGateway) SetS3Port(v int64)`

SetS3Port sets S3Port field to given value.

### HasS3Port

`func (o *DfsGateway) HasS3Port() bool`

HasS3Port returns a boolean if a field has been set.

### GetS3Scheme

`func (o *DfsGateway) GetS3Scheme() string`

GetS3Scheme returns the S3Scheme field if non-nil, zero value otherwise.

### GetS3SchemeOk

`func (o *DfsGateway) GetS3SchemeOk() (*string, bool)`

GetS3SchemeOk returns a tuple with the S3Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Scheme

`func (o *DfsGateway) SetS3Scheme(v string)`

SetS3Scheme sets S3Scheme field to given value.

### HasS3Scheme

`func (o *DfsGateway) HasS3Scheme() bool`

HasS3Scheme returns a boolean if a field has been set.

### GetSslCertificate

`func (o *DfsGateway) GetSslCertificate() SSLCertificateNestview`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *DfsGateway) GetSslCertificateOk() (*SSLCertificateNestview, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *DfsGateway) SetSslCertificate(v SSLCertificateNestview)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *DfsGateway) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetStatus

`func (o *DfsGateway) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsGateway) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsGateway) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsGateway) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsGateway) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsGateway) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsGateway) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsGateway) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


