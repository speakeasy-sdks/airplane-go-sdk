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

	var runID string = "Bicycle"

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