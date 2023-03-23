// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airplane-api/pkg/models/shared"
	"net/http"
)

type GetRunQueryParams struct {
	// ID of the run to retrieve.
	ID string `queryParam:"style=form,explode=true,name=id"`
}

type GetRunRequest struct {
	QueryParams GetRunQueryParams
}

type GetRunResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ApiextGetRunResponse *shared.ApiextGetRunResponse
}