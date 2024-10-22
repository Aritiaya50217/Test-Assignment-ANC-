package domain

type History struct {
	Id        int    `json:"id"`
	OrderId   int    `json:"order_id"`
	StatusId  int    `json:"status_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
