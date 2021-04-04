package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// position=<%2d,  1> velocity=< 0,  2>
	f, err := os.Open("input.htm")
	if err != nil {
		log.Fatalf("can't open input: %v", err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	r, err := regexp.Compile(`^position=<\s*(\-?\d+),\s*(\-?\d+)> velocity=<\s*(\-?\d+),\s*(\-?\d+)>$`)
	if err != nil {
		log.Fatal(err)
	}
	pp := []*point{}
	for sc.Scan() {
		strs := r.FindAllStringSubmatch(sc.Text(), -1)
		x, err := strconv.Atoi(strs[0][1])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(strs[0][2])
		if err != nil {
			log.Fatal(err)
		}
		vX, err := strconv.Atoi(strs[0][3])
		if err != nil {
			log.Fatal(err)
		}
		vY, err := strconv.Atoi(strs[0][4])
		if err != nil {
			log.Fatal(err)
		}
		pp = append(pp, &point{x, y, vX, vY})
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	item := points(pp)
	i := 0
	for {
		if i > 10500 {
			break
		}
		if math.Abs(10407-float64(i)) < 3 {
			log.Println(i)
			item.render()
		}
		item.move()
		i++
	}
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

type points []*point

func (self points) move() {
	for _, p := range self {
		p.move()
	}
}

func (self points) render() {
	var left, bottom = MaxInt, MaxInt
	var right, top = MinInt, MinInt
	for _, p := range self {
		if p.x < left {
			left = p.x
		}
		if p.y < bottom {
			bottom = p.y
		}
		if p.x > right {
			right = p.x
		}
		if p.y > top {
			top = p.y
		}
	}
	dict := map[string]bool{}
	for _, p := range self {
		cor := fmt.Sprintf("%d_%d", p.x-left, p.y-bottom)
		dict[cor] = true
	}
	dx := right - left + 1
	dy := top - bottom + 1
	// log.Println(dx, dy)
	// return
	for j := 0; j < dy; j++ {
		for i := 0; i < dx; i++ {
			cor := fmt.Sprintf("%d_%d", i, j)
			if dict[cor] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}

}

type point struct {
	x, y   int
	vX, vY int
}

func (self *point) move() {
	self.x += self.vX
	self.y += self.vY
}
