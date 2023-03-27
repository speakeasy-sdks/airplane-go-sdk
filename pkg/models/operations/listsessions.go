// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/shared"
	"net/http"
)

type ListSessionsRequest struct {
	// Number of results per call. Accepted values: 0 - 50. The default: 50.
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// The offset used for this page of results.
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Only fetch sessions for this specific runbook.
	RunbookID *string `queryParam:"style=form,explode=true,name=runbookID"`
	// Fetch sessions last updated after this time (RFC 3339 format).
	UpdatedAfter *string `queryParam:"style=form,explode=true,name=updatedAfter"`
	// Fetch sessions last updated before this time (RFC 3339 format).
	UpdatedBefore *string `queryParam:"style=form,explode=true,name=updatedBefore"`
}

type ListSessionsResponse struct {
	ContentType string
	// OK
	ListSessionResponse *shared.ListSessionResponse
	StatusCode          int
	RawResponse         *http.Response
}
