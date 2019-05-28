package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getBoard() ([][]*boardPoint, int, int, []*point) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalf("can't open file: %v", err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	original := []*point{}
	var id int
	var maxI, maxJ int
	for sc.Scan() {
		var i, j int
		_, err := fmt.Sscanf(sc.Text(), "%d, %d", &i, &j)
		if err != nil {
			log.Fatalf("format err: %v", err)
		}
		original = append(original, &point{id, i, j})
		if i > maxI {
			maxI = i
		}
		if j > maxJ {
			maxJ = j
		}
		id++
	}
	points := original
	board := make([][]*boardPoint, maxJ+1)
	for i := range board {
		board[i] = make([]*boardPoint, maxI+1)
	}
	return board, maxI, maxJ, points
}
func main() {
	var max int
	_, maxI, maxJ, points := getBoard()
	for i := 0; i <= maxI; i++ {
		for j := 0; j <= maxJ; j++ {
			tmp := 0
			for _, p := range points {
				tmp += distance(p.i, p.j, i, j)
			}
			if tmp < 10000 {
				max++
			}
		}
	}
	log.Printf("answer is %d", max)
}

type boardPoint struct {
	id       int
	distance int
}

type point struct {
	id int
	i  int
	j  int
}

func distance(i1, j1, i2, j2 int) int {
	return abs(i1-i2) + abs(j1-j2)
}
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
