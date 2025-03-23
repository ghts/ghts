package lib

import (
	"strconv"
	"strings"
	"time"
)

//goland:noinspection GoUnusedConst
const (
	P일자_형식     = "2006-01-02"
	P시간_형식     = "2006-01-02 15:04:05.999999999 -0700 MST"
	P간략한_시간_형식 = "2006-01-02 15:04:05"
	P시분초_형식    = "15:04:05"

	P마이너스1초 = -1 * time.Second
	P10밀리초  = 10 * time.Millisecond
	P30밀리초  = 30 * time.Millisecond
	P50밀리초  = 50 * time.Millisecond
	P100밀리초 = 100 * time.Millisecond
	P300밀리초 = 300 * time.Millisecond
	P500밀리초 = 500 * time.Millisecond
	P1초     = 1 * time.Second
	P3초     = 3 * time.Second
	P4초     = 4 * time.Second
	P5초     = 5 * time.Second
	P10초    = 10 * time.Second
	P15초    = 15 * time.Second
	P20초    = 20 * time.Second
	P30초    = 30 * time.Second
	P40초    = 40 * time.Second
	P50초    = 50 * time.Second
	P1분     = 1 * time.Minute
	P2분     = 2 * time.Minute
	P3분     = 3 * time.Minute
	P4분     = 4 * time.Minute
	P5분     = 5 * time.Minute
	P6분     = 6 * time.Minute
	P7분     = 7 * time.Minute
	P8분     = 8 * time.Minute
	P9분     = 9 * time.Minute
	P10분    = 10 * time.Minute
	P15분    = 15 * time.Minute
	P20분    = 20 * time.Minute
	P30분    = 30 * time.Minute
	P40분    = 40 * time.Minute
	P50분    = 50 * time.Minute
	P1시간    = time.Hour
	P1일     = 24 * time.Hour
	P1주     = 7 * P1일
	P1년     = 365 * P1일
	P무기한    = 9999 * time.Hour

	P에러_자료형 = "error"

	// P정규식_실수는 자주 사용되는 정규식 표현
	P정규식_실수 = `[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?`

	// P긴_공백문자열는 입력 구조체 바이트 복사에 사용됨.
	P긴_공백문자열 = "         " +
		"                           " +
		"                           " +
		"                           " +
		"                           " +
		"                           " +
		"                           " +
		"                           " +
		"                           " +
		"                           "

	// P긴_0_문자열는 입력 구조체 바이트 복사에 사용됨.
	P긴_0_문자열 = "0000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000"
)

const (
	P증권사_NH   = T증권사(byte('N'))
	P증권사_Xing = T증권사(byte('X'))
)

type T증권사 byte

func (v T증권사) String() string {
	switch v {
	case P증권사_NH:
		return "NH"
	case P증권사_Xing:
		return "Xing"
	default:
		return F2문자열("예상하지 못한 값 : '%v'", int8(v))
	}
}

const (
	P양수 T부호 = 1
	P영  T부호 = 0
	P음수 T부호 = -1
)

type T부호 int8

func (v T부호) String() string {
	switch v {
	case P양수:
		return "+"
	case P영:
		return "0(zero)"
	case P음수:
		return "-"
	default:
		return F2문자열("예상하지 못한 값 : '%v'", int8(v))
	}
}

const (
	P비교_같음 T비교 = 0
	P비교_큼  T비교 = 1
	P비교_작음 T비교 = -1
	P비교_불가 T비교 = -99
)

type T비교 int8

func (v T비교) String() string {
	switch v {
	case P비교_같음:
		return "같음"
	case P비교_큼:
		return "큼"
	case P비교_작음:
		return "작음"
	case P비교_불가:
		return "비교 불가"
	default:
		return F2문자열("예상하지 못한 값 : '%v'", int8(v))
	}
}

const (
	JSON      = T변환(byte('J'))
	GOB       = T변환(byte('G'))
	Raw       = T변환(byte('R'))
	P변환형식_기본값 = GOB
)

type T변환 byte

func (t T변환) G검사() error {
	switch t {
	case JSON, GOB, Raw:
		return nil
	default:
		return New에러with출력("예상하지 못한 변환 형식. '%v'", t)
	}
}

func (t T변환) String() string {
	switch t {
	case JSON:
		return "JSON"
	case GOB:
		return "GOB"
	case Raw:
		return "Raw"
	default:
		return f포맷된_문자열("예상하지 못한 변환 형식 : '%v'", string(t))
	}
}

