package main

type AuthRoles interface {
	Value() string
}

type UserRole struct {
	value string
}

func (u UserRole) Value() string {
	return u.value
}

type AdminRole struct {
	value string
}

func (a AdminRole) Value() string {
	return a.value
}
