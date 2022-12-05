package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scores := [][]int{
		{3, 4, 8},
		{1, 5, 9},
		{2, 6, 7},
	}
	total := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		play := strings.Split(scanner.Text(), " ")
		op := []byte(play[0])[0]
		me := []byte(play[1])[0]
		total += scores[op-65][me-88]
	}
	for scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	println(strconv.Itoa(total))
}
