# ClusterInstallation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Finished** | Pointer to **bool** |  | [optional] 
**Steps** | Pointer to [**ClusterInstallationSteps**](ClusterInstallationSteps.md) |  | [optional] 

## Methods

### NewClusterInstallation

`func NewClusterInstallation() *ClusterInstallation`

NewClusterInstallation instantiates a new ClusterInstallation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInstallationWithDefaults

`func NewClusterInstallationWithDefaults() *ClusterInstallation`

NewClusterInstallationWithDefaults instantiates a new ClusterInstallation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinished

`func (o *ClusterInstallation) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *ClusterInstallation) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *ClusterInstallation) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *ClusterInstallation) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetSteps

`func (o *ClusterInstallation) GetSteps() ClusterInstallationSteps`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ClusterInstallation) GetStepsOk() (*ClusterInstallationSteps, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ClusterInstallation) SetSteps(v ClusterInstallationSteps)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ClusterInstallation) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


