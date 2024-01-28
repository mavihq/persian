package persian

import (
	"fmt"
	"regexp"
	"strings"
)

// ToPersianDigits Converts all English digits in the string to Persian digits.
func ToPersianDigits(text string) string {
	return strings.NewReplacer(
		"0", "۰",
		"1", "۱",
		"2", "۲",
		"3", "۳",
		"4", "۴",
		"5", "۵",
		"6", "۶",
		"7", "۷",
		"8", "۸",
		"9", "۹",
	).Replace(text)
}

// ToPersianDigitsFromInt Converts integer value to string with Persian digits.
func ToPersianDigitsFromInt(value int) string {
	return ToPersianDigits(fmt.Sprintf("%d", value))
}

// ToEnglishDigits Converts all Persian digits in the string to English digits.
func ToEnglishDigits(text string) string {
	return strings.NewReplacer(
		"۰", "0",
		"۱", "1",
		"۲", "2",
		"۳", "3",
		"۴", "4",
		"۵", "5",
		"۶", "6",
		"۷", "7",
		"۸", "8",
		"۹", "9",
		// Arabic Number
		"٠", "0",
		"١", "1",
		"٢", "2",
		"٣", "3",
		"٤", "4",
		"٥", "5",
		"٦", "6",
		"٧", "7",
		"٨", "8",
		"٩", "9",
	).Replace(text)
}

// OnlyEnglishNumbers extracts only English digits from string.
func OnlyEnglishNumbers(text string) string {
	re := regexp.MustCompile("[^0-9.]")
	return re.ReplaceAllLiteralString(text, "")
}

// OnlyPersianNumbers extracts only Persian digits from string.
func OnlyPersianNumbers(text string) string {
	re := regexp.MustCompile("[^۰-۹.]")
	return re.ReplaceAllLiteralString(text, "")
}

// OnlyNumbers extracts only digits from string.
func OnlyNumbers(text string) string {
	re := regexp.MustCompile("[^۰-۹0-9.]")
	return re.ReplaceAllLiteralString(text, "")
}

// OnlyPersianAlpha extracts only persian alphabetes from string.
func OnlyPersianAlpha(text string) string {
	re := regexp.MustCompile("[^\u0600-\u06FF.]")
	return re.ReplaceAllLiteralString(text, "")
}

// SwitchToPersianKey converts English chars to their equivalent Persian char on keyboard.
func SwitchToPersianKey(text string) string {
	return strings.NewReplacer(
		"q", "ض",
		"w", "ص",
		"e", "ث",
		"r", "ق",
		"t", "ف",
		"y", "غ",
		"u", "ع",
		"i", "ه",
		"o", "خ",
		"p", "ح",
		"\\[", "ج",
		"\\]", "چ",
		"a", "ش",
		"s", "س",
		"d", "ی",
		"f", "ب",
		"g", "ل",
		"h", "ا",
		"j", "ت",
		"k", "ن",
		"l", "م",
		";", "ک",
		"'", "گ",
		"z", "ظ",
		"x", "ط",
		"c", "ز",
		"v", "ر",
		"b", "ذ",
		"n", "د",
		"m", "پ",
		",", "و",
		"?", "؟",
	).Replace(text)
}

// SwitchToEnglishKey converts Persian chars to their equivalent English char on keyboard.
func SwitchToEnglishKey(text string) string {
	return strings.NewReplacer(
		"ض", "q",
		"ص", "w",
		"ث", "e",
		"ق", "r",
		"ف", "t",
		"غ", "y",
		"ع", "u",
		"ه", "i",
		"خ", "o",
		"ح", "p",
		"ج", "\\[",
		"چ", "\\]",
		"ش", "a",
		"س", "s",
		"ی", "d",
		"ب", "f",
		"ل", "g",
		"ا", "h",
		"ت", "j",
		"ن", "k",
		"م", "l",
		"ک", ";",
		"گ", "'",
		"ظ", "z",
		"ط", "x",
		"ز", "c",
		"ر", "v",
		"ذ", "b",
		"د", "n",
		"پ", "m",
		"و", ",",
		"؟", "?",
	).Replace(text)
}

// FixArabic used for converting Arabic characters to Persian.
func FixArabic(text string) string {
	return strings.NewReplacer(
		"ي", "ی",
		"ك", "ک",
		"دِ", "د",
		"بِ", "ب",
		"زِ", "ز",
		"ذِ", "ذ",
		"ِشِ", "ش",
		"ِسِ", "س",
		"ى", "ی",
	).Replace(text)
}

