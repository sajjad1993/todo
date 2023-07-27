package todo

type Item struct {
	ID       uint
	Title    string
	Priority uint
	TodoList *List
	ListId   uint
	UserId   uint
}

type List struct {
	ID          uint
	Name        string
	Description string
	Todos       []*Item
	UserID      uint
}
