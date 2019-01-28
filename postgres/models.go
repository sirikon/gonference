package postgres

// PersonModel .
type PersonModel struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}
