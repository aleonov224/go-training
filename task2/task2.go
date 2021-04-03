/* Compresseo/decompression task
* Usage test2 < file_to_compress > copressed_file
* test2 -d < copressed_file > decompressed_file
*
* Examples:
* echo -n "кооооооординатааааа3" | task2
* echo -n к#7#ординат#5#а3 | task2 -d
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if len(os.Args) == 2 && strings.HasSuffix(os.Args[1], "d") {
			qstr := Decompress(input.Text())
			fmt.Println(qstr)
		} else {
			qstr := Compress(input.Text())
			fmt.Println(qstr)
		}
	}
}

func Decompress(s string) string {
	var out string = ""
	var letter rune
	runeS := []rune(s)

	for i := 0; i < len(runeS); i++ {
		letter = runeS[i]
		if letter == '#' {
			clen := ""
			var cr rune
			var foundSize bool = false
			var foundRune bool = false
			for j := i + 1; j < len(runeS); j++ {
				if unicode.IsDigit(runeS[j]) {
					clen += string(runeS[j])
					foundSize = true
				} else if (runeS[j] == '#') && ((j + 1) < len(runeS)) && foundSize {
					cr = runeS[j+1]
					foundRune = true
				}
				if foundRune && foundSize {
					i = j + 1
					rcount, _ := strconv.ParseInt(clen, 10, 64)
					out += Repeat(rcount, cr)
					break
				}
			}
			if !(foundSize || foundRune) {
				out += string(letter)
			}
		} else {
			out += string(letter)
		}
	}
	return out
}

func Repeat(count int64, c rune) string {
	var rs string = ""
	var i int64 = 0
	for i < count {
		rs += string(c)
		i++
	}
	return rs
}

func Compress(s string) string {
	var buf string = ""
	var out string = ""
	var letter rune
	runeS := []rune(s)
	for i, count := 0, 0; i < len(runeS); i++ {
		letter = runeS[i]
		count += 1
		if count > 1 {
			//next not equals to previous ans buf length less of 5
			if !strings.HasSuffix(buf, string(letter)) && len([]rune(buf)) <= 4 {
				out += buf
				buf = ""
				count = 1
			}
			//need to compress
			if !strings.HasSuffix(buf, string(letter)) && len([]rune(buf)) >= 5 {
				out += Compcalc(buf)
				buf = ""
				count = 1
			}
		}
		buf += string(letter)
	}
	runeBuf := []rune(buf)
	if len(runeBuf) >= 5 {
		out += Compcalc(buf)
	} else {
		out += buf
	}
	return out
}

func Compcalc(s string) string {
	var res string = ""
	rbuf := []rune(s)
	res += "#"
	res += fmt.Sprintf("%d", len(rbuf))
	res += "#"
	res += string(rbuf[len(rbuf)-1])
	return res
}
