package main
import (
	"fmt"
	"os"
	"flag"
)
func main() {
	filePath := flag.String("file", "", "Path to the file")
	topWords := flag.Int("top", 5, "Number of top words to display")

	flag.Parse()

	if *filePath == "" {
		fmt.Fprintln(os.Stderr, "Failure: You did not provide a file path!")
		flag.Usage()
		os.Exit(1)
	}

	if *topWords <= 0 {
		fmt.Fprintln(os.Stderr, "Failure: The number of top words to display must be a positive integer!")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("File path: %s\n", *filePath)
	fmt.Printf("Top words to display: %d\n", *topWords)

	file, err := os.Open(*filePath)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failure: Could not open the file!")
		os.Exit(1)
	}

	defer file.Close()

}