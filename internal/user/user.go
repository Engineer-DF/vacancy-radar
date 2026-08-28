package user

type User struct {
	Id int64 // ID пользователя в Telegram
}

// TODO: Добавить больше фильтров в будущем для более точечного поиска.
type Filter struct {
	Text       string
	Experience string
	Area       string
	Salary     int
}
