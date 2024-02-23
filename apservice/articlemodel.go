package apservice

import "fmt"

type ArticleModel struct {
	Name      string  `json:"name"`
	Stock     int     `json:"stock"`
	Unitprice float64 `json:"unitprice"`
}

func (am ArticleModel) ValidateStock(quantity int) (bool, string) {
	var issuccess bool
	var response string
	switch {
	case am.Stock == 0:
		issuccess = false
		response = "There is no article available in stock"

	case am.Stock < quantity:
		issuccess = false
		verb := "is"
		if am.Stock > 1 {
			verb = "are"
		}
		response = fmt.Sprintf("There %v just %v(%v) available in stock", verb, am.Stock, am.Name)

	default:
		issuccess = true
		response = ""
	}

	return issuccess, response
}
