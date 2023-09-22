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
	airplanegosdk "github.com/speakeasy-sdks/airplane-go-sdk"
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/shared"
)

func main() {
    s := airplanegosdk.New(
        airplanegosdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Prompts.Cancel(ctx, shared.CancelPromptRequest{
        ID: airplanegosdk.String("pmt20221122zyydx3rho2t"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CancelPromptResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.CancelPromptRequest](../../models/shared/cancelpromptrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.CancelPromptResponse](../../models/operations/cancelpromptresponse.md), error**


## Get

Get information about an existing prompt.

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
    id := "distinctio"

    ctx := context.Background()
    res, err := s.Prompts.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetPromptResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the prompt to retrieve.                         |


### Response

**[*operations.GetPromptResponse](../../models/operations/getpromptresponse.md), error**


## List

List prompts from existing runs.

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
    runID := "quibusdam"

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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `runID`                                               | *string*                                              | :heavy_check_mark:                                    | ID of the run to retrieve prompts for.                |


### Response

**[*operations.ListPromptsResponse](../../models/operations/listpromptsresponse.md), error**


## Submit

Submit a prompt with a set of parameter values.

### Example Usage

```go
package main

import(
	"context"
	"log"
	airplanegosdk "github.com/speakeasy-sdks/airplane-go-sdk"
	"github.com/speakeasy-sdks/airplane-go-sdk/pkg/models/shared"
)

func main() {
    s := airplanegosdk.New(
        airplanegosdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Prompts.Submit(ctx, shared.SubmitPromptRequest{
        ID: airplanegosdk.String("pmt20221122zyydx3rho2t"),
        Values: map[string]string{
            "unde": "nulla",
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

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.SubmitPromptRequest](../../models/shared/submitpromptrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.SubmitPromptResponse](../../models/operations/submitpromptresponse.md), error**

