# Host

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

## Methods

### NewHost

`func NewHost(adminIp string, ) *Host`

NewHost instantiates a new Host object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostWithDefaults

`func NewHostWithDefaults() *Host`

NewHostWithDefaults instantiates a new Host object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *Host) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *Host) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *Host) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *Host) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAdminIp

`func (o *Host) GetAdminIp() string`

GetAdminIp returns the AdminIp field if non-nil, zero value otherwise.

### GetAdminIpOk

`func (o *Host) GetAdminIpOk() (*string, bool)`

GetAdminIpOk returns a tuple with the AdminIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminIp

`func (o *Host) SetAdminIp(v string)`

SetAdminIp sets AdminIp field to given value.


### GetAdminVip

`func (o *Host) GetAdminVip() AdminVIPNestview`

GetAdminVip returns the AdminVip field if non-nil, zero value otherwise.

### GetAdminVipOk

`func (o *Host) GetAdminVipOk() (*AdminVIPNestview, bool)`

GetAdminVipOk returns a tuple with the AdminVip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminVip

`func (o *Host) SetAdminVip(v AdminVIPNestview)`

SetAdminVip sets AdminVip field to given value.

### HasAdminVip

`func (o *Host) HasAdminVip() bool`

HasAdminVip returns a boolean if a field has been set.

### GetClockDiff

`func (o *Host) GetClockDiff() int64`

GetClockDiff returns the ClockDiff field if non-nil, zero value otherwise.

### GetClockDiffOk

`func (o *Host) GetClockDiffOk() (*int64, bool)`

GetClockDiffOk returns a tuple with the ClockDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockDiff

`func (o *Host) SetClockDiff(v int64)`

SetClockDiff sets ClockDiff field to given value.

### HasClockDiff

`func (o *Host) HasClockDiff() bool`

HasClockDiff returns a boolean if a field has been set.

### GetCluster

`func (o *Host) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Host) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Host) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Host) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCores

`func (o *Host) GetCores() int64`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *Host) GetCoresOk() (*int64, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *Host) SetCores(v int64)`

SetCores sets Cores field to given value.

### HasCores

`func (o *Host) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetCpuArch

`func (o *Host) GetCpuArch() string`

GetCpuArch returns the CpuArch field if non-nil, zero value otherwise.

### GetCpuArchOk

`func (o *Host) GetCpuArchOk() (*string, bool)`

GetCpuArchOk returns a tuple with the CpuArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArch

`func (o *Host) SetCpuArch(v string)`

SetCpuArch sets CpuArch field to given value.

### HasCpuArch

`func (o *Host) HasCpuArch() bool`

HasCpuArch returns a boolean if a field has been set.

### GetCpuModel

`func (o *Host) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *Host) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *Host) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *Host) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetCreate

`func (o *Host) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *Host) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *Host) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *Host) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *Host) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Host) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Host) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Host) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiskNum

`func (o *Host) GetDiskNum() int64`

GetDiskNum returns the DiskNum field if non-nil, zero value otherwise.

### GetDiskNumOk

`func (o *Host) GetDiskNumOk() (*int64, bool)`

GetDiskNumOk returns a tuple with the DiskNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskNum

`func (o *Host) SetDiskNum(v int64)`

SetDiskNum sets DiskNum field to given value.

### HasDiskNum

`func (o *Host) HasDiskNum() bool`

HasDiskNum returns a boolean if a field has been set.

### GetEnclosures

`func (o *Host) GetEnclosures() []map[string]interface{}`

GetEnclosures returns the Enclosures field if non-nil, zero value otherwise.

### GetEnclosuresOk

`func (o *Host) GetEnclosuresOk() (*[]map[string]interface{}, bool)`

GetEnclosuresOk returns a tuple with the Enclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosures

`func (o *Host) SetEnclosures(v []map[string]interface{})`

SetEnclosures sets Enclosures field to given value.

### HasEnclosures

`func (o *Host) HasEnclosures() bool`

HasEnclosures returns a boolean if a field has been set.

### GetFcports

`func (o *Host) GetFcports() []FCPort`

GetFcports returns the Fcports field if non-nil, zero value otherwise.

### GetFcportsOk

`func (o *Host) GetFcportsOk() (*[]FCPort, bool)`

GetFcportsOk returns a tuple with the Fcports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcports

`func (o *Host) SetFcports(v []FCPort)`

SetFcports sets Fcports field to given value.

### HasFcports

`func (o *Host) HasFcports() bool`

HasFcports returns a boolean if a field has been set.

### GetGatewayIps

`func (o *Host) GetGatewayIps() string`

GetGatewayIps returns the GatewayIps field if non-nil, zero value otherwise.

### GetGatewayIpsOk

`func (o *Host) GetGatewayIpsOk() (*string, bool)`

GetGatewayIpsOk returns a tuple with the GatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIps

`func (o *Host) SetGatewayIps(v string)`

SetGatewayIps sets GatewayIps field to given value.

### HasGatewayIps

`func (o *Host) HasGatewayIps() bool`

HasGatewayIps returns a boolean if a field has been set.

### GetId

`func (o *Host) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Host) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Host) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Host) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsMasterDb

