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
	inpNum, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Printf("%q looks like is not integer number or too big\n", os.Args[1])
		fmt.Println(err.Error())
		return
	}
	var x int64 = inpNum
	if inpNum < 0 {
		x = -x
	}
	if IsAutomorphic(x) {
		fmt.Printf("%d is automorphic number\n", inpNum)
	} else {
		fmt.Printf("%d is not automorphic number\n", inpNum)
	}
}

func IsAutomorphic(num int64) bool {
	sqn := num * num
	origStr := fmt.Sprintf("%d", num)
	sqStr := fmt.Sprintf("%d", sqn)
	return strings.HasSuffix(sqStr, origStr)
}
