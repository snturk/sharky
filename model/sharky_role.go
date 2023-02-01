package model

import "github.com/voicera/gooseberry/urn"

type SharkyRole struct {
	Name        string
	Permissions []Permission
}

func (r SharkyRole) AddPermission(p Permission) {
	r.Permissions = append(r.Permissions, p)
}

func (r SharkyRole) RemovePermission(p Permission) {
	for i, permission := range r.Permissions {
		if permission == p {
			r.Permissions = append(r.Permissions[:i], r.Permissions[i+1:]...)
		}
	}
}

func (r SharkyRole) HasPermission(p Permission) bool {
	for _, permission := range r.Permissions {
		if permission == p {
			return true
		}
	}
	return false
}

func (r SharkyRole) GetPermissionsForURN(urn urn.URN) ([]Permission, bool) {
	var permissions []Permission
	for _, permission := range r.Permissions {
		if permission.GetUrn() == urn {
			permissions = append(permissions, permission)
		}
	}
	return permissions, len(permissions) > 0
}

func (r SharkyRole) GetPermissionsForNamespace(namespaceId string) []Permission {
	var permissions []Permission
	for _, permission := range r.Permissions {
		if permission.urn.GetNamespaceID() == namespaceId {
			permissions = append(permissions, permission)
		}
	}
	return permissions
}
