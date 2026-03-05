package domain

import "time"

type Order struct {
    ID        int64     `json:"id"`
    UserID    int64     `json:"user_id"`
    Status    string    `json:"status"`
    CreatedAt time.Time `json:"created_at"`
}

type OrderItem struct {
    OrderID   int64   `json:"order_id"`
    ProductID int64   `json:"product_id"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}