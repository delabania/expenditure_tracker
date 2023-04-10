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
