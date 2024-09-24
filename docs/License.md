# License

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**ClusterType** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Customer** | Pointer to **string** |  | [optional] 
**ExpiredTime** | Pointer to **time.Time** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**FsId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Limits** | Pointer to **map[string]string** | ProductLimits defines product limits | [optional] 
**MinorVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NameCn** | Pointer to **string** |  | [optional] 
**ProductInfo** | Pointer to [**ProductInfo**](ProductInfo.md) |  | [optional] 
**Protocols** | Pointer to **[]string** |  | [optional] 
**Registered** | Pointer to **bool** |  | [optional] 
**Services** | Pointer to [**[]ProductService**](ProductService.md) |  | [optional] 
**SignedTime** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Type is Result type | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** | load from license info | [optional] 

## Methods

### NewLicense

`func NewLicense() *License`

NewLicense instantiates a new License object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseWithDefaults

`func NewLicenseWithDefaults() *License`

NewLicenseWithDefaults instantiates a new License object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *License) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *License) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *License) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *License) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCluster

`func (o *License) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *License) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *License) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *License) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterType

`func (o *License) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *License) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *License) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *License) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### GetCreate

`func (o *License) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *License) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *License) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *License) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCustomer

`func (o *License) GetCustomer() string`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *License) GetCustomerOk() (*string, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *License) SetCustomer(v string)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *License) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetExpiredTime

`func (o *License) GetExpiredTime() time.Time`

GetExpiredTime returns the ExpiredTime field if non-nil, zero value otherwise.

### GetExpiredTimeOk

`func (o *License) GetExpiredTimeOk() (*time.Time, bool)`

GetExpiredTimeOk returns a tuple with the ExpiredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredTime

`func (o *License) SetExpiredTime(v time.Time)`

SetExpiredTime sets ExpiredTime field to given value.

### HasExpiredTime

`func (o *License) HasExpiredTime() bool`

HasExpiredTime returns a boolean if a field has been set.

### GetFeatures

`func (o *License) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *License) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *License) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *License) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFsId

`func (o *License) GetFsId() string`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *License) GetFsIdOk() (*string, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *License) SetFsId(v string)`

SetFsId sets FsId field to given value.

### HasFsId

`func (o *License) HasFsId() bool`

HasFsId returns a boolean if a field has been set.

### GetId

`func (o *License) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *License) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *License) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *License) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLimits

`func (o *License) GetLimits() map[string]string`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *License) GetLimitsOk() (*map[string]string, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *License) SetLimits(v map[string]string)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *License) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetMinorVersion

`func (o *License) GetMinorVersion() string`

GetMinorVersion returns the MinorVersion field if non-nil, zero value otherwise.

### GetMinorVersionOk

`func (o *License) GetMinorVersionOk() (*string, bool)`

GetMinorVersionOk returns a tuple with the MinorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorVersion

`func (o *License) SetMinorVersion(v string)`

SetMinorVersion sets MinorVersion field to given value.

### HasMinorVersion

`func (o *License) HasMinorVersion() bool`

HasMinorVersion returns a boolean if a field has been set.

### GetName

`func (o *License) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *License) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *License) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *License) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameCn

`func (o *License) GetNameCn() string`

GetNameCn returns the NameCn field if non-nil, zero value otherwise.

### GetNameCnOk

`func (o *License) GetNameCnOk() (*string, bool)`

GetNameCnOk returns a tuple with the NameCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameCn

`func (o *License) SetNameCn(v string)`

SetNameCn sets NameCn field to given value.

### HasNameCn

`func (o *License) HasNameCn() bool`

HasNameCn returns a boolean if a field has been set.

### GetProductInfo

`func (o *License) GetProductInfo() ProductInfo`

GetProductInfo returns the ProductInfo field if non-nil, zero value otherwise.

### GetProductInfoOk

`func (o *License) GetProductInfoOk() (*ProductInfo, bool)`

GetProductInfoOk returns a tuple with the ProductInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductInfo

`func (o *License) SetProductInfo(v ProductInfo)`

SetProductInfo sets ProductInfo field to given value.

### HasProductInfo

`func (o *License) HasProductInfo() bool`

HasProductInfo returns a boolean if a field has been set.

### GetProtocols

`func (o *License) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *License) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *License) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *License) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetRegistered

`func (o *License) GetRegistered() bool`

GetRegistered returns the Registered field if non-nil, zero value otherwise.

### GetRegisteredOk

`func (o *License) GetRegisteredOk() (*bool, bool)`

GetRegisteredOk returns a tuple with the Registered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistered

`func (o *License) SetRegistered(v bool)`

SetRegistered sets Registered field to given value.

### HasRegistered

`func (o *License) HasRegistered() bool`

HasRegistered returns a boolean if a field has been set.

### GetServices

`func (o *License) GetServices() []ProductService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *License) GetServicesOk() (*[]ProductService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *License) SetServices(v []ProductService)`

SetServices sets Services field to given value.

### HasServices

`func (o *License) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetSignedTime

`func (o *License) GetSignedTime() time.Time`

GetSignedTime returns the SignedTime field if non-nil, zero value otherwise.

### GetSignedTimeOk

`func (o *License) GetSignedTimeOk() (*time.Time, bool)`

GetSignedTimeOk returns a tuple with the SignedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTime

`func (o *License) SetSignedTime(v time.Time)`

SetSignedTime sets SignedTime field to given value.

### HasSignedTime

`func (o *License) HasSignedTime() bool`

HasSignedTime returns a boolean if a field has been set.

### GetStatus

`func (o *License) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *License) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *License) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *License) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *License) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *License) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *License) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *License) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *License) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *License) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *License) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *License) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *License) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *License) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *License) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *License) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


