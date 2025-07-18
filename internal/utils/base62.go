package utils

import "strings"

const dict string = "JhzfMpXQS3OWvTIoLHBxNs9eF7yVCrYtl2jb60dkAwmK1Z8Piuc5EgDnUa4RGq"

var intToChar []rune
var charToInt map[rune]int64

func init() {
	dict_rune := []rune(dict)
	charToInt = make(map[rune]int64, 62)
	intToChar = make([]rune, 62)
	for i := int64(0); i < 62; i++ {
		charToInt[dict_rune[i]] = i
		intToChar[i] = dict_rune[i]
	}
}

func IntToBase62(num int64) string {

	var builder strings.Builder

	for ; num > 0; num /= 62 {
		t := num % 62
		builder.WriteRune(intToChar[t])
	}

	return builder.String()
}

func Base62ToInt(str string) int64 {
	var t int64 = 1
	var res int64 = 0
	str_rune := []rune(str)
	for i := 0; i < len(str); i++ {
		res += charToInt[str_rune[i]] * t
		t *= 62
	}

	return res
}
