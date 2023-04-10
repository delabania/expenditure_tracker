package expenditure_tracker

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddExpenditure() {}

func AddSettlement() {}

func GetExpenditures(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{})
}
