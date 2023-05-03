# Sessions

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
    res, err := s.Sessions.Get(ctx, operations.GetSessionRequest{
        ID: "7739251a-a52c-43f5-ad01-9da1ffe78f09",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSessionResponse != nil {
        // handle response
    }
}
```

## List

List Sessions

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
    res, err := s.Sessions.List(ctx, operations.ListSessionsRequest{
        Limit: airplane.Int64(451159),
        Page: airplane.Int64(739264),
        RunbookID: airplane.String("perferendis"),
        UpdatedAfter: airplane.String("doloremque"),
        UpdatedBefore: airplane.String("reprehenderit"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListSessionResponse != nil {
        // handle response
    }
}
```
