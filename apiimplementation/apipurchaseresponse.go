package apiimplementation

type ApiPurchaseReponse struct {
	Purchaseid string   `json:"purchaseid"`
	Total      float64  `json:"total"`
	Articles   []string `json:"articles"`
	Date       string   `json:"date"`
}

func NewApiPurchaseReponse(purchaseid, date string, total float64, articles []string) *ApiPurchaseReponse {
	return &ApiPurchaseReponse{
		Purchaseid: purchaseid,
		Date:       date,
		Total:      total,
		Articles:   articles,
	}
}
