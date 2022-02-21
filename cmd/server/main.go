package main

import (
	"fmt"
	"log"
	"net"

	"todo-app/handlers"
	"todo-app/proto/todo"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	err := orm.RegisterDataBase("default", "mysql", "root:dhh13072001@tcp(127.0.0.1:3306)/todo_app?charset=utf8")

	if err != nil {
		log.Panicf("Connect database error : %v", err)
	}

	// orm.RegisterModel(new(models.Todo))

	name := "default"

	// drop table and re-create
	force := false

	// print log
	verbose := true

	// error
	err = orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connect database successfully !!!")
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50001")

	h := handlers.New()

	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()
	todo.RegisterTodoServiceServer(s, h)

	fmt.Println("Gateway service is running...")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
