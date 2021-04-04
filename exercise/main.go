package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	var firstLine string
	for sc.Scan() {
		secondLine := sc.Text()
		wholeLine := firstLine + secondLine
		if findsum49(wholeLine) {
			log.Println("got it")
			return
		}
		firstLine = secondLine
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
}

var primes = getInitPrimes()

func find(str string) bool {
	for i := 0; i <= len(str)-10; i++ {
		n, err := strconv.Atoi(str[i : i+10])
		if err != nil {
			log.Fatal(err)
		}
		if isPrimeAlone(n, primes) {
			log.Printf("answer is %d", n)
			return true
		}
	}
	return false
}
func findsum49(str string) bool {
	for i := 0; i <= len(str)-10; i++ {
		sum := 0
		for j := 0; j < 10; j++ {
			v := str[i+j]
			n, err := strconv.Atoi(string(v))
			if err != nil {
				log.Fatal(err)
			}
			sum += n
		}
		if sum == 49 {
			log.Printf("answer is %s", str[i:i+10])
			return true
		}
	}
	return false
}

func getInitPrimes() []int {
	// init primes
	primes := []int{2, 3}
	curr := 3
	for i := 0; i < 5000000; i++ {
		if curr > 99999 {
			break
		}
		curr = nextPrime(curr, primes)
		primes = append(primes, curr)
	}
	return primes
}
func nextPrime(n int, primes []int) int {
	for {
		n += 2
		if isPrime(n, primes) {
			return n
		}
	}
}

func isPrime(n int, primes []int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	if sqrt*sqrt == n {
		return false
	}
	for _, p := range primes {
		if p > sqrt {
			return true
		}
		if n%p == 0 {
			return false
		}
	}
	return false
}

func isPrimeAlone(n int, primes []int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	if sqrt*sqrt == n {
		return false
	}
	for _, p := range primes {
		if p > sqrt {
			return true
		}
		if n%p == 0 {
			return false
		}
	}
	return false
}
