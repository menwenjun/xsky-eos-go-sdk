# \OsZonePairsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListOSZonePairs**](OsZonePairsAPI.md#ListOSZonePairs) | **Get** /os-zone-pairs/ | 



## ListOSZonePairs

> OSZonePairsResp ListOSZonePairs(ctx).MinClockDiffAbs(minClockDiffAbs).ClusterId(clusterId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	minClockDiffAbs := "minClockDiffAbs_example" // string | min clock diff absolute value for zone pairs (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsZonePairsAPI.ListOSZonePairs(context.Background()).MinClockDiffAbs(minClockDiffAbs).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsZonePairsAPI.ListOSZonePairs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOSZonePairs`: OSZonePairsResp
	fmt.Fprintf(os.Stdout, "Response from `OsZonePairsAPI.ListOSZonePairs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOSZonePairsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minClockDiffAbs** | **string** | min clock diff absolute value for zone pairs | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSZonePairsResp**](OSZonePairsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

