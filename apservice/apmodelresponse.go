package apservice

type APModelRespose struct {
	Nat map[int]ArticleModel `json:"articlemodel"`
	Npm *PurchaseModels      `json:"purchasemodel"`
}

func NewAPModelRespose(at map[int]ArticleModel, pm PurchaseModels) *APModelRespose {
	return &APModelRespose{
		Nat: at,
		Npm: &pm,
	}
}
