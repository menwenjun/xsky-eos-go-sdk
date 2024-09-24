# \FsUsersAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFSUser**](FsUsersAPI.md#CreateFSUser) | **Post** /fs-users/ | 
[**DeleteFSUser**](FsUsersAPI.md#DeleteFSUser) | **Delete** /fs-users/{fs_user_id} | 
[**GetFSUser**](FsUsersAPI.md#GetFSUser) | **Get** /fs-users/{fs_user_id} | 
[**ListFSUsers**](FsUsersAPI.md#ListFSUsers) | **Get** /fs-users/ | 
[**UpdateFSUser**](FsUsersAPI.md#UpdateFSUser) | **Patch** /fs-users/{fs_user_id} | 
[**UpdateFSUserSecondaryGroups**](FsUsersAPI.md#UpdateFSUserSecondaryGroups) | **Patch** /fs-users/{fs_user_id}:update-secondary-groups | 
[**VerifyFSUser**](FsUsersAPI.md#VerifyFSUser) | **Post** /fs-users/:verify | 



## CreateFSUser

> FSUserResp CreateFSUser(ctx).Body(body).ClusterId(clusterId).AllowPathCreate(allowPathCreate).Execute()





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
	body := *openapiclient.NewFSUserCreateReq(*openapiclient.NewFSUserCreateReqUser()) // FSUserCreateReq | user info
	clusterId := "clusterId_example" // string | cluster id (optional)
	allowPathCreate := true // bool | allow create path for s3 when not existed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsUsersAPI.CreateFSUser(context.Background()).Body(body).ClusterId(clusterId).AllowPathCreate(allowPathCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsUsersAPI.CreateFSUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFSUser`: FSUserResp
	fmt.Fprintf(os.Stdout, "Response from `FsUsersAPI.CreateFSUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFSUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FSUserCreateReq**](FSUserCreateReq.md) | user info | 
 **clusterId** | **string** | cluster id | 
 **allowPathCreate** | **bool** | allow create path for s3 when not existed | 

### Return type

[**FSUserResp**](FSUserResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFSUser

> DeleteFSUser(ctx, fsUserId).Execute()





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
	fsUserId := int64(789) // int64 | user id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FsUsersAPI.DeleteFSUser(context.Background(), fsUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsUsersAPI.DeleteFSUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsUserId** | **int64** | user id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFSUserRequest struct via the builder pattern


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


## GetFSUser

> FSUserResp GetFSUser(ctx, fsUserId).Execute()





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
	fsUserId := int64(789) // int64 | user id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsUsersAPI.GetFSUser(context.Background(), fsUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsUsersAPI.GetFSUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFSUser`: FSUserResp
	fmt.Fprintf(os.Stdout, "Response from `FsUsersAPI.GetFSUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsUserId** | **int64** | user id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFSUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FSUserResp**](FSUserResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFSUsers

> FSUsersResp ListFSUsers(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).Security(security).FsUserGroupId(fsUserGroupId).NotFsUserGroupId(notFsUserGroupId).NotDfsSmbShareId(notDfsSmbShareId).NotDfsHdfsId(notDfsHdfsId).NotDfsFtpShare(notDfsFtpShare).DfsGatewayGroupId(dfsGatewayGroupId).S3Enabled(s3Enabled).UserId(userId).Execute()





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
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)
	security := "security_example" // string | security of file storage users (optional)
	fsUserGroupId := int64(789) // int64 | file storage user group id (optional)
	notFsUserGroupId := int64(789) // int64 | file storage user group id (optional)
	notDfsSmbShareId := int64(789) // int64 | id of dfs smb share users not in (optional)
	notDfsHdfsId := int64(789) // int64 | id of dfs hdfs users not in (optional)
	notDfsFtpShare := true // bool | value must be true, means available add to ftp share (optional)
	dfsGatewayGroupId := int64(789) // int64 | dfs gateway group id, used with not_dfs_ftp_share (optional)
	s3Enabled := true // bool | is s3 enabled (optional)
	userId := int64(789) // int64 | user id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsUsersAPI.ListFSUsers(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).Security(security).FsUserGroupId(fsUserGroupId).NotFsUserGroupId(notFsUserGroupId).NotDfsSmbShareId(notDfsSmbShareId).NotDfsHdfsId(notDfsHdfsId).NotDfsFtpShare(notDfsFtpShare).DfsGatewayGroupId(dfsGatewayGroupId).S3Enabled(s3Enabled).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsUsersAPI.ListFSUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFSUsers`: FSUsersResp
	fmt.Fprintf(os.Stdout, "Response from `FsUsersAPI.ListFSUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFSUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **clusterId** | **string** | cluster id | 
 **security** | **string** | security of file storage users | 
 **fsUserGroupId** | **int64** | file storage user group id | 
 **notFsUserGroupId** | **int64** | file storage user group id | 
 **notDfsSmbShareId** | **int64** | id of dfs smb share users not in | 
 **notDfsHdfsId** | **int64** | id of dfs hdfs users not in | 
 **notDfsFtpShare** | **bool** | value must be true, means available add to ftp share | 
 **dfsGatewayGroupId** | **int64** | dfs gateway group id, used with not_dfs_ftp_share | 
 **s3Enabled** | **bool** | is s3 enabled | 
 **userId** | **int64** | user id | 

### Return type

[**FSUsersResp**](FSUsersResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFSUser

> FSUserResp UpdateFSUser(ctx, fsUserId).Body(body).AllowPathCreate(allowPathCreate).Execute()





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
	fsUserId := int64(789) // int64 | user id
	body := *openapiclient.NewFSUserUpdateReq(*openapiclient.NewFSUserUpdateReqUser()) // FSUserUpdateReq | user info
	allowPathCreate := true // bool | allow create path for s3 when not existed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsUsersAPI.UpdateFSUser(context.Background(), fsUserId).Body(body).AllowPathCreate(allowPathCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsUsersAPI.UpdateFSUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFSUser`: FSUserResp
	fmt.Fprintf(os.Stdout, "Response from `FsUsersAPI.UpdateFSUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsUserId** | **int64** | user id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFSUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FSUserUpdateReq**](FSUserUpdateReq.md) | user info | 
 **allowPathCreate** | **bool** | allow create path for s3 when not existed | 

### Return type

[**FSUserResp**](FSUserResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFSUserSecondaryGroups

> FSUserResp UpdateFSUserSecondaryGroups(ctx, fsUserId).Body(body).Execute()





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
	fsUserId := int64(789) // int64 | user id
	body := *openapiclient.NewFSUserUpdateSecondaryGroupsReq(*openapiclient.NewFSUserUpdateSecondaryGroupsReqUser()) // FSUserUpdateSecondaryGroupsReq | user info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsUsersAPI.UpdateFSUserSecondaryGroups(context.Background(), fsUserId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsUsersAPI.UpdateFSUserSecondaryGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFSUserSecondaryGroups`: FSUserResp
	fmt.Fprintf(os.Stdout, "Response from `FsUsersAPI.UpdateFSUserSecondaryGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsUserId** | **int64** | user id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFSUserSecondaryGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FSUserUpdateSecondaryGroupsReq**](FSUserUpdateSecondaryGroupsReq.md) | user info | 

### Return type

[**FSUserResp**](FSUserResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyFSUser

> FSUserResp VerifyFSUser(ctx).Body(body).Execute()





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
	body := *openapiclient.NewFSUserVerifyReq(*openapiclient.NewFSUserVerifyReqUser("Name_example")) // FSUserVerifyReq | the identity credential

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsUsersAPI.VerifyFSUser(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsUsersAPI.VerifyFSUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyFSUser`: FSUserResp
	fmt.Fprintf(os.Stdout, "Response from `FsUsersAPI.VerifyFSUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyFSUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FSUserVerifyReq**](FSUserVerifyReq.md) | the identity credential | 

### Return type

[**FSUserResp**](FSUserResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

