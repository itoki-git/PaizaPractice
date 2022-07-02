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

func Min(arr []int) int {
	min := arr[0]
	result := 0
	for index, value := range arr {
		if min > value {
			min = value
			result = index
		}
	}
	return result
}

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		index := Min(arr[i:])
		arr[i], arr[i+index] = arr[i+index], arr[i]
	}
	return arr
}

func main() {
	sc.Split(bufio.ScanWords)
	// 数式
	length := nextInt()

	var arr []int
	for i := 0; i < length; i++ {
		arr = append(arr, nextInt())
	}

	fmt.Println(SelectionSort(arr))

}
