package apmodulev1

import (
	"time"
)

// method to performed a purchase processs and update existing stock in base on the array (using receiver param)
func (aap *ArrayArticlePurchase) PerformPurchasev3(quantity int, article string) {
	time := time.Now()

	for i := 0; i < len(aap.Arraynode); i++ {
		attr := &aap.Arraynode[i]
		if attr.Name == article {
			attr.Quantity = quantity
			attr.Date = time.Format("2006-01-02")
			attr.Stock = attr.Stock - quantity
			attr.Total = attr.Unitprice * float64(quantity)
		}
	}

}
