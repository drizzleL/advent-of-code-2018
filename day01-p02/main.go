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
	sc := bufio.NewScanner(f)
	seen := map[int]bool{}
	seen[0] = true
	nums := []int{}
	for sc.Scan() {
		i := 0
		_, err := fmt.Sscanf(sc.Text(), "%d", &i)
		if err != nil {
			log.Fatalf("can't format: %v", err)
		}
		nums = append(nums, i)
	}
	sum := 0
	for {
		for _, i := range nums {
			sum += i
			if seen[sum] {
				log.Printf("answer is %d", sum)
				return
			}
			seen[sum] = true
		}
	}
}
