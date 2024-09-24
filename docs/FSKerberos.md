# FSKerberos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DfsHdfsNum** | Pointer to **int64** |  | [optional] 
**DfsHdfses** | Pointer to [**[]DfsHdfs**](DfsHdfs.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Kdc** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Realm** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**VerifyTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFSKerberos

`func NewFSKerberos() *FSKerberos`

NewFSKerberos instantiates a new FSKerberos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFSKerberosWithDefaults

`func NewFSKerberosWithDefaults() *FSKerberos`

NewFSKerberosWithDefaults instantiates a new FSKerberos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *FSKerberos) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *FSKerberos) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *FSKerberos) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *FSKerberos) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *FSKerberos) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *FSKerberos) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *FSKerberos) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *FSKerberos) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *FSKerberos) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *FSKerberos) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *FSKerberos) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *FSKerberos) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *FSKerberos) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FSKerberos) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FSKerberos) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FSKerberos) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsHdfsNum

`func (o *FSKerberos) GetDfsHdfsNum() int64`

GetDfsHdfsNum returns the DfsHdfsNum field if non-nil, zero value otherwise.

### GetDfsHdfsNumOk

`func (o *FSKerberos) GetDfsHdfsNumOk() (*int64, bool)`

GetDfsHdfsNumOk returns a tuple with the DfsHdfsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfsNum

`func (o *FSKerberos) SetDfsHdfsNum(v int64)`

SetDfsHdfsNum sets DfsHdfsNum field to given value.

### HasDfsHdfsNum

`func (o *FSKerberos) HasDfsHdfsNum() bool`

HasDfsHdfsNum returns a boolean if a field has been set.

### GetDfsHdfses

`func (o *FSKerberos) GetDfsHdfses() []DfsHdfs`

GetDfsHdfses returns the DfsHdfses field if non-nil, zero value otherwise.

### GetDfsHdfsesOk

`func (o *FSKerberos) GetDfsHdfsesOk() (*[]DfsHdfs, bool)`

GetDfsHdfsesOk returns a tuple with the DfsHdfses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfses

`func (o *FSKerberos) SetDfsHdfses(v []DfsHdfs)`

SetDfsHdfses sets DfsHdfses field to given value.

### HasDfsHdfses

`func (o *FSKerberos) HasDfsHdfses() bool`

HasDfsHdfses returns a boolean if a field has been set.

### GetId

`func (o *FSKerberos) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FSKerberos) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FSKerberos) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FSKerberos) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *FSKerberos) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *FSKerberos) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *FSKerberos) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *FSKerberos) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetKdc

`func (o *FSKerberos) GetKdc() string`

GetKdc returns the Kdc field if non-nil, zero value otherwise.

### GetKdcOk

`func (o *FSKerberos) GetKdcOk() (*string, bool)`

GetKdcOk returns a tuple with the Kdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdc

`func (o *FSKerberos) SetKdc(v string)`

SetKdc sets Kdc field to given value.

### HasKdc

`func (o *FSKerberos) HasKdc() bool`

HasKdc returns a boolean if a field has been set.

### GetName

`func (o *FSKerberos) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FSKerberos) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FSKerberos) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FSKerberos) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRealm

`func (o *FSKerberos) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *FSKerberos) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *FSKerberos) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *FSKerberos) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetStatus

`func (o *FSKerberos) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FSKerberos) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FSKerberos) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FSKerberos) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *FSKerberos) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FSKerberos) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FSKerberos) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FSKerberos) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *FSKerberos) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *FSKerberos) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *FSKerberos) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *FSKerberos) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUsername

`func (o *FSKerberos) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *FSKerberos) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *FSKerberos) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *FSKerberos) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVerifyTime

`func (o *FSKerberos) GetVerifyTime() time.Time`

GetVerifyTime returns the VerifyTime field if non-nil, zero value otherwise.

### GetVerifyTimeOk

`func (o *FSKerberos) GetVerifyTimeOk() (*time.Time, bool)`

GetVerifyTimeOk returns a tuple with the VerifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyTime

`func (o *FSKerberos) SetVerifyTime(v time.Time)`

SetVerifyTime sets VerifyTime field to given value.

### HasVerifyTime

`func (o *FSKerberos) HasVerifyTime() bool`

HasVerifyTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


