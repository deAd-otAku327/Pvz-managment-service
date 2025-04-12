package enum

type ProductType int

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
