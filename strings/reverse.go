package strings

/*
	- Ngày: 21.07.2021
	- By: ĐứcNĐ
*/
func Reverse(s string) string {
	runes := []rune(s)
	reversedRunes := reverseRunes(runes)
	return string(reversedRunes)
}
