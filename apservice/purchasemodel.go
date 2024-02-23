package apservice

type PurchaseModels struct {
	ArticleId int     `json:"articleId"`
	Date      string  `json:"date"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}

func NewListMapPruchase(articleId, quantity int, date string, total float64) *PurchaseModels {
	return &PurchaseModels{
		ArticleId: articleId,
		Date:      date,
		Quantity:  quantity,
		Total:     total,
	}
}
