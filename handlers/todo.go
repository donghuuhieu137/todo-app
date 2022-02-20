package handlers

import (
	"context"
	"todo-app/data"
	"todo-app/proto/todo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TodoInteractor interface {
	Get(ctx context.Context, in *todo.Nil) (*todo.GetResponse, error)
	GetById(ctx context.Context, in *todo.Todo) (*todo.Todo, error)
	Create(ctx context.Context, in *todo.Todo) (*todo.Todo, error)
	Update(ctx context.Context, in *todo.Todo) (*todo.Todo, error)
	Delete(ctx context.Context, in *todo.Todo) (*todo.Todo, error)
}

func (h *Handler) Get(ctx context.Context, in *todo.Nil) (*todo.GetResponse, error) {
	res := []*todo.Todo{}
	for _, element := range data.Todos {
		appendTodo := toProtoToDo(&element)
		res = append(res, appendTodo)
	}
	ans := todo.GetResponse{Todos: res}
	return &ans, nil
}
func (h *Handler) GetById(ctx context.Context, payload *todo.Todo) (*todo.Todo, error) {
	for _, element := range data.Todos {
		if element.Id == payload.Id {
			return toProtoToDo(&element), nil
		}
	}
	return nil, status.Error(codes.InvalidArgument, "Id not found !!!")
}
func (h *Handler) Create(ctx context.Context, payload *todo.Todo) (*todo.Todo, error) {
	data.Todos = append(data.Todos, *payload)
	return payload, nil
}
func (h *Handler) Update(ctx context.Context, payload *todo.Todo) (*todo.Todo, error) {
	for index, element := range data.Todos {
		if element.Id == payload.Id {
			data.Todos[index] = *payload
			return payload, nil
		}
	}
	return nil, status.Error(codes.Internal, "Update fail !!!")
}
func (h *Handler) Delete(ctx context.Context, payload *todo.Todo) (*todo.Todo, error) {
	for index, element := range data.Todos {
		if element.Id == payload.Id {
			data.Todos = removeIndex(data.Todos, int(index))
			return payload, nil
		}
	}
	return nil, status.Error(codes.InvalidArgument, "Id not found !!!")
}
func removeIndex(slice []todo.Todo, s int) []todo.Todo {
	ret := make([]todo.Todo, 0)
	ret = append(ret, slice[:s]...)
	return append(ret, slice[s+1:]...)
}
func toProtoToDo(payload *todo.Todo) *todo.Todo {
	return &todo.Todo{
		Id:      payload.GetId(),
		Name:    payload.GetName(),
		Content: payload.GetContent(),
	}
}
