package orderrepository

import (
	"context"
	"fmt"
	"sync"

	test "example.com/m/pkg/api/order/api"
)

type Repository struct {
	mu     sync.Mutex
	orders map[string]*test.Order
}

func NewRepository() *Repository {
	return &Repository{
		orders: make(map[string]*test.Order),
		mu:     sync.Mutex{},
	}
}

func (r *Repository) Create(ctx context.Context, req *test.Order) (string, error) {
	orderID := string(len(r.orders) + 1)
	r.orders[orderID] = &test.Order{
		Id:       orderID,
		Item:     req.Item,
		Quantity: req.Quantity,
	}
	return orderID, nil
}

func (r *Repository) Get(ctx context.Context, id string) (*test.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	orderID := id
	order, exists := r.orders[orderID]
	if !exists {
		return nil, fmt.Errorf("orderrepository.Get: doesnt exists")
	}
	return order, nil
}

func (r *Repository) Update(ctx context.Context, req *test.Order) (*test.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	orderID := string(req.Id)

	order := &test.Order{
		Id:       orderID,
		Item:     req.Item,
		Quantity: req.Quantity,
	}

	r.orders[orderID] = order

	return order, nil
}

func (r *Repository) Delete(ctx context.Context, id string) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	orderID := id
	delete(r.orders, orderID)

	return true, nil
}

func (r *Repository) List(ctx context.Context) ([]*test.Order, error) {
	var orders []*test.Order
	for _, order := range r.orders {
		orders = append(orders, order)
	}

	return orders, nil
}
