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
            APIKeyAuth: "",
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

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [shared.CancelRunRequest](../../models/shared/cancelrunrequest.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[*operations.CancelRunResponse](../../models/operations/cancelrunresponse.md), error**


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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Runs.Get(ctx, "suscipit")
    if err != nil {
        log.Fatal(err)
    }

    if res.GetRunResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the run to retrieve.                            |


### Response

**[*operations.GetRunResponse](../../models/operations/getrunresponse.md), error**


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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Runs.GetOutputs(ctx, "molestiae")
    if err != nil {
        log.Fatal(err)
    }

    if res.GetOutputsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the run to retrieve.                            |


### Response

**[*operations.GetOutputsResponse](../../models/operations/getoutputsresponse.md), error**


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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Runs.List(ctx, 791725, 812169, "voluptatum", "iusto", "excepturi", "nisi")
    if err != nil {
        log.Fatal(err)
    }

    if res.ListRunsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |
| `limit`                                                           | **int64*                                                          | :heavy_minus_sign:                                                | Number of results per call. Accepted values: 0 - 50. Default: 50. |
| `page`                                                            | **int64*                                                          | :heavy_minus_sign:                                                | The offset used for this page of results.                         |
| `since`                                                           | **string*                                                         | :heavy_minus_sign:                                                | Fetch runs last updated after this time (RFC 3339 format).        |
| `taskID`                                                          | **string*                                                         | :heavy_minus_sign:                                                | Unique ID of the task to fetch runs for.                          |
| `taskSlug`                                                        | **string*                                                         | :heavy_minus_sign:                                                | Unique slug of the task to fetch runs for.                        |
| `until`                                                           | **string*                                                         | :heavy_minus_sign:                                                | Fetch runs last updated before this time (RFC 3339 format).       |


### Response

**[*operations.ListRunsResponse](../../models/operations/listrunsresponse.md), error**

