package utils

import (
	"sharky/security"
	"sharky/security/model"
)

// Example Role & Action definitions
var (
	ROLE_ADMIN = model.SharkyRole{Name: "ADMIN"}
	ROLE_USER  = model.SharkyRole{Name: "USER"}

	ACTION_VIEW   = model.SharkyAction{Name: "VIEW"}
	ACTION_DELETE = model.SharkyAction{Name: "DELETE"}
)

func CheckPermission(obj model.SharkyObject, role model.SharkyRole, action model.SharkyAction) bool {
	context := model.NewPermissionDecisionContext(role, action, obj)

	enforcer, err := security.GetEnforcer()
	if err != nil {
		return false
	}
	//TODO: Logging is necessary here.
	return enforcer.Enforce(context)
}
