package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	max := 0
	cur := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) == 0 {
			cur = 0
			continue
		}
		n, err := strconv.Atoi(t)
		if err != nil {
			log.Fatal(err)
		}
		cur += n
		if cur > max {
			max = cur
		}
	}
	for scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	println(strconv.Itoa(max))
}
