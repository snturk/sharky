package utils

import (
	"errors"
	"fmt"
	"sharky/model"
)

func RoleNotExistError(role model.SharkyRole) error {
	return newError("SharkyRole: %s is not exists on permission map for related SharkyObject", role.Name)
}

func newError(msg string, args ...string) error {
	message := fmt.Sprintf(msg, args)

	return errors.New(message)
}
