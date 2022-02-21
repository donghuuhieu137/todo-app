package handlers

import (
	"context"
	"fmt"
	"todo-app/models"
	"todo-app/proto/todo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TodoInteractor interface {
	Get(ctx context.Context, in *todo.Nil) (*todo.GetResponse, error)
	GetById(ctx context.Context, in *todo.Todo) (*todo.Todo, error)
	Create(ctx context.Context, in *todo.Todo) (*todo.ResponseDto, error)
	Update(ctx context.Context, in *todo.Todo) (*todo.ResponseDto, error)
	Delete(ctx context.Context, in *todo.Todo) (*todo.ResponseDto, error)
}

func (h *Handler) Get(ctx context.Context, in *todo.Nil) (*todo.GetResponse, error) {
	res := []*todo.Todo{}
	array := models.Get()
	for _, element := range array {
		appendTodo := toProtoToDo(element)
		res = append(res, appendTodo)
	}
	ans := todo.GetResponse{
		Status:  200,
		Message: "Success",
		Data:    res,
	}
	return &ans, nil
}
func (h *Handler) GetById(ctx context.Context, payload *todo.Todo) (*todo.Todo, error) {
	todoModel, err := models.GetById(payload.Id)
	if err != nil {
		fmt.Println(err)
		return nil, status.Error(codes.Internal, "Fail to get todo")
	}
	todo := toProtoToDo(todoModel)
	return todo, nil
}
func (h *Handler) Create(ctx context.Context, payload *todo.Todo) (*todo.ResponseDto, error) {
	fmt.Println(payload)
	ci := toModelToDo(payload)
	err := models.Create(ci)
	if err != nil {
		fmt.Println(err)
		return nil, status.Error(codes.Internal, "Create fail !!!")
	}
	ans := &todo.ResponseDto{
		Status:  200,
		Message: "Success",
	}
	return ans, nil
}
func (h *Handler) Update(ctx context.Context, payload *todo.Todo) (*todo.ResponseDto, error) {
	result := models.Update(payload)
	if result != nil {
		fmt.Println(result)
		return nil, status.Error(codes.Internal, "Update fail !!!")
	}
	ans := todo.ResponseDto{
		Status:  200,
		Message: "Success",
	}
	return &ans, nil
}
func (h *Handler) Delete(ctx context.Context, payload *todo.Todo) (*todo.ResponseDto, error) {
	fmt.Println(payload)
	result := models.Delete(payload)
	if result != nil {
		fmt.Println(result)
		return nil, status.Error(codes.InvalidArgument, "Delete fail !!!")
	}
	ans := todo.ResponseDto{
		Status:  200,
		Message: "Success",
	}
	return &ans, nil
}
func toProtoToDo(payload *models.Todo) *todo.Todo {
	return &todo.Todo{
		Id:      payload.Id,
		Name:    payload.Name,
		Content: payload.Content,
	}
}
func toModelToDo(payload *todo.Todo) *models.Todo {
	return &models.Todo{
		Id:      payload.Id,
		Name:    payload.Name,
		Content: payload.Content,
	}
}
