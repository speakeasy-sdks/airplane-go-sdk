<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/airplane-go-sdk"
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/operations"
)

func main() {
    s := airplane.New(
        airplane.WithSecurity(shared.Security{
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Prompts.List(ctx, "corrupti")
    if err != nil {
        log.Fatal(err)
    }

    if res.ListPromptsResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->