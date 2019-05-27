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
	ids := map[int]bool{}
	seen := map[string]int{}
	for sc.Scan() {
		var id, left, top, wide, height int
		_, err := fmt.Sscanf(sc.Text(), "#%d @ %d,%d: %dx%d", &id, &left, &top, &wide, &height)
		if err != nil {
			log.Fatalf("can't format: %v", err)
		}
		ids[id] = false
		for i := left; i < left+wide; i++ {
			for j := top; j < top+height; j++ {
				loc := fmt.Sprintf("%d,%d", i, j)
				if oldid, ok := seen[loc]; ok { // existed before
					ids[oldid] = true
					ids[id] = true
				} else { // not existed before
					seen[loc] = id
				}
			}
		}
	}
	for id, val := range ids {
		if !val {
			log.Printf("answer is %d", id)
			return
		}
	}
	log.Fatal("no answer")

}
