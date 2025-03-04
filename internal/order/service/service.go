package orderservice

import (
	"context"

	test "example.com/m/pkg/api/order/api"
)

type Repository interface {
	Create(context.Context, *test.Order) (string, error)
	Get(context.Context, string) (*test.Order, error)
	Update(context.Context, *test.Order) (*test.Order, error)
	Delete(context.Context, string) (bool, error)
	List(context.Context) ([]*test.Order, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateOrder(ctx context.Context, req *test.CreateOrderRequest) (*test.CreateOrderResponse, error) {
	order := &test.Order{
		Id:       "",
		Item:     req.Item,
		Quantity: req.Quantity,
	}

	createdOrderId, _ := s.repo.Create(ctx, order)
	return &test.CreateOrderResponse{Id: createdOrderId}, nil
}

func (s Service) GetOrder(ctx context.Context, req *test.GetOrderRequest) (*test.GetOrderResponse, error) {
	order, _ := s.repo.Get(ctx, req.Id)
	return &test.GetOrderResponse{Order: order}, nil
}

func (s Service) UpdateOrder(ctx context.Context, req *test.UpdateOrderRequest) (*test.UpdateOrderResponse, error) {
	order := &test.Order{
		Id:       req.Id,
		Item:     req.Item,
		Quantity: req.Quantity,
	}
	updatedOrder, _ := s.repo.Update(ctx, order)
	return &test.UpdateOrderResponse{Order: updatedOrder}, nil
}

func (s Service) DeleteOrder(ctx context.Context, req *test.DeleteOrderRequest) (*test.DeleteOrderResponse, error) {
	flag, _ := s.repo.Delete(ctx, req.Id)
	return &test.DeleteOrderResponse{Success: flag}, nil
}

func (s Service) ListOrders(ctx context.Context, req *test.ListOrdersRequest) (*test.ListOrdersResponse, error) {
	orders, _ := s.repo.List(ctx)
	return &test.ListOrdersResponse{Orders: orders}, nil
}
