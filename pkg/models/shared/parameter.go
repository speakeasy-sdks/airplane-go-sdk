// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ParameterComponentEnum - Used to specify UI-only type modifiers
type ParameterComponentEnum string

const (
	ParameterComponentEnumUnknown   ParameterComponentEnum = ""
	ParameterComponentEnumEditorSQL ParameterComponentEnum = "editor-sql"
	ParameterComponentEnumTextarea  ParameterComponentEnum = "textarea"
)

func (e ParameterComponentEnum) ToPointer() *ParameterComponentEnum {
	return &e
}

func (e *ParameterComponentEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "":
		fallthrough
	case "editor-sql":
		fallthrough
	case "textarea":
		*e = ParameterComponentEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ParameterComponentEnum: %s", s)
	}
}

// ParameterTypeEnum - Parameter data type.
type ParameterTypeEnum string

const (
	ParameterTypeEnumAny       ParameterTypeEnum = "any"
	ParameterTypeEnumString    ParameterTypeEnum = "string"
	ParameterTypeEnumBoolean   ParameterTypeEnum = "boolean"
	ParameterTypeEnumUpload    ParameterTypeEnum = "upload"
	ParameterTypeEnumInteger   ParameterTypeEnum = "integer"
	ParameterTypeEnumFloat     ParameterTypeEnum = "float"
	ParameterTypeEnumDate      ParameterTypeEnum = "date"
	ParameterTypeEnumDatetime  ParameterTypeEnum = "datetime"
	ParameterTypeEnumConfigvar ParameterTypeEnum = "configvar"
	ParameterTypeEnumList      ParameterTypeEnum = "list"
	ParameterTypeEnumMap       ParameterTypeEnum = "map"
	ParameterTypeEnumObject    ParameterTypeEnum = "object"
)

func (e ParameterTypeEnum) ToPointer() *ParameterTypeEnum {
	return &e
}

func (e *ParameterTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
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
		*e = ParameterTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ParameterTypeEnum: %s", s)
	}
}

type Parameter struct {
	// Used to specify UI-only type modifiers
	Component   *ParameterComponentEnum `json:"component,omitempty"`
	Constraints *Constraints            `json:"constraints,omitempty"`
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
	Type   *ParameterTypeEnum `json:"type,omitempty"`
	Values *Parameter         `json:"values,omitempty"`
}
