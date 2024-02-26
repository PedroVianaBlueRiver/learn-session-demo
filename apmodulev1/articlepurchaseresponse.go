package apmodulev1

type ArticlePurchaseResponse struct {
	Response  string
	Issuccess bool
	AT        ArticlePurchase
}

func NewArticlePurchaseResponse(name string, stock int, unitprice float64) *ArticlePurchaseResponse {
	return &ArticlePurchaseResponse{
		Response:  "SUCCESS process",
		Issuccess: true,
		AT: ArticlePurchase{
			Name:      name,
			Stock:     stock,
			Unitprice: unitprice,
			Date:      "",
			Quantity:  0,
			Total:     0,
		},
	}
}
