package trade

import (
	"github.com/ghts/ghts/lib"

	"math"
	"strconv"
	"strings"
)

func F이동_범위_최대값(값_모음 []float64, 윈도우_크기 int) []float64 {
	고가_모음 := make([]float64, len(값_모음))
	고가_모음[0] = 값_모음[0]

	for i := 1; i < 윈도우_크기; i++ {
		고가_모음[i] = lib.F최대값(값_모음[:i])
	}

	for i := 윈도우_크기; i < len(값_모음); i++ {
		고가_모음[i] = lib.F최대값(값_모음[i-윈도우_크기+1 : i])
	}

	return 고가_모음
}

func F이동_범위_최소값(값_모음 []float64, 윈도우_크기 int) []float64 {
	저가_모음 := make([]float64, len(값_모음))
	저가_모음[0] = 값_모음[0]

	for i := 1; i < 윈도우_크기; i++ {
		저가_모음[i] = lib.F최소값(값_모음[:i])
	}

	for i := 윈도우_크기; i < len(값_모음); i++ {
		저가_모음[i] = lib.F최소값(값_모음[i-윈도우_크기+1 : i])
	}

	return 저가_모음
}

func F단순_이동_평균(값_모음 []float64, 윈도우_크기 int) []float64 {
	윈도우_크기 = lib.F최소값_정수(윈도우_크기, len(값_모음))

	이동_평균_모음 := make([]float64, len(값_모음))
	윈도우_크기_실수값 := float64(윈도우_크기)
	합계 := 0.0

	for i := 0; i < 윈도우_크기; i++ {
		합계 += 값_모음[i]
		이동_평균_모음[i] = 합계 / float64(i+1)
	}

	for i := 윈도우_크기; i < len(값_모음); i++ {
		합계 += 값_모음[i] - 값_모음[i-윈도우_크기]
		이동_평균_모음[i] = 합계 / 윈도우_크기_실수값
	}

	return 이동_평균_모음
}

func F지수_이동_평균(값_모음 []float64, 윈도우_크기 int) []float64 {
	윈도우_크기 = lib.F최소값_정수(윈도우_크기, len(값_모음))

	이동_평균_모음 := make([]float64, len(값_모음))

	합계 := 0.0

	for i := 0; i < 윈도우_크기; i++ {
		합계 += 값_모음[i]
		이동_평균_모음[i] = 합계 / float64(i+1)
	}

	승수 := 2.0 / float64(1+윈도우_크기)
	승수_나머지 := 1.0 - 승수

	for i := 윈도우_크기; i < len(값_모음); i++ {
		이동_평균_모음[i] = 값_모음[i]*승수 + 이동_평균_모음[i-1]*승수_나머지
	}

	return 이동_평균_모음
}

func F가중_이동_평균(값_모음, 가중치 []float64, 윈도우_크기 int) []float64 {
	윈도우_크기 = lib.F최소값_정수(윈도우_크기, len(가중치))
	가중치_합계 := lib.F합계(가중치[:윈도우_크기-1])
	가중_이동_평균 := make([]float64, len(값_모음))

	for i := 윈도우_크기; i < len(값_모음); i++ {
		가중치_합계 += 가중치[i] - 가중치[i-윈도우_크기]

		for j := i - 윈도우_크기 + 1; j <= i; j++ {
			가중_이동_평균[i] += 값_모음[j] * 가중치[j] / 가중치_합계
		}
	}

	return 가중_이동_평균
}

func F이동_Z점수(값_모음 []float64, 윈도우_크기 int) []float64 {
	return f이동_Z점수_도우미(값_모음, 윈도우_크기, false)
}

func F지수_이동_Z점수(값_모음 []float64, 윈도우_크기 int) []float64 {
	return f이동_Z점수_도우미(값_모음, 윈도우_크기, true)
}

func F볼린저_밴드(값_모음 []float64, 윈도우_크기 int, 표준편차_배율 float64) []float64 {
	return f볼린저_밴드_도우미(값_모음, 윈도우_크기, 표준편차_배율, false)
}

