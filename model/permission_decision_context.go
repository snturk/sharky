package model

import "github.com/voicera/gooseberry/urn"

type Permission struct {
	urn    urn.URN
	action SharkyAction
}

type PermissionBuilder struct {
	urn    urn.URN
	action SharkyAction
}

func (p Permission) GetUrn() urn.URN {
	return p.urn
}

func (p Permission) GetAction() SharkyAction {
	return p.action
}

func (p Permission) Builder() PermissionBuilder {
	return PermissionBuilder{}
}

func (builder PermissionBuilder) SetUrn(urn urn.URN) PermissionBuilder {
	builder.urn = urn
	return builder
}

func (builder PermissionBuilder) SetAction(action SharkyAction) PermissionBuilder {
	builder.action = action
	return builder
}

func (builder PermissionBuilder) Build() Permission {
	var permission = Permission{
		urn:    builder.urn,
		action: builder.action,
	}
	return permission
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
