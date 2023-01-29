package model

import "sharky/utils"

type permissionMap map[SharkyRole][]SharkyAction

func (pm permissionMap) GetActions(role SharkyRole) ([]SharkyAction, error) {
	actions, exists := pm[role]

	if exists {
		return actions, nil
	}

	return nil, utils.RoleNotExistError(role)
}

type SharkyObject struct {
	Name string
	permissionMap
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
