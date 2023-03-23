// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airplane-api/pkg/models/shared"
	"net/http"
)

type GetOutputsQueryParams struct {
	// ID of the run to retrieve.
	ID string `queryParam:"style=form,explode=true,name=id"`
}

type GetOutputsRequest struct {
	QueryParams GetOutputsQueryParams
}

type GetOutputsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ApiextGetOutputsResponse *shared.ApiextGetOutputsResponse
}