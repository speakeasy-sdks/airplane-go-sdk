// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CancelPromptResponse - OK
type CancelPromptResponse struct {
	// Unique ID of the prompt.
	ID *string `json:"id,omitempty"`
}

func (o *CancelPromptResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