const (
	P메시지_질의  = "G" // GET.
	P메시지_설정  = "S" // SET. 있으면 갱신. 없으면 생성 후 갱신.
	P메시지_삭제  = "D" // DELETE
	P메시지_종료  = "Q"
	P메시지_초기화 = "I" // 주로 테스트 할 때 사용.

	P메시지_OK = "O"
	P메시지_에러 = "E"

	P메시지_생성 = "C" // CREATE
	P메시지_읽기 = "R" // READ. 질의(GET)와 중복된다고 판단되면 삭제될 수 있음.
	P메시지_갱신 = "U" // UPDATE
)

const 포트_번호_최소값 = 2000

type T주소 int

func (p T주소) G포트_번호() int {
	return int(p) + 포트_번호_최소값
}

func (p T주소) G단축값() string {
	return "127.0.0.1:" + strconv.Itoa(p.G포트_번호())
}

func (p T주소) TCP주소() string {
	return "tcp://" + p.G단축값()
}

func (p T주소) HTTPS주소(추가_인수 ...string) string {
	if len(추가_인수) > 0 {
		추가url := 추가_인수[0]

		if !strings.HasPrefix(추가url, "/") {
			추가url = "/" + 추가url
		}

		return "https://" + p.G단축값() + 추가url
	}

	return "https://" + p.G단축값()
}

func (p T주소) WS주소(추가_인수 ...string) string {
	if len(추가_인수) > 0 {
		추가url := 추가_인수[0]

		if !strings.HasPrefix(추가url, "/") {
			추가url = "/" + 추가url
		}

		return "ws://" + p.G단축값() + 추가url
	}

	return "ws://" + p.G단축값()
}

func (p T주소) String() string {
	return p.G단축값()
}

const (
	P신호 T신호 = iota
	P신호_OK
	P신호_에러
	P신호_타임아웃
	P신호_초기화
	P신호_종료
)

type T신호 uint8

func (t T신호) String() string {
	switch t {
	case P신호:
		return "신호"
	case P신호_OK:
		return "OK"
	case P신호_에러:
		return "에러"
	case P신호_타임아웃:
		return "타임아웃"
	case P신호_초기화:
		return "초기화"
	case P신호_종료:
		return "종료"
	default:
		return F2문자열("예상하지 못한 신호값 : '%v'", uint8(t))
	}
}

type T신호_32비트_모듈 uint8

const (
	P신호_DLL32_초기화 = iota
	P신호_DLL32_LOGIN
	P신호_DLL32_접속_끊김
	P신호_DLL32_종료
)

func (p T신호_32비트_모듈) String() string {
	switch p {
	case P신호_DLL32_초기화:
		return "C32 READY"
	case P신호_DLL32_LOGIN:
		return "C32 LOGIN"
	case P신호_DLL32_접속_끊김:
		return "C32 접속 끊김"
	case P신호_DLL32_종료:
		return "C32 종료"
	default:
		return F2문자열("예상하지 못한 T신호_32비트_모듈 값 : '%v'", p)
	}
}

const (
	P자료형_Int            = "int"
	P자료형_Int64          = "int64"
	P자료형_Float64        = "float64"
	P자료형_Bool           = "bool"
	P자료형_String         = "string"
	P자료형_StringArray    = "[]string"
	P자료형_Time           = "Time"
	P자료형_Error          = "error"
	P자료형_T신호            = "T신호"
	P자료형_S질의값_기본형       = "S질의값_기본형"
	P자료형_S질의값_정수        = "S질의값_정수"
	P자료형_S질의값_문자열       = "S질의값_문자열"
	P자료형_S질의값_문자열_모음    = "S질의값_문자열_모음"
	P자료형_S질의값_바이트_변환    = "S질의값_바이트_변환"
	P자료형_S질의값_바이트_변환_모음 = "S질의값_바이트_변환_모음"
	P자료형_S질의값_단일_종목     = "S질의값_단일_종목"
	P자료형_S질의값_단일종목_연속키  = "S질의값_단일종목_연속키"
	P자료형_S질의값_복수_종목     = "S질의값_복수_종목"
	P자료형_S질의값_정상_주문     = "S질의값_정상_주문"
	P자료형_S질의값_정정_주문     = "S질의값_정정_주문"
	P자료형_S질의값_취소_주문     = "S질의값_취소_주문"
	P자료형_S콜백_기본형        = "S콜백_기본형"
	P자료형_S콜백_정수값        = "S콜백_정수값"
	P자료형_S콜백_문자열        = "S콜백_문자열"
	P자료형_S콜백_TR데이터      = "S콜백_TR데이터"
	P자료형_S콜백_메시지_및_에러   = "S콜백_메시지_및_에러"
)

