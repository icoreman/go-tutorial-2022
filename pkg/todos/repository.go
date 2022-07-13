package todos

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lithammer/shortuuid"
)

type InMemoryTodoRepository struct {
	Todos []*Todo
}

func NewInMemoryTodoRepository() InMemoryTodoRepository {
	t := new(InMemoryTodoRepository)
	t.Todos = make([]*Todo, 0)
	return *t
}

func (r *InMemoryTodoRepository) Create(c echo.Context, todo *TodoForCreate) *Todo {
	id := shortuuid.New()
	newTodo := &Todo{
		Id:    id,
		Url:   fmt.Sprintf("http://%s%s/%s", c.Request().Host, c.Request().RequestURI, id),
		Title: todo.Title,
	}
	if todo.Completed == nil {
		newTodo.Completed = false
	} else {
		newTodo.Completed = *todo.Completed
	}
	if todo.Order != nil {
		newTodo.Order = todo.Order
	}
	r.Todos = append(r.Todos, newTodo)
	return newTodo
}

func (r *InMemoryTodoRepository) GetAll() []*Todo {
	return r.Todos
}

func (r *InMemoryTodoRepository) DeleteAll() {
	r.Todos = make([]*Todo, 0)
}

func (r *InMemoryTodoRepository) Get(id string) (t *Todo, err error) {
	for _, t = range r.Todos {
		if t.Id == id {
			return t, nil
		}
	}
	return nil, errors.New("todo not found")
}

func (r *InMemoryTodoRepository) Delete(id string) (err error) {
	for i, t := range r.Todos {
		if t.Id == id {
			r.Todos = append(r.Todos[:i], r.Todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}

func (r *InMemoryTodoRepository) Update(todo *TodoForCreate) (t *Todo, err error) {
	for i, t := range r.Todos {
		if t.Id == todo.Id {
			newTodo := &Todo{
				Id:    todo.Id,
				Url:   todo.Url,
				Title: todo.Title,
			}
			if todo.Completed == nil {
				newTodo.Completed = false
			} else {
				newTodo.Completed = *todo.Completed
			}
			if todo.Order == nil && t.Order != nil {
				newTodo.Order = t.Order
			} else if todo.Order != nil {
				newTodo.Order = todo.Order
			}
			r.Todos[i] = newTodo
			return newTodo, nil
		}
	}
	return nil, errors.New("todo not found")
}
