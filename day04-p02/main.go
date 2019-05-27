package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type SortByTime []*TimeCmd

func (self SortByTime) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}
func (self SortByTime) Less(i, j int) bool {
	return self[i].t.Unix() < self[j].t.Unix()
}
func (self SortByTime) Len() int {
	return len(self)
}

type TimeCmd struct {
	t time.Time
	s string
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalf("can't open file: %v", err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	cmds := []*TimeCmd{}
	for sc.Scan() {
		t, err := time.Parse("2006-01-02 15:04", sc.Text()[1:17])
		if err != nil {
			log.Fatalf("time format: %v", err)
		}
		cmds = append(cmds, &TimeCmd{
			t: t,
			s: sc.Text()[19:],
		})
	}
	sort.Sort(SortByTime(cmds))
	maxID := struct {
		id  int
		max int
	}{0, 0}
	var guardID, startAt int
	dict := map[int]map[int]int{}
	for _, cmd := range cmds {
		switch cmd.s {
		case "falls asleep":
			startAt = cmd.t.Minute()
		case "wakes up":
			for i := startAt; i < cmd.t.Minute(); i++ {
				if dict[guardID] == nil { // init if not exist
					dict[guardID] = map[int]int{}
				}
				dict[guardID][i]++
				if dict[guardID][i] > maxID.max {
					maxID.max = dict[guardID][i]
					maxID.id = guardID * i
				}
			}
		default:
			guardID = findGuardID(cmd.s, guardID)
		}
	}
	log.Printf("answer is %d", maxID.id)
}

func findGuardID(str string, defaultID int) int {
	reg, err := regexp.Compile(`Guard #(\d+)`)
	if err != nil {
		log.Fatalf("regexp err: %v", err)
	}
	strs := reg.FindStringSubmatch(str)
	if len(strs) <= 1 || strs[1] == "" {
		return defaultID
	}
	guardID, err := strconv.Atoi(strs[1])
	if err != nil {
		log.Fatalf("guard_id parse failed: %v", err)
	}
	return guardID
}
