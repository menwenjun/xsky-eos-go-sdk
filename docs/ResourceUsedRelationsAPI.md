# \ResourceUsedRelationsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetResourceRelatedResources**](ResourceUsedRelationsAPI.md#GetResourceRelatedResources) | **Get** /resource-used-relations/related-resources | 



## GetResourceRelatedResources

> ResourceRelatedResourcesResp GetResourceRelatedResources(ctx).Resource(resource).Ids(ids).Execute()





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
	resource := "resource_example" // string | resource type
	ids := "ids_example" // string | resource ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceUsedRelationsAPI.GetResourceRelatedResources(context.Background()).Resource(resource).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceUsedRelationsAPI.GetResourceRelatedResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourceRelatedResources`: ResourceRelatedResourcesResp
	fmt.Fprintf(os.Stdout, "Response from `ResourceUsedRelationsAPI.GetResourceRelatedResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceRelatedResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resource** | **string** | resource type | 
 **ids** | **string** | resource ids | 

### Return type

[**ResourceRelatedResourcesResp**](ResourceRelatedResourcesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

