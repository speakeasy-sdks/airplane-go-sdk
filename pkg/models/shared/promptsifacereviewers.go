// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PromptsifaceReviewers struct {
	// Whether or not self approvals are allowed.
	AllowSelfApprovals *bool `json:"allowSelfApprovals,omitempty"`
	// Groups allowed to review a given prompt indentified by their slugs.
	Groups []string `json:"groups,omitempty"`
	// Users allowed to review a given prompt indentified by their emails.
	Users []string `json:"users,omitempty"`
}