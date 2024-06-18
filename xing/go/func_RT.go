package xing

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
)

func F실시간_정보_구독_단순TR(RT코드 string) (에러 error) {
	return F질의(lib.New질의값_기본형(xt.TR실시간_정보_구독, RT코드)).G에러()
}

func F실시간_정보_해지_단순TR(RT코드 string) (에러 error) {
	return F질의(lib.New질의값_기본형(xt.TR실시간_정보_해지, RT코드)).G에러()
}

func F실시간_정보_구독_단일_종목(RT코드 string, 종목코드 string) (에러 error) {
	return F질의(lib.New질의값_단일_종목(xt.TR실시간_정보_구독, RT코드, 종목코드)).G에러()
}

func F실시간_정보_해지_단일_종목(RT코드 string, 종목코드 string) (에러 error) {
	return F질의(lib.New질의값_단일_종목(xt.TR실시간_정보_해지, RT코드, 종목코드)).G에러()
}

func F실시간_정보_구독_복수_종목(RT코드 string, 종목코드_모음 []string) (에러 error) {
	return F질의(lib.New질의값_복수_종목(xt.TR실시간_정보_구독, RT코드, 종목코드_모음)).G에러()
}

func F실시간_정보_해지_복수_종목(RT코드 string, 종목코드_모음 []string) (에러 error) {
	return F질의(lib.New질의값_복수_종목(xt.TR실시간_정보_해지, RT코드, 종목코드_모음)).G에러()
}

func F실시간_정보_일괄_해지() (에러 error) {
	return F질의(lib.New질의값_기본형(xt.TR실시간_정보_일괄_해지, "")).G에러()
}

func F실시간_데이터_구독_ETF(종목코드 string, 종목코드_모음 ...string) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	RT코드_모음 := []string{
		xt.RT코스피_호가_잔량_H1,
		xt.RT코스피_체결_S3,
		xt.RT코스피_ETF_NAV_I5,
		xt.RT코스피_시간외_호가_잔량_H2,
		xt.RT코스피_예상_체결_YS3}

	종목코드_모음 = append([]string{종목코드}, 종목코드_모음...)

	for _, RT코드 := range RT코드_모음 {
		lib.F확인1(F실시간_정보_구독_복수_종목(RT코드, 종목코드_모음))
	}

	return nil
}

func F실시간_데이터_해지_ETF(종목코드_모음 []string) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	RT코드_모음 := []string{
		xt.RT코스피_호가_잔량_H1,
		xt.RT코스피_체결_S3,
		xt.RT코스피_ETF_NAV_I5,
		xt.RT코스피_시간외_호가_잔량_H2,
		xt.RT코스피_예상_체결_YS3}

	for _, RT코드 := range RT코드_모음 {
		lib.F확인1(F실시간_정보_해지_복수_종목(RT코드, 종목코드_모음))
	}

	return nil
}

func F주문_응답_실시간_정보_구독() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	if 주문_응답_구독_중.G값() {
		return
	}

	lib.F확인1(F실시간_정보_구독_단순TR(xt.RT현물_주문_접수_SC0))
	lib.F확인1(F실시간_정보_구독_단순TR(xt.RT현물_주문_체결_SC1))
	lib.F확인1(F실시간_정보_구독_단순TR(xt.RT현물_주문_정정_SC2))
	lib.F확인1(F실시간_정보_구독_단순TR(xt.RT현물_주문_취소_SC3))
	lib.F확인1(F실시간_정보_구독_단순TR(xt.RT현물_주문_거부_SC4))

	return nil
}

func F주문_응답_실시간_정보_해지() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	if !주문_응답_구독_중.G값() {
		return
	}

	defer 주문_응답_구독_중.S값(false)

	lib.F확인1(F실시간_정보_해지_단순TR(xt.RT현물_주문_접수_SC0))
	lib.F확인1(F실시간_정보_해지_단순TR(xt.RT현물_주문_체결_SC1))
	lib.F확인1(F실시간_정보_해지_단순TR(xt.RT현물_주문_정정_SC2))
	lib.F확인1(F실시간_정보_해지_단순TR(xt.RT현물_주문_취소_SC3))
	lib.F확인1(F실시간_정보_해지_단순TR(xt.RT현물_주문_거부_SC4))

	return nil
}

func F호가_잔량_실시간_정보_구독(종목코드 string) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	종목 := lib.F확인2(F종목by코드(종목코드))

	switch 종목.G시장구분() {
	case lib.P시장구분_코스피, lib.P시장구분_ETF, lib.P시장구분_ETN:
		return F실시간_정보_구독_단일_종목(xt.RT코스피_호가_잔량_H1, 종목코드)
	case lib.P시장구분_코스닥:
		return F실시간_정보_구독_단일_종목(xt.RT코스닥_호가_잔량_HA, 종목코드)
	default:
		return lib.New에러("미구현 시장 구분 : '%v' '%v'", 종목코드, 종목.G시장구분())
	}
}

func F호가_잔량_실시간_정보_해지(종목코드 string) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	종목 := lib.F확인2(F종목by코드(종목코드))

	switch 종목.G시장구분() {
	case lib.P시장구분_코스피, lib.P시장구분_ETF, lib.P시장구분_ETN:
		return F실시간_정보_해지_단일_종목(xt.RT코스피_호가_잔량_H1, 종목코드)
	case lib.P시장구분_코스닥:
		return F실시간_정보_해지_단일_종목(xt.RT코스닥_호가_잔량_HA, 종목코드)
	default:
		return lib.New에러("미구현 시장 구분 : '%v' '%v'", 종목코드, 종목.G시장구분())
	}
}

func F체결_실시간_정보_구독(종목코드 string) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	종목 := lib.F확인2(F종목by코드(종목코드))

	switch 종목.G시장구분() {
	case lib.P시장구분_코스피, lib.P시장구분_ETF, lib.P시장구분_ETN:
		return F실시간_정보_구독_단일_종목(xt.RT코스피_체결_S3, 종목코드)
	case lib.P시장구분_코스닥:
		return F실시간_정보_구독_단일_종목(xt.RT코스닥_체결_K3, 종목코드)
	default:
		return lib.New에러("미구현 시장 구분 : '%v' '%v'", 종목코드, 종목.G시장구분())
	}
}

func F체결_실시간_정보_해지(종목코드 string) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	종목 := lib.F확인2(F종목by코드(종목코드))

	switch 종목.G시장구분() {
	case lib.P시장구분_코스피, lib.P시장구분_ETF, lib.P시장구분_ETN:
		return F실시간_정보_해지_단일_종목(xt.RT코스피_체결_S3, 종목코드)
	case lib.P시장구분_코스닥:
		return F실시간_정보_해지_단일_종목(xt.RT코스닥_체결_K3, 종목코드)
	default:
		return lib.New에러("미구현 시장 구분 : '%v' '%v'", 종목코드, 종목.G시장구분())
	}
}
