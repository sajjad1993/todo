package todo

type Item struct {
	Title    string
	Priority uint
	TodoList List
}

type List struct {
	Name        string
	Description string
	Todos       []Item
}
