package xing

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/trade"
	"strings"
)

func F종목by코드(종목코드 string) (종목 *lib.S종목, 에러 error) {
	if len(종목맵_전체) == 0 {
		return nil, lib.New에러("Xing API가 초기화 되어 있지 않습니다.")
	}

	종목코드 = trade.F종목코드_보정(종목코드)

	if 종목, ok := 종목맵_전체[종목코드]; !ok {
		return nil, lib.New에러("해당 종목코드가 존재하지 않습니다. '%v'", 종목코드)
	} else {
		return 종목, nil
	}
}

func F전일_종가(종목코드 string) (전일_종가 int64, 에러 error) {
	if 종목, 에러 := F종목by코드(종목코드); 에러 != nil {
		return 0, 에러
	} else {
		return 종목.G전일_종가(), nil
	}
}

func F상한가(종목코드 string) (상한가 int64, 에러 error) {
	if 종목, 에러 := F종목by코드(종목코드); 에러 != nil {
		return 0, 에러
	} else {
		return 종목.G상한가(), nil
	}
}

func F하한가(종목코드 string) (하한가 int64, 에러 error) {
	if 종목, 에러 := F종목by코드(종목코드); 에러 != nil {
		return 0, 에러
	} else {
		return 종목.G하한가(), nil
	}
}

func F기준가(종목코드 string) (기준가 int64, 에러 error) {
	if 종목, 에러 := F종목by코드(종목코드); 에러 != nil {
		return 0, 에러
	} else {
		return 종목.G기준가(), nil
	}
}

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

