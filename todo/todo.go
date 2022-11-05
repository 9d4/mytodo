package todo

type ToDo struct {
	ID          int
	Title       string `storm:"index"`
	Description string
	ListID      int `storm:"index"`
}

type ToDoList struct {
	ID          int
	Title       string `storm:"index"`
	Description string
}
