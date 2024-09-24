# \VipGroupsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVIPGroup**](VipGroupsAPI.md#CreateVIPGroup) | **Post** /vip-groups/ | 
[**DeleteVIPGroup**](VipGroupsAPI.md#DeleteVIPGroup) | **Delete** /vip-groups/{vip_group_id} | 
[**GetVIPGroup**](VipGroupsAPI.md#GetVIPGroup) | **Get** /vip-groups/{vip_group_id} | 
[**ListVIPGroups**](VipGroupsAPI.md#ListVIPGroups) | **Get** /vip-groups/ | 
[**RedeployVIPGroup**](VipGroupsAPI.md#RedeployVIPGroup) | **Post** /vip-groups/{vip_group_id}:redeploy | 
[**UpdateVIPGroup**](VipGroupsAPI.md#UpdateVIPGroup) | **Patch** /vip-groups/{vip_group_id} | 



## CreateVIPGroup

> VIPGroupResp CreateVIPGroup(ctx).VipGroup(vipGroup).Execute()





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
	vipGroup := *openapiclient.NewVIPGroupCreateReq(*openapiclient.NewVIPGroupCreateReqVIPGroup("Network_example", int64(123), "ResourceType_example", []openapiclient.VIPArgs{*openapiclient.NewVIPArgs("Ip_example", int64(123))})) // VIPGroupCreateReq | vip group info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipGroupsAPI.CreateVIPGroup(context.Background()).VipGroup(vipGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipGroupsAPI.CreateVIPGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVIPGroup`: VIPGroupResp
	fmt.Fprintf(os.Stdout, "Response from `VipGroupsAPI.CreateVIPGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVIPGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vipGroup** | [**VIPGroupCreateReq**](VIPGroupCreateReq.md) | vip group info | 

### Return type

[**VIPGroupResp**](VIPGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVIPGroup

> VIPGroupResp DeleteVIPGroup(ctx, vipGroupId).Execute()





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
	vipGroupId := int64(789) // int64 | vip group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipGroupsAPI.DeleteVIPGroup(context.Background(), vipGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipGroupsAPI.DeleteVIPGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVIPGroup`: VIPGroupResp
	fmt.Fprintf(os.Stdout, "Response from `VipGroupsAPI.DeleteVIPGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vipGroupId** | **int64** | vip group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVIPGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VIPGroupResp**](VIPGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVIPGroup

> VIPGroupResp GetVIPGroup(ctx, vipGroupId).Execute()





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
	vipGroupId := int64(789) // int64 | vip group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipGroupsAPI.GetVIPGroup(context.Background(), vipGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipGroupsAPI.GetVIPGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVIPGroup`: VIPGroupResp
	fmt.Fprintf(os.Stdout, "Response from `VipGroupsAPI.GetVIPGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vipGroupId** | **int64** | vip group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVIPGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VIPGroupResp**](VIPGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVIPGroups

> VIPGroupResps ListVIPGroups(ctx).Limit(limit).Offset(offset).ResourceType(resourceType).ResourceId(resourceId).Execute()





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
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	resourceType := "resourceType_example" // string | resource type (optional)
	resourceId := int64(789) // int64 | resource id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipGroupsAPI.ListVIPGroups(context.Background()).Limit(limit).Offset(offset).ResourceType(resourceType).ResourceId(resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipGroupsAPI.ListVIPGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVIPGroups`: VIPGroupResps
	fmt.Fprintf(os.Stdout, "Response from `VipGroupsAPI.ListVIPGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVIPGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **resourceType** | **string** | resource type | 
 **resourceId** | **int64** | resource id | 

### Return type

[**VIPGroupResps**](VIPGroupResps.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedeployVIPGroup

> VIPGroupResp RedeployVIPGroup(ctx, vipGroupId).Execute()





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
	vipGroupId := int64(789) // int64 | vip group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipGroupsAPI.RedeployVIPGroup(context.Background(), vipGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipGroupsAPI.RedeployVIPGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedeployVIPGroup`: VIPGroupResp
	fmt.Fprintf(os.Stdout, "Response from `VipGroupsAPI.RedeployVIPGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vipGroupId** | **int64** | vip group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeployVIPGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VIPGroupResp**](VIPGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVIPGroup

> VIPGroupResp UpdateVIPGroup(ctx, vipGroupId).VipGroup(vipGroup).Execute()





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
	vipGroupId := int64(789) // int64 | vip group id
	vipGroup := *openapiclient.NewVIPGroupUpdateReq(*openapiclient.NewVIPGroupUpdateReqVIPGroup()) // VIPGroupUpdateReq | vip group info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipGroupsAPI.UpdateVIPGroup(context.Background(), vipGroupId).VipGroup(vipGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipGroupsAPI.UpdateVIPGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVIPGroup`: VIPGroupResp
	fmt.Fprintf(os.Stdout, "Response from `VipGroupsAPI.UpdateVIPGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vipGroupId** | **int64** | vip group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVIPGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vipGroup** | [**VIPGroupUpdateReq**](VIPGroupUpdateReq.md) | vip group info | 

### Return type

[**VIPGroupResp**](VIPGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

