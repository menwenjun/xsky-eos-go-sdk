# OspMetadataServer

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

## Methods

### NewOspMetadataServer

`func NewOspMetadataServer() *OspMetadataServer`

NewOspMetadataServer instantiates a new OspMetadataServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspMetadataServerWithDefaults

`func NewOspMetadataServerWithDefaults() *OspMetadataServer`

NewOspMetadataServerWithDefaults instantiates a new OspMetadataServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *OspMetadataServer) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *OspMetadataServer) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *OspMetadataServer) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *OspMetadataServer) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetCluster

`func (o *OspMetadataServer) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OspMetadataServer) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OspMetadataServer) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OspMetadataServer) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OspMetadataServer) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OspMetadataServer) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OspMetadataServer) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OspMetadataServer) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDatacenter

`func (o *OspMetadataServer) GetDatacenter() PlacementNodeNestview`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *OspMetadataServer) GetDatacenterOk() (*PlacementNodeNestview, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *OspMetadataServer) SetDatacenter(v PlacementNodeNestview)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *OspMetadataServer) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetHost

`func (o *OspMetadataServer) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *OspMetadataServer) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *OspMetadataServer) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *OspMetadataServer) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *OspMetadataServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OspMetadataServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OspMetadataServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OspMetadataServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *OspMetadataServer) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OspMetadataServer) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OspMetadataServer) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OspMetadataServer) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMetadataCluster

`func (o *OspMetadataServer) GetMetadataCluster() OspMetadataClusterNestview`

GetMetadataCluster returns the MetadataCluster field if non-nil, zero value otherwise.

### GetMetadataClusterOk

`func (o *OspMetadataServer) GetMetadataClusterOk() (*OspMetadataClusterNestview, bool)`

GetMetadataClusterOk returns a tuple with the MetadataCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCluster

`func (o *OspMetadataServer) SetMetadataCluster(v OspMetadataClusterNestview)`

SetMetadataCluster sets MetadataCluster field to given value.

### HasMetadataCluster

`func (o *OspMetadataServer) HasMetadataCluster() bool`

HasMetadataCluster returns a boolean if a field has been set.

### GetOsd

`func (o *OspMetadataServer) GetOsd() OsdNestview`

GetOsd returns the Osd field if non-nil, zero value otherwise.

### GetOsdOk

`func (o *OspMetadataServer) GetOsdOk() (*OsdNestview, bool)`

GetOsdOk returns a tuple with the Osd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsd

`func (o *OspMetadataServer) SetOsd(v OsdNestview)`

SetOsd sets Osd field to given value.

### HasOsd

`func (o *OspMetadataServer) HasOsd() bool`

HasOsd returns a boolean if a field has been set.

### GetPartition

`func (o *OspMetadataServer) GetPartition() PartitionNestview`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *OspMetadataServer) GetPartitionOk() (*PartitionNestview, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *OspMetadataServer) SetPartition(v PartitionNestview)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *OspMetadataServer) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPort

`func (o *OspMetadataServer) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OspMetadataServer) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OspMetadataServer) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *OspMetadataServer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRoles

`func (o *OspMetadataServer) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OspMetadataServer) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OspMetadataServer) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *OspMetadataServer) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetStatus

`func (o *OspMetadataServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OspMetadataServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OspMetadataServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OspMetadataServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *OspMetadataServer) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OspMetadataServer) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OspMetadataServer) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OspMetadataServer) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVersion

`func (o *OspMetadataServer) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OspMetadataServer) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OspMetadataServer) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OspMetadataServer) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


