package todo

type Item struct {
	ID       uint
	Title    string
	Priority uint
	TodoList *List
	ListID   uint
	UserID   uint
}

type List struct {
	ID          uint
	Name        string
	Description string
	Todos       []*Item
	UserID      uint
}
