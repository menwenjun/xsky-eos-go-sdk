# \DnsGatewayGroupsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDNSGatewayToGroup**](DnsGatewayGroupsAPI.md#AddDNSGatewayToGroup) | **Post** /dns-gateway-groups/{group_id}:add-gateway | 
[**CreateDNSGatewayGroup**](DnsGatewayGroupsAPI.md#CreateDNSGatewayGroup) | **Post** /dns-gateway-groups/ | 
[**DeleteDNSGatewayGroup**](DnsGatewayGroupsAPI.md#DeleteDNSGatewayGroup) | **Delete** /dns-gateway-groups/{group_id} | 
[**GetDNSGatewayGroup**](DnsGatewayGroupsAPI.md#GetDNSGatewayGroup) | **Get** /dns-gateway-groups/{group_id} | 
[**ListDNSGatewayGroups**](DnsGatewayGroupsAPI.md#ListDNSGatewayGroups) | **Get** /dns-gateway-groups/ | 
[**RedeployDNSGatewayGroup**](DnsGatewayGroupsAPI.md#RedeployDNSGatewayGroup) | **Post** /dns-gateway-groups/{group_id}:redeploy | 
[**RemoveDNSGatewayFromGroup**](DnsGatewayGroupsAPI.md#RemoveDNSGatewayFromGroup) | **Post** /dns-gateway-groups/{group_id}:remove-gateway | 



## AddDNSGatewayToGroup

> DNSGatewayGroupResp AddDNSGatewayToGroup(ctx, groupId).Body(body).Execute()





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
	groupId := int64(789) // int64 | dns gateway group id
	body := *openapiclient.NewDNSGatewayGroupAddGatewayReq() // DNSGatewayGroupAddGatewayReq | DNS gateway info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsGatewayGroupsAPI.AddDNSGatewayToGroup(context.Background(), groupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsGatewayGroupsAPI.AddDNSGatewayToGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDNSGatewayToGroup`: DNSGatewayGroupResp
	fmt.Fprintf(os.Stdout, "Response from `DnsGatewayGroupsAPI.AddDNSGatewayToGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | dns gateway group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDNSGatewayToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DNSGatewayGroupAddGatewayReq**](DNSGatewayGroupAddGatewayReq.md) | DNS gateway info | 

### Return type

[**DNSGatewayGroupResp**](DNSGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDNSGatewayGroup

> DNSGatewayGroupResp CreateDNSGatewayGroup(ctx).DnsGatewayGroup(dnsGatewayGroup).Execute()





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
	dnsGatewayGroup := *openapiclient.NewDNSGatewayGroupCreateReq(*openapiclient.NewDNSGatewayGroupCreateReqGroup([]int64{int64(123)}, "Name_example", "Origin_example", int64(123))) // DNSGatewayGroupCreateReq | DNS gateway group info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsGatewayGroupsAPI.CreateDNSGatewayGroup(context.Background()).DnsGatewayGroup(dnsGatewayGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsGatewayGroupsAPI.CreateDNSGatewayGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDNSGatewayGroup`: DNSGatewayGroupResp
	fmt.Fprintf(os.Stdout, "Response from `DnsGatewayGroupsAPI.CreateDNSGatewayGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDNSGatewayGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsGatewayGroup** | [**DNSGatewayGroupCreateReq**](DNSGatewayGroupCreateReq.md) | DNS gateway group info | 

### Return type

[**DNSGatewayGroupResp**](DNSGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDNSGatewayGroup

> DNSGatewayGroupResp DeleteDNSGatewayGroup(ctx, groupId).Force(force).Execute()





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
	groupId := int64(789) // int64 | dns gateway group id
	force := true // bool | force delete or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsGatewayGroupsAPI.DeleteDNSGatewayGroup(context.Background(), groupId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsGatewayGroupsAPI.DeleteDNSGatewayGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDNSGatewayGroup`: DNSGatewayGroupResp
	fmt.Fprintf(os.Stdout, "Response from `DnsGatewayGroupsAPI.DeleteDNSGatewayGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | dns gateway group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDNSGatewayGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 

### Return type

[**DNSGatewayGroupResp**](DNSGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDNSGatewayGroup

> DNSGatewayGroupResp GetDNSGatewayGroup(ctx, groupId).Execute()





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
	groupId := int64(789) // int64 | dns gateway group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsGatewayGroupsAPI.GetDNSGatewayGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsGatewayGroupsAPI.GetDNSGatewayGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDNSGatewayGroup`: DNSGatewayGroupResp
	fmt.Fprintf(os.Stdout, "Response from `DnsGatewayGroupsAPI.GetDNSGatewayGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | dns gateway group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDNSGatewayGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DNSGatewayGroupResp**](DNSGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDNSGatewayGroups

> DNSGatewayGroupsResp ListDNSGatewayGroups(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsGatewayGroupsAPI.ListDNSGatewayGroups(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsGatewayGroupsAPI.ListDNSGatewayGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDNSGatewayGroups`: DNSGatewayGroupsResp
	fmt.Fprintf(os.Stdout, "Response from `DnsGatewayGroupsAPI.ListDNSGatewayGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDNSGatewayGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 

### Return type

[**DNSGatewayGroupsResp**](DNSGatewayGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedeployDNSGatewayGroup

> DNSGatewayGroupResp RedeployDNSGatewayGroup(ctx, groupId).Execute()





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
	groupId := int64(789) // int64 | dns gateway group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsGatewayGroupsAPI.RedeployDNSGatewayGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsGatewayGroupsAPI.RedeployDNSGatewayGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedeployDNSGatewayGroup`: DNSGatewayGroupResp
	fmt.Fprintf(os.Stdout, "Response from `DnsGatewayGroupsAPI.RedeployDNSGatewayGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | dns gateway group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeployDNSGatewayGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DNSGatewayGroupResp**](DNSGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDNSGatewayFromGroup

> DNSGatewayGroupResp RemoveDNSGatewayFromGroup(ctx, groupId).Body(body).Force(force).Execute()





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
	groupId := int64(789) // int64 | dns gateway group id
	body := *openapiclient.NewDNSGatewayGroupRemoveGatewayReq() // DNSGatewayGroupRemoveGatewayReq | DNS gateway info
	force := true // bool | force delete or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsGatewayGroupsAPI.RemoveDNSGatewayFromGroup(context.Background(), groupId).Body(body).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsGatewayGroupsAPI.RemoveDNSGatewayFromGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveDNSGatewayFromGroup`: DNSGatewayGroupResp
	fmt.Fprintf(os.Stdout, "Response from `DnsGatewayGroupsAPI.RemoveDNSGatewayFromGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | dns gateway group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDNSGatewayFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DNSGatewayGroupRemoveGatewayReq**](DNSGatewayGroupRemoveGatewayReq.md) | DNS gateway info | 
 **force** | **bool** | force delete or not | 

### Return type

[**DNSGatewayGroupResp**](DNSGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

