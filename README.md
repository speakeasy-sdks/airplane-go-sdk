# airplane-api

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/airplane-go-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "airplane-api"
    "airplane-api/pkg/models/shared"
    "airplane-api/pkg/models/operations"
)

func main() {
    s := airplane.New(
        airplane.WithSecurity(shared.Security{
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    req := shared.ApiextCancelPromptRequest{
        ID: "pmt20221122zyydx3rho2t",
    }

    ctx := context.Background()
    res, err := s.Prompts.Cancel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextCancelPromptResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations


### Prompts

* `Cancel` - Cancel Prompt
* `Get` - Get Prompt
* `List` - List Prompts
* `Submit` - Submit Prompt

### Runbooks

* `Execute` - Execute Runbook

### Runs

* `Cancel` - Cancel Run
* `Get` - Cancel Run
* `GetOutputs` - Get Run Outputs
* `List` - List Runs

### Sessions

* `Get` - Get Session
* `List` - List Sessions

### Tasks

* `Execute` - Execute Task
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
