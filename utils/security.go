package utils

import (
	"sharky/security/model"
)

// Example Role & Action definitions
var (
	ROLE_ADMIN = model.SharkyRole{Name: "ADMIN"}
	ROLE_USER = model.SharkyRole{Name: "USER"}

	ACTION_VIEW = model.SharkyAction{Name: "VIEW"}
)