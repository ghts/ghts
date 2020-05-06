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

package pd

import (
	"github.com/ghts/ghts/lib"
)

// 각 항목의 활용법은 다음 서적 참조할 것.
// '할 수 있다! 퀀트투자'(강환국 저).
// 'What works on Wall Street'(James P, O'Shaughnessy 저).
type S종목별_멀티_팩터_데이터 struct {
	*S종목별_공통_데이터
	*lib.S내재가치_정보
	M시가총액 float64
	M시가총액_순위 int
	EV float64
	// -- 가치 팩터 --
	EV_EBITDA float64
	EV_EBITDA순위 int
	EV_Sales float64
	EV_Sales순위 int
	EV_FCF float64
	EV_FCF순위 int
	PBR float64
	PBR순위 int
	PER float64
	PER순위 int
	PSR float64
	PSR순위 int
	PCR float64
	PCR순위 int
	DPR float64 // DPS / Price Rate : 배당수익율과 같은 효과. 높을 수록 좋다.
	DPR순위 int
	// --- 추세 팩터 --
	M추세점수 float64
	M추세점수_순위 int
	// --- 퀄리티 팩터 --
	GPA float64	// 매출총이익 / 자산
	GPA순위 int
	//ROIC float64	// '*S내재가치_정보'에 이미 포함되어 있음.
	ROIC순위 int
	//ROE float64	// '*S내재가치_정보'에 이미 포함되어 있음. // 하위 10% 그룹은 걸러야 한다.
	ROE순위 int
	//ROA float64	// '*S내재가치_정보'에 이미 포함되어 있음. // 하위 10% 그룹은 걸러야 한다.
	ROA순위 int
	APR float64 // = Accrual / Price  = (당기순이익 - 영업현금흐름) / 현재가 : 상위 10% 그룹은 걸러야 한다.
	APR순위 int
	AAR float64 // = Accrual / Asseet = (당기순이익 - 영업현금흐름) / 총자산 : 상위 10% 그룹은 걸러야 한다.
	AAR순위 int
	CDR float64 // Cash flow / Debt = 현금 흐름 / 부채 : 하위 10% 그룹은 걸러야 한다.
	CDR순위 int
	M부채증가율 float64 // 부채증가율 : 상위 10% 그룹은 걸러내어야 한다.
	M부채증가율_순위 int
}

