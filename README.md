# persian
![cover.run go](https://cover.run/go/github.com/mavihq/persian.svg)
Some utilities for persian language in go

## Installation
```
go get github.com/mavihq/persian
```

## API
### .ToPersianDigits()
Converts all English digits in the string to Persian digits
```
persian.ToPersianDigits("123salam456")
=> "۱۲۳salam۴۵۶"
```

### .ToPersianDigitsFromInt()
Converts integer value to string with Persian digits
```
persian.ToPersianDigitsFromInt(123)
=> "۱۲۳"
```

### .ToEnglishDigits()
Converts all Persian digits in the string to English digits
```
persian.ToEnglishDigits("۱۲۳salam۴۵۶")
=> "123salam456"
```