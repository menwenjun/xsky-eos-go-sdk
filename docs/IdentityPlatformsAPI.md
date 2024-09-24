# \IdentityPlatformsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentityPlatform**](IdentityPlatformsAPI.md#CreateIdentityPlatform) | **Post** /identity-platforms/ | 
[**DeleteIdentityPlatform**](IdentityPlatformsAPI.md#DeleteIdentityPlatform) | **Delete** /identity-platforms/{identity_platform_id} | 
[**GetIdentityPlatform**](IdentityPlatformsAPI.md#GetIdentityPlatform) | **Get** /identity-platforms/{identity_platform_id} | 
[**ListIdentityPlatforms**](IdentityPlatformsAPI.md#ListIdentityPlatforms) | **Get** /identity-platforms/ | 
[**LoginSSOTypes**](IdentityPlatformsAPI.md#LoginSSOTypes) | **Get** /identity-platforms/types | 
[**RegenerateIdentityKey**](IdentityPlatformsAPI.md#RegenerateIdentityKey) | **Post** /identity-platforms/{identity_platform_id}:regenerate | 
[**UpdateIdentityPlatform**](IdentityPlatformsAPI.md#UpdateIdentityPlatform) | **Patch** /identity-platforms/{identity_platform_id} | 



## CreateIdentityPlatform

> IdentityPlatformResp CreateIdentityPlatform(ctx).Body(body).Execute()





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
	body := *openapiclient.NewIdentityPlatformCreateReq() // IdentityPlatformCreateReq | identity platform info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityPlatformsAPI.CreateIdentityPlatform(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityPlatformsAPI.CreateIdentityPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIdentityPlatform`: IdentityPlatformResp
	fmt.Fprintf(os.Stdout, "Response from `IdentityPlatformsAPI.CreateIdentityPlatform`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IdentityPlatformCreateReq**](IdentityPlatformCreateReq.md) | identity platform info | 

### Return type

[**IdentityPlatformResp**](IdentityPlatformResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityPlatform

> DeleteIdentityPlatform(ctx, identityPlatformId).Execute()





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
	identityPlatformId := int64(789) // int64 | identity platform id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityPlatformsAPI.DeleteIdentityPlatform(context.Background(), identityPlatformId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityPlatformsAPI.DeleteIdentityPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityPlatformId** | **int64** | identity platform id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityPlatformRequest struct via the builder pattern


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


## GetIdentityPlatform

> IdentityPlatformResp GetIdentityPlatform(ctx, identityPlatformId).Execute()





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
	identityPlatformId := int64(789) // int64 | identity platform id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityPlatformsAPI.GetIdentityPlatform(context.Background(), identityPlatformId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityPlatformsAPI.GetIdentityPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentityPlatform`: IdentityPlatformResp
	fmt.Fprintf(os.Stdout, "Response from `IdentityPlatformsAPI.GetIdentityPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityPlatformId** | **int64** | identity platform id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityPlatformResp**](IdentityPlatformResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentityPlatforms

> IdentityPlatformsResp ListIdentityPlatforms(ctx).Limit(limit).Offset(offset).Execute()





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
	resp, r, err := apiClient.IdentityPlatformsAPI.ListIdentityPlatforms(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityPlatformsAPI.ListIdentityPlatforms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdentityPlatforms`: IdentityPlatformsResp
	fmt.Fprintf(os.Stdout, "Response from `IdentityPlatformsAPI.ListIdentityPlatforms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdentityPlatformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 

### Return type

[**IdentityPlatformsResp**](IdentityPlatformsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginSSOTypes

> LoginSsoTypesResp LoginSSOTypes(ctx).Execute()





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
	resp, r, err := apiClient.IdentityPlatformsAPI.LoginSSOTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityPlatformsAPI.LoginSSOTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoginSSOTypes`: LoginSsoTypesResp
	fmt.Fprintf(os.Stdout, "Response from `IdentityPlatformsAPI.LoginSSOTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLoginSSOTypesRequest struct via the builder pattern


### Return type

[**LoginSsoTypesResp**](LoginSsoTypesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegenerateIdentityKey

> IdentityPlatformResp RegenerateIdentityKey(ctx, identityPlatformId).Execute()





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
	identityPlatformId := int64(789) // int64 | identity platform id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityPlatformsAPI.RegenerateIdentityKey(context.Background(), identityPlatformId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityPlatformsAPI.RegenerateIdentityKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegenerateIdentityKey`: IdentityPlatformResp
	fmt.Fprintf(os.Stdout, "Response from `IdentityPlatformsAPI.RegenerateIdentityKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityPlatformId** | **int64** | identity platform id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateIdentityKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityPlatformResp**](IdentityPlatformResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentityPlatform

> IdentityPlatformResp UpdateIdentityPlatform(ctx, identityPlatformId).Body(body).Execute()





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
	identityPlatformId := int64(789) // int64 | identity platform id
	body := *openapiclient.NewIdentityPlatformUpdateReq() // IdentityPlatformUpdateReq | identity platform info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityPlatformsAPI.UpdateIdentityPlatform(context.Background(), identityPlatformId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityPlatformsAPI.UpdateIdentityPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIdentityPlatform`: IdentityPlatformResp
	fmt.Fprintf(os.Stdout, "Response from `IdentityPlatformsAPI.UpdateIdentityPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityPlatformId** | **int64** | identity platform id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentityPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IdentityPlatformUpdateReq**](IdentityPlatformUpdateReq.md) | identity platform info | 

### Return type

[**IdentityPlatformResp**](IdentityPlatformResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

