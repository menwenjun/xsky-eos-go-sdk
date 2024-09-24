# FSKerberosCreateReqInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description of kerberos | [optional] 
**Ip** | Pointer to **string** | dns ip for domain name | [optional] 
**Kdc** | **string** | domain name or ip address of kdc server | 
**Name** | **string** | name of kerberos | 
**Password** | **string** | password of kerberos | 
**Realm** | **string** | realm of kerberos | 
**Type** | Pointer to **string** | type of kerberos | [optional] 
**Username** | **string** | username of kerberos | 

## Methods

### NewFSKerberosCreateReqInfo

`func NewFSKerberosCreateReqInfo(kdc string, name string, password string, realm string, username string, ) *FSKerberosCreateReqInfo`

NewFSKerberosCreateReqInfo instantiates a new FSKerberosCreateReqInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFSKerberosCreateReqInfoWithDefaults

`func NewFSKerberosCreateReqInfoWithDefaults() *FSKerberosCreateReqInfo`

NewFSKerberosCreateReqInfoWithDefaults instantiates a new FSKerberosCreateReqInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FSKerberosCreateReqInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FSKerberosCreateReqInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FSKerberosCreateReqInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FSKerberosCreateReqInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIp

`func (o *FSKerberosCreateReqInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *FSKerberosCreateReqInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *FSKerberosCreateReqInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *FSKerberosCreateReqInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetKdc

`func (o *FSKerberosCreateReqInfo) GetKdc() string`

GetKdc returns the Kdc field if non-nil, zero value otherwise.

### GetKdcOk

`func (o *FSKerberosCreateReqInfo) GetKdcOk() (*string, bool)`

GetKdcOk returns a tuple with the Kdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdc

`func (o *FSKerberosCreateReqInfo) SetKdc(v string)`

SetKdc sets Kdc field to given value.


### GetName

`func (o *FSKerberosCreateReqInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FSKerberosCreateReqInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FSKerberosCreateReqInfo) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *FSKerberosCreateReqInfo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FSKerberosCreateReqInfo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FSKerberosCreateReqInfo) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRealm

`func (o *FSKerberosCreateReqInfo) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *FSKerberosCreateReqInfo) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *FSKerberosCreateReqInfo) SetRealm(v string)`

SetRealm sets Realm field to given value.


### GetType

`func (o *FSKerberosCreateReqInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FSKerberosCreateReqInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FSKerberosCreateReqInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FSKerberosCreateReqInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *FSKerberosCreateReqInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *FSKerberosCreateReqInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *FSKerberosCreateReqInfo) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


