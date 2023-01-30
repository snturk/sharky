package model

type PermissionDecisionContext struct {
	subject SharkySubject
	role    SharkyRole
	action  SharkyAction
	object  SharkyObject
}

func NewPermissionDecisionContext(subject SharkySubject, role SharkyRole,
	action SharkyAction, object SharkyObject) PermissionDecisionContext {
	return PermissionDecisionContext{
		subject: subject,
		role:    role,
		action:  action,
		object:  object,
	}
}

func (context PermissionDecisionContext) GetSubject() SharkySubject {
	return context.subject
}

func (context PermissionDecisionContext) GetRole() SharkyRole {
	return context.role
}

func (context PermissionDecisionContext) GetAction() SharkyAction {
	return context.action
}

func (context PermissionDecisionContext) GetObject() SharkyObject {
	return context.object
}
