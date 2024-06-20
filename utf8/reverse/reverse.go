//go:build !solution

package reverse

func Reverse(input string) string {
	runes := []rune(input)
	i, j := 0, len(runes)-1
	tmp := rune(2)
	for i < j {
		tmp = runes[i]
		runes[i] = runes[j]
		runes[j] = tmp
		i++
		j--
	}
	return string(runes)
}
