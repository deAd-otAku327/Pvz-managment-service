package consts

const (
	ID = "id"

	PvzsTable           = "pvzs"
	PvzRegistrationDate = "registration_date"
	PvzCity             = "city"

	ReceptionsTable   = "receptions"
	ReceptionDateTime = "date_time"
	ReceptionPvzID    = "pvz_id"
	ReceptionStatus   = "status"

	ProductsTable      = "products"
	ProductDateTime    = "date_time"
	ProductReceptionID = "reception_id"
	ProductType        = "type"

	// From http://www.postgresql.org/docs/9.3/static/errcodes-appendix.html
	PQInvalidTextRepresentation = "invalid_text_representation"
	PQForeignKeyViolation       = "foreign_key_violation"
)
