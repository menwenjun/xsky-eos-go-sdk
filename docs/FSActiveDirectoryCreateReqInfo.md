# FSActiveDirectoryCreateReqInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | ip of dns server | [optional] 
**Name** | **string** | name of active directory | 
**Password** | **string** | password of active directory | 
**Realm** | **string** | realm of active directory | 
**Username** | **string** | username of active directory | 
**Workgroup** | **string** | workgroup of active directory | 

## Methods

### NewFSActiveDirectoryCreateReqInfo

`func NewFSActiveDirectoryCreateReqInfo(name string, password string, realm string, username string, workgroup string, ) *FSActiveDirectoryCreateReqInfo`

NewFSActiveDirectoryCreateReqInfo instantiates a new FSActiveDirectoryCreateReqInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFSActiveDirectoryCreateReqInfoWithDefaults

`func NewFSActiveDirectoryCreateReqInfoWithDefaults() *FSActiveDirectoryCreateReqInfo`

NewFSActiveDirectoryCreateReqInfoWithDefaults instantiates a new FSActiveDirectoryCreateReqInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *FSActiveDirectoryCreateReqInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *FSActiveDirectoryCreateReqInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *FSActiveDirectoryCreateReqInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *FSActiveDirectoryCreateReqInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetName

`func (o *FSActiveDirectoryCreateReqInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FSActiveDirectoryCreateReqInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FSActiveDirectoryCreateReqInfo) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *FSActiveDirectoryCreateReqInfo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FSActiveDirectoryCreateReqInfo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FSActiveDirectoryCreateReqInfo) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRealm

`func (o *FSActiveDirectoryCreateReqInfo) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *FSActiveDirectoryCreateReqInfo) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *FSActiveDirectoryCreateReqInfo) SetRealm(v string)`

SetRealm sets Realm field to given value.


### GetUsername

`func (o *FSActiveDirectoryCreateReqInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *FSActiveDirectoryCreateReqInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *FSActiveDirectoryCreateReqInfo) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetWorkgroup

`func (o *FSActiveDirectoryCreateReqInfo) GetWorkgroup() string`

GetWorkgroup returns the Workgroup field if non-nil, zero value otherwise.

### GetWorkgroupOk

`func (o *FSActiveDirectoryCreateReqInfo) GetWorkgroupOk() (*string, bool)`

GetWorkgroupOk returns a tuple with the Workgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkgroup

`func (o *FSActiveDirectoryCreateReqInfo) SetWorkgroup(v string)`

SetWorkgroup sets Workgroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


