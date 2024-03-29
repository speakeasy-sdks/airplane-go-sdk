// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package airplanegosdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/airplane-go-sdk/v3/internal/hooks"
	"github.com/speakeasy-sdks/airplane-go-sdk/v3/pkg/models/operations"
	"github.com/speakeasy-sdks/airplane-go-sdk/v3/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/airplane-go-sdk/v3/pkg/models/shared"
	"github.com/speakeasy-sdks/airplane-go-sdk/v3/pkg/utils"
	"io"
	"net/http"
	"net/url"
)

// Tasks - A Task is a lightweight app that represents a single business operation for people at your company to execute.
type Tasks struct {
	sdkConfiguration sdkConfiguration
}

func newTasks(sdkConfig sdkConfiguration) *Tasks {
	return &Tasks{
		sdkConfiguration: sdkConfig,
	}
}

// Execute Task
// Execute a task with a set of parameter values and receive a run ID to track the task's execution.
// Check on the status of your newly created run with [/runs/get](/api/runs#runs-get).
func (s *Tasks) Execute(ctx context.Context, executeTaskRequest shared.ExecuteTaskRequest, envSlug *string) (*operations.ExecuteTaskResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "executeTask",
		SecuritySource: s.sdkConfiguration.Security,
	}

	request := operations.ExecuteTaskRequest{
		ExecuteTaskRequest: executeTaskRequest,
		EnvSlug:            envSlug,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := url.JoinPath(baseURL, "/tasks/execute")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "ExecuteTaskRequest", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.ExecuteTaskResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out shared.ExecuteTaskResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ExecuteTaskResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}
