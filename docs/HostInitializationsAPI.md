# \HostInitializationsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInitializationConf**](HostInitializationsAPI.md#GetInitializationConf) | **Get** /host-initializations/conf | 
[**HostInitialization**](HostInitializationsAPI.md#HostInitialization) | **Post** /host-initializations/ | 



## GetInitializationConf

> HostInitializationConfResp GetInitializationConf(ctx).Execute()





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
	resp, r, err := apiClient.HostInitializationsAPI.GetInitializationConf(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostInitializationsAPI.GetInitializationConf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInitializationConf`: HostInitializationConfResp
	fmt.Fprintf(os.Stdout, "Response from `HostInitializationsAPI.GetInitializationConf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInitializationConfRequest struct via the builder pattern


### Return type

[**HostInitializationConfResp**](HostInitializationConfResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostInitialization

> HostInitializationResp HostInitialization(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewHostInitializationReq() // HostInitializationReq | host initialization info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostInitializationsAPI.HostInitialization(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostInitializationsAPI.HostInitialization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostInitialization`: HostInitializationResp
	fmt.Fprintf(os.Stdout, "Response from `HostInitializationsAPI.HostInitialization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHostInitializationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HostInitializationReq**](HostInitializationReq.md) | host initialization info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**HostInitializationResp**](HostInitializationResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

