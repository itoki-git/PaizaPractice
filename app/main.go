package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

type Mark struct {
	startY int
	startX int
	endY   int
	endX   int
}

func main() {
	sc.Split(bufio.ScanWords)

	a := nextInt()
	b := nextInt()
	arr := make([][]string, a)
	for i := 0; i < a; i++ {
		sc.Scan()
		input := strings.Split(sc.Text(), "")
		arr[i] = input
	}
	count := 0
	var mark Mark

	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			if arr[i][j] == "o" {
				if count == 0 {
					count++
					mark.startX = j
					mark.startY = i
				} else {
					mark.endX = j
					mark.endY = i
				}
			}
		}
	}
	result := math.Abs(float64(mark.startY-mark.endY)) + math.Abs(float64(mark.startX-mark.endX))
	fmt.Print(result)

}
