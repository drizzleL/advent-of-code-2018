package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input")
	defer f.Close()
	if err != nil {
		log.Fatalf("can't open input: %v", err)
	}
	sc := bufio.NewScanner(f)
	var twos, threes int
	for sc.Scan() {
		cntMap := map[rune]int{}
		for _, c := range sc.Text() {
			cntMap[c]++
		}
		var gotTwo, gotThree bool
		for _, v := range cntMap {
			switch v {
			case 2:
				gotTwo = true
			case 3:
				gotThree = true
			}
		}
		if gotTwo {
			twos++
		}
		if gotThree {
			threes++
		}
	}
	log.Printf("answer is %d", twos*threes)
}
