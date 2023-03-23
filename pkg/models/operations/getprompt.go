// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airplane-api/pkg/models/shared"
	"net/http"
)

type GetPromptRequest struct {
	// ID of the prompt to retrieve.
	ID string `queryParam:"style=form,explode=true,name=id"`
}

type GetPromptResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ApiextGetPromptResponse *shared.ApiextGetPromptResponse
}
