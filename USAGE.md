<!-- Start SDK Example Usage -->


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