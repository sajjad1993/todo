package grpc

import (
	"context"
	"github.com/sajjad1993/todo/internal/common"
	rpc "github.com/sajjad1993/todo/internal/common/rpc/todo_list/api/protobuf"
	"github.com/sajjad1993/todo/internal/todo_list/app"
	"github.com/sajjad1993/todo/internal/todo_list/domain/todo"
	"github.com/sajjad1993/todo/pkg/log"
	"google.golang.org/grpc/status"
)

type Handler struct {
	rpc.UnimplementedTodoServiceServer
	todoService app.UseCase
	logger      log.Logger
}

func New(todoService app.UseCase, logger log.Logger) *Handler {
	return &Handler{
		todoService: todoService,
		logger:      logger,
	}
}

func (h *Handler) ListTodoList(ctx context.Context, request *rpc.ListTodoListRequest) (*rpc.ListTodoListResponse, error) {
	entites, err := h.todoService.GetTodoListByUserID(ctx, uint(request.UserID))
	if err != nil {
		return nil, status.Errorf(common.GetGrpcStatusCodeByError(err), err.Error())
	}
	return &rpc.ListTodoListResponse{TodoLists: toProto(entites)}, nil
}

func toProto(lists []*todo.List) []*rpc.TodoList {
	var result []*rpc.TodoList
	for _, list := range lists {
		todoList := &rpc.TodoList{
			ID:          uint64(list.ID),
			Name:        list.Name,
			Description: list.Description,
		}
		var items []*rpc.TodoItem
		for _, item := range list.Todos {
			todoItem := &rpc.TodoItem{
				ID:       uint64(item.ID),
				Title:    item.Title,
				Priority: uint64(item.Priority),
			}
			items = append(items, todoItem)
		}
		todoList.TodoItems = items
		result = append(result, todoList)
	}
	return result
}
