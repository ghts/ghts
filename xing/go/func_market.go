package xing

import (
	"strings"

	lb "github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/trade"
	xt "github.com/ghts/ghts/xing/base"
)

// f종목코드_모음_추출 : 종목 모음에서 종목코드 슬라이스를 추출.
func f종목코드_모음_추출(모음 []*lb.S종목) []string {
	종목코드_모음 := make([]string, len(모음))

	for i, 종목 := range 모음 {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F종목코드_모음_전체() []string {
	정보 := f종목_정보()
	lb.F조건부_패닉(정보 == nil || len(정보.M전체) == 0, "xing 초기화 안 됨.")

	return f종목코드_모음_추출(정보.M전체)
}

func F종목코드_모음_KOSPI() []string {
	정보 := f종목_정보()
	lb.F조건부_패닉(정보 == nil || len(정보.M코스피) == 0, "xing 초기화 안 됨.")

	return f종목코드_모음_추출(정보.M코스피)
}

func F종목코드_모음_KOSDAQ() []string {
	정보 := f종목_정보()
	lb.F조건부_패닉(정보 == nil || len(정보.M코스닥) == 0, "xing 초기화 안 됨.")

	return f종목코드_모음_추출(정보.M코스닥)
}

func F종목코드_모음_ETF() []string {
	정보 := f종목_정보()
	lb.F조건부_패닉(정보 == nil || len(정보.ETF) == 0, "xing 초기화 안 됨.")

	return f종목코드_모음_추출(정보.ETF)
}

func F종목코드_모음_ETN() []string {
	정보 := f종목_정보()
	lb.F조건부_패닉(정보 == nil || len(정보.ETN) == 0, "xing 초기화 안 됨.")

	return f종목코드_모음_추출(정보.ETN)
}

func F질의값_종목코드_검사(질의값_원본 lb.I질의값) (에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	switch 질의값 := 질의값_원본.(type) {
	case lb.I종목코드:
		lb.F조건부_패닉(!F종목코드_존재함(질의값.G종목코드()),
			"존재하지 않는 종목코드 : '%v'", 질의값.G종목코드())
	case lb.I종목코드_모음:
		종목코드_모음 := 질의값.G종목코드_모음()

		for _, 종목코드 := range 종목코드_모음 {
			lb.F조건부_패닉(!F종목코드_존재함(종목코드), "존재하지 않는 종목코드 : '%v'", 종목코드)
		}
	}

	return nil
}

func F종목코드_존재함(종목코드 string) bool {
	종목, 에러 := F종목by코드(종목코드)

	return 에러 == nil && 종목 != nil
}

func F종목코드_검사(종목코드 string) error {
	_, 에러 := F종목by코드(종목코드)

	return 에러
}

// F종목_정보_설정 : 종목 정보를 (재)조회하여 스냅샷으로 원자 교체한다.
// 실패 시 기존 스냅샷은 유지된다.
func F종목_정보_설정() (에러 error) {
	종목정보_초기화_잠금.Lock()
	defer 종목정보_초기화_잠금.Unlock()

	return f종목_정보_설정_실행()
}

// f종목_정보_설정_실행 : TR 조회 + 스냅샷 Store.
// 반드시 종목정보_초기화_잠금 상태에서 호출할 것 (중복 TR 조회 방지).
func f종목_정보_설정_실행() (에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	// 당일 이미 설정된 경우 TR 재조회 생략.
	if s := 종목정보_저장소.Load(); s != nil && len(s.M전체) > 0 && s.M설정일.Equal(lb.F금일()) {
		return nil
	}

	종목_정보_모음 := lb.F확인2(TrT8436_주식종목_조회(lb.P시장구분_전체))

	종목정보_저장소.Store(new종목_정보_저장소(종목_정보_모음))

	return nil
}

// new종목_정보_저장소 : 전역 상태와 잠금에 접근하지 않는 순수 생성자 함수이다.
func new종목_정보_저장소(종목_정보_모음 []*xt.T8436_현물_종목조회_응답_반복값) *s종목_정보_저장소 {
	저장소 := &s종목_정보_저장소{
		M설정일:   lb.F금일(),
		M전체:    make([]*lb.S종목, 0, len(종목_정보_모음)),
		M맵_전체:  make(map[string]*lb.S종목, len(종목_정보_모음)),
		M코스피:   make([]*lb.S종목, 0),
		M맵_코스피: make(map[string]*lb.S종목),
		M코스닥:   make([]*lb.S종목, 0),
		M맵_코스닥: make(map[string]*lb.S종목),
		ETF:    make([]*lb.S종목, 0),
		M맵_ETF: make(map[string]*lb.S종목),
		ETN:    make([]*lb.S종목, 0),
		M맵_ETN: make(map[string]*lb.S종목),
		M특수_맵:  make(map[string]*lb.S종목),
		M기준가_맵: make(map[string]int64, len(종목_정보_모음)),
		M하한가_맵: make(map[string]int64, len(종목_정보_모음)),
	}

	for _, s := range 종목_정보_모음 {
		종목 := lb.New종목with가격정보(s.M종목코드, s.M종목명, s.M시장구분, s.M전일가, s.M상한가, s.M하한가, s.M기준가)

		저장소.M전체 = append(저장소.M전체, 종목)
		저장소.M맵_전체[종목.G코드()] = 종목
		저장소.M기준가_맵[s.M종목코드] = s.M기준가
		저장소.M하한가_맵[s.M종목코드] = s.M하한가

		switch s.M시장구분 {
		case lb.P시장구분_코스피:
			저장소.M코스피 = append(저장소.M코스피, 종목)
			저장소.M맵_코스피[종목.G코드()] = 종목
		case lb.P시장구분_코스닥:
			저장소.M코스닥 = append(저장소.M코스닥, 종목)
			저장소.M맵_코스닥[종목.G코드()] = 종목
		case lb.P시장구분_ETF:
			저장소.ETF = append(저장소.ETF, 종목)
			저장소.M맵_ETF[종목.G코드()] = 종목
		case lb.P시장구분_ETN:
			저장소.ETN = append(저장소.ETN, 종목)
			저장소.M맵_ETN[종목.G코드()] = 종목
		default:
			// PASS. 코드 검사 통과를 위해서 default문 추가.
		}

		switch s.M증권그룹 {
		case xt.P증권그룹_예탁증서,
			xt.P증권그룹_증권투자회사_뮤추얼펀드,
			xt.P증권그룹_Reits종목,
			xt.P증권그룹_선박투자회사,
			xt.P증권그룹_인프라투융자회사,
			xt.P증권그룹_해외ETF,
			xt.P증권그룹_해외원주:
			저장소.M특수_맵[s.M종목코드] = 종목
		}
	}

	return 저장소
}

// f종목_정보 : 현재 스냅샷을 원자적으로 반환. 없거나 전일 이하이면 (재)설정한다.
// 읽기 경로의 정상 흐름은 잠금 없이(원자 로드 1회) 통과한다.
func f종목_정보() *s종목_정보_저장소 {
	if s := 종목정보_저장소.Load(); s != nil && s.M설정일.Equal(lb.F금일()) {
		return s
	}

	// 지연 초기화: 여러 루틴이 동시에 들어와도 TR 조회는 1회만.
	종목정보_초기화_잠금.Lock()
	defer 종목정보_초기화_잠금.Unlock()

	if s := 종목정보_저장소.Load(); s != nil && s.M설정일.Equal(lb.F금일()) {
		return s
	}

	_ = f종목_정보_설정_실행() // 실패 시 기존(전일) 스냅샷 유지. 에러는 S예외처리에서 출력됨.

	return 종목정보_저장소.Load()
}

func F종목by코드(종목코드 string) (종목 *lb.S종목, 에러 error) {
	정보 := f종목_정보()

	if 정보 == nil || len(정보.M맵_전체) == 0 {
		return nil, lb.New에러("xing 모듈 초기화 되지 않음.")
	}

	종목코드 = trade.F종목코드_보정(종목코드)

	if strings.HasPrefix(종목코드, "B") {
		return nil, lb.New에러("%v : B로 시작하는 채권 종목입니다.", 종목코드)
	}

	종목, ok := 정보.M맵_전체[종목코드]

	if !ok {
		return nil, lb.New에러("해당 종목코드가 존재하지 않습니다. '%v'", 종목코드)
	}

	return 종목, nil
}

func F종목명by코드(종목코드 string) (종목명 string, 에러 error) {
	if 종목, 에러 := F종목by코드(종목코드); 에러 != nil {
		return "", 에러
	} else if 종목명 := 종목.G이름(); 종목명 == "" {
		return "", lb.New에러("%v : 종목명 없음", 종목코드)
	} else {
		return 종목명, nil
	}
}

// F하한가by종목코드 : 종목코드로 하한가 조회. (테스트 등 전역 맵 직접 참조 대체용)
func F하한가by종목코드(종목코드 string) (int64, bool) {
	if s := 종목정보_저장소.Load(); s != nil {
		값, ok := s.M하한가_맵[종목코드]
		return 값, ok
	}

	return 0, false
}

func F임의_종목() *lb.S종목 {
	if 정보 := f종목_정보(); 정보 != nil {
		return f임의_종목_추출(정보.M전체)
	}

	return nil
}

func F임의_종목_코스피_주식() *lb.S종목 {
	if 정보 := f종목_정보(); 정보 != nil {
		return f임의_종목_추출(정보.M코스피)
	}

	return nil
}

func F임의_종목_코스닥_주식() *lb.S종목 {
	if 정보 := f종목_정보(); 정보 != nil {
		return f임의_종목_추출(정보.M코스닥)
	}

	return nil
}

func F임의_종목_ETF() *lb.S종목 {
	if 정보 := f종목_정보(); 정보 != nil {
		return f임의_종목_추출(정보.ETF)
	}

	return nil
}

func f임의_종목_추출(종목_모음 []*lb.S종목) *lb.S종목 {
	if len(종목_모음) == 0 {
		return nil
	}

	return 종목_모음[lb.F임의_범위_이내_정수값(0, len(종목_모음)-1)].G복제본()
}

func F코스피_종목_여부(종목코드 string) bool {
	정보 := f종목_정보()
	if 정보 == nil {
		return false
	}

	_, 존재함 := 정보.M맵_코스피[trade.F종목코드_보정(종목코드)]

	return 존재함
}

func F코스닥_종목_여부(종목코드 string) bool {
	정보 := f종목_정보()
	if 정보 == nil {
		return false
	}

	_, 존재함 := 정보.M맵_코스닥[trade.F종목코드_보정(종목코드)]

	return 존재함
}

func ETF_ETN_종목_여부(종목코드 string) bool {
	정보 := f종목_정보()
	if 정보 != nil {
		if _, 존재함 := 정보.M맵_ETF[trade.F종목코드_보정(종목코드)]; 존재함 {
			return true
		}

		if _, 존재함 := 정보.M맵_ETN[trade.F종목코드_보정(종목코드)]; 존재함 {
			return true
		}
	}

	종목, 에러 := F종목by코드(종목코드)
	if 에러 != nil {
		return false
	}

	switch {
	case 종목.G시장구분() == lb.P시장구분_ETF,
		종목.G시장구분() == lb.P시장구분_ETN,
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
	return strings.HasPrefix(종목_코드, "D")
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
	if 정보 := f종목_정보(); 정보 != nil {
		if _, 존재함 := 정보.M특수_맵[종목코드]; 존재함 {
			return true
		}
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
		strings.HasSuffix(종목.G이름(), "우"),
		strings.HasSuffix(종목.G이름(), "B"),
		strings.HasSuffix(종목.G이름(), "C"),
		strings.Contains(종목.G이름(), "전환"),
		strings.HasSuffix(종목.G이름(), "1"),
		strings.HasSuffix(종목.G이름(), "2"),
		strings.HasSuffix(종목.G이름(), "3"),
		strings.HasSuffix(종목.G이름(), "4"),
		strings.HasSuffix(종목.G이름(), "5"),
		strings.HasSuffix(종목.G이름(), "6"),
		strings.HasSuffix(종목.G이름(), "7"),
		strings.HasSuffix(종목.G이름(), "8"),
		strings.HasSuffix(종목.G이름(), "9"):
		return true
	}

	return false
}

func F호가_단위by종목코드(종목코드 string) (값 int64, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M에러_실행: func() { 값 = 0 }}.S실행()

	종목 := lb.F확인2(F종목by코드(종목코드))

	return F호가_단위by종목(종목)
}

func F호가_단위by종목(종목 *lb.S종목) (값 int64, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M에러_실행: func() { 값 = 0 }}.S실행()

	// 오류 발생 예방을 위해서 상한가 기준으로 호가 단위 산출.

	if ETF_ETN_종목_여부(종목.G코드()) {
		return f호가_단위_ETF_ETN(종목.G상한가()), nil
	}

	return f호가_단위_개별_종목(종목.G상한가()), nil
}

func f호가_단위_개별_종목(기준가 int64) int64 {
	switch {
	case 기준가 < 2000:
		return 1
	case 기준가 < 5000:
		return 5
	case 기준가 < 20_000:
		return 10
	case 기준가 < 50_000:
		return 50
	case 기준가 < 200_000:
		return 100
	case 기준가 < 500_000:
		return 500
	default:
		return 1000
	}
}

func f호가_단위_ETF_ETN(기준가 int64) int64 {
	switch {
	case 기준가 < 2000:
		return 1
	default:
		return 5
	}
}

func F호가_필터(종목코드 string, 호가 int64) int64 {
	if 호가 <= 0 {
		return 0
	} else if 호가_단위, 에러 := F호가_단위by종목코드(종목코드); 에러 != nil {
		return 호가 / 호가_단위 * 호가_단위
	}

	return 호가
}

func F호가_필터by종목(종목 *lb.S종목, 호가 int64) int64 {
	if 호가 <= 0 {
		return 0
	} else if 호가_단위, 에러 := F호가_단위by종목(종목); 에러 != nil {
		return 호가 / 호가_단위 * 호가_단위
	}

	return 호가
}

func F금일_한국증시_개장() bool {
	return F당일().Equal(lb.F금일())
}
