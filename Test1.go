package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter an integer ")
	//reader := bufio.NewReader(os.Stdin)
	var i int
	fmt.Scanf("%d", &i)
	fmt.Println(i)
	integers := make([]int, i+1)
	var numbers []int
	for j := 0; j < i; j++ {
		numbers[j] = j
	}
	fmt.Println("The length is", len(integers))

}
