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
