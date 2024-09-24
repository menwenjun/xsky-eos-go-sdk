# OSBuiltinBluRay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminEndpoint** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OsPolicies** | Pointer to [**[]ObjectStoragePolicy**](ObjectStoragePolicy.md) |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewOSBuiltinBluRay

`func NewOSBuiltinBluRay() *OSBuiltinBluRay`

NewOSBuiltinBluRay instantiates a new OSBuiltinBluRay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSBuiltinBluRayWithDefaults

`func NewOSBuiltinBluRayWithDefaults() *OSBuiltinBluRay`

NewOSBuiltinBluRayWithDefaults instantiates a new OSBuiltinBluRay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminEndpoint

`func (o *OSBuiltinBluRay) GetAdminEndpoint() string`

GetAdminEndpoint returns the AdminEndpoint field if non-nil, zero value otherwise.

### GetAdminEndpointOk

`func (o *OSBuiltinBluRay) GetAdminEndpointOk() (*string, bool)`

GetAdminEndpointOk returns a tuple with the AdminEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEndpoint

`func (o *OSBuiltinBluRay) SetAdminEndpoint(v string)`

SetAdminEndpoint sets AdminEndpoint field to given value.

### HasAdminEndpoint

`func (o *OSBuiltinBluRay) HasAdminEndpoint() bool`

HasAdminEndpoint returns a boolean if a field has been set.

### GetCluster

`func (o *OSBuiltinBluRay) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSBuiltinBluRay) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSBuiltinBluRay) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSBuiltinBluRay) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OSBuiltinBluRay) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSBuiltinBluRay) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSBuiltinBluRay) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSBuiltinBluRay) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *OSBuiltinBluRay) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSBuiltinBluRay) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSBuiltinBluRay) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OSBuiltinBluRay) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OSBuiltinBluRay) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OSBuiltinBluRay) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OSBuiltinBluRay) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OSBuiltinBluRay) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsPolicies

`func (o *OSBuiltinBluRay) GetOsPolicies() []ObjectStoragePolicy`

GetOsPolicies returns the OsPolicies field if non-nil, zero value otherwise.

### GetOsPoliciesOk

`func (o *OSBuiltinBluRay) GetOsPoliciesOk() (*[]ObjectStoragePolicy, bool)`

GetOsPoliciesOk returns a tuple with the OsPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPolicies

`func (o *OSBuiltinBluRay) SetOsPolicies(v []ObjectStoragePolicy)`

SetOsPolicies sets OsPolicies field to given value.

### HasOsPolicies

`func (o *OSBuiltinBluRay) HasOsPolicies() bool`

HasOsPolicies returns a boolean if a field has been set.

### GetPassword

`func (o *OSBuiltinBluRay) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OSBuiltinBluRay) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OSBuiltinBluRay) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OSBuiltinBluRay) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetStatus

`func (o *OSBuiltinBluRay) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSBuiltinBluRay) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSBuiltinBluRay) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OSBuiltinBluRay) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserName

`func (o *OSBuiltinBluRay) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *OSBuiltinBluRay) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *OSBuiltinBluRay) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *OSBuiltinBluRay) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


