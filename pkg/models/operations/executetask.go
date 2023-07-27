// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/shared"
	"net/http"
)

type ExecuteTaskRequest struct {
	// ExecuteTaskRequest
	ExecuteTaskRequest shared.ExecuteTaskRequest `request:"mediaType=application/json"`
	// Environment to execute the task in.
	EnvSlug *string `queryParam:"style=form,explode=true,name=envSlug"`
}

func (o *ExecuteTaskRequest) GetExecuteTaskRequest() shared.ExecuteTaskRequest {
	if o == nil {
		return shared.ExecuteTaskRequest{}
	}
	return o.ExecuteTaskRequest
}

func (o *ExecuteTaskRequest) GetEnvSlug() *string {
	if o == nil {
		return nil
	}
	return o.EnvSlug
}

type ExecuteTaskResponse struct {
	ContentType string
	// OK
	ExecuteTaskResponse *shared.ExecuteTaskResponse
	StatusCode          int
	RawResponse         *http.Response
}

func (o *ExecuteTaskResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ExecuteTaskResponse) GetExecuteTaskResponse() *shared.ExecuteTaskResponse {
	if o == nil {
		return nil
	}
	return o.ExecuteTaskResponse
}

func (o *ExecuteTaskResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ExecuteTaskResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
