# OspMetadataServerRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Datacenter** | Pointer to [**PlacementNodeNestview**](PlacementNodeNestview.md) |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**MetadataCluster** | Pointer to [**OspMetadataClusterNestview**](OspMetadataClusterNestview.md) |  | [optional] 
**Osd** | Pointer to [**OsdNestview**](OsdNestview.md) |  | [optional] 
**Partition** | Pointer to [**PartitionNestview**](PartitionNestview.md) |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**Samples** | Pointer to [**[]OspMetadataServerStat**](OspMetadataServerStat.md) |  | [optional] 

## Methods

### NewOspMetadataServerRecord

`func NewOspMetadataServerRecord() *OspMetadataServerRecord`

NewOspMetadataServerRecord instantiates a new OspMetadataServerRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspMetadataServerRecordWithDefaults

`func NewOspMetadataServerRecordWithDefaults() *OspMetadataServerRecord`

NewOspMetadataServerRecordWithDefaults instantiates a new OspMetadataServerRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *OspMetadataServerRecord) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *OspMetadataServerRecord) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *OspMetadataServerRecord) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *OspMetadataServerRecord) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetCluster

`func (o *OspMetadataServerRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OspMetadataServerRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OspMetadataServerRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OspMetadataServerRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OspMetadataServerRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OspMetadataServerRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OspMetadataServerRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OspMetadataServerRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDatacenter

`func (o *OspMetadataServerRecord) GetDatacenter() PlacementNodeNestview`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *OspMetadataServerRecord) GetDatacenterOk() (*PlacementNodeNestview, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *OspMetadataServerRecord) SetDatacenter(v PlacementNodeNestview)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *OspMetadataServerRecord) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetHost

`func (o *OspMetadataServerRecord) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *OspMetadataServerRecord) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *OspMetadataServerRecord) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *OspMetadataServerRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *OspMetadataServerRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OspMetadataServerRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OspMetadataServerRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OspMetadataServerRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *OspMetadataServerRecord) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OspMetadataServerRecord) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OspMetadataServerRecord) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OspMetadataServerRecord) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMetadataCluster

`func (o *OspMetadataServerRecord) GetMetadataCluster() OspMetadataClusterNestview`

GetMetadataCluster returns the MetadataCluster field if non-nil, zero value otherwise.

### GetMetadataClusterOk

`func (o *OspMetadataServerRecord) GetMetadataClusterOk() (*OspMetadataClusterNestview, bool)`

GetMetadataClusterOk returns a tuple with the MetadataCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCluster

`func (o *OspMetadataServerRecord) SetMetadataCluster(v OspMetadataClusterNestview)`

SetMetadataCluster sets MetadataCluster field to given value.

### HasMetadataCluster

`func (o *OspMetadataServerRecord) HasMetadataCluster() bool`

HasMetadataCluster returns a boolean if a field has been set.

### GetOsd

`func (o *OspMetadataServerRecord) GetOsd() OsdNestview`

GetOsd returns the Osd field if non-nil, zero value otherwise.

### GetOsdOk

`func (o *OspMetadataServerRecord) GetOsdOk() (*OsdNestview, bool)`

GetOsdOk returns a tuple with the Osd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsd

`func (o *OspMetadataServerRecord) SetOsd(v OsdNestview)`

SetOsd sets Osd field to given value.

### HasOsd

`func (o *OspMetadataServerRecord) HasOsd() bool`

HasOsd returns a boolean if a field has been set.

### GetPartition

`func (o *OspMetadataServerRecord) GetPartition() PartitionNestview`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *OspMetadataServerRecord) GetPartitionOk() (*PartitionNestview, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *OspMetadataServerRecord) SetPartition(v PartitionNestview)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *OspMetadataServerRecord) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPort

`func (o *OspMetadataServerRecord) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OspMetadataServerRecord) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OspMetadataServerRecord) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *OspMetadataServerRecord) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRoles

`func (o *OspMetadataServerRecord) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OspMetadataServerRecord) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OspMetadataServerRecord) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *OspMetadataServerRecord) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetStatus

`func (o *OspMetadataServerRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OspMetadataServerRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OspMetadataServerRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OspMetadataServerRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *OspMetadataServerRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OspMetadataServerRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OspMetadataServerRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OspMetadataServerRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVersion

`func (o *OspMetadataServerRecord) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OspMetadataServerRecord) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OspMetadataServerRecord) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OspMetadataServerRecord) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSamples

`func (o *OspMetadataServerRecord) GetSamples() []OspMetadataServerStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *OspMetadataServerRecord) GetSamplesOk() (*[]OspMetadataServerStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *OspMetadataServerRecord) SetSamples(v []OspMetadataServerStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *OspMetadataServerRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


