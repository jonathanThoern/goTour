package main

import (
	"golang.org/x/tour/wc"
	"strings"
	//"fmt"
)

func WordCount(s string) map[string]int {
	
	words := strings.Fields(s) 	// Divide string into separate words
	
	m := make(map[string]int) 	// Initialize map
	
	for i := range words{		// Range over words
		_, ok := m[words[i]]	// Check if word exists as key
		if !ok {
			m[words[i]] = 1		// If not add the word to map
		} else{
			m[words[i]] += 1	// Else increment counter of word
		}
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}
