package todo

import "github.com/asdine/storm"

type ToDoStore interface {
	GetAll() ([]ToDo, error)
	GetAllFromList(listID int) ([]ToDo, error)
	GetByID(todoID int) (ToDo, error)
	Insert(listID int, todo ToDo) error
	Delete(todoID int) error
	Update(todo ToDo) error

	GetLists() ([]ToDoList, error)
	GetListByID(listID int) (ToDoList, error)
	CreateList(list ToDoList) error
	UpdateList(list ToDoList) error
	DeleteList(listID int) error
}

type toDoStore struct {
	db *storm.DB
}

func (s toDoStore) GetAll() ([]ToDo, error) {
	panic("not implemented") // TODO: Implement
}

func (s toDoStore) GetAllFromList(listID int) ([]ToDo, error) {
	panic("not implemented") // TODO: Implement
}

func (s toDoStore) GetByID(todoID int) (ToDo, error) {
	panic("not implemented") // TODO: Implement
}

func (s toDoStore) Insert(listID int, todo ToDo) error {
	panic("not implemented") // TODO: Implement
}

func (s toDoStore) Delete(todoID int) error {
	panic("not implemented") // TODO: Implement
}

func (s toDoStore) Update(todo ToDo) error {
	panic("not implemented") // TODO: Implement
}

func (s toDoStore) GetLists() ([]ToDoList, error) {
	panic("not implemented") // TODO: Implement
}

func (s toDoStore) GetListByID(listID int) (ToDoList, error) {
	panic("not implemented") // TODO: Implement
}

func (s toDoStore) CreateList(list ToDoList) error {
	panic("not implemented") // TODO: Implement
}

func (s toDoStore) UpdateList(list ToDoList) error {
	panic("not implemented") // TODO: Implement
}

func (s toDoStore) DeleteList(listID int) error {
	panic("not implemented") // TODO: Implement
}

func NewStore(stormDB *storm.DB) ToDoStore {
	return toDoStore{db: stormDB}
}
