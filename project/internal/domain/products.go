package domain
type Product struct {
    ID       int64   `json:"id"`
    SellerID int64   `json:"seller_id"`
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    Stock    int     `json:"stock"`
}