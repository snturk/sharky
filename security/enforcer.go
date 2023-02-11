package security

import (
	"log"
	"sharky/model"
)

type enforcer struct {
	logger log.Logger
}

func GetEnforcer(logger log.Logger) (enforcer, error) {
	return enforcer{logger: logger}, nil
}

func (e enforcer) Enforce(context model.PermissionDecisionContext) bool {
	e.logger.Printf("Enforcing context: %v", context)
	obj := context.GetObject()
	role := context.GetRole()
	action := context.GetAction()

	permissions, ok := role.GetPermissionsForURN(obj.Urn)
	if !ok {
		e.logger.Printf("No permissions found for role %v and object %v", role, obj)
		return false
	}

	for _, permission := range permissions {
		if permission.GetAction() == action {
			e.logger.Printf("PermissionRule found for role %v and object %v", role, obj)
			return true
		}
	}

	e.logger.Printf("No permission found for role %v and object %v", role, obj)
	return false
}

func (e enforcer) SetLogger(logger log.Logger) {
	e.logger = logger
	e.logger.Printf("Enforcer logger set to %v", logger)
}
