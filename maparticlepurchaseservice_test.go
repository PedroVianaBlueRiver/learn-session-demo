package main

import (
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
		expected    ArticleModel
	}{
		{
			name: "test GetItem #1 success",
			id:   1,
			expected: ArticleModel{
				name:      "TV",
				stock:     4,
				unitprice: 130,
			},
			expectedmsn: "",
			expectedok:  true,
		},
		{
			name:        "test GetItem #2 error",
			id:          70,
			expected:    ArticleModel{},
			expectedmsn: "article with id: 70 doesn't exist",
			expectedok:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articleList := NewListMapArticle()
			actual, ok, msn := GetItem(articleList, tt.id)
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
		stock       int
		namet       string
		unitprice   float64
		expected    map[int]ArticleModel
	}{
		{
			name:        "test AddItem #1 success",
			expectedmsn: "",
			expectedok:  true,
			stock:       20,
			id:          6,
			namet:       "CPU",
			unitprice:   1300,
			expected: map[int]ArticleModel{
				1: {
					name:      "TV",
					stock:     4,
					unitprice: 130,
				},
				2: {
					name:      "Phone",
					stock:     5,
					unitprice: 210,
				},
				4: {
					name:      "AirFryer",
					stock:     30,
					unitprice: 190,
				},
				6: {
					name:      "CPU",
					stock:     20,
					unitprice: 1300,
				},
			},
		},
		{
			name:        "test AddItem #2 error",
			expectedmsn: "article with id: 1 already exists",
			expectedok:  false,
			id:          1,
			expected: map[int]ArticleModel{
				1: {
					name:      "TV",
					stock:     4,
					unitprice: 130,
				},
				2: {
					name:      "Phone",
					stock:     5,
					unitprice: 210,
				},
				4: {
					name:      "AirFryer",
					stock:     30,
					unitprice: 190,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articleList := NewListMapArticle()
			actual, ok, msn := AddItem(articleList, tt.stock, tt.id, tt.namet, tt.unitprice)
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
		expected    map[int]ArticleModel
	}{
		{
			name:        "test DeleteItem #1 success",
			id:          4,
			expectedmsn: "",
			expectedok:  true,
			expected: map[int]ArticleModel{
				1: {
					name:      "TV",
					stock:     4,
					unitprice: 130,
				},
				2: {
					name:      "Phone",
					stock:     5,
					unitprice: 210,
				},
			},
		},
		{
			name:        "test DeleteItem #2 error",
			id:          7,
			expectedmsn: "article with id: 7 doesn't exist",
			expectedok:  false,
			expected: map[int]ArticleModel{
				1: {
					name:      "TV",
					stock:     4,
					unitprice: 130,
				},
				2: {
					name:      "Phone",
					stock:     5,
					unitprice: 210,
				},
				4: {
					name:      "AirFryer",
					stock:     30,
					unitprice: 190,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articleList := NewListMapArticle()
			actual, ok, msn := DeleteItem(articleList, tt.id)
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
		stock       int
		unitprice   float64
		expected    ArticleModel
	}{
		{
			name:        "test updateItem #1 success",
			id:          1,
			expectedmsn: "",
			expectedok:  true,
			nameat:      "Smart TV",
			stock:       90,
			unitprice:   1200,
			expected: ArticleModel{
				name:      "Smart TV",
				stock:     90,
				unitprice: 1200,
			},
		},
		{
			name:        "test updateItem #2 error",
			id:          9,
			expectedmsn: "article with id: 9 doesn't exist",
			expectedok:  false,
			expected:    ArticleModel{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articleList := NewListMapArticle()
			actual, ok, msn := updateItem(articleList, tt.id, tt.stock, tt.nameat, tt.unitprice)
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
		expectedarticle  ArticleModel
		expectedpurchase PurchaseModels
	}{
		{
			name:        "test CreatePurchase #1 success",
			id:          1,
			expectedmsn: "",
			expectedok:  true,
			quantity:    2,
			expectedarticle: ArticleModel{
				name:      "TV",
				stock:     2,
				unitprice: 130,
			},
			expectedpurchase: PurchaseModels{
				articleId: 1,
				date:      time.Format("2006-01-02"),
				quantity:  2,
				total:     260,
			},
		},
		{
			name:             "test CreatePurchase #2 error",
			id:               30,
			expectedmsn:      "article with id: 30 doesn't exist",
			expectedok:       false,
			quantity:         2,
			expectedarticle:  ArticleModel{},
			expectedpurchase: PurchaseModels{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := NewAPModelRespose(NewListMapArticle(), PurchaseModels{})
			ok, msn := implementcreatePurchase(tt.id, tt.quantity, response)
			assert.Equal(t, ok, tt.expectedok)
			assert.Equal(t, msn, tt.expectedmsn)
			actualItem := response.at[tt.id]
			if actualItem != tt.expectedarticle {
				t.Errorf("there was an error during the test  expected, actualItem = %v; expectedarticle %v", actualItem, tt.expectedarticle)
			}
			if *response.pm != tt.expectedpurchase {
				t.Errorf("there was an error during the test  expected, actualpm = %v; expectedpurchase %v", *response.pm, tt.expectedpurchase)
			}

		})
	}
}
