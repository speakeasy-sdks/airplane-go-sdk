# ListSessionsRequest


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Limit`                                                               | **int64*                                                              | :heavy_minus_sign:                                                    | Number of results per call. Accepted values: 0 - 50. The default: 50. |
| `Page`                                                                | **int64*                                                              | :heavy_minus_sign:                                                    | The offset used for this page of results.                             |
| `RunbookID`                                                           | **string*                                                             | :heavy_minus_sign:                                                    | Only fetch sessions for this specific runbook.                        |
| `UpdatedAfter`                                                        | **string*                                                             | :heavy_minus_sign:                                                    | Fetch sessions last updated after this time (RFC 3339 format).        |
| `UpdatedBefore`                                                       | **string*                                                             | :heavy_minus_sign:                                                    | Fetch sessions last updated before this time (RFC 3339 format).       |