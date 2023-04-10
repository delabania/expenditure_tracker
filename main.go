package expenditure_tracker

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/expenditures", GetExpenditures)

	router.Run("localhost:8080")
}
