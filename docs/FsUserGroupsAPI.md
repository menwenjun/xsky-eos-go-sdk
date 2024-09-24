# \FsUserGroupsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFSUsers**](FsUserGroupsAPI.md#AddFSUsers) | **Put** /fs-user-groups/{fs_user_group_id}/fs-users | 
[**CreateFSUserGroup**](FsUserGroupsAPI.md#CreateFSUserGroup) | **Post** /fs-user-groups/ | 
[**DeleteFSUserGroup**](FsUserGroupsAPI.md#DeleteFSUserGroup) | **Delete** /fs-user-groups/{fs_user_group_id} | 
[**GetFSUserGroup**](FsUserGroupsAPI.md#GetFSUserGroup) | **Get** /fs-user-groups/{fs_user_group_id} | 
[**ListFSUserGroups**](FsUserGroupsAPI.md#ListFSUserGroups) | **Get** /fs-user-groups/ | 
[**RemoveFSUsers**](FsUserGroupsAPI.md#RemoveFSUsers) | **Delete** /fs-user-groups/{fs_user_group_id}/fs-users | 
[**UpdateFSUserGroup**](FsUserGroupsAPI.md#UpdateFSUserGroup) | **Patch** /fs-user-groups/{fs_user_group_id} | 



## AddFSUsers

> FSUserGroupResp AddFSUsers(ctx, fsUserGroupId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fsUserGroupId := int64(789) // int64 | user group id
	body := *openapiclient.NewFSUserGroupAddUsersReq(*openapiclient.NewFSUserGroupAddUsersReqUserGroup()) // FSUserGroupAddUsersReq | users info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsUserGroupsAPI.AddFSUsers(context.Background(), fsUserGroupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsUserGroupsAPI.AddFSUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFSUsers`: FSUserGroupResp
	fmt.Fprintf(os.Stdout, "Response from `FsUserGroupsAPI.AddFSUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsUserGroupId** | **int64** | user group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFSUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FSUserGroupAddUsersReq**](FSUserGroupAddUsersReq.md) | users info | 

### Return type

[**FSUserGroupResp**](FSUserGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFSUserGroup

> FSUserGroupResp CreateFSUserGroup(ctx).Body(body).ClusterId(clusterId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	body := *openapiclient.NewFSUserGroupCreateReq(*openapiclient.NewFSUserGroupCreateReqUserGroup("Name_example")) // FSUserGroupCreateReq | user group info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsUserGroupsAPI.CreateFSUserGroup(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsUserGroupsAPI.CreateFSUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFSUserGroup`: FSUserGroupResp
	fmt.Fprintf(os.Stdout, "Response from `FsUserGroupsAPI.CreateFSUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFSUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FSUserGroupCreateReq**](FSUserGroupCreateReq.md) | user group info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**FSUserGroupResp**](FSUserGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFSUserGroup

> DeleteFSUserGroup(ctx, fsUserGroupId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fsUserGroupId := int64(789) // int64 | user group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FsUserGroupsAPI.DeleteFSUserGroup(context.Background(), fsUserGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsUserGroupsAPI.DeleteFSUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsUserGroupId** | **int64** | user group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFSUserGroupRequest struct via the builder pattern


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


## GetFSUserGroup

> FSUserGroupResp GetFSUserGroup(ctx, fsUserGroupId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fsUserGroupId := int64(789) // int64 | user group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsUserGroupsAPI.GetFSUserGroup(context.Background(), fsUserGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsUserGroupsAPI.GetFSUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFSUserGroup`: FSUserGroupResp
	fmt.Fprintf(os.Stdout, "Response from `FsUserGroupsAPI.GetFSUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsUserGroupId** | **int64** | user group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFSUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FSUserGroupResp**](FSUserGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFSUserGroups

> FSUserGroupsResp ListFSUserGroups(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).Type_(type_).FsUserId(fsUserId).NotDfsHdfsId(notDfsHdfsId).NotDfsSmbShareId(notDfsSmbShareId).GroupId(groupId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)
	type_ := "type__example" // string | security type of file storage user group (optional)
	fsUserId := int64(789) // int64 | file storage user id (optional)
	notDfsHdfsId := int64(789) // int64 | id of dfs hdfs user groups not in (optional)
	notDfsSmbShareId := int64(789) // int64 | id of dfs smb share user groups not in (optional)
	groupId := int64(789) // int64 | group id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsUserGroupsAPI.ListFSUserGroups(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).Type_(type_).FsUserId(fsUserId).NotDfsHdfsId(notDfsHdfsId).NotDfsSmbShareId(notDfsSmbShareId).GroupId(groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsUserGroupsAPI.ListFSUserGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFSUserGroups`: FSUserGroupsResp
	fmt.Fprintf(os.Stdout, "Response from `FsUserGroupsAPI.ListFSUserGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFSUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **clusterId** | **string** | cluster id | 
 **type_** | **string** | security type of file storage user group | 
 **fsUserId** | **int64** | file storage user id | 
 **notDfsHdfsId** | **int64** | id of dfs hdfs user groups not in | 
 **notDfsSmbShareId** | **int64** | id of dfs smb share user groups not in | 
 **groupId** | **int64** | group id | 

### Return type

[**FSUserGroupsResp**](FSUserGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFSUsers

> FSUserGroupResp RemoveFSUsers(ctx, fsUserGroupId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fsUserGroupId := int64(789) // int64 | user group id
	body := *openapiclient.NewFSUserGroupRemoveUsersReq(*openapiclient.NewFSUserGroupRemoveUsersReqUserGroup()) // FSUserGroupRemoveUsersReq | users info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsUserGroupsAPI.RemoveFSUsers(context.Background(), fsUserGroupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsUserGroupsAPI.RemoveFSUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveFSUsers`: FSUserGroupResp
	fmt.Fprintf(os.Stdout, "Response from `FsUserGroupsAPI.RemoveFSUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsUserGroupId** | **int64** | user group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFSUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FSUserGroupRemoveUsersReq**](FSUserGroupRemoveUsersReq.md) | users info | 

### Return type

[**FSUserGroupResp**](FSUserGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFSUserGroup

> FSUserGroupResp UpdateFSUserGroup(ctx, fsUserGroupId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fsUserGroupId := int64(789) // int64 | user group id
	body := *openapiclient.NewFSUserGroupUpdateReq(*openapiclient.NewFSUserGroupUpdateReqUserGroup()) // FSUserGroupUpdateReq | user group info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsUserGroupsAPI.UpdateFSUserGroup(context.Background(), fsUserGroupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsUserGroupsAPI.UpdateFSUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFSUserGroup`: FSUserGroupResp
	fmt.Fprintf(os.Stdout, "Response from `FsUserGroupsAPI.UpdateFSUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsUserGroupId** | **int64** | user group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFSUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FSUserGroupUpdateReq**](FSUserGroupUpdateReq.md) | user group info | 

### Return type

[**FSUserGroupResp**](FSUserGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

