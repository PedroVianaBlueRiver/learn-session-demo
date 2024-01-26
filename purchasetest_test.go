package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
				unitprice: 150,
				quantity:  2,
				total:     300,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ap := NewArticlePurchase(tt.namearticle, tt.stock, tt.unitprice)
			actual := PerformPurchase(tt.quantity, ap)

			if actual != tt.expected {
				t.Errorf("method PerformPurchase() = %v; expected %v", actual, tt.expected)
			} else {
				assert.Equal(t, actual.name, tt.expected.name)
				assert.Equal(t, actual.stock, tt.expected.stock)
				assert.Equal(t, actual.total, tt.expected.total)
				t.Logf("SUCCESS : method PerformPurchase() = %v; expected %v", actual, tt.expected)
			}
		})
	}

}
