package main

import "fmt"

func getDiscount() (int, error) {
	return 3, nil
}

func recordDiscount(discount int) {
	fmt.Printf("debug:x the discount is: %d\n", discount)
}

func main() {
	var discount int
	carTotal := 500

	if carTotal > 100 {
		discount, err := getDiscount()
		if err != nil {
			panic(err)
		}

		recordDiscount(discount)
	}

	fmt.Printf("You saved: %d\n", discount)
}