const (
	P시장구분_전체 T시장구분 = iota
	P시장구분_코스피
	P시장구분_코스닥
	P시장구분_코넥스
	P시장구분_ETF
	P시장구분_ETN
	P시장구분_선물옵션
	P시장구분_CME야간선물옵션
	P시장구분_EUREX야간선물옵션
)

type T시장구분 int8

func (p *T시장구분) String() string {
	switch *p {
	case P시장구분_전체:
		return "전체"
	case P시장구분_코스피:
		return "코스피"
	case P시장구분_코스닥:
		return "코스닥"
	case P시장구분_코넥스:
		return "코넥스"
	case P시장구분_ETF:
		return "ETF"
	case P시장구분_ETN:
		return "ETN"
	case P시장구분_선물옵션:
		return "선물옵션"
	case P시장구분_CME야간선물옵션:
		return "CME야간선물옵션"
	case P시장구분_EUREX야간선물옵션:
		return "EUREX야간선물옵션"
	default:
		return F2문자열("예상하지 못한 시장구분값 : '%v'", p)
	}
}

func (p *T시장구분) Parse(값 string) error {
	switch 값 {
	case "코스피":
		*p = P시장구분_코스피
	case "코스닥":
		*p = P시장구분_코스닥
	case "코넥스":
		*p = P시장구분_코넥스
	case "ETF":
		*p = P시장구분_ETF
	case "ETN":
		*p = P시장구분_ETN
	default:
		return New에러("예상하지 못한 M값. %v", 값)
	}

	return nil
}

type T매도_매수_구분 uint8

const (
	P매도_매수_전체 T매도_매수_구분 = iota
	P매도
	P매수
	P매도_정정
	P매수_정정
	P매도_취소
	P매수_취소
)

func (p T매도_매수_구분) String() string {
	switch p {
	case P매도:
		return "매도"
	case P매수:
		return "매수"
	case P매도_정정:
		return "매도 정정"
	case P매수_정정:
		return "매수 정정"
	case P매도_취소:
		return "매도 취소"
	case P매수_취소:
		return "매수 취소"
	default:
		return F2문자열("예상하지 못한 값 : %v", int(p))
	}
}

func (p T매도_매수_구분) F해석(값 interface{}) T매도_매수_구분 {
	문자열 := F2문자열_EUC_KR_공백제거(값)

	switch 문자열 {
	case P매도.String():
		return P매도
	case P매수.String():
		return P매수
	case P매도.String():
		return P매도
	case P매수.String():
		return P매수
	case strings.ReplaceAll(P매도_정정.String(), " ", ""):
		return P매도_정정
	case strings.ReplaceAll(P매수_정정.String(), " ", ""):
		return P매수_정정
	case strings.ReplaceAll(P매도_취소.String(), " ", ""):
		return P매도_취소
	case strings.ReplaceAll(P매수_취소.String(), " ", ""):
		return P매수_취소
	default:
		panic(New에러("예상하지 못한 값 : '%v'", 문자열))
	}
}

func (p T매도_매수_구분) G검사() error {
	switch p {
	case P매도_매수_전체, P매도, P매수, P매도_정정, P매수_정정, P매도_취소, P매수_취소:
		return nil
	default:
		return New에러("잘못된 매수 매도 구분값 : %v", int(p))
	}
}

type T체결_구분 uint8

const (
	P체결구분_전체 T체결_구분 = iota
	P체결구분_체결
	P체결구분_미체결
)