//func F선물옵션_종목코드_검사(종목코드 string) (에러 error) {
//	switch 종목코드[:1] {
//	case "1", "2", "3", "4": // 1:선물, 2:콜옵션, 3:풋옵션, 4:스프레드
//		// OK
//	default:
//		return lib.New에러("예상하지 못한 1번째 자리값 : '%v'", 종목코드[:1])
//	}
//
//	if 정수값, 에러 := lib.F2정수(종목코드[1:3]); 에러 != nil {
//		return 에러
//	} else if 정수값 <= 0 || 정수값 >= 60 { // 지수(01~09), 주식(10~59), 01:코스피200 지수, 10:국민은행, 11:삼성전자 등
//		return lib.New에러("예상하지 못한 2~3번째 자리값 : '%v'", 종목코드[1:3])
//	}
//
//	switch 종목코드[3:4] {
//	case "6", "7", "8", "9", "0", "1", "2", "3", "4", "5",
//		"A", "B", "C", "D", "E", "F", "G", "H", "J", "K",
//		"L", "M", "N", "P", "Q", "R", "S", "T", "V", "W":
//		// PASS	// 1996년부터 시작. 30년마다 순환. 알파벳"I/O/U"는 혼동의 위험이 있어서 제외.
//	default:
//		return lib.New에러("예상하지 못한 4번째 자리값 : '%v'", 종목코드[1:3])
//	}
//
//	switch 종목코드[:1] {
//	case "1": // 선물
//		switch 종목코드[4:5] { // 선물 결산월은 3, 6, 9, 12
//		case "3", "6", "9", "12":
//		default:
//			return lib.New에러("예상하지 못한 5번째 자리값 : '%v'", 종목코드[4:5])
//		}
//
//		if 종목코드[5:] != "000" {
//			return lib.New에러("예상하지 못한 6~8번째 자리값 : '%v'", 종목코드[5:])
//		}
//	case "2", "3": // 옵션
//		switch 종목코드[4:5] { // 옵션 결산월은 매월
//		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C":
//		default:
//			return lib.New에러("예상하지 못한 5번째 자리값 : '%v'", 종목코드[4:5])
//		}
//
//		if _, 에러 := lib.F2정수(종목코드[5:]); 에러 != nil {
//			return lib.New에러("예상하지 못한 6~8번째 자리값 : '%v'", 종목코드[5:])
//		}
//	case "4": // 스프레드
//		switch 종목코드[4:5] { // 옵션 결산월은 매월
//		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C":
//		default:
//			return lib.New에러("예상하지 못한 5번째 자리값 : '%v'", 종목코드[4:5])
//		}
//
//		switch 종목코드[5:6] { // 원월물 만기연도
//		case "6", "7", "8", "9", "0", "1", "2", "3", "4", "5",
//			"A", "B", "C", "D", "E", "F", "G", "H", "J", "K",
//			"L", "M", "N", "P", "Q", "R", "S", "T", "V", "W":
//			// PASS	// 1996년부터 시작. 30년마다 순환. 알파벳"I/O/U"는 혼동의 위험이 있어서 제외.
//		default:
//			return lib.New에러("예상하지 못한 6번째 자리값 : '%v'", 종목코드[5:6])
//		}
//
//		switch 종목코드[6:7] { // 원월물 만기월
//		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C":
//			// PASS
//		default:
//			return lib.New에러("예상하지 못한 7번째 자리값 : '%v'", 종목코드[6:7])
//		}
//
//		if 종목코드[7:] != "S" {
//			return lib.New에러("예상하지 못한 8번째 자리값 : '%v'", 종목코드[6:7])
//		}
//	}
//
//	return nil
//}

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
		strings.HasPrefix(종목.G이름(), "흥국 "),
		strings.HasPrefix(종목.G이름(), "QV "),
		strings.HasPrefix(종목.G이름(), "삼성 "),
		strings.HasPrefix(종목.G이름(), "TRUE "),
		strings.HasPrefix(종목.G이름(), "KB "),
		strings.HasPrefix(종목.G이름(), "대신 "),
		strings.HasPrefix(종목.G이름(), "신한 "),
		strings.HasPrefix(종목.G이름(), "미래에셋 "):
		return true
	default:
		return false
	}
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
		"000180", // 성창기업지주
		"000230", // 일동홀딩스
		"000240", // 한국테크놀로지그룹
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
		"004840", // DRB동일
		"004870", // 티웨이홀딩스
		"004990", // 롯데지주
		"005250", // 녹십자홀딩스
		"005257", // 녹십자홀딩스2우
		"005720", // 넥센
		"005725", // 넥센우
		"005740", // 크라운해태홀딩스
		"005745", // 크라운해태홀딩스우
		"005990", // 매일홀딩스
		"006120", // SK디스커버리
		"006200", // 한국전자홀딩스
		"006260", // LS
		"006840", // AK홀딩스
		"006880", // 신송홀딩스
		"008930", // 한미사이언스
		"009440", // KC그린홀딩스
		"009970", // 영원무역홀딩스
		"010770", // 평화홀딩스
		"012630", // HDC
		"015360", // 예스코홀딩스
		"015860", // 일진홀딩스
		"016450", // 한세예스24홀딩스
		"016710", // 대성홀딩스
		"017810", // 풀무원
		"024720", // 한국콜마홀딩스
		"025530", // SJM홀딩스
		"026960", // 동서
		"027410", // BGF
		"028080", // 휴맥스홀딩스
		"030530", // 원익홀딩스
		"031980", // 피에스케이홀딩스
		"034310", // NICE
		"034730", // SK
		"03473K", // SK우
		"035080", // 인터파크홀딩스
		"036420", // 제이콘텐트리
		"036530", // S&T홀딩스
		"036710", // 심텍홀딩스
		"039020", // 이건홀딩스
		"042420", // 네오위즈홀딩스
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
		"078070", // 유비쿼스홀딩스
		"078930", // GS
		"078935", // GS우
		"081660", // 휠라홀딩스
		"084690", // 대상홀딩스
		"084695", // 대상홀딩스우
		"086790", // 하나금융지주
		"092230", // KPX홀딩스
		"096760", // JW홀딩스
		"100250", // 진양홀딩스
		"101060", // SBS미디어홀딩스
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
		"267250", // 현대중공업지주
		"307520", // TIGER 지주회사
		"316140", // 우리금융지주
		"900110", // 이스트아시아홀딩스
		"900140": // 엘브이엠씨홀딩스
		return true
	}

	return false
}

func F특수_종목_여부(종목코드 string) bool {
	종목, 에러 := F종목by코드(종목코드)
	if 에러 != nil {
		return false
	}

	종목명 := 종목.G이름()

	switch {
	case strings.Contains(종목명, "스팩") &&
		strings.Contains(종목명, "호"):
		return true
	case strings.HasPrefix(종목명, "하이골드") &&
		strings.Contains(종목명, "호"):
		return true
	case strings.Contains(종목명, "리츠"):
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

	return F최소_호가단위by시장구분_기준가(종목.G시장구분(), 종목맵_전체[종목.G코드()].G기준가())
}

func F최소_호가단위by시장구분_기준가(시장구분 lib.T시장구분, 기준가 int64) (값 int64, 에러 error) {
	switch 시장구분 {
	case lib.P시장구분_ETF, lib.P시장구분_ETN:
		return 5, nil
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
	}

	return 0, lib.New에러with출력("예상하지 못한 시장구분. %v", 시장구분)
}

func F금일_한국증시_개장() bool {
	return F당일().Equal(lib.F금일())
}
