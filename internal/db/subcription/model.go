package subcription

type Subcription struct {
	Email *string `db:"EMAIL" json:"email"`
	Url   *string `db:"URL" json:"URL"`
}
