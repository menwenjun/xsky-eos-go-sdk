/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// DfsHdfsIpWhiteListsAPIService DfsHdfsIpWhiteListsAPI service
type DfsHdfsIpWhiteListsAPIService service

type ApiListDfsHdfsIPWhiteListRequest struct {
	ctx context.Context
	ApiService *DfsHdfsIpWhiteListsAPIService
	limit *int64
	offset *int64
	clusterId *string
	dfsHdfsId *int64
	permission *string
	q *string
	sort *string
}

// paging param
func (r ApiListDfsHdfsIPWhiteListRequest) Limit(limit int64) ApiListDfsHdfsIPWhiteListRequest {
	r.limit = &limit
	return r
}

// paging param
func (r ApiListDfsHdfsIPWhiteListRequest) Offset(offset int64) ApiListDfsHdfsIPWhiteListRequest {
	r.offset = &offset
	return r
}

// cluster id
func (r ApiListDfsHdfsIPWhiteListRequest) ClusterId(clusterId string) ApiListDfsHdfsIPWhiteListRequest {
	r.clusterId = &clusterId
	return r
}

// dfs hdfs id
func (r ApiListDfsHdfsIPWhiteListRequest) DfsHdfsId(dfsHdfsId int64) ApiListDfsHdfsIPWhiteListRequest {
	r.dfsHdfsId = &dfsHdfsId
	return r
}

// permission of ip white list:RW, RO
func (r ApiListDfsHdfsIPWhiteListRequest) Permission(permission string) ApiListDfsHdfsIPWhiteListRequest {
	r.permission = &permission
	return r
}

// query param of search
func (r ApiListDfsHdfsIPWhiteListRequest) Q(q string) ApiListDfsHdfsIPWhiteListRequest {
	r.q = &q
	return r
}

// sort param of search
func (r ApiListDfsHdfsIPWhiteListRequest) Sort(sort string) ApiListDfsHdfsIPWhiteListRequest {
	r.sort = &sort
	return r
}

func (r ApiListDfsHdfsIPWhiteListRequest) Execute() (*DfsHdfsIPWhiteListsResp, *http.Response, error) {
	return r.ApiService.ListDfsHdfsIPWhiteListExecute(r)
}

/*
ListDfsHdfsIPWhiteList Method for ListDfsHdfsIPWhiteList

List dfs hdfs ip white list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDfsHdfsIPWhiteListRequest
*/
func (a *DfsHdfsIpWhiteListsAPIService) ListDfsHdfsIPWhiteList(ctx context.Context) ApiListDfsHdfsIPWhiteListRequest {
	return ApiListDfsHdfsIPWhiteListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DfsHdfsIPWhiteListsResp
func (a *DfsHdfsIpWhiteListsAPIService) ListDfsHdfsIPWhiteListExecute(r ApiListDfsHdfsIPWhiteListRequest) (*DfsHdfsIPWhiteListsResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsHdfsIPWhiteListsResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsHdfsIpWhiteListsAPIService.ListDfsHdfsIPWhiteList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-hdfs-ip-white-lists/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.clusterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cluster_id", r.clusterId, "form", "")
	}
	if r.dfsHdfsId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dfs_hdfs_id", r.dfsHdfsId, "form", "")
	}
	if r.permission != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "permission", r.permission, "form", "")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenInHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Xms-Auth-Token"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenInQuery"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