`func (o *Host) GetIsMasterDb() bool`

GetIsMasterDb returns the IsMasterDb field if non-nil, zero value otherwise.

### GetIsMasterDbOk

`func (o *Host) GetIsMasterDbOk() (*bool, bool)`

GetIsMasterDbOk returns a tuple with the IsMasterDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMasterDb

`func (o *Host) SetIsMasterDb(v bool)`

SetIsMasterDb sets IsMasterDb field to given value.

### HasIsMasterDb

`func (o *Host) HasIsMasterDb() bool`

HasIsMasterDb returns a boolean if a field has been set.

### GetKvmStatus

`func (o *Host) GetKvmStatus() int64`

GetKvmStatus returns the KvmStatus field if non-nil, zero value otherwise.

### GetKvmStatusOk

`func (o *Host) GetKvmStatusOk() (*int64, bool)`

GetKvmStatusOk returns a tuple with the KvmStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmStatus

`func (o *Host) SetKvmStatus(v int64)`

SetKvmStatus sets KvmStatus field to given value.

### HasKvmStatus

`func (o *Host) HasKvmStatus() bool`

HasKvmStatus returns a boolean if a field has been set.

### GetMemoryKbyte

`func (o *Host) GetMemoryKbyte() int64`

GetMemoryKbyte returns the MemoryKbyte field if non-nil, zero value otherwise.

### GetMemoryKbyteOk

`func (o *Host) GetMemoryKbyteOk() (*int64, bool)`

GetMemoryKbyteOk returns a tuple with the MemoryKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryKbyte

`func (o *Host) SetMemoryKbyte(v int64)`

SetMemoryKbyte sets MemoryKbyte field to given value.

### HasMemoryKbyte

`func (o *Host) HasMemoryKbyte() bool`

HasMemoryKbyte returns a boolean if a field has been set.

### GetModel

`func (o *Host) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Host) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Host) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Host) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *Host) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Host) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Host) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Host) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumaNodes

`func (o *Host) GetNumaNodes() string`

GetNumaNodes returns the NumaNodes field if non-nil, zero value otherwise.

### GetNumaNodesOk

`func (o *Host) GetNumaNodesOk() (*string, bool)`

GetNumaNodesOk returns a tuple with the NumaNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaNodes

`func (o *Host) SetNumaNodes(v string)`

SetNumaNodes sets NumaNodes field to given value.

### HasNumaNodes

`func (o *Host) HasNumaNodes() bool`

HasNumaNodes returns a boolean if a field has been set.

### GetOs

`func (o *Host) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Host) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Host) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *Host) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOspCluster

`func (o *Host) GetOspCluster() Cluster`

GetOspCluster returns the OspCluster field if non-nil, zero value otherwise.

### GetOspClusterOk

`func (o *Host) GetOspClusterOk() (*Cluster, bool)`

GetOspClusterOk returns a tuple with the OspCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspCluster

`func (o *Host) SetOspCluster(v Cluster)`

SetOspCluster sets OspCluster field to given value.

### HasOspCluster

`func (o *Host) HasOspCluster() bool`

HasOspCluster returns a boolean if a field has been set.

### GetOspClusterIp

`func (o *Host) GetOspClusterIp() string`

GetOspClusterIp returns the OspClusterIp field if non-nil, zero value otherwise.

### GetOspClusterIpOk

`func (o *Host) GetOspClusterIpOk() (*string, bool)`

GetOspClusterIpOk returns a tuple with the OspClusterIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspClusterIp

`func (o *Host) SetOspClusterIp(v string)`

SetOspClusterIp sets OspClusterIp field to given value.

### HasOspClusterIp

`func (o *Host) HasOspClusterIp() bool`

HasOspClusterIp returns a boolean if a field has been set.

### GetOspGatewayIps

`func (o *Host) GetOspGatewayIps() string`

GetOspGatewayIps returns the OspGatewayIps field if non-nil, zero value otherwise.

### GetOspGatewayIpsOk

`func (o *Host) GetOspGatewayIpsOk() (*string, bool)`

GetOspGatewayIpsOk returns a tuple with the OspGatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspGatewayIps

`func (o *Host) SetOspGatewayIps(v string)`

SetOspGatewayIps sets OspGatewayIps field to given value.

### HasOspGatewayIps

`func (o *Host) HasOspGatewayIps() bool`

HasOspGatewayIps returns a boolean if a field has been set.

