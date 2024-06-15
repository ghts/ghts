/* Copyright (C) 2015-2024 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2024년 UnHa Kim (unha.kim@ghts.org)

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

package util

import (
	dd "github.com/ghts/ghts/data/daily_price"
	"github.com/ghts/ghts/lib"
	mt "github.com/ghts/ghts/lib/market_time"
	xt "github.com/ghts/ghts/xing/base"
	xing "github.com/ghts/ghts/xing/go"

	"database/sql"
	"time"
)

func F당일_일일_가격정보_수집(db *sql.DB) (에러 error) {
	defer func() {
		lib.S예외처리{M에러: &에러}.S실행()

		if 에러 == nil {
			lib.F문자열_출력("** 당일 가격정보 수집 완료 **")
		}
	}()

	if lib.F지금().Before(mt.F금일_보정_시각(15, 31, 0)) && lib.F금일().Equal(xing.F당일()) {
		lib.F문자열_출력("폐장 이후 당일 일일 가격정보 수집 예정.")
		mt.F대기_한국_시각(15, 31, 00)
		lib.F문자열_출력("폐장 대기 완료.")
	}

	당일 := lib.F일자2정수(xing.F당일())
	한달전 := lib.F금일().Add(-30 * lib.P1일)
	종목코드_모음 := make([]string, 0)

	for _, 종목코드 := range xing.F종목코드_모음_전체() {
		s := new(dd.S종목별_일일_가격정보_모음)
		s.DB읽기with시작일(db, 종목코드, 한달전)

		if len(s.M저장소) > 0 && s.M저장소[len(s.M저장소)-1].M일자 == 당일 && s.M저장소[len(s.M저장소)-1].M거래량 > 0 {
			continue
		}

		// 당일 데이터 없는 경우에만 수집 대상에 추가.
		종목코드_모음 = append(종목코드_모음, 종목코드)
	}

	if len(종목코드_모음) == 0 {
		return
	}

	dd.F일일_가격정보_테이블_생성(db)

	당일_가격정보_맵 := lib.F확인2(xing.TrT8407_현물_멀티_현재가_조회(종목코드_모음))

	for 종목코드, 값 := range 당일_가격정보_맵 {
		s := new(dd.S일일_가격정보)
		s.M종목코드 = 종목코드
		s.M일자 = 당일
		s.M시가 = float64(값.M시가)
		s.M고가 = float64(값.M고가)
		s.M저가 = float64(값.M저가)
		s.M종가 = float64(값.M현재가)
		s.M거래량 = float64(값.M누적_거래량)

		종목별_일일_가격정보_모음 := lib.F확인2(dd.New종목별_일일_가격정보_모음([]*dd.S일일_가격정보{s}))
		lib.F확인1(종목별_일일_가격정보_모음.DB저장(db))
	}

	return nil
}

func F일개월_일일_가격정보_수집(db *sql.DB, 종목코드_모음 []string) (에러 error) {
	return F고정_기간_일일_가격정보_수집(db, 종목코드_모음, 31*lib.P1일, true)
}

func F일년_일일_가격정보_수집(db *sql.DB, 종목코드_모음 []string) (에러 error) {
	return F고정_기간_일일_가격정보_수집(db, 종목코드_모음, lib.P1년, true)
}

func F고정_기간_일일_가격정보_수집(db *sql.DB, 종목코드_모음 []string, 기간 time.Duration, 추가_인수 ...bool) (에러 error) {
	if len(종목코드_모음) == 0 {
		return nil
	}

	dd.F일일_가격정보_테이블_생성(db)

	시작일 := lib.F금일().Add(-1 * 기간)

	출력_여부 := true
	if len(추가_인수) > 0 {
		출력_여부 = 추가_인수[0]
	}

	종목코드_맵 := lib.F2맵(종목코드_모음) // 종목 순서를 랜덤화

	i := 0

	for 종목코드 := range 종목코드_맵 {
		if lib.F공통_종료_채널_닫힘() {
			return nil
		}

		f일일_가격정보_수집_도우미(db, 종목코드, 시작일, i, len(종목코드_맵), 출력_여부)
		i++

		lib.F대기(lib.P4초) // TR 한도 초과 방지.
	}

	return nil
}

func F일일_가격정보_수집(db *sql.DB, 종목코드_모음 []string, 추가_인수 ...bool) (에러 error) {
	var 시작일, 마지막_저장일 time.Time
	var 종목별_일일_가격정보_모음 *dd.S종목별_일일_가격정보_모음

	dd.F일일_가격정보_테이블_생성(db)

	출력_여부 := lib.F조건값(len(추가_인수) > 0, 추가_인수[0], true)

	for i, 종목코드 := range 종목코드_모음 {
		종목별_일일_가격정보_모음 = lib.F확인2(dd.New종목별_일일_가격정보_모음_DB읽기(db, 종목코드))

		// 시작일 설정
		시작일 = lib.F지금().AddDate(-30, 0, 0)
		if 에러 == nil && len(종목별_일일_가격정보_모음.M저장소) > 0 {
			// lib.S종목별_일일_가격정보_모음 는 일자 순서로 정렬되어 있음.
			마지막_저장일 = 종목별_일일_가격정보_모음.M저장소[len(종목별_일일_가격정보_모음.M저장소)-1].G일자2()
			시작일 = 마지막_저장일.AddDate(0, 0, 1)
		}

		if 시작일.After(xing.F당일()) {
			//lib.F문자열_출력("%v [%v] : 최신 데이터 업데이트.", i, 종목코드)
			continue
		} else if 시작일.After(lib.F금일().AddDate(0, 0, -14)) {
			// 데이터 수량이 1개이나 100개이나 소요 시간은 비슷함.
			시작일 = lib.F금일().AddDate(0, 0, -14)
		}

		f일일_가격정보_수집_도우미(db, 종목코드, 시작일, i, len(종목코드_모음), 출력_여부)
	}

	return nil
}

func f일일_가격정보_수집_도우미(db *sql.DB, 종목코드 string, 시작일 time.Time, i, 전체_수량 int, 출력_여부 bool) {
	var 종료일 time.Time

	// 종료일 설정
	if lib.F지금().After(xing.F당일().Add(15*lib.P1시간 + lib.P30분)) {
		종료일 = xing.F당일()
	} else {
		종료일 = xing.F전일()
	}

	// 시작일 오류 확인
	if 시작일 = lib.F2일자(시작일); 시작일.After(종료일) {
		return
	} else if 시작일.Equal(종료일) { // 시작일과 종료일이 같으면 수천 개의 데이터를 불러오는 현상이 있음.
		시작일 = 시작일.AddDate(0, 0, -1)
	}

	// 데이터 수집
	값_모음, 에러 := xing.TrT8410_현물_차트_일주월년(종목코드, 시작일, 종료일, xt.P일주월_일, false)
	if 에러 != nil {
		lib.F에러_출력(에러)
		return
	} else if len(값_모음) == 0 {
		return // 추가 저장할 데이터가 없음.
	}

	금일 := lib.F금일()
	일일_가격정보_슬라이스 := make([]*dd.S일일_가격정보, len(값_모음))
	폐장_전 := lib.F지금().Before(mt.F금일_보정_시각(15, 30, 0))

	for j, 일일_데이터 := range 값_모음 {
		if 일일_데이터.M거래량 == 0 {
			continue
		} else if 일일_데이터.M일자.Equal(금일) && 폐장_전 {
			continue // 폐장 전에 수집된 금일 데이터 제외.
		}

		일일_가격정보_슬라이스[j] = dd.New일일_가격정보(
			일일_데이터.M종목코드,
			일일_데이터.M일자,
			일일_데이터.M시가,
			일일_데이터.M고가,
			일일_데이터.M저가,
			일일_데이터.M종가,
			일일_데이터.M거래량)
	}

	if 출력_여부 {
		lib.F문자열_출력("%v 일일 가격 정보 수집 (%v/%v) : %v %v~%v %v개\n",
			lib.F지금().Format("15:04"), i+1, 전체_수량,
			xing.F종목_식별_문자열(종목코드), 시작일.Format(lib.P일자_형식), 종료일.Format(lib.P일자_형식), len(값_모음))
	} else if i > 0 && i%100 == 0 {
		lib.F문자열_출력("%v 일일 가격 정보 수집 (%v/%v) %.1f%%",
			lib.F지금().Format("15:04"), i+1, 전체_수량, float64(i+1)/float64(전체_수량)*100)
	} else if (출력_여부 || 전체_수량 > 100) && i == 전체_수량-1 {
		lib.F문자열_출력("%v 일일 가격 정보 수집 (%v/%v) 100%%", lib.F지금().Format("15:04"), 전체_수량, 전체_수량)
	}

	종목별_일일_가격정보_모음, 에러 := dd.New종목별_일일_가격정보_모음(일일_가격정보_슬라이스)
	if 에러 != nil {
		lib.F에러_출력(에러)
		return
	}

	lib.F확인1(종목별_일일_가격정보_모음.DB저장(db))
}
