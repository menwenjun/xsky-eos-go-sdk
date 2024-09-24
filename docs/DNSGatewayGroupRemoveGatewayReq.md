# DNSGatewayGroupRemoveGatewayReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsGateways** | Pointer to [**[]DNSGatewayGroupRemoveGatewayReqGatewaysElt**](DNSGatewayGroupRemoveGatewayReqGatewaysElt.md) |  | [optional] 

## Methods

### NewDNSGatewayGroupRemoveGatewayReq

`func NewDNSGatewayGroupRemoveGatewayReq() *DNSGatewayGroupRemoveGatewayReq`

NewDNSGatewayGroupRemoveGatewayReq instantiates a new DNSGatewayGroupRemoveGatewayReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSGatewayGroupRemoveGatewayReqWithDefaults

`func NewDNSGatewayGroupRemoveGatewayReqWithDefaults() *DNSGatewayGroupRemoveGatewayReq`

NewDNSGatewayGroupRemoveGatewayReqWithDefaults instantiates a new DNSGatewayGroupRemoveGatewayReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsGateways

`func (o *DNSGatewayGroupRemoveGatewayReq) GetDnsGateways() []DNSGatewayGroupRemoveGatewayReqGatewaysElt`

GetDnsGateways returns the DnsGateways field if non-nil, zero value otherwise.

### GetDnsGatewaysOk

`func (o *DNSGatewayGroupRemoveGatewayReq) GetDnsGatewaysOk() (*[]DNSGatewayGroupRemoveGatewayReqGatewaysElt, bool)`

GetDnsGatewaysOk returns a tuple with the DnsGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsGateways

`func (o *DNSGatewayGroupRemoveGatewayReq) SetDnsGateways(v []DNSGatewayGroupRemoveGatewayReqGatewaysElt)`

SetDnsGateways sets DnsGateways field to given value.

### HasDnsGateways

`func (o *DNSGatewayGroupRemoveGatewayReq) HasDnsGateways() bool`

HasDnsGateways returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


