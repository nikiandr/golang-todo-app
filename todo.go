package todo

type List struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type Item struct {
	Id          int    `json:"id" db:"id"`
	ListId      int    `json:"-" db:"list_id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done" binding:"required"`
}
