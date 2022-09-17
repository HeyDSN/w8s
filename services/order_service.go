package services

import (
	"w8s/models"
	"w8s/repository"
)

type OrderSvc struct {
	OrderRepo repository.OrderRepo
}

type IOrderSvc interface {
	CreateOrder(order *models.Order) error
	GetOrders(orders *[]models.Order) error
	GetOrder(id uint64, order *models.Order) error
	UpdateOrder(id uint64, order *models.Order) error
	DeleteOrder(id uint64) error
}

var IOrderRepo repository.IOrderRepo

func (s *OrderSvc) CreateOrder(order *models.Order) error {
	IOrderRepo = &s.OrderRepo
	return IOrderRepo.CreateOrder(order)
}

func (s *OrderSvc) GetOrders(orders *[]models.Order) error {
	IOrderRepo = &s.OrderRepo
	return IOrderRepo.GetOrders(orders)

}

func (s *OrderSvc) GetOrder(id uint64, order *models.Order) error {
	IOrderRepo = &s.OrderRepo
	return IOrderRepo.GetOrder(id, order)
}

func (s *OrderSvc) UpdateOrder(id uint64, order *models.Order) error {
	IOrderRepo = &s.OrderRepo
	return IOrderRepo.UpdateOrder(id, order)
}

func (s *OrderSvc) DeleteOrder(id uint64) error {
	IOrderRepo = &s.OrderRepo
	return IOrderRepo.DeleteOrder(id)
}
