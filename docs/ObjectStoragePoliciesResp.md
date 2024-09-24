# ObjectStoragePoliciesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsPolicies** | [**[]ObjectStoragePolicy**](ObjectStoragePolicy.md) | object storage policies | 

## Methods

### NewObjectStoragePoliciesResp

`func NewObjectStoragePoliciesResp(osPolicies []ObjectStoragePolicy, ) *ObjectStoragePoliciesResp`

NewObjectStoragePoliciesResp instantiates a new ObjectStoragePoliciesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStoragePoliciesRespWithDefaults

`func NewObjectStoragePoliciesRespWithDefaults() *ObjectStoragePoliciesResp`

NewObjectStoragePoliciesRespWithDefaults instantiates a new ObjectStoragePoliciesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsPolicies

`func (o *ObjectStoragePoliciesResp) GetOsPolicies() []ObjectStoragePolicy`

GetOsPolicies returns the OsPolicies field if non-nil, zero value otherwise.

### GetOsPoliciesOk

`func (o *ObjectStoragePoliciesResp) GetOsPoliciesOk() (*[]ObjectStoragePolicy, bool)`

GetOsPoliciesOk returns a tuple with the OsPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPolicies

`func (o *ObjectStoragePoliciesResp) SetOsPolicies(v []ObjectStoragePolicy)`

SetOsPolicies sets OsPolicies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


