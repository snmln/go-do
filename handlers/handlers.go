package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"todos/todos"

	// "github.com/Pungyeon/golang-auth0-example/todo"
	"github.com/gin-gonic/gin"
)

// GetTodoListHandler returns all current todo items
func GetTodoListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, todos.Get())
}

// AddTodoHandler adds a new todo to the todo list
func AddTodoHandler(c *gin.Context) {
	todoItem, statusCode, err := convertHTTPBodyToTodo(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	c.JSON(statusCode, gin.H{"id": todos.Add(todoItem.Message)})
}

// DeleteTodoHandler will delete a specified todo based on user http input
func DeleteTodoHandler(c *gin.Context) {
	todoID := c.Param("id")
	if err := todos.Delete(todoID); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "")
}

// CompleteTodoHandler will complete a specified todo based on user http input
func CompleteTodoHandler(c *gin.Context) {
	todoItem, statusCode, err := convertHTTPBodyToTodo(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	if todos.Complete(todoItem.ID) != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "")
}

func RevertTodoHandler(c *gin.Context) {
	todoItem, statusCode, err := convertHTTPBodyToTodo(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	if todos.Revert(todoItem.ID) != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "")
}

func convertHTTPBodyToTodo(httpBody io.ReadCloser) (todos.Todo, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return todos.Todo{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	return convertJSONBodyToTodo(body)
}

func convertJSONBodyToTodo(jsonBody []byte) (todos.Todo, int, error) {
	var todoItem todos.Todo
	err := json.Unmarshal(jsonBody, &todoItem)
	if err != nil {
		return todos.Todo{}, http.StatusBadRequest, err
	}
	return todoItem, http.StatusOK, nil
}
