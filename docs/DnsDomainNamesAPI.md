# \DnsDomainNamesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDNSDomainName**](DnsDomainNamesAPI.md#CreateDNSDomainName) | **Post** /dns-domain-names/ | 
[**DeleteDNSDomainName**](DnsDomainNamesAPI.md#DeleteDNSDomainName) | **Delete** /dns-domain-names/{name_id} | 
[**GetDNSDomainName**](DnsDomainNamesAPI.md#GetDNSDomainName) | **Get** /dns-domain-names/{name_id} | 
[**ListDNSDomainNames**](DnsDomainNamesAPI.md#ListDNSDomainNames) | **Get** /dns-domain-names/ | 
[**UpdateDNSDomainName**](DnsDomainNamesAPI.md#UpdateDNSDomainName) | **Patch** /dns-domain-names/{name_id} | 



## CreateDNSDomainName

> DNSDomainNameResp CreateDNSDomainName(ctx).DnsDomainName(dnsDomainName).Execute()





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
	dnsDomainName := *openapiclient.NewDNSDomainNameCreateReq(*openapiclient.NewDNSDomainNameCreateReqDomainName(int64(123), int64(123), "Name_example", int64(123), "ResourceType_example")) // DNSDomainNameCreateReq | DNS domain name info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsDomainNamesAPI.CreateDNSDomainName(context.Background()).DnsDomainName(dnsDomainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsDomainNamesAPI.CreateDNSDomainName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDNSDomainName`: DNSDomainNameResp
	fmt.Fprintf(os.Stdout, "Response from `DnsDomainNamesAPI.CreateDNSDomainName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDNSDomainNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsDomainName** | [**DNSDomainNameCreateReq**](DNSDomainNameCreateReq.md) | DNS domain name info | 

### Return type

[**DNSDomainNameResp**](DNSDomainNameResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDNSDomainName

> DeleteDNSDomainName(ctx, nameId).Execute()





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
	nameId := int64(789) // int64 | DNS domain name id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsDomainNamesAPI.DeleteDNSDomainName(context.Background(), nameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsDomainNamesAPI.DeleteDNSDomainName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameId** | **int64** | DNS domain name id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDNSDomainNameRequest struct via the builder pattern


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


## GetDNSDomainName

> DNSDomainNameResp GetDNSDomainName(ctx, nameId).Execute()





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
	nameId := int64(789) // int64 | dns domain name id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsDomainNamesAPI.GetDNSDomainName(context.Background(), nameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsDomainNamesAPI.GetDNSDomainName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDNSDomainName`: DNSDomainNameResp
	fmt.Fprintf(os.Stdout, "Response from `DnsDomainNamesAPI.GetDNSDomainName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameId** | **int64** | dns domain name id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDNSDomainNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DNSDomainNameResp**](DNSDomainNameResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDNSDomainNames

> DNSDomainNamesResp ListDNSDomainNames(ctx).Limit(limit).Offset(offset).DnsZoneId(dnsZoneId).DfsGatewayZoneId(dfsGatewayZoneId).Execute()





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
	dnsZoneId := int64(789) // int64 | dns zone id (optional)
	dfsGatewayZoneId := int64(789) // int64 | dfs gateway zone id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsDomainNamesAPI.ListDNSDomainNames(context.Background()).Limit(limit).Offset(offset).DnsZoneId(dnsZoneId).DfsGatewayZoneId(dfsGatewayZoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsDomainNamesAPI.ListDNSDomainNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDNSDomainNames`: DNSDomainNamesResp
	fmt.Fprintf(os.Stdout, "Response from `DnsDomainNamesAPI.ListDNSDomainNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDNSDomainNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **dnsZoneId** | **int64** | dns zone id | 
 **dfsGatewayZoneId** | **int64** | dfs gateway zone id | 

### Return type

[**DNSDomainNamesResp**](DNSDomainNamesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDNSDomainName

> DNSDomainNameResp UpdateDNSDomainName(ctx, nameId).DnsDomainName(dnsDomainName).Execute()





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
	nameId := int64(789) // int64 | DNS domain name id
	dnsDomainName := *openapiclient.NewDNSDomainNameUpdateReq(*openapiclient.NewDNSDomainNameUpdateReqDomainName(int64(123))) // DNSDomainNameUpdateReq | DNS domain name info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsDomainNamesAPI.UpdateDNSDomainName(context.Background(), nameId).DnsDomainName(dnsDomainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsDomainNamesAPI.UpdateDNSDomainName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDNSDomainName`: DNSDomainNameResp
	fmt.Fprintf(os.Stdout, "Response from `DnsDomainNamesAPI.UpdateDNSDomainName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameId** | **int64** | DNS domain name id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDNSDomainNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsDomainName** | [**DNSDomainNameUpdateReq**](DNSDomainNameUpdateReq.md) | DNS domain name info | 

### Return type

[**DNSDomainNameResp**](DNSDomainNameResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

