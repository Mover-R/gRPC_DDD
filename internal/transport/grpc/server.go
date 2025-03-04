package server

import (
	"context"
	"fmt"

	test "example.com/m/pkg/api/order/api"
)

type Service interface {
	CreateOrder(context.Context, *test.CreateOrderRequest) (*test.CreateOrderResponse, error)
	GetOrder(context.Context, *test.GetOrderRequest) (*test.GetOrderResponse, error)
	UpdateOrder(context.Context, *test.UpdateOrderRequest) (*test.UpdateOrderResponse, error)
	DeleteOrder(context.Context, *test.DeleteOrderRequest) (*test.DeleteOrderResponse, error)
	ListOrders(context.Context, *test.ListOrdersRequest) (*test.ListOrdersResponse, error)
}

type Handler struct {
	test.UnimplementedOrderServiceServer
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h Handler) CreateOrder(ctx context.Context, req *test.CreateOrderRequest) (*test.CreateOrderResponse, error) {
	s, err := h.service.CreateOrder(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("server.CreateOrder %w", err)
	}
	return &test.CreateOrderResponse{Id: s.Id}, nil
}

func (h Handler) GetOrder(ctx context.Context, req *test.GetOrderRequest) (*test.GetOrderResponse, error) {
	order, err := h.service.GetOrder(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("server.GetOrder: %w", err)
	}
	return &test.GetOrderResponse{Order: order.Order}, nil
}

func (h Handler) UpdateOrder(ctx context.Context, req *test.UpdateOrderRequest) (*test.UpdateOrderResponse, error) {
	order, err := h.service.UpdateOrder(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("server.UpdateOrder: %w", err)
	}
	return &test.UpdateOrderResponse{Order: order.Order}, nil
}

func (h Handler) DeleteOrder(ctx context.Context, req *test.DeleteOrderRequest) (*test.DeleteOrderResponse, error) {
	_, err := h.service.DeleteOrder(ctx, req)
	flag := true
	if err != nil {
		flag = false
	}
	return &test.DeleteOrderResponse{Success: flag}, nil
}

func (h Handler) ListOrders(ctx context.Context, req *test.ListOrdersRequest) (*test.ListOrdersResponse, error) {
	orders, err := h.service.ListOrders(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("server.ListOrder: %w", err)
	}

	return &test.ListOrdersResponse{Orders: orders.Orders}, nil
}
