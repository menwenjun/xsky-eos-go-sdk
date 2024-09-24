# HostRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**AdminIp** | **string** |  | 
**AdminVip** | Pointer to [**AdminVIPNestview**](AdminVIPNestview.md) |  | [optional] 
**ClockDiff** | Pointer to **int64** | clock diff in milliseconds with primary host | [optional] 
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**Cores** | Pointer to **int64** |  | [optional] 
**CpuArch** | Pointer to **string** |  | [optional] 
**CpuModel** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DiskNum** | Pointer to **int64** |  | [optional] 
**Enclosures** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Fcports** | Pointer to [**[]FCPort**](FCPort.md) | fc ports of host | [optional] 
**GatewayIps** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IsMasterDb** | Pointer to **bool** |  | [optional] 
**KvmStatus** | Pointer to **int64** | HostKVMStatus represent KVM status | [optional] 
**MemoryKbyte** | Pointer to **int64** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumaNodes** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**OspCluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**OspClusterIp** | Pointer to **string** |  | [optional] 
**OspGatewayIps** | Pointer to **string** |  | [optional] 
**OspPlacementNode** | Pointer to [**PlacementNode**](PlacementNode.md) |  | [optional] 
**OspZoneId** | Pointer to **int64** |  | [optional] 
**PlacementNode** | Pointer to [**PlacementNode**](PlacementNode.md) |  | [optional] 
**PrivateIp** | Pointer to **string** |  | [optional] 
**ProtectionDomain** | Pointer to [**ProtectionDomainNestview**](ProtectionDomainNestview.md) |  | [optional] 
**PublicIps** | Pointer to **string** |  | [optional] 
**Rack** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **string** |  | [optional] 
**RootDisk** | Pointer to [**DiskNestview**](DiskNestview.md) |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Up** | Pointer to **bool** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Uptime** | Pointer to **time.Time** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Samples** | Pointer to [**[]HostStat**](HostStat.md) |  | [optional] 

## Methods

### NewHostRecord

`func NewHostRecord(adminIp string, ) *HostRecord`

NewHostRecord instantiates a new HostRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostRecordWithDefaults

`func NewHostRecordWithDefaults() *HostRecord`

NewHostRecordWithDefaults instantiates a new HostRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *HostRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *HostRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *HostRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *HostRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAdminIp

`func (o *HostRecord) GetAdminIp() string`

GetAdminIp returns the AdminIp field if non-nil, zero value otherwise.

### GetAdminIpOk

`func (o *HostRecord) GetAdminIpOk() (*string, bool)`

GetAdminIpOk returns a tuple with the AdminIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminIp

`func (o *HostRecord) SetAdminIp(v string)`

SetAdminIp sets AdminIp field to given value.


### GetAdminVip

`func (o *HostRecord) GetAdminVip() AdminVIPNestview`

GetAdminVip returns the AdminVip field if non-nil, zero value otherwise.

### GetAdminVipOk

`func (o *HostRecord) GetAdminVipOk() (*AdminVIPNestview, bool)`

GetAdminVipOk returns a tuple with the AdminVip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminVip

`func (o *HostRecord) SetAdminVip(v AdminVIPNestview)`

SetAdminVip sets AdminVip field to given value.

### HasAdminVip

`func (o *HostRecord) HasAdminVip() bool`

HasAdminVip returns a boolean if a field has been set.

### GetClockDiff

`func (o *HostRecord) GetClockDiff() int64`

GetClockDiff returns the ClockDiff field if non-nil, zero value otherwise.

### GetClockDiffOk

`func (o *HostRecord) GetClockDiffOk() (*int64, bool)`

GetClockDiffOk returns a tuple with the ClockDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockDiff

`func (o *HostRecord) SetClockDiff(v int64)`

SetClockDiff sets ClockDiff field to given value.

### HasClockDiff

`func (o *HostRecord) HasClockDiff() bool`

HasClockDiff returns a boolean if a field has been set.

### GetCluster

`func (o *HostRecord) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HostRecord) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HostRecord) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HostRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCores

`func (o *HostRecord) GetCores() int64`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *HostRecord) GetCoresOk() (*int64, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *HostRecord) SetCores(v int64)`

SetCores sets Cores field to given value.

### HasCores

