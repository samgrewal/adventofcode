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

func GetCommonFromHalves(s string) (c rune) {
	s1 := s[:len(s)/2]
	s2 := s[len(s)/2:]

	for _, r1 := range s1 {
		found := IsRuneInString(r1, s2)
		if found {
			c = r1
			break
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
	input, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	sum := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		r := GetCommonFromHalves(scanner.Text())
		s := GetScore(r)
		fmt.Printf("%c, %d\n", r, s)
		sum += s
	}
	for scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	println(strconv.Itoa(sum))
}
