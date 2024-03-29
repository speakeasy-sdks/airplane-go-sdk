// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ParameterInput struct {
	// Used to specify UI-only type modifiers
	Component   *Component   `json:"component,omitempty"`
	Constraints *Constraints `json:"constraints,omitempty"`
	// Optional default value for this parameter, used if not set.
	Default interface{} `json:"default,omitempty"`
	// Description for this parameter.
	Desc *string `json:"desc,omitempty"`
	// Name for this parameter.
	Name *string `json:"name,omitempty"`
	// If this parameter has an object data type, represents an ordered list of key-value pairs that can be included in this object.
	Params []Parameter `json:"params,omitempty"`
	// A human-friendly identifier for the parameter that can be referenced inside a task or runbook.
	// Airplane automatically generates a slug when provided a parameter name.
	Slug *string `json:"slug,omitempty"`
	// Parameter data type.
	Type   *Type      `json:"type,omitempty"`
	Values *Parameter `json:"values,omitempty"`
}

func (o *ParameterInput) GetComponent() *Component {
	if o == nil {
		return nil
	}
	return o.Component
}

func (o *ParameterInput) GetConstraints() *Constraints {
	if o == nil {
		return nil
	}
	return o.Constraints
}

func (o *ParameterInput) GetDefault() interface{} {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *ParameterInput) GetDesc() *string {
	if o == nil {
		return nil
	}
	return o.Desc
}

func (o *ParameterInput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ParameterInput) GetParams() []Parameter {
	if o == nil {
		return nil
	}
	return o.Params
}

func (o *ParameterInput) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *ParameterInput) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ParameterInput) GetValues() *Parameter {
	if o == nil {
		return nil
	}
	return o.Values
}
