package xt

import (
	lb "github.com/ghts/ghts/lib"
	"strings"
	"testing"
)

func TestF바이트_변환값_해석(t *testing.T) {
	변환_형식_모음 := []lb.T변환{lb.JSON, lb.GOB}

	원본값_모음 := []interface{}{
		new(S현물_주문_응답_실시간_정보),
		new(CSPAT00600_현물_정상_주문_질의값), new(CSPAT00700_현물_정정_주문_질의값), new(T1310_현물_전일당일분틱조회_질의값),
		new(T1305_현물_기간별_조회_질의값), new(lb.S질의값_단일종목_연속키), new(T8428_증시주변_자금추이_질의값),
		new(T1101_현물_호가_조회_응답), new(T1102_현물_시세_조회_응답), new(T1901_ETF_시세_조회_응답),
		new(T1301_현물_시간대별_체결_응답), new(T1301_현물_시간대별_체결_응답_헤더), new(T1301_현물_시간대별_체결_응답_반복값),
		new(T1301_현물_시간대별_체결_응답_반복값_모음),
		new(T1305_현물_기간별_조회_응답), new(T1305_현물_기간별_조회_응답_헤더), new(T1305_현물_기간별_조회_응답_반복값),
		new(T1305_현물_기간별_조회_응답_반복값_모음),
		new(T1310_현물_전일당일분틱조회_응답), new(T1310_현물_전일당일분틱조회_응답_헤더), new(T1310_현물_전일당일분틱조회_응답_반복값),
		new(T1310_현물_전일당일분틱조회_응답_반복값_모음),
		new(T1901_ETF_시세_조회_응답),
		new(T1902_ETF시간별_추이_응답), new(T1902_ETF시간별_추이_응답_헤더), new(T1902_ETF시간별_추이_응답_반복값),
		new(T1902_ETF시간별_추이_응답_반복값_모음), new(T1906_ETF_LP_호가_조회_응답),
		new(T8411_현물_차트_틱_질의값), new(T8411_현물_차트_틱_응답_헤더), new(T8411_현물_차트_틱_응답_반복값),
		new(T8428_증시주변_자금추이_응답), new(T8428_증시주변_자금추이_응답_헤더), new(T8428_증시주변_자금추이_응답_반복값),
		new(T8428_증시주변_자금추이_응답_반복값_모음),
		new(T8436_현물_종목조회_응답_반복값), new(T8436_현물_종목조회_응답)}

	for _, 변환_형식 := range 변환_형식_모음 {
		for _, 원본값 := range 원본값_모음 {
			매개체, 에러 := lb.New바이트_변환(변환_형식, 원본값)
			lb.F테스트_에러없음(t, 에러)

			해석값, 에러 := 매개체.S해석기(F바이트_변환값_해석).G해석값()
			lb.F테스트_에러없음(t, 에러)

			자료형_문자열 := strings.Replace(f자료형_문자열(해석값), "*", "", -1)
			lb.F테스트_같음(t, 자료형_문자열, 매개체.G자료형_문자열())
		}
	}
}
