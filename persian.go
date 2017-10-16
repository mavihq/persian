package persian

import "regexp"

func ToPersianDigits(text string) string {
	var persian = map[string]string{
		"0": "۰",
		"1": "۱",
		"2": "۲",
		"3": "۳",
		"4": "۴",
		"5": "۵",
		"6": "۶",
		"7": "۷",
		"8": "۸",
		"9": "۹",
	}
	re := regexp.MustCompile("[0-9]+")
	out := re.ReplaceAllFunc([]byte(text), func(s []byte) []byte {
		out := ""
		ss := string(s)
		for _, ch := range ss {
			o := persian[string(ch)]
			out = out + o
		}
		return []byte(out)
	})
	return string(out)
}

func ToEnglishDigits(text string) string {
	var persian = map[string]string{
		"۰": "0",
		"۱": "1",
		"۲": "2",
		"۳": "3",
		"۴": "4",
		"۵": "5",
		"۶": "6",
		"۷": "7",
		"۸": "8",
		"۹": "9",
	}
	re := regexp.MustCompile("[۰-۹]+")
	out := re.ReplaceAllFunc([]byte(text), func(s []byte) []byte {
		out := ""
		ss := string(s)
		for _, ch := range ss {
			o := persian[string(ch)]
			out = out + o
		}
		return []byte(out)
	})
	return string(out)
}
