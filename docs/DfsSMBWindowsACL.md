# DfsSMBWindowsACL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aces** | Pointer to [**[]DfsSMBWindowsACE**](DfsSMBWindowsACE.md) |  | [optional] 
**Control** | Pointer to **string** |  | [optional] 
**DfsRootfs** | Pointer to [**DfsRootfs**](DfsRootfs.md) |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**InheritanceEnabled** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 

## Methods

### NewDfsSMBWindowsACL

`func NewDfsSMBWindowsACL() *DfsSMBWindowsACL`

NewDfsSMBWindowsACL instantiates a new DfsSMBWindowsACL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBWindowsACLWithDefaults

`func NewDfsSMBWindowsACLWithDefaults() *DfsSMBWindowsACL`

NewDfsSMBWindowsACLWithDefaults instantiates a new DfsSMBWindowsACL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAces

`func (o *DfsSMBWindowsACL) GetAces() []DfsSMBWindowsACE`

GetAces returns the Aces field if non-nil, zero value otherwise.

### GetAcesOk

`func (o *DfsSMBWindowsACL) GetAcesOk() (*[]DfsSMBWindowsACE, bool)`

GetAcesOk returns a tuple with the Aces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAces

`func (o *DfsSMBWindowsACL) SetAces(v []DfsSMBWindowsACE)`

SetAces sets Aces field to given value.

### HasAces

`func (o *DfsSMBWindowsACL) HasAces() bool`

HasAces returns a boolean if a field has been set.

### GetControl

`func (o *DfsSMBWindowsACL) GetControl() string`

GetControl returns the Control field if non-nil, zero value otherwise.

### GetControlOk

`func (o *DfsSMBWindowsACL) GetControlOk() (*string, bool)`

GetControlOk returns a tuple with the Control field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControl

`func (o *DfsSMBWindowsACL) SetControl(v string)`

SetControl sets Control field to given value.

### HasControl

`func (o *DfsSMBWindowsACL) HasControl() bool`

HasControl returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *DfsSMBWindowsACL) GetDfsRootfs() DfsRootfs`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsSMBWindowsACL) GetDfsRootfsOk() (*DfsRootfs, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsSMBWindowsACL) SetDfsRootfs(v DfsRootfs)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *DfsSMBWindowsACL) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetGroup

`func (o *DfsSMBWindowsACL) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DfsSMBWindowsACL) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DfsSMBWindowsACL) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DfsSMBWindowsACL) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetInheritanceEnabled

`func (o *DfsSMBWindowsACL) GetInheritanceEnabled() bool`

GetInheritanceEnabled returns the InheritanceEnabled field if non-nil, zero value otherwise.

### GetInheritanceEnabledOk

`func (o *DfsSMBWindowsACL) GetInheritanceEnabledOk() (*bool, bool)`

GetInheritanceEnabledOk returns a tuple with the InheritanceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceEnabled

`func (o *DfsSMBWindowsACL) SetInheritanceEnabled(v bool)`

SetInheritanceEnabled sets InheritanceEnabled field to given value.

### HasInheritanceEnabled

`func (o *DfsSMBWindowsACL) HasInheritanceEnabled() bool`

HasInheritanceEnabled returns a boolean if a field has been set.

### GetOwner

`func (o *DfsSMBWindowsACL) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DfsSMBWindowsACL) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DfsSMBWindowsACL) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DfsSMBWindowsACL) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPath

`func (o *DfsSMBWindowsACL) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsSMBWindowsACL) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsSMBWindowsACL) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DfsSMBWindowsACL) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRevision

`func (o *DfsSMBWindowsACL) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DfsSMBWindowsACL) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DfsSMBWindowsACL) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *DfsSMBWindowsACL) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


