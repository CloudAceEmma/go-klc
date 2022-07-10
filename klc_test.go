package klc

import (
	"fmt"
	"testing"
)

func TestLunarIsoFormat(t *testing.T) {
	SetSolarDate(2022, 7, 10)
	lunar := GetLunarIsoFormat()
	want := "2022-06-12"

	fmt.Println(lunar)

	if lunar != want {
		t.Errorf("got %q want %q", lunar, want)
	}
}

func TestGapjaString(t *testing.T) {
	SetSolarDate(2022, 7, 10)
	lunarGapja := GetGapjaString()
	want := "임인년 정미월 갑자일"

	fmt.Println(lunarGapja)

	if lunarGapja != want {
		t.Errorf("got %q want %q", lunarGapja, want)
	}
}

func TestChineseGapjaString(t *testing.T) {
	SetSolarDate(2022, 7, 10)
	lunarChineseGapja := GetChineseGapJaString()
	want := "壬寅年 丁未月 甲子日"

	fmt.Println(lunarChineseGapja)

	if lunarChineseGapja != want {
		t.Errorf("got %q want %q", lunarChineseGapja, want)
	}
}

func TestSolarIsoFormat(t *testing.T) {
	SetLunarDate(2022, 6, 12, false)
	solar := GetSolarIsoFormat()
	want := "2022-07-10"

	fmt.Println(solar)

	if solar != want {
		t.Errorf("got %q want %q", solar, want)
	}
}
