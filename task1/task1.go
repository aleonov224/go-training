/*
 Package automorphic
 Training GO, task1, check automorphic number
*/
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: task1 <number_to_checK_automorphic")
		return
	}
	inpNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q looks like is not number or too big\n", os.Args[1])
		fmt.Println(err.Error())
		return
	}
	if inpNum < 0 {
		fmt.Println("Use positive number please")
		return
	}
	if IsAutomorphic(inpNum) {
		fmt.Printf("%d is automorphic number\n", inpNum)
	} else {
		fmt.Printf("%d is not automorphic number\n", inpNum)
	}
}

func IsAutomorphic(num int) bool {
	sqn := num * num
	origStr := fmt.Sprintf("%d", num)
	sqStr := fmt.Sprintf("%d", sqn)
	return strings.HasSuffix(sqStr, origStr)
}
