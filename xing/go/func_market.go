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

package xing

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"time"
)

func F종목코드_모음_전체() []string {
	종목코드_모음 := make([]string, len(종목모음_전체), len(종목모음_전체))

	for i, 종목 := range 종목모음_전체 {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F종목코드_모음_KOSPI() []string {
	종목코드_모음 := make([]string, len(종목모음_코스피), len(종목모음_코스피))

	for i, 종목 := range 종목모음_코스피 {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F종목코드_모음_KOSDAQ() []string {
	종목코드_모음 := make([]string, len(종목모음_코스닥), len(종목모음_코스닥))

	for i, 종목 := range 종목모음_코스닥 {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F종목코드_모음_ETF() []string {
	종목코드_모음 := make([]string, len(종목모음_ETF), len(종목모음_ETF))

	for i, 종목 := range 종목모음_ETF {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F종목코드_모음_ETN() []string {
	종목코드_모음 := make([]string, len(종목모음_ETN), len(종목모음_ETN))

	for i, 종목 := range 종목모음_ETN {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F종목코드_모음_ETF_ETN() []string {
	종목코드_모음 := make([]string, len(종목모음_ETF_ETN), len(종목모음_ETF_ETN))

	for i, 종목 := range 종목모음_ETF_ETN {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F질의값_종목코드_검사(질의값_원본 lib.I질의값) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	switch 질의값_원본.TR코드() {
	case xt.TR선물옵션_정상주문_CFOAT00100,
		xt.TR선물옵션_정정주문_CFOAT00200,
		xt.TR선물옵션_취소주문_CFOAT00300,
		xt.TR선물옵션_체결_미체결_조회_t0434: // 선물옵션은 종목코드 규칙이 현물과 다르다.
		return F선물옵션_종목코드_검사(질의값_원본.(lib.I종목코드).G종목코드())
	}

	switch 질의값 := 질의값_원본.(type) {
	case lib.I종목코드:
		lib.F조건부_패닉(!F종목코드_존재함(질의값.G종목코드()),
			"존재하지 않는 종목코드 : '%v'", 질의값.G종목코드())
	case lib.I종목코드_모음:
		종목코드_모음 := 질의값.G종목코드_모음()

		for _, 종목코드 := range 종목코드_모음 {
			lib.F조건부_패닉(!F종목코드_존재함(종목코드), "존재하지 않는 종목코드 : '%v'", 종목코드)
		}
	}

	return nil
}

func F선물옵션_종목코드_검사(종목코드 string) (에러 error) {
	switch 종목코드[:1] {
	case "1", "2", "3", "4": // 1:선물, 2:콜옵션, 3:풋옵션, 4:스프레드
		// OK
	default:
		return lib.New에러("예상하지 못한 1번째 자리값 : '%v'", 종목코드[:1])
	}

	if 정수값, 에러 := lib.F2정수(종목코드[1:3]); 에러 != nil {
		return 에러
	} else if 정수값 <= 0 || 정수값 >= 60 { // 지수(01~09), 주식(10~59), 01:코스피200 지수, 10:국민은행, 11:삼성전자 등
		return lib.New에러("예상하지 못한 2~3번째 자리값 : '%v'", 종목코드[1:3])
	}

	switch 종목코드[3:4] {
	case "6", "7", "8", "9", "0", "1", "2", "3", "4", "5",
		"A", "B", "C", "D", "E", "F", "G", "H", "J", "K",
		"L", "M", "N", "P", "Q", "R", "S", "T", "V", "W":
		// PASS	// 1996년부터 시작. 30년마다 순환. 알파벳"I/O/U"는 혼동의 위험이 있어서 제외.
	default:
		return lib.New에러("예상하지 못한 4번째 자리값 : '%v'", 종목코드[1:3])
	}

	switch 종목코드[:1] {
	case "1": // 선물
		switch 종목코드[4:5] { // 선물 결산월은 3, 6, 9, 12
		case "3", "6", "9", "12":
		default:
			return lib.New에러("예상하지 못한 5번째 자리값 : '%v'", 종목코드[4:5])
		}

		if 종목코드[5:] != "000" {
			return lib.New에러("예상하지 못한 6~8번째 자리값 : '%v'", 종목코드[5:])
		}
	case "2", "3": // 옵션
		switch 종목코드[4:5] { // 옵션 결산월은 매월
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C":
		default:
			return lib.New에러("예상하지 못한 5번째 자리값 : '%v'", 종목코드[4:5])
		}

		if _, 에러 := lib.F2정수(종목코드[5:]); 에러 != nil {
			return lib.New에러("예상하지 못한 6~8번째 자리값 : '%v'", 종목코드[5:])
		}
	case "4": // 스프레드
		switch 종목코드[4:5] { // 옵션 결산월은 매월
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C":
		default:
			return lib.New에러("예상하지 못한 5번째 자리값 : '%v'", 종목코드[4:5])
		}

		switch 종목코드[5:6] { // 원월물 만기연도
		case "6", "7", "8", "9", "0", "1", "2", "3", "4", "5",
			"A", "B", "C", "D", "E", "F", "G", "H", "J", "K",
			"L", "M", "N", "P", "Q", "R", "S", "T", "V", "W":
			// PASS	// 1996년부터 시작. 30년마다 순환. 알파벳"I/O/U"는 혼동의 위험이 있어서 제외.
		default:
			return lib.New에러("예상하지 못한 6번째 자리값 : '%v'", 종목코드[5:6])
		}

		switch 종목코드[6:7] { // 원월물 만기월
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C":
			// PASS
		default:
			return lib.New에러("예상하지 못한 7번째 자리값 : '%v'", 종목코드[6:7])
		}

		if 종목코드[7:] != "S" {
			return lib.New에러("예상하지 못한 8번째 자리값 : '%v'", 종목코드[6:7])
		}
	}

	return nil
}

func F종목코드_존재함(종목코드 string) bool {
	_, 존재함 := 종목맵_전체[종목코드]

	return 존재함
}

func F종목코드_검사(종목코드 string) error {
	if !F종목코드_존재함(종목코드) {
		return lib.New에러("존재하지 않는 종목코드 : '%s'.", 종목코드)
	}

	return nil
}

func f종목모음_설정() (에러 error) {
	종목모음_설정_잠금.Lock()
	defer 종목모음_설정_잠금.Unlock()

	defer lib.S예외처리{
		M에러: &에러,
		M함수: func() {
			종목모음_코스피 = make([]*lib.S종목, 0)
			종목모음_코스닥 = make([]*lib.S종목, 0)
			종목모음_ETF = make([]*lib.S종목, 0)
			종목모음_ETN = make([]*lib.S종목, 0)
			종목모음_ETF_ETN = make([]*lib.S종목, 0)
			종목모음_전체 = make([]*lib.S종목, 0)
			종목맵_전체 = make(map[string]*lib.S종목)
			기준가_맵 = make(map[string]int64)
			하한가_맵 = make(map[string]int64)
			종목모음_설정일 = lib.New안전한_시각(time.Time{})
		}}.S실행()

	if len(종목모음_코스피) > 0 &&
		len(종목모음_코스닥) > 0 &&
		len(종목모음_ETF) > 0 &&
		len(종목모음_ETN) > 0 &&
		len(종목모음_ETF_ETN) > 0 &&
		len(종목모음_전체) > 0 &&
		len(종목맵_전체) > 0 &&
		len(기준가_맵) > 0 &&
		len(하한가_맵) > 0 &&
		종목모음_설정일.G값().Equal(lib.F금일()) {
		return nil
	}

	종목_정보_모음, 에러 := TrT8436_주식종목_조회(lib.P시장구분_전체)
	lib.F확인(에러)

	종목모음_코스피 = make([]*lib.S종목, 0)
	종목모음_코스닥 = make([]*lib.S종목, 0)
	종목모음_ETF = make([]*lib.S종목, 0)
	종목모음_ETN = make([]*lib.S종목, 0)
	종목모음_ETF_ETN = make([]*lib.S종목, 0)
	종목모음_전체 = make([]*lib.S종목, 0)
	종목맵_전체 = make(map[string]*lib.S종목)
	기준가_맵 = make(map[string]int64)
	하한가_맵 = make(map[string]int64)

	for _, s := range 종목_정보_모음 {
		종목 := lib.New종목(s.M종목코드, s.M종목명, s.M시장구분)

		기준가_맵[s.M종목코드] = s.M기준가
		하한가_맵[s.M종목코드] = s.M하한가
		종목맵_전체[종목.G코드()] = 종목
		종목모음_전체 = append(종목모음_전체, 종목)

		switch s.M시장구분 {
		case lib.P시장구분_코스피:
			종목모음_코스피 = append(종목모음_코스피, 종목)
		case lib.P시장구분_코스닥:
			종목모음_코스닥 = append(종목모음_코스닥, 종목)
		case lib.P시장구분_ETF:
			종목모음_ETF = append(종목모음_ETF, 종목)
			종목모음_ETF_ETN = append(종목모음_ETF_ETN, 종목)
		case lib.P시장구분_ETN:
			종목모음_ETN = append(종목모음_ETN, 종목)
			종목모음_ETF_ETN = append(종목모음_ETF_ETN, 종목)
		}
	}

	종목모음_설정일 = lib.New안전한_시각(lib.F금일())

	return nil
}

func F종목by코드(종목코드 string) (종목 *lib.S종목, 에러 error) {
	if 종목, ok := 종목맵_전체[종목코드]; !ok {
		return nil, lib.New에러("해당 종목코드가 존재하지 않습니다. '%v'", 종목코드)
	} else {
		return 종목, nil
	}
}

func F임의_종목() *lib.S종목 {
	return f임의_종목_추출(종목모음_전체)
}

func F임의_종목_코스피_주식() *lib.S종목 {
	return f임의_종목_추출(종목모음_코스피)
}

func F임의_종목_코스닥_주식() *lib.S종목 {
	return f임의_종목_추출(종목모음_코스닥)
}

func F임의_종목_ETF() *lib.S종목 {
	return f임의_종목_추출(종목모음_ETF)
}

func f임의_종목_추출(종목_모음 []*lib.S종목) *lib.S종목 {
	return 종목_모음[lib.F임의_범위_이내_정수값(0, len(종목_모음))].G복제본()
}

func ETF종목_여부(종목_코드 string) bool {
	종목, 에러 := F종목by코드(종목_코드)

	return 에러 == nil && 종목.G시장구분() == lib.P시장구분_ETF
}

func ETN종목_여부(종목_코드 string) bool {
	종목, 에러 := F종목by코드(종목_코드)

	if 에러 != nil && 종목.G시장구분() == lib.P시장구분_ETN {
		return true
	}

	return false
}

func F최소_호가단위by종목코드(종목코드 string) (값 int64, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = 0 }}.S실행()

	종목 := lib.F확인(F종목by코드(종목코드)).(*lib.S종목)

	return F최소_호가단위by종목(종목)
}

func F최소_호가단위by종목(종목 *lib.S종목) (값 int64, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = 0 }}.S실행()

	return F최소_호가단위by시장구분_기준가(종목.G시장구분(), 기준가_맵[종목.G코드()])
}

func F최소_호가단위by시장구분_기준가(시장구분 lib.T시장구분, 기준가 int64) (값 int64, 에러 error) {
	switch 시장구분 {
	case lib.P시장구분_코스피:
		switch {
		case 기준가 < 1000:
			return 1, nil
		case 기준가 >= 1000 && 기준가 < 5000:
			return 5, nil
		case 기준가 >= 5000 && 기준가 < 10000:
			return 10, nil
		case 기준가 >= 10000 && 기준가 < 50000:
			return 50, nil
		case 기준가 >= 50000 && 기준가 < 100000:
			return 100, nil
		case 기준가 >= 100000 && 기준가 < 500000:
			return 500, nil
		case 기준가 >= 500000:
			return 1000, nil
		default:
			panic(lib.New에러with출력("예상하지 못한 경우. %v", 기준가))
		}
	case lib.P시장구분_코스닥:
		switch {
		case 기준가 < 1000:
			return 1, nil
		case 기준가 >= 1000 && 기준가 < 5000:
			return 5, nil
		case 기준가 >= 5000 && 기준가 < 10000:
			return 10, nil
		case 기준가 >= 10000 && 기준가 < 50000:
			return 50, nil
		case 기준가 >= 50000:
			return 100, nil
		default:
			panic(lib.New에러with출력("예상하지 못한 경우. %v", 기준가))
		}
	case lib.P시장구분_코넥스:
		switch {
		case 기준가 < 5000:
			return 5, nil
		case 기준가 >= 5000 && 기준가 < 10000:
			return 10, nil
		case 기준가 >= 10000 && 기준가 < 50000:
			return 50, nil
		case 기준가 >= 50000 && 기준가 < 100000:
			return 100, nil
		case 기준가 >= 100000 && 기준가 < 500000:
			return 500, nil
		case 기준가 >= 500000:
			return 1000, nil
		default:
			panic(lib.New에러with출력("예상하지 못한 경우. %v", 기준가))
		}
	case lib.P시장구분_ETF:
		return 5, nil
	}

	return 0, lib.New에러with출력("예상하지 못한 시장구분. %v", 시장구분)
}

func F금일_한국증시_개장() bool {
	return F당일().Equal(lib.F금일())
}