`func (o *HostRecord) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetCpuArch

`func (o *HostRecord) GetCpuArch() string`

GetCpuArch returns the CpuArch field if non-nil, zero value otherwise.

### GetCpuArchOk

`func (o *HostRecord) GetCpuArchOk() (*string, bool)`

GetCpuArchOk returns a tuple with the CpuArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArch

`func (o *HostRecord) SetCpuArch(v string)`

SetCpuArch sets CpuArch field to given value.

### HasCpuArch

`func (o *HostRecord) HasCpuArch() bool`

HasCpuArch returns a boolean if a field has been set.

### GetCpuModel

`func (o *HostRecord) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *HostRecord) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *HostRecord) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *HostRecord) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetCreate

`func (o *HostRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *HostRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *HostRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *HostRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *HostRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HostRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HostRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HostRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiskNum

`func (o *HostRecord) GetDiskNum() int64`

GetDiskNum returns the DiskNum field if non-nil, zero value otherwise.

### GetDiskNumOk

`func (o *HostRecord) GetDiskNumOk() (*int64, bool)`

GetDiskNumOk returns a tuple with the DiskNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskNum

`func (o *HostRecord) SetDiskNum(v int64)`

SetDiskNum sets DiskNum field to given value.

### HasDiskNum

`func (o *HostRecord) HasDiskNum() bool`

HasDiskNum returns a boolean if a field has been set.

### GetEnclosures

`func (o *HostRecord) GetEnclosures() []map[string]interface{}`

GetEnclosures returns the Enclosures field if non-nil, zero value otherwise.

### GetEnclosuresOk

`func (o *HostRecord) GetEnclosuresOk() (*[]map[string]interface{}, bool)`

GetEnclosuresOk returns a tuple with the Enclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosures

`func (o *HostRecord) SetEnclosures(v []map[string]interface{})`

SetEnclosures sets Enclosures field to given value.

### HasEnclosures

`func (o *HostRecord) HasEnclosures() bool`

HasEnclosures returns a boolean if a field has been set.

### GetFcports

`func (o *HostRecord) GetFcports() []FCPort`

GetFcports returns the Fcports field if non-nil, zero value otherwise.

### GetFcportsOk

`func (o *HostRecord) GetFcportsOk() (*[]FCPort, bool)`

GetFcportsOk returns a tuple with the Fcports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcports

`func (o *HostRecord) SetFcports(v []FCPort)`

SetFcports sets Fcports field to given value.

### HasFcports

`func (o *HostRecord) HasFcports() bool`

HasFcports returns a boolean if a field has been set.

### GetGatewayIps

`func (o *HostRecord) GetGatewayIps() string`

GetGatewayIps returns the GatewayIps field if non-nil, zero value otherwise.

### GetGatewayIpsOk

`func (o *HostRecord) GetGatewayIpsOk() (*string, bool)`

GetGatewayIpsOk returns a tuple with the GatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIps

`func (o *HostRecord) SetGatewayIps(v string)`

SetGatewayIps sets GatewayIps field to given value.

### HasGatewayIps

`func (o *HostRecord) HasGatewayIps() bool`

HasGatewayIps returns a boolean if a field has been set.

### GetId

`func (o *HostRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *HostRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsMasterDb

`func (o *HostRecord) GetIsMasterDb() bool`

GetIsMasterDb returns the IsMasterDb field if non-nil, zero value otherwise.

### GetIsMasterDbOk

`func (o *HostRecord) GetIsMasterDbOk() (*bool, bool)`

GetIsMasterDbOk returns a tuple with the IsMasterDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMasterDb

`func (o *HostRecord) SetIsMasterDb(v bool)`

SetIsMasterDb sets IsMasterDb field to given value.

### HasIsMasterDb

`func (o *HostRecord) HasIsMasterDb() bool`

HasIsMasterDb returns a boolean if a field has been set.

### GetKvmStatus

`func (o *HostRecord) GetKvmStatus() int64`

GetKvmStatus returns the KvmStatus field if non-nil, zero value otherwise.

### GetKvmStatusOk

`func (o *HostRecord) GetKvmStatusOk() (*int64, bool)`

GetKvmStatusOk returns a tuple with the KvmStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmStatus

`func (o *HostRecord) SetKvmStatus(v int64)`

SetKvmStatus sets KvmStatus field to given value.

### HasKvmStatus

`func (o *HostRecord) HasKvmStatus() bool`

HasKvmStatus returns a boolean if a field has been set.

### GetMemoryKbyte

`func (o *HostRecord) GetMemoryKbyte() int64`

GetMemoryKbyte returns the MemoryKbyte field if non-nil, zero value otherwise.

### GetMemoryKbyteOk

`func (o *HostRecord) GetMemoryKbyteOk() (*int64, bool)`

GetMemoryKbyteOk returns a tuple with the MemoryKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryKbyte

`func (o *HostRecord) SetMemoryKbyte(v int64)`

SetMemoryKbyte sets MemoryKbyte field to given value.

### HasMemoryKbyte

`func (o *HostRecord) HasMemoryKbyte() bool`

HasMemoryKbyte returns a boolean if a field has been set.

### GetModel

`func (o *HostRecord) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *HostRecord) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *HostRecord) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *HostRecord) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *HostRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HostRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumaNodes

`func (o *HostRecord) GetNumaNodes() string`

GetNumaNodes returns the NumaNodes field if non-nil, zero value otherwise.

### GetNumaNodesOk

`func (o *HostRecord) GetNumaNodesOk() (*string, bool)`

GetNumaNodesOk returns a tuple with the NumaNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaNodes

`func (o *HostRecord) SetNumaNodes(v string)`

SetNumaNodes sets NumaNodes field to given value.

### HasNumaNodes

`func (o *HostRecord) HasNumaNodes() bool`

HasNumaNodes returns a boolean if a field has been set.

### GetOs

`func (o *HostRecord) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *HostRecord) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *HostRecord) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *HostRecord) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOspCluster

