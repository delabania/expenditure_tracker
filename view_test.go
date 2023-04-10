package main

import (
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
