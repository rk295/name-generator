package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	coloursFile = "colours.txt"
)

type permutations struct {
	colours []string
}

func main() {

	perms, err := readData()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, c := range perms.colours {
		fmt.Println(c)
	}

	entry := ran(len(perms.colours))
	fmt.Println("Random selection:", perms.colours[entry])

}

func ran(length int) int {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	return r.Intn(length)
}

func readData() (permutations, error) {
	var perms permutations

	// Read colours
	lines, err := readLines(coloursFile)
	if err != nil {
		return perms, err
	}
	perms.colours = lines

	return perms, nil
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
