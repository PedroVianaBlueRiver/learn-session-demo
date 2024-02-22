package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPerformPurchasev4_MessageValidation(t *testing.T) {
	tests := []struct {
		nametest    string
		quantity    int
		namearticle string
		stock       int
		unitprice   float64
		expected    ArticlePurchaseResponse
	}{
		{
			nametest:    "test #1 error message",
			quantity:    3,
			namearticle: "TV",
			stock:       2,
			unitprice:   300,
			expected:    ArticlePurchaseResponse{issuccess: false, response: "There are just 2(TV) available in stock"},
		},
		{
			nametest:    "test #2 error message",
			quantity:    2,
			namearticle: "TV",
			stock:       0,
			unitprice:   300,
			expected:    ArticlePurchaseResponse{issuccess: false, response: "There is no article available in stock"},
		},
		{
			nametest:    "test #3 success message",
			quantity:    2,
			namearticle: "TV",
			stock:       3,
			unitprice:   300,
			expected:    ArticlePurchaseResponse{issuccess: true, response: "SUCCESS process"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.nametest, func(t *testing.T) {
			ap := NewArticlePurchaseResponse(tt.namearticle, tt.stock, tt.unitprice)
			ap.PerformPurchasev4(tt.quantity)

			if ap.response != tt.expected.response {
				t.Errorf("message received = %v; expected %v", ap.response, tt.expected.response)
			} else {
				assert.Equal(t, ap.response, tt.expected.response)
				assert.Equal(t, ap.issuccess, tt.expected.issuccess)

			}
		})
	}
}

func TestPerformPurchasev4_ArticlePurchaseValidation(t *testing.T) {
	time := time.Now()
	tests := []struct {
		nametest    string
		quantity    int
		namearticle string
		stock       int
		unitprice   float64
		expected    ArticlePurchaseResponse
	}{

		{
			nametest:    "test #1 Article Purchase Validation",
			quantity:    2,
			namearticle: "TV",
			stock:       3,
			unitprice:   300,
			expected: ArticlePurchaseResponse{
				issuccess: true,
				response:  "SUCCESS process",
				at: ArticlePurchase{
					name:      "TV",
					stock:     1,
					unitprice: 300,
					quantity:  2,
					date:      time.Format("2006-01-02"),
					total:     600,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.nametest, func(t *testing.T) {
			ap := NewArticlePurchaseResponse(tt.namearticle, tt.stock, tt.unitprice)
			ap.PerformPurchasev4(tt.quantity)

			if ap.at != tt.expected.at {
				t.Fatalf("received object  = %v; expected %v", ap.at, tt.expected.at)
			}
			assert.Equal(t, ap.at.total, tt.expected.at.total)
			assert.Equal(t, ap.at.stock, tt.expected.at.stock)
			assert.Equal(t, ap.at.date, tt.expected.at.date)
		})
	}
}
