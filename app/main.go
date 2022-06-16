package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 特定の文字を除去する
func removeParticularCharacter(target string, delimiter string, len int) []string {
	var slices []string
	splited := strings.SplitN(target, delimiter, len)
	for _, str := range splited {
		if str != "" {
			slices = append(slices, str)
		}
	}
	return slices
}

func convertingDividedInteger(target, delimiter string, len int) []int {
	var slices []int
	splited := removeParticularCharacter(target, delimiter, len)
	for i := range splited {
		var value int
		value, _ = strconv.Atoi(splited[i])
		slices = append(slices, value)
	}
	return slices
}

func getNum() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	num := convertingDividedInteger(str, " ", -1)
	return num
}

/*
// 文字取得
func getWord() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}


// 小数の場合は一旦分数にしてから計算を行う(big.NewRat)を使用する
*/
func main() {
	var line []int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := scanner.Text()
	num, _ := strconv.Atoi(str)

	for i := 0; i < num; i++ {
		scanner.Scan()
		str := scanner.Text()
		elevator, _ := strconv.Atoi(str)
		line = append(line, elevator)
	}
	result := 0
	for i, v := range line {
		if i == 0 {
			result += v - 1
		} else if v > line[i-1] {
			result += v - line[i-1]
		} else {
			result += line[i-1] - v
		}
	}
	fmt.Println(result)
}
