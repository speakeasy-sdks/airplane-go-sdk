// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package airplanegosdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/airplane-go-sdk/v3/pkg/models/shared"
	"github.com/speakeasy-sdks/airplane-go-sdk/v3/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Airplane API
	"https://api.airplane.dev/v0",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Airplane API: Airplane API
type Airplane struct {
	// A prompt is used to gather user input during a task's execution. See Prompts to see how prompts are used.
	Prompts *Prompts
	// A Runbook is a multi-step, human-in-the-loop workflow. Runbooks are able to take a set of top-level parameters, run one or more functions, and generate output at each step of the way.
	Runbooks *Runbooks
	// A run represents an instance of a task's execution. See Tasks API for how to execute tasks.
	Runs *Runs
	// A session represents an instance of a runbook's execution. See Runbooks API for how to execute runbooks.
	Sessions *Sessions
	// A Task is a lightweight app that represents a single business operation for people at your company to execute.
	Tasks *Tasks

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Airplane)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Airplane) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Airplane) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Airplane) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Airplane) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(apiKeyAuth string) SDKOption {
	return func(sdk *Airplane) {
		security := shared.Security{APIKeyAuth: apiKeyAuth}
		sdk.sdkConfiguration.Security = withSecurity(&security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *Airplane) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *Airplane) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Airplane {
	sdk := &Airplane{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "0.0.1",
			SDKVersion:        "3.0.1",
			GenVersion:        "2.250.2",
			UserAgent:         "speakeasy-sdk/go 3.0.1 2.250.2 0.0.1 github.com/speakeasy-sdks/airplane-go-sdk",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Prompts = newPrompts(sdk.sdkConfiguration)

	sdk.Runbooks = newRunbooks(sdk.sdkConfiguration)

	sdk.Runs = newRuns(sdk.sdkConfiguration)

	sdk.Sessions = newSessions(sdk.sdkConfiguration)

	sdk.Tasks = newTasks(sdk.sdkConfiguration)

	return sdk
}
