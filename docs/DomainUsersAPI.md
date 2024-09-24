# \DomainUsersAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDomainUser**](DomainUsersAPI.md#DeleteDomainUser) | **Delete** /domain-users/{domain_user_id} | 
[**ListDomainUsers**](DomainUsersAPI.md#ListDomainUsers) | **Get** /domain-users/ | 



## DeleteDomainUser

> DeleteDomainUser(ctx, domainUserId).Execute()





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
	domainUserId := int64(789) // int64 | domain user id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DomainUsersAPI.DeleteDomainUser(context.Background(), domainUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainUsersAPI.DeleteDomainUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainUserId** | **int64** | domain user id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainUserRequest struct via the builder pattern


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


## ListDomainUsers

> DomainUsersResp ListDomainUsers(ctx).Limit(limit).Offset(offset).Execute()





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
	resp, r, err := apiClient.DomainUsersAPI.ListDomainUsers(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainUsersAPI.ListDomainUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDomainUsers`: DomainUsersResp
	fmt.Fprintf(os.Stdout, "Response from `DomainUsersAPI.ListDomainUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDomainUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 

### Return type

[**DomainUsersResp**](DomainUsersResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

