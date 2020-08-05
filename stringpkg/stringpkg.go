package stringpkg

import (
	"strconv"
	"strings"
)

func itoa(k int) string{
	var isnegative bool
	var i int
	if k == 0 {
		return "0"
	}
	if k == -9223372036854775808 {
		return "-9223372036854775808"
	}
	if k < 0 {
		k *= -1
		isnegative = true
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
	if isnegative == true {
		sl[i - 1] = "-"
	}
	return strings.Join(sl, "")
}

func stringpkg(str string) string {
	sl := strings.Split(str, "")
	var result strings.Builder
	lenstr := len(sl)
	for i := 0; i < lenstr; i++ {
		if !(sl[i] >= "0" && sl[i] <= "9") {
			if sl[i] == "\\" && i + 1 < lenstr {
				i++
			}
			if i + 1 < lenstr && (sl[i + 1] >= "0" && sl[i + 1] <= "9") {
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