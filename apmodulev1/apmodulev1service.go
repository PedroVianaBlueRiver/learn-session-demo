package apmodulev1

import (
	"time"
)

// method to performed a purchase processs and update existing stock
func PerformPurchase(quantity int, purchase *ArticlePurchase) ArticlePurchase {
	time := time.Now()
	purchase.Date = time.Format("2006-01-02")
	purchase.Quantity = quantity
	purchase.Stock = purchase.Stock - quantity
	purchase.Total = purchase.Unitprice * float64(quantity)
	return *purchase
}

// method to performed a purchase processs and update existing stock (using receiver param)
func (ap *ArticlePurchase) PerformPurchasev2(quantity int) {
	time := time.Now()
	ap.Date = time.Format("2006-01-02")
	ap.Quantity = quantity
	ap.Stock = ap.Stock - quantity
	ap.Total = ap.Unitprice * float64(quantity)
}
