package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	input := os.Args[1]
	value, _ := strconv.ParseInt(input, 10, 64)
	fmt.Printf("The scanned input from the commandline is %d\n", int(value))
	var slicedint = make([]int, int(value+1))
	var slicedintcopy = make([]int, int(value+1))
	for i := 0; i < len(slicedint); i++ {
		slicedint[i] = rand.Intn(int(value))
		slicedintcopy[i] = slicedint[i]
	}
	start := time.Now()
	sort.Ints(slicedint)
	slicesort := time.Since(start)
	startstable := time.Now()
	sort.SliceStable(slicedintcopy, func(i, j int) bool { return slicedintcopy[i] < slicedintcopy[j] })
	endstable := time.Since(startstable)
	fmt.Printf("It takes %d nanoseconds to run Sort to sort the array, and %d nanoseconds using stable to sor"+
		"t the array\n", slicesort.Nanoseconds(), endstable.Nanoseconds())

}
