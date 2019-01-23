package main

import (
	"fmt"

	"blog/model"
)

func main() {
	fmt.Println(model.GeneratePasswordHash("123456"))
}
