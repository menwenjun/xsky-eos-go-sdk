# \NetworkDiagnosesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkDiagnosis**](NetworkDiagnosesAPI.md#CreateNetworkDiagnosis) | **Post** /network-diagnoses/ | 
[**DeleteNetworkDiagnosis**](NetworkDiagnosesAPI.md#DeleteNetworkDiagnosis) | **Delete** /network-diagnoses/{network_diagnosis_id} | 
[**GetNetworkDiagnosis**](NetworkDiagnosesAPI.md#GetNetworkDiagnosis) | **Get** /network-diagnoses/{network_diagnosis_id} | 
[**ListNetworkDiagnoses**](NetworkDiagnosesAPI.md#ListNetworkDiagnoses) | **Get** /network-diagnoses/ | 
[**StopNetworkDiagnosis**](NetworkDiagnosesAPI.md#StopNetworkDiagnosis) | **Post** /network-diagnoses/{network_diagnosis_id}:stop | 



## CreateNetworkDiagnosis

> NetworkDiagnosisResp CreateNetworkDiagnosis(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewNetworkDiagnosisCreateReq(*openapiclient.NewNetworkDiagnosisCreateReqDiagnosis([]int64{int64(123)}, "HostScope_example", []string{"Networks_example"})) // NetworkDiagnosisCreateReq | network diagnosis info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDiagnosesAPI.CreateNetworkDiagnosis(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDiagnosesAPI.CreateNetworkDiagnosis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkDiagnosis`: NetworkDiagnosisResp
	fmt.Fprintf(os.Stdout, "Response from `NetworkDiagnosesAPI.CreateNetworkDiagnosis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDiagnosisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NetworkDiagnosisCreateReq**](NetworkDiagnosisCreateReq.md) | network diagnosis info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**NetworkDiagnosisResp**](NetworkDiagnosisResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkDiagnosis

> DeleteNetworkDiagnosis(ctx, networkDiagnosisId).Execute()





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
	networkDiagnosisId := int64(789) // int64 | network diagnosis id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDiagnosesAPI.DeleteNetworkDiagnosis(context.Background(), networkDiagnosisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDiagnosesAPI.DeleteNetworkDiagnosis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDiagnosisId** | **int64** | network diagnosis id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDiagnosisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDiagnosis

> NetworkDiagnosisResp GetNetworkDiagnosis(ctx, networkDiagnosisId).Execute()





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
	networkDiagnosisId := int64(789) // int64 | network diagnosis id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDiagnosesAPI.GetNetworkDiagnosis(context.Background(), networkDiagnosisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDiagnosesAPI.GetNetworkDiagnosis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDiagnosis`: NetworkDiagnosisResp
	fmt.Fprintf(os.Stdout, "Response from `NetworkDiagnosesAPI.GetNetworkDiagnosis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDiagnosisId** | **int64** | network diagnosis id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDiagnosisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDiagnosisResp**](NetworkDiagnosisResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkDiagnoses

> NetworkDiagnosesResp ListNetworkDiagnoses(ctx).Limit(limit).Offset(offset).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDiagnosesAPI.ListNetworkDiagnoses(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDiagnosesAPI.ListNetworkDiagnoses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkDiagnoses`: NetworkDiagnosesResp
	fmt.Fprintf(os.Stdout, "Response from `NetworkDiagnosesAPI.ListNetworkDiagnoses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkDiagnosesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 

### Return type

[**NetworkDiagnosesResp**](NetworkDiagnosesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopNetworkDiagnosis

> NetworkDiagnosisResp StopNetworkDiagnosis(ctx, networkDiagnosisId).Execute()





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
	networkDiagnosisId := int64(789) // int64 | network diagnosis id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDiagnosesAPI.StopNetworkDiagnosis(context.Background(), networkDiagnosisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDiagnosesAPI.StopNetworkDiagnosis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopNetworkDiagnosis`: NetworkDiagnosisResp
	fmt.Fprintf(os.Stdout, "Response from `NetworkDiagnosesAPI.StopNetworkDiagnosis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDiagnosisId** | **int64** | network diagnosis id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopNetworkDiagnosisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDiagnosisResp**](NetworkDiagnosisResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

