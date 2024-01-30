package main

// Create an array of ArticlePurchase
type ArrayArticlePurchase struct {
	arraynode []ArticlePurchase
}

// Initialize  struct array with defailt values
func NewMultipleArticlePurchase() *ArrayArticlePurchase {
	return &ArrayArticlePurchase{
		[]ArticlePurchase{
			{
				name:      "TV",
				stock:     10,
				unitprice: 300,
				date:      "",
				quantity:  0,
				total:     0,
			},
			{
				name:      "Phone",
				stock:     90,
				unitprice: 120,
				date:      "",
				quantity:  0,
				total:     0,
			},
		},
	}
}
