package main

import (
	"purchasetest/apmodulev1"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPerformPurchase(t *testing.T) {
	time := time.Now()

	tests := []struct {
		name        string
		Quantity    int
		namearticle string
		Stock       int
		Unitprice   float64
		expected    apmodulev1.ArticlePurchase
	}{
		{
			name:        "Purchase a tv",
			Quantity:    2,
			namearticle: "TV",
			Stock:       20,
			Unitprice:   150,
			expected: apmodulev1.ArticlePurchase{
				Name:      "TV",
				Stock:     18,
				Unitprice: 150,
				Date:      time.Format("2006-01-02"),
				Quantity:  2,
				Total:     300,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ap := apmodulev1.NewArticlePurchase(tt.namearticle, tt.Stock, tt.Unitprice)
			actual := apmodulev1.PerformPurchase(tt.Quantity, ap)

			if actual != tt.expected {
				t.Errorf("method PerformPurchase() = %v; expected %v", actual, tt.expected)
			} else {
				assert.Equal(t, actual.Name, tt.expected.Name)
				assert.Equal(t, actual.Stock, tt.expected.Stock)
				assert.Equal(t, actual.Total, tt.expected.Total)
				t.Logf("SUCCESS : method PerformPurchase() = %v; expected %v", actual, tt.expected)
			}
		})
	}

}

func TestPerformPurchasev2(t *testing.T) {

	time := time.Now()
	tests := []struct {
		name     string
		Quantity int
		art      apmodulev1.ArticlePurchase
		expected apmodulev1.ArticlePurchase
	}{
		{
			name:     "PerformPurchasev2 AirFryer",
			Quantity: 2,
			art: apmodulev1.ArticlePurchase{
				Name:      "AirFryer",
				Stock:     20,
				Unitprice: 150,
				Quantity:  0,
				Total:     0,
			},
			expected: apmodulev1.ArticlePurchase{
				Name:      "AirFryer",
				Stock:     18,
				Unitprice: 150,
				Date:      time.Format("2006-01-02"),
				Quantity:  2,
				Total:     300,
			},
		},
		{
			name:     "PerformPurchasev2 Phone",
			Quantity: 3,
			art: apmodulev1.ArticlePurchase{
				Name:      "Phone",
				Stock:     20,
				Unitprice: 600,
				Quantity:  0,
				Total:     0,
			},
			expected: apmodulev1.ArticlePurchase{
				Name:      "Phone",
				Stock:     17,
				Unitprice: 600,
				Date:      time.Format("2006-01-02"),
				Quantity:  3,
				Total:     1800,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			art := tt.art
			art.PerformPurchasev2(tt.Quantity)

			if art != tt.expected {
				t.Errorf("method PerformPurchasev2(%v) = %v; expected %v", tt.Quantity, art, tt.expected)
			} else {

				t.Logf("SUCCESS : PerformPurchasev2 PerformPurchasev2(%v) = %v; expected %v", tt.Quantity, art, tt.expected)
			}
		})
	}
}
