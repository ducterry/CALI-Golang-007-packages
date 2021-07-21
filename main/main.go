package main

import (
	"fmt"
	"math"
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

}
