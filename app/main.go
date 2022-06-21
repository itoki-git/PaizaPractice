package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)

	sc.Scan()
	N := sc.Text()
	var before = map[string]string{"A": "4", "E": "3", "G": "6", "I": "1", "O": "0", "S": "5", "Z": "2"}

	strArray := strings.Split(N, "")

	for _, str := range strArray {
		if val, ok := before[str]; ok {
			fmt.Print(val)
		} else {
			fmt.Print(str)
		}
	}

}
