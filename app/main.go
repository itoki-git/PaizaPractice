package main

import (
	"bufio"
	"fmt"
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

func toCharStr(i int) string {
	return string('A' - 1 + i)
}

func main() {
	sc.Split(bufio.ScanWords)

	a := nextInt()
	fmt.Print(strings.ToLower(toCharStr(a - 96)))

}
