// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GetPromptResponse struct {
	Prompt *PromptsifacePrompt `json:"prompt,omitempty"`
}

func (o *GetPromptResponse) GetPrompt() *PromptsifacePrompt {
	if o == nil {
		return nil
	}
	return o.Prompt
}
