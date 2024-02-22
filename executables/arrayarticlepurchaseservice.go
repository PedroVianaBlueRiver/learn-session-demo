package main

import (
	"time"
)

// method to performed a purchase processs and update existing stock in base on the array (using receiver param)
func (aap *ArrayArticlePurchase) PerformPurchasev3(quantity int, article string) {
	time := time.Now()

	for i := 0; i < len(aap.arraynode); i++ {
		attr := &aap.arraynode[i]
		if attr.name == article {
			attr.quantity = quantity
			attr.date = time.Format("2006-01-02")
			attr.stock = attr.stock - quantity
			attr.total = attr.unitprice * float64(quantity)
		}
	}

}
