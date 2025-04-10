package enum

type UserRole int
type City int
type ReceptionStatus int
type ProductType int

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

const (
	Moscow City = iota
	StPetersburg
	Kazan
)

var cityToName = map[City]string{
	Moscow:       "Москва",
	StPetersburg: "Санкт-Петербург",
	Kazan:        "Казань",
}

var nameToCity = map[string]City{
	"Москва":          Moscow,
	"Санкт-Петербург": StPetersburg,
	"Казань":          Kazan,
}

func CheckCity(city string) bool {
	_, ok := nameToCity[city]
	return ok
}

func (c City) String() string {
	return cityToName[c]
}

const (
	Electrinics ProductType = iota
	Cloth
	Shoes
)

var productTypeToName = map[ProductType]string{
	Electrinics: "электроника",
	Cloth:       "одежда",
	Shoes:       "обувь",
}

var nameToProductType = map[string]ProductType{
	"электроника": Electrinics,
	"одежда":      Cloth,
	"обувь":       Shoes,
}

func CheckProductType(pt string) bool {
	_, ok := nameToProductType[pt]
	return ok
}

func (pt ProductType) String() string {
	return productTypeToName[pt]
}
