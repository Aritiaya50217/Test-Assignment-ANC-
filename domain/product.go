package domain

type Product struct {
	Id        int64  `json:"id"`
	Style     string `json:"style"`
	GenderId  int64  `json:"gender_id"`
	SizeId    int64  `json:"size_id"`
	ColorId   int64  `json:"color_id"`
	Price     string `json:"price"`
	PatternId int64  `json:"pattern_id"`
	FigureId  int64  `json:"figure_id"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

		