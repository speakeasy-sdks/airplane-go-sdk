// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SubmitPromptResponse struct {
	// Unique ID of the prompt.
	ID *string `json:"id,omitempty"`
}

func (o *SubmitPromptResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
