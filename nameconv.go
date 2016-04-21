package fadb

import (
	"fmt"
	i "github.com/skilstak/go-input"
	"log"
	"regexp"
	"strings"
)

func capstring(input string) string {
	words := strings.Fields(input)

	for index, word := range words {
		words[index] = strings.Title(word)
	}
	return strings.Join(words, " ")
}

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}
func remove(str, chr string) string {
	reg, err := regexp.Compile("[^A-Za-z0-9_-]+")
	if err != nil {
		log.Fatal(err)
	}
	safe := reg.ReplaceAllString(str, "")
	return safe
}
func NameConv(input string) {
	name := input
	if strings.ContainsAny(name, "_") {
		name1 := remove(name, " ")
		name2 := strings.Replace(name1, "_", " ", -1)
		name3 := capstring(name2)
		name4 := stripchars(name3, " ")
		fmt.Println(name4)
	} else if strings.ContainsAny(name, "-") {
		name1 := remove(name, "")
		name2 := strings.Replace(name1, "-", " ", -1)
		name3 := capstring(name2)
		name4 := stripchars(name3, " ")
		fmt.Println(name4)
	} else {
		fmt.Println(remove(name, ""))
	}
}
