package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	a := nextInt()
	b := nextInt()
	c := nextInt()

	arr := []int{a, b, c}
	sort.Ints(arr)

	if arr[1] == b {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
