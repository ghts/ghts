/* Copyright (C) 2015-2019 김운하(UnHa Kim)  unha.kim.ghts@gmail.com

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

Copyright (C) 2015-2019년 UnHa Kim (unha.kim.ghts@gmail.com)

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

package sh

import (
	st "github.com/ghts/ghts/shinhan/base"
	"github.com/ghts/ghts/lib"
	"strings"
	"testing"
)

func TestStockMst_현물_종목코드_조회(t *testing.T) {
	t.Parallel()

	값_모음, 에러 := Tr현물_종목코드_조회_stock_mst()
	lib.F테스트_에러없음(t, 에러)

	for _, 값 := range 값_모음 {
		lib.F테스트_참임(t, strings.Contains(값.M표준코드, 값.M종목코드), 값.M표준코드, 값.M종목코드)
		lib.F테스트_같음(t, len(값.M종목코드), 6)
		lib.F테스트_같음(t, 값.M장구분, lib.P시장구분_코스피, lib.P시장구분_코스닥)
		lib.F체크포인트(값.M종목명)
		lib.F테스트_다름(t, 값.M종목명, "")
		lib.F테스트_같음(t, 값.M업종, st.P업종_미분류, st.P업종_제조업, st.P업종_전기통신, st.P업종_건설, st.P업종_유통서비스, st.P업종_금융)
		lib.F체크포인트(값.M결산월일)
		lib.F테스트_다름(t, 값.M결산월일, "")
		lib.F테스트_같음(t, 값.M거래정지구분, st.P거래_정상, st.P거래_정지, st.P거래_CB발동)
		lib.F테스트_같음(t, 값.M관리구분, st.P종목_정상, st.P종목_관리)
		lib.F테스트_같음(t, 값.M시장경보구분, st.P시장_정상, st.P시장_주의, st.P시장_경고, st.P시장_위험)
		lib.F테스트_같음(t, 값.M락구분, st.P락_미발생, st.P락_권리락, st.P락_배당락, st.P락_분배락, st.P락_권배락, st.P락_중간배당락, st.P락_권리중간배당락, st.P락_기타)
		lib.F테스트_같음(t, 값.M불성실공시지정여부, true, false)
		lib.F테스트_같음(t, 값.M증거금_퍼센트, 15, 20, 25, 100)
		lib.F테스트_같음(t, 값.M신용증거금_구분, "A", "B", "C")
		lib.F테스트_같음(t, 값.ETF_구분, st.ETF_일반형, st.ETF_투자회사형, st.ETF_수익증권형)
		lib.F테스트_같음(t, 값.M증권그룹, st.P증권_주식, st.P증권_증권투자회사, st.P증권_리츠, st.P증권_선박투자회사,
			st.P증권_인프라투융자회사, st.P증권_예탁증서, st.P증권_신주인수권증권, st.P증권_신주인수권증서,
			st.P증권_주식워런트증권_ELW, st.P증권_선물, st.P증권_옵션, st.P증권_상장지수펀드_ETF,
			st.P증권_수익증권, st.P증권_해외_ETF, st.P증권_해외_주식, st.P증권_ETN)
	}
}
