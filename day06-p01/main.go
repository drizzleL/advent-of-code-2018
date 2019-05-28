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
	board, maxI, maxJ, points := getBoard()
	var distance = 0
	for len(points) != 0 {
		newPoints := []*point{}
		readyPoints := map[string]*point{}
		for _, p := range points {
			if board[p.j][p.i] == nil {
				board[p.j][p.i] = &boardPoint{p.id, distance}
				readyPoints[fmt.Sprintf("%d_%d", p.i, p.j)] = p // ready to add
				continue
			}
			if board[p.j][p.i].id == p.id || board[p.j][p.i].id == -1 {
				continue
			}
			if board[p.j][p.i].distance == distance { // just added before
				delete(readyPoints, fmt.Sprintf("%d_%d", p.i, p.j))
				board[p.j][p.i].id = -1
			}
		}
		for _, p := range readyPoints {
			if p.i-1 >= 0 && board[p.j][p.i-1] == nil {
				newPoints = append(newPoints, &point{p.id, p.i - 1, p.j})
			}
			if p.i+1 <= maxI && board[p.j][p.i+1] == nil {
				newPoints = append(newPoints, &point{p.id, p.i + 1, p.j})
			}
			if p.j+1 <= maxJ && board[p.j+1][p.i] == nil {
				newPoints = append(newPoints, &point{p.id, p.i, p.j + 1})
			}
			if p.j-1 >= 0 && board[p.j-1][p.i] == nil {
				newPoints = append(newPoints, &point{p.id, p.i, p.j - 1})
			}
		}
		points = newPoints
		distance++
	}
	not := map[int]bool{}
	for i := 0; i <= maxI; i++ {
		if board[0][i] != nil {
			not[board[0][i].id] = true
		}
		if board[maxJ][i] != nil {
			not[board[maxJ][i].id] = true
		}
	}
	for j := 0; j <= maxJ; j++ {
		if board[j][0] != nil {
			not[board[j][0].id] = true
		}
		if board[j][maxI] != nil {
			not[board[j][maxI].id] = true
		}
	}
	var max struct {
		id  int
		val int
	}
	dict := map[int]int{}
	for _, line := range board {
		for _, v := range line {
			if not[v.id] {
				continue
			}
			dict[v.id]++
			if dict[v.id] > max.val {
				max.id = v.id
				max.val = dict[v.id]
			}
		}
	}
	log.Printf("answer is %d", max.val)
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
