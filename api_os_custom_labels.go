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
	"strings"
)


// OsCustomLabelsAPIService OsCustomLabelsAPI service
type OsCustomLabelsAPIService service

type ApiGetOSCustomLabelRequest struct {
	ctx context.Context
	ApiService *OsCustomLabelsAPIService
	osCustomLabelId int64
}

func (r ApiGetOSCustomLabelRequest) Execute() (*OSCustomLabelResp, *http.Response, error) {
	return r.ApiService.GetOSCustomLabelExecute(r)
}

/*
GetOSCustomLabel Method for GetOSCustomLabel

Get object storage custom label

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param osCustomLabelId os custom label id
 @return ApiGetOSCustomLabelRequest
*/
func (a *OsCustomLabelsAPIService) GetOSCustomLabel(ctx context.Context, osCustomLabelId int64) ApiGetOSCustomLabelRequest {
	return ApiGetOSCustomLabelRequest{
		ApiService: a,
		ctx: ctx,
		osCustomLabelId: osCustomLabelId,
	}
}

// Execute executes the request
//  @return OSCustomLabelResp
func (a *OsCustomLabelsAPIService) GetOSCustomLabelExecute(r ApiGetOSCustomLabelRequest) (*OSCustomLabelResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OSCustomLabelResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsCustomLabelsAPIService.GetOSCustomLabel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/os-custom-labels/{os_custom_label_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"os_custom_label_id"+"}", url.PathEscape(parameterValueToString(r.osCustomLabelId, "osCustomLabelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListOSCustomLabelsRequest struct {
	ctx context.Context
	ApiService *OsCustomLabelsAPIService
	bucketId *int64
	category *string
	clusterId *string
	limit *int64
	offset *int64
}

// bucket id
func (r ApiListOSCustomLabelsRequest) BucketId(bucketId int64) ApiListOSCustomLabelsRequest {
	r.bucketId = &bucketId
	return r
}

// label category: meta, default, tagging
func (r ApiListOSCustomLabelsRequest) Category(category string) ApiListOSCustomLabelsRequest {
	r.category = &category
	return r
}

// cluster id
func (r ApiListOSCustomLabelsRequest) ClusterId(clusterId string) ApiListOSCustomLabelsRequest {
	r.clusterId = &clusterId
	return r
}

// paging param
func (r ApiListOSCustomLabelsRequest) Limit(limit int64) ApiListOSCustomLabelsRequest {
	r.limit = &limit
	return r
}

// paging param
func (r ApiListOSCustomLabelsRequest) Offset(offset int64) ApiListOSCustomLabelsRequest {
	r.offset = &offset
	return r
}

func (r ApiListOSCustomLabelsRequest) Execute() (*OSCustomLabelsResp, *http.Response, error) {
	return r.ApiService.ListOSCustomLabelsExecute(r)
}

/*
ListOSCustomLabels Method for ListOSCustomLabels

List object storage custom labels

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListOSCustomLabelsRequest
*/
func (a *OsCustomLabelsAPIService) ListOSCustomLabels(ctx context.Context) ApiListOSCustomLabelsRequest {
	return ApiListOSCustomLabelsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OSCustomLabelsResp
func (a *OsCustomLabelsAPIService) ListOSCustomLabelsExecute(r ApiListOSCustomLabelsRequest) (*OSCustomLabelsResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OSCustomLabelsResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsCustomLabelsAPIService.ListOSCustomLabels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/os-custom-labels/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.bucketId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "bucket_id", r.bucketId, "form", "")
	}
	if r.category != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category", r.category, "form", "")
	}
	if r.clusterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cluster_id", r.clusterId, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
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
