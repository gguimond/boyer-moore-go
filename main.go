package main

import (
	"fmt"
	"strings"
)

// buildBadCharTable builds a table that stores
// the rightmost occurrence of each character in the pattern
func buildBadCharTable(pattern string) map[byte]int {
	badCharTable := make(map[byte]int)
	
	// Initialize all possible characters to -1 (not found)
	for c := 0; c < 256; c++ {
		badCharTable[byte(c)] = -1
	}
	
	// Fill the actual positions of characters in pattern
	for i := 0; i < len(pattern); i++ {
		badCharTable[pattern[i]] = i
	}
	
	return badCharTable
}

// BoyerMooreSearch finds all occurrences of pattern in text
// Returns a slice of starting indices where pattern is found
func BoyerMooreSearch(text, pattern string) []int {
	if len(pattern) == 0 {
		return nil
	}
	
	// Build the bad character table
	badCharTable := buildBadCharTable(pattern)
	
	// Store the found occurrences
	occurrences := []int{}
	
	// Start searching
	i := 0
	for i <= len(text)-len(pattern) {
		j := len(pattern) - 1
		
		// Compare pattern with text from right to left
		for j >= 0 && pattern[j] == text[i+j] {
			j--
		}
		
		// If we completed all comparisons, we found a match
		if j < 0 {
			occurrences = append(occurrences, i)
			i++
		} else {
			// Calculate shift using the bad character rule
			charInText := text[i+j]
			lastOccurrence := badCharTable[charInText]
			
			// Calculate how much to shift
			// We want to align the bad character in text with its
			// rightmost occurrence in pattern
			if lastOccurrence < j {
				// If character exists in pattern, align it
				// If it doesn't exist (lastOccurrence == -1), shift by j+1
				i += j - lastOccurrence
			} else {
				// If the character exists in pattern but only at or to the right of the mismatch,
				// shift by 1
				i++
			}
		}
	}
	
	return occurrences
}


func main() {
	fmt.Println("Hello")
	//pattern := "abcefg"
	// Build the bad character table
	//badCharTable := buildBadCharTable(pattern)
	//fmt.Println(badCharTable[98])
	//fmt.Printf("%v", badCharTable)
	text := "ABAAABCDBBABCDDEBCABC"
	pattern := "ABC"
	
	fmt.Printf("Searching for '%s' in '%s'\n", pattern, text)
	
	occurrences := BoyerMooreSearch(text, pattern)
	
	if len(occurrences) == 0 {
		fmt.Println("Pattern not found")
	} else {
		fmt.Printf("Pattern found at positions: ")
		for i, pos := range occurrences {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(pos)
		}
		fmt.Println()
		
		// Visualize the matches
		fmt.Println("\nVisualization:")
		for _, pos := range occurrences {
			fmt.Println(text)
			fmt.Printf("%s%s\n", strings.Repeat(" ", pos), pattern)
		}
	}
}