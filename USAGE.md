<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	airplanegosdk "github.com/speakeasy-sdks/airplane-go-sdk/v3"
	"github.com/speakeasy-sdks/airplane-go-sdk/v3/pkg/models/shared"
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