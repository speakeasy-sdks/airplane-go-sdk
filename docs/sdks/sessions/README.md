# Sessions
(*Sessions*)

## Overview

A session represents an instance of a runbook's execution. See Runbooks API for how to execute runbooks.

### Available Operations

* [Get](#get) - Get Session
* [List](#list) - List Sessions

## Get

Get information about an existing session.

### Example Usage

```go
package main

import(
	"context"
	"log"
	airplanegosdk "github.com/speakeasy-sdks/airplane-go-sdk"
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/operations"
)

func main() {
    s := airplanegosdk.New(
        airplanegosdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )
    id := "female"

    ctx := context.Background()
    res, err := s.Sessions.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSessionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the session to retrieve.                        |


### Response

**[*operations.GetSessionResponse](../../models/operations/getsessionresponse.md), error**


## List

List Sessions

### Example Usage

```go
package main

import(
	"context"
	"log"
	airplanegosdk "github.com/speakeasy-sdks/airplane-go-sdk"
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/operations"
)

func main() {
    s := airplanegosdk.New(
        airplanegosdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )
    limit := 768578
    page := 99895
    runbookID := "compress"
    updatedAfter := "Product"
    updatedBefore := "Canada"

    ctx := context.Background()
    res, err := s.Sessions.List(ctx, limit, page, runbookID, updatedAfter, updatedBefore)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListSessionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |
| `limit`                                                               | **int64*                                                              | :heavy_minus_sign:                                                    | Number of results per call. Accepted values: 0 - 50. The default: 50. |
| `page`                                                                | **int64*                                                              | :heavy_minus_sign:                                                    | The offset used for this page of results.                             |
| `runbookID`                                                           | **string*                                                             | :heavy_minus_sign:                                                    | Only fetch sessions for this specific runbook.                        |
| `updatedAfter`                                                        | **string*                                                             | :heavy_minus_sign:                                                    | Fetch sessions last updated after this time (RFC 3339 format).        |
| `updatedBefore`                                                       | **string*                                                             | :heavy_minus_sign:                                                    | Fetch sessions last updated before this time (RFC 3339 format).       |


### Response

**[*operations.ListSessionsResponse](../../models/operations/listsessionsresponse.md), error**

