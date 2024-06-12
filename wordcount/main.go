//go:build !solution

package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	argsWithoutProg := os.Args[1:]
	resMap := make(map[string]int)

	for _, arg := range argsWithoutProg {
		f, err := os.Open(arg)
		check(err)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			resMap[line]++
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Ошибка чтения файла:", err)
		}
		err = f.Close()
		if err != nil {
			return
		}
	}
	for key, value := range resMap {
		if value >= 2 {
			fmt.Printf("%d\t%s\n", value, key)
		}
	}
}
