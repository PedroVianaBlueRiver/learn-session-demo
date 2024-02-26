package apmodulev1

// Create an array of ArticlePurchase
type ArrayArticlePurchase struct {
	Arraynode []ArticlePurchase
}

// Initialize  struct array with defailt values
func NewMultipleArticlePurchase() *ArrayArticlePurchase {
	return &ArrayArticlePurchase{
		[]ArticlePurchase{
			{
				Name:      "TV",
				Stock:     10,
				Unitprice: 300,
				Date:      "",
				Quantity:  0,
				Total:     0,
			},
			{
				Name:      "Phone",
				Stock:     90,
				Unitprice: 120,
				Date:      "",
				Quantity:  0,
				Total:     0,
			},
		},
	}
}
