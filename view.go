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
	// read all expenditures from db
	// return them as json

	e.db.GetAllExpenditures()
	c.IndentedJSON(http.StatusOK, gin.H{})
}
