# \ClientLunMappingsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClientLunMapping**](ClientLunMappingsAPI.md#GetClientLunMapping) | **Get** /client-lun-mappings/{client_lun_mapping_id} | 
[**ListClientLunMappings**](ClientLunMappingsAPI.md#ListClientLunMappings) | **Get** /client-lun-mappings/ | 



## GetClientLunMapping

> ClientLunMappingResp GetClientLunMapping(ctx, clientLunMappingId).IsFetchLunInfo(isFetchLunInfo).Execute()





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
	clientLunMappingId := int64(789) // int64 | client lun mapping id
	isFetchLunInfo := true // bool | whether to fetch lun info from target (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientLunMappingsAPI.GetClientLunMapping(context.Background(), clientLunMappingId).IsFetchLunInfo(isFetchLunInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientLunMappingsAPI.GetClientLunMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientLunMapping`: ClientLunMappingResp
	fmt.Fprintf(os.Stdout, "Response from `ClientLunMappingsAPI.GetClientLunMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientLunMappingId** | **int64** | client lun mapping id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientLunMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isFetchLunInfo** | **bool** | whether to fetch lun info from target | 

### Return type

[**ClientLunMappingResp**](ClientLunMappingResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClientLunMappings

> ClientLunMappingsResp ListClientLunMappings(ctx).MappingGroupId(mappingGroupId).Limit(limit).Offset(offset).Execute()





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
	mappingGroupId := int64(789) // int64 | mapping group id
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientLunMappingsAPI.ListClientLunMappings(context.Background()).MappingGroupId(mappingGroupId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientLunMappingsAPI.ListClientLunMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClientLunMappings`: ClientLunMappingsResp
	fmt.Fprintf(os.Stdout, "Response from `ClientLunMappingsAPI.ListClientLunMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClientLunMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mappingGroupId** | **int64** | mapping group id | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 

### Return type

[**ClientLunMappingsResp**](ClientLunMappingsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

