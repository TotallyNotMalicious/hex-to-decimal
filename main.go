package main

import (
	"fmt"
	"strconv"
)

func main() {
	var hex string

	fmt.Print("Hex: ")
	fmt.Scan(&hex)

	num, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		fmt.Print("Invalid input")
		return
	}
	fmt.Print("Decimal: ", num)
}
