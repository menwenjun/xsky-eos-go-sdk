# \AccessTokensAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessToken**](AccessTokensAPI.md#CreateAccessToken) | **Post** /access-tokens/ | 
[**DeleteAccessToken**](AccessTokensAPI.md#DeleteAccessToken) | **Delete** /access-tokens/{access_token_id} | 
[**GetAccessToken**](AccessTokensAPI.md#GetAccessToken) | **Get** /access-tokens/{access_token_id} | 
[**GetAccessTokenByUUID**](AccessTokensAPI.md#GetAccessTokenByUUID) | **Get** /access-tokens/{uuid} | 
[**ListAccessTokens**](AccessTokensAPI.md#ListAccessTokens) | **Get** /access-tokens/ | 
[**RegenerateAccessToken**](AccessTokensAPI.md#RegenerateAccessToken) | **Post** /access-tokens/{access_token_id}:regenerate | 
[**UpdateAccessToken**](AccessTokensAPI.md#UpdateAccessToken) | **Patch** /access-tokens/{access_token_id} | 
[**ValidateAccessToken**](AccessTokensAPI.md#ValidateAccessToken) | **Post** /access-tokens:validate | 



## CreateAccessToken

> AccessTokenCreateResp CreateAccessToken(ctx).Body(body).Execute()





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
	body := *openapiclient.NewAccessTokenCreateReq() // AccessTokenCreateReq | access token info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessTokensAPI.CreateAccessToken(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPI.CreateAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccessToken`: AccessTokenCreateResp
	fmt.Fprintf(os.Stdout, "Response from `AccessTokensAPI.CreateAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccessTokenCreateReq**](AccessTokenCreateReq.md) | access token info | 

### Return type

[**AccessTokenCreateResp**](AccessTokenCreateResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessToken

> DeleteAccessToken(ctx, accessTokenId).Execute()





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
	accessTokenId := int64(789) // int64 | access token id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessTokensAPI.DeleteAccessToken(context.Background(), accessTokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPI.DeleteAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessTokenId** | **int64** | access token id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessTokenRequest struct via the builder pattern


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


## GetAccessToken

> AccessTokenResp GetAccessToken(ctx, accessTokenId).Execute()





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
	accessTokenId := int64(789) // int64 | access token id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessTokensAPI.GetAccessToken(context.Background(), accessTokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPI.GetAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessToken`: AccessTokenResp
	fmt.Fprintf(os.Stdout, "Response from `AccessTokensAPI.GetAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessTokenId** | **int64** | access token id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessTokenResp**](AccessTokenResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessTokenByUUID

> AccessTokenResp GetAccessTokenByUUID(ctx, uuid).Execute()





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
	uuid := "uuid_example" // string | access token uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessTokensAPI.GetAccessTokenByUUID(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPI.GetAccessTokenByUUID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessTokenByUUID`: AccessTokenResp
	fmt.Fprintf(os.Stdout, "Response from `AccessTokensAPI.GetAccessTokenByUUID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | access token uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessTokenByUUIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessTokenResp**](AccessTokenResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessTokens

> AccessTokensResp ListAccessTokens(ctx).Limit(limit).Offset(offset).UserId(userId).Role(role).All(all).Execute()





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
	userId := int64(789) // int64 | The id of user access tokens belong to (optional)
	role := "role_example" // string | filter access tokens by role (optional)
	all := true // bool | show all access tokens (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessTokensAPI.ListAccessTokens(context.Background()).Limit(limit).Offset(offset).UserId(userId).Role(role).All(all).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPI.ListAccessTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccessTokens`: AccessTokensResp
	fmt.Fprintf(os.Stdout, "Response from `AccessTokensAPI.ListAccessTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccessTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **userId** | **int64** | The id of user access tokens belong to | 
 **role** | **string** | filter access tokens by role | 
 **all** | **bool** | show all access tokens | 

### Return type

[**AccessTokensResp**](AccessTokensResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegenerateAccessToken

> AccessTokenCreateResp RegenerateAccessToken(ctx, accessTokenId).Execute()





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
	accessTokenId := int64(789) // int64 | access token id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessTokensAPI.RegenerateAccessToken(context.Background(), accessTokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPI.RegenerateAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegenerateAccessToken`: AccessTokenCreateResp
	fmt.Fprintf(os.Stdout, "Response from `AccessTokensAPI.RegenerateAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessTokenId** | **int64** | access token id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessTokenCreateResp**](AccessTokenCreateResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessToken

> AccessTokenResp UpdateAccessToken(ctx, accessTokenId).Body(body).Execute()





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
	accessTokenId := int64(789) // int64 | access token id
	body := *openapiclient.NewAccessTokenUpdateReq() // AccessTokenUpdateReq | access token info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessTokensAPI.UpdateAccessToken(context.Background(), accessTokenId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPI.UpdateAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccessToken`: AccessTokenResp
	fmt.Fprintf(os.Stdout, "Response from `AccessTokensAPI.UpdateAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessTokenId** | **int64** | access token id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AccessTokenUpdateReq**](AccessTokenUpdateReq.md) | access token info | 

### Return type

[**AccessTokenResp**](AccessTokenResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAccessToken

> AccessTokenResp ValidateAccessToken(ctx).SubjectAccessToken(subjectAccessToken).Execute()





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
	subjectAccessToken := "subjectAccessToken_example" // string | access token to be validated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessTokensAPI.ValidateAccessToken(context.Background()).SubjectAccessToken(subjectAccessToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPI.ValidateAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateAccessToken`: AccessTokenResp
	fmt.Fprintf(os.Stdout, "Response from `AccessTokensAPI.ValidateAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subjectAccessToken** | **string** | access token to be validated | 

### Return type

[**AccessTokenResp**](AccessTokenResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

