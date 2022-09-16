package services

import (
	"w8s/models"
	"w8s/repository"
)

type OrderSvc struct {
	OrderRepo repository.OrderRepo
}

func (s *OrderSvc) CreateOrder(order *models.Order) error {
	return s.OrderRepo.CreateOrder(order)
}

func (s *OrderSvc) GetOrders(orders *[]models.Order) error {
	return s.OrderRepo.GetOrders(orders)
}

func (s *OrderSvc) GetOrder(id uint64, order *models.Order) error {
	return s.OrderRepo.GetOrder(id, order)
}

func (s *OrderSvc) UpdateOrder(id uint64, order *models.Order) error {
	return s.OrderRepo.UpdateOrder(id, order)
}

func (s *OrderSvc) DeleteOrder(id uint64) error {
	return s.OrderRepo.DeleteOrder(id)
}
