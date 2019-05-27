package main

import (
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
		log.Fatalf("read failed: %v", err)
	}
	minLen := len(b)
	for base := 'a'; base <= 'z'; base++ {
		var result []byte
		for i := 0; i < len(b); i++ {
			if b[i] == byte(base) || helper(byte(base), b[i]) {
				continue
			}
			if len(result) == 0 || !helper(result[len(result)-1], b[i]) {
				result = append(result, b[i])
			} else {
				result = result[:len(result)-1]
			}
		}
		if len(result) < minLen {
			minLen = len(result)
		}
	}
	log.Printf("answer is %d", minLen)
}

func helper(a, b byte) bool {
	tmp := int(a) - int(b)
	return tmp == 32 || tmp == -32
}
