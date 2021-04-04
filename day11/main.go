package main

import (
	"fmt"
	"log"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

type maxVal struct {
	val  int
	x    int
	y    int
	size int
}

func main() {
	board := make([][]int, 301)
	for i := range board {
		board[i] = make([]int, 301)
	}
	serial := 9435
	for i := 1; i < 301; i++ {
		for j := 1; j < 301; j++ {
			board[i][j] = powerLevel(i, j, serial)
		}
	}
	max := maxVal{val: MinInt}
	ret := map[string]int{} // key: size:x_y
	for size := 1; size <= 300; size++ {
		for i := 1; i <= 300-size+1; i++ {
			for j := 1; j <= 300-size+1; j++ {
				val := grab(size, board, i, j, ret)
				ret[key(size, i, j)] = val
				if val > max.val {
					max.val = val
					max.x = i
					max.y = j
					max.size = size
				}
			}
		}
		log.Println(size)

	}
	log.Println(max)
}

func powerLevel(x, y, serial int) int {
	return ((((x+10)*y + serial) * (x + 10) / 100) % 10) - 5
}

func grab(size int, board [][]int, x, y int, dict map[string]int) int {
	sum := board[x][y]
	if size == 1 {
		return sum
	}
	k := key(size-1, x+1, y+1)
	sum += dict[k]
	for i := 1; i < size; i++ {
		sum += board[x][y+i] + board[x+i][y]
	}
	return sum
}

func key(size, x, y int) string {
	return fmt.Sprintf("%d:%d_%d", size, x, y)
}