// Normalize used for Normalize Persian for sort and equality check.
// TODO: Complete list according to Persian Collation
func Normalize(text string) string {
	return strings.NewReplacer(
		"ي", "ی",
		"ك", "ک",
		"دِ", "د",
		"بِ", "ب",
		"زِ", "ز",
		"ذِ", "ذ",
		"ِشِ", "ش",
		"ِسِ", "س",
		"‍", " ",
		"‌", " ",
		"ى", "ی",
		"ٱ", "ا",
		"آ", "ا",
		"ء", "ا",
		"ئ", "ی",
		"أ", "ا",
		"ة", "ه",
	).Replace(text)
}

// Reverse reverses the given string.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Currency formats number to Persian currency style.
func Currency(amount string) string {
	countThrees := 0
	out := ""
	_amount := []rune(OnlyNumbers(amount))
	for i := len(_amount) - 1; i >= 0; i-- {
		if countThrees == 3 {
			out += "،" + string(_amount[i])
			countThrees = 1
		} else {
			out += string(_amount[i])
			countThrees++
		}
	}
	return ToPersianDigits(Reverse(out))
}

// Toman formats number to Persian currency style with تومان postfix.
func Toman(amount string) string {
	return Currency(amount) + " تومان"
}

// Rial  formats number to Persian currency style with ﷼ postfix.
func Rial(amount string) string {
	return Currency(amount) + " ﷼"
}

func CheckIsEnglish(text string) bool {
	for _, r := range text {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

// finglish generator

func Finglish(farsi string) (finglish string) {

	f := newF2F()
	finglish = f.SimpleF2F(farsi)

	return
}

type f2F struct {
	Words  map[rune][]string
	Vowels []rune
}

func newF2F() *f2F {
	f := &f2F{
		Words: map[rune][]string{
			'ا': {"a", "a"},
			'آ': {"a", "a"},
			'ئ': {"a"},
			'ء': {"a"},
			'ب': {"b"},
			'پ': {"p"},
			'ت': {"t"},
			'ث': {"s"},
			'ج': {"j"},
			'چ': {"ch"},
			'ح': {"h"},
			'خ': {"kh"},
			'د': {"d"},
			'ذ': {"z"},
			'ر': {"r"},
			'ز': {"z"},
			'ژ': {"zh"},
			'س': {"s"},
			'ش': {"sh"},
			'ص': {"s"},
			'ض': {"z"},
			'ط': {"t"},
			'ظ': {"z"},
			'ع': {"", "a"},
			'غ': {"gh"},
			'ف': {"f"},
			'ق': {"gh"},
			'ک': {"k"},
			'گ': {"g"},
			'ل': {"l"},
			'م': {"m"},
			'ن': {"n"},
			'و': {"v", "o"},
			'ه': {"h"},
			'ی': {"y", "i"},
			'ي': {"y", "i"},
		},
	}

	//f.Vowels = []rune{'ﺕَ', 'ﺕِ', 'ﺕُ'}
	//f.Words['ﺕَ'] = []string{"a", "a"}
	//f.Words['ﺕِ'] = []string{"e", "e"}
	//f.Words['ﺕُ'] = []string{"o", "o"}
	//f.Words['ﻁْ'] = []string{"."}

	return f
}

func (f *f2F) SimpleF2F(word string) string {
	result := ""
	state := 0
	wordsRunes := []rune(word)

	for i := 0; i < len(wordsRunes); i++ {
		ch := wordsRunes[i]
		if ch == ' ' {
			result += " "
		}

		_, ok := f.Words[ch]
		if !ok {
			continue
		}

		if f.Words[ch][0] == "." {
			state = 0
			continue
		}

		if state == 0 {
			result += f.Words[ch][0]
			state = 1
		} else if state == 1 {
			if len(f.Words[ch]) > 1 {
				result += f.Words[ch][1]
			} else {
				result += "a"
				i--
			}
			state = 2
		} else if state == 2 {
			result += f.Words[ch][0]
			state = 3
		} else if state == 3 {
			if len(f.Words[ch]) > 1 {
				result += f.Words[ch][1]
				state = 2
			} else {
				result += f.Words[ch][0]
				state = 4
			}
		} else if state == 4 {
			if len(f.Words[ch]) > 1 {
				result += f.Words[ch][1]
				state = 2
			} else {
				result += f.Words[ch][0]
				state = 5
			}
		} else if state >= 5 {
			if len(f.Words[ch]) > 1 {
				result += f.Words[ch][1]
				state = 2
			} else {
				result += f.Words[ch][0]
				state++
			}
		}
	}

	result = strings.ReplaceAll(result, "aa", "a")
	result = strings.ReplaceAll(result, "ao", "o")
	result = strings.ReplaceAll(result, "ae", "e")

	return result
}
