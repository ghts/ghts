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
	"database/sql"
	"github.com/ghts/ghts/lib"
	xing "github.com/ghts/ghts/xing/go"
	"math"
)

func New종목별_멀티_팩터_데이터(
	종목별_일일_가격정보_모음 *lib.S종목별_일일_가격정보_모음,
	내재가치_정보_모음 *lib.S내재가치_정보_모음) (s *S종목별_멀티_팩터_데이터) {

	// 추세 점수 산출에는 11개월치 데이터가 필요함.
	if len(종목별_일일_가격정보_모음.M저장소) < 232 {
		return nil
	}

	defer lib.S예외처리{M함수: func() { s = nil }}.S실행()

	종목코드 := 종목별_일일_가격정보_모음.M저장소[0].M종목코드
	_6개월전_종가 := 종목별_일일_가격정보_모음.M저장소[len(종목별_일일_가격정보_모음.M저장소)-(6*21)].M종가
	_9개월전_종가 := 종목별_일일_가격정보_모음.M저장소[len(종목별_일일_가격정보_모음.M저장소)-(9*21)].M종가

	s = new(S종목별_멀티_팩터_데이터)
	s.S종목별_공통_데이터 = new(S종목별_공통_데이터)
	s.S내재가치_정보 = 내재가치_정보_모음.G종목별_최신_정보(종목코드)

	if s.S내재가치_정보 == nil ||
		s.S내재가치_식별정보 == nil ||
		s.S재무제표_정보_내용 == nil ||
		s.S재무비율_정보_내용 == nil ||
		s.S투자지표_정보_내용 == nil {
		return nil
	}

	s.M종목코드 = 종목코드
	s.M기준가 = 종목별_일일_가격정보_모음.G종가()
	s.M시가총액 = s.M상장주식수 * s.M기준가
	s.EV = s.M시가총액 + s.M부채 - s.M현금및현금성자산
	s.EV_EBITDA = s.EV / (s.EBITDAPS * s.M상장주식수)
	s.EV_Sales = s.EV / s.M매출액
	s.EV_FCF = s.EV / (s.CFPS * s.M상장주식수)
	s.PBR = s.M기준가 / s.BPS
	s.PER = s.M기준가 / s.EPS
	s.PSR = s.M기준가 / s.SPS
	s.PCR = s.M기준가 / s.CFPS
	s.DPR = s.DPS / s.M기준가
	s.M추세점수 = 종목별_일일_가격정보_모음.G추세_점수()
	s.M6개월_수익율 = (s.M기준가 - _6개월전_종가) / _6개월전_종가
	s.M9개월_수익율 = (s.M기준가 - _9개월전_종가) / _9개월전_종가
	s.M수익율_변동성 = 종목별_일일_가격정보_모음.G월수익율_변동성_평균(21)
	s.GPA = s.M매출총이익 / s.M자산
	s.APR = (s.M당기순이익 - s.M영업_현금흐름) / s.M기준가
	s.AAR = (s.M당기순이익 - s.M영업_현금흐름) / s.M자산
	s.CDR = s.M현금_증가 / s.M부채
	s.VPCI = 종목별_일일_가격정보_모음.VPCI(10, 50)
	s.VPCIs = 종목별_일일_가격정보_모음.VPCIs(10, 50)
	s.MFI = 종목별_일일_가격정보_모음.MFI(14)
	s.MFIs = 종목별_일일_가격정보_모음.MFIs(14, 9)

	차최신_정보 := 내재가치_정보_모음.G종목별_차최신_정보(종목코드)
	if 차최신_정보 == nil ||
		차최신_정보.S내재가치_식별정보 == nil ||
		차최신_정보.S재무제표_정보_내용 == nil ||
		차최신_정보.S재무비율_정보_내용 == nil ||
		차최신_정보.S투자지표_정보_내용 == nil {
		return nil
	}

	s.M부채증가율 = 차최신_정보.M부채_비율 - s.M부채_비율

	return s
}

