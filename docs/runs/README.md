# Runs

## Overview

A run represents an instance of a task's execution. See Tasks API for how to execute tasks.

### Available Operations

* [Cancel](#cancel) - Cancel Run
* [Get](#get) - Cancel Run
* [GetOutputs](#getoutputs) - Get Run Outputs
* [List](#list) - List Runs

## Cancel

Cancel a run.
Check on the status of your run with [/runs/get](/api/runs#runs-get).

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
    res, err := s.Runs.Cancel(ctx, shared.CancelRunRequest{
        RunID: airplane.String("run20220111zlq2ig4"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## Get

Get information about an existing run.

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
    res, err := s.Runs.Get(ctx, operations.GetRunRequest{
        ID: "1ba928fc-8167-442c-b739-205929396fea",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetRunResponse != nil {
        // handle response
    }
}
```

## GetOutputs

Get outputs from an existing run.

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
    res, err := s.Runs.GetOutputs(ctx, operations.GetOutputsRequest{
        ID: "7596eb10-faaa-4235-ac59-55907aff1a3a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetOutputsResponse != nil {
        // handle response
    }
}
```

## List

List Runs

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
    res, err := s.Runs.List(ctx, operations.ListRunsRequest{
        Limit: airplane.Int64(161309),
        Page: airplane.Int64(995300),
        Since: airplane.String("mollitia"),
        TaskID: airplane.String("occaecati"),
        TaskSlug: airplane.String("numquam"),
        Until: airplane.String("commodi"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListRunsResponse != nil {
        // handle response
    }
}
```
