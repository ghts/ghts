/* This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>.

@author: UnHa Kim <unha.kim.ghts@gmail.com> */

package common

import (
	"math"
)

func F샘플_종목_모음_ETF() []I종목 {
	종목_모음 := make([]I종목, 0)
	종목_모음 = append(종목_모음, New종목("069500", "KODEX 200"))
	종목_모음 = append(종목_모음, New종목("069660", "KOSEF 200"))
	종목_모음 = append(종목_모음, New종목("091160", "KODEX 반도체"))
	종목_모음 = append(종목_모음, New종목("091170", "KODEX 은행"))
	종목_모음 = append(종목_모음, New종목("091180", "KODEX 자동차"))
	종목_모음 = append(종목_모음, New종목("091210", "TIGER KRX100"))
	종목_모음 = append(종목_모음, New종목("091220", "TIGER 은행"))
	종목_모음 = append(종목_모음, New종목("091230", "TIGER 반도체"))
	종목_모음 = append(종목_모음, New종목("097710", "TIGER 가치주"))
	종목_모음 = append(종목_모음, New종목("097720", "TIGER 미드캡"))
	종목_모음 = append(종목_모음, New종목("097750", "TREX 중소형가치"))
	종목_모음 = append(종목_모음, New종목("098560", "TIGER 미디어통신"))
	종목_모음 = append(종목_모음, New종목("099140", "KODEX China H"))
	종목_모음 = append(종목_모음, New종목("100910", "KOSEF KRX100"))
	종목_모음 = append(종목_모음, New종목("101280", "KODEX Japan"))
	종목_모음 = append(종목_모음, New종목("102110", "TIGER 200"))
	종목_모음 = append(종목_모음, New종목("102780", "KODEX 삼성그룹"))
	종목_모음 = append(종목_모음, New종목("102960", "KODEX 조선"))
	종목_모음 = append(종목_모음, New종목("102970", "KODEX 증권"))
	종목_모음 = append(종목_모음, New종목("104520", "KOSEF 블루칩"))
	종목_모음 = append(종목_모음, New종목("104530", "KOSEF 고배당"))
	종목_모음 = append(종목_모음, New종목("105010", "TIGER 라틴"))
	종목_모음 = append(종목_모음, New종목("105190", "KINDEX 200"))
	종목_모음 = append(종목_모음, New종목("105270", "KINDEX 성장대형F15"))
	종목_모음 = append(종목_모음, New종목("105780", "KStar 5대그룹주"))
	종목_모음 = append(종목_모음, New종목("107560", "GIANT 현대차그룹"))
	종목_모음 = append(종목_모음, New종목("108440", "KINDEX 코스닥스타"))
	종목_모음 = append(종목_모음, New종목("108450", "KINDEX 삼성그룹SW"))
	종목_모음 = append(종목_모음, New종목("108480", "KStar 코스닥엘리트"))
	종목_모음 = append(종목_모음, New종목("108590", "TREX 200"))
	종목_모음 = append(종목_모음, New종목("108630", "FIRST 스타우량"))
	종목_모음 = append(종목_모음, New종목("114100", "KStar 국고채"))
	종목_모음 = append(종목_모음, New종목("114260", "KODEX 국고채"))
	종목_모음 = append(종목_모음, New종목("114460", "KINDEX 국고채"))
	종목_모음 = append(종목_모음, New종목("114470", "KOSEF 국고채"))
	종목_모음 = append(종목_모음, New종목("114800", "KODEX 인버스"))
	종목_모음 = append(종목_모음, New종목("114820", "TIGER 국채3"))
	종목_모음 = append(종목_모음, New종목("117460", "KODEX 에너지화학"))
	종목_모음 = append(종목_모음, New종목("117680", "KODEX 철강"))
	종목_모음 = append(종목_모음, New종목("117690", "TIGER 차이나"))
	종목_모음 = append(종목_모음, New종목("117700", "KODEX 건설"))
	종목_모음 = append(종목_모음, New종목("120210", "GREAT SRI"))
	종목_모음 = append(종목_모음, New종목("122090", "ARIRANG KOSPI50"))
	종목_모음 = append(종목_모음, New종목("122260", "KOSEF 통안채"))
	종목_모음 = append(종목_모음, New종목("122390", "TIGER 코스닥프리미어"))
	종목_모음 = append(종목_모음, New종목("122630", "KODEX 레버리지"))
	종목_모음 = append(종목_모음, New종목("123310", "TIGER 인버스"))
	종목_모음 = append(종목_모음, New종목("123320", "TIGER 레버리지"))
	종목_모음 = append(종목_모음, New종목("123760", "KStar 레버리지"))
	종목_모음 = append(종목_모음, New종목("130680", "TIGER 원유선물(H)"))
	종목_모음 = append(종목_모음, New종목("130730", "KOSEF 단기자금"))
	종목_모음 = append(종목_모음, New종목("131890", "KINDEX 삼성그룹EW"))
	종목_모음 = append(종목_모음, New종목("132030", "KODEX 골드선물(H)"))
	종목_모음 = append(종목_모음, New종목("133690", "TIGER 나스닥100"))
	종목_모음 = append(종목_모음, New종목("136280", "KODEX 소비재"))
	종목_모음 = append(종목_모음, New종목("136340", "KStar 우량회사채"))
	종목_모음 = append(종목_모음, New종목("137610", "TIGER 농산물선물(H)"))
	종목_모음 = append(종목_모음, New종목("137930", "마이다스 커버드콜"))
	종목_모음 = append(종목_모음, New종목("137990", "TIGER 그린"))
	종목_모음 = append(종목_모음, New종목("138230", "KOSEF 달러선물"))
	종목_모음 = append(종목_모음, New종목("138520", "TIGER 삼성그룹"))
	종목_모음 = append(종목_모음, New종목("138530", "TIGER LG그룹+"))
	종목_모음 = append(종목_모음, New종목("138540", "TIGER 현대차그룹+"))
	종목_모음 = append(종목_모음, New종목("138910", "KODEX 구리선물(H)"))
	종목_모음 = append(종목_모음, New종목("138920", "KODEX 콩선물(H)"))
	종목_모음 = append(종목_모음, New종목("139220", "TIGER 200 건설"))
	종목_모음 = append(종목_모음, New종목("139230", "TIGER 200 중공업"))
	종목_모음 = append(종목_모음, New종목("139240", "TIGER 200 철강소재"))
	종목_모음 = append(종목_모음, New종목("139250", "TIGER 200 에너지화학"))
	종목_모음 = append(종목_모음, New종목("139260", "TIGER 200 IT"))
	종목_모음 = append(종목_모음, New종목("139270", "TIGER 200 금융"))
	종목_모음 = append(종목_모음, New종목("139280", "TIGER 경기방어"))
	종목_모음 = append(종목_모음, New종목("139290", "TIGER 200 경기소비재"))
	종목_모음 = append(종목_모음, New종목("139310", "TIGER 금속선물(H)"))
	종목_모음 = append(종목_모음, New종목("139320", "TIGER 금은선물(H)"))
	종목_모음 = append(종목_모음, New종목("139660", "KOSEF 달러인버스선물"))
	종목_모음 = append(종목_모음, New종목("140570", "KStar 수출주"))
	종목_모음 = append(종목_모음, New종목("140580", "KStar 우량업종"))
	종목_모음 = append(종목_모음, New종목("140700", "KODEX 보험"))
	종목_모음 = append(종목_모음, New종목("140710", "KODEX 운송"))
	종목_모음 = append(종목_모음, New종목("140950", "파워 K100"))
	종목_모음 = append(종목_모음, New종목("141240", "ARIRANG K100EW"))
	종목_모음 = append(종목_모음, New종목("143460", "KINDEX 밸류대형"))
	종목_모음 = append(종목_모음, New종목("143850", "TIGER S&P500선물(H)"))
	종목_모음 = append(종목_모음, New종목("143860", "TIGER 헬스케어"))
	종목_모음 = append(종목_모음, New종목("144600", "KODEX 은선물(H)"))
	종목_모음 = append(종목_모음, New종목("145670", "KINDEX 인버스"))
	종목_모음 = append(종목_모음, New종목("145850", "TREX 펀더멘탈 200"))
	종목_모음 = append(종목_모음, New종목("147970", "TIGER 모멘텀"))
	종목_모음 = append(종목_모음, New종목("148020", "KStar 200"))
	종목_모음 = append(종목_모음, New종목("148040", "PIONEER SRI"))
	종목_모음 = append(종목_모음, New종목("148070", "KOSEF 10년 국고채"))
	종목_모음 = append(종목_모음, New종목("150460", "TIGER 중국소비테마"))
	종목_모음 = append(종목_모음, New종목("152100", "ARIRANG 200"))
	종목_모음 = append(종목_모음, New종목("152180", "TIGER 생활필수품"))
	종목_모음 = append(종목_모음, New종목("152280", "KOSEF 200 선물"))
	종목_모음 = append(종목_모음, New종목("152380", "KODEX 10년 국채선물"))
	종목_모음 = append(종목_모음, New종목("152500", "KINDEX 레버리지"))
	종목_모음 = append(종목_모음, New종목("152870", "파워 K200"))
	종목_모음 = append(종목_모음, New종목("153130", "KODEX 단기채권"))
	종목_모음 = append(종목_모음, New종목("153270", "iKon 100"))
	종목_모음 = append(종목_모음, New종목("156080", "KODEX MSCI KOREA"))
	종목_모음 = append(종목_모음, New종목("157450", "TIGER 유동자금"))
	종목_모음 = append(종목_모음, New종목("157490", "TIGER 소프트웨어"))
	종목_모음 = append(종목_모음, New종목("157500", "TIGER 증권"))
	종목_모음 = append(종목_모음, New종목("157510", "TIGER 자동차"))
	종목_모음 = append(종목_모음, New종목("157520", "TIGER 화학"))
	종목_모음 = append(종목_모음, New종목("157650", "KStar 5대그룹주 장기"))
	종목_모음 = append(종목_모음, New종목("159800", "마이티 K100"))
	종목_모음 = append(종목_모음, New종목("160580", "TIGER 구리실물"))
	종목_모음 = append(종목_모음, New종목("161490", "ARIRANG 방어주"))
	종목_모음 = append(종목_모음, New종목("161500", "ARIRANG 주도주"))
	종목_모음 = append(종목_모음, New종목("161510", "ARIRANG 고배당주"))
	종목_모음 = append(종목_모음, New종목("166400", "TIGER 커버드C200"))
	종목_모음 = append(종목_모음, New종목("167860", "KOSEF 10년 국고채 레버리지"))
	종목_모음 = append(종목_모음, New종목("168300", "KTOP50"))
	종목_모음 = append(종목_모음, New종목("168580", "KINDEX 중국본토CSI30"))
	종목_모음 = append(종목_모음, New종목("169950", "KODEX 중국본토 A50"))
	종목_모음 = append(종목_모음, New종목("170350", "TIGER 베타플러스"))
	종목_모음 = append(종목_모음, New종목("174350", "TIGER 로우볼"))
	종목_모음 = append(종목_모음, New종목("174360", "KStar 중국본토 대형"))
	종목_모음 = append(종목_모음, New종목("176710", "파워 국고채"))
	종목_모음 = append(종목_모음, New종목("176950", "KODEX 인버스국채선물"))
	종목_모음 = append(종목_모음, New종목("181450", "KINDEX 선진국 하이일드"))
	종목_모음 = append(종목_모음, New종목("181480", "KINDEX 미국 리츠 부동산"))
	종목_모음 = append(종목_모음, New종목("182480", "TIGER US리츠(합성 H)"))
	종목_모음 = append(종목_모음, New종목("182490", "TIGER 단기선진하이일드"))
	종목_모음 = append(종목_모음, New종목("183700", "KStar 채권혼합"))
	종목_모음 = append(종목_모음, New종목("183710", "KStar 주식혼합"))
	종목_모음 = append(종목_모음, New종목("185680", "KODEX 미국바이오(합성)"))
	종목_모음 = append(종목_모음, New종목("189400", "ARIRANG AC 월드(합성)"))
	종목_모음 = append(종목_모음, New종목("190150", "ARIRANG 바벨 채권"))
	종목_모음 = append(종목_모음, New종목("190160", "ARIRANG 단기유동성"))
	종목_모음 = append(종목_모음, New종목("190620", "KINDEX 단기자금"))
	종목_모음 = append(종목_모음, New종목("192090", "TIGER 차이나A300"))
	종목_모음 = append(종목_모음, New종목("192720", "파워고배당저변동성"))
	종목_모음 = append(종목_모음, New종목("195920", "TIGER 일본(합성 H)"))
	종목_모음 = append(종목_모음, New종목("195930", "TIGER 유로스탁스50(합성 H)"))
	종목_모음 = append(종목_모음, New종목("195970", "ARIRANG 선진국(합성 H)"))
	종목_모음 = append(종목_모음, New종목("195980", "ARIRANG 신흥국(합성 H)"))
	종목_모음 = append(종목_모음, New종목("196030", "KINDEX 일본레버리지(H)"))
	종목_모음 = append(종목_모음, New종목("196220", "KStar 일본레버리지(H)"))
	종목_모음 = append(종목_모음, New종목("196230", "KStar 단기통안채"))
	종목_모음 = append(종목_모음, New종목("200020", "KODEX 미국IT(합성)"))
	종목_모음 = append(종목_모음, New종목("200030", "KODEX 미국산업재(합성)"))
	종목_모음 = append(종목_모음, New종목("200040", "KODEX 미국금융(합성)"))
	종목_모음 = append(종목_모음, New종목("200050", "KODEX MSCI독일(합성)"))
	종목_모음 = append(종목_모음, New종목("200250", "KOSEF 인디아(합성)"))
	종목_모음 = append(종목_모음, New종목("203780", "TIGER 나스닥 바이오"))
	종목_모음 = append(종목_모음, New종목("204420", "ARIRANG 차이나H 레버리지"))
	종목_모음 = append(종목_모음, New종목("204450", "KODEX China H 레버리지"))
	종목_모음 = append(종목_모음, New종목("204480", "TIGER 차이나A레버리지"))
	종목_모음 = append(종목_모음, New종목("205720", "KINDEX 일본인버스(합성 H)"))
	종목_모음 = append(종목_모음, New종목("208470", "SMART MSCI선진국(합성 H)"))
	종목_모음 = append(종목_모음, New종목("210780", "TIGER 코스피 고배당"))
	종목_모음 = append(종목_모음, New종목("211210", "마이티 코스피 고배당"))
	종목_모음 = append(종목_모음, New종목("211260", "KINDEX 배당성장"))
	종목_모음 = append(종목_모음, New종목("211560", "TIGER 배당성장"))
	종목_모음 = append(종목_모음, New종목("211900", "KODEX 배당성장"))
	종목_모음 = append(종목_모음, New종목("213610", "KODEX 삼성그룹밸류"))
	종목_모음 = append(종목_모음, New종목("213630", "ARIRANG 미국고배당주(합성 H)"))
	종목_모음 = append(종목_모음, New종목("214980", "KODEX 단기채권 PLUS"))
	종목_모음 = append(종목_모음, New종목("215620", "흥국 S&P 로우볼"))
	종목_모음 = append(종목_모음, New종목("217770", "TIGER 원유선물 인버스, TIGER 원유 인버스"))
	종목_모음 = append(종목_모음, New종목("217780", "TIGER 차이나A 인버스"))
	종목_모음 = append(종목_모음, New종목("217790", "TIGER 가격조정"))
	종목_모음 = append(종목_모음, New종목("218420", "KODEX 미국에너지(합성)"))
	종목_모음 = append(종목_모음, New종목("219390", "KStar 미국원유생산기업(합성 H)"))
	종목_모음 = append(종목_모음, New종목("219480", "KODEX S&P500선물(H)"))
	종목_모음 = append(종목_모음, New종목("219900", "KINDEX 중국본토레버리지"))
	종목_모음 = append(종목_모음, New종목("220130", "SMART 중국본토 중소형 CSI500(합성 H)"))
	종목_모음 = append(종목_모음, New종목("222170", "ARIRANG S&P 배당성장"))
	종목_모음 = append(종목_모음, New종목("222180", "ARIRANG 스마트베타 V"))
	종목_모음 = append(종목_모음, New종목("222190", "ARIRANG 스마트베타 M"))
	종목_모음 = append(종목_모음, New종목("222200", "ARIRANG 스마트베타 Q"))
	종목_모음 = append(종목_모음, New종목("223190", "KODEX 200 내재가치"))
	종목_모음 = append(종목_모음, New종목("225030", "TIGER S&P500 인버스 선물"))
	종목_모음 = append(종목_모음, New종목("225040", "TIGER S&P500 레버리지"))
	종목_모음 = append(종목_모음, New종목("225050", "TIGER 유로스탁스 레버리지"))
	종목_모음 = append(종목_모음, New종목("225060", "TIGER 이머징마켓 레버리지"))
	종목_모음 = append(종목_모음, New종목("225130", "KINDEX 골드선물 레버리지"))
	종목_모음 = append(종목_모음, New종목("225800", "KOSEF 미국달러선물 레버리지(합성)"))
	종목_모음 = append(종목_모음, New종목("226380", "KINDEX 한류"))
	종목_모음 = append(종목_모음, New종목("226490", "KODEX 코스피"))
	종목_모음 = append(종목_모음, New종목("226810", "파워 단기채"))
	종목_모음 = append(종목_모음, New종목("226980", "KODEX 200 중소형"))
	종목_모음 = append(종목_모음, New종목("227540", "TIGER 200 건강관리"))
	종목_모음 = append(종목_모음, New종목("227550", "TIGER 200 산업재"))
	종목_모음 = append(종목_모음, New종목("227560", "TIGER 200 생활소비재"))
	종목_모음 = append(종목_모음, New종목("227570", "TIGER 우량가치"))
	종목_모음 = append(종목_모음, New종목("227830", "ARIRANG 코스피"))
	종목_모음 = append(종목_모음, New종목("227930", "KINDEX 코스닥 150"))
	종목_모음 = append(종목_모음, New종목("228790", "TIGER 화장품"))
	종목_모음 = append(종목_모음, New종목("228800", "TIGER 여행레저"))
	종목_모음 = append(종목_모음, New종목("228810", "TIGER 미디어컨텐츠"))
	종목_모음 = append(종목_모음, New종목("228820", "TIGER KTOP30"))
	종목_모음 = append(종목_모음, New종목("229200", "KODEX 코스닥 150"))
	종목_모음 = append(종목_모음, New종목("229720", "KODEX KTOP30"))
	return 종목_모음
}

