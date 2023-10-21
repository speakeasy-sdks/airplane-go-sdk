# Tasks
(*Tasks*)

## Overview

A Task is a lightweight app that represents a single business operation for people at your company to execute.

### Available Operations

* [Execute](#execute) - Execute Task

## Execute

Execute a task with a set of parameter values and receive a run ID to track the task's execution.
Check on the status of your newly created run with [/runs/get](/api/runs#runs-get).

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
        airplanegosdk.WithSecurity(""),
    )


    executeTaskRequest := shared.ExecuteTaskRequest{
        ID: airplanegosdk.String("tsk20210728zxb2vxn"),
        ParamValues: map[string]string{
            "user": "eric",
            "limit": "20",
        },
        Resources: map[string]string{
            "key": "string",
        },
        Slug: airplanegosdk.String("hello_world"),
    }

    var envSlug *string = "string"

    ctx := context.Background()
    res, err := s.Tasks.Execute(ctx, executeTaskRequest, envSlug)
    if err != nil {
        log.Fatal(err)
    }

    if res.ExecuteTaskResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `executeTaskRequest`                                                   | [shared.ExecuteTaskRequest](../../models/shared/executetaskrequest.md) | :heavy_check_mark:                                                     | ExecuteTaskRequest                                                     |
| `envSlug`                                                              | **string*                                                              | :heavy_minus_sign:                                                     | Environment to execute the task in.                                    |


### Response

**[*operations.ExecuteTaskResponse](../../models/operations/executetaskresponse.md), error**

