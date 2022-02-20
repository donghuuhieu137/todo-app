package data

import "todo-app/proto/todo"

var Todos []todo.Todo

func init() {
	Todos = []todo.Todo{
		{Id: 1, Name: "Monday", Content: "Learn Java"},
		{Id: 2, Name: "Tuesday", Content: "Learn Nodejs"},
		{Id: 3, Name: "Wednesday", Content: "Learn Go Lang"},
		{Id: 4, Name: "Thursday", Content: "Learn Ruby"},
		{Id: 5, Name: "Friday", Content: "Learn C#"},
	}
}
