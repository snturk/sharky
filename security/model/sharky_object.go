package model

import "sharky/utils"

type permissionMap map[SharkyRole][]SharkyAction

func (obj SharkyObject) GetActions(role SharkyRole) ([]SharkyAction, error) {
	actions, exists := obj.permissionMap[role]

	if exists {
		return actions, nil
	}
	return nil, utils.RoleNotExistError(role)
}

type SharkyObject struct {
	Name string
	permissionMap
	Type string // TODO: Can type be changed to another data type? (URN, URI, etc.) Issue #1 go get
	// TODO: Another field needed to implement Object-Specific permissions. Issue #1
}

type SharkyObjectBuilder struct {
	name string
	permissionMap
}

func (obj SharkyObject) Builder() *SharkyObjectBuilder {
	return &SharkyObjectBuilder{}
}

func (builder SharkyObjectBuilder) SetName(name string) SharkyObjectBuilder {
	builder.name = name
	return builder
}

func (builder SharkyObjectBuilder) SetPermissionMap(pm map[SharkyRole][]SharkyAction) SharkyObjectBuilder {
	builder.permissionMap = pm
	return builder
}

func (builder SharkyObjectBuilder) Build() SharkyObject {
	var obj = SharkyObject{
		Name:          builder.name,
		permissionMap: builder.permissionMap,
	}
	return obj
}