// 멀티 팩터 방법론은 다음 서적 참조할 것. (간단한 서적 -> 복잡한 서적)
// - '주식 시장을 이기는 작은 책 (조엘 그린블라트 저).
// - '할 수 있다! 퀀트투자'(강환국 저).
// - 'What works on Wall Street'(James P, O'Shaughnessy 저).
type S종목별_멀티_팩터_데이터 struct {
	*S종목별_공통_데이터
	*lib.S내재가치_정보
	M기준가  float64 // 산출 시점에서 가장 최근의 종가. 대부분의 경우 '전일 종가'가 될 것임.
	M시가총액 float64
	EV    float64
	// -- 가치 팩터 --
	EV_EBITDA   float64
	EV_EBITDA등급 float64
	EV_Sales    float64
	EV_Sales등급  float64
	EV_FCF      float64
	EV_FCF등급    float64
	PBR         float64
	PBR등급       float64
	PER         float64
	PER등급       float64
	PSR         float64
	PSR등급       float64
	PCR         float64
	PCR등급       float64
	DPR         float64 // DPS / Price Rate : 배당수익율과 같은 효과. 다른 가치 지수와 반대로 높을수록 좋다. 배당금이 없는 경우가 많아서 가격이 분모임.
	DPR등급       float64
	// --- 추세 팩터 --
	M추세점수       float64
	M추세점수_등급    float64
	M6개월_수익율    float64
	M6개월_수익율_등급 float64
	M9개월_수익율    float64
	M9개월_수익율_등급 float64
	// --- 변동성 팩터 ---
	M수익율_변동성    float64
	M수익율_변동성_등급 float64
	// --- 퀄리티 팩터 --
	GPA   float64 // 매출총이익 / 자산
	GPA등급 float64
	//ROIC float64	// '*S내재가치_정보'에 이미 포함되어 있음.
	ROIC등급 float64
	//ROE float64	// '*S내재가치_정보'에 이미 포함되어 있음. // 하위 10% 그룹은 걸러야 한다.
	ROE등급 float64
	//ROA float64	// '*S내재가치_정보'에 이미 포함되어 있음. // 하위 10% 그룹은 걸러야 한다.
	ROA등급     float64
	APR       float64 // = Accrual / Price  = (당기순이익 - 영업현금흐름) / 현재가 : 낮을수록 분식회계 위험이 낮다. 상위 10% 그룹은 걸러야 한다.
	APR등급     float64
	AAR       float64 // = Accrual / Asseet = (당기순이익 - 영업현금흐름) / 총자산 : 낮을수록 분식회계 위험이 낮다. 상위 10% 그룹은 걸러야 한다.
	AAR등급     float64
	CDR       float64 // Cash flow / Debt = 현금 흐름 / 부채 : 하위 10% 그룹은 걸러야 한다.
	CDR등급     float64
	M부채증가율    float64 // 부채증가율 : 상위 10% 그룹은 걸러내어야 한다.
	M부채증가율_등급 float64
	// --- 거래량 팩터 --
	MFI   float64
	MFIs  float64
	VPCI  float64
	VPCIs float64
	// --- 최종 등급 --
	M복합_등급 float64
}

func F종목별_멀티_팩터_데이터_DB읽기(db *sql.DB) (값_모음 []*S종목별_멀티_팩터_데이터, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	값_모음 = make([]*S종목별_멀티_팩터_데이터, 0)
	내재가치_정보_모음 := lib.New내재가치_정보_모음_DB읽기(db)

	for i, 종목코드 := range xing.F종목코드_모음_전체() {
		if xing.ETF_ETN_종목_여부(종목코드) {
			continue
		}

		종목별_일일_가격정보_모음, 에러 := lib.New종목별_일일_가격정보_모음_DB읽기(db, 종목코드)
		if 에러 != nil || 종목별_일일_가격정보_모음 == nil {
			continue
		}

		if 값 := New종목별_멀티_팩터_데이터(종목별_일일_가격정보_모음, 내재가치_정보_모음); 값 != nil {
			값_모음 = append(값_모음, 값)
		}

		if i > 0 && i%100 == 0 {
			lib.F체크포인트(i, len(값_모음))
		}
	}

	return 값_모음, nil
}

