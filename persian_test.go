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
