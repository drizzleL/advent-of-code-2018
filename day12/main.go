package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func initInput() (string, map[string]byte) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	i := 0
	sc := bufio.NewScanner(f)
	state := ""
	ruleDict := map[string]byte{}
	for sc.Scan() {
		switch i {
		case 0:
			_, err := fmt.Sscanf(sc.Text(), "initial state: %s", &state)
			if err != nil {
				log.Fatal(err)
			}
		case 1:
		default:
			var rule string
			var chr byte
			_, err := fmt.Sscanf(sc.Text(), "%5s => %c", &rule, &chr)
			if err != nil {
				log.Fatal(err)
			}
			ruleDict[rule] = chr
		}
		i++
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	return state, ruleDict
}

func main() {
	var str string
	str, ruleDict = initInput()
	state := State{
		str: str,
	}
	for i := 0; i < 500; i++ {
		if i == 80 {
			break
		}
		if re, ok := resultDict[state.str]; ok {
			log.Println(re, i)
			break
		}
		next, _ := nextState(state)
		resultDict[state.str] = Result{
			idx: state.startIndex,
		}
		state = next
	}
	log.Printf("answer is %s %d, count %d", state.str, state.startIndex, state.Count())
}

func nextState(state State) (State, int) {
	nextStr, diff := nextStateDiff(state.str)
	return State{
		str:        nextStr,
		startIndex: state.startIndex + diff,
	}, diff
}

type State struct {
	str        string
	startIndex int
}

func (self State) Count() (sum int) {
	for i, c := range self.str {
		if c != '#' {
			continue
		}
		sum += self.startIndex + i
	}
	return
}

var ruleDict = map[string]byte{}

func nextStateDiff(state string) (string, int) {
	next := []byte{}
	for i := -2; i < len(state)+2; i++ {
		var str string
		switch {
		case i <= 1:
			for j := 0; j < 2-i; j++ {
				str += "."
			}
			str += state[0 : i+3]
		case i >= len(state)-2:
			for j := 0; j < i-len(state)+3; j++ {
				str += "."
			}
			str = state[i-2:len(state)] + str
		default:
			str = state[i-2 : i+3]
		}
		ans, ok := ruleDict[str]
		if !ok {
			ans = '.'
		}
		next = append(next, ans)
	}
	leftCnt := 0
	next = bytes.TrimLeftFunc(next, func(c rune) bool {
		if c == '.' {
			leftCnt++
			return true
		}
		return false
	})
	return strings.TrimRight(string(next), "."), leftCnt - 2
}

var resultDict = map[string]Result{}

type Result struct {
	idx int
}
