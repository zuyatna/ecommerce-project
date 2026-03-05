package repository

import (
	"log"

	"github.com/zuyatna/ecommerce-project/internal/order/domain"
	"gorm.io/gorm"
)

type OrderModel struct {
	ID         string `gorm:"primaryKey"`
	UserID     string
	TotalPrice string `gorm:"type:decimal(15,2)"`
	Status     string
	Items      []OrderItemModel `gorm:"foreignKey:OrderID"`
}

type OrderItemModel struct {
	ID          string `gorm:"primaryKey"`
	OrderID     string
	ProductID   string
	ProductName string
	Price       string `gorm:"type:decimal(15,2)"`
	Quantity    int
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(order *domain.Order) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		dbOrder := OrderModel{
			ID:         order.ID,
			UserID:     order.UserID,
			TotalPrice: order.TotalPrice.String(),
			Status:     order.Status,
		}

		if err := tx.Create(&dbOrder).Error; err != nil {
			log.Printf("Error creating order, rolling back transaction: %v", err)
			return err
		}

		for _, item := range order.Items {
			result := tx.Exec("UPDATE products SET stock = stock - ? WHERE id = ? AND stock >= ?",
				item.Quantity, item.ProductID, item.Quantity)

			if result.Error != nil {
				log.Printf("Error updating product stock, rolling back transaction: %v", result.Error)
				return result.Error
			}

			if result.RowsAffected == 0 {
				log.Printf("Insufficient stock for product %s, rolling back transaction", item.ProductID)
				return gorm.ErrRecordNotFound
			}
		}

		return nil
	})

	return err
}

func (r *orderRepository) UpdateProductStock(productID string, quantity int) error {
	log.Printf("Restoring stock for product %s by quantity %d", productID, quantity)
	return r.db.Exec("UPDATE products SET stock = stock + ? WHERE id = ?", quantity, productID).Error
}
