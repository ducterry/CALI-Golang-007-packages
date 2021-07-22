package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*
	- Ngày: 21.07.2021
	- By: Nguyễn Đăng Đức
*/
func main() {
	// Khai báo import
	number1 := 12.5
	number2 := 10.5
	var maxNumber = math.Max(number1, number2)
	fmt.Printf("Số lớn nhất giữa %.2f bà %.2f là %.2f", number1, number2, maxNumber)

	// Calculate the square root of a number
	fmt.Println(math.Sqrt(225))

	// Printing the value of `𝜋`
	fmt.Println(math.Pi)

	// Epoch time in milliseconds
	epoch := time.Now().Unix()
	fmt.Println(epoch)

	// Generating a random integer between 0 to 100
	rand.Seed(epoch)
	fmt.Println(rand.Intn(100))

}
