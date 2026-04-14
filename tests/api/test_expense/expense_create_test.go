package test_expense

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tmazitov/ayda-order-service.git/tests/api/utils"
)

type CreateRequest struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Date       string  `json:"date"`
	CategoryId string  `json:"categoryId"`
}

type CreateResponse struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryId string  `json:"categoryId"`
	CreatedAt  string  `json:"createdAt"`
}

func TestExpenseCreate(t *testing.T) {

	app := utils.SetupAppInstance()

	groups := []*utils.TestGroup[CreateRequest, CreateResponse]{

		utils.NewTestGroup[CreateRequest, CreateResponse]("Invalid Create", 400).
			Output(CreateResponse{}).
			Case("Empty body", CreateRequest{}).
			// Invalid Name
			Case("No name", CreateRequest{Price: 1, Date: "02.01.2006"}).
			Case("Empty name", CreateRequest{Price: 1, Date: "02.01.2006", Name: ""}).
			Case("Too large name", CreateRequest{Price: 1, Date: "02.01.2006", Name: utils.TooLargeString(256)}).
			// Invalid Date
			Case("No date", CreateRequest{Name: "expense", Price: 1}).
			Case("Empty date", CreateRequest{Name: "expense", Price: 1, Date: ""}).
			Case("Wrong format date #1", CreateRequest{Name: "expense", Price: 1, Date: "02/01/2006"}).
			Case("Wrong format date #2", CreateRequest{Name: "expense", Price: 1, Date: "02.01.2006"}).
			Case("Too large date", CreateRequest{Name: "expense", Price: 1, Date: utils.TooLargeString(256)}).
			// Invalid Price
			Case("No price", CreateRequest{Name: "expense", Date: "02.01.2006"}).
			Case("Zero price", CreateRequest{Name: "expense", Date: "02.01.2006", Price: 0}).
			Case("Negative price", CreateRequest{Name: "expense", Date: "02.01.2006", Price: -10}).
			// Invalid Category
			Case("Zero price", CreateRequest{Name: "expense", Date: utils.ISODate(2006, 01, 02), Price: 1, CategoryId: "132"}).
			Case("Zero price", CreateRequest{Name: "expense", Date: utils.ISODate(2006, 01, 02), Price: 1, CategoryId: utils.TooLargeString(256)}),

		utils.NewTestGroup[CreateRequest, CreateResponse]("Valid Create", 201).
			Output(CreateResponse{
				Name:  "expense",
				Price: 1,
			}).
			Case("Basic", CreateRequest{Name: "expense", Price: 1, Date: utils.ISODate(2006, 01, 02)}).
			Case("Basic", CreateRequest{Name: "expense", Price: 1, Date: utils.ISODate(2006, 01, 02), CategoryId: uuid.NewString()}),
	}

	for _, testGroup := range groups {
		t.Run(testGroup.Name, func(t *testing.T) {
			for _, testCase := range testGroup.Cases {
				t.Run(testCase.Name, func(t *testing.T) {

					data, _ := json.Marshal(testCase.Input)

					req := httptest.NewRequest(http.MethodPost, "/expense/", bytes.NewReader(data))
					req.Header.Set("Content-Type", "application/json")

					resp, err := app.Test(req)
					if err != nil {
						t.Fatalf("app.Test failed: %s", err.Error())
					}

					assert.Equal(t, testGroup.Status, resp.StatusCode)

					fmt.Println("status", testGroup.Status)

					if resp.StatusCode == http.StatusCreated {
						var got CreateResponse
						json.NewDecoder(resp.Body).Decode(&got)

						assert.Equal(t, testGroup.Expected.Name, got.Name)
						assert.Equal(t, testGroup.Expected.Price, got.Price)

						assert.NotEmpty(t, got.Id)
						assert.NotEmpty(t, got.CreatedAt)
					}
				})
			}
		})
	}
}
