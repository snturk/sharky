package model

type SharkySubject struct {

}

type SharkyRole struct {
	Name string
}

type SharkyAction struct {
	Name string
}

type SharkySecurityContext interface {
	GetSubject() SharkySubject
	GetRole() SharkyRole
	GetAction() SharkyAction
	GetObject() SharkyObject
}
