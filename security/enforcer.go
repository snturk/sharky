package security

import (
	"sharky/model"
)

type enforcer struct {
}

func GetEnforcer() (enforcer, error) {
	return enforcer{}, nil
}

func (e enforcer) Enforce(context model.PermissionDecisionContext) bool {
	obj := context.GetObject()
	role := context.GetRole()
	action := context.GetAction()

	permissions, ok := role.GetPermissionsForURN(obj.Urn)
	if !ok {
		return false
	}

	for _, permission := range permissions {
		if permission.GetAction() == action {
			return true
		}
	}

	return false
}
