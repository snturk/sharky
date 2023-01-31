package security

import "sharky/security/model"

var lastResult bool

type enforcer struct {
	// TODO: Fields needed. Issue #1
}

func InitializeEnforcer(params ...string) {
	// TODO: Enforcer needs parameters to initialize, relevant params will be added. Issue #1
}

func GetEnforcer() (enforcer, error) {
	// TODO: Implement GetEnforcer method. Issue #1
	return enforcer{}, nil
}

func (e enforcer) Enforce(context model.PermissionDecisionContext) bool {
	// TODO: Improve basic implementation Issue #1
	obj := context.GetObject()
	role := context.GetRole()
	action := context.GetAction()

	permissions, ok := role.GetPermissionsForURN(obj.Urn)
	if !ok {
		return false
	}

	for _, permission := range permissions {
		if permission.Action == action {
			return true
		}
	}

	return false
}
