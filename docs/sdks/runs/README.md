# Runs
(*Runs*)

## Overview

A run represents an instance of a task's execution. See Tasks API for how to execute tasks.

### Available Operations

* [Get](#get) - Cancel Run
* [GetOutputs](#getoutputs) - Get Run Outputs
* [List](#list) - List Runs

## Get

Get information about an existing run.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/airplane-go-sdk/v3/pkg/models/shared"
	airplanegosdk "github.com/speakeasy-sdks/airplane-go-sdk/v3"
	"context"
	"log"
)

func main() {
    s := airplanegosdk.New(
        airplanegosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.Runs.Get(ctx, id)
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

**[*operations.GetRunResponse](../../pkg/models/operations/getrunresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetOutputs

Get outputs from an existing run.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/airplane-go-sdk/v3/pkg/models/shared"
	airplanegosdk "github.com/speakeasy-sdks/airplane-go-sdk/v3"
	"context"
	"log"
)

func main() {
    s := airplanegosdk.New(
        airplanegosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.Runs.GetOutputs(ctx, id)
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

**[*operations.GetOutputsResponse](../../pkg/models/operations/getoutputsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

List Runs

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/airplane-go-sdk/v3/pkg/models/shared"
	airplanegosdk "github.com/speakeasy-sdks/airplane-go-sdk/v3"
	"context"
	"log"
)

func main() {
    s := airplanegosdk.New(
        airplanegosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    var limit *int64 = airplanegosdk.Int64(768578)

    var page *int64 = airplanegosdk.Int64(99895)

    var since *string = airplanegosdk.String("<value>")

    var taskID *string = airplanegosdk.String("<value>")

    var taskSlug *string = airplanegosdk.String("<value>")

    var until *string = airplanegosdk.String("<value>")

    ctx := context.Background()
    res, err := s.Runs.List(ctx, limit, page, since, taskID, taskSlug, until)
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

**[*operations.ListRunsResponse](../../pkg/models/operations/listrunsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
