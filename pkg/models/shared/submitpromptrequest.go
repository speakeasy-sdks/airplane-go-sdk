// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SubmitPromptRequest struct {
	// Unique ID of the prompt.
	ID *string `json:"id,omitempty"`
	// Mapping of parameter slug to value. You can find your prompt's parameter slugs on
	// the Airplane runs page or by fetching the prompt via the API.
	Values map[string]string `json:"values,omitempty"`
}

func (o *SubmitPromptRequest) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SubmitPromptRequest) GetValues() map[string]string {
	if o == nil {
		return nil
	}
	return o.Values
}
