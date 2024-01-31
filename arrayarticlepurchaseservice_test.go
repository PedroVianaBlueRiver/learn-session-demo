package main

import (
	"testing"
	"time"
)

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
