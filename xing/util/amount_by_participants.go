/* Copyright (C) 2015-2023 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2023년 UnHa Kim (unha.kim@ghts.org)

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
	"database/sql"
	"github.com/ghts/ghts/lib"
	mt "github.com/ghts/ghts/lib/market_time"
	dd "github.com/ghts/ghts/lib/unused"
	xt "github.com/ghts/ghts/xing/base"
	xing "github.com/ghts/ghts/xing/go"
	"time"
)

func F매매주체_동향_수집(db *sql.DB, 종목코드_모음 []string, 시작일 time.Time, 출력_여부 bool) (에러 error) {
	if len(종목코드_모음) == 0 {
		return nil
	}

	dd.F매매주체_동향_정보_테이블_생성(db)

	종목코드_맵 := lib.F2맵(종목코드_모음) // 종목 순서를 랜덤화

	i := 0

	for 종목코드 := range 종목코드_맵 {
		if lib.F공통_종료_채널_닫힘() {
			return nil
		}

		f매매주체_동향_수집_도우미(db, 종목코드, 시작일, i, len(종목코드_맵), 출력_여부)
		i++

		lib.F대기(lib.P4초) // TR 한도 초과 방지.
	}

	return nil
}

func f매매주체_동향_수집_도우미(db *sql.DB, 종목코드 string, 시작일 time.Time, i, 전체_수량 int, 출력_여부 bool) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	종료일 := lib.F금일()

	if lib.F지금().Before(mt.F금일_보정_시각(6, 0, 0)) {
		종료일 = 종료일.AddDate(0, 0, -1) // 폐장 이전 데이터는 확정 데이터가 아니다.
	}

	// 시작일 오류 확인
	if 시작일 = lib.F2일자(시작일); !시작일.Before(종료일) {
		시작일 = 종료일.AddDate(0, 0, -1)
	}

	// 데이터 수집
	값_모음, 에러 := xing.TrT1717_종목별_매매주체_동향(종목코드, xt.P일별_순매수_T1717, 시작일, 종료일)
	if 에러 != nil {
		lib.F에러_출력(에러)
		return
	} else if len(값_모음) == 0 {
		return // 추가 저장할 데이터가 없음.
	}

	매매주체_동향_모음 := make([]*dd.S종목별_매매주체_동향, len(값_모음))

	for j, 값 := range 값_모음 {
		if 값.M거래량 == 0 || (값.M개인_순매수량 == 0 && 값.M기관_순매수량 == 0 && 값.M외인계_순매수량 == 0) {
			continue // 오류가 의심되거나 의미없는 데이터는 건너뜀.
		} else if 값.M일자.Weekday() == time.Saturday || 값.M일자.Weekday() == time.Sunday {
			continue // 주말 데이터 수집 중 발생하는 데이터 오류 건너뜀.
		}

		매매주체_동향_모음[j] = dd.New종목별_매매주체_동향(
			값.M종목코드,
			값.M일자,
			float64(값.M기관_순매수량*값.M기관_단가),
			float64(값.M외인계_순매수량*값.M외인계_단가),
			float64(값.M개인_순매수량*값.M개인_단가))
	}

	if 출력_여부 {
		lib.F문자열_출력("매매주체 동향 정보 수집 (%v/%v) : %v %v개\n",
			i+1, 전체_수량, xing.F종목_식별_문자열(종목코드), len(값_모음))
	} else if i > 0 && i%100 == 0 {
		lib.F문자열_출력("매매주체 동향 정보 수집 (%v/%v) %.1f%%",
			i+1, 전체_수량, float64(i+1)/float64(전체_수량)*100)
	} else if (출력_여부 || 전체_수량 > 100) && i == 전체_수량-1 {
		lib.F문자열_출력("매매주체 동향 정보 수집 (%v/%v) 100%%", 전체_수량, 전체_수량)
	}

	return dd.F종목별_매매주체_동향_모음_DB저장(db, 매매주체_동향_모음)
}
