package main

import (
	"fmt"

	"github.com/cerncorp/larago/routes"
)

func main() {
	fmt.Println("Larago Hello!")

	r, _ := routes.LoadRoutes()

	r.Run()
}
