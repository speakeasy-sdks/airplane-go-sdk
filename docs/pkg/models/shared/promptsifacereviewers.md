# PromptsifaceReviewers


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `AllowSelfApprovals`                                                | **bool*                                                             | :heavy_minus_sign:                                                  | Whether or not self approvals are allowed.                          | false                                                               |
| `Groups`                                                            | []*string*                                                          | :heavy_minus_sign:                                                  | Groups allowed to review a given prompt indentified by their slugs. |                                                                     |
| `Users`                                                             | []*string*                                                          | :heavy_minus_sign:                                                  | Users allowed to review a given prompt indentified by their emails. |                                                                     |