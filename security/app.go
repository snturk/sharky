package security

import "sharky/security/model"

func CanWithContext(ctx model.SharkySecurityContext) (bool, error) {
	return true, nil
}

// Rolün, Obje üzerinde Action uygulama yetkisi var mı?
func CanWithRole(role model.SharkyRole, action model.SharkyAction, obj model.SharkyObject) (bool, error) {
	return true, nil
}

// Öznenin, Obje üzerinde Action uygulama yetkisi var mı?
func CanWithSubject(sbj model.SharkySubject, action model.SharkyAction, obj model.SharkyObject) (bool, error) {
	return true, nil
}

