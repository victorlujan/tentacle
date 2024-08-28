package models

type User struct {
	ID         int    `db:"id"`
	Email      string `db:"email"`
	Nif        string `db:"nif"`
	Delegation string `db:"delegation"`
}
