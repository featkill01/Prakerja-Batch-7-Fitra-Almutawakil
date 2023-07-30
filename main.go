package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todos []Todo

func main() {
	e := echo.New()

	// Routes
	e.GET("/todos", getTodos)
	e.POST("/todos", createTodo)
	e.PUT("/todos/:id", updateTodo)
	e.DELETE("/todos/:id", deleteTodo)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler functions
func getTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}

func createTodo(c echo.Context) error {
	todo := new(Todo)
	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON format"})
	}

	todo.ID = len(todos) + 1
	todo.Completed = false
	todo.CreatedAt = time.Now()

	todos = append(todos, *todo)

	return c.JSON(http.StatusCreated, todo)
}

func updateTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid todo ID"})
	}

	found := false
	for i, todo := range todos {
		if todo.ID == id {
			found = true
			todos[i].Title = c.FormValue("title")
			todos[i].Completed, _ = strconv.ParseBool(c.FormValue("completed"))
			return c.JSON(http.StatusOK, todos[i])
		}
	}

	if !found {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
	}

	return nil
}

func deleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid todo ID"})
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
}
