package xing

//import (
//	lb "github.com/ghts/ghts/lib"
//	"strings"
//
//	"testing"
//)
//
//func TestT8432_지수선물_마스터_조회(t *testing.T) {
//	t.Parallel()
//
//	// 'V' : 변동성 지수 선물, 'S' : 섹터 지수 선물, 그 외 코스피200 지수 선물
//	값_모음, 에러 := TrT8432_지수선물_마스터_조회("K")
//	lb.F테스트_에러없음(t, 에러)
//
//	for _, 값 := range 값_모음 {
//		//lb.F체크포인트(값.M종목명, 값.M상한가, 값.M하한가, 값.M전일종가, 값.M전일고가, 값.M전일저가, 값.M기준가)
//
//		lb.F테스트_참임(t, strings.Contains(값.M확장코드, 값.M종목코드), 값.M종목코드, 값.M확장코드)
//
//		// (-) 값이 나올 수 있다.
//		//lb.F테스트_참임(t, 값.M상한가 > 0, 값.M상한가)
//		//lb.F테스트_참임(t, 값.M하한가 > 0, 값.M하한가)
//		//lb.F테스트_참임(t, 값.M전일종가 > 0, 값.M전일종가)
//		//lb.F테스트_참임(t, 값.M전일고가 >= 0, 값.M전일고가)
//		//lb.F테스트_참임(t, 값.M전일저가 >= 0, 값.M전일저가)
//		//lb.F테스트_참임(t, 값.M기준가 > 0, 값.M기준가)
//
//		if 값.M전일고가 != 0 {
//			lb.F테스트_참임(t, 값.M전일고가 <= 값.M상한가)
//			lb.F테스트_참임(t, 값.M전일고가 >= 값.M하한가)
//			lb.F테스트_참임(t, 값.M전일고가 >= 값.M전일저가)
//
//			if 값.M전일종가 != 0 {
//				lb.F테스트_참임(t, 값.M전일고가 >= 값.M전일종가)
//			}
//		}
//
//		if 값.M전일저가 != 0 {
//			lb.F테스트_참임(t, 값.M전일저가 <= 값.M상한가)
//			lb.F테스트_참임(t, 값.M전일저가 >= 값.M하한가)
//
//			if 값.M전일종가 != 0 {
//				lb.F테스트_참임(t, 값.M전일저가 <= 값.M전일종가, 값.M전일저가, 값.M전일종가)
//			}
//		}
//	}
//}
