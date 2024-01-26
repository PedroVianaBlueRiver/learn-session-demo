package main

import (
	"testing"
)

func TestPerformPurchase(t *testing.T) {
	tests := []struct {
		name        string
		quantity    int
		namearticle string
		stock       int
		unitprice   float64
		expected    ArticlePurchase
	}{
		{
			name:        "Purchase a tv",
			quantity:    2,
			namearticle: "AirFryer",
			stock:       20,
			unitprice:   150,
			expected: ArticlePurchase{
				name:      "AirFryer",
				stock:     18,
				unitprice: 300,
				quantity:  2,
				total:     600,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//ap := NewArticlePurchase(tt.namearticle, tt.stock, tt.unitprice)
			//actual := PerformPurchase(tt.quantity)

			// if actual != tt.expected {
			// 	t.Errorf("method Drive() = %v; expected %v", actual, tt.expected)
			// } else {
			// 	t.Logf("SUCCESS : method PerformPurchase() = %v; expected %v", actual, tt.expected)
			// }
		})
	}

}
