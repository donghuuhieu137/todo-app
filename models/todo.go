package models

import (
	"fmt"
	"todo-app/proto/todo"

	"github.com/beego/beego/v2/client/orm"
)

func init() {
	orm.RegisterModel(new(Todo))
}

type Todo struct {
	Id      int32  `orm:"auto"`
	Name    string `orm:null`
	Content string `orm:"type(text)"`
}

func Create(todo *Todo) error {
	o := orm.NewOrm()
	id, err := o.Insert(todo)
	if err != nil {
		fmt.Println(id)
		return err
	}
	return nil
}

func Get() []*Todo {
	o := orm.NewOrm()
	var todos []*Todo
	o.QueryTable(new(Todo)).All(&todos)
	return todos
}

func GetById(id int32) (*Todo, error) {
	o := orm.NewOrm()
	todo := &Todo{
		Id:      id,
		Name:    "",
		Content: "",
	}
	err := o.Read(todo)

	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(todo)
	}
	return todo, nil
}

func Update(todo *todo.Todo) error {
	o := orm.NewOrm()
	modelTodo := toModelToDo(todo)
	num, err := o.Update(modelTodo)
	if err != nil {
		return err
	}
	fmt.Println(num)
	return nil
}

func Delete(todo *todo.Todo) error {
	o := orm.NewOrm()
	todoModel := toModelToDo(todo)
	num, err := o.Delete(todoModel)
	if err == nil {
		fmt.Println(num)
		return nil
	}
	return err
}

func toModelToDo(payload *todo.Todo) *Todo {
	return &Todo{
		Id:      payload.Id,
		Name:    payload.Name,
		Content: payload.Content,
	}
}
