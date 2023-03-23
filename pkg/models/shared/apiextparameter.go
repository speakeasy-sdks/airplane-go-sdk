// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ApiextParameterComponentEnum - Used to specify UI-only type modifiers
type ApiextParameterComponentEnum string

const (
	ApiextParameterComponentEnumUnknown   ApiextParameterComponentEnum = ""
	ApiextParameterComponentEnumEditorSQL ApiextParameterComponentEnum = "editor-sql"
	ApiextParameterComponentEnumTextarea  ApiextParameterComponentEnum = "textarea"
)

// ApiextParameterTypeEnum - Parameter data type.
type ApiextParameterTypeEnum string

const (
	ApiextParameterTypeEnumAny       ApiextParameterTypeEnum = "any"
	ApiextParameterTypeEnumString    ApiextParameterTypeEnum = "string"
	ApiextParameterTypeEnumBoolean   ApiextParameterTypeEnum = "boolean"
	ApiextParameterTypeEnumUpload    ApiextParameterTypeEnum = "upload"
	ApiextParameterTypeEnumInteger   ApiextParameterTypeEnum = "integer"
	ApiextParameterTypeEnumFloat     ApiextParameterTypeEnum = "float"
	ApiextParameterTypeEnumDate      ApiextParameterTypeEnum = "date"
	ApiextParameterTypeEnumDatetime  ApiextParameterTypeEnum = "datetime"
	ApiextParameterTypeEnumConfigvar ApiextParameterTypeEnum = "configvar"
	ApiextParameterTypeEnumList      ApiextParameterTypeEnum = "list"
	ApiextParameterTypeEnumMap       ApiextParameterTypeEnum = "map"
	ApiextParameterTypeEnumObject    ApiextParameterTypeEnum = "object"
)

type ApiextParameter struct {
	// Used to specify UI-only type modifiers
	Component   *ApiextParameterComponentEnum `json:"component,omitempty"`
	Constraints *ApiextConstraints            `json:"constraints,omitempty"`
	// Optional default value for this parameter, used if not set.
	Default interface{} `json:"default,omitempty"`
	// Description for this parameter.
	Desc *string `json:"desc,omitempty"`
	// Name for this parameter.
	Name *string `json:"name,omitempty"`
	// If this parameter has an object data type, represents an ordered list of key-value pairs that can be included in this object.
	Params []ApiextParameter `json:"params,omitempty"`
	// A human-friendly identifier for the parameter that can be referenced inside a task or runbook.
	// Airplane automatically generates a slug when provided a parameter name.
	Slug *string `json:"slug,omitempty"`
	// Parameter data type.
	Type   *ApiextParameterTypeEnum `json:"type,omitempty"`
	Values *ApiextParameter         `json:"values,omitempty"`
}
