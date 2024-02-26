package main

import (
	"purchasetest/apmodulev1"
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
		expected    apmodulev1.ArticlePurchaseResponse
	}{
		{
			nametest:    "test #1 error message",
			quantity:    3,
			namearticle: "TV",
			stock:       2,
			unitprice:   300,
			expected:    apmodulev1.ArticlePurchaseResponse{Issuccess: false, Response: "There are just 2(TV) available in stock"},
		},
		{
			nametest:    "test #2 error message",
			quantity:    2,
			namearticle: "TV",
			stock:       0,
			unitprice:   300,
			expected:    apmodulev1.ArticlePurchaseResponse{Issuccess: false, Response: "There is no article available in stock"},
		},
		{
			nametest:    "test #3 success message",
			quantity:    2,
			namearticle: "TV",
			stock:       3,
			unitprice:   300,
			expected:    apmodulev1.ArticlePurchaseResponse{Issuccess: true, Response: "SUCCESS process"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.nametest, func(t *testing.T) {
			ap := apmodulev1.NewArticlePurchaseResponse(tt.namearticle, tt.stock, tt.unitprice)
			ap.PerformPurchasev4(tt.quantity)

			if ap.Response != tt.expected.Response {
				t.Errorf("message received = %v; expected %v", ap.Response, tt.expected.Response)
			} else {
				assert.Equal(t, ap.Response, tt.expected.Response)
				assert.Equal(t, ap.Issuccess, tt.expected.Issuccess)

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
		expected    apmodulev1.ArticlePurchaseResponse
	}{

		{
			nametest:    "test #1 Article Purchase Validation",
			quantity:    2,
			namearticle: "TV",
			stock:       3,
			unitprice:   300,
			expected: apmodulev1.ArticlePurchaseResponse{
				Issuccess: true,
				Response:  "SUCCESS process",
				AT: apmodulev1.ArticlePurchase{
					Name:      "TV",
					Stock:     1,
					Unitprice: 300,
					Quantity:  2,
					Date:      time.Format("2006-01-02"),
					Total:     600,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.nametest, func(t *testing.T) {
			ap := apmodulev1.NewArticlePurchaseResponse(tt.namearticle, tt.stock, tt.unitprice)
			ap.PerformPurchasev4(tt.quantity)

			if ap.AT != tt.expected.AT {
				t.Fatalf("received object  = %v; expected %v", ap.AT, tt.expected.AT)
			}
			assert.Equal(t, ap.AT.Total, tt.expected.AT.Total)
			assert.Equal(t, ap.AT.Stock, tt.expected.AT.Stock)
			assert.Equal(t, ap.AT.Date, tt.expected.AT.Date)
		})
	}
}
