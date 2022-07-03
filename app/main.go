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
	//sc.Split(bufio.ScanWords)

	// バンドA, バンドBのライブ日を格納
	var liveDays [31][2]int

	// バンドAの入力
	bandA := nextInt()
	for i := 0; i < bandA; i++ {

		liveDays[nextInt()-1][0] = 1
	}

	bandB := nextInt()
	for i := 0; i < bandB; i++ {
		liveDays[nextInt()-1][1] = 1

	}
	sameDays := 0
	for _, day := range liveDays {

		if day[0] == 1 && day[1] == 1 {
			if sameDays%2 == 0 {
				fmt.Println("A")
			} else {
				fmt.Println("B")
			}
			sameDays += 1
		} else if day[0] == 1 {
			fmt.Println("A")
		} else if day[1] == 1 {
			fmt.Println("B")
		} else {
			fmt.Println("x")
		}
	}
}
