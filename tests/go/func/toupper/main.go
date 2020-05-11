package main

import (
	"strings"

	"./student"
)

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	table := lib.MultRandASCII()

	table = append(table,
		"Hello! How are you?",
	)

	for _, arg := range table {
		lib.Challenge("ToUpper", student.ToUpper, toUpper, arg)
	}
}