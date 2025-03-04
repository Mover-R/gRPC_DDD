package service

import (
	"context"
	"fmt"
	"sync"

	test "example.com/m/pkg/api/order/api"
)

type Service struct {
	test.OrderServiceServer
}

func New() *Service {
	return &Service{}
}

type orderServiceServer struct {
	test.UnimplementedOrderServiceServer
	mu     sync.Mutex
	orders map[string]*test.Order
}

func NewOrderServiceServer() *orderServiceServer {
	return &orderServiceServer{
		orders: make(map[string]*test.Order),
	}
}

func (s *orderServiceServer) CreateOrder(ctx context.Context, req *test.CreateOrderRequest) (*test.CreateOrderResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	orderID := fmt.Sprintf("order-%d", len(s.orders)+1)
	order := &test.Order{
		Id:       orderID,
		Item:     req.Item,
		Quantity: req.Quantity,
	}

	s.orders[orderID] = order

	return &test.CreateOrderResponse{Id: orderID}, nil
}

func (s *orderServiceServer) GetOrder(ctx context.Context, req *test.GetOrderRequest) (*test.GetOrderResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, exists := s.orders[req.Id]
	if !exists {
		return nil, fmt.Errorf("order with ID %s not found", req.Id)
	}

	return &test.GetOrderResponse{Order: order}, nil
}

func (s *orderServiceServer) UpdateOrder(ctx context.Context, req *test.UpdateOrderRequest) (*test.UpdateOrderResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, exists := s.orders[req.Id]
	if !exists {
		return nil, fmt.Errorf("order with ID %s not found", req.Id)
	}

	order.Item = req.Item
	order.Quantity = req.Quantity

	return &test.UpdateOrderResponse{Order: order}, nil
}

func (s *orderServiceServer) DeleteOrder(ctx context.Context, req *test.DeleteOrderRequest) (*test.DeleteOrderResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.orders[req.Id]
	if !exists {
		return nil, fmt.Errorf("order with ID %s not found", req.Id)
	}

	delete(s.orders, req.Id)

	return &test.DeleteOrderResponse{Success: true}, nil
}

func (s *orderServiceServer) ListOrders(ctx context.Context, req *test.ListOrdersRequest) (*test.ListOrdersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	orders := make([]*test.Order, 0, len(s.orders))
	for _, order := range s.orders {
		orders = append(orders, order)
	}

	return &test.ListOrdersResponse{Orders: orders}, nil
}
