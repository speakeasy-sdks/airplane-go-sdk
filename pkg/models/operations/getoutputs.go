// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/airplane-go-sdk/v3/pkg/models/shared"
	"net/http"
)

type GetOutputsRequest struct {
	// ID of the run to retrieve.
	ID string `queryParam:"style=form,explode=true,name=id"`
}

func (o *GetOutputsRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetOutputsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	GetOutputsResponse *shared.GetOutputsResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetOutputsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetOutputsResponse) GetGetOutputsResponse() *shared.GetOutputsResponse {
	if o == nil {
		return nil
	}
	return o.GetOutputsResponse
}

func (o *GetOutputsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetOutputsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
