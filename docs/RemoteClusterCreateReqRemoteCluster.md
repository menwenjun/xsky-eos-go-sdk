# RemoteClusterCreateReqRemoteCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | access token of remote cluster | 
**RemoteClusterId** | Pointer to **int64** | id of remote cluster | [optional] 
**Url** | **string** | URL of remote cluster | 

## Methods

### NewRemoteClusterCreateReqRemoteCluster

`func NewRemoteClusterCreateReqRemoteCluster(accessToken string, url string, ) *RemoteClusterCreateReqRemoteCluster`

NewRemoteClusterCreateReqRemoteCluster instantiates a new RemoteClusterCreateReqRemoteCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteClusterCreateReqRemoteClusterWithDefaults

`func NewRemoteClusterCreateReqRemoteClusterWithDefaults() *RemoteClusterCreateReqRemoteCluster`

NewRemoteClusterCreateReqRemoteClusterWithDefaults instantiates a new RemoteClusterCreateReqRemoteCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *RemoteClusterCreateReqRemoteCluster) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *RemoteClusterCreateReqRemoteCluster) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *RemoteClusterCreateReqRemoteCluster) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetRemoteClusterId

`func (o *RemoteClusterCreateReqRemoteCluster) GetRemoteClusterId() int64`

GetRemoteClusterId returns the RemoteClusterId field if non-nil, zero value otherwise.

### GetRemoteClusterIdOk

`func (o *RemoteClusterCreateReqRemoteCluster) GetRemoteClusterIdOk() (*int64, bool)`

GetRemoteClusterIdOk returns a tuple with the RemoteClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClusterId

`func (o *RemoteClusterCreateReqRemoteCluster) SetRemoteClusterId(v int64)`

SetRemoteClusterId sets RemoteClusterId field to given value.

### HasRemoteClusterId

`func (o *RemoteClusterCreateReqRemoteCluster) HasRemoteClusterId() bool`

HasRemoteClusterId returns a boolean if a field has been set.

### GetUrl

`func (o *RemoteClusterCreateReqRemoteCluster) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RemoteClusterCreateReqRemoteCluster) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RemoteClusterCreateReqRemoteCluster) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


