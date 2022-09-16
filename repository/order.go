package repository

import (
	"w8s/models"

	"gorm.io/gorm"
)

type OrderRepo struct {
	Conn *gorm.DB
}

func (r *OrderRepo) CreateOrder(order *models.Order) error {
	return r.Conn.Create(order).Error
}

func (r *OrderRepo) GetOrders(orders *[]models.Order) error {
	return r.Conn.Preload("Items").Find(orders).Error
}

func (r *OrderRepo) GetOrder(id uint64, order *models.Order) error {
	return r.Conn.Preload("Items").Where("order_id = ?", id).First(&order).Error
}

func (r *OrderRepo) UpdateOrder(id uint64, order *models.Order) error {
	return r.Conn.Transaction(func(tx *gorm.DB) error {
		// update order
		if err := tx.Model(&models.Order{}).Where("order_id = ?", id).Updates(order).Error; err != nil {
			return err
		}
		// update items
		// delete old data items
		if err := tx.Model(&models.Item{}).Where("order_id = ?", id).Delete(&models.Item{}).Error; err != nil {
			return err
		}
		// insert new data items
		for _, item := range order.Items {
			item.OrderID = id
			if err := tx.Create(&item).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *OrderRepo) DeleteOrder(id uint64) error {
	return r.Conn.Transaction(func(tx *gorm.DB) error {
		// delete items
		if err := tx.Model(&models.Item{}).Where("order_id = ?", id).Delete(&models.Item{}).Error; err != nil {
			return err
		}
		// delete order
		if err := tx.Where("order_id = ?", id).Delete(&models.Order{}).Error; err != nil {
			return err
		}
		return nil
	})
}
