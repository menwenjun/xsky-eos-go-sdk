/*
XMS API

Testing DfsAuditLogsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_DfsAuditLogsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DfsAuditLogsAPIService ListDfsAuditLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DfsAuditLogsAPI.ListDfsAuditLogs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsAuditLogsAPIService UpdateDfsAuditLog", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dfsAuditLogId int64

		resp, httpRes, err := apiClient.DfsAuditLogsAPI.UpdateDfsAuditLog(context.Background(), dfsAuditLogId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
