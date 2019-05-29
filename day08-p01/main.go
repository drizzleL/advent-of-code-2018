package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalf("can't open file: %v", err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("read err: %v", err)
	}

	bs := bytes.Split(b, []byte(` `))
	nums := []int{}
	for _, b := range bs {
		i, err := strconv.Atoi(string(b))
		if err != nil {
			log.Fatalf("int trans failed: %v", err)
		}
		nums = append(nums, i)
	}
	_, sum := parse(nums)
	log.Printf("answer is %d", sum)
}

// parse first node
func parse(nums []int) (forward int, sum int) {
	childNum, metaNum := nums[0], nums[1]
	forward = 2
	for childNum != 0 { // deal with all children
		f, newsum := parse(nums[forward:])
		forward += f
		sum += newsum
		childNum--
	}
	for i := 0; i < metaNum; i++ {
		sum += nums[forward+i]
	}
	forward += metaNum
	return forward, sum
}
