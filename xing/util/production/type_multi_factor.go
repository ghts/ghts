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

func New종목별_멀티_팩터_데이터(
	종목별_일일_가격정보_모음 *lib.S종목별_일일_가격정보_모음,
	내재가치_정보_모음 *lib.S내재가치_정보_모음) (s *S종목별_멀티_팩터_데이터) {
	defer lib.S예외처리{M함수: func() { s = nil }}.S실행()

	종목코드 := 종목별_일일_가격정보_모음.M저장소[0].M종목코드

	s = new(S종목별_멀티_팩터_데이터)
	s.S종목별_공통_데이터 = new(S종목별_공통_데이터)
	s.S내재가치_정보 = 내재가치_정보_모음.G종목별_최신_정보(종목코드)
	s.M종목코드 = 종목코드
	s.M기준가 = 종목별_일일_가격정보_모음.G최근_종가()
	s.M시가총액 = s.M상장주식수 * s.M기준가
	s.EV = s.M시가총액 + s.M부채 - s.M현금및현금성자산
	s.EV_EBITDA = s.EV / (s.EBITDAPS * s.M상장주식수)
	s.EV_Sales = s.EV / s.M매출액
	s.EV_FCF = s.EV / s.FCFF
	s.PBR = s.M기준가 / s.BPS
	s.PER = s.M기준가 / s.EPS
	s.PSR = s.M기준가 / s.SPS
	s.PCR = s.M기준가 / s.CFPS
	s.DPR = s.DPS / s.M기준가
	s.M추세점수 = 종목별_일일_가격정보_모음.G최근_추세_점수()
	s.GPA = s.M매출총이익 / s.M자산
	s.APR = (s.M당기순이익 - s.M영업_현금흐름) / s.M기준가
	s.AAR = (s.M당기순이익 - s.M영업_현금흐름) / s.M자산
	s.CDR = s.M현금_증가 / s.M부채
	s.M부채증가율 = 내재가치_정보_모음.G종목별_차최신_정보(종목코드).M부채_비율 - s.M부채_비율

	return s
}

// 각 항목의 활용법은 다음 서적 참조할 것.
// '할 수 있다! 퀀트투자'(강환국 저).
// 'What works on Wall Street'(James P, O'Shaughnessy 저).
type S종목별_멀티_팩터_데이터 struct {
	*S종목별_공통_데이터
	*lib.S내재가치_정보
	M기준가     float64 // 산출 시점에서 가장 최근의 종가. 대부분의 경우 '전일 종가'가 될 것임.
	M시가총액    float64
	EV       float64
	// -- 가치 팩터 --
	EV_EBITDA   float64
	EV_Sales    float64
	EV_FCF      float64
	PBR         float64
	PER         float64
	PSR         float64
	PCR         float64
	DPR         float64 // DPS / Price Rate : 배당수익율과 같은 효과. 높을 수록 좋다. 배당금이 없는 경우가 많아서 가격이 분모임.
	// --- 추세 팩터 --
	M추세점수    float64
	// --- 퀄리티 팩터 --
	GPA   float64 // 매출총이익 / 자산
	//ROIC float64	// '*S내재가치_정보'에 이미 포함되어 있음.
	//ROE float64	// '*S내재가치_정보'에 이미 포함되어 있음. // 하위 10% 그룹은 걸러야 한다.
	//ROA float64	// '*S내재가치_정보'에 이미 포함되어 있음. // 하위 10% 그룹은 걸러야 한다.
	APR       float64 // = Accrual / Price  = (당기순이익 - 영업현금흐름) / 현재가 : 상위 10% 그룹은 걸러야 한다.
	AAR       float64 // = Accrual / Asseet = (당기순이익 - 영업현금흐름) / 총자산 : 상위 10% 그룹은 걸러야 한다.
	CDR       float64 // Cash flow / Debt = 현금 흐름 / 부채 : 하위 10% 그룹은 걸러야 한다.
	M부채증가율    float64 // 부채증가율 : 상위 10% 그룹은 걸러내어야 한다.
}
