package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
	for i := range not {
		delete(ready, i)
	}
	out := strings.Builder{}
	for len(ready) != 0 {
		id := findFirst(ready)
		delete(ready, id)
		idP := getPoint(id)
		for _, child := range idP.children {
			if len(child.parents) == 1 { // idP is the only parent
				ready[child.id] = true // now this kid is ready
				continue
			}
			// remove this idP parent from child
			newParents := []*point{}
			for _, parent := range child.parents {
				if parent.id == idP.id {
					continue
				}
				newParents = append(newParents, parent)
			}
			child.parents = newParents
		}
		_, err := fmt.Fprintf(&out, "%c", id)
		if err != nil {
			log.Fatalf("fprint err: %v", err)
		}
	}
	log.Printf("answer is %s", out.String())
}

type point struct {
	id       rune
	children []*point
	parents  []*point
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

func findFirst(ready map[rune]bool) rune {
	min := rune(10000)
	for i := range ready {
		if i < min {
			min = i
		}
	}
	return min
}
