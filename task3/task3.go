/*

Дана текстовая строка с различными разделителями между словами.
Требуется найти количество вхождений каждого слова в этот текст и вывести слова отсортированными по невозрастанию этого количества.
Для каждого слова при выводе в скобках указать количество его вхождений в текст.
Например, для текста:
   	 один, два - это 2, три один два, много слов: один
вывести:
   	 один(3) два(2) это(1) 2(1) три(1) много(1) слов(1)
Слова, которые встречаются в тексте одинаковое количество раз, должны быть выведены в том порядке, в котором они впервые встречаются в тексте.

	Для проверки запустить программу с данными в системном потоке ввода.

	echo "один, два - это 2, три один два, много слов: один" | task3


*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

//word count pair
type kv struct {
	Key   string
	Value int
}

//equal frequency indexes
type ndx_pair struct {
	ValueA int
	ValueB int
}

func main() {
	//firts word index in the source text
	ffound_map := make(map[string]int)

	//word counts map
	count_map := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	position := 1
	for input.Scan() {
		line := input.Text()
		line_words := strings.FieldsFunc(line, func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		})
		for i := 0; i < len(line_words); i++ {
			if len(line_words[i]) > 0 {
				word := line_words[i]
				count_map[word]++
				if ffound_map[word] == 0 {
					ffound_map[word] = position
				}
				position++
			}
		}
	}

	PrintWordsOrder(ffound_map, count_map)
}

func PrintWordsOrder(ffound_map map[string]int, count_map map[string]int) {

	//create array pairs to sort
	var ss []kv
	for k, v := range count_map {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	//prepare equal word's count sub array indexes
	var eq_ndx []ndx_pair
	var begin int = 0
	for i := 0; i < len(ss); i++ {
		if (ss[begin].Value != ss[i].Value) && ((i - begin) == 1) {
			begin = i
		}
		if (ss[begin].Value != ss[i].Value) && ((i - begin) > 1) {
			eq_ndx = append(eq_ndx, ndx_pair{begin, i - 1})
			begin = i
		}
		if (ss[begin].Value == ss[i].Value) && ((i - begin) >= 1) && (len(ss) == (i + 1)) {
			eq_ndx = append(eq_ndx, ndx_pair{begin, i})
		}
	}

	//Sort array parts for equal word frequency
	for i := 0; i < len(eq_ndx); i++ {
		sort.Slice(ss[eq_ndx[i].ValueA:eq_ndx[i].ValueB+1], func(i, j int) bool {
			return ffound_map[ss[i].Key] < ffound_map[ss[j].Key]
		})
	}

	for i := 0; i < len(ss); i++ {
		fmt.Printf("%s(%d) ", ss[i].Key, ss[i].Value)
	}
}
