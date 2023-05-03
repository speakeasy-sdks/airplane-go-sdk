# Tasks

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
    res, err := s.Tasks.Execute(ctx, operations.ExecuteTaskRequest{
        ExecuteTaskRequest: shared.ExecuteTaskRequest{
            ID: airplane.String("tsk20210728zxb2vxn"),
            ParamValues: map[string]string{
                "maiores": "dicta",
                "corporis": "dolore",
            },
            Resources: map[string]string{
                "dicta": "harum",
                "enim": "accusamus",
            },
            Slug: airplane.String("hello_world"),
        },
        EnvSlug: airplane.String("commodi"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ExecuteTaskResponse != nil {
        // handle response
    }
}
```
