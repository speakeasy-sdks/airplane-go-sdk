// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ExecuteRunbookRequest
type ExecuteRunbookRequest struct {
	// Unique ID of the runbook. You can find your runbook's ID by visiting the runbook's page on Airplane.
	// The runbook ID is located at the end of the url.
	//
	// e.g. the runbook ID for `https://app.airplane.dev/runbooks/rbk20220120z15kl79` is `rbk20220120z15kl79`
	ID string `json:"id"`
	// Mapping of parameter slug to value. You can find your runbooks's parameter slugs inside the
	// runbook editor on Airplane.
	ParamValues map[string]string `json:"paramValues,omitempty"`
	// Unique slug of the runbook. You can find your runbook's slug next to the runbook's name within the
	// runbook editor on Airplane.
	//
	// Either an ID or a slug must be provided.
	Slug *string `json:"slug,omitempty"`
}

func (o *ExecuteRunbookRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ExecuteRunbookRequest) GetParamValues() map[string]string {
	if o == nil {
		return nil
	}
	return o.ParamValues
}

func (o *ExecuteRunbookRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}
