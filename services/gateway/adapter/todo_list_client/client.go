package todo_list_client

import (
	"context"
	"fmt"
	"github.com/sajjad1993/todo/pkg/log"
	rpc "github.com/sajjad1993/todo/pkg/rpc/proto/todo_list/api/protobuf"
	"github.com/sajjad1993/todo/services/gateway/config"
	"github.com/sajjad1993/todo/services/gateway/domain/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type todoListService struct {
	client rpc.TodoServiceClient
	logger log.Logger
}

func (s *todoListService) GetByUserID(ctx context.Context, userID uint) ([]*todo.List, error) {
	request := &rpc.ListTodoListRequest{
		UserID: uint64(userID),
	}
	response, err := s.client.ListTodoList(ctx, request)
	if err != nil {
		return nil, err
	}

	return toEntity(response), nil
}
func toEntity(response *rpc.ListTodoListResponse) []*todo.List {
	var result []*todo.List
	for _, list := range response.TodoLists {
		todoList := &todo.List{
			ID:          uint(list.ID),
			Name:        list.Name,
			Description: list.Description,
		}
		var items []*todo.Item
		for _, item := range list.TodoItems {
			todoItem := &todo.Item{
				ID:       uint(item.ID),
				Title:    item.Title,
				Priority: uint(item.Priority),
			}
			items = append(items, todoItem)
		}
		todoList.Todos = items
		result = append(result, todoList)
	}
	return result
}
func New(logger log.Logger, config config.Config) (todo.Reader, error) {
	cc, err := grpc.Dial(fmt.Sprintf("%s", config.GetTodoServiceAddress()), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal(fmt.Sprintf("cant connect to user service: %s", err))
		cc.Close()
		return nil, err
	}
	client := rpc.NewTodoServiceClient(cc)
	service := todoListService{
		client: client,
		logger: logger,
	}
	return &service, nil

}
