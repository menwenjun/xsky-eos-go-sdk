# \NetworkDiagnosisItemsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNetworkDiagnosisItems**](NetworkDiagnosisItemsAPI.md#ListNetworkDiagnosisItems) | **Get** /network-diagnosis-items/ | 



## ListNetworkDiagnosisItems

> NetworkDiagnosisItemsResp ListNetworkDiagnosisItems(ctx).Limit(limit).Offset(offset).Networks(networks).Finished(finished).Execute()





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
	networks := "networks_example" // string | network type (optional)
	finished := true // bool | diagnosis item finished or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDiagnosisItemsAPI.ListNetworkDiagnosisItems(context.Background()).Limit(limit).Offset(offset).Networks(networks).Finished(finished).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDiagnosisItemsAPI.ListNetworkDiagnosisItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkDiagnosisItems`: NetworkDiagnosisItemsResp
	fmt.Fprintf(os.Stdout, "Response from `NetworkDiagnosisItemsAPI.ListNetworkDiagnosisItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkDiagnosisItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **networks** | **string** | network type | 
 **finished** | **bool** | diagnosis item finished or not | 

### Return type

[**NetworkDiagnosisItemsResp**](NetworkDiagnosisItemsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

