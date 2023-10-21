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

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/airplane-go-sdk
```
<!-- End SDK Installation -->

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

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	airplanegosdk "github.com/speakeasy-sdks/airplane-go-sdk"
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := airplanegosdk.New(
		airplanegosdk.WithSecurity(""),
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
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
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
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
