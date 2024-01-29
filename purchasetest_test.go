package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPerformPurchase(t *testing.T) {
	time := time.Now()

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
			namearticle: "TV",
			stock:       20,
			unitprice:   150,
			expected: ArticlePurchase{
				name:      "TV",
				stock:     18,
				unitprice: 150,
				date:      time.Format("2006-01-02"),
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

func TestPerformPurchasev2(t *testing.T) {

	time := time.Now()
	tests := []struct {
		name     string
		quantity int
		art      ArticlePurchase
		expected ArticlePurchase
	}{
		{
			name:     "PerformPurchasev2 AirFryer",
			quantity: 2,
			art: ArticlePurchase{
				name:      "AirFryer",
				stock:     20,
				unitprice: 150,
				quantity:  0,
				total:     0,
			},
			expected: ArticlePurchase{
				name:      "AirFryer",
				stock:     18,
				unitprice: 150,
				date:      time.Format("2006-01-02"),
				quantity:  2,
				total:     300,
			},
		},
		{
			name:     "PerformPurchasev2 Phone",
			quantity: 3,
			art: ArticlePurchase{
				name:      "Phone",
				stock:     20,
				unitprice: 600,
				quantity:  0,
				total:     0,
			},
			expected: ArticlePurchase{
				name:      "Phone",
				stock:     17,
				unitprice: 600,
				date:      time.Format("2006-01-02"),
				quantity:  3,
				total:     1800,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			art := tt.art
			art.PerformPurchasev2(tt.quantity)

			if art != tt.expected {
				t.Errorf("method PerformPurchasev2(%v) = %v; expected %v", tt.quantity, art, tt.expected)
			} else {

				t.Logf("SUCCESS : PerformPurchasev2 PerformPurchasev2(%v) = %v; expected %v", tt.quantity, art, tt.expected)
			}
		})
	}
}

func TestPerformPurchasev3(t *testing.T) {
	time := time.Now()
	tests := []struct {
		name     string
		quantity int
		article  string
		art      ArrayArticlePurchase
		expected ArrayArticlePurchase
	}{
		{
			name:     "TestPerformPurchasev3 TV",
			quantity: 2,
			article:  "TV",
			art: ArrayArticlePurchase{
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
			},
			expected: ArrayArticlePurchase{
				[]ArticlePurchase{
					{
						name:      "TV",
						stock:     8,
						unitprice: 300,
						quantity:  2,
						date:      time.Format("2006-01-02"),
						total:     600,
					},
					{
						name:      "Phone",
						stock:     90,
						unitprice: 120,
						quantity:  0,
						total:     0,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			art := tt.art
			art.PerformPurchasev3(tt.quantity, tt.article)

			for i := 0; i < len(art.arraynode); i++ {
				if art.arraynode[i] != tt.expected.arraynode[i] {
					t.Errorf("method PerformPurchasev3(%v, %v) = %v; expected %v", tt.article, tt.quantity, art.arraynode[i], tt.expected.arraynode[i])
				} else {
					t.Logf("SUCCESS : method PerformPurchasev3(%v, %v) = %v; expected %v", tt.article, tt.quantity, art.arraynode[i], tt.expected.arraynode[i])
				}
			}

		})
	}

}
