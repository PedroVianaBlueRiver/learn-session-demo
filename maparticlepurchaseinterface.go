package main

type APInterface interface {
	createPurchase(id, quantity int) (bool, string)
}
