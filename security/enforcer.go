package security

import "sharky/security/model"

var lastResult bool

type enforcer struct {
	// TODO: Fields needed. Issue #1
}

func InitializeEnforcer(params ...string) {
	// TODO: Enforcer needs parameters to initialize, relevant params will be added. Issue #1
}

func (e enforcer) GetEnforcer() (enforcer, error) {
	// TODO: Implement GetEnforcer method. Issue #1
	return enforcer{}, nil
}

func (e enforcer) Enforce(context model.PermissionDecisionContext) bool {
	// TODO: Implement Enforce method. Issue #1
	return false
}
