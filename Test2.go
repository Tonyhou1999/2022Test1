package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	var size int = int(16)
	size, err = strconv.Atoi(os.Args[1])
	if err != nil {
		//This will be executed if something is not right
		fmt.Println(err)
	}
	var numbers []int
	for i := 0; i < size; i++ {
		numbers[i] = rand.Intn(5)
	}
	integers := make([]int, size)
	for i := range integers {
		integers[i] = rand.Intn(20)
	}
}
