package model

import "github.com/voicera/gooseberry/urn"

type Permission struct {
	ObjectURN urn.URN
	Action    SharkyAction
}

type PermissionDecisionContext struct {
	role   SharkyRole
	action SharkyAction
	object SharkyObject
}

func NewPermissionDecisionContext(role SharkyRole, action SharkyAction, object SharkyObject) PermissionDecisionContext {
	return PermissionDecisionContext{
		role:   role,
		action: action,
		object: object,
	}
}

func (pdc PermissionDecisionContext) GetRole() SharkyRole {
	return pdc.role
}

func (pdc PermissionDecisionContext) GetAction() SharkyAction {
	return pdc.action
}

func (pdc PermissionDecisionContext) GetObject() SharkyObject {
	return pdc.object
}