func F종목별_멀티_팩터_데이터_등급_산출(멀티_팩터_데이터_모음 []*S종목별_멀티_팩터_데이터) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	s := new(S종목별_멀티_팩터_데이터_정렬_도우미)
	s.M저장소 = 멀티_팩터_데이터_모음

	// EV_EBITDA 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].EV_EBITDA < s.M저장소[j].EV_EBITDA
	}).S정렬()

	for i, 값 := range s.M저장소 {
		값.EV_EBITDA등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// EV_Sales 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].EV_Sales < s.M저장소[j].EV_Sales
	}).S정렬()

	for i, 값 := range s.M저장소 {
		값.EV_Sales등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// EV_FCF 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].EV_FCF < s.M저장소[j].EV_FCF
	}).S정렬()

	for i, 값 := range s.M저장소 {
		값.EV_FCF등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// PBR 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].PBR < s.M저장소[j].PBR
	}).S정렬()

	for i, 값 := range s.M저장소 {
		값.PBR등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// PER 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].PER < s.M저장소[j].PER
	}).S정렬()

	for i, 값 := range s.M저장소 {
		값.PER등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// PSR 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].PSR < s.M저장소[j].PSR
	}).S정렬()

	for i, 값 := range s.M저장소 {
		값.PSR등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// PCR 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].PCR < s.M저장소[j].PCR
	}).S정렬()

	for i, 값 := range s.M저장소 {
		값.PCR등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// DPR 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].DPR < s.M저장소[j].DPR
	}).S역순_정렬()

	for i, 값 := range s.M저장소 {
		값.DPR등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// 추세 등급 백분율 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].M추세점수 < s.M저장소[j].M추세점수
	}).S역순_정렬()

	for i, 값 := range s.M저장소 {
		값.M추세점수_등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// 6개월 수익율 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].M6개월_수익율 < s.M저장소[j].M6개월_수익율
	}).S역순_정렬()

	for i, 값 := range s.M저장소 {
		값.M6개월_수익율_등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// 9개월 수익율 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].M9개월_수익율 < s.M저장소[j].M9개월_수익율
	}).S역순_정렬()

	for i, 값 := range s.M저장소 {
		값.M9개월_수익율_등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// 수익율 변동성 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].M수익율_변동성 < s.M저장소[j].M수익율_변동성
	}).S정렬()

	for i, 값 := range s.M저장소 {
		값.M수익율_변동성_등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// GPA 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].GPA < s.M저장소[j].GPA
	}).S역순_정렬()

	for i, 값 := range s.M저장소 {
		값.GPA등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// ROIC 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].ROIC < s.M저장소[j].ROIC
	}).S역순_정렬()

	for i, 값 := range s.M저장소 {
		값.ROIC등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// ROE 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].ROE < s.M저장소[j].ROE
	}).S역순_정렬()

	for i, 값 := range s.M저장소 {
		값.ROE등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// ROA 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].ROA < s.M저장소[j].ROA
	}).S역순_정렬()

	for i, 값 := range s.M저장소 {
		값.ROA등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
		//lib.F체크포인트(i, 값.M종목코드, 값.ROA등급, 값.ROA)
	}

	// APR 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].APR < s.M저장소[j].APR
	}).S정렬()

	for i, 값 := range s.M저장소 {
		값.APR등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// AAR 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].AAR < s.M저장소[j].AAR
	}).S정렬()

	for i, 값 := range s.M저장소 {
		값.AAR등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// CDR 등급 산출
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].CDR < s.M저장소[j].CDR
	}).S역순_정렬()

	for i, 값 := range s.M저장소 {
		값.CDR등급 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	// 부채비율 증가율 등급
	s.S정렬_함수_설정(func(s *S종목별_멀티_팩터_데이터_정렬_도우미, i, j int) bool {
		return s.M저장소[i].M부채증가율 < s.M저장소[j].M부채증가율
	}).S정렬()

	for i, 값 := range s.M저장소 {
		값.M부채증가율 = math.Floor(float64(i)/float64(len(s.M저장소))*100.0) + 1
	}

	return nil
}
