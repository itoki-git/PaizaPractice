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

func BubbleSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			fmt.Println(arr)
		}
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

	fmt.Println(BubbleSort(arr))

}
