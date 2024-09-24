# OSOriginPullRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketId** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**EscapeToSlash** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ModeType** | Pointer to **string** |  | [optional] 
**OriginConf** | Pointer to [**OriginConf**](OriginConf.md) |  | [optional] 
**OriginInfo** | Pointer to [**OriginInfo**](OriginInfo.md) |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**S3LoadBalancerGroup** | Pointer to [**S3LoadBalancerGroup**](S3LoadBalancerGroup.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Suffix** | Pointer to **string** |  | [optional] 

## Methods

### NewOSOriginPullRule

`func NewOSOriginPullRule() *OSOriginPullRule`

NewOSOriginPullRule instantiates a new OSOriginPullRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSOriginPullRuleWithDefaults

`func NewOSOriginPullRuleWithDefaults() *OSOriginPullRule`

NewOSOriginPullRuleWithDefaults instantiates a new OSOriginPullRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketId

`func (o *OSOriginPullRule) GetBucketId() int64`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *OSOriginPullRule) GetBucketIdOk() (*int64, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *OSOriginPullRule) SetBucketId(v int64)`

SetBucketId sets BucketId field to given value.

### HasBucketId

`func (o *OSOriginPullRule) HasBucketId() bool`

HasBucketId returns a boolean if a field has been set.

### GetCluster

`func (o *OSOriginPullRule) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSOriginPullRule) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSOriginPullRule) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSOriginPullRule) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConnected

`func (o *OSOriginPullRule) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *OSOriginPullRule) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *OSOriginPullRule) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *OSOriginPullRule) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetCreate

`func (o *OSOriginPullRule) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSOriginPullRule) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSOriginPullRule) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSOriginPullRule) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetEscapeToSlash

`func (o *OSOriginPullRule) GetEscapeToSlash() bool`

GetEscapeToSlash returns the EscapeToSlash field if non-nil, zero value otherwise.

### GetEscapeToSlashOk

`func (o *OSOriginPullRule) GetEscapeToSlashOk() (*bool, bool)`

GetEscapeToSlashOk returns a tuple with the EscapeToSlash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscapeToSlash

`func (o *OSOriginPullRule) SetEscapeToSlash(v bool)`

SetEscapeToSlash sets EscapeToSlash field to given value.

### HasEscapeToSlash

`func (o *OSOriginPullRule) HasEscapeToSlash() bool`

HasEscapeToSlash returns a boolean if a field has been set.

### GetId

`func (o *OSOriginPullRule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSOriginPullRule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSOriginPullRule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OSOriginPullRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModeType

`func (o *OSOriginPullRule) GetModeType() string`

GetModeType returns the ModeType field if non-nil, zero value otherwise.

### GetModeTypeOk

`func (o *OSOriginPullRule) GetModeTypeOk() (*string, bool)`

GetModeTypeOk returns a tuple with the ModeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeType

`func (o *OSOriginPullRule) SetModeType(v string)`

SetModeType sets ModeType field to given value.

### HasModeType

`func (o *OSOriginPullRule) HasModeType() bool`

HasModeType returns a boolean if a field has been set.

### GetOriginConf

`func (o *OSOriginPullRule) GetOriginConf() OriginConf`

GetOriginConf returns the OriginConf field if non-nil, zero value otherwise.

### GetOriginConfOk

`func (o *OSOriginPullRule) GetOriginConfOk() (*OriginConf, bool)`

GetOriginConfOk returns a tuple with the OriginConf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginConf

`func (o *OSOriginPullRule) SetOriginConf(v OriginConf)`

SetOriginConf sets OriginConf field to given value.

### HasOriginConf

`func (o *OSOriginPullRule) HasOriginConf() bool`

HasOriginConf returns a boolean if a field has been set.

### GetOriginInfo

`func (o *OSOriginPullRule) GetOriginInfo() OriginInfo`

GetOriginInfo returns the OriginInfo field if non-nil, zero value otherwise.

### GetOriginInfoOk

`func (o *OSOriginPullRule) GetOriginInfoOk() (*OriginInfo, bool)`

GetOriginInfoOk returns a tuple with the OriginInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginInfo

`func (o *OSOriginPullRule) SetOriginInfo(v OriginInfo)`

SetOriginInfo sets OriginInfo field to given value.

### HasOriginInfo

`func (o *OSOriginPullRule) HasOriginInfo() bool`

HasOriginInfo returns a boolean if a field has been set.

### GetPrefix

`func (o *OSOriginPullRule) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *OSOriginPullRule) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *OSOriginPullRule) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *OSOriginPullRule) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetS3LoadBalancerGroup

`func (o *OSOriginPullRule) GetS3LoadBalancerGroup() S3LoadBalancerGroup`

GetS3LoadBalancerGroup returns the S3LoadBalancerGroup field if non-nil, zero value otherwise.

### GetS3LoadBalancerGroupOk

`func (o *OSOriginPullRule) GetS3LoadBalancerGroupOk() (*S3LoadBalancerGroup, bool)`

GetS3LoadBalancerGroupOk returns a tuple with the S3LoadBalancerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3LoadBalancerGroup

`func (o *OSOriginPullRule) SetS3LoadBalancerGroup(v S3LoadBalancerGroup)`

SetS3LoadBalancerGroup sets S3LoadBalancerGroup field to given value.

### HasS3LoadBalancerGroup

`func (o *OSOriginPullRule) HasS3LoadBalancerGroup() bool`

HasS3LoadBalancerGroup returns a boolean if a field has been set.

### GetStatus

`func (o *OSOriginPullRule) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSOriginPullRule) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSOriginPullRule) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OSOriginPullRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuffix

`func (o *OSOriginPullRule) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *OSOriginPullRule) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *OSOriginPullRule) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *OSOriginPullRule) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


