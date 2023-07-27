// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package airplane

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/utils"
	"io"
	"net/http"
	"strings"
)

// tasks - A Task is a lightweight app that represents a single business operation for people at your company to execute.
type tasks struct {
	sdkConfiguration sdkConfiguration
}

func newTasks(sdkConfig sdkConfiguration) *tasks {
	return &tasks{
		sdkConfiguration: sdkConfig,
	}
}

// Execute - Execute Task
// Execute a task with a set of parameter values and receive a run ID to track the task's execution.
// Check on the status of your newly created run with [/runs/get](/api/runs#runs-get).
func (s *tasks) Execute(ctx context.Context, executeTaskRequest shared.ExecuteTaskRequest, envSlug *string) (*operations.ExecuteTaskResponse, error) {
	request := operations.ExecuteTaskRequest{
		ExecuteTaskRequest: executeTaskRequest,
		EnvSlug:            envSlug,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/tasks/execute"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "ExecuteTaskRequest", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ExecuteTaskResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ExecuteTaskResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.ExecuteTaskResponse = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}
