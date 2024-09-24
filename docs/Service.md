# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Etag** | Pointer to **string** |  | [optional] 
**Heartbeat** | Pointer to **time.Time** |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Pid** | Pointer to **int32** |  | [optional] 
**Primary** | Pointer to **bool** |  | [optional] 
**SslCertificate** | Pointer to [**SSLCertificateNestview**](SSLCertificateNestview.md) |  | [optional] 
**StartedTime** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Up** | Pointer to **bool** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewService

`func NewService() *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *Service) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Service) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Service) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Service) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *Service) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *Service) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *Service) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *Service) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetEnabled

`func (o *Service) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Service) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Service) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Service) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEtag

`func (o *Service) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *Service) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *Service) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *Service) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetHeartbeat

`func (o *Service) GetHeartbeat() time.Time`

GetHeartbeat returns the Heartbeat field if non-nil, zero value otherwise.

### GetHeartbeatOk

`func (o *Service) GetHeartbeatOk() (*time.Time, bool)`

GetHeartbeatOk returns a tuple with the Heartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeat

`func (o *Service) SetHeartbeat(v time.Time)`

SetHeartbeat sets Heartbeat field to given value.

### HasHeartbeat

`func (o *Service) HasHeartbeat() bool`

HasHeartbeat returns a boolean if a field has been set.

### GetHost

`func (o *Service) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Service) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Service) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *Service) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *Service) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Service) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Service) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPid

`func (o *Service) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *Service) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *Service) SetPid(v int32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *Service) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPrimary

`func (o *Service) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *Service) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *Service) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *Service) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetSslCertificate

`func (o *Service) GetSslCertificate() SSLCertificateNestview`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *Service) GetSslCertificateOk() (*SSLCertificateNestview, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *Service) SetSslCertificate(v SSLCertificateNestview)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *Service) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetStartedTime

`func (o *Service) GetStartedTime() time.Time`

GetStartedTime returns the StartedTime field if non-nil, zero value otherwise.

### GetStartedTimeOk

`func (o *Service) GetStartedTimeOk() (*time.Time, bool)`

GetStartedTimeOk returns a tuple with the StartedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedTime

`func (o *Service) SetStartedTime(v time.Time)`

SetStartedTime sets StartedTime field to given value.

### HasStartedTime

`func (o *Service) HasStartedTime() bool`

HasStartedTime returns a boolean if a field has been set.

### GetStatus

`func (o *Service) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Service) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Service) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Service) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Service) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Service) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Service) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Service) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUp

`func (o *Service) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *Service) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *Service) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *Service) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUpdate

`func (o *Service) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *Service) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *Service) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *Service) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


