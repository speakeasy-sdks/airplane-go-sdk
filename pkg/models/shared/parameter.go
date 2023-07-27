// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ParameterComponent - Used to specify UI-only type modifiers
type ParameterComponent string

const (
	ParameterComponentUnknown   ParameterComponent = ""
	ParameterComponentEditorSQL ParameterComponent = "editor-sql"
	ParameterComponentTextarea  ParameterComponent = "textarea"
)

func (e ParameterComponent) ToPointer() *ParameterComponent {
	return &e
}

func (e *ParameterComponent) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "editor-sql":
		fallthrough
	case "textarea":
		*e = ParameterComponent(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ParameterComponent: %v", v)
	}
}

// ParameterType - Parameter data type.
type ParameterType string

const (
	ParameterTypeAny       ParameterType = "any"
	ParameterTypeString    ParameterType = "string"
	ParameterTypeBoolean   ParameterType = "boolean"
	ParameterTypeUpload    ParameterType = "upload"
	ParameterTypeInteger   ParameterType = "integer"
	ParameterTypeFloat     ParameterType = "float"
	ParameterTypeDate      ParameterType = "date"
	ParameterTypeDatetime  ParameterType = "datetime"
	ParameterTypeConfigvar ParameterType = "configvar"
	ParameterTypeList      ParameterType = "list"
	ParameterTypeMap       ParameterType = "map"
	ParameterTypeObject    ParameterType = "object"
)

func (e ParameterType) ToPointer() *ParameterType {
	return &e
}

func (e *ParameterType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "any":
		fallthrough
	case "string":
		fallthrough
	case "boolean":
		fallthrough
	case "upload":
		fallthrough
	case "integer":
		fallthrough
	case "float":
		fallthrough
	case "date":
		fallthrough
	case "datetime":
		fallthrough
	case "configvar":
		fallthrough
	case "list":
		fallthrough
	case "map":
		fallthrough
	case "object":
		*e = ParameterType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ParameterType: %v", v)
	}
}

type Parameter struct {
	// Used to specify UI-only type modifiers
	Component   *ParameterComponent `json:"component,omitempty"`
	Constraints *Constraints        `json:"constraints,omitempty"`
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
	Type   *ParameterType `json:"type,omitempty"`
	Values *Parameter     `json:"values,omitempty"`
}

func (o *Parameter) GetComponent() *ParameterComponent {
	if o == nil {
		return nil
	}
	return o.Component
}

func (o *Parameter) GetConstraints() *Constraints {
	if o == nil {
		return nil
	}
	return o.Constraints
}

func (o *Parameter) GetDefault() interface{} {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *Parameter) GetDesc() *string {
	if o == nil {
		return nil
	}
	return o.Desc
}

func (o *Parameter) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Parameter) GetParams() []Parameter {
	if o == nil {
		return nil
	}
	return o.Params
}

func (o *Parameter) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *Parameter) GetType() *ParameterType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Parameter) GetValues() *Parameter {
	if o == nil {
		return nil
	}
	return o.Values
}
