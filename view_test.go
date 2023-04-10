package main

import (
	"bytes"
	"encoding/json"
	"expenditure_tracker/db"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"time"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEnv_GetExpenditures(t *testing.T) {
	type fields struct {
		db db.Database
	}

	tests := []struct {
		name   string
		fields fields
		want   []db.Expenditure
	}{
		{
			name: "test get expenditures",
			fields: fields{db: &db.InMemoryDatabase{
				Persons: []db.Person{},
				Expenditures: []db.Expenditure{
					{ID: 1, Description: "Shopping", Amount: 100, SplitRatio: 0.5, Date: time.Date(2022, 9, 1, 16, 0, 0, 0, time.UTC)},
					{ID: 2, Description: "Gas", Amount: 50, SplitRatio: 0.25, Date: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)},
				},
				Settlements: []db.Settlement{},
			}},
			want: []db.Expenditure{
				{ID: 1, Description: "Shopping", Amount: 100, SplitRatio: 0.5, Date: time.Date(2022, 9, 1, 16, 0, 0, 0, time.UTC)},
				{ID: 2, Description: "Gas", Amount: 50, SplitRatio: 0.25, Date: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Env{
				db: tt.fields.db,
			}
			writer := httptest.NewRecorder()

			request, err := http.NewRequest("GET", "/expenditures", nil)
			if err != nil {
				t.Fatal(err)
			}
			router := gin.Default()
			router.GET("/expenditures", e.GetExpenditures)
			router.ServeHTTP(writer, request)

			assert.Equal(t, http.StatusOK, writer.Code)
			var expenditures []db.Expenditure
			_ = json.Unmarshal(writer.Body.Bytes(), &expenditures)

			assert.Equal(t, tt.want, expenditures)
		})
	}
}

func TestEnv_AddExpenditure(t *testing.T) {
	type fields struct {
		db          db.Database
		expenditure db.Expenditure
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"test add expenditure",
			fields{
				db: &db.InMemoryDatabase{
					Persons:      []db.Person{},
					Expenditures: []db.Expenditure{},
					Settlements:  []db.Settlement{},
				}, expenditure: db.Expenditure{
					Description: "Shopping",
					Amount:      100,
					SplitRatio:  0.5,
					Date:        time.Date(2022, 9, 1, 16, 0, 0, 0, time.UTC),
				},
			}, args{c: &gin.Context{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Env{
				db: tt.fields.db,
			}
			writer := httptest.NewRecorder()

			expenditure, _ := json.Marshal(tt.fields.expenditure)
			request, err := http.NewRequest(
				"POST", "/expenditures/add", bytes.NewBuffer(expenditure))
			if err != nil {
				t.Fatal(err)
			}
			router := gin.Default()
			router.POST("/expenditures/add", e.AddExpenditure)
			router.ServeHTTP(writer, request)

			assert.Equal(t, http.StatusCreated, writer.Code)
			var id int64
			_ = json.Unmarshal(writer.Body.Bytes(), &id)
			assert.Equal(t, int64(0), id)

			expenditureReturned, _ := e.db.GetExpenditure(id)
			assert.Equal(t, tt.fields.expenditure, expenditureReturned)
		})
	}
}
