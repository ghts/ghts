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

import (
	"github.com/ghts/ghts/lib"
	"math"
	"strings"
	"testing"
)

func TestT3320_F기업정보_요약(t *testing.T) {
	t.Parallel()

	종목코드 := "051910" // LG전자

	값, 에러 := TrT3320_F기업정보_요약(종목코드)
	lib.F테스트_에러없음(t, 에러)

	값1 := 값.M응답1
	lib.F테스트_다름(t, strings.TrimSpace(값1.M업종구분명), "")
	lib.F테스트_다름(t, strings.TrimSpace(값1.M시장구분), "")
	lib.F테스트_다름(t, strings.TrimSpace(값1.M시장구분명), "")
	lib.F테스트_다름(t, strings.TrimSpace(값1.M한글기업명), "")
	lib.F테스트_다름(t, strings.TrimSpace(값1.M본사주소), "")
	lib.F테스트_다름(t, lib.F정규식_검색(값1.M본사전화번호, []string{`[1-9]+`}), "")
	lib.F테스트_다름(t, lib.F정규식_검색(값1.M최근결산년도, []string{`[1-9]+`}), "")
	lib.F테스트_다름(t, lib.F정규식_검색(값1.M결산월, []string{`[1-9]+`}), "")
	lib.F테스트_다름(t, lib.F정규식_검색(값1.M최근결산년월, []string{`[1-9]+`}), "")
	lib.F테스트_참임(t, 값1.M주당액면가 > 0)
	lib.F테스트_참임(t, 값1.M주식수 > 0)
	lib.F테스트_다름(t, strings.TrimSpace(값1.M홈페이지), "")
	//lib.F테스트_다름(t, strings.TrimSpace(값1.M그룹명), "")
	lib.F테스트_참임(t, 값1.M외국인_비중 > 0)
	lib.F테스트_다름(t, lib.F정규식_검색(값1.M주담전화, []string{`[1-9]+`}), "")
	lib.F테스트_참임(t, 값1.M자본금_억 > 0)
	lib.F테스트_참임(t, 값1.M시가총액 > 0)
	lib.F테스트_참임(t, 값1.M배당금 >= 0)
	lib.F테스트_참임(t, 값1.M배당수익율 >= 0)
	lib.F테스트_참임(t, 값1.M현재가 > 0)
	lib.F테스트_참임(t, 값1.M전일종가 > 0)
	lib.F테스트_참임(t, 값1.M현재가 > int64(float64(값1.M전일종가)*0.4) || 값1.M현재가 < int64(float64(값1.M전일종가)*1.4))

	값2 := 값.M응답2
	lib.F테스트_같음(t, strings.TrimSpace(값2.M종목코드), 종목코드)
	lib.F테스트_다름(t, strings.TrimSpace(값2.M결산년월), "")
	lib.F테스트_다름(t, strings.TrimSpace(값2.M결산구분), "")
	lib.F테스트_참임(t, math.Abs(값2.PER) < 500, 값2.PER)
	//값2.EPS      = lib.F2실수_단순형(g.Eps)
	lib.F테스트_참임(t, 값2.PBR > 0)
	lib.F테스트_참임(t, math.Abs(값2.ROA) < 100, 값2.ROA)
	lib.F테스트_참임(t, math.Abs(값2.ROE) < 100, 값2.ROE)
	//값2.EBITDA   = lib.F2실수_단순형(g.Ebitda)
	lib.F테스트_참임(t, math.Abs(값2.EVEBITDA) < 30)
	lib.F테스트_참임(t, 값2.M액면가 > 0)
	//값2.SPS      = lib.F2실수_단순형(g.Sps)
	//값2.CPS      = lib.F2실수_단순형(g.Cps)
	//값2.BPS      = lib.F2실수_단순형(g.Bps)
	//값2.T_PER    = lib.F2실수_단순형(g.Tper)
	//값2.T_EPS    = lib.F2실수_단순형(g.Teps)
	//값2.PEG      = lib.F2실수_단순형(g.Peg)
	//값2.T_PEG    = lib.F2실수_단순형(g.Tpeg)
	lib.F테스트_다름(t, lib.F정규식_검색(값2.M최근분기년도, []string{`[1-9]+`}), "")
}
