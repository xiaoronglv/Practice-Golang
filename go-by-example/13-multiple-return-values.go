package main

import (
	"fmt"
)

func geo() (int, int) {
	return 13123, 13421412
}

func main() {
	longitude, latitude := geo()

	fmt.Println("longitude is ", longitude)
	fmt.Println("latitude is ", latitude)

	_, l := geo()
	fmt.Println("l is ", l)
}
