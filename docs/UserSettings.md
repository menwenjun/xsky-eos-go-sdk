# UserSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**UserSettingsUser**](UserSettingsUser.md) |  | 

## Methods

### NewUserSettings

`func NewUserSettings(user UserSettingsUser, ) *UserSettings`

NewUserSettings instantiates a new UserSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingsWithDefaults

`func NewUserSettingsWithDefaults() *UserSettings`

NewUserSettingsWithDefaults instantiates a new UserSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserSettings) GetUser() UserSettingsUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserSettings) GetUserOk() (*UserSettingsUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserSettings) SetUser(v UserSettingsUser)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


