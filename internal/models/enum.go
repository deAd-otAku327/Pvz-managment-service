package models

type UserRole int
type City int
type ReceptionStatus int
type ProductType int

const (
	RoleEmployye UserRole = iota
	RoleModerator
)

var roleName = map[UserRole]string{
	RoleEmployye:  "employee",
	RoleModerator: "moderator",
}

func (r UserRole) String() string {
	return roleName[r]
}
