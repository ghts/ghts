/* Copyright (C) 2015-2020 김운하 (unha.kim@ghts.org)

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGPL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGPL 2.1판을 참고하시기 바랍니다.
GNU LGPL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2020년 UnHa Kim (unha.kim@ghts.org)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package xing

//import (
//	"github.com/ghts/ghts/lib"
//	xt "github.com/ghts/ghts/xing/base"
//	"testing"
//)
//
//func TestCFOFQ02400_선물옵션_미결제약정_질의값(t *testing.T) {
//	t.Parallel()
//
//	_, ok := interface{}(new(xt.CFOFQ02400_선물옵션_미결제약정_질의값)).(lib.I질의값)
//
//	lib.F테스트_참임(t, ok)
//}
//
//func TestCFOFQ02400_선물옵션_미결제약정(t *testing.T) {
//	t.Parallel()
//
//	t.SkipNow()
//
//	//등록시장_`모음 := []xt.CFOFQ02400_등록시장{xt.CFOFQ02400_전체, xt.CFOFQ02400_KOSPI,
//	//	xt.CFOFQ02400_KOSDAQ, xt.CFOFQ02400_KSE, xt.CFOFQ02400_KOFEX}
//
//	계좌번호 := lib.F확인(F계좌_번호(1)).(string) // 선물옵션 계좌를 선택해야 함.
//	계좌_상세명, 에러 := F계좌_상세명(계좌번호)
//	lib.F테스트_에러없음(t, 에러)
//	lib.F테스트_같음(t, 계좌_상세명, "선물옵션")
//
//	등록시장 := xt.CFOFQ02400_전체 // 등록시장_모음[lib.F임의_범위_이내_정수값(0, 4)]
//	매수일자 := F당일()
//
//	응답값, 에러 := TrCFOFQ02400_선물옵션_미결제약정(계좌번호, 등록시장, 매수일자)
//	lib.F테스트_에러없음(t, 에러)
//
//	값1 := 응답값.M응답1
//	lib.F테스트_같음(t, 값1.M레코드수량, 1)
//	lib.F테스트_같음(t, 값1.M계좌번호, 계좌번호)
//	lib.F테스트_같음(t, 값1.M등록시장, xt.CFOFQ02400_전체, xt.CFOFQ02400_KOSPI,
//		xt.CFOFQ02400_KOSDAQ, xt.CFOFQ02400_KSE, xt.CFOFQ02400_KOFEX)
//	lib.F테스트_같음(t, 값1.M매수일자, 매수일자)
//
//	값2 := 응답값.M응답2
//	lib.F테스트_같음(t, 값2.M레코드수량, 1)
//	lib.F테스트_다름(t, 값2.M계좌명, "")
//	lib.F테스트_참임(t, 값2.M선물약정수량 >= 0)
//	lib.F테스트_참임(t, 값2.M옵션약정수량 >= 0)
//	lib.F테스트_참임(t, 값2.M약정수량 >= 0)
//	lib.F테스트_참임(t, 값2.M선물약정금액 >= 0)
//	lib.F테스트_참임(t, 값2.M선물매수약정금액 >= 0)
//	lib.F테스트_참임(t, 값2.M선물매도약정금액 >= 0)
//	lib.F테스트_참임(t, 값2.M콜옵션약정금액 >= 0)
//	lib.F테스트_참임(t, 값2.M콜매수금액 >= 0)
//	lib.F테스트_참임(t, 값2.M콜매도금액 >= 0)
//	lib.F테스트_참임(t, 값2.M풋옵션약정금액 >= 0)
//	lib.F테스트_참임(t, 값2.M풋매수금액 >= 0)
//	lib.F테스트_참임(t, 값2.M풋매도금액 >= 0)
//	lib.F테스트_참임(t, 값2.M전체약정금액 >= 0)
//	lib.F테스트_참임(t, 값2.M매수약정누계금액 >= 0)
//	lib.F테스트_참임(t, 값2.M매도약정누계금액 >= 0)
//	lib.F테스트_같음(t, 값2.M전체손익합계, 값2.M선물손익합계+값2.M옵션손익합계)
//
//	for _, 반복값1 := range 응답값.M반복값1_모음 {
//		lib.F테스트_같음(t, 반복값1.M선물옵션품목구분,
//			xt.P선옵품목_코스피200_관련, xt.P선옵품목_코스피200_제외, xt.P선옵품목_코스닥50_관련)
//		lib.F테스트_참임(t, 반복값1.M선물매도수량 >= 0)
//		//lib.F체크포인트(반복값1.M선물매도손익)
//		lib.F테스트_참임(t, 반복값1.M선물매수수량 >= 0)
//		//lib.F체크포인트(반복값1.M선물매수손익)
//		lib.F테스트_참임(t, 반복값1.M콜매도수량 >= 0)
//		//lib.F체크포인트(반복값1.M콜매도손익)
//		lib.F테스트_참임(t, 반복값1.M콜매수수량 >= 0)
//		//lib.F체크포인트(반복값1.M콜매수손익)
//		lib.F테스트_참임(t, 반복값1.M풋매도수량 >= 0)
//		//lib.F체크포인트(반복값1.M풋매도손익)
//		lib.F테스트_참임(t, 반복값1.M풋매수수량 >= 0)
//		//lib.F체크포인트(반복값1.M풋매수손익)
//	}
//
//	for _, 반복값2 := range 응답값.M반복값2_모음 {
//		lib.F테스트_다름(t, 반복값2.M종목코드, "")
//		lib.F테스트_다름(t, 반복값2.M종목명, "")
//		lib.F테스트_같음(t, 반복값2.M매도_매수_구분, lib.P매도_매수_전체, lib.P매도, lib.P매수)
//		lib.F테스트_참임(t, 반복값2.M잔고수량 >= 0)
//		lib.F테스트_참임(t, 반복값2.M평균가 >= 0)
//		lib.F테스트_참임(t, 반복값2.M당초금액 >= 0)
//		lib.F테스트_참임(t, 반복값2.M당일청산수량 >= 0)
//		lib.F테스트_참임(t, 반복값2.M현재가 >= 0)
//		lib.F테스트_참임(t, 반복값2.M평가금액 >= 0)
//		lib.F테스트_참임(t, float64(반복값2.M평가손익금액)*반복값2.M평가수익률 >= 0)
//	}
//}
