package main

type ArticlePurchaseResponse struct {
	response  string
	issuccess bool
	at        ArticlePurchase
}

func NewArticlePurchaseResponse(name string, stock int, unitprice float64) *ArticlePurchaseResponse {
	return &ArticlePurchaseResponse{
		response:  "SUCCESS process",
		issuccess: true,
		at: ArticlePurchase{
			name:      name,
			stock:     stock,
			unitprice: unitprice,
			date:      "",
			quantity:  0,
			total:     0,
		},
	}
}
