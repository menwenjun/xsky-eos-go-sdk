# \LicensesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableLicense**](LicensesAPI.md#DisableLicense) | **Post** /licenses/{license_id}:disable | 
[**DownloadLicenseKey**](LicensesAPI.md#DownloadLicenseKey) | **Get** /licenses/key | 
[**GetLicense**](LicensesAPI.md#GetLicense) | **Get** /licenses/{license_id} | 
[**GetLicenseSummary**](LicensesAPI.md#GetLicenseSummary) | **Get** /licenses/summary | 
[**ListLicenses**](LicensesAPI.md#ListLicenses) | **Get** /licenses/ | 
[**RegisterLicense**](LicensesAPI.md#RegisterLicense) | **Post** /licenses/ | 



## DisableLicense

> LicenseResp DisableLicense(ctx, licenseId).Force(force).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/xsky-eos-go-sdk"
)

func main() {
	licenseId := int64(789) // int64 | the license id
	force := true // bool | force disable not expired license (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.DisableLicense(context.Background(), licenseId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.DisableLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableLicense`: LicenseResp
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.DisableLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licenseId** | **int64** | the license id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force disable not expired license | 

### Return type

[**LicenseResp**](LicenseResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadLicenseKey

> *os.File DownloadLicenseKey(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/xsky-eos-go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.DownloadLicenseKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.DownloadLicenseKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadLicenseKey`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.DownloadLicenseKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLicenseKeyRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicense

> LicenseResp GetLicense(ctx, licenseId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/xsky-eos-go-sdk"
)

func main() {
	licenseId := int64(789) // int64 | the license id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.GetLicense(context.Background(), licenseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.GetLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicense`: LicenseResp
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.GetLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licenseId** | **int64** | the license id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LicenseResp**](LicenseResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseSummary

> LicenseSummaryResp GetLicenseSummary(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/xsky-eos-go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.GetLicenseSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.GetLicenseSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicenseSummary`: LicenseSummaryResp
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.GetLicenseSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseSummaryRequest struct via the builder pattern


### Return type

[**LicenseSummaryResp**](LicenseSummaryResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLicenses

> LicensesResp ListLicenses(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).Active(active).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/xsky-eos-go-sdk"
)

func main() {
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	active := true // bool | filter license by active status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.ListLicenses(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).Active(active).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.ListLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLicenses`: LicensesResp
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.ListLicenses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **active** | **bool** | filter license by active status | 

### Return type

[**LicensesResp**](LicensesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterLicense

> LicenseResp RegisterLicense(ctx).License(license).Force(force).DryRun(dryRun).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/xsky-eos-go-sdk"
)

func main() {
	license := os.NewFile(1234, "some_file") // *os.File | license info
	force := true // bool | force activate (optional)
	dryRun := true // bool | dry run upload license file (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.RegisterLicense(context.Background()).License(license).Force(force).DryRun(dryRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.RegisterLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterLicense`: LicenseResp
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.RegisterLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **license** | ***os.File** | license info | 
 **force** | **bool** | force activate | 
 **dryRun** | **bool** | dry run upload license file | 

### Return type

[**LicenseResp**](LicenseResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

