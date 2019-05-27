package main

import (
	"bufio"
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
	seen := map[string]map[int]bool{}
	for sc.Scan() {
		for i := range sc.Text() {
			s := sc.Text()[:i] + sc.Text()[i+1:]
			if seen[s] == nil {
				seen[s] = map[int]bool{i: true}
				continue
			}
			if seen[s][i] {
				log.Printf("answer is %s", s)
				return
			}
			seen[s][i] = true
		}
	}
}
