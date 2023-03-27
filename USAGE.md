<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/airplane-go-sdk"
    "github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/operations"
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