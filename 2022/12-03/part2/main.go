package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func IsRuneInString(r rune, s string) bool {
	for _, x := range s {
		if x == r {
			return true
		}
	}
	return false
}

func GetCommonAmong3(s1, s2, s3 string) (c rune) {
	for _, r := range s1 {
		if IsRuneInString(r, s2) {
			if IsRuneInString(r, s3) {
				c = r
				break
			}
		}
	}
	return c
}

func GetScore(r rune) (s int) {
	// is lowercase
	if r >= 97 {
		s = int(r) - 96
	} else { // is uppercase
		s = int(r) - 38
	}
	return s
}

func main() {
	input, err := os.Open("../input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	sum := 0
	var e1, e2, e3 string
	for scanner.Scan() {
		// e := scanner.Text()

		r := GetCommonAmong3(e1, e2, e3)
		s := GetScore(r)
		fmt.Printf("%c, %d\n", r, s)
		sum += s
	}
	for scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	println(strconv.Itoa(sum))
}
