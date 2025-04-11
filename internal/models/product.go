package models

type AddProduct struct {
	Type  string
	PvzID int
}

type DeleteProduct struct {
	PvzID int
}

func (ap *AddProduct) Validate() error {
	return nil
}

func (dp *DeleteProduct) Validate() error {
	return nil
}
