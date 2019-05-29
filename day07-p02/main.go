package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var dict = map[rune]*point{}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalf("can't open file: %v", err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	ready := map[rune]bool{}
	not := map[rune]bool{}
	for sc.Scan() {
		var dep, k rune
		_, err := fmt.Sscanf(sc.Text(), "Step %c must be finished before step %c can begin.", &dep, &k)
		if err != nil {
			log.Fatalf("format err: %v", err)
		}
		depP := getPoint(dep)
		itemP := getPoint(k)
		depP.children = append(depP.children, itemP)
		itemP.parents = append(itemP.parents, depP)
		ready[dep] = true
		not[k] = true
	}
	stack := []*point{}
	for i := range ready {
		if not[i] {
			continue
		}
		stack = append(stack, getPoint(i))
	}
	workers := makeWorkers(5)
	// init workers
	for _, worker := range workers {
		if len(stack) == 0 { // nothing to assign
			break
		}
		worker.working = stack[0]
		worker.ttl = cost(stack[0].id)
		stack = stack[1:]
	}
	var sec int
	for {
		var notDone bool
		// work
		for _, worker := range workers {
			if worker.working == nil {
				continue
			}
			notDone = true
			worker.ttl--
			if worker.ttl == 0 { // this job done
				for _, child := range worker.working.children {
					if len(child.parents) == 1 { // idP is the only parent
						stack = append(stack, child)
						continue
					}
					// remove this idP parent from child
					newParents := []*point{}
					for _, parent := range child.parents {
						if parent.id == worker.working.id {
							continue
						}
						newParents = append(newParents, parent)
					}
					child.parents = newParents
				}
				worker.working = nil
			}
		}
		// work assign
		for _, worker := range workers {
			if len(stack) == 0 { // nothing to assign
				continue
			}
			notDone = true
			if worker.working != nil { // working now
				continue
			}
			worker.working = stack[0]
			worker.ttl = cost(stack[0].id)
			stack = stack[1:]
		}
		if !notDone {
			break
		}
		sec++
	}
	log.Printf("answer is %d", sec)
}

type point struct {
	id       rune
	children []*point
	parents  []*point
}

type worker struct {
	working *point
	ttl     int
}

func getPoint(id rune) *point {
	if p, ok := dict[id]; ok {
		return p
	}
	p := &point{
		id: id,
	}
	dict[id] = p
	return p
}

func cost(i rune) int {
	return int(i-'A') + 1 + 60
}

func makeWorkers(n int) (workers []*worker) {
	for i := 0; i < n; i++ {
		workers = append(workers, &worker{})
	}
	return
}
