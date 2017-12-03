package persian

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toPersianDigits(t *testing.T) {
	str := "123salam456"
	assert.Equal(t, "۱۲۳salam۴۵۶", ToPersianDigits(str))
}

func Test_toPersianDigitsFromInt(t *testing.T) {
	value := 123
	assert.Equal(t, "۱۲۳", ToPersianDigitsFromInt(value))
}

func Test_toEnglishDigits(t *testing.T) {
	str := "۱۲۳salam۴۵۶"
	assert.Equal(t, "123salam456", ToEnglishDigits(str))
}

func Test_onlyEnglishNumbers(t *testing.T) {
	value := "salam123hello456۱۲۳سلام۴۵۶"
	assert.Equal(t, "123456", OnlyEnglishNumbers(value))
}

func Test_onlyPersianNumbers(t *testing.T) {
	value := "salam123hello456۱۲۳سلام۴۵۶"
	assert.Equal(t, "۱۲۳۴۵۶", OnlyPersianNumbers(value))
}

func Test_switchToPersianKey(t *testing.T) {
	value := "sghl o,fd ? o,fl llk,k"
	assert.Equal(t, "سلام خوبی ؟ خوبم ممنون", SwitchToPersianKey(value))
}

func Test_switchToEnglishKey(t *testing.T) {
	value := "اثغ صاشفس عح ؟"
	assert.Equal(t, "hey whats up ?", SwitchToEnglishKey(value))
}

func Test_Currency(t *testing.T) {
	value := "123۴۵۶7"
	assert.Equal(t, "۱،۲۳۴،۵۶۷", Currency(value))
}

func Test_toman(t *testing.T) {
	value := "123۴۵۶7"
	assert.Equal(t, "۱،۲۳۴،۵۶۷ تومان", Toman(value))
}

func Test_rial(t *testing.T) {
	value := "123۴۵۶7"
	assert.Equal(t, "۱،۲۳۴،۵۶۷ ﷼", Rial(value))
}

func Test_fixArabic(t *testing.T) {
	value := "علي"
	assert.Equal(t, "علی", FixArabic(value))
}

func Test_checkIsEnglish(t *testing.T) {
	value := "علي"
	assert.Equal(t, false, CheckIsEnglish(value))

	value = "ali"
	assert.Equal(t, true, CheckIsEnglish(value))
}
