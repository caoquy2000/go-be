package main

import (
	"fmt"
	"go-be/internal/routers"
)

func main() {
	fmt.Println("Hello World!")
	r := routers.NewRouter()

	r.Run(":8000")
}