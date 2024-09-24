# DNSZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DnsGatewayGroup** | Pointer to [**DNSGatewayGroupNestview**](DNSGatewayGroupNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDNSZone

`func NewDNSZone() *DNSZone`

NewDNSZone instantiates a new DNSZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSZoneWithDefaults

`func NewDNSZoneWithDefaults() *DNSZone`

NewDNSZoneWithDefaults instantiates a new DNSZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DNSZone) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DNSZone) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DNSZone) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DNSZone) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCreate

`func (o *DNSZone) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DNSZone) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DNSZone) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DNSZone) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDnsGatewayGroup

`func (o *DNSZone) GetDnsGatewayGroup() DNSGatewayGroupNestview`

GetDnsGatewayGroup returns the DnsGatewayGroup field if non-nil, zero value otherwise.

### GetDnsGatewayGroupOk

`func (o *DNSZone) GetDnsGatewayGroupOk() (*DNSGatewayGroupNestview, bool)`

GetDnsGatewayGroupOk returns a tuple with the DnsGatewayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsGatewayGroup

`func (o *DNSZone) SetDnsGatewayGroup(v DNSGatewayGroupNestview)`

SetDnsGatewayGroup sets DnsGatewayGroup field to given value.

### HasDnsGatewayGroup

`func (o *DNSZone) HasDnsGatewayGroup() bool`

HasDnsGatewayGroup returns a boolean if a field has been set.

### GetId

`func (o *DNSZone) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DNSZone) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DNSZone) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DNSZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrigin

`func (o *DNSZone) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *DNSZone) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *DNSZone) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *DNSZone) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetStatus

`func (o *DNSZone) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DNSZone) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DNSZone) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DNSZone) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTtl

`func (o *DNSZone) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DNSZone) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DNSZone) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DNSZone) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUpdate

`func (o *DNSZone) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DNSZone) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DNSZone) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DNSZone) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


