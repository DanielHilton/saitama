package main

import (
	"flag"
	"fmt"
)

func main() {
	endpoint := flag.String("endpoint", "http://localhost", "The endpoint you want to attack")
	standQuantity := flag.Int("quantity", 1, "The amount of stands you want to attack your endpoint with")

	flag.Parse()

	notice := fmt.Sprintf("Attacking %s endpoint with %d stands", *endpoint, *standQuantity)
	fmt.Println(notice)
}
