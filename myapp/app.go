package main

import (
	"07-packages/numbers"
	"07-packages/strings"
	"07-packages/strings/greeting"
	"fmt"
	str "strings"
)

func main() {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greeting.WelcomeText)

	fmt.Println(strings.Reverse("callicoder"))

	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))
}
