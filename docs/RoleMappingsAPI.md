# \RoleMappingsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoleMapping**](RoleMappingsAPI.md#CreateRoleMapping) | **Post** /role-mappings/ | 
[**DeleteRoleMapping**](RoleMappingsAPI.md#DeleteRoleMapping) | **Delete** /role-mappings/{role_mapping_id} | 
[**GetRoleMapping**](RoleMappingsAPI.md#GetRoleMapping) | **Get** /role-mappings/{role_mapping_id} | 
[**ListRoleMappings**](RoleMappingsAPI.md#ListRoleMappings) | **Get** /role-mappings/ | 
[**UpdateRoleMapping**](RoleMappingsAPI.md#UpdateRoleMapping) | **Patch** /role-mappings/{role_mapping_id} | 



## CreateRoleMapping

> RoleMappingResp CreateRoleMapping(ctx).Body(body).Execute()





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
	body := *openapiclient.NewRoleMappingCreateReq() // RoleMappingCreateReq | role mapping info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleMappingsAPI.CreateRoleMapping(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMappingsAPI.CreateRoleMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoleMapping`: RoleMappingResp
	fmt.Fprintf(os.Stdout, "Response from `RoleMappingsAPI.CreateRoleMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RoleMappingCreateReq**](RoleMappingCreateReq.md) | role mapping info | 

### Return type

[**RoleMappingResp**](RoleMappingResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoleMapping

> DeleteRoleMapping(ctx, roleMappingId).Execute()





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
	roleMappingId := int64(789) // int64 | role mapping id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleMappingsAPI.DeleteRoleMapping(context.Background(), roleMappingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMappingsAPI.DeleteRoleMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleMappingId** | **int64** | role mapping id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleMappingRequest struct via the builder pattern


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


## GetRoleMapping

> RoleMappingResp GetRoleMapping(ctx, roleMappingId).Execute()





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
	roleMappingId := int64(789) // int64 | role mapping id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleMappingsAPI.GetRoleMapping(context.Background(), roleMappingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMappingsAPI.GetRoleMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleMapping`: RoleMappingResp
	fmt.Fprintf(os.Stdout, "Response from `RoleMappingsAPI.GetRoleMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleMappingId** | **int64** | role mapping id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoleMappingResp**](RoleMappingResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleMappings

> RoleMappingsResp ListRoleMappings(ctx).Limit(limit).Offset(offset).IdentityPlatformId(identityPlatformId).Execute()





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
	identityPlatformId := int64(789) // int64 | identity platform id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleMappingsAPI.ListRoleMappings(context.Background()).Limit(limit).Offset(offset).IdentityPlatformId(identityPlatformId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMappingsAPI.ListRoleMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoleMappings`: RoleMappingsResp
	fmt.Fprintf(os.Stdout, "Response from `RoleMappingsAPI.ListRoleMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoleMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **identityPlatformId** | **int64** | identity platform id | 

### Return type

[**RoleMappingsResp**](RoleMappingsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleMapping

> RoleMappingResp UpdateRoleMapping(ctx, roleMappingId).Body(body).Execute()





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
	roleMappingId := int64(789) // int64 | role mapping id
	body := *openapiclient.NewRoleMappingUpdateReq() // RoleMappingUpdateReq | role mapping info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleMappingsAPI.UpdateRoleMapping(context.Background(), roleMappingId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMappingsAPI.UpdateRoleMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoleMapping`: RoleMappingResp
	fmt.Fprintf(os.Stdout, "Response from `RoleMappingsAPI.UpdateRoleMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleMappingId** | **int64** | role mapping id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RoleMappingUpdateReq**](RoleMappingUpdateReq.md) | role mapping info | 

### Return type

[**RoleMappingResp**](RoleMappingResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

