# OspGatewayCreateReqGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**GatewayIp** | Pointer to **string** |  | [optional] 
**HostId** | **int64** |  | 
**Name** | **string** |  | 
**OspZoneId** | Pointer to **int64** |  | [optional] 
**Role** | **string** |  | 
**S3Port** | **int64** |  | 
**XmsPort** | Pointer to **int64** |  | [optional] 

## Methods

### NewOspGatewayCreateReqGateway

`func NewOspGatewayCreateReqGateway(hostId int64, name string, role string, s3Port int64, ) *OspGatewayCreateReqGateway`

NewOspGatewayCreateReqGateway instantiates a new OspGatewayCreateReqGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspGatewayCreateReqGatewayWithDefaults

`func NewOspGatewayCreateReqGatewayWithDefaults() *OspGatewayCreateReqGateway`

NewOspGatewayCreateReqGatewayWithDefaults instantiates a new OspGatewayCreateReqGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OspGatewayCreateReqGateway) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OspGatewayCreateReqGateway) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OspGatewayCreateReqGateway) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OspGatewayCreateReqGateway) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGatewayIp

`func (o *OspGatewayCreateReqGateway) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *OspGatewayCreateReqGateway) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *OspGatewayCreateReqGateway) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *OspGatewayCreateReqGateway) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetHostId

`func (o *OspGatewayCreateReqGateway) GetHostId() int64`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *OspGatewayCreateReqGateway) GetHostIdOk() (*int64, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *OspGatewayCreateReqGateway) SetHostId(v int64)`

SetHostId sets HostId field to given value.


### GetName

`func (o *OspGatewayCreateReqGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OspGatewayCreateReqGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OspGatewayCreateReqGateway) SetName(v string)`

SetName sets Name field to given value.


### GetOspZoneId

`func (o *OspGatewayCreateReqGateway) GetOspZoneId() int64`

GetOspZoneId returns the OspZoneId field if non-nil, zero value otherwise.

### GetOspZoneIdOk

`func (o *OspGatewayCreateReqGateway) GetOspZoneIdOk() (*int64, bool)`

GetOspZoneIdOk returns a tuple with the OspZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspZoneId

`func (o *OspGatewayCreateReqGateway) SetOspZoneId(v int64)`

SetOspZoneId sets OspZoneId field to given value.

### HasOspZoneId

`func (o *OspGatewayCreateReqGateway) HasOspZoneId() bool`

HasOspZoneId returns a boolean if a field has been set.

### GetRole

`func (o *OspGatewayCreateReqGateway) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OspGatewayCreateReqGateway) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OspGatewayCreateReqGateway) SetRole(v string)`

SetRole sets Role field to given value.


### GetS3Port

`func (o *OspGatewayCreateReqGateway) GetS3Port() int64`

GetS3Port returns the S3Port field if non-nil, zero value otherwise.

### GetS3PortOk

`func (o *OspGatewayCreateReqGateway) GetS3PortOk() (*int64, bool)`

GetS3PortOk returns a tuple with the S3Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Port

`func (o *OspGatewayCreateReqGateway) SetS3Port(v int64)`

SetS3Port sets S3Port field to given value.


### GetXmsPort

`func (o *OspGatewayCreateReqGateway) GetXmsPort() int64`

GetXmsPort returns the XmsPort field if non-nil, zero value otherwise.

### GetXmsPortOk

`func (o *OspGatewayCreateReqGateway) GetXmsPortOk() (*int64, bool)`

GetXmsPortOk returns a tuple with the XmsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmsPort

`func (o *OspGatewayCreateReqGateway) SetXmsPort(v int64)`

SetXmsPort sets XmsPort field to given value.

### HasXmsPort

`func (o *OspGatewayCreateReqGateway) HasXmsPort() bool`

HasXmsPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


