package enum

type City int

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
