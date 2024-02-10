package api

import (
	"fmt"
	"github.com/graphql-go/handler"
	"net/http"
)

const RootEndpoint = "/api/v1"

func main() {
	http.Handle(RootEndpoint, handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	}))

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
