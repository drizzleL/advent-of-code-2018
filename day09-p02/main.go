package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalf("can't open file: %v", err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("read file: %v", err)
	}
	var playerNum, turns int
	_, err = fmt.Sscanf(string(b), "%d players; last marble is worth %d points", &playerNum, &turns)
	if err != nil {
		log.Fatalf("format err: %v", err)
	}
	turns *= 100
	curr := &marble{val: 0}
	curr.prev = curr
	curr.next = curr
	scores := make([]int, playerNum)
	for i := 0; i < turns; i++ {
		val := i + 1
		if val%23 == 0 {
			player := i % playerNum
			rem := curr.prev.prev.prev.prev.prev.prev.prev
			scores[player] += val + rem.val
			rem.prev.next = rem.next
			rem.next.prev = rem.prev
			curr = rem.next
			continue
		}
		newMarble := &marble{val: val}
		newMarble.next = curr.next.next
		newMarble.prev = curr.next
		newMarble.prev.next = newMarble
		newMarble.next.prev = newMarble
		curr = newMarble
	}
	max := 0
	for _, score := range scores {
		if score > max {
			max = score
		}
	}
	log.Printf("answer is %d", max)
}

type marble struct {
	val  int
	prev *marble
	next *marble
}
