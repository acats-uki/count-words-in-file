package main

import (
	"strings"
	"bufio"
	"os"
)

func CountWords(file *os.File) (map[string]int, error) {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		cleanWord := strings.Trim(word, ".,!?\"'():;")
		if (cleanWord != "") {
			counts[cleanWord]++;
		}
	}

	return counts, scanner.Err()
}