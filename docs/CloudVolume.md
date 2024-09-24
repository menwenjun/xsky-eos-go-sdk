# CloudVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockVolume** | Pointer to [**VolumeNestview**](VolumeNestview.md) |  | [optional] 
**CloudDatacenter** | Pointer to [**CloudDatacenterNestview**](CloudDatacenterNestview.md) |  | [optional] 
**CloudPlatform** | Pointer to [**CloudPlatformNestview**](CloudPlatformNestview.md) |  | [optional] 
**CloudVolumeId** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MultiAttach** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCloudVolume

`func NewCloudVolume() *CloudVolume`

NewCloudVolume instantiates a new CloudVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVolumeWithDefaults

`func NewCloudVolumeWithDefaults() *CloudVolume`

NewCloudVolumeWithDefaults instantiates a new CloudVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockVolume

`func (o *CloudVolume) GetBlockVolume() VolumeNestview`

GetBlockVolume returns the BlockVolume field if non-nil, zero value otherwise.

### GetBlockVolumeOk

`func (o *CloudVolume) GetBlockVolumeOk() (*VolumeNestview, bool)`

GetBlockVolumeOk returns a tuple with the BlockVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolume

`func (o *CloudVolume) SetBlockVolume(v VolumeNestview)`

SetBlockVolume sets BlockVolume field to given value.

### HasBlockVolume

`func (o *CloudVolume) HasBlockVolume() bool`

HasBlockVolume returns a boolean if a field has been set.

### GetCloudDatacenter

`func (o *CloudVolume) GetCloudDatacenter() CloudDatacenterNestview`

GetCloudDatacenter returns the CloudDatacenter field if non-nil, zero value otherwise.

### GetCloudDatacenterOk

`func (o *CloudVolume) GetCloudDatacenterOk() (*CloudDatacenterNestview, bool)`

GetCloudDatacenterOk returns a tuple with the CloudDatacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDatacenter

`func (o *CloudVolume) SetCloudDatacenter(v CloudDatacenterNestview)`

SetCloudDatacenter sets CloudDatacenter field to given value.

### HasCloudDatacenter

`func (o *CloudVolume) HasCloudDatacenter() bool`

HasCloudDatacenter returns a boolean if a field has been set.

### GetCloudPlatform

`func (o *CloudVolume) GetCloudPlatform() CloudPlatformNestview`

GetCloudPlatform returns the CloudPlatform field if non-nil, zero value otherwise.

### GetCloudPlatformOk

`func (o *CloudVolume) GetCloudPlatformOk() (*CloudPlatformNestview, bool)`

GetCloudPlatformOk returns a tuple with the CloudPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPlatform

`func (o *CloudVolume) SetCloudPlatform(v CloudPlatformNestview)`

SetCloudPlatform sets CloudPlatform field to given value.

### HasCloudPlatform

`func (o *CloudVolume) HasCloudPlatform() bool`

HasCloudPlatform returns a boolean if a field has been set.

### GetCloudVolumeId

`func (o *CloudVolume) GetCloudVolumeId() string`

GetCloudVolumeId returns the CloudVolumeId field if non-nil, zero value otherwise.

### GetCloudVolumeIdOk

`func (o *CloudVolume) GetCloudVolumeIdOk() (*string, bool)`

GetCloudVolumeIdOk returns a tuple with the CloudVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudVolumeId

`func (o *CloudVolume) SetCloudVolumeId(v string)`

SetCloudVolumeId sets CloudVolumeId field to given value.

### HasCloudVolumeId

`func (o *CloudVolume) HasCloudVolumeId() bool`

HasCloudVolumeId returns a boolean if a field has been set.

### GetCreate

`func (o *CloudVolume) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *CloudVolume) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *CloudVolume) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *CloudVolume) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *CloudVolume) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudVolume) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudVolume) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CloudVolume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMultiAttach

`func (o *CloudVolume) GetMultiAttach() bool`

GetMultiAttach returns the MultiAttach field if non-nil, zero value otherwise.

### GetMultiAttachOk

`func (o *CloudVolume) GetMultiAttachOk() (*bool, bool)`

GetMultiAttachOk returns a tuple with the MultiAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAttach

`func (o *CloudVolume) SetMultiAttach(v bool)`

SetMultiAttach sets MultiAttach field to given value.

### HasMultiAttach

`func (o *CloudVolume) HasMultiAttach() bool`

HasMultiAttach returns a boolean if a field has been set.

### GetName

`func (o *CloudVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CloudVolume) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudVolume) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudVolume) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudVolume) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *CloudVolume) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *CloudVolume) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *CloudVolume) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *CloudVolume) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


