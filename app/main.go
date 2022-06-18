package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	a := make([][]int, N)
	for i := 0; i < N; i++ {
		var arr []int
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				arr = append(arr, 1)
			} else {
				arr = append(arr, a[i-1][j-1]+a[i-1][j])

			}
		}
		a[i] = arr
	}

	for _, i := range a {
		for _, j := range i {
			fmt.Print(j)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
