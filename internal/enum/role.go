package enum

type UserRole int

const (
	Employye UserRole = iota
	Moderator
)

var roleToName = map[UserRole]string{
	Employye:  "employee",
	Moderator: "moderator",
}

var nameToRole = map[string]UserRole{
	"employee":  Employye,
	"moderator": Moderator,
}

func CheckRole(role string) bool {
	_, ok := nameToRole[role]
	return ok
}

func (r UserRole) String() string {
	return roleToName[r]
}
