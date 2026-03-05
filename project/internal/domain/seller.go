package domain

type Seller struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	ShopName    string `json:"shop_name"`
	Description string `json:"description"`
	Phone       string `json:"phone"`
}
