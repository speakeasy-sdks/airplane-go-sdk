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
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: shared.SchemeAPIKeyAuth{
                APIKey: "YOUR_API_KEY_HERE",
            },
        }),
    )

    req := operations.CancelPromptRequest{
        Request: shared.ApiextCancelPromptRequest{
            ID: "pmt20221122zyydx3rho2t",
        },
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