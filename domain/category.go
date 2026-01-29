package domain

type Category struct {
	ID          int64  `db:"id" json:"-"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Seq         int16  `db:"seq" json:"seq"`
}
