# Prompts

## Overview

A prompt is used to gather user input during a task's execution. See Prompts to see how prompts are used.

### Available Operations

* [Cancel](#cancel) - Cancel Prompt
* [Get](#get) - Get Prompt
* [List](#list) - List Prompts
* [Submit](#submit) - Submit Prompt

## Cancel

Cancel a prompt.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/airplane-go-sdk"
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/shared"
)

func main() {
    s := airplane.New(
        airplane.WithSecurity(shared.Security{
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Prompts.Cancel(ctx, shared.CancelPromptRequest{
        ID: airplane.String("pmt20221122zyydx3rho2t"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CancelPromptResponse != nil {
        // handle response
    }
}
```

## Get

Get information about an existing prompt.

### Example Usage

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
    res, err := s.Prompts.Get(ctx, "provident")
    if err != nil {
        log.Fatal(err)
    }

    if res.GetPromptResponse != nil {
        // handle response
    }
}
```

## List

List prompts from existing runs.

### Example Usage

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
    res, err := s.Prompts.List(ctx, "distinctio")
    if err != nil {
        log.Fatal(err)
    }

    if res.ListPromptsResponse != nil {
        // handle response
    }
}
```

## Submit

Submit a prompt with a set of parameter values.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/airplane-go-sdk"
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/shared"
)

func main() {
    s := airplane.New(
        airplane.WithSecurity(shared.Security{
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Prompts.Submit(ctx, shared.SubmitPromptRequest{
        ID: airplane.String("pmt20221122zyydx3rho2t"),
        Values: map[string]string{
            "unde": "nulla",
            "corrupti": "illum",
            "vel": "error",
            "deserunt": "suscipit",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SubmitPromptResponse != nil {
        // handle response
    }
}
```
