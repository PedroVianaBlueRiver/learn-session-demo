package main

import (
	"fmt"
)

type ArticlePurchase struct {
	name      string
	stock     int
	unitprice float64
	quantity  int
	total     float64
}

type ArrayArticlePurchase struct {
	arraynode []ArticlePurchase
}

func NewArticlePurchase(name string, stock int, unitprice float64) *ArticlePurchase {
	return &ArticlePurchase{
		name:      name,
		stock:     stock,
		unitprice: unitprice,
		quantity:  0,
		total:     0,
	}
}

func NewMultipleArticlePurchase() *ArrayArticlePurchase {
	return &ArrayArticlePurchase{
		[]ArticlePurchase{
			{
				name:      "TV",
				stock:     10,
				unitprice: 300,
				quantity:  0,
				total:     0,
			},
			{
				name:      "Phone",
				stock:     90,
				unitprice: 120,
				quantity:  0,
				total:     0,
			},
		},
	}
}

func PerformPurchase(quantity int, purchase *ArticlePurchase) ArticlePurchase {
	purchase.quantity = quantity
	purchase.stock = purchase.stock - quantity
	purchase.total = purchase.unitprice * float64(quantity)
	return *purchase
}

func (ap *ArticlePurchase) PerformPurchasev2(quantity int) {
	ap.quantity = quantity
	ap.stock = ap.stock - quantity
	ap.total = ap.unitprice * float64(quantity)
}

func (aap *ArrayArticlePurchase) PerformPurchasev3(quantity int, article string) {

	for i := 0; i < len(aap.arraynode); i++ {
		attr := &aap.arraynode[i]
		if attr.name == article {
			attr.quantity = quantity
			attr.stock = attr.stock - quantity
			attr.total = attr.unitprice * float64(quantity)
		}
	}

}

func main() {

	//Create new ArticlePurchase
	ap := NewArticlePurchase("TV", 3, 300)
	fmt.Printf("ArticlePurchase struct %v", ap)
	fmt.Println()
	p1 := PerformPurchase(3, ap)
	fmt.Printf("PerformPurchase() result = %v", p1)
	fmt.Println()

	//method with a receiver argument
	p2 := NewArticlePurchase("TV", 3, 300)
	fmt.Printf("ArticlePurchase struct %v", p2)
	fmt.Println()
	p2.PerformPurchasev2(2)
	fmt.Printf("PerformPurchasev2() result = %v", p2)
	fmt.Println()

	arrayp := NewMultipleArticlePurchase()
	arrayp.PerformPurchasev3(3, "Phone")
	fmt.Println(*arrayp)
}
