package main

import (
	"time"
)

// method to performed a purchase processs and update existing stock
func PerformPurchase(quantity int, purchase *ArticlePurchase) ArticlePurchase {
	time := time.Now()
	purchase.date = time.Format("2006-01-02")
	purchase.quantity = quantity
	purchase.stock = purchase.stock - quantity
	purchase.total = purchase.unitprice * float64(quantity)
	return *purchase
}

// method to performed a purchase processs and update existing stock (using receiver param)
func (ap *ArticlePurchase) PerformPurchasev2(quantity int) {
	time := time.Now()
	ap.date = time.Format("2006-01-02")
	ap.quantity = quantity
	ap.stock = ap.stock - quantity
	ap.total = ap.unitprice * float64(quantity)
}
