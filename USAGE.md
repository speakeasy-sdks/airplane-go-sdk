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

    req := operations.ListPromptsRequest{
        RunID: "unde",
    }

    ctx := context.Background()
    res, err := s.Prompts.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextListPromptsResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->