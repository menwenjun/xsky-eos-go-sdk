# \ExtCompatConfigsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListExtCompatConf**](ExtCompatConfigsAPI.md#ListExtCompatConf) | **Get** /ext-compat-configs/ | 
[**SetExtCompatConf**](ExtCompatConfigsAPI.md#SetExtCompatConf) | **Post** /ext-compat-configs/ | 



## ListExtCompatConf

> ExtCompatConfResp ListExtCompatConf(ctx).ExtName(extName).FuncName(funcName).Policy(policy).HostId(hostId).Execute()





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
	extName := "extName_example" // string | filter the external interface name of ext compat conf (optional)
	funcName := "funcName_example" // string | filter the function name of ext compat conf (optional)
	policy := "policy_example" // string | filter the policy of ext compat conf (optional)
	hostId := int64(789) // int64 | filter the host of ext compat conf, 0 for global config (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtCompatConfigsAPI.ListExtCompatConf(context.Background()).ExtName(extName).FuncName(funcName).Policy(policy).HostId(hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtCompatConfigsAPI.ListExtCompatConf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExtCompatConf`: ExtCompatConfResp
	fmt.Fprintf(os.Stdout, "Response from `ExtCompatConfigsAPI.ListExtCompatConf`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExtCompatConfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extName** | **string** | filter the external interface name of ext compat conf | 
 **funcName** | **string** | filter the function name of ext compat conf | 
 **policy** | **string** | filter the policy of ext compat conf | 
 **hostId** | **int64** | filter the host of ext compat conf, 0 for global config | 

### Return type

[**ExtCompatConfResp**](ExtCompatConfResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetExtCompatConf

> SetExtCompatConfResp SetExtCompatConf(ctx).Body(body).Execute()





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
	body := *openapiclient.NewSetExtCompatConfReq(*openapiclient.NewExtCompatConf()) // SetExtCompatConfReq | ext compat conf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtCompatConfigsAPI.SetExtCompatConf(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtCompatConfigsAPI.SetExtCompatConf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetExtCompatConf`: SetExtCompatConfResp
	fmt.Fprintf(os.Stdout, "Response from `ExtCompatConfigsAPI.SetExtCompatConf`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetExtCompatConfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SetExtCompatConfReq**](SetExtCompatConfReq.md) | ext compat conf | 

### Return type

[**SetExtCompatConfResp**](SetExtCompatConfResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