func F임의_종목_ETF() I종목 {
	r := F임의값_생성기()
	종목_모음 := F샘플_종목_모음_ETF()

	return 종목_모음[r.Intn(len(종목_모음))]
}

func F샘플_종목_모음_주식() []I종목 {
	종목_모음 := make([]I종목, 0)
	종목_모음 = append(종목_모음, New종목("000020", "동화약품"))
	종목_모음 = append(종목_모음, New종목("000030", "우리은행"))
	종목_모음 = append(종목_모음, New종목("000040", "KR모터스"))
	종목_모음 = append(종목_모음, New종목("000050", "경방"))
	종목_모음 = append(종목_모음, New종목("000060", "메리츠화재"))

	return 종목_모음
}

func F임의_종목_주식() I종목 {
	r := F임의값_생성기()
	종목_모음 := F샘플_종목_모음_주식()

	return 종목_모음[r.Intn(len(종목_모음))]
}

func F샘플_통화단위_모음() []string {
	샘플_통화단위_모음 := make([]string, 0)
	샘플_통화단위_모음 = append(샘플_통화단위_모음, KRW)
	샘플_통화단위_모음 = append(샘플_통화단위_모음, USD)
	샘플_통화단위_모음 = append(샘플_통화단위_모음, CNY)
	샘플_통화단위_모음 = append(샘플_통화단위_모음, EUR)

	return 샘플_통화단위_모음
}

func F임의_통화단위() string {
	r := F임의값_생성기()
	통화단위_모음 := F샘플_통화단위_모음()

	return 통화단위_모음[r.Intn(len(통화단위_모음))]
}

func F임의_통화값_모음(수량 int) []I통화 {
	통화_모음 := make([]I통화, 수량)
	통화단위_모음 := F샘플_통화단위_모음()
	r := F임의값_생성기()

	통화단위 := ""
	금액 := 0.0

	for i := 0; i < 수량; i++ {
		통화단위 = 통화단위_모음[r.Intn(len(통화단위_모음))]
		금액 = math.Trunc(r.Float64()*math.Pow10(r.Intn(5))*100) / 100

		통화_모음[i] = New통화(통화단위, 금액)
	}

	return 통화_모음
}

func F임의_통화값() I통화 {
	return F임의_통화값_모음(1)[0]
}

func F임의_정수값() int {
	r := F임의값_생성기()

	if r.Intn(1) == 0 {
		return r.Int()
	} else {
		return -1 * r.Int()
	}
}
