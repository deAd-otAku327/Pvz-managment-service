package models

type CreateReception struct {
	PvzID int
}

type CloseReception struct {
	PvzID int
}

func (cr *CreateReception) Validate() error {
	return nil
}

func (cr *CloseReception) Validate() error {
	return nil
}
