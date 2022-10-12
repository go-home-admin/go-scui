package base

import (
	"unicode"
)

// 标识符
func IsIdentifier(r int32) bool {
	if unicode.IsLetter(r) {
		return true
	} else if unicode.IsDigit(r) {
		return true
	} else if 95 == r {
		return true
	}

	return false
}

// 判断是否以某字符串作为结尾 HasSuffix
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// 判断是否以某字符串作为开始 HasPrefix
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

func InArrString(str string, arr []string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}

// 最后一个有意义的符号是否是start(后面跟踪回车、空格、注释不影响, 前面也不能有start)
func GetLastIsIdentifier(l []*word, start string) (bool, int) {
	ok := false
	i := 0
	var w *word
	for i, w = range l {
		if w.Ty == wordT_line {
			return ok, i
		} else if w.Str == start {
			ok = true
		} else if ok && !(w.Ty == wordT_doc || InArrString(w.Str, []string{" ", "\t"})) {
			ok = false
		}
	}

	return ok, i
}

// 获取第一个单词
func GetFistWord(l []*word) (string, int) {
	for i, w := range l {
		if w.Ty == wordT_word {
			return w.Str, i
		}
	}
	return "", len(l)
}

// 获取第n个单词
func GetWord(l []*word, n int) (string, int) {
	for i, w := range l {
		if w.Ty == wordT_word {
			n--
			if n <= 0 {
				return w.Str, i
			}
		}
	}
	return "", len(l)
}

// 获取第一个字符(不包括空格, 换行符, 制表)
func GetFistStr(l []*word) (string, int) {
	for i, w := range l {
		if !InArrString(w.Str, []string{" ", "\n", "\t"}) {
			return w.Str, i
		}
	}
	return "", len(l)
}

// 获取换行前所有词
func GetStrAtEnd(l []*word) (string, int) {
	s := ""
	for i, w := range l {
		if w.Str == "\n" {
			return s, i
		} else if w.Str != " " {
			s = s + w.Str
		}
	}
	return "", len(l)
}

// 获取下一行开始
func NextLine(l []*word) int {
	for i, w := range l {
		if w.Ty == wordT_line {
			return i
		}
	}
	return len(l)
}

// 获取某字符串后第一个单词
func GetFistWordBehindStr(l []*word, behind string) (string, int) {
	init := false
	for i, w := range l {
		if w.Ty == wordT_word {
			if init {
				return w.Str, i
			} else if w.Str == behind {
				init = true
			}
		}
	}
	return "", len(l)
}

// 括号引用起来的块, 或者第一行没有块就换行
func GetBracketsOrLn(l []*word, start, end string) (int, int, bool) {
	for i, w := range l {
		if w.Ty == wordT_line {
			return 0, i, false
		} else if w.Ty == wordT_division && w.Str == start {
			startInt, endInt := GetBrackets(l, start, end)
			return startInt, endInt, true
		}
	}

	return 0, 0, false
}

// 括号引用起来的块, 不限制分隔符
func GetBracketsAtString(l []*word, start, end string) (int, int) {
	var startInt, endInt int

	bCount := 0
	for i, w := range l {
		if bCount == 0 {
			if w.Str == start {
				startInt = i
				bCount++
			}
		} else {
			switch w.Str {
			case start:
				bCount++
			case end:
				bCount--
				if bCount <= 0 {
					endInt = i
					return startInt, endInt
				}
			}
		}
	}

	return startInt, endInt
}

// 括号引用起来的块, 词性必须是分隔符
// 返回是开始偏移和结束偏移
func GetBrackets(l []*word, start, end string) (int, int) {
	var startInt, endInt int

	bCount := 0
	for i, w := range l {
		if bCount == 0 {
			if w.Ty == wordT_division && w.Str == start {
				startInt = i
				bCount++
			}
		} else {
			if w.Ty == wordT_division {
				switch w.Str {
				case start:
					bCount++
				case end:
					bCount--
					if bCount <= 0 {
						endInt = i
						return startInt, endInt
					}
				}
			}
		}
	}

	return startInt, endInt
}

// 丢掉注释
func WlToString(l []*word) string {
	s := ""
	for _, w := range l {
		if w.Ty != wordT_doc {
			s = s + w.Str
		}
	}
	return s
}