func F지수_볼린저_밴드(값_모음 []float64, 윈도우_크기 int, 표준편차_배율 float64) []float64 {
	return f볼린저_밴드_도우미(값_모음, 윈도우_크기, 표준편차_배율, true)
}

func F볼린저_밴드_폭(값_모음 []float64, 윈도우_크기 int, 표준편차_배율 float64) []float64 {
	return f볼린저_밴드_폭_도우미(값_모음, 윈도우_크기, 표준편차_배율, false)
}

func F지수_볼린저_밴드_폭(값_모음 []float64, 윈도우_크기 int, 표준편차_배율 float64) []float64 {
	return f볼린저_밴드_폭_도우미(값_모음, 윈도우_크기, 표준편차_배율, true)
}

func F이동_표준_편차(값_모음 []float64, 윈도우_크기 int) []float64 {
	return F이동_표준_편차_도우미(값_모음, 윈도우_크기, false)
}

func F지수_이동_표준_편차(값_모음 []float64, 윈도우_크기 int) []float64 {
	return F이동_표준_편차_도우미(값_모음, 윈도우_크기, true)
}

func F이동_평균_도우미(값_모음 []float64, 윈도우_크기 int, EMA bool) []float64 {
	if EMA {
		return F지수_이동_평균(값_모음, 윈도우_크기)
	} else {
		return F단순_이동_평균(값_모음, 윈도우_크기)
	}
}

func F이동_표준_편차_도우미(값_모음 []float64, 윈도우_크기 int, EMA bool) []float64 {
	이동_평균_모음 := F이동_평균_도우미(값_모음, 윈도우_크기, EMA)
	편차_제곱_모음 := make([]float64, len(값_모음))
	표준_편차_모음 := make([]float64, len(값_모음))
	윈도우_크기_실수값 := float64(윈도우_크기)
	합계 := 0.0

	for i := 0; i < len(값_모음); i++ {
		편차_제곱_모음[i] = math.Pow(이동_평균_모음[i]-값_모음[i], 2)
	}

	for i := 0; i < 윈도우_크기; i++ {
		합계 += 편차_제곱_모음[i]
		표준_편차_모음[i] = math.Sqrt(합계 / float64(i+1))
	}

	for i := 윈도우_크기; i < len(값_모음); i++ {
		합계 = 합계 + 편차_제곱_모음[i] - 편차_제곱_모음[i-윈도우_크기]
		표준_편차_모음[i] = math.Sqrt(합계 / 윈도우_크기_실수값)
	}

	return 표준_편차_모음
}

func f이동_Z점수_도우미(값_모음 []float64, 윈도우_크기 int, EMA bool) []float64 {
	이동_평균 := F이동_평균_도우미(값_모음, 윈도우_크기, EMA)
	표준_편차 := F이동_표준_편차_도우미(값_모음, 윈도우_크기, EMA)
	z점수 := make([]float64, len(값_모음))

	for i := 0; i < len(값_모음); i++ {
		if 표준_편차[i] == 0 && 값_모음[i]-이동_평균[i] >= 0 {
			z점수[i] = math.Inf(1)
		} else if 표준_편차[i] == 0 && 값_모음[i]-이동_평균[i] < 0 {
			z점수[i] = math.Inf(-1)
		} else {
			z점수[i] = (값_모음[i] - 이동_평균[i]) / 표준_편차[i]
		}
	}

	return z점수
}

func f볼린저_밴드_도우미(값_모음 []float64, 윈도우_크기 int, 표준편차_배율 float64, EMA bool) []float64 {
	이동_평균 := F이동_평균_도우미(값_모음, 윈도우_크기, EMA)
	표준_편차 := F이동_표준_편차_도우미(값_모음, 윈도우_크기, EMA)
	볼린저_밴드 := make([]float64, len(값_모음))

	for i := 0; i < len(값_모음); i++ {
		볼린저_밴드[i] = 이동_평균[i] + 표준편차_배율*표준_편차[i]
	}

	return 볼린저_밴드
}

