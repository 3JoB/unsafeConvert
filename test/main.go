package main

import (
	"fmt"

	"github.com/3JoB/unsafeConvert"
)

var f = "12345 float32"

func main() {
	fmt.Println(unsafeConvert.Atoi(f))
}
