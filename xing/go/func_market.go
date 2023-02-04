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

package xing

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/trade"
	xt "github.com/ghts/ghts/xing/base"
	"strings"
	"time"
)

func F종목코드_모음_전체() []string {
	lib.F조건부_패닉(len(종목모음_전체) == 0, "xing 초기화 안 됨.")

	종목코드_모음 := make([]string, len(종목모음_전체), len(종목모음_전체))

	for i, 종목 := range 종목모음_전체 {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F종목코드_모음_KOSPI() []string {
	lib.F조건부_패닉(len(종목모음_전체) == 0, "xing 초기화 안 됨.")

	종목코드_모음 := make([]string, len(종목모음_코스피), len(종목모음_코스피))

	for i, 종목 := range 종목모음_코스피 {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F종목코드_모음_KOSDAQ() []string {
	lib.F조건부_패닉(len(종목모음_전체) == 0, "xing 초기화 안 됨.")

	종목코드_모음 := make([]string, len(종목모음_코스닥), len(종목모음_코스닥))

	for i, 종목 := range 종목모음_코스닥 {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F종목코드_모음_ETF() []string {
	lib.F조건부_패닉(len(종목모음_전체) == 0, "xing 초기화 안 됨.")

	종목코드_모음 := make([]string, len(종목모음_ETF), len(종목모음_ETF))

	for i, 종목 := range 종목모음_ETF {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F종목코드_모음_ETN() []string {
	lib.F조건부_패닉(len(종목모음_전체) == 0, "xing 초기화 안 됨.")

	종목코드_모음 := make([]string, len(종목모음_ETN), len(종목모음_ETN))

	for i, 종목 := range 종목모음_ETN {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F종목코드_모음_ETF_ETN() []string {
	lib.F조건부_패닉(len(종목모음_전체) == 0, "xing 초기화 안 됨.")

	종목코드_모음 := make([]string, len(종목모음_ETF_ETN), len(종목모음_ETF_ETN))

	for i, 종목 := range 종목모음_ETF_ETN {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F질의값_종목코드_검사(질의값_원본 lib.I질의값) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	//switch 질의값_원본.TR코드() {
	//case xt.TR선물옵션_정상주문_CFOAT00100,
	//	xt.TR선물옵션_정정주문_CFOAT00200,
	//	xt.TR선물옵션_취소주문_CFOAT00300,
	//	xt.TR선물옵션_체결_미체결_조회_t0434: // 선물옵션은 종목코드 규칙이 현물과 다르다.
	//	return F선물옵션_종목코드_검사(질의값_원본.(lib.I종목코드).G종목코드())
	//}

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
	if len(종목맵_전체) == 0 {
		panic(lib.New에러("xing 모듈 초기화 되지 않음."))
	}

	종목코드 = trade.F종목코드_보정(종목코드)

	_, 존재함 := 종목맵_전체[종목코드]

	return 존재함
}

func F종목코드_검사(종목코드 string) error {
	if !F종목코드_존재함(종목코드) {
		return lib.New에러("존재하지 않는 종목코드 : '%s'.", 종목코드)
	}

	return nil
}

func F종목_정보_설정() (에러 error) {
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
			특수_종목_맵 = make(map[string]*lib.S종목)
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

	종목_정보_모음 := lib.F확인2(TrT8436_주식종목_조회(lib.P시장구분_전체))

	종목모음_코스피 = make([]*lib.S종목, 0)
	종목모음_코스닥 = make([]*lib.S종목, 0)
	종목모음_ETF = make([]*lib.S종목, 0)
	종목모음_ETN = make([]*lib.S종목, 0)
	종목모음_ETF_ETN = make([]*lib.S종목, 0)
	특수_종목_맵 = make(map[string]*lib.S종목)
	종목모음_전체 = make([]*lib.S종목, 0)
	종목맵_전체 = make(map[string]*lib.S종목)
	기준가_맵 = make(map[string]int64)
	하한가_맵 = make(map[string]int64)

	for _, s := range 종목_정보_모음 {
		종목 := lib.New종목with가격정보(s.M종목코드, s.M종목명, s.M시장구분, s.M전일가, s.M상한가, s.M하한가, s.M기준가)

		기준가_맵[s.M종목코드] = s.M기준가
		하한가_맵[s.M종목코드] = s.M하한가
		종목맵_전체[종목.G코드()] = 종목
		종목모음_전체 = append(종목모음_전체, 종목)

		switch s.M시장구분 {
		case lib.P시장구분_코스피:
			종목모음_코스피 = append(종목모음_코스피, 종목)
			종목맵_코스피[종목.G코드()] = 종목
		case lib.P시장구분_코스닥:
			종목모음_코스닥 = append(종목모음_코스닥, 종목)
			종목맵_코스닥[종목.G코드()] = 종목
		case lib.P시장구분_ETF:
			종목모음_ETF = append(종목모음_ETF, 종목)
			종목모음_ETF_ETN = append(종목모음_ETF_ETN, 종목)
		case lib.P시장구분_ETN:
			종목모음_ETN = append(종목모음_ETN, 종목)
			종목모음_ETF_ETN = append(종목모음_ETF_ETN, 종목)
		}

		switch s.M증권그룹 {
		case xt.P증권그룹_예탁증서,
			xt.P증권그룹_증권투자회사_뮤추얼펀드,
			xt.P증권그룹_Reits종목,
			xt.P증권그룹_선박투자회사,
			xt.P증권그룹_인프라투융자회사,
			xt.P증권그룹_해외ETF,
			xt.P증권그룹_해외원주:
			특수_종목_맵[s.M종목코드] = 종목
		}
	}

	종목모음_설정일 = lib.New안전한_시각(lib.F금일())

	return nil
}

func F종목by코드(종목코드 string) (종목 *lib.S종목, 에러 error) {
	if len(종목맵_전체) == 0 {
		return nil, lib.New에러("Xing API가 초기화 되어 있지 않습니다.")
	} else if strings.HasPrefix(종목코드, "B") {
		return nil, lib.New에러("%v : B로 시작하는 채권 종목입니다.", 종목코드)
	}

	종목코드 = trade.F종목코드_보정(종목코드)

	if 종목, ok := 종목맵_전체[종목코드]; !ok {
		return nil, lib.New에러("해당 종목코드가 존재하지 않습니다. '%v'", 종목코드)
	} else {
		return 종목, nil
	}
}

func F종목명by코드(종목코드 string) (종목명 string, 에러 error) {
	if 종목, 에러 := F종목by코드(종목코드); 에러 != nil {
		return "", 에러
	} else if 종목명 := 종목.G이름(); 종목명 == "" {
		return "", lib.New에러("%v : 종목명 없음", 종목코드)
	} else {
		return 종목명, nil
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

func F코스피_종목_여부(종목코드 string) bool {
	_, 존재함 := 종목맵_코스피[종목코드]

	return 존재함
}

func F코스닥_종목_여부(종목코드 string) bool {
	_, 존재함 := 종목맵_코스닥[종목코드]

	return 존재함
}

func ETF_ETN_종목_여부(종목_코드 string) bool {
	종목, 에러 := F종목by코드(종목_코드)

	switch {
	case 에러 != nil:
		return false
	case 종목.G시장구분() == lib.P시장구분_ETF,
		종목.G시장구분() == lib.P시장구분_ETN,
		strings.Contains(종목.G이름(), "ETN"),
		strings.HasPrefix(종목.G이름(), "KODEX "),
		strings.HasPrefix(종목.G이름(), "TIGER "),
		strings.HasPrefix(종목.G이름(), "KOSEF "),
		strings.HasPrefix(종목.G이름(), "KINDEX "),
		strings.HasPrefix(종목.G이름(), "KBSTAR "),
		strings.HasPrefix(종목.G이름(), "HANARO "),
		strings.HasPrefix(종목.G이름(), "ARIRANG "),
		strings.HasPrefix(종목.G이름(), "SMART "),
		strings.HasPrefix(종목.G이름(), "파워 "),
		strings.HasPrefix(종목.G이름(), "TREX "),
		strings.HasPrefix(종목.G이름(), "KTOP "),
		strings.HasPrefix(종목.G이름(), "마이티 "),
		strings.HasPrefix(종목.G이름(), "FOCUS "),
		strings.HasPrefix(종목.G이름(), "QV "),
		strings.HasPrefix(종목.G이름(), "TRUE "),
		strings.HasPrefix(종목.G이름(), "ACE "),
		strings.HasPrefix(종목.G이름(), "흥국 "),
		strings.HasPrefix(종목.G이름(), "삼성 "),
		strings.HasPrefix(종목.G이름(), "KB "),
		strings.HasPrefix(종목.G이름(), "대신 "),
		strings.HasPrefix(종목.G이름(), "신한 "),
		strings.HasPrefix(종목.G이름(), "미래에셋 "),
		strings.HasPrefix(종목.G이름(), "메리츠 "):
		return true
	default:
		return false
	}
}

func F채권_종목_여부(종목_코드 string) bool {
	return strings.HasPrefix(종목_코드, "D") || (!F코스피_종목_여부(종목_코드) && !F코스닥_종목_여부(종목_코드))
}

func F레버리지_종목_여부(종목코드 string) bool {
	if !ETF_ETN_종목_여부(종목코드) {
		return false
	} else if 종목, 에러 := F종목by코드(종목코드); 에러 != nil {
		return false
	} else {
		return strings.Contains(종목.G이름(), "레버")
	}
}

func F지주회사_종목_여부(종목코드 string) bool {
	switch 종목코드 {
	case "000070", // 삼양홀딩스
		"000075", // 삼양홀딩스우
		"000140", // 하이트진로홀딩스
		"000145", // 하이트진로홀딩스우
		"000150", // 두산
		"000180", // 성창기업지주
		"000210", // DL
		"000230", // 일동홀딩스
		"000240", // 한국앤컴퍼니
		"000320", // 노루홀딩스
		"000325", // 노루홀딩스우
		"000590", // CS홀딩스
		"000640", // 동아쏘시오홀딩스
		"000670", // 영풍
		"000700", // 유수홀딩스
		"000880", // 한화
		"000885", // 한화우
		"00088K", // 한화3우B
		"001040", // CJ
		"001045", // CJ우
		"001630", // 종근당홀딩스
		"001800", // 오리온홀딩스
		"001940", // KISCO홀딩스
		"002020", // 코오롱
		"002025", // 코오롱우
		"002030", // 아세아
		"002620", // 제일파마홀딩스
		"002790", // 아모레G
		"002990", // 금호건설
		"003030", // 세아제강지주
		"003090", // 대웅
		"003300", // 한일홀딩스
		"003380", // 하림지주
		"003480", // 한진중공업홀딩스
		"003550", // LG
		"003555", // LG우
		"004150", // 한솔홀딩스
		"004360", // 세방
		"004365", // 세방우
		"004800", // 효성
		"004840", // DRB동일
		"004870", // 티웨이홀딩스
		"004990", // 롯데지주
		"005250", // 녹십자홀딩스
		"005257", // 녹십자홀딩스2우
		"005490", // POSCO홀딩스
		"005720", // 넥센
		"005725", // 넥센우
		"005740", // 크라운해태홀딩스
		"005745", // 크라운해태홀딩스우
		"005810", // 풍산홀딩스
		"005990", // 매일홀딩스
		"006120", // SK디스커버리
		"006200", // 한국전자홀딩스
		"006260", // LS
		"006840", // AK홀딩스
		"006880", // 신송홀딩스
		"007540", // 샘표
		"007700", // F&F홀딩스
		"007860", // 서연
		"008060", // 대덕
		"008930", // 한미사이언스
		"009440", // KC그린홀딩스
		"009970", // 영원무역홀딩스
		"009540", // 한국조선해양
		"010770", // 평화홀딩스
		"012030", // DB
		"012320", // 경동인베스트
		"012630", // HDC
		"013570", // 디와이
		"015360", // 예스코홀딩스
		"015860", // 일진홀딩스
		"016450", // 한세예스24홀딩스
		"016710", // 대성홀딩스
		"016880", // 웅진
		"017810", // 풀무원
		"023460", // CNH
		"024720", // 한국콜마홀딩스
		"025530", // SJM홀딩스
		"026960", // 동서
		"027410", // BGF
		"028080", // 휴맥스홀딩스
		"028260", // 삼성물산
		"030530", // 원익홀딩스
		"031980", // 피에스케이홀딩스
		"034310", // NICE
		"034730", // SK
		"03473K", // SK우
		"035080", // 인터파크홀딩스
		"035610", // 솔본
		"035810", // 이지홀딩스
		"036420", // 제이콘텐트리
		"036530", // S&T홀딩스
		"036710", // 심텍홀딩스
		"036830", // 솔브레인홀딩스
		"039020", // 이건홀딩스
		"042420", // 네오위즈홀딩스
		"044820", // 코스맥스비티아이
		"045970", // 코아시아
		"051780", // 큐로홀딩스
		"054620", // APS홀딩스
		"054800", // 아이디스홀딩스
		"055550", // 신한지주
		"057050", // 현대홈쇼핑
		"058650", // 세아홀딩스
		"060560", // 홈센타홀딩스
		"060980", // 한라홀딩스
		"071050", // 한국금융지주
		"071055", // 한국금융지주우
		"072470", // 우리산업홀딩스
		"072710", // 농심홀딩스
		"077360", // 덕산하이메탈
		"078070", // 유비쿼스홀딩스
		"078930", // GS
		"078935", // GS우
		"081660", // 휠라홀딩스
		"084110", // 휴온스글로벌
		"084690", // 대상홀딩스
		"084695", // 대상홀딩스우
		"086520", // 에코프로
		"086790", // 하나금융지주
		"088390", // 이녹스
		"092230", // KPX홀딩스
		"095570", // AJ네트웍스
		"096760", // JW홀딩스
		"100250", // 진양홀딩스
		"101060", // SBS미디어홀딩스
		"102260", // 동성케이컬
		"105560", // KB금융
		"107590", // 미원홀딩스
		"117670", // 알파홀딩스
		"121440", // 골프존뉴딘홀딩스
		"138040", // 메리츠금융지주
		"138930", // BNK금융지주
		"139130", // DGB금융지주
		"175330", // JB금융지주
		"180640", // 한진칼
		"192400", // 쿠쿠홀딩스
		"227840", // 현대코퍼레이션홀딩스
		"229640", // LS전선아시아
		"241560", // 두산밥캣
		"241590", // 화승엔터프라이즈
		"267250", // 현대중공업지주
		"307520", // TIGER 지주회사
		"316140", // 우리금융지주
		"383800", // LX홀딩스
		"900070", // 글로벌에스엠
		"900110", // 이스트아시아홀딩스
		"900140", // 엘브이엠씨홀딩스
		"900260", // 로스웰
		"900270", // 헝셩그룹
		"900280", // 골든센츄리
		"900300", // 오가님티코스메틱
		"900340": // 윙입푸드
		return true
	}

	if 종목명, 에러 := F종목명by코드(종목코드); 에러 != nil {
		return false
	} else if strings.HasSuffix(종목명, "홀딩스") ||
		strings.HasSuffix(종목명, "지주") {
		return true
	}

	return false
}

func F금융사_종목_여부(종목코드 string) bool {
	if 종목명, 에러 := F종목명by코드(종목코드); 에러 != nil {
		return false
	} else if strings.Contains(종목명, "금융") ||
		strings.HasSuffix(종목명, "은행") ||
		strings.HasSuffix(종목명, "뱅크") ||
		strings.Contains(종목명, "저축") ||
		strings.HasSuffix(종목명, "증권") ||
		strings.Contains(종목명, "카드") ||
		strings.Contains(종목명, "보험") ||
		strings.HasSuffix(종목명, "생명") ||
		strings.Contains(종목명, "손해") ||
		strings.Contains(종목명, "화재") ||
		strings.Contains(종목명, "해상") ||
		strings.Contains(종목명, "캐피탈") ||
		strings.Contains(종목명, "인베스트") ||
		strings.Contains(종목명, "투자") ||
		strings.HasPrefix(종목명, "신한") ||
		strings.HasPrefix(종목명, "하나") ||
		strings.HasPrefix(종목명, "KB") ||
		strings.HasPrefix(종목명, "BNK") ||
		strings.HasPrefix(종목명, "DGB") ||
		strings.HasPrefix(종목명, "JB") ||
		strings.HasPrefix(종목명, "메리츠") ||
		strings.Contains(종목명, "리드코프") ||
		strings.Contains(종목명, "코리안리") {
		return true
	}

	return false
}

func F특수_종목_여부(종목코드 string) bool {
	if _, 존재함 := 특수_종목_맵[종목코드]; 존재함 {
		return true
	}

	종목, 에러 := F종목by코드(종목코드)
	if 에러 != nil {
		return false
	}

	종목명 := 종목.G이름()

	switch {
	case strings.Contains(종목명, "리츠"),
		strings.Contains(종목명, "스팩"),
		strings.Contains(종목명, "SPAC"),
		strings.Contains(종목명, "하이골드"),
		strings.Contains(종목명, "1호"),
		strings.Contains(종목명, "2호"),
		strings.Contains(종목명, "3호"),
		strings.Contains(종목명, "4호"),
		strings.Contains(종목명, "5호"),
		strings.Contains(종목명, "6호"),
		strings.Contains(종목명, "7호"),
		strings.Contains(종목명, "8호"),
		strings.Contains(종목명, "9호"),
		strings.Contains(종목명, "10호"),
		strings.HasSuffix(종목.G이름(), "우") ||
			strings.HasSuffix(종목.G이름(), "B") ||
			strings.HasSuffix(종목.G이름(), "C") ||
			strings.Contains(종목.G이름(), "전환") ||
			strings.HasSuffix(종목.G이름(), "1") ||
			strings.HasSuffix(종목.G이름(), "2") ||
			strings.HasSuffix(종목.G이름(), "3") ||
			strings.HasSuffix(종목.G이름(), "4") ||
			strings.HasSuffix(종목.G이름(), "5") ||
			strings.HasSuffix(종목.G이름(), "6") ||
			strings.HasSuffix(종목.G이름(), "7") ||
			strings.HasSuffix(종목.G이름(), "8") ||
			strings.HasSuffix(종목.G이름(), "9"):
		return true

	}

	return false
}

func F최소_호가단위by종목코드(종목코드 string) (값 int64, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = 0 }}.S실행()

	종목 := lib.F확인2(F종목by코드(종목코드))

	return F최소_호가단위by종목(종목)
}

func F최소_호가단위by종목(종목 *lib.S종목) (값 int64, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = 0 }}.S실행()

	// 오류 발생 예방을 위해서 (기준가가 아닌) 상한가 기준으로 호가 단위 산출.
	return f최소_호가단위by시장구분_기준가_20230125(종목.G시장구분(), 종목.G상한가())
}

func f최소_호가단위by시장구분_기준가_20230125(시장구분 lib.T시장구분, 기준가 int64) (값 int64, 에러 error) {
	switch 시장구분 {
	case lib.P시장구분_ETF, lib.P시장구분_ETN:
		return 5, nil
	}

	switch {
	case 기준가 < 2000:
		return 1, nil
	case 기준가 < 5000:
		return 5, nil
	case 기준가 < 20_000:
		return 10, nil
	case 기준가 < 50_000:
		return 50, nil
	case 기준가 < 200_000:
		return 100, nil
	case 기준가 < 500_000:
		return 500, nil
	case 기준가 >= 500_000:
		return 1000, nil
	default:
		panic(lib.New에러with출력("예상하지 못한 경우. %v", 기준가))
	}
}

func F금일_한국증시_개장() bool {
	return F당일().Equal(lib.F금일())
}
