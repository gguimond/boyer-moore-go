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

// buildGoodSuffixTable builds tables for the good suffix rule
func buildGoodSuffixTable(pattern string) ([]int, []int) {
	m := len(pattern)
	
	// Border positions
	border := make([]int, m+1)
	// Position of the widest border
	shift := make([]int, m+1)
	
	// Preprocessing for case 2
	i := m
	j := m + 1
	border[i] = j
	
	for i > 0 {
		for j <= m && pattern[i-1] != pattern[j-1] {
			if shift[j] == 0 {
				shift[j] = j - i
			}
			j = border[j]
		}
		i--
		j--
		border[i] = j
	}
	
	// Preprocessing for case 1
	j = border[0]
	for i := 0; i <= m; i++ {
		if shift[i] == 0 {
			shift[i] = j
		}
		if i == j {
			j = border[j]
		}
	}
	
	return border, shift
}

// BoyerMooreSearch finds all occurrences of pattern in text
// Returns a slice of starting indices where pattern is found
func BoyerMooreSearch(text, pattern string) []int {
	if len(pattern) == 0 {
		return nil
	}
	
	// Build the bad character table
	badCharTable := buildBadCharTable(pattern)
	_, goodSuffixShift := buildGoodSuffixTable(pattern)
	
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
			i += goodSuffixShift[0] // Shift by good suffix rule
		} else {
			// Calculate shift using the bad character rule
			charInText := text[i+j]
			lastOccurrence := badCharTable[charInText]
			bcShift := j - lastOccurrence

			if bcShift < 1 {
				bcShift = 1
			}
			
			// Good suffix rule
			gsShift := goodSuffixShift[j+1]
			
			// Take the maximum of both shifts
			shift := bcShift
			if gsShift > bcShift {
				shift = gsShift
			}
			
			i += shift
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