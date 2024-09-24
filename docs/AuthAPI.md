# \AuthAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRSAKey**](AuthAPI.md#CreateRSAKey) | **Post** /auth/rsa-keys | 
[**CreateToken**](AuthAPI.md#CreateToken) | **Post** /auth/tokens | 
[**GetAuthSecurityPolicy**](AuthAPI.md#GetAuthSecurityPolicy) | **Get** /auth/security-policy | 
[**Login**](AuthAPI.md#Login) | **Post** /auth/tokens:login | 
[**Logout**](AuthAPI.md#Logout) | **Post** /auth/tokens:logout | 
[**UpdateAuthSecurityPolicy**](AuthAPI.md#UpdateAuthSecurityPolicy) | **Patch** /auth/security-policy | 
[**ValidatePrivilegedToken**](AuthAPI.md#ValidatePrivilegedToken) | **Post** /auth/:validate-privileged-token | 
[**Verification**](AuthAPI.md#Verification) | **Post** /auth/verifications | 



## CreateRSAKey

> AuthRSAKeyResp CreateRSAKey(ctx).Execute()





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
	resp, r, err := apiClient.AuthAPI.CreateRSAKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.CreateRSAKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRSAKey`: AuthRSAKeyResp
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.CreateRSAKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRSAKeyRequest struct via the builder pattern


### Return type

[**AuthRSAKeyResp**](AuthRSAKeyResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateToken

> TokenResp CreateToken(ctx).Body(body).Execute()





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
	body := *openapiclient.NewTokenCreateReq(*openapiclient.NewTokenCreateReqAuth(*openapiclient.NewTokenCreateReqAuthIdentity())) // TokenCreateReq | the identity credential

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.CreateToken(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.CreateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateToken`: TokenResp
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.CreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TokenCreateReq**](TokenCreateReq.md) | the identity credential | 

### Return type

[**TokenResp**](TokenResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthSecurityPolicy

> AuthSecurityPolicyResp GetAuthSecurityPolicy(ctx).Execute()





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
	resp, r, err := apiClient.AuthAPI.GetAuthSecurityPolicy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.GetAuthSecurityPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthSecurityPolicy`: AuthSecurityPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.GetAuthSecurityPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthSecurityPolicyRequest struct via the builder pattern


### Return type

[**AuthSecurityPolicyResp**](AuthSecurityPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> TokenResp Login(ctx).Body(body).Execute()





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
	body := *openapiclient.NewAuthLoginReq(*openapiclient.NewAuthLoginReqAuth()) // AuthLoginReq | the identity credential

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.Login(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.Login``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Login`: TokenResp
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthLoginReq**](AuthLoginReq.md) | the identity credential | 

### Return type

[**TokenResp**](TokenResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> Logout(ctx).Execute()





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
	r, err := apiClient.AuthAPI.Logout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.Logout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


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


## UpdateAuthSecurityPolicy

> AuthSecurityPolicyResp UpdateAuthSecurityPolicy(ctx).Body(body).Execute()





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
	body := *openapiclient.NewAuthSecurityPolicyUpdateReq() // AuthSecurityPolicyUpdateReq | auth security policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.UpdateAuthSecurityPolicy(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.UpdateAuthSecurityPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAuthSecurityPolicy`: AuthSecurityPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.UpdateAuthSecurityPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthSecurityPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthSecurityPolicyUpdateReq**](AuthSecurityPolicyUpdateReq.md) | auth security policy info | 

### Return type

[**AuthSecurityPolicyResp**](AuthSecurityPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePrivilegedToken

> PrivilegedTokenResp ValidatePrivilegedToken(ctx).Body(body).Execute()





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
	body := *openapiclient.NewPrivilegedTokenReq("PrivilegedToken_example") // PrivilegedTokenReq | privileged token info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.ValidatePrivilegedToken(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.ValidatePrivilegedToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidatePrivilegedToken`: PrivilegedTokenResp
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.ValidatePrivilegedToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidatePrivilegedTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PrivilegedTokenReq**](PrivilegedTokenReq.md) | privileged token info | 

### Return type

[**PrivilegedTokenResp**](PrivilegedTokenResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Verification

> Verification(ctx).Body(body).Execute()





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
	body := *openapiclient.NewSendVerificationReq() // SendVerificationReq | verification info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.Verification(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.Verification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SendVerificationReq**](SendVerificationReq.md) | verification info | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

