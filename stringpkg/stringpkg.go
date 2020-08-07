package main

import (
	"sort"
	"strconv"
	"strings"
)

func Itoa(k int) string{
	var isNegative bool
	var i int
	if k == 0 {
		return "0"
	}
	if k == -9223372036854775808 {
		return "-9223372036854775808"
	}
	if k < 0 {
		k *= -1
		isNegative = true
		i++
	}
	for temp := k; temp != 0; temp /= 10 {
		i++
	}
	var sl = make([]string, i, i)
	for k != 0 {
		sl[i - 1] = string(rune(k % 10 + '0'))
		i--
		k /= 10
	}
	if isNegative == true {
		sl[i - 1] = "-"
	}
	return strings.Join(sl, "")
}

func StringPkg(str string) string {
	sl := strings.Split(str, "")
	var result strings.Builder
	lenStr := len(sl)
	for i := 0; i < lenStr; i++ {
		if !(sl[i] >= "0" && sl[i] <= "9") {
			if sl[i] == "\\" && i + 1 < lenStr {
				i++
			}
			if i + 1 < lenStr && (sl[i + 1] >= "0" && sl[i + 1] <= "9") {
				for k, err := strconv.Atoi(sl[i + 1]); k > 0; k-- {
					if err != nil {
						println(err)
						break
					}
					result.WriteString(sl[i])
				}
				i++
			} else {
				result.WriteString(sl[i])
			}
		} else {
			return ""
		}
	}
	return result.String()
}

type words struct {
	word string
	size int
}

func tenWords(text string) []string{
	sli := strings.Split(text, " ")
	result := make([]words, 0, len(sli))
	k := 0
	lenWord := 0
	for _, word := range sli {
		for i := 0; i < lenWord; i++ {
			if word == result[i].word {
				result[i].size++
				k = 1
				break
			}
		}
		if k != 1 {
			result = append(result, words{word, 1})
			lenWord++
		}
		k = 0
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].size > result[j].size
	})
	arr := make([]string, 0, len(sli))
	for i := 0; i < 10 && i < lenWord; i++ {
		arr = append(arr, result[i].word)
	}
	return arr
}

func tenWordsMap(text string) []string{
	var isSet bool
	sli := strings.Split(text, " ")
	cache := make(map[string]int, len(sli))
	for _, word := range sli {
		_, isSet = cache[word]
		if isSet == false {
			cache[word] = 1
		} else {
			cache[word]++
		}
	}
	arr := make([]string, 0, len(cache))
	for word := range cache {
		arr = append(arr, word)
	}
	sort.Slice(arr, func(i, j int) bool {
		return cache[arr[i]] > cache[arr[j]]
	})
	if cap(arr) > 10 {
		arr = arr[0:10]
	}
	return arr
}
