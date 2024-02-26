package apmodulev1

// Create struct ArticlePurchase
type ArticlePurchase struct {
	Name      string
	Stock     int
	Unitprice float64
	Date      string
	Quantity  int
	Total     float64
}

// Initialize a single struct with defailt values
func NewArticlePurchase(name string, stock int, unitprice float64) *ArticlePurchase {
	return &ArticlePurchase{
		Name:      name,
		Stock:     stock,
		Unitprice: unitprice,
		Date:      "",
		Quantity:  0,
		Total:     0,
	}
}
