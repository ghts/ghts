/* Copyright (C) 2015-2020 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright (C) 2015-2020년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	xing "github.com/ghts/ghts/xing/go"

	"database/sql"
	"time"
)

func F일일_가격정보_수집(db *sql.DB, 종목코드_모음 []string) (에러 error) {
	var 시작일, 종료일, 마지막_저장일 time.Time
	var 종목별_일일_가격정보_모음 *lib.S종목별_일일_가격정보_모음

	lib.F일일_가격정보_테이블_생성(db)

	for i, 종목코드 := range 종목코드_모음 {
		for {
			if xing.C32_재시작_실행_중.G값() {
				lib.F대기(lib.P10초)
				continue
			}

			break
		}

		종목별_일일_가격정보_모음, 에러 = lib.New종목별_일일_가격정보_모음_DB읽기(db, 종목코드)
		lib.F확인(에러)

		// 시작일 설정
		시작일 = lib.F지금().AddDate(-30, 0, 0)
		if 에러 == nil && len(종목별_일일_가격정보_모음.M저장소) > 0 {
			// lib.S종목별_일일_가격정보_모음 는 일자 순서로 정렬되어 있음.
			마지막_저장일 = 종목별_일일_가격정보_모음.M저장소[len(종목별_일일_가격정보_모음.M저장소)-1].G일자2()
			시작일 = 마지막_저장일.AddDate(0, 0, 1)
		}

		if 시작일.After(xing.F당일()) {
			lib.F문자열_출력("%v [%v] : 최신 데이터 업데이트.", i, 종목코드)
			continue
		}

		// 종료일 설정
		if lib.F지금().After(xing.F당일().Add(15*lib.P1시간 + lib.P30분)) {
			종료일 = xing.F당일()
		} else {
			종료일 = xing.F전일()
		}

		if 시작일.After(종료일) {
			continue
		} else if 시작일.Equal(종료일) { // 시작일과 종료일이 같으면 수천 개의 데이터를 불러오는 현상이 있음.
			시작일 = 시작일.AddDate(0, 0, -1)
		}

		// 데이터 수집
		값_모음, 에러 := xing.TrT8413_현물_차트_일주월(종목코드, 시작일, 종료일, xt.P일주월_일)
		if 에러 != nil {
			lib.F에러_출력(에러)
			continue
		} else if len(값_모음) == 0 {
			lib.F체크포인트(i, 종목코드, "추가 저장할 데이터가 없음.")
			continue // 추가 저장할 데이터가 없음.
		}

		일일_가격정보_슬라이스 := make([]*lib.S일일_가격정보, len(값_모음))

		for i, 일일_데이터 := range 값_모음 {
			일일_가격정보_슬라이스[i] = lib.New일일_가격정보(
				일일_데이터.M종목코드,
				일일_데이터.M일자,
				일일_데이터.M시가,
				일일_데이터.M고가,
				일일_데이터.M저가,
				일일_데이터.M종가,
				일일_데이터.M거래량)
		}

		if 종목별_일일_가격정보_모음 != nil && len(종목별_일일_가격정보_모음.M저장소) > 0 {
			일일_가격정보_슬라이스 = append(일일_가격정보_슬라이스, 종목별_일일_가격정보_모음.M저장소...)
		}

		lib.F문자열_출력("%v %v %v~%v %v개", i, 종목코드, 시작일.Format(lib.P일자_형식), 종료일.Format(lib.P일자_형식), len(값_모음))

		// 새로 생성하는 과정에서 중복 제거됨.
		종목별_일일_가격정보_모음, 에러 = lib.New종목별_일일_가격정보_모음(일일_가격정보_슬라이스)
		if 에러 != nil {
			lib.F에러_출력(에러)
			continue
		}

		lib.F확인(종목별_일일_가격정보_모음.DB저장(db))
	}

	return nil
}
