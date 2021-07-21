package numbers

import "math"

/*
   - Ngày: 21.07.2021
   - By: Đức NĐ
*/
func isPrime(num int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(num)))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return num > 1
}
