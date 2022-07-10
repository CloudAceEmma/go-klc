# KoreanLunarCalendar (한국 양음력 변환)
## Getting Started

#### Overview
Here is a library to convert Korean lunar-calendar to Gregorian calendar.

Korean and Chinese lunar-calendars are not always same.

This follows the Gregorian calendar of KARI(Korea Astronomy and Space Science Institute) - https://astro.kasi.re.kr/life/pageView/8

한국 양음력 변환 (한국천문연구원 기준) - 네트워크 연결 불필요

음력 변환은 1391년 1월 1일 부터 2050년 11월 18일까지 지원

````
Gregorian calendar (1391. 2. 5. ~ 2050. 12. 31) <--> Korean lunar-calendar (1391. 1. 1. ~ 2050. 11. 18)
````

#### To use
Only requires the single go file.

```go
import github.com/chungha/go-klc
```

#### Sample code
1. 양력 2022년 7월 10일을 음력으로 변환
```go

// param : year(년), month(월), day(일)
SetSolarDate(2022, 7, 10)

// Lunar Date (ISO format)
lunar := GetLunarIsoFormat()
fmt.Println(lunar)

// Korean GapJa String
lunarGapja := GetGapjaString()
fmt.Println(lunarGapja)

// Chinese GapJa String
lunarChineseGapja := GetChineseGapJaString()
fmt.Println(lunarChineseGapja)
```

[Result]
```go
2022-06-12
임인년 정미월 갑자일
壬寅年 丁未月 甲子日
```

2. 음력 2022년 6월 12일을 양력으로 변환
```go
// param : year(년), month(월), day(일), intercalation(윤달여부)
calendar.setLuarDate(2022, 6, 12, false);

// Solar Date (ISO format)
solar := GetSolarIsoFormat()
fmt.Println(solar)
```

[Result]
```go
2022-07-10
```

**CREDIT**

Thanks for the original Java code here which is now converted in Go:

[KoreanLunarCalendar](https://github.com/usingsky/KoreanLunarCalendar)

