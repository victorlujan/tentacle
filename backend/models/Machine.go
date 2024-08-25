package models

type Machine struct {
	ID          int    `db:"id"`
	Description string `db:"description"`
}
