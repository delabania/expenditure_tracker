package main

import (
	"expenditure_tracker/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	env := &Env{db: &db.InMemoryDatabase{}}
	router.GET("/expenditures", env.GetExpenditures)
	router.POST("/expenditures/add", env.AddExpenditure)

	router.Run("localhost:8080")
}
