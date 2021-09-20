package todo

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}
