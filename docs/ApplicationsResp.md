# ApplicationsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]Application**](Application.md) |  | [optional] 

## Methods

### NewApplicationsResp

`func NewApplicationsResp() *ApplicationsResp`

NewApplicationsResp instantiates a new ApplicationsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationsRespWithDefaults

`func NewApplicationsRespWithDefaults() *ApplicationsResp`

NewApplicationsRespWithDefaults instantiates a new ApplicationsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *ApplicationsResp) GetApplications() []Application`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ApplicationsResp) GetApplicationsOk() (*[]Application, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ApplicationsResp) SetApplications(v []Application)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *ApplicationsResp) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


