# DNSGatewayGroupCreateReqGroupVIPGroupsElt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** | network of vip group | 
**Vips** | Pointer to [**[]VIPArgs**](VIPArgs.md) | vip list | [optional] 

## Methods

### NewDNSGatewayGroupCreateReqGroupVIPGroupsElt

`func NewDNSGatewayGroupCreateReqGroupVIPGroupsElt(network string, ) *DNSGatewayGroupCreateReqGroupVIPGroupsElt`

NewDNSGatewayGroupCreateReqGroupVIPGroupsElt instantiates a new DNSGatewayGroupCreateReqGroupVIPGroupsElt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSGatewayGroupCreateReqGroupVIPGroupsEltWithDefaults

`func NewDNSGatewayGroupCreateReqGroupVIPGroupsEltWithDefaults() *DNSGatewayGroupCreateReqGroupVIPGroupsElt`

NewDNSGatewayGroupCreateReqGroupVIPGroupsEltWithDefaults instantiates a new DNSGatewayGroupCreateReqGroupVIPGroupsElt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *DNSGatewayGroupCreateReqGroupVIPGroupsElt) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *DNSGatewayGroupCreateReqGroupVIPGroupsElt) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *DNSGatewayGroupCreateReqGroupVIPGroupsElt) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetVips

`func (o *DNSGatewayGroupCreateReqGroupVIPGroupsElt) GetVips() []VIPArgs`

GetVips returns the Vips field if non-nil, zero value otherwise.

### GetVipsOk

`func (o *DNSGatewayGroupCreateReqGroupVIPGroupsElt) GetVipsOk() (*[]VIPArgs, bool)`

GetVipsOk returns a tuple with the Vips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVips

`func (o *DNSGatewayGroupCreateReqGroupVIPGroupsElt) SetVips(v []VIPArgs)`

SetVips sets Vips field to given value.

### HasVips

`func (o *DNSGatewayGroupCreateReqGroupVIPGroupsElt) HasVips() bool`

HasVips returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


