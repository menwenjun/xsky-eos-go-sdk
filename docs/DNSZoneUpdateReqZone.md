# DNSZoneUpdateReqZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origin** | Pointer to **string** | dns origin in zone | [optional] 
**Ttl** | Pointer to **int64** | dns ttl with zone | [optional] 

## Methods

### NewDNSZoneUpdateReqZone

`func NewDNSZoneUpdateReqZone() *DNSZoneUpdateReqZone`

NewDNSZoneUpdateReqZone instantiates a new DNSZoneUpdateReqZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSZoneUpdateReqZoneWithDefaults

`func NewDNSZoneUpdateReqZoneWithDefaults() *DNSZoneUpdateReqZone`

NewDNSZoneUpdateReqZoneWithDefaults instantiates a new DNSZoneUpdateReqZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigin

`func (o *DNSZoneUpdateReqZone) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *DNSZoneUpdateReqZone) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *DNSZoneUpdateReqZone) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *DNSZoneUpdateReqZone) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetTtl

`func (o *DNSZoneUpdateReqZone) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DNSZoneUpdateReqZone) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DNSZoneUpdateReqZone) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DNSZoneUpdateReqZone) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


