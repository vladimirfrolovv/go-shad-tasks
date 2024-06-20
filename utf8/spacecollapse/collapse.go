//go:build !solution

package spacecollapse

func CollapseSpaces(input string) string {
	runes := []rune(input)
	l, r, countSpace := 0, 0, 0
	for r < len(runes) {
		if runes[r] == ' ' || runes[r] == '\t' || runes[r] == '\n' || runes[r] == '\r' {
			countSpace++
			if countSpace == 1 {
				runes[l] = ' '
				l++
			}
		} else {
			runes[l] = runes[r]
			l++
			countSpace = 0
		}
		r++
	}
	return string(runes[:l])
}
