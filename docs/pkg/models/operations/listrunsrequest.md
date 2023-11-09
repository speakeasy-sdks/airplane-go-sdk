# ListRunsRequest


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Limit`                                                           | **int64*                                                          | :heavy_minus_sign:                                                | Number of results per call. Accepted values: 0 - 50. Default: 50. |
| `Page`                                                            | **int64*                                                          | :heavy_minus_sign:                                                | The offset used for this page of results.                         |
| `Since`                                                           | **string*                                                         | :heavy_minus_sign:                                                | Fetch runs last updated after this time (RFC 3339 format).        |
| `TaskID`                                                          | **string*                                                         | :heavy_minus_sign:                                                | Unique ID of the task to fetch runs for.                          |
| `TaskSlug`                                                        | **string*                                                         | :heavy_minus_sign:                                                | Unique slug of the task to fetch runs for.                        |
| `Until`                                                           | **string*                                                         | :heavy_minus_sign:                                                | Fetch runs last updated before this time (RFC 3339 format).       |