`func (o *HostRecord) GetOspCluster() Cluster`

GetOspCluster returns the OspCluster field if non-nil, zero value otherwise.

### GetOspClusterOk

`func (o *HostRecord) GetOspClusterOk() (*Cluster, bool)`

GetOspClusterOk returns a tuple with the OspCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspCluster

`func (o *HostRecord) SetOspCluster(v Cluster)`

SetOspCluster sets OspCluster field to given value.

### HasOspCluster

`func (o *HostRecord) HasOspCluster() bool`

HasOspCluster returns a boolean if a field has been set.

### GetOspClusterIp

`func (o *HostRecord) GetOspClusterIp() string`

GetOspClusterIp returns the OspClusterIp field if non-nil, zero value otherwise.

### GetOspClusterIpOk

`func (o *HostRecord) GetOspClusterIpOk() (*string, bool)`

GetOspClusterIpOk returns a tuple with the OspClusterIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspClusterIp

`func (o *HostRecord) SetOspClusterIp(v string)`

SetOspClusterIp sets OspClusterIp field to given value.

### HasOspClusterIp

`func (o *HostRecord) HasOspClusterIp() bool`

HasOspClusterIp returns a boolean if a field has been set.

### GetOspGatewayIps

`func (o *HostRecord) GetOspGatewayIps() string`

GetOspGatewayIps returns the OspGatewayIps field if non-nil, zero value otherwise.

### GetOspGatewayIpsOk

`func (o *HostRecord) GetOspGatewayIpsOk() (*string, bool)`

GetOspGatewayIpsOk returns a tuple with the OspGatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspGatewayIps

`func (o *HostRecord) SetOspGatewayIps(v string)`

SetOspGatewayIps sets OspGatewayIps field to given value.

### HasOspGatewayIps

`func (o *HostRecord) HasOspGatewayIps() bool`

HasOspGatewayIps returns a boolean if a field has been set.

### GetOspPlacementNode

`func (o *HostRecord) GetOspPlacementNode() PlacementNode`

GetOspPlacementNode returns the OspPlacementNode field if non-nil, zero value otherwise.

### GetOspPlacementNodeOk

`func (o *HostRecord) GetOspPlacementNodeOk() (*PlacementNode, bool)`

GetOspPlacementNodeOk returns a tuple with the OspPlacementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspPlacementNode

`func (o *HostRecord) SetOspPlacementNode(v PlacementNode)`

SetOspPlacementNode sets OspPlacementNode field to given value.

### HasOspPlacementNode

`func (o *HostRecord) HasOspPlacementNode() bool`

HasOspPlacementNode returns a boolean if a field has been set.

### GetOspZoneId

`func (o *HostRecord) GetOspZoneId() int64`

GetOspZoneId returns the OspZoneId field if non-nil, zero value otherwise.

### GetOspZoneIdOk

