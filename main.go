package main

import (
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Complete one rest api", Completed: false},
	{ID: "3", Item: "learn game dev in go", Completed: true},
}

func get_todo_and_id(id string)(*todo,error){
	for i,t := range todos{
		if t.ID == id{
			return &todos[i],nil
		}
	}
	return nil,errors.New("todo not found")
}

func get_todos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func add_todos(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err!=nil{
		return
	}
	
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}
func get_todo(context *gin.Context){
	id:=context.Param("id")
	todo ,err := get_todo_and_id(id)
	if err!= nil{
		context.IndentedJSON(http.StatusNotFound,gin.H{"message":"todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK,todo)
}
func updateTodo(context *gin.Context){
	var newTodo todo
	if err := context.BindJSON(&newTodo); err!=nil{
		return
	}
	id:=context.Param("id")
	todo , err:= get_todo_and_id(id)
	if err!= nil{
		
		context.IndentedJSON(http.StatusNotFound,gin.H{"message":"todo not found"})
		return
	}
	*todo=newTodo
	context.IndentedJSON(http.StatusFound,todo)
}


func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			context.IndentedJSON(http.StatusOK, gin.H{"message": "todo deleted"})
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func main() {
	router := gin.Default()
	router.GET("/todos", get_todos)
	router.POST("/todos", add_todos)
	router.GET("/todos/:id",get_todo)
	router.PATCH("/todos/:id",updateTodo)
	router.DELETE("/todos/:id",deleteTodo)
	router.Run("localhost:9090")
}
