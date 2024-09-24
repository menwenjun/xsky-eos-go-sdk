# LicenseSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityBase** | **int64** |  | 
**ExpiredFeatures** | **[]string** |  | 
**ExpiredProtocols** | **[]string** |  | 
**Features** | **[]string** |  | 
**LicenseInvalid** | **bool** |  | 
**Limits** | **map[string]string** | ProductLimits defines product limits | 
**OspCapacityBase** | **int64** |  | 
**ProductInfo** | [**ProductInfo**](ProductInfo.md) |  | 
**Protocols** | **[]string** |  | 
**UsedQuota** | [**UsedQuota**](UsedQuota.md) |  | 

## Methods

### NewLicenseSummary

`func NewLicenseSummary(capacityBase int64, expiredFeatures []string, expiredProtocols []string, features []string, licenseInvalid bool, limits map[string]string, ospCapacityBase int64, productInfo ProductInfo, protocols []string, usedQuota UsedQuota, ) *LicenseSummary`

NewLicenseSummary instantiates a new LicenseSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseSummaryWithDefaults

`func NewLicenseSummaryWithDefaults() *LicenseSummary`

NewLicenseSummaryWithDefaults instantiates a new LicenseSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityBase

`func (o *LicenseSummary) GetCapacityBase() int64`

GetCapacityBase returns the CapacityBase field if non-nil, zero value otherwise.

### GetCapacityBaseOk

`func (o *LicenseSummary) GetCapacityBaseOk() (*int64, bool)`

GetCapacityBaseOk returns a tuple with the CapacityBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityBase

`func (o *LicenseSummary) SetCapacityBase(v int64)`

SetCapacityBase sets CapacityBase field to given value.


### GetExpiredFeatures

`func (o *LicenseSummary) GetExpiredFeatures() []string`

GetExpiredFeatures returns the ExpiredFeatures field if non-nil, zero value otherwise.

### GetExpiredFeaturesOk

`func (o *LicenseSummary) GetExpiredFeaturesOk() (*[]string, bool)`

GetExpiredFeaturesOk returns a tuple with the ExpiredFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredFeatures

`func (o *LicenseSummary) SetExpiredFeatures(v []string)`

SetExpiredFeatures sets ExpiredFeatures field to given value.


### GetExpiredProtocols

`func (o *LicenseSummary) GetExpiredProtocols() []string`

GetExpiredProtocols returns the ExpiredProtocols field if non-nil, zero value otherwise.

### GetExpiredProtocolsOk

`func (o *LicenseSummary) GetExpiredProtocolsOk() (*[]string, bool)`

GetExpiredProtocolsOk returns a tuple with the ExpiredProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredProtocols

`func (o *LicenseSummary) SetExpiredProtocols(v []string)`

SetExpiredProtocols sets ExpiredProtocols field to given value.


### GetFeatures

`func (o *LicenseSummary) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *LicenseSummary) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *LicenseSummary) SetFeatures(v []string)`

SetFeatures sets Features field to given value.


### GetLicenseInvalid

`func (o *LicenseSummary) GetLicenseInvalid() bool`

GetLicenseInvalid returns the LicenseInvalid field if non-nil, zero value otherwise.

### GetLicenseInvalidOk

`func (o *LicenseSummary) GetLicenseInvalidOk() (*bool, bool)`

GetLicenseInvalidOk returns a tuple with the LicenseInvalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseInvalid

`func (o *LicenseSummary) SetLicenseInvalid(v bool)`

SetLicenseInvalid sets LicenseInvalid field to given value.


### GetLimits

`func (o *LicenseSummary) GetLimits() map[string]string`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *LicenseSummary) GetLimitsOk() (*map[string]string, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *LicenseSummary) SetLimits(v map[string]string)`

SetLimits sets Limits field to given value.


### GetOspCapacityBase

`func (o *LicenseSummary) GetOspCapacityBase() int64`

GetOspCapacityBase returns the OspCapacityBase field if non-nil, zero value otherwise.

### GetOspCapacityBaseOk

`func (o *LicenseSummary) GetOspCapacityBaseOk() (*int64, bool)`

GetOspCapacityBaseOk returns a tuple with the OspCapacityBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspCapacityBase

`func (o *LicenseSummary) SetOspCapacityBase(v int64)`

SetOspCapacityBase sets OspCapacityBase field to given value.


### GetProductInfo

`func (o *LicenseSummary) GetProductInfo() ProductInfo`

GetProductInfo returns the ProductInfo field if non-nil, zero value otherwise.

### GetProductInfoOk

`func (o *LicenseSummary) GetProductInfoOk() (*ProductInfo, bool)`

GetProductInfoOk returns a tuple with the ProductInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductInfo

`func (o *LicenseSummary) SetProductInfo(v ProductInfo)`

SetProductInfo sets ProductInfo field to given value.


### GetProtocols

`func (o *LicenseSummary) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *LicenseSummary) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *LicenseSummary) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.


### GetUsedQuota

`func (o *LicenseSummary) GetUsedQuota() UsedQuota`

GetUsedQuota returns the UsedQuota field if non-nil, zero value otherwise.

### GetUsedQuotaOk

`func (o *LicenseSummary) GetUsedQuotaOk() (*UsedQuota, bool)`

GetUsedQuotaOk returns a tuple with the UsedQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuota

`func (o *LicenseSummary) SetUsedQuota(v UsedQuota)`

SetUsedQuota sets UsedQuota field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


