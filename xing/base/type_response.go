package xt

import (
	"errors"
	lb "github.com/ghts/ghts/lib"
	"strings"
)

func NewJSON응답(값 interface{}) *JSON응답 {
	switch 값.(type) {
	case error:
		return &JSON응답{V: nil, E: 값.(error).Error()}
	default:
		return &JSON응답{V: 값, E: ""}
	}
}

type JSON응답 struct {
	V interface{}
	E string
}

func (s JSON응답) Error() error {
	if strings.TrimSpace(s.E) == "" {
		return nil
	} else {
		return errors.New(s.E)
	}
}

type I이중_응답 interface {
	I이중_응답1
	I이중_응답2
}

type I이중_응답1 interface {
	G응답1() I이중_응답1
}

type I이중_응답2 interface {
	G응답2() I이중_응답2
}

type S이중_응답_일반형 struct {
	M응답1 I이중_응답1
	M응답2 I이중_응답2
}

func (s *S이중_응답_일반형) G응답1() I이중_응답1 { return s.M응답1 }
func (s *S이중_응답_일반형) G응답2() I이중_응답2 { return s.M응답2 }

func (s *S이중_응답_일반형) G값(TR코드 string) interface{} {
	switch TR코드 {
	case TR현물_정상_주문_CSPAT00600:
		g := new(CSPAT00600_현물_정상_주문_응답)
		g.M응답1 = s.M응답1.(*CSPAT00600_현물_정상_주문_응답1)
		g.M응답2 = s.M응답2.(*CSPAT00600_현물_정상_주문_응답2)
		return g
	case TR현물_정정_주문_CSPAT00700:
		g := new(CSPAT00700_현물_정정_주문_응답)
		g.M응답1 = s.M응답1.(*CSPAT00700_현물_정정_주문_응답1)
		g.M응답2 = s.M응답2.(*CSPAT00700_현물_정정_주문_응답2)
		return g
	case TR현물_취소_주문_CSPAT00800:
		g := new(CSPAT00800_현물_취소_주문_응답)
		g.M응답1 = s.M응답1.(*CSPAT00800_현물_취소_주문_응답1)
		g.M응답2 = s.M응답2.(*CSPAT00800_S현물_취소_주문_응답2)
		return g
	case TR기업정보_요약_t3320:
		g := new(T3320_기업정보_요약_응답)
		g.M응답1 = s.M응답1.(*T3320_기업정보_요약_응답1)
		g.M응답2 = s.M응답2.(*T3320_기업정보_요약_응답2)
		return g
	default:
		panic(lb.New에러("예상하지 못한 TR코드 : '%v'", TR코드))
	}
}

type I헤더_반복값_TR데이터 interface {
	I헤더_TR데이터
	I반복값_모음_TR데이터
}

type I헤더_TR데이터 interface {
	G헤더_TR데이터() I헤더_TR데이터
}

type I반복값_모음_TR데이터 interface {
	G반복값_모음_TR데이터() I반복값_모음_TR데이터
}

type S헤더_반복값 struct {
	M헤더     I헤더_TR데이터
	M반복값_모음 I반복값_모음_TR데이터
}

func (s *S헤더_반복값) G헤더_TR데이터() I헤더_TR데이터 { return s.M헤더 }
func (s *S헤더_반복값) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

func (s *S헤더_반복값) G값(TR코드 string) interface{} {
	switch TR코드 {
	default:
		panic(lb.New에러with출력("예상하지 못한 TR코드 : '%v' %v", TR코드, lb.F소스코드_위치(0))) // 패닉 출력을 삭제하지 말 것.
	//case TR현물_시간대별_체결_조회:
	//	값 := new(T1301_현물_시간대별_체결_응답)
	//	값.M헤더 = s.M헤더.(*T1301_현물_시간대별_체결_응답_헤더)
	//	값.M반복값_모음 = s.M반복값_모음.(*T1301_현물_시간대별_체결_응답_반복값_모음)
	//	return 값
	case TR현물_기간별_조회_t1305:
		값 := new(T1305_현물_기간별_조회_응답)
		값.M헤더 = s.M헤더.(*T1305_현물_기간별_조회_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*T1305_현물_기간별_조회_응답_반복값_모음)
		return 값
	case TR현물_당일_전일_분틱_조회_t1310:
		값 := new(T1310_현물_전일당일분틱조회_응답)
		값.M헤더 = s.M헤더.(*T1310_현물_전일당일분틱조회_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*T1310_현물_전일당일분틱조회_응답_반복값_모음)
		return 값
	case TR관리_불성실_투자유의_조회_t1404:
		값 := new(T1404_관리종목_조회_응답)
		값.M헤더 = s.M헤더.(*T1404_관리종목_조회_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*T1404_관리종목_조회_응답_반복값_모음)
		return 값
	case TR투자경고_매매정지_정리매매_조회_t1405:
		값 := new(T1405_투자경고_조회_응답)
		값.M헤더 = s.M헤더.(*T1405_투자경고_조회_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*T1405_투자경고_조회_응답_반복값_모음)
		return 값
	case TR_ETF_시간별_추이_t1902:
		값 := new(T1902_ETF시간별_추이_응답)
		값.M헤더 = s.M헤더.(*T1902_ETF시간별_추이_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*T1902_ETF시간별_추이_응답_반복값_모음)
		return 값
	case TR재무순위_종합_t3341:
		값 := new(T3341_재무순위_응답)
		값.M헤더 = s.M헤더.(*T3341_재무순위_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*T3341_재무순위_응답_반복값_모음)
		return 값
	case TR현물_차트_일주월년_t8410:
		값 := new(T8410_현물_차트_일주월년_응답)
		값.M헤더 = s.M헤더.(*T8410_현물_차트_일주월년_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*T8410_현물_차트_일주월년_응답_반복값_모음)
		return 값
	case TR현물_차트_틱_t8411:
		값 := new(T8411_현물_차트_틱_응답)
		값.M헤더 = s.M헤더.(*T8411_현물_차트_틱_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*T8411_현물_차트_틱_응답_반복값_모음)
		return 값
	case TR현물_차트_분_t8412:
		값 := new(T8412_현물_차트_분_응답)
		값.M헤더 = s.M헤더.(*T8412_현물_차트_분_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*T8412_현물_차트_분_응답_반복값_모음)
		return 값
	case TR현물_차트_일주월_t8413:
		값 := new(T8413_현물_차트_일주월_응답)
		값.M헤더 = s.M헤더.(*T8413_현물_차트_일주월_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*T8413_현물_차트_일주월_응답_반복값_모음)
		return 값
	case TR증시_주변_자금_추이_t8428:
		값 := new(T8428_증시주변_자금추이_응답)
		값.M헤더 = s.M헤더.(*T8428_증시주변_자금추이_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*T8428_증시주변_자금추이_응답_반복값_모음)
		return 값
	}
}
