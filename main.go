package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	// Context.BindJSON means whatever json inside our request body and its gonna bind
	// it to the variable newTodo that has a todo type it will return an error
	// if the request doesnt match to todo struct format like id, item, completed

	// catch the error and if the error is not null return
	if err := context.BindJSON(&newTodo); err != nil {

		// if there is an error do not execute context and return the request
		return
	}

	todos = append(todos, newTodo)

	// Return with status created and pass the newtodo
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {

	// context.Param("id") are from your route dynamic /todo/:id
	id := context.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

// (*todo, error) this will return todo and error
// if it will return an error todo will be null vice versa
func getTodoById(id string) (*todo, error) {
	for index, t := range todos {

		// Check if the loop id and the id requested is match
		if t.ID == id {
			return &todos[index], nil
		}
	}

	// If ever the request id doesn't exist on our todo slice
	return nil, errors.New("Todo not found")
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	// flip the current value vice versa
	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todo", addTodo)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todo/:id", toggleTodoStatus)
	router.Run("localhost:9090")
}
