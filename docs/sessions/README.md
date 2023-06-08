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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Sessions.Get(ctx, "recusandae")
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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Sessions.List(ctx, 836079, 71036, "quis", "veritatis", "deserunt")
    if err != nil {
        log.Fatal(err)
    }

    if res.ListSessionResponse != nil {
        // handle response
    }
}
```
