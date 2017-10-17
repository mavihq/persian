# persian
[![cover.run go](https://cover.run/go/github.com/mavihq/persian.svg)](https://cover.run/go/github.com/mavihq/persian)
[![Go Report Card](https://goreportcard.com/badge/github.com/mavihq/persian)](https://goreportcard.com/report/github.com/mavihq/persian)
[![GoDoc](https://godoc.org/github.com/mavihq/persian?status.svg)](https://godoc.org/github.com/mavihq/persian)

Some utilities for persian language in go.

## Installation
```
go get github.com/mavihq/persian
```

## API
### .ToPersianDigits
Converts all English digits in the string to Persian digits.
```
persian.ToPersianDigits("123salam456")
=> "۱۲۳salam۴۵۶"
```

### .ToPersianDigitsFromInt
Converts integer value to string with Persian digits.
```
persian.ToPersianDigitsFromInt(123)
=> "۱۲۳"
```

### .ToEnglishDigits
Converts all Persian digits in the string to English digits.
```
persian.ToEnglishDigits("۱۲۳salam۴۵۶")
=> "123salam456"
```

### .OnlyEnglishNumbers
Extracts only English digits from string.
```
persian.ToEnglishDigits("123salam۴۵۶")
=> "123"
```

### .OnlyPersianNumbers
Extracts only Persian digits from string.
```
persian.ToEnglishDigits("123salam۴۵۶")
=> "۴۵۶"
```

### .SwitchToPersianKey
Converts English chars to their equivalent Persian char on keyboard.
```
persian.ToEnglishDigits("sghl o,fd ? o,fl llk,k")
=> "سلام خوبی ؟ خوبم ممنون"
```

### .SwitchToEnglishKey
Converts Persian chars to their equivalent English char on keyboard.
```
persian.ToEnglishDigits("اثغ صاشفس عح ؟")
=> "hey whats up ?"
```

### .Currency
Formats number to Persian currency style.
```
persian.Currency("1234567")
=> "۱،۲۳۴،۵۶۷"
```

### .Toman
Formats number to Persian currency style with تومان postfix.
```
persian.Toman("1234567")
=> "۱،۲۳۴،۵۶۷ تومان"
```

### .Rial
Formats number to Persian currency style with ﷼ postfix.
```
persian.Rial("1234567")
=> "۱،۲۳۴،۵۶۷ ﷼"
```

### .FixArabic
Used for converting Arabic characters to Persian.
```
persian.FixArabic("علي")
=> "علی"
```