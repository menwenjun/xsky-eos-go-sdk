# EmailGroupsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailGroups** | [**[]EmailGroupRecord**](EmailGroupRecord.md) | email groups | 

## Methods

### NewEmailGroupsResp

`func NewEmailGroupsResp(emailGroups []EmailGroupRecord, ) *EmailGroupsResp`

NewEmailGroupsResp instantiates a new EmailGroupsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailGroupsRespWithDefaults

`func NewEmailGroupsRespWithDefaults() *EmailGroupsResp`

NewEmailGroupsRespWithDefaults instantiates a new EmailGroupsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailGroups

`func (o *EmailGroupsResp) GetEmailGroups() []EmailGroupRecord`

GetEmailGroups returns the EmailGroups field if non-nil, zero value otherwise.

### GetEmailGroupsOk

`func (o *EmailGroupsResp) GetEmailGroupsOk() (*[]EmailGroupRecord, bool)`

GetEmailGroupsOk returns a tuple with the EmailGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailGroups

`func (o *EmailGroupsResp) SetEmailGroups(v []EmailGroupRecord)`

SetEmailGroups sets EmailGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


