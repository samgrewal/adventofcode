package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func InsertTop3(e int, t *[]int) {
	for i := len(*t) - 1; i >= 0; i-- {
		if e > (*t)[i] {
			for j := 0; j < i; j++ {
				(*t)[j] = (*t)[j+1]
			}
			(*t)[i] = e
			break
		}
	}
}

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	top3 := []int{0, 0, 0}
	cur := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) == 0 {
			InsertTop3(cur, &top3)
			cur = 0
			continue
		}
		n, err := strconv.Atoi(t)
		if err != nil {
			log.Fatal(err)
		}
		cur += n
	}
	InsertTop3(cur, &top3)

	total := 0
	for _, v := range top3 {
		total += v
	}
	println(strconv.Itoa(total))
}
