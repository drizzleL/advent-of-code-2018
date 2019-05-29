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
	_, _, root := parse(nums)

	log.Printf("answer is %d", root.val())
}

// parse first node
func parse(nums []int) (forward int, sum int, n *node) {
	n = &node{}
	childNum, metaNum := nums[0], nums[1]
	forward = 2
	for childNum != 0 { // deal with all children
		f, newsum, child := parse(nums[forward:])
		n.children = append(n.children, child)
		forward += f
		sum += newsum
		childNum--
	}
	for i := 0; i < metaNum; i++ {
		n.valIDs = append(n.valIDs, nums[forward+i])
		sum += nums[forward+i]
	}
	forward += metaNum
	return forward, sum, n
}

type node struct {
	children []*node
	valIDs   []int
}

func (self *node) val() (sum int) {
	if len(self.children) == 0 {
		for _, valID := range self.valIDs {
			sum += valID
		}
		return
	}
	valDict := map[int]int{}
	for i, child := range self.children {
		valDict[i] = child.val()
	}
	for _, valID := range self.valIDs {
		sum += valDict[valID-1]
	}
	return
}
