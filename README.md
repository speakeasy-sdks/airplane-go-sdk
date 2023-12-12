<div align="center">
    <picture>
        <source srcset="https://user-images.githubusercontent.com/6267663/227311185-62d5759f-743c-488b-8b97-09eae1dac881.png" media="(prefers-color-scheme: dark)" width="100">
        <img src="https://user-images.githubusercontent.com/6267663/227311185-62d5759f-743c-488b-8b97-09eae1dac881.png" width="100">
    </picture>
    <h1>Airplane Go SDK</h1>
   <p>Developer infrastructure for internal tools</p>
   <a href="https://docs.airplane.dev/"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=5444e4&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/airplane-go-sdk/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/airplane-go-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/speakeasy-sdks/airplane-go-sdk/releases"><img src="https://img.shields.io/github/v/release/speakeasy-sdks/airplane-go-sdk?sort=semver&style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/airplane-go-sdk
```
<!-- End SDK Installation [installation] -->

## Authentication

The Airplane API uses API keys to authenticate your requests. API keys are prefixed with `tkn_`. Authentication to the API is performed by supplying the `X-Airplane-API-Key` header. Runs and sessions created through the API are recorded as being run by `API user`.
You can view and manage your API keys with the [Airplane CLI](https://docs.airplane.dev/platform/airplane-cli). To generate a new token:

```bash
airplane apikeys create <token name>
```

API keys can be created by `Team admins` and `Team developers`. See [Team Roles](https://docs.airplane.dev/platform/team-roles) for details.

Your API keys carry the privileges to execute every task and runbook your team owns so be sure to keep them secure! Do not share your API keys in publicly accessible areas such as GitHub, client-side code, etc.

All API requests must be made over HTTPS. Any requests made over plain HTTP will fail. Requests without a valid API key will fail and return a `401` error.

```bash
# Authenticated request example
curl https://api.airplane.dev/v0/runs/get \
    -H 'X-Airplane-API-Key: tkn_examplekey123456789EXAMPLEKEY1234567' \
    -d id=run20220222example \
    -G
```

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	airplanegosdk "github.com/speakeasy-sdks/airplane-go-sdk/v2"
	"github.com/speakeasy-sdks/airplane-go-sdk/v2/pkg/models/shared"
	"log"
)

func main() {
	s := airplanegosdk.New(
		airplanegosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	var runID string = "string"

	ctx := context.Background()
	res, err := s.Prompts.List(ctx, runID)
	if err != nil {
		log.Fatal(err)
	}

	if res.ListPromptsResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Prompts](docs/sdks/prompts/README.md)

* [Cancel](docs/sdks/prompts/README.md#cancel) - Cancel Prompt
* [Get](docs/sdks/prompts/README.md#get) - Get Prompt
* [List](docs/sdks/prompts/README.md#list) - List Prompts
* [Submit](docs/sdks/prompts/README.md#submit) - Submit Prompt

### [Runbooks](docs/sdks/runbooks/README.md)

* [Execute](docs/sdks/runbooks/README.md#execute) - Execute Runbook

### [Runs](docs/sdks/runs/README.md)

* [Get](docs/sdks/runs/README.md#get) - Cancel Run
* [GetOutputs](docs/sdks/runs/README.md#getoutputs) - Get Run Outputs
* [List](docs/sdks/runs/README.md#list) - List Runs

### [Sessions](docs/sdks/sessions/README.md)

* [Get](docs/sdks/sessions/README.md#get) - Get Session
* [List](docs/sdks/sessions/README.md#list) - List Sessions

### [Tasks](docs/sdks/tasks/README.md)

* [Execute](docs/sdks/tasks/README.md#execute) - Execute Task
<!-- End Available Resources and Operations [operations] -->







<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	airplanegosdk "github.com/speakeasy-sdks/airplane-go-sdk/v2"
	"github.com/speakeasy-sdks/airplane-go-sdk/v2/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/airplane-go-sdk/v2/pkg/models/shared"
	"log"
)

func main() {
	s := airplanegosdk.New(
		airplanegosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Prompts.Cancel(ctx, shared.CancelPromptRequest{
		ID: airplanegosdk.String("pmt20221122zyydx3rho2t"),
	})
	if err != nil {

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.airplane.dev/v0` | None |

#### Example

```go
package main

import (
	"context"
	airplanegosdk "github.com/speakeasy-sdks/airplane-go-sdk/v2"
	"github.com/speakeasy-sdks/airplane-go-sdk/v2/pkg/models/shared"
	"log"
)

func main() {
	s := airplanegosdk.New(
		airplanegosdk.WithServerIndex(0),
		airplanegosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Prompts.Cancel(ctx, shared.CancelPromptRequest{
		ID: airplanegosdk.String("pmt20221122zyydx3rho2t"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CancelPromptResponse != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	airplanegosdk "github.com/speakeasy-sdks/airplane-go-sdk/v2"
	"github.com/speakeasy-sdks/airplane-go-sdk/v2/pkg/models/shared"
	"log"
)

func main() {
	s := airplanegosdk.New(
		airplanegosdk.WithServerURL("https://api.airplane.dev/v0"),
		airplanegosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Prompts.Cancel(ctx, shared.CancelPromptRequest{
		ID: airplanegosdk.String("pmt20221122zyydx3rho2t"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CancelPromptResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `APIKeyAuth` | apiKey       | API key      |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	airplanegosdk "github.com/speakeasy-sdks/airplane-go-sdk/v2"
	"github.com/speakeasy-sdks/airplane-go-sdk/v2/pkg/models/shared"
	"log"
)

func main() {
	s := airplanegosdk.New(
		airplanegosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Prompts.Cancel(ctx, shared.CancelPromptRequest{
		ID: airplanegosdk.String("pmt20221122zyydx3rho2t"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CancelPromptResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
