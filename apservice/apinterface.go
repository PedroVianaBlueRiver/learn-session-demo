package apservice

type APInterface interface {
	CreatePurchase(id, quantity int) (bool, string)
}