func f볼린저_밴드_폭_도우미(값_모음 []float64, 윈도우_크기 int, 표준편차_배율 float64, EMA bool) []float64 {
	표준편차_배율 = math.Abs(표준편차_배율)
	이동_평균 := F이동_평균_도우미(값_모음, 윈도우_크기, EMA)
	표준_편차 := F이동_표준_편차_도우미(값_모음, 윈도우_크기, EMA)
	밴드_폭 := make([]float64, len(값_모음))

	for i := 0; i < len(값_모음); i++ {
		밴드_폭[i] = 표준편차_배율 * 표준_편차[i] * 2 / 이동_평균[i]
	}

	return 밴드_폭
}

func F모의_매수_거래가(기준값, 슬리피지_비용 float64) float64 {
	거래가_후보 := int64(math.Ceil(기준값))

	for {
		if 거래가_후보%5 == 0 {
			return float64(거래가_후보) + 슬리피지_비용
		}

		거래가_후보++
	}
}

func F모의_매도_거래가(기준값, 슬리피지_비용 float64) float64 {
	거래가_후보 := int64(math.Floor(기준값))

	for {
		if 거래가_후보%5 == 0 {
			return float64(거래가_후보) - 슬리피지_비용
		}

		거래가_후보--
	}
}

func F상향_돌파(전일_고가, 당일_고가, 기준가 float64) bool {
	return 전일_고가 < 기준가 && 당일_고가 > 기준가
}

func F하향_돌파(전일_저가, 당일_저가, 기준가 float64) bool {
	return 전일_저가 > 기준가 && 당일_저가 < 기준가
}

func F기하_수익율(수익율_모음 []float64) float64 {
	기하_수익율 := 1.0

	for _, 수익율 := range 수익율_모음 {
		기하_수익율 *= 1 + (수익율 / 100)
	}

	return math.Round((기하_수익율*100-100)*1000) / 1000
}

func F기간(시작일, 종료일 uint32) int {
	시작, 에러 := lib.F2포맷된_일자("20060102", strconv.Itoa(int(시작일)))
	if 에러 != nil {
		return 0
	}

	종료, 에러 := lib.F2포맷된_일자("20060102", strconv.Itoa(int(종료일)))
	if 에러 != nil {
		return 0
	}

	return int(종료.Sub(시작).Hours() / 24)
}

func F주문가by퍼센트(매도_매수_구분 lib.T매도_매수_구분, 현재가 int64, 퍼센트 float64) int64 {
	switch 매도_매수_구분 {
	case lib.P매도:
		return F매도_주문가by퍼센트(현재가, 퍼센트)
	case lib.P매수:
		return F매수_주문가by퍼센트(현재가, 퍼센트)
	default:
		panic(lib.New에러("예상하지 못한 ''매도_매수_구분'값 : %v", 매도_매수_구분))
	}
}

func F매도_주문가by퍼센트(현재가 int64, 퍼센트 float64) int64 {
	기준가 := float64(현재가) * (1 + math.Abs(퍼센트)/100)
	주문가_후보 := int64(math.Ceil(기준가))

	for {
		if 주문가_후보%5 == 0 {
			return 주문가_후보
		} else {
			주문가_후보++
		}
	}
}

func F매수_주문가by퍼센트(현재가 int64, 퍼센트 float64) int64 {
	기준가 := float64(현재가) * (1 - math.Abs(퍼센트)/100)
	주문가_후보 := int64(math.Ceil(기준가))

	for {
		if 주문가_후보%5 == 0 {
			return 주문가_후보
		} else {
			주문가_후보--
		}
	}
}

func F종목코드_보정(종목코드 string) string {
	if len(종목코드) == 7 && strings.HasPrefix(종목코드, "A") {
		return 종목코드[1:]
	} else if len(종목코드) == 7 && strings.HasPrefix(종목코드, "Q") {
		return 종목코드[1:]
	}

	return 종목코드
}