`func (o *HostRecord) GetOspZoneIdOk() (*int64, bool)`

GetOspZoneIdOk returns a tuple with the OspZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspZoneId

`func (o *HostRecord) SetOspZoneId(v int64)`

SetOspZoneId sets OspZoneId field to given value.

### HasOspZoneId

`func (o *HostRecord) HasOspZoneId() bool`

HasOspZoneId returns a boolean if a field has been set.

### GetPlacementNode

`func (o *HostRecord) GetPlacementNode() PlacementNode`

GetPlacementNode returns the PlacementNode field if non-nil, zero value otherwise.

### GetPlacementNodeOk

`func (o *HostRecord) GetPlacementNodeOk() (*PlacementNode, bool)`

GetPlacementNodeOk returns a tuple with the PlacementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementNode

`func (o *HostRecord) SetPlacementNode(v PlacementNode)`

SetPlacementNode sets PlacementNode field to given value.

### HasPlacementNode

`func (o *HostRecord) HasPlacementNode() bool`

HasPlacementNode returns a boolean if a field has been set.

### GetPrivateIp

`func (o *HostRecord) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *HostRecord) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *HostRecord) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *HostRecord) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetProtectionDomain

`func (o *HostRecord) GetProtectionDomain() ProtectionDomainNestview`

GetProtectionDomain returns the ProtectionDomain field if non-nil, zero value otherwise.

### GetProtectionDomainOk

`func (o *HostRecord) GetProtectionDomainOk() (*ProtectionDomainNestview, bool)`

GetProtectionDomainOk returns a tuple with the ProtectionDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionDomain

`func (o *HostRecord) SetProtectionDomain(v ProtectionDomainNestview)`

SetProtectionDomain sets ProtectionDomain field to given value.

### HasProtectionDomain

`func (o *HostRecord) HasProtectionDomain() bool`

HasProtectionDomain returns a boolean if a field has been set.

### GetPublicIps

`func (o *HostRecord) GetPublicIps() string`

GetPublicIps returns the PublicIps field if non-nil, zero value otherwise.

### GetPublicIpsOk

`func (o *HostRecord) GetPublicIpsOk() (*string, bool)`

GetPublicIpsOk returns a tuple with the PublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIps

`func (o *HostRecord) SetPublicIps(v string)`

SetPublicIps sets PublicIps field to given value.

### HasPublicIps

`func (o *HostRecord) HasPublicIps() bool`

HasPublicIps returns a boolean if a field has been set.

### GetRack

`func (o *HostRecord) GetRack() string`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *HostRecord) GetRackOk() (*string, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *HostRecord) SetRack(v string)`

SetRack sets Rack field to given value.

### HasRack

`func (o *HostRecord) HasRack() bool`

HasRack returns a boolean if a field has been set.

### GetRoles

`func (o *HostRecord) GetRoles() string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *HostRecord) GetRolesOk() (*string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *HostRecord) SetRoles(v string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *HostRecord) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRootDisk

`func (o *HostRecord) GetRootDisk() DiskNestview`

GetRootDisk returns the RootDisk field if non-nil, zero value otherwise.

### GetRootDiskOk

`func (o *HostRecord) GetRootDiskOk() (*DiskNestview, bool)`

GetRootDiskOk returns a tuple with the RootDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDisk

`func (o *HostRecord) SetRootDisk(v DiskNestview)`

SetRootDisk sets RootDisk field to given value.

### HasRootDisk

`func (o *HostRecord) HasRootDisk() bool`

HasRootDisk returns a boolean if a field has been set.

### GetSerial

`func (o *HostRecord) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *HostRecord) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *HostRecord) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *HostRecord) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetStatus

`func (o *HostRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HostRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *HostRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HostRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HostRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HostRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUp

`func (o *HostRecord) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *HostRecord) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *HostRecord) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *HostRecord) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUpdate

`func (o *HostRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *HostRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *HostRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *HostRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUptime

`func (o *HostRecord) GetUptime() time.Time`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *HostRecord) GetUptimeOk() (*time.Time, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *HostRecord) SetUptime(v time.Time)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *HostRecord) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVendor

`func (o *HostRecord) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HostRecord) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HostRecord) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *HostRecord) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetSamples

`func (o *HostRecord) GetSamples() []HostStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *HostRecord) GetSamplesOk() (*[]HostStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *HostRecord) SetSamples(v []HostStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *HostRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


