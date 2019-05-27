package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalf("can't open input: %v", err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	seen := map[string]int{}
	overlapped := 0
	for sc.Scan() {
		var id, left, top, wide, height int
		_, err := fmt.Sscanf(sc.Text(), "#%d @ %d,%d: %dx%d", &id, &left, &top, &wide, &height)
		if err != nil {
			log.Fatalf("can't format: %v", err)
		}
		for i := left; i < left+wide; i++ {
			for j := top; j < top+height; j++ {
				loc := fmt.Sprintf("%d,%d", i, j)
				seen[loc]++
				if seen[loc] == 2 {
					overlapped++
				}
			}
		}
	}
	log.Printf("answer is %d", overlapped)
}
