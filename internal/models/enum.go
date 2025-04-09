package models

type UserRole int
type City int
type ReceptionStatus int
type ProductType int

const (
	RoleEmployye UserRole = iota
	RoleModerator
)

var roleToName = map[UserRole]string{
	RoleEmployye:  "employee",
	RoleModerator: "moderator",
}

var nameToRole = map[string]UserRole{
	"employee":  RoleEmployye,
	"moderator": RoleModerator,
}

func CheckRole(role string) bool {
	_, ok := nameToRole[role]
	return ok
}

func (r UserRole) String() string {
	return roleToName[r]
}