### GetOspPlacementNode

`func (o *Host) GetOspPlacementNode() PlacementNode`

GetOspPlacementNode returns the OspPlacementNode field if non-nil, zero value otherwise.

### GetOspPlacementNodeOk

`func (o *Host) GetOspPlacementNodeOk() (*PlacementNode, bool)`

GetOspPlacementNodeOk returns a tuple with the OspPlacementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspPlacementNode

`func (o *Host) SetOspPlacementNode(v PlacementNode)`

SetOspPlacementNode sets OspPlacementNode field to given value.

### HasOspPlacementNode

`func (o *Host) HasOspPlacementNode() bool`

HasOspPlacementNode returns a boolean if a field has been set.

### GetOspZoneId

`func (o *Host) GetOspZoneId() int64`

GetOspZoneId returns the OspZoneId field if non-nil, zero value otherwise.

### GetOspZoneIdOk

`func (o *Host) GetOspZoneIdOk() (*int64, bool)`

GetOspZoneIdOk returns a tuple with the OspZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspZoneId

`func (o *Host) SetOspZoneId(v int64)`

SetOspZoneId sets OspZoneId field to given value.

### HasOspZoneId

`func (o *Host) HasOspZoneId() bool`

HasOspZoneId returns a boolean if a field has been set.

### GetPlacementNode

`func (o *Host) GetPlacementNode() PlacementNode`

GetPlacementNode returns the PlacementNode field if non-nil, zero value otherwise.

### GetPlacementNodeOk

`func (o *Host) GetPlacementNodeOk() (*PlacementNode, bool)`

GetPlacementNodeOk returns a tuple with the PlacementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementNode

`func (o *Host) SetPlacementNode(v PlacementNode)`

SetPlacementNode sets PlacementNode field to given value.

### HasPlacementNode

`func (o *Host) HasPlacementNode() bool`

HasPlacementNode returns a boolean if a field has been set.

### GetPrivateIp

`func (o *Host) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *Host) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *Host) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *Host) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetProtectionDomain

`func (o *Host) GetProtectionDomain() ProtectionDomainNestview`

GetProtectionDomain returns the ProtectionDomain field if non-nil, zero value otherwise.

### GetProtectionDomainOk

`func (o *Host) GetProtectionDomainOk() (*ProtectionDomainNestview, bool)`

GetProtectionDomainOk returns a tuple with the ProtectionDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionDomain

`func (o *Host) SetProtectionDomain(v ProtectionDomainNestview)`

SetProtectionDomain sets ProtectionDomain field to given value.

### HasProtectionDomain

`func (o *Host) HasProtectionDomain() bool`

HasProtectionDomain returns a boolean if a field has been set.

### GetPublicIps

`func (o *Host) GetPublicIps() string`

GetPublicIps returns the PublicIps field if non-nil, zero value otherwise.

### GetPublicIpsOk

`func (o *Host) GetPublicIpsOk() (*string, bool)`

GetPublicIpsOk returns a tuple with the PublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIps

`func (o *Host) SetPublicIps(v string)`

SetPublicIps sets PublicIps field to given value.

### HasPublicIps

`func (o *Host) HasPublicIps() bool`

HasPublicIps returns a boolean if a field has been set.

### GetRack

`func (o *Host) GetRack() string`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *Host) GetRackOk() (*string, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *Host) SetRack(v string)`

SetRack sets Rack field to given value.

### HasRack

`func (o *Host) HasRack() bool`

HasRack returns a boolean if a field has been set.

### GetRoles

`func (o *Host) GetRoles() string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Host) GetRolesOk() (*string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Host) SetRoles(v string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Host) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRootDisk

`func (o *Host) GetRootDisk() DiskNestview`

GetRootDisk returns the RootDisk field if non-nil, zero value otherwise.

### GetRootDiskOk

`func (o *Host) GetRootDiskOk() (*DiskNestview, bool)`

GetRootDiskOk returns a tuple with the RootDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDisk

`func (o *Host) SetRootDisk(v DiskNestview)`

SetRootDisk sets RootDisk field to given value.

### HasRootDisk

`func (o *Host) HasRootDisk() bool`

HasRootDisk returns a boolean if a field has been set.

### GetSerial

`func (o *Host) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Host) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Host) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *Host) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetStatus

`func (o *Host) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Host) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Host) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Host) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Host) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Host) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Host) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Host) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUp

`func (o *Host) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *Host) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *Host) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *Host) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUpdate

`func (o *Host) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *Host) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *Host) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *Host) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUptime

`func (o *Host) GetUptime() time.Time`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *Host) GetUptimeOk() (*time.Time, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *Host) SetUptime(v time.Time)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *Host) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVendor

`func (o *Host) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Host) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Host) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *Host) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


