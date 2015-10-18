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
	// ETF 종목이 너무 많아서 다 못 넣겠음. ETF시장이 엄청나게 컸다는 것을 실감함.
	종목_모음 = append(종목_모음, New종목("152100", "ARIRANG 200"))

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
