package authservice_dao

type User struct {
	ID             int64  `db:"id"`
	UserName       string `db:"username"`
	HashedPassword string `db:"password"`
	Salt           string `db:"salt"`
	Email          string `db:"email"`
	Info           string `db:"info"`
	Role           int    `db:"role"`
}
