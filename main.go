package main

import (
	"fmt"

	Result "github.com/golang-oop/result/src"
)

func main() {
	r := Result.New()
	fmt.Println(r.HasError())
}
