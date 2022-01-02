package krx

import (
	"github.com/ghts/ghts/lib"
	"time"
)

func F임시_지연_시간() time.Duration {
	if 지금 := time.Now(); 지금.Year() == 2022 && 지금.Month() == time.January && 지금.Day() == 3 {
		return time.Hour // 2022년 1월 3일 개장 1시간 순연.
	} else {
		return 0
	}
}

func F한국증시_정규_거래_시간임() bool {
	return f한국증시_거래시간_도우미(9, 0, 15, 20)
}

func F한국증시_동시호가_시간임() bool {
	return F한국증시_장전_동시호가_시간임() || F한국증시_장후_동시호가_시간임()
}

func F한국증시_장전_동시호가_시간임() bool {
	return f한국증시_거래시간_도우미(8, 40, 9, 0)
}

func F한국증시_장후_동시호가_시간임() bool {
	return f한국증시_거래시간_도우미(15, 20, 15, 30)
}

func F한국증시_장전_시간외_종가매매_시간임() bool {
	return f한국증시_거래시간_도우미(8, 30, 8, 40)
}

func F한국증시_장후_시간외_종가매매_시간임() bool {
	return f한국증시_거래시간_도우미(15, 40, 16, 0)
}

func F한국증시_시간외_단일가매매_시간임() bool {
	return f한국증시_거래시간_도우미(16, 0, 18, 0)
}

func F한국증시_정규경쟁대량매매_거래시간임() bool {
	return f한국증시_거래시간_도우미(9, 0, 15, 00)
}

func F한국증시_시간외_대량바스켓매매_거래시간임() bool {
	return f한국증시_거래시간_도우미(15, 40, 18, 0)
}

func F한국증시_ETF_LP_의무_호가제출_시간임() bool {
	// 규정상 9시 5분부터 LP 호가 제출이 시작되지만,
	// 9시 10분이 지나서도 NAV에서 벗어나 심하게 출렁이는 사례가 있으므로, 안전하게 9시 20분부터 시작하는 것으로 설정함.

	return f한국증시_거래시간_도우미(9, 20, 15, 20)
}

func F한국증시_폐장_시간임() bool {
	return !F한국증시_정규_거래_시간임() &&
		!F한국증시_동시호가_시간임() &&
		!F한국증시_장전_시간외_종가매매_시간임() &&
		!F한국증시_장후_시간외_종가매매_시간임() &&
		!F한국증시_시간외_단일가매매_시간임()
}

func f한국증시_거래시간_도우미(시작_시간, 시작_분, 종료_시간, 종료_분 int) bool {
	값 := lib.F금일()
	지금 := time.Now()
	로케일 := 지금.Location()

	시작_시각 := time.Date(값.Year(), 값.Month(), 값.Day(), 시작_시간, 시작_분, 0, 0, 로케일)
	종료_시각 := time.Date(값.Year(), 값.Month(), 값.Day(), 종료_시간, 종료_분, 0, 0, 로케일)

	if int64(F임시_지연_시간()) != 0 {
		시작_시각 = 시작_시각.Add(F임시_지연_시간())
		종료_시각 = 종료_시각.Add(F임시_지연_시간())
	}

	if 지금.After(시작_시각) && 지금.Before(종료_시각) {
		return true
	}

	return false
}

func F대기_한국_시각(시, 분, 초 int) {
	목표_시각 := lib.F2금일_한국_시각_단순형(시, 분, 초).Add(F임시_지연_시간())
	지금 := lib.F지금()

	if 목표_시각.After(지금) {
		대기_시간 := 목표_시각.Sub(지금)
		lib.F대기(대기_시간)
	}
}

func F금일_보정_시각(시, 분, 초 int) time.Time {
	return lib.F금일().Add(F임시_지연_시간() + time.Duration(시)*lib.P1시간 + time.Duration(분)*lib.P1분 + time.Duration(초)*lib.P1초)
}
