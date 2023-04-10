package main

import (
	"expenditure_tracker/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Env struct {
	db db.Database
}

func (e *Env) GetExpenditures(c *gin.Context) {
	expenditures, err := e.db.GetAllExpenditures()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, expenditures)
}

func (e *Env) AddExpenditure(c *gin.Context) {
	var expenditure db.Expenditure
	if err := c.BindJSON(&expenditure); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := e.db.SaveExpenditure(expenditure)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	expenditure.ID = id
	c.IndentedJSON(http.StatusCreated, expenditure)
}
