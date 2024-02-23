package main

import (
	"purchasetest/apservice"
	"purchasetest/mapstructmodel"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetItem(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		expectedmsn string
		expectedok  bool
		expected    apservice.ArticleModel
	}{
		{
			name: "test GetItem #1 success",
			id:   1,
			expected: apservice.ArticleModel{
				Name:      "TV",
				Stock:     4,
				Unitprice: 130,
			},
			expectedmsn: "",
			expectedok:  true,
		},
		{
			name:        "test GetItem #2 error",
			id:          70,
			expected:    apservice.ArticleModel{},
			expectedmsn: "article with id: 70 doesn't exist",
			expectedok:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articleList := mapstructmodel.NewListMapArticle()
			actual, ok, msn := apservice.GetItem(articleList, tt.id)
			assert.Equal(t, ok, tt.expectedok)
			assert.Equal(t, msn, tt.expectedmsn)
			if actual != tt.expected {
				t.Errorf("there was an error during the test actual = %v; expected %v", actual, tt.expected)
			}

		})
	}

}

func TestAddItem(t *testing.T) {
	tests := []struct {
		name        string
		expectedmsn string
		expectedok  bool
		id          int
		Stock       int
		namet       string
		Unitprice   float64
		expected    map[int]apservice.ArticleModel
	}{
		{
			name:        "test AddItem #1 success",
			expectedmsn: "",
			expectedok:  true,
			Stock:       20,
			id:          6,
			namet:       "CPU",
			Unitprice:   1300,
			expected: map[int]apservice.ArticleModel{
				1: {
					Name:      "TV",
					Stock:     4,
					Unitprice: 130,
				},
				2: {
					Name:      "Phone",
					Stock:     5,
					Unitprice: 210,
				},
				4: {
					Name:      "AirFryer",
					Stock:     30,
					Unitprice: 190,
				},
				6: {
					Name:      "CPU",
					Stock:     20,
					Unitprice: 1300,
				},
			},
		},
		{
			name:        "test AddItem #2 error",
			expectedmsn: "article with id: 1 already exists",
			expectedok:  false,
			id:          1,
			expected: map[int]apservice.ArticleModel{
				1: {
					Name:      "TV",
					Stock:     4,
					Unitprice: 130,
				},
				2: {
					Name:      "Phone",
					Stock:     5,
					Unitprice: 210,
				},
				4: {
					Name:      "AirFryer",
					Stock:     30,
					Unitprice: 190,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articleList := mapstructmodel.NewListMapArticle()
			actual, ok, msn := apservice.AddItem(articleList, tt.Stock, tt.id, tt.namet, tt.Unitprice)
			assert.Equal(t, ok, tt.expectedok)
			assert.Equal(t, msn, tt.expectedmsn)
			if actual == nil {
				t.Errorf("there was an error during the test actual = %v; expected %v", actual, tt.expected)
			}
			if len(actual) != len(tt.expected) {
				t.Errorf("there was an error during the test len expected, actual = %v; expected %v", actual, tt.expected)
			}

		})
	}

}

func TestDeleteItem(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		expectedmsn string
		expectedok  bool
		expected    map[int]apservice.ArticleModel
	}{
		{
			name:        "test DeleteItem #1 success",
			id:          4,
			expectedmsn: "",
			expectedok:  true,
			expected: map[int]apservice.ArticleModel{
				1: {
					Name:      "TV",
					Stock:     4,
					Unitprice: 130,
				},
				2: {
					Name:      "Phone",
					Stock:     5,
					Unitprice: 210,
				},
			},
		},
		{
			name:        "test DeleteItem #2 error",
			id:          7,
			expectedmsn: "article with id: 7 doesn't exist",
			expectedok:  false,
			expected: map[int]apservice.ArticleModel{
				1: {
					Name:      "TV",
					Stock:     4,
					Unitprice: 130,
				},
				2: {
					Name:      "Phone",
					Stock:     5,
					Unitprice: 210,
				},
				4: {
					Name:      "AirFryer",
					Stock:     30,
					Unitprice: 190,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articleList := mapstructmodel.NewListMapArticle()
			actual, ok, msn := apservice.DeleteItem(articleList, tt.id)
			assert.Equal(t, ok, tt.expectedok)
			assert.Equal(t, msn, tt.expectedmsn)

			if len(actual) != len(tt.expected) {
				t.Errorf("there was an error during the test len expected, actual = %v; expected %v", actual, tt.expected)
			}

		})
	}

}

func TestUpdateItem(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		expectedmsn string
		expectedok  bool
		nameat      string
		Stock       int
		Unitprice   float64
		expected    apservice.ArticleModel
	}{
		{
			name:        "test updateItem #1 success",
			id:          1,
			expectedmsn: "",
			expectedok:  true,
			nameat:      "Smart TV",
			Stock:       90,
			Unitprice:   1200,
			expected: apservice.ArticleModel{
				Name:      "Smart TV",
				Stock:     90,
				Unitprice: 1200,
			},
		},
		{
			name:        "test updateItem #2 error",
			id:          9,
			expectedmsn: "article with id: 9 doesn't exist",
			expectedok:  false,
			expected:    apservice.ArticleModel{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articleList := mapstructmodel.NewListMapArticle()
			actual, ok, msn := apservice.UpdateItem(articleList, tt.id, tt.Stock, tt.nameat, tt.Unitprice)
			assert.Equal(t, ok, tt.expectedok)
			assert.Equal(t, msn, tt.expectedmsn)
			actualItem := actual[tt.id]
			if actualItem != tt.expected {
				t.Errorf("there was an error during the test  expected, actualItem = %v; expected %v", actualItem, tt.expected)
			}

		})
	}

}

func TestCreatePurchase(t *testing.T) {
	time := time.Now()
	tests := []struct {
		name             string
		id               int
		expectedmsn      string
		expectedok       bool
		quantity         int
		expectedarticle  apservice.ArticleModel
		expectedpurchase apservice.PurchaseModels
	}{
		{
			name:        "test CreatePurchase #1 success",
			id:          1,
			expectedmsn: "",
			expectedok:  true,
			quantity:    2,
			expectedarticle: apservice.ArticleModel{
				Name:      "TV",
				Stock:     2,
				Unitprice: 130,
			},
			expectedpurchase: apservice.PurchaseModels{
				ArticleId: 1,
				Date:      time.Format("2006-01-02"),
				Quantity:  2,
				Total:     260,
			},
		},
		{
			name:             "test CreatePurchase #2 error",
			id:               30,
			expectedmsn:      "article with id: 30 doesn't exist",
			expectedok:       false,
			quantity:         2,
			expectedarticle:  apservice.ArticleModel{},
			expectedpurchase: apservice.PurchaseModels{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := apservice.NewAPModelRespose(mapstructmodel.NewListMapArticle(), apservice.PurchaseModels{})
			ok, msn := apservice.CreatePurchase(tt.id, tt.quantity, response)
			assert.Equal(t, ok, tt.expectedok)
			assert.Equal(t, msn, tt.expectedmsn)
			actualItem := response.Nat[tt.id]
			if actualItem != tt.expectedarticle {
				t.Errorf("there was an error during the test  expected, actualItem = %v; expectedarticle %v", actualItem, tt.expectedarticle)
			}
			if *response.Npm != tt.expectedpurchase {
				t.Errorf("there was an error during the test  expected, actualpm = %v; expectedpurchase %v", *response.Npm, tt.expectedpurchase)
			}

		})
	}
}
