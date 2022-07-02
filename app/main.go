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

func Merge(left, right []int) []int {
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = right[1:]
		} else {
			result = append(result, right[0])
			right = left[1:]
		}
	}
	return append(result, append(left, right...)...)
}

func MergeSort(arr []int) []int {
	if len(arr) > 1 {
		n := len(arr) / 2
		arr = Merge(MergeSort(arr[:n]), MergeSort(arr[n:]))
		fmt.Println(arr)
	}
	return arr
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Scan()
	// 数式
	length := nextInt()

	var arr []int
	for i := 0; i < length; i++ {
		arr = append(arr, nextInt())
	}

	fmt.Println(MergeSort(arr))

}
