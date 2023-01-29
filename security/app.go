package security

import "sharky/security/model"

func IsContextExecutable(ctx model.PermissionDecisionContext) (bool, error) {
	return true, nil
}
