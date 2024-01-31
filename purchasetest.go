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

//func main() {

//Create new ArticlePurchase
// time1 := time.Now()
// fmt.Println(time1.Format("2006-01-01 "))
// ap := NewArticlePurchase("TV", 3, 300)
// fmt.Printf("ArticlePurchase struct %v", ap)
// fmt.Println()
// p1 := PerformPurchase(3, ap)
// fmt.Printf("PerformPurchase() result = %v", p1)
// fmt.Println()

// //method with a receiver argument
// p2 := NewArticlePurchase("TV", 3, 300)
// fmt.Printf("ArticlePurchase struct %v", p2)
// fmt.Println()
// p2.PerformPurchasev2(2)
// fmt.Printf("PerformPurchasev2() result = %v", p2)
// fmt.Println()

// arrayp := NewMultipleArticlePurchase()
// arrayp.PerformPurchasev3(3, "Phone")
// fmt.Println(*arrayp)

//}
