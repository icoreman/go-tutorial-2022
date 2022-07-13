package todos

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var todoRepository = NewInMemoryTodoRepository()

func RegisterEndPoints(e *echo.Echo) {
	e.Static("/", "static")
	e.GET("/todos", getAllTodosHandler)
	e.POST("/todos", createTodoHandler)
	e.DELETE("/todos", deleteAllTodosHandler)

	e.GET("/todos/:id", getTodoHandler)
	e.DELETE("/todos/:id", deleteTodoHandler)
	e.PATCH("/todos/:id", updateTodoHandler)
}

func getAllTodosHandler(c echo.Context) error {
	todos := todoRepository.GetAll()
	return c.JSON(http.StatusOK, todos)
}

func createTodoHandler(c echo.Context) (err error) {
	todo := new(TodoForCreate)
	if err = c.Bind(todo); err != nil {
		return err
	}

	newTodo := todoRepository.Create(c, todo)
	return c.JSON(http.StatusOK, newTodo)
}

func deleteAllTodosHandler(c echo.Context) (err error) {
	todoRepository.DeleteAll()
	return c.NoContent(http.StatusOK)
}

func getTodoHandler(c echo.Context) (err error) {
	id := c.Param("id")
	if todo, err := todoRepository.Get(id); err != nil {
		return c.JSON(http.StatusNotFound, "Todo note was not found")
	} else {
		return c.JSON(http.StatusOK, todo)
	}
}

func deleteTodoHandler(c echo.Context) (err error) {
	id := c.Param("id")
	if err := todoRepository.Delete(id); err != nil {
		return c.JSON(http.StatusNotFound, "Todo note was not found")
	} else {
		return c.NoContent(http.StatusNoContent)
	}
}

func updateTodoHandler(c echo.Context) (err error) {
	id := c.Param("id")

	_, err = todoRepository.Get(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Todo note was not found")
	}
	todoForUpdate := new(TodoForCreate)
	if err = c.Bind(todoForUpdate); err != nil {
		return err
	}

	if updatedTodo, err := todoRepository.Update(todoForUpdate); err != nil {
		return c.JSON(http.StatusNotFound, "Todo note was not found")
	} else {
		return c.JSON(http.StatusOK, updatedTodo)
	}
}
