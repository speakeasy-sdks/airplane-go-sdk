// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/shared"
	"net/http"
)

type GetOutputsRequest struct {
	// ID of the run to retrieve.
	ID string `queryParam:"style=form,explode=true,name=id"`
}

type GetOutputsResponse struct {
	ContentType string
	// OK
	GetOutputsResponse *shared.GetOutputsResponse
	StatusCode         int
	RawResponse        *http.Response
}
