# ClusterInstallationSteps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bootnode** | Pointer to **bool** | step of setting boot node | [optional] 
**License** | Pointer to **bool** | step of registering license | [optional] 
**User** | Pointer to **bool** | step of creating user | [optional] 

## Methods

### NewClusterInstallationSteps

`func NewClusterInstallationSteps() *ClusterInstallationSteps`

NewClusterInstallationSteps instantiates a new ClusterInstallationSteps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInstallationStepsWithDefaults

`func NewClusterInstallationStepsWithDefaults() *ClusterInstallationSteps`

NewClusterInstallationStepsWithDefaults instantiates a new ClusterInstallationSteps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootnode

`func (o *ClusterInstallationSteps) GetBootnode() bool`

GetBootnode returns the Bootnode field if non-nil, zero value otherwise.

### GetBootnodeOk

`func (o *ClusterInstallationSteps) GetBootnodeOk() (*bool, bool)`

GetBootnodeOk returns a tuple with the Bootnode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootnode

`func (o *ClusterInstallationSteps) SetBootnode(v bool)`

SetBootnode sets Bootnode field to given value.

### HasBootnode

`func (o *ClusterInstallationSteps) HasBootnode() bool`

HasBootnode returns a boolean if a field has been set.

### GetLicense

`func (o *ClusterInstallationSteps) GetLicense() bool`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ClusterInstallationSteps) GetLicenseOk() (*bool, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ClusterInstallationSteps) SetLicense(v bool)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ClusterInstallationSteps) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetUser

`func (o *ClusterInstallationSteps) GetUser() bool`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ClusterInstallationSteps) GetUserOk() (*bool, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ClusterInstallationSteps) SetUser(v bool)`

SetUser sets User field to given value.

### HasUser

`func (o *ClusterInstallationSteps) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


