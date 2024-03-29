// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/airplane-go-sdk/v3/pkg/models/shared"
	"net/http"
)

type ListRunsRequest struct {
	// Number of results per call. Accepted values: 0 - 50. Default: 50.
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// The offset used for this page of results.
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Fetch runs last updated after this time (RFC 3339 format).
	Since *string `queryParam:"style=form,explode=true,name=since"`
	// Unique ID of the task to fetch runs for.
	TaskID *string `queryParam:"style=form,explode=true,name=taskID"`
	// Unique slug of the task to fetch runs for.
	TaskSlug *string `queryParam:"style=form,explode=true,name=taskSlug"`
	// Fetch runs last updated before this time (RFC 3339 format).
	Until *string `queryParam:"style=form,explode=true,name=until"`
}

func (o *ListRunsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListRunsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListRunsRequest) GetSince() *string {
	if o == nil {
		return nil
	}
	return o.Since
}

func (o *ListRunsRequest) GetTaskID() *string {
	if o == nil {
		return nil
	}
	return o.TaskID
}

func (o *ListRunsRequest) GetTaskSlug() *string {
	if o == nil {
		return nil
	}
	return o.TaskSlug
}

func (o *ListRunsRequest) GetUntil() *string {
	if o == nil {
		return nil
	}
	return o.Until
}

type ListRunsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	ListRunsResponse *shared.ListRunsResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListRunsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListRunsResponse) GetListRunsResponse() *shared.ListRunsResponse {
	if o == nil {
		return nil
	}
	return o.ListRunsResponse
}

func (o *ListRunsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListRunsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
