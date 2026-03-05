package repository

import (
	"database/sql"
	"errors"

	"github.com/l0ng7h0r/internal/domain"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}


func (r *OrderRepository) CreateOrders(order *domain.Order, items []*domain.OrderItem) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    err = tx.QueryRow(
        `INSERT INTO orders (user_id, status) VALUES ($1, $2) RETURNING id, status, created_at`,
        order.UserID, "pending",
    ).Scan(&order.ID, &order.Status, &order.CreatedAt)
    if err != nil {
        return err
    }

    for _, item := range items {
        // เช็ค stock ก่อนว่าพอไหม
        var stock int
        err := tx.QueryRow(
            `SELECT stock FROM products WHERE id = $1`, item.ProductID,
        ).Scan(&stock)
        if err != nil {
            return err
        }
        if stock < item.Quantity {
            return errors.New("insufficient stock")
        }

        // insert order_items
        _, err = tx.Exec(
            `INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)`,
            order.ID, item.ProductID, item.Quantity, item.Price,
        )
        if err != nil {
            return err
        }

        // ลด stock
        _, err = tx.Exec(
            `UPDATE products SET stock = stock - $1 WHERE id = $2`,
            item.Quantity, item.ProductID,
        )
        if err != nil {
            return err
        }
    }

    return tx.Commit()
}

func (r *OrderRepository) GetOrdersByID(id int64) (*domain.Order, []*domain.OrderItem, error) {
    // ดึง order หลักก่อน
    order := &domain.Order{}
    err := r.db.QueryRow(
        `SELECT id, user_id, status, created_at FROM orders WHERE id = $1`, id,
    ).Scan(&order.ID, &order.UserID, &order.Status, &order.CreatedAt)
    if err != nil {
        return nil, nil, err
    }

    // แล้วค่อยดึง items
    var orderItems []*domain.OrderItem
    rows, err := r.db.Query(
        `SELECT order_id, product_id, quantity, price FROM order_items WHERE order_id = $1`, id,
    )
    if err != nil {
        return nil, nil, err
    }
    defer rows.Close()

    for rows.Next() {
        orderItem := &domain.OrderItem{}
        err := rows.Scan(&orderItem.OrderID, &orderItem.ProductID, &orderItem.Quantity, &orderItem.Price)
        if err != nil {
            return nil, nil, err
        }
        orderItems = append(orderItems, orderItem)
    }

    return order, orderItems, nil
}

func (r *OrderRepository)GetOrdersByUserId(id int64) ([]*domain.Order,error) {
	var orders []*domain.Order

	rows, err := r.db.Query(`SELECT id, user_id, status, created_at FROM orders WHERE user_id = $1`, id)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		order := &domain.Order{}
		err := rows.Scan(&order.ID, &order.UserID, &order.Status, &order.CreatedAt)
		if err != nil{
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *OrderRepository)UpdateOrders(order *domain.Order) error {
	_, err := r.db.Exec(`UPDATE orders SET status = $1 WHERE id = $2`, order.Status, order.ID)
	if err != nil{
		return err
	}
	return nil
}

func (r *OrderRepository)DeleteOrders(id int64) error {
	tx, err := r.db.Begin()
	if err != nil{
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`DELETE FROM order_items WHERE order_id = $1`, id)
	if err != nil{
		return err
	}
	_, err = tx.Exec(`DELETE FROM orders WHERE id = $1`, id)
	if err != nil{
		return err
	}
	return tx.Commit()
}