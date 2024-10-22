package domain

type Order struct {
	Id        int    `json:"id"`
	ProductId int    `json:"product_id"`
	Amount    int    `json:"amount"`
	StatusId  int    `json:"status_id"`
	UserId    int    `json:"user_id"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
