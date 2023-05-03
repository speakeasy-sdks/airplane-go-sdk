# Runbooks

## Overview

A Runbook is a multi-step, human-in-the-loop workflow. Runbooks are able to take a set of top-level parameters, run one or more functions, and generate output at each step of the way.

### Available Operations

* [Execute](#execute) - Execute Runbook

## Execute

Execute a runbook and receive a session ID to track the runbook's execution.
Check on the status of your newly created session with [/sessions/get](/api/sessions#sessions-get).

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/airplane-go-sdk"
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/shared"
)

func main() {
    s := airplane.New(
        airplane.WithSecurity(shared.Security{
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Runbooks.Execute(ctx, shared.ExecuteRunbookRequest{
        ID: "rbk20220120z15kl79",
        ParamValues: map[string]string{
            "magnam": "debitis",
            "ipsa": "delectus",
        },
        Slug: airplane.String("hello_world"),
    }, "tempora")
    if err != nil {
        log.Fatal(err)
    }

    if res.ExecuteRunbookResponse != nil {
        // handle response
    }
}
```
