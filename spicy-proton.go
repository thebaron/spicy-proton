package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

//go:embed nouns.txt
var nounList string

//go:embed adjectives.txt
var adjectiveList string

func readLines(content string) []string {
	return strings.Split(content, "\n")
}

func generateHostname(adjectives, nouns []string) string {
	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	return strings.TrimSpace(adjective) + "-" + strings.TrimSpace(noun)
}

func main() {

	var count int = 10

	if len(os.Args) >= 2 {
		var err error
		count, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("usage: spicy-proton [number of hostnames]\ninvalid number of hostnames:", err)
			return
		}
	}

	nouns := readLines(nounList)
	adjectives := readLines(adjectiveList)

	for i := 0; i < count; i++ {
		fmt.Println(generateHostname(adjectives, nouns))
	}
}
