package main

import (
	"purchasetest/apmodulev1"
	"testing"
	"time"
)

func TestPerformPurchasev3(t *testing.T) {
	time := time.Now()
	tests := []struct {
		name     string
		quantity int
		article  string
		art      apmodulev1.ArrayArticlePurchase
		expected apmodulev1.ArrayArticlePurchase
	}{
		{
			name:     "TestPerformPurchasev3 TV",
			quantity: 2,
			article:  "TV",
			art: apmodulev1.ArrayArticlePurchase{
				Arraynode: []apmodulev1.ArticlePurchase{
					{
						Name:      "TV",
						Stock:     10,
						Unitprice: 300,
						Quantity:  0,
						Total:     0,
					},
					{
						Name:      "Phone",
						Stock:     90,
						Unitprice: 120,
						Quantity:  0,
						Total:     0,
					},
				},
			},
			expected: apmodulev1.ArrayArticlePurchase{
				Arraynode: []apmodulev1.ArticlePurchase{
					{
						Name:      "TV",
						Stock:     8,
						Unitprice: 300,
						Quantity:  2,
						Date:      time.Format("2006-01-02"),
						Total:     600,
					},
					{
						Name:      "Phone",
						Stock:     90,
						Unitprice: 120,
						Quantity:  0,
						Total:     0,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			art := tt.art
			art.PerformPurchasev3(tt.quantity, tt.article)

			for i := 0; i < len(art.Arraynode); i++ {
				if art.Arraynode[i] != tt.expected.Arraynode[i] {
					t.Errorf("method PerformPurchasev3(%v, %v) = %v; expected %v", tt.article, tt.quantity, art.Arraynode[i], tt.expected.Arraynode[i])
				} else {
					t.Logf("SUCCESS : method PerformPurchasev3(%v, %v) = %v; expected %v", tt.article, tt.quantity, art.Arraynode[i], tt.expected.Arraynode[i])
				}
			}

		})
	}

}
