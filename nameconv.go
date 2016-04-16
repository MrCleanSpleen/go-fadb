package fadb

import (
	"fmt"
	//c "github.com/skilstak/go-colors"
	i "github.com/skilstak/go-input"
	"strings"
)

func CapString(input string) string {
	words := strings.Fields(input)

	for index, word := range words {
		words[index] = strings.Title(word)
	}
	return strings.Join(words, " ")
}

func StripChars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

func NameConv() {
	name := i.Ask("Please put in your name: ")
	if strings.ContainsAny(name, "_") {
		name1 := StripChars(name, "~!@#$%^&*()<>.,?/|}{[]:;-'+=")
		name2 := strings.Replace(name1, "_", " ", -1)
		name3 := CapString(name2)
		name4 := StripChars(name3, " ")
		fmt.Println(name4)
	} else if strings.ContainsAny(name, "-") {
		name1 := StripChars(name, "~!@#$%^&*()<>.,?/|}{[]:;_'+=")
		name2 := strings.Replace(name1, "-", " ", -1)
		name3 := CapString(name2)
		name4 := StripChars(name3, " ")
		fmt.Println(name4)
	} else {
		fmt.Println(StripChars(name, "~!@#$%^&*()<>.,?/|}{[]:;_- +='"))
	}
}
