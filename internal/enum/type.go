package enum

type ProductType int

const (
	TypeElectrinics ProductType = iota
	TypeCloth
	TypeShoes
)

var productTypeToName = map[ProductType]string{
	TypeElectrinics: "электроника",
	TypeCloth:       "одежда",
	TypeShoes:       "обувь",
}

var nameToProductType = map[string]ProductType{
	"электроника": TypeElectrinics,
	"одежда":      TypeCloth,
	"обувь":       TypeShoes,
}

func CheckProductType(pt string) bool {
	_, ok := nameToProductType[pt]
	return ok
}

func (pt ProductType) String() string {
	return productTypeToName[pt]
}