func (p T체결_구분) String() string {
	switch p {
	case P체결구분_전체:
		return "전체"
	case P체결구분_체결:
		return "체결"
	case P체결구분_미체결:
		return "미체결"
	default:
		panic(New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

const (
	P조정시가 T가격데이터_구분 = iota
	P조정종가
)

type T가격데이터_구분 uint8

func (v T가격데이터_구분) String() string {
	switch v {
	case P조정시가:
		return "조정시가"
	case P조정종가:
		return "조정종가"
	default:
		return F2문자열("잘못된 구분값. %v", v)
	}
}

const (
	P장_개시 T장_개시_종료 = iota
	P장_종료
)

type T장_개시_종료 uint8

func (v T장_개시_종료) String() string {
	switch v {
	case P장_개시:
		return "장_개시"
	case P장_종료:
		return "장_종료"
	default:
		return F2문자열("잘못된 구분값. %v", v)
	}
}

const (
	P장_중 T장_정보 = iota
	P장_후_시간외
	P장_전_시간외
)

type T장_정보 uint8

func (v T장_정보) String() string {
	switch v {
	case P장_중:
		return "장_중"
	case P장_후_시간외:
		return "장_후_시간외"
	case P장_전_시간외:
		return "장_전_시간외"
	default:
		return F2문자열("잘못된 장 정보 M값. %v", v)
	}
}

type T정렬_구분 uint8

const (
	P정렬구분_해당없음 T정렬_구분 = iota
	P정렬_정순
	P정렬_역순
)

func (p T정렬_구분) String() string {
	switch p {
	case P정렬구분_해당없음:
		return "해당 없음"
	case P정렬_정순:
		return "정순"
	case P정렬_역순:
		return "역순"
	default:
		return F2문자열("예상하지 못한 M값. %v", int(p))
	}
}

type T신규_정정_취소 int8

const (
	P신규 T신규_정정_취소 = iota
	P정정
	P취소
)

func (v T신규_정정_취소) String() string {
	switch v {
	case P신규:
		return "신규"
	case P정정:
		return "정정"
	case P취소:
		return "취소"
	default:
		return F2문자열("예상하지 못한 M값. %v", v)
	}
}

func (v T신규_정정_취소) G검사() error {
	switch v {
	case P신규, P정정, P취소:
		return nil
	}

	return New에러with출력("잘못된 신규 정정 취소 구분값. %v", v)
}

const (
	P주문응답_정상 T주문응답_구분 = iota
	P주문응답_정정
	P주문응답_취소
	P주문응답_거부
	P주문응답_체결
	P주문응답_IOC취소
	P주문응답_FOK취소
)

type T주문응답_구분 uint8

func (v T주문응답_구분) G검사() error {
	switch v {
	case P주문응답_정상, P주문응답_정정, P주문응답_취소,
		P주문응답_거부, P주문응답_체결,
		P주문응답_IOC취소, P주문응답_FOK취소:
		return nil
	default:
		return New에러("잘못된 주문응답 구분값. %v", uint8(v))
	}
}

func (v T주문응답_구분) String() string {
	switch v {
	case P주문응답_정상:
		return "접수"
	case P주문응답_정정:
		return "정정"
	case P주문응답_취소:
		return "취소"
	case P주문응답_거부:
		return "거부"
	case P주문응답_체결:
		return "체결"
	case P주문응답_IOC취소:
		return "IOC 취소"
	case P주문응답_FOK취소:
		return "FOC 취소"
	default:
		return F2문자열("잘못된 주문응답 구분값. %v", v)
	}
}

const (
	P호가_지정가 T호가유형 = iota
	P호가_시장가
	P호가_조건부_지정가
	P호가_최유리_지정가
	P호가_최우선_지정가
	P호가_중간가
	P호가_장전_시간외
	P호가_장후_시간외
	P호가_시간외_단일가
	P호가_해당없음
)

type T호가유형 uint8

func (v T호가유형) Xing코드() uint8 {
	switch v {
	case P호가_지정가:
		return 0
	case P호가_시장가:
		return 3
	case P호가_조건부_지정가:
		return 5
	case P호가_최유리_지정가:
		return 6
	case P호가_최우선_지정가:
		return 7
	case P호가_중간가:
		return 12
	case P호가_장전_시간외:
		return 61
	case P호가_장후_시간외:
		return 81
	case P호가_시간외_단일가:
		return 82
	default:
		return 0 // 예상하지 못한 입력값은 지정가로 대체.
	}
}

func (v T호가유형) String() string {
	switch v {
	case P호가_지정가:
		return "지정가"
	case P호가_시장가:
		return "시장가"
	case P호가_조건부_지정가:
		return "조건부 지정가"
	case P호가_최유리_지정가:
		return "최유리 지정가"
	case P호가_최우선_지정가:
		return "최우선 지정가"
	case P호가_장전_시간외:
		return "장전 시간외"
	case P호가_장후_시간외:
		return "장후 시간외"
	case P호가_시간외_단일가:
		return "시간외 단일가"
	case P호가_해당없음:
		return "해당없음"
	default:
		return F2문자열("예상하지 못한 M값. %v", uint8(v))
	}
}

func (v T호가유형) G검사() {
	switch v {
	case P호가_지정가, P호가_시장가, P호가_조건부_지정가,
		P호가_최유리_지정가, P호가_최우선_지정가,
		P호가_장전_시간외, P호가_장후_시간외,
		P호가_시간외_단일가, P호가_해당없음:
		return
	default:
		panic(New에러("잘못된 지정가 시장가 구분값. %v", uint8(v)))
	}
}

const (
	P주문조건_없음 T주문조건 = iota
	P주문조건_IOC
	P주문조건_FOK
)

type T주문조건 uint8

func (v T주문조건) String() string {
	switch v {
	case P주문조건_없음:
		return "없음"
	case P주문조건_IOC:
		return "IOC"
	case P주문조건_FOK:
		return "FOK"
	default:
		return F2문자열("잘못된 주문조건 구분값. %v", uint8(v))
	}
}

func (v T주문조건) G검사() {
	switch v {
	case P주문조건_없음, P주문조건_IOC, P주문조건_FOK:
		return
	default:
		panic(New에러("잘못된 주문 조건 구분값. %v", uint8(v)))
	}
}

const (
	P롱 = T롱숏(byte('L'))
	P숏 = T롱숏(byte('S'))
)

type T롱숏 byte

func (v T롱숏) String() string {
	switch v {
	case P롱:
		return "롱"
	case P숏:
		return "숏"
	default:
		return F2문자열("예상하지 못한 값 : '%v'", byte(v))
	}
}

type T소켓_접속방식 uint

const (
	P소켓_접속_BIND T소켓_접속방식 = iota
	P소켓_접속_CONNECT
)

func (t T소켓_접속방식) String() string {
	switch t {
	case P소켓_접속_BIND:
		return "BIND"
	case P소켓_접속_CONNECT:
		return "CONNECT"
	default:
		return "예상하지 못한 접속방식 : '" + t.String() + "'"
	}
}

type T소켓_종류 uint8

const (
	P소켓_종류_REQ = iota
	P소켓_종류_REP
	P소켓_종류_XREP
	P소켓_종류_DEALER
	P소켓_종류_ROUTER
	P소켓_종류_PUB
	P소켓_종류_SUB
	P소켓_종류_PUSH
	P소켓_종류_PULL
	P소켓_종류_PAIR
)

func (t T소켓_종류) String() string {
	switch t {
	case P소켓_종류_REQ:
		return "REQ"
	case P소켓_종류_REP:
		return "REP"
	case P소켓_종류_XREP:
		return "XREP"
	case P소켓_종류_DEALER:
		return "DEALER"
	case P소켓_종류_ROUTER:
		return "ROUTER"
	case P소켓_종류_PUB:
		return "PUB"
	case P소켓_종류_SUB:
		return "SUB"
	case P소켓_종류_PUSH:
		return "PUSH"
	case P소켓_종류_PULL:
		return "PULL"
	case P소켓_종류_PAIR:
		return "PAIR"
	default:
		return F2문자열("예상하지 못한 소켓 종류 : '%v'", uint8(t))
	}
}

type T콜백 uint8

const (
	P콜백_TR데이터 = iota
	P콜백_메시지_및_에러
	P콜백_TR완료
	P콜백_타임아웃
	P콜백_링크_데이터
	P콜백_실시간_차트_데이터
	P콜백_신호
	P콜백_소켓_테스트
)

func (p T콜백) String() string {
	switch p {
	case P콜백_TR데이터:
		return "데이터"
	case P콜백_메시지_및_에러:
		return "메시지 및 에러"
	case P콜백_TR완료:
		return "TR완료"
	case P콜백_타임아웃:
		return "타임아웃"
	case P콜백_링크_데이터:
		return "링크_데이터"
	case P콜백_실시간_차트_데이터:
		return "실시간_차트_데이터"
	case P콜백_신호:
		return "신호"
	case P콜백_소켓_테스트:
		return "소켓_테스트"
	default:
		return F2문자열("예상하지 못한 콜백값 : '%v'", p)
	}
}

const (
	P리밸런싱_주기_연 T리밸런싱_주기 = iota
	P리밸런싱_주기_반기
	P리밸런싱_주기_분기
	P리밸런싱_주기_월
	P리밸런싱_주기_주
	P리밸런싱_주기_일
)

type T리밸런싱_주기 uint8

func (t T리밸런싱_주기) String() string {
	switch t {
	case P리밸런싱_주기_연:
		return "연"
	case P리밸런싱_주기_반기:
		return "반기"
	case P리밸런싱_주기_분기:
		return "분기"
	case P리밸런싱_주기_월:
		return "월"
	case P리밸런싱_주기_주:
		return "주"
	case P리밸런싱_주기_일:
		return "일"
	default:
		return F2문자열("T리밸런싱_주기 예상하지 못한 값 : '%v'", int8(t))
	}
}
