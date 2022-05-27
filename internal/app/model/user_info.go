package model

type UserInfo struct {
	Id    int64  `db:"id"`
	Name  string `db:"name"`
	Age   int64  `db:"age"`
	Email string `db:"email"`
}
