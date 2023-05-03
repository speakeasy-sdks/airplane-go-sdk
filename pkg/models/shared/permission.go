// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PermissionRoleIDEnum - Which role is assigned to this permission.
type PermissionRoleIDEnum string

const (
	PermissionRoleIDEnumTeamAdmin        PermissionRoleIDEnum = "team_admin"
	PermissionRoleIDEnumTeamDeveloper    PermissionRoleIDEnum = "team_developer"
	PermissionRoleIDEnumTaskViewer       PermissionRoleIDEnum = "task_viewer"
	PermissionRoleIDEnumTaskRequester    PermissionRoleIDEnum = "task_requester"
	PermissionRoleIDEnumTaskExecuter     PermissionRoleIDEnum = "task_executer"
	PermissionRoleIDEnumTaskAdmin        PermissionRoleIDEnum = "task_admin"
	PermissionRoleIDEnumRunViewer        PermissionRoleIDEnum = "run_viewer"
	PermissionRoleIDEnumRunbookViewer    PermissionRoleIDEnum = "runbook_viewer"
	PermissionRoleIDEnumRunbookRequester PermissionRoleIDEnum = "runbook_requester"
	PermissionRoleIDEnumRunbookExecuter  PermissionRoleIDEnum = "runbook_executer"
	PermissionRoleIDEnumRunbookAdmin     PermissionRoleIDEnum = "runbook_admin"
	PermissionRoleIDEnumSessionViewer    PermissionRoleIDEnum = "session_viewer"
	PermissionRoleIDEnumSessionExecuter  PermissionRoleIDEnum = "session_executer"
	PermissionRoleIDEnumSessionAdmin     PermissionRoleIDEnum = "session_admin"
	PermissionRoleIDEnumResourceUser     PermissionRoleIDEnum = "resource_user"
	PermissionRoleIDEnumDeployCreator    PermissionRoleIDEnum = "deploy_creator"
	PermissionRoleIDEnumGroupAdmin       PermissionRoleIDEnum = "group_admin"
)

func (e PermissionRoleIDEnum) ToPointer() *PermissionRoleIDEnum {
	return &e
}

func (e *PermissionRoleIDEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "team_admin":
		fallthrough
	case "team_developer":
		fallthrough
	case "task_viewer":
		fallthrough
	case "task_requester":
		fallthrough
	case "task_executer":
		fallthrough
	case "task_admin":
		fallthrough
	case "run_viewer":
		fallthrough
	case "runbook_viewer":
		fallthrough
	case "runbook_requester":
		fallthrough
	case "runbook_executer":
		fallthrough
	case "runbook_admin":
		fallthrough
	case "session_viewer":
		fallthrough
	case "session_executer":
		fallthrough
	case "session_admin":
		fallthrough
	case "resource_user":
		fallthrough
	case "deploy_creator":
		fallthrough
	case "group_admin":
		*e = PermissionRoleIDEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PermissionRoleIDEnum: %v", v)
	}
}

type Permission struct {
	// Which action this permission applies to.
	Action *string `json:"action,omitempty"`
	// Which role is assigned to this permission.
	RoleID *PermissionRoleIDEnum `json:"roleID,omitempty"`
	// ID of the group  this permission applies to if assigned to a group.
	SubGroupID *string `json:"subGroupID,omitempty"`
	// ID of the user this permission applies to if assigned directly to a user.
	SubUserID *string `json:"subUserID,omitempty"`
}
