package main
import (
	"slices"
	"cmp"
	"fmt"
)
type wordFreq struct {
	word string
	freq int
}
func PrintTopWords(counts map[string]int, top int) {
	list := make([]wordFreq, 0, len(counts))

	for w, c := range(counts) {
		list = append(list, wordFreq{word : w, freq : c})
	}	

	slices.SortFunc(list, func (a, b wordFreq) int {
		if a.freq != b.freq {
			return cmp.Compare(b.freq, a.freq)
		}
		return cmp.Compare(a.word, b.word)
	})

	if top > len(list) {
		top = len(list)
	}

	for i := 0; i < top; i++ {
		fmt.Printf("%-20s %d\n", list[i].word, list[i].freq)
	}
}