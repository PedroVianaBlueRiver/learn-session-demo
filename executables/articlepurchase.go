package main

// Create struct ArticlePurchase
type ArticlePurchase struct {
	name      string
	stock     int
	unitprice float64
	date      string
	quantity  int
	total     float64
}

// Initialize a single struct with defailt values
func NewArticlePurchase(name string, stock int, unitprice float64) *ArticlePurchase {
	return &ArticlePurchase{
		name:      name,
		stock:     stock,
		unitprice: unitprice,
		date:      "",
		quantity:  0,
		total:     0,
	}
}
