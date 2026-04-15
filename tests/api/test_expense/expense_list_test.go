package test_expense

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tmazitov/ayda-order-service.git/tests/api/utils"
)

type ListFilters struct {
	Name       string `query:"name"`
	Page       int    `query:"page"`
	Limit      int    `query:"limit"`
	Date       string `query:"date"`
	CategoryId string `query:"category"`
}

type ListResponse struct {
	Items []ListItem `json:"items"`
}

type ListItem struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	CategoryId string    `json:"categoryId,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
	Price      float64   `json:"price"`
}

func TestExpenseList(t *testing.T) {

	app := utils.SetupAppInstance()

	groups := []*utils.TestGroup[ListFilters, ListResponse]{
		utils.NewTestGroup[ListFilters, ListResponse]("Invalid Filters", 400).
			Case("Empty Query", ListFilters{}).
			// Date
			Case("Empty date", ListFilters{Date: ""}).
			Case("Invalid date format #1", ListFilters{Date: "1234"}).
			Case("Invalid date format #2", ListFilters{Date: "02/01/2006"}).
			// Page
			Case("Negative page", ListFilters{Date: "02.01.2006", Page: -1}).
			// Limit
			Case("Negative limit", ListFilters{Date: "02.01.2006", Limit: -1}).
			Case("Too big limit", ListFilters{Date: "02.01.2006", Limit: 101}).
			// Category
			Case("Invalid category format #1", ListFilters{Date: "02.01.2006", CategoryId: "152"}).
			Case("Invalid category format #2", ListFilters{Date: "02.01.2006", CategoryId: "1a5a2"}),

		utils.NewTestGroup[ListFilters, ListResponse]("Valid Filters", 200).
			Case("Basic", ListFilters{Date: "02.01.2006"}).
			Case("With name", ListFilters{Date: "02.01.2006", Name: "expense"}).
			Case("With category", ListFilters{Date: "02.01.2006", CategoryId: uuid.NewString()}).
			Case("With page", ListFilters{Date: "02.01.2006", Page: 1}).
			Case("With limit", ListFilters{Date: "02.01.2006", Limit: 10}).
			Case("All", ListFilters{
				Date:       "02.01.2006",
				Name:       "expense",
				CategoryId: uuid.NewString(),
				Page:       1,
				Limit:      10,
			}),
	}

	for _, testGroup := range groups {
		t.Run(testGroup.Name, func(t *testing.T) {
			for _, testCase := range testGroup.Cases {
				t.Run(testCase.Name, func(t *testing.T) {

					query := utils.StructToQuery(testCase.Input)

					req := httptest.NewRequest(http.MethodGet, "/expense/?"+query, nil)

					resp, err := app.Test(req)
					if err != nil {
						t.Fatalf("app.Test failed: %s", err.Error())
					}

					assert.Equal(t, testGroup.Status, resp.StatusCode)

					if testGroup.Status == fiber.StatusOK {
						var got ListResponse

						err := json.NewDecoder(resp.Body).Decode(&got)
						assert.NoError(t, err)

						assert.NotEmpty(t, got.Items)
						for _, item := range got.Items {
							assert.NotEmpty(t, item.Id)
							assert.NotEmpty(t, item.Name)
							assert.Greater(t, item.Price, 0.0)
							assert.False(t, item.CreatedAt.IsZero())
						}
					}
				})
			}
		})
	}
}
