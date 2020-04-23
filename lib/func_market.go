package lib

import "time"

func F한국증시_정규시장_거래시간임() bool {
	return f한국증시_거래시간_도우미(9, 0, 15, 30)
}

func F한국증시_정규경쟁대량매매_거래시간임() bool {
	return f한국증시_거래시간_도우미(9, 0, 15, 00)
}

func F한국증시_동시호가_시간임() bool {
	return f한국증시_거래시간_도우미(8, 30, 9, 0) ||
		f한국증시_거래시간_도우미(15, 20, 15, 30)
}

func F한국증시_시간외_종가매매_시간임() bool {
	return f한국증시_거래시간_도우미(8, 30, 8, 40) ||
		f한국증시_거래시간_도우미(15, 40, 16, 0)
}

func F한국증시_시간외_단일가매매_시간임() bool {
	return f한국증시_거래시간_도우미(16, 0, 18, 0)
}

func F한국증시_시간외_대량바스켓매매_거래시간임() bool {
	return f한국증시_거래시간_도우미(15, 40, 18, 0)
}

func f한국증시_거래시간_도우미(시작_시간, 시작_분, 종료_시간, 종료_분 int) bool {
	값 := F금일()
	지금 := time.Now()
	로케일 := 지금.Location()

	시작_시각 := time.Date(값.Year(), 값.Month(), 값.Day(), 시작_시간, 시작_분, 0, 0, 로케일)
	종료_시각 := time.Date(값.Year(), 값.Month(), 값.Day(), 종료_시간, 종료_분, 0, 0, 로케일)

	if 지금.After(시작_시각) && 지금.Before(종료_시각) {
		return true
	}

	return false
}
