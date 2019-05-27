package main

import (
	"bufio"
	"fmt"
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
	sum := 0
	for sc.Scan() {
		i := 0
		_, err := fmt.Sscanf(sc.Text(), "%d", &i)
		if err != nil {
			log.Fatalf("format err: %v", err)
		}
		sum += i
	}
	log.Printf("answer is %d", sum)
}
