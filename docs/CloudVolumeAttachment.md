# CloudVolumeAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockVolume** | Pointer to [**VolumeNestview**](VolumeNestview.md) |  | [optional] 
**CloudInstance** | Pointer to [**CloudInstanceNestview**](CloudInstanceNestview.md) |  | [optional] 
**CloudVolume** | Pointer to [**CloudVolumeNestview**](CloudVolumeNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Device** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCloudVolumeAttachment

`func NewCloudVolumeAttachment() *CloudVolumeAttachment`

NewCloudVolumeAttachment instantiates a new CloudVolumeAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVolumeAttachmentWithDefaults

`func NewCloudVolumeAttachmentWithDefaults() *CloudVolumeAttachment`

NewCloudVolumeAttachmentWithDefaults instantiates a new CloudVolumeAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockVolume

`func (o *CloudVolumeAttachment) GetBlockVolume() VolumeNestview`

GetBlockVolume returns the BlockVolume field if non-nil, zero value otherwise.

### GetBlockVolumeOk

`func (o *CloudVolumeAttachment) GetBlockVolumeOk() (*VolumeNestview, bool)`

GetBlockVolumeOk returns a tuple with the BlockVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolume

`func (o *CloudVolumeAttachment) SetBlockVolume(v VolumeNestview)`

SetBlockVolume sets BlockVolume field to given value.

### HasBlockVolume

`func (o *CloudVolumeAttachment) HasBlockVolume() bool`

HasBlockVolume returns a boolean if a field has been set.

### GetCloudInstance

`func (o *CloudVolumeAttachment) GetCloudInstance() CloudInstanceNestview`

GetCloudInstance returns the CloudInstance field if non-nil, zero value otherwise.

### GetCloudInstanceOk

`func (o *CloudVolumeAttachment) GetCloudInstanceOk() (*CloudInstanceNestview, bool)`

GetCloudInstanceOk returns a tuple with the CloudInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInstance

`func (o *CloudVolumeAttachment) SetCloudInstance(v CloudInstanceNestview)`

SetCloudInstance sets CloudInstance field to given value.

### HasCloudInstance

`func (o *CloudVolumeAttachment) HasCloudInstance() bool`

HasCloudInstance returns a boolean if a field has been set.

### GetCloudVolume

`func (o *CloudVolumeAttachment) GetCloudVolume() CloudVolumeNestview`

GetCloudVolume returns the CloudVolume field if non-nil, zero value otherwise.

### GetCloudVolumeOk

`func (o *CloudVolumeAttachment) GetCloudVolumeOk() (*CloudVolumeNestview, bool)`

GetCloudVolumeOk returns a tuple with the CloudVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudVolume

`func (o *CloudVolumeAttachment) SetCloudVolume(v CloudVolumeNestview)`

SetCloudVolume sets CloudVolume field to given value.

### HasCloudVolume

`func (o *CloudVolumeAttachment) HasCloudVolume() bool`

HasCloudVolume returns a boolean if a field has been set.

### GetCreate

`func (o *CloudVolumeAttachment) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *CloudVolumeAttachment) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *CloudVolumeAttachment) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *CloudVolumeAttachment) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDevice

`func (o *CloudVolumeAttachment) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *CloudVolumeAttachment) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *CloudVolumeAttachment) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *CloudVolumeAttachment) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetId

`func (o *CloudVolumeAttachment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudVolumeAttachment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudVolumeAttachment) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CloudVolumeAttachment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdate

`func (o *CloudVolumeAttachment) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *CloudVolumeAttachment) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *CloudVolumeAttachment) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *CloudVolumeAttachment) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


