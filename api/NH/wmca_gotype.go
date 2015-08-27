package NH

import (
	"C"
	공용 "github.com/ghts/ghts/common"
	
	"time"
	"unsafe"
)

//----------------------------------------------------------------------//
// WMCA_CONNECTED 로그인 구조체
//----------------------------------------------------------------------//
type S로그인_정보_블록 struct {
	TR구분번호 int
	M로그인_정보 S로그인_정보
}

func New로그인_정보_블록(c LoginBlock) S로그인_정보_블록 {
	return S로그인_정보_블록{
		TR구분번호: c.TrIndex,
		M로그인_정보: New로그인_정보(*(c.LoginInfo))}
}

type S로그인_정보 struct {
	M접속시각 time.Time
	M접속서버 string
	M접속ID string
	M계좌_목록 []S계좌_정보
}

func New로그인_정보(c LoginInfo) S로그인_정보 {
	시각, 에러 := time.Parse(공용.F2문자열(c.Date), "값을 확인한 후 포맷 문자열 수정할 것.")
	공용.F에러_패닉(에러)
	
	계좌_수량, 에러 := 공용.F2정수(공용.F2문자열(c.AccountCount))
	공용.F에러_패닉(에러)
	
	계좌_목록 := make([]S계좌_정보, 계좌_수량)
	for i:=0 ; i < 계좌_수량 ; i++ {
		계좌_목록[i] = New계좌_정보(c.Accountlist[i])
	} 
	
	return S로그인_정보{
		M접속시각: 시각 ,
		M접속서버: 공용.F2문자열(c.ServerName),
		M접속ID: 공용.F2문자열(c.UserID),
		M계좌_목록: 계좌_목록,
	}
}

type S계좌_정보 struct {
	M계좌_번호 string
	M계좌명 string
	M상품_코드 string
	M관리점_코드 string
	M위임_만기일 time.Time
	M일괄주문_허용계좌 bool // ('G': 허용)
	M주석 string
}

func New계좌_정보(c AccountInfo) S계좌_정보 {
	위임_만기일, 에러 := time.Parse(공용.F2문자열(c.ExpirationDate8), "값을 확인한 후 포맷 문자열 수정할 것.")
	공용.F에러_패닉(에러)
	
	일괄주문_허용계좌 := false
	if 공용.F2문자열(c.Granted) == "G" {
		일괄주문_허용계좌 = true
	}
	
	return S계좌_정보{
		M계좌_번호: 공용.F2문자열(c.AccountNo),
		M계좌명: 공용.F2문자열(c.AccountName),
		M상품_코드: 공용.F2문자열(c.Act_pdt_cdz3),
		M관리점_코드: 공용.F2문자열(c.Amn_tab_cdz4),
		M위임_만기일: 위임_만기일,
		M일괄주문_허용계좌: 일괄주문_허용계좌,
		M주석: 공용.F2문자열(c.Filler),
	}
}

//----------------------------------------------------------------------//
// WMCA 문자 message 구조체
//----------------------------------------------------------------------//
type S메시지 struct {
	M메시지_코드 string	//00000:정상, 기타:비정상(코드값은 언제든지 변경될 수 있음.)
	M메시지_내용 string
}

func New메시지(c MsgHeader) S메시지 {
	return S메시지{
		M메시지_코드: 공용.F2문자열(c.MsgCode),
		M메시지_내용: 공용.F2문자열(c.UsrMsg),
	}
}

//----------------------------------------------------------------------//
// WMCA TR 응답 구조체
//----------------------------------------------------------------------//
type S수신_블록 struct {
	TR구분번호 int
	M수신_데이터 S수신_데이터
}

func New수신_블록(c OutDataBlock) S수신_블록 {
	return S수신_블록{
		TR구분번호: int(TrIndex),
		M수신_데이터: New수신_데이터(c.DataStruct)}
}

type S수신_데이터 struct {
	M블록_이름 string
	M데이터 []byte
}

func New수신_데이터(c Received) S수신_데이터 {
	데이터 := C.GoBytes(unsafe.Pointer(*(c.DataString)), C.int(c.Length))
	// 반대는 (*C.char)(unsafe.Pointer(&b[0]))
	 
	return S수신_데이터{
		M블록_이름: 공용.F2문자열(c.BlockName),
		M데이터: 데이터}
}

//----------------------------------------------------------------------//
// 주식 현재가 조회 (c1101)
//----------------------------------------------------------------------//
type S주식_현재가_조회_질의 struct {
	M한영구분 string
	M종목코드 string
}

func New주식_현재가_조회_질의(종목코드 string) C.Tc1101InBlock {
	c := Tc1101InBlock{}
	c.Lang = []byte("K")
	copy(c.Code, 종목코드)
	
	return *((*C.Tc1101InBlock)(unsafe.Pointer(c))) 
}

type S주식_현재가_조회_기본_자료 struct {
	M종목코드 string
	M종목명 string	// 첫자리는 KOSPI200은 ‘*’, 스타지수종목은 ‘#’
	M현재가 int64
	M등락부호 string	// 0x18 :상한, 0x1E :상승, 0x20 :보합, 0x19 :하한, 0x1F :하락
	M등락폭 int64
	M등락률 float64
	M매도호가 int64
	M매수호가 int64
	M거래량 int64
	M거래비율 float64
	M유동주_회전율 float64
	M거래대금 int64
	M상한가 int64
	M고가 int64
	M시가 int64
	M시가_대비_부호 string
	M시가_대비_등락폭 int64
	M저가 int64
	M하한가 int64
	M시각 time.Time
	M매도_호가_최우선 int64
	M매도_호가_차선 int64
	M매도_호가_차차선 int64
	M매도_호가_4차선 int64
	M매도_호가_5차선 int64
	M매도_호가_6차선 int64
	M매도_호가_7차선 int64
	M매도_호가_8차선 int64
	M매도_호가_9차선 int64
	M매도_호가_10차선 int64
	M매수_호가_최우선 int64
	M매수_호가_차선 int64
	M매수_호가_차차선 int64
	M매수_호가_4차선 int64
	M매수_호가_5차선 int64
	M매수_호가_6차선 int64
	M매수_호가_7차선 int64
	M매수_호가_8차선 int64
	M매수_호가_9차선 int64
	M매수_호가_10차선 int64
	M매도_최우선_잔량 int64
	M매도_차선_잔량 int64
	M매도_차차선_잔량 int64
	M매도_4차선_잔량 int64
	M매도_5차선_잔량 int64
	M매도_6차선_잔량 int64
	M매도_7차선_잔량 int64
	M매도_8차선_잔량 int64
	M매도_9차선_잔량 int64
	M매도_10차선_잔량 int64
	M매수_최우선_잔량 int64
	M매수_차선_잔량 int64
	M매수_차차선_잔량 int64
	M매수_4차선_잔량 int64
	M매수_5차선_잔량 int64
	M매수_6차선_잔량 int64
	M매수_7차선_잔량 int64
	M매수_8차선_잔량 int64
	M매수_9차선_잔량 int64
	M매수_10차선_잔량 int64
	M매도_잔량_총합 int64
	M매수_잔량_총합 int64
	M시간외_매도_잔량 int64
	M시간외_매수_잔량 int64
	M피봇_2차_저항 int64	// 피봇가 + 전일 고가 – 전일 저가
	M피봇_1차_저항 int64	// (피봇가 * 2) – 전일 저가
	M피봇가 int64			// (전일 고가 + 전일 저가 + 전일 종가) / 3
	M피봇_1차_지지 int64	// (피봇가 * 2) – 전일 고가
	M피봇_2차_지지 int64	// 피봇가 – 전일고가 + 전일 저가
	M코스피_코스닥_구분 string	// '코스피' , '코스닥'
	M업종명 string
	M자본금_규모 string
	M결산월 string
	M시장조치1 string
	M시장조치2 string
	M시장조치3 string
	M시장조치4 string
	M시장조치5 string
	M시장조치6 string
	M전환사채_구분 string
	M액면가 int64
	M전일종가_타이틀 string
	M전일종가 int64
	M대용가 int64		// 담보가치인 듯
	M공모가 int64
	M5일_고가 int64
	M5일_저가 int64
	M20일_고가 int64
	M20일_저가 int64
	M52주_고가 int64	// 거래일만 따지면 1년이 52주이었던 것으로 기억함.
	M52주_고가일 time.Time
	M52주_저가 int64	
	M52주_저가일 time.Time
	M유동_주식수 int64
	M상장_주식수_1000주_단위 int64
	M시가_총액 int64
	M거래원_정보_수신_시간 time.Time
	M매도_거래원_1 string
	M매수_거래원_1 string
	M매도_거래량_1 int64
	M매수_거래량_1 int64
	M매도_거래원_2 string
	M매수_거래원_2 string
	M매도_거래량_2 int64
	M매수_거래량_2 int64
	M매도_거래원_3 string
	M매수_거래원_3 string
	M매도_거래량_3 int64
	M매수_거래량_3 int64
	M매도_거래원_4 string
	M매수_거래원_4 string
	M매도_거래량_4 int64
	M매수_거래량_4 int64
	M매도_거래원_5 string
	M매수_거래원_5 string
	M매도_거래량_5 int64
	M매수_거래량_5 int64
	M외국인_매도_거래량 int64
	M외국인_매수_거래량 int64
	M외국인_시간 time.Time
	M외국인_지분율 float64
	M결제일 time.Time
	M신용잔고_비율 float64
	M유상_배정_기준일 time.Time
	M무상_배정_기준일 time.Time
	M유상_배정_비율 float64
	M외국인_변동주_수량 int64
	M무상_배정_비율 float64
	M당일_자사주_신청_여부 bool
	M상장일 time.Time
	M대주주_지분율 float64
	M대주주_지분율_정보_일자 time.Time
	M네잎클로버_종목_여부 bool
	M증거금_비율 ??
	M자본금 int64
	M전체_거래원_매도_합계 int64
	M전체_거래원_매수_합계 int64
	M종목명2 string
	M우회_상장_여부 bool
	M유동주_회전율_2 float64
	M코스피_구분_2 string // 앞에 나온 '코스피/코스닥 구분'과 중복 아닌가?
	M공여율_기준일 time.Time // 공여율은 '신용거래 관련 비율'이라고 함.
	M공여율 float64	// 공여율(%)
	PER float64 // PER
	M종목별_신용_한도 ??무슨 형식? bool? int64? float64?
	M가중_평균_가격 int64
	M상장_주식수 int64
	M추가_상장_주식수 int64
	M종목_코멘트 string
	M전일_거래량 int64
	M전일대비_등락부호 string
	M전일대비_등락폭 int64
	M연중_최고가 int64	// 52주 최고가와 중복.
	M연중_최고가_일자 time.Time // 연중 최고가일
	M연중_최저가 int64
	M연중_최저가_일자 time.Time // 연중 최저가일
	M외국인_보유_주식수 int64	// 외국인 보유 주식수
	M외국인_지분_한도 float64 // % 단위
	M매매_수량_단위 int64 // 매매 수량 단위
	M대량_매매_방향 int8	// 0: 해당없음 1: 매도 2: 매수
	M대량_매매_존재 bool
}

type S주식_현재가_조회_변동_거래량_자료 struct { // 변동거래량자료[반복]
	M시간 time.Time
	M현재가 int64
	M등락부호 string
	M등락폭 int64
	M매도_호가 int64
	M매수_호가 int64
	M변동_거래량 int64
	M거래량 int64
}

type S주식_현재가_조회_종목_지표 { // 종목지표
	M동시_호가_구분 string	// 0:동시호가 아님 1:동시호가 2:동시호가연장 3:시가범위연장 4:종가범위연장 5:배분개시 6:변동성 완화장치 발동
	M예상_체결가 int64
	M예상_체결부호 string
	M예상_등락폭 int64
	M예상_등락율 float64
	M예상_체결수량 int64
	ECN정보_유무 bool	// 우리나라에는 아직 ECN이 없는 것으로 알고 있음.
	ECN전일_종가 int64
	ECN등락_부호 string
	ECN등락폭 int64
	ECN등락률 float64
	ECN체결_수량 int64
	ECN대비_예상_체결_부호 string
	ECN대비_예상_체결_등락폭 int64
	ECN대비_예상_체결_등락율 float64
}

//----------------------------------------------------------------------//
// ETF 현재가 조회 (c1151)
//----------------------------------------------------------------------//
type S_ETF_현재가_조회_질의 struct { // 기본입력
	M한영구분 string	// 기본값 'K'
	M종목코드 string
}

type S_ETF_현재가_조회_기본_자료 struct { // 종목마스타기본자료
	M종목코드 string
	M종목명 string
	M현재가 int64
	M등락부호 string
	M등락폭 int64
	M등락율 float64
	M매도_호가 int64
	M매수_호가 int64
	M거래량 int64
	M거래_비율 float64
	M유동주_회전율 float64
	M거래_대금 int64
	M상한가 int64
	M고가 int64
	M시가 int64
	M시가_대비_부호 string
	M시가_대비_등락폭 int64
	M저가 int64
	M하한가 int64
	M호가_시각 time.Time
	M매도_최우선_호가 int64
	M매도_차선_호가 int64
	M매도_차차선_호가 int64
	M매도_4차선_호가 int64
	M매도_5차선_호가 int64
	M매도_6차선_호가 int64
	M매도_7차선_호가 int64
	M매도_8차선_호가 int64
	M매도_9차선_호가 int64
	M매도_10차선_호가 int64
	M매수_최우선_호가 int64
	M매수_차선_호가 int64
	M매수_차차선_호가 int64
	M매수_4차선_호가 int64
	M매수_5차선_호가 int64
	M매수_6차선_호가 int64
	M매수_7차선_호가 int64
	M매수_8차선_호가 int64
	M매수_9차선_호가 int64
	M매수_10차선_호가 int64
	M매도_최우선_잔량 int64
	M매도_차선_잔량 int64
	M매도_차차선_잔량 int64
	M매도_4차선_잔량 int64
	M매도_5차선_잔량 int64
	M매도_6차선_잔량 int64
	M매도_7차선_잔량 int64
	M매도_8차선_잔량 int64
	M매도_9차선_잔량 int64
	M매도_10차선_잔량 int64
	M매수_최우선_잔량 int64
	M매수_차선_잔량 int64
	M매수_차차선_잔량 int64
	M매수_4차선_잔량 int64
	M매수_5차선_잔량 int64
	M매수_6차선_잔량 int64
	M매수_7차선_잔량 int64
	M매수_8차선_잔량 int64
	M매수_9차선_잔량 int64
	M매수_10차선_잔량 int64
	M총_매도_잔량 int64
	M총_매수_잔량 int64
	M시간외_매도_잔량 int64
	M시간외_매수_잔량 int64
	M피봇_2차_저항 int64
	M피봇_1차_저항 int64
	M피봇_가격 int64
	M피봇_1차_지지 int64
	M피봇_2차_지지 int64
	M코스피_코스닥_구분 string
	M업종명 string
	M자본금_규모 string
	M결산월 string
	M시장_조치_1 string
	M시장_조치_2 string
	M시장_조치_3 string
	M시장_조치_4 string
	M시장_조치_5 string
	M시장_조치_6 string
	M전환사채_구분 string
	M액면가 int64
	M전일_종가_타이틀 string	// GUI화면에 쓸 자료이니 별 필요 없을 듯...
	M전일_종가 int64
	M대용가 int64
	M공모가 int64
	M5일_고가 int64
	M5일_저가 int64
	M20일_고가 int64
	M20일_저가 int64
	M52주_고가 int64
	M52주_고가_일자 time.Time
	M52주_저가 int64
	M52주_저가_일자 time.Time
	M유동_주식수 int64
	M상장_주식수_1000주_단위 int64
	M시가_총액 int64
	M거래원_정보_수신_시점 time.Time
	M매도_거래원_1 string
	M매수_거래원_1 string
	M매도_거래원_1_거래량 int64
	M매수_거래원_1_거래량 int64
	M매도_거래원_2 string
	M매수_거래원_2 string
	M매도_거래원_2_거래량 int64
	M매수_거래원_2_거래량 int64
	M매도_거래원_3 string
	M매수_거래원_3 string
	M매도_거래원_3_거래량 int64
	M매수_거래원_3_거래량 int64
	M매도_거래원_4 string
	M매수_거래원_4 string
	M매도_거래원_4_거래량 int64
	M매수_거래원_4_거래량 int64
	M매도_거래원_5 string
	M매수_거래원_5 string
	M매도_거래원_5_거래량 int64
	M매수_거래원_5_거래량 int64
	M외국인_매도_거래량 int64
	M외국인_매수_거래량 int64
	M외국인_시간 ??
	M외국인_지분율 float64
	M결제일 time.Time
	M신용잔고_퍼센트 float64
	M유상_배정_기준일 time.Time
	M무상_배정_기준일 time.Time
	M유상_배정_비율 float64
	M무상_배정_비율 float64
	M상장일 time.Time
	M상장_주식_수량 int64
	M전체_거래원_매도_합계 int64
	M전체_거래원_매수_합계 int64
}

type S_ETF_현재가_조회_변동_거래량 struct {
	M시간 time.Time
	M현재가 int64
	M등락_부호 string
	M등락폭 int64
	M매도_호가 int64
	M매수_호가 int64
	M변동_거래량 int64
	M거래량 int64
}

type S_ETF_현재가_조회_예상_체결 struct {
	M동시_호가_구분 string
	M예상_체결가 int64
	M예상_체결_부호 string
	M예상_체결_등락폭 int64
	M예상_체결_등락율 float64
	M예상_체결_수량 int64
}

type S_ETF_현재가_조회_ETF자료 struct {
	ETF구분 string
	NAV float64
	NAV등락부호 string
	NAV등락폭 int64
	M전일NAV int64
	M괴리율 float64
	M괴리율_부호 string
	M설정단위_당_현금_배당액 int64
	M구성_종목수 int64
	M순자산_총액_억원 int64
	M추적_오차율 float64
	LP_매도_최우선_잔량 int64
	LP_매도_차선_잔량 int64
	LP_매도_차차선_잔량 int64
	LP_매도_4차선_잔량 int64
	LP_매도_5차선_잔량 int64
	LP_매도_6차선_잔량 int64
	LP_매도_7차선_잔량 int64
	LP_매도_8차선_잔량 int64
	LP_매도_9차선_잔량 int64
	LP_매도_10차선_잔량 int64
	LP_매수_최우선_잔량 int64
	LP_매수_차선_잔량 int64
	LP_매수_차차선_잔량 int64
	LP_매수_4차선_잔량 int64
	LP_매수_5차선_잔량 int64
	LP_매수_6차선_잔량 int64
	LP_매수_7차선_잔량 int64
	LP_매수_8차선_잔량 int64
	LP_매수_9차선_잔량 int64
	LP_매수_10차선_잔량 int64
	ETF_복제_방법_구분_코드 string
	ETF_상품_유형_코드 string
} Tc1151OutBlock4;

type S_ETF_현재가_조회_기반_지수_자료 {
	M지수_코드 string
	M업종_코드 string
	M지수_이름 string
	M지수 float64
	M코스피200_등락부호 string
	M코스피200_등락폭 float64
	M채권_지수 float64
	M채권_지수_등락부호 string
	M채권_지수_등락폭 float64
	M해외_지수_코드 string
	M기타_업종_코드 string
	M채권_지수_코드 string
	M채권_지수_세부_코드 string
}

//----------------------------------------------------------------------//
// 코스피 호가 잔량 (h1)
//----------------------------------------------------------------------//
type S코스피_호가_잔량_질의 struct {
	M종목_코드 string
}

type S코스피_호가_잔량_응답 struct {
	M종목_코드 string
	M시점 time.Time
	M매도_호가 int64
	M매수_호가 int64
	M매도_호가_잔량 int64
	M매수_호가_잔량 int64
	M차선_매도_호가 int64
	M차선_매수_호가 int64
	M차선_매도_호가_잔량 int64
	M차선_매수_호가_잔량 int64
	M차차선_매도_호가 int64
	M차차선_매수_호가 int64
	M차차선_매도_호가_잔량 int64
	M차차선_매수_호가_잔량 int64
	M4차선_매도_호가 int64
	M4차선_매수_호가 int64
	M4차선_매도_호가_잔량 int64
	M4차선_매수_호가_잔량 int64
	M5차선_매도_호가 int64
	M5차선_매수_호가 int64
	M5차선_매도_호가_잔량 int64
	M5차선_매수_호가_잔량 int64
	M6차선_매도_호가 int64
	M6차선_매수_호가 int64
	M6차선_매도_호가_잔량 int64
	M6차선_매수_호가_잔량 int64
	M7차선_매도_호가 int64
	M7차선_매수_호가 int64
	M7차선_매도_호가_잔량 int64
	M7차선_매수_호가_잔량 int64
	M8차선_매도_호가 int64
	M8차선_매수_호가 int64
	M8차선_매도_호가_잔량 int64
	M8차선_매수_호가_잔량 int64
	M9차선_매도_호가 int64
	M9차선_매수_호가 int64
	M9차선_매도_호가_잔량 int64
	M9차선_매수_호가_잔량 int64
	M10차선_매도_호가 int64
	M10차선_매수_호가 int64
	M10차선_매도_호가_잔량 int64
	M10차선_매수_호가_잔량 int64
	M누적_거래량 int64
}

//----------------------------------------------------------------------//
// 코스닥 호가 잔량 (k3)
//----------------------------------------------------------------------//
type S코스닥_호가_잔량_질의 struct {
	M종목_코드 string
}

type S코스닥_호가_잔량_응답 struct {
	M종목_코드 string
	M시점 time.Time
	M매도_호가 int64
	M매수_호가 int64
	M매도_호가_잔량 int64
	M매수_호가_잔량 int64
	M차선_매도_호가 int64
	M차선_매수_호가 int64
	M차선_매도_호가_잔량 int64
	M차선_매수_호가_잔량 int64
	M차차선_매도_호가 int64
	M차차선_매수_호가 int64
	M차차선_매도_호가_잔량 int64
	M차차선_매수_호가_잔량 int64
	M4차선_매도_호가 int64
	M4차선_매수_호가 int64
	M4차선_매도_호가_잔량 int64
	M4차선_매수_호가_잔량 int64
	M5차선_매도_호가 int64
	M5차선_매수_호가 int64
	M5차선_매도_호가_잔량 int64
	M5차선_매수_호가_잔량 int64
	M6차선_매도_호가 int64
	M6차선_매수_호가 int64
	M6차선_매도_호가_잔량 int64
	M6차선_매수_호가_잔량 int64
	M7차선_매도_호가 int64
	M7차선_매수_호가 int64
	M7차선_매도_호가_잔량 int64
	M7차선_매수_호가_잔량 int64
	M8차선_매도_호가 int64
	M8차선_매수_호가 int64
	M8차선_매도_호가_잔량 int64
	M8차선_매수_호가_잔량 int64
	M9차선_매도_호가 int64
	M9차선_매수_호가 int64
	M9차선_매도_호가_잔량 int64
	M9차선_매수_호가_잔량 int64
	M10차선_매도_호가 int64
	M10차선_매수_호가 int64
	M10차선_매도_호가_잔량 int64
	M10차선_매수_호가_잔량 int64
	M누적_거래량 int64
}

//----------------------------------------------------------------------//
// 코스피 시간외 호가 잔량 (h2)
//----------------------------------------------------------------------//

type S코스피_시간외_호가_잔량_질의 struct {
	M종목_코드 string
}

type S코스피_시간외_호가_잔량_응답 struct {
	M종목_코드 string
	M시점 time.Time
	M매도_호가_잔량 int64
	M매수_호가_잔량 int64
}

//----------------------------------------------------------------------//
// 코스닥 시간외 호가 잔량 (k4)
//----------------------------------------------------------------------//

type S코스닥_시간외_호가_잔량_질의 struct {
	M종목_코드 string
}

type S코스닥_시간외_호가_잔량_응답 struct {
	M종목_코드 string
	M시점 time.Time
	M매도_호가_잔량 int64
	M매수_호가_잔량 int64
}

//----------------------------------------------------------------------//
// 코스피 예상 호가 잔량 (h3)
//----------------------------------------------------------------------//

type S코스피_예상_호가_잔량_질의 struct {
	M종목_코드 string
}

type S코스피_예상_호가_잔량_응답 struct {
	M종목_코드 string
	M시점 string
	M동시_호가_구분 string
	M예상_체결가 int64
	M예상_등락_부호 string
	M예상_등락폭 int64
	M예상_등락율 float64
	M예상_체결_수량 int64
	M매도_호가 int64
	M매수_호가 int64
	M매도_호가_잔량 int64
	M매수_호가_잔량 int64
}

//----------------------------------------------------------------------//
// 코스닥 예상 호가 잔량 (k5)
//----------------------------------------------------------------------//

type S코스닥_예상_호가_잔량_질의 struct {
	M종목_코드 string
}

type S코스닥_예상_호가_잔량_응답 struct {
	M종목_코드 string
	M시점 string
	M동시_호가_구분 string
	M예상_체결가 int64
	M예상_등락_부호 string
	M예상_등락폭 int64
	M예상_등락율 float64
	M예상_체결_수량 int64
	M매도_호가 int64
	M매수_호가 int64
	M매도_호가_잔량 int64
	M매수_호가_잔량 int64
}

//----------------------------------------------------------------------//
// 코스피 체결 (j8)
//----------------------------------------------------------------------//

type S코스피_체결_질의 struct {
	M종목_코드 string
}

type S코스피_체결_응답 struct {
	M종목_코드 string
	M시각 time.Time
	M등락_부호 string
	M등락폭 int64
	M현재가 int64
	M등락율 float64
	M고가 int64
	M저가 int64
	M매도_호가 int64
	M매수_호가 int64
	M누적_거래량 int64
	M전일_거래량_대비_비율 float64
	M변동_거래량 int64
	M거래_대금 int64
	M시가 int64
	M가중_평균_가격 int64
	M장구분 string
}

//----------------------------------------------------------------------//
// 코스닥 체결 (k8)
//----------------------------------------------------------------------//

type S코스닥_체결_질의 struct {
	M종목_코드 string
}

type S코스닥_체결_응답 struct {
	M종목_코드 string
	M시각 time.Time
	M등락_부호 string
	M등락폭 int64
	M현재가 int64
	M등락율 float64
	M고가 int64
	M저가 int64
	M매도_호가 int64
	M매수_호가 int64
	M누적_거래량 int64
	M전일_거래량_대비_비율 float64
	M변동_거래량 int64
	M거래_대금 int64
	M시가 int64
	M가중_평균_가격 int64
	M장구분 string
}

//----------------------------------------------------------------------//
// 코스피 ETF NAV (j1) (예제코드가 없음. 패딩 필드가 필요한 지 추가 확인 필요함.)
//----------------------------------------------------------------------//

type S_코스피_ETF_NAV struct {
	M종목_코드 string
}

type S_코스피_ETF_NAV_응답 struct {
	M종목_코드 string
	M시점 time.Time
	M등락_부호 string
	M등락폭 int64
	NAV_현재가 int64
	NAV_시가 int64
	NAV_고가 int64
	NAV_저가 int64
	M추적_오차_부호 string
	M추적_오차 float64?? int64??
	M괴리율_부호 string
	M괴리율 float64
}

//----------------------------------------------------------------------//
// 코스닥 ETF NAV (j0) (예제코드가 없음. 패딩 필드가 필요한 지 추가 확인 필요함.)
//----------------------------------------------------------------------//

type S_코스닥_ETF_NAV struct {
	M종목_코드 string
}

type S_코스닥_ETF_NAV_응답 struct {
	M종목_코드 string
	M시점 time.Time
	M등락_부호 string
	M등락폭 int64
	NAV_현재가 int64
	NAV_시가 int64
	NAV_고가 int64
	NAV_저가 int64
	M추적_오차_부호 string
	M추적_오차 float64?? int64??
	M괴리율_부호 string
	M괴리율 float64
}

/* 코스피/코스닥 업종코드 참고표
코스피 업종명			코스닥 업종명
00 	KRX 100			01 	코스닥지수
01 	코스피지수			03 	기타서비스
02 	대형주			04 	코스닥 IT
03 	중형주			06 	제조
04 	소형주			07 	건설
05 	음식료품			08 	유통
06 	섬유의복			10 	운송
07 	종이목재			11 	금융
08 	화학				12 	통신방송서비스
09 	의약품			13 	IT S/W & SVC
10 	비금속광물			14 IT H/W
11 	철강금속			15 	음식료담배
12 	기계				16 	섬유의류
13 	전기전자			17 	종이목재
14 	의료정밀			18 	출판매체복제
15 	운수장비			19 	화학
16 	유통업			20 	제약
17 	전기가스업			21 	비금속
18 	건설업			22 	금속
19 	운수창고			23 	기계장비
20 	통신업			24 	일반전기전자
21 	금융업			25 	의료정밀기기
22 	은행				26 	운송장비부품
24 	증권				27 	기타 제조
25 	보험				28 	통신서비스
26 	서비스업			29 	방송서비스
27 	제조업			30 	인터넷
28 	코스피 200		31 	디지털컨텐츠
29 	코스피 100		32 	소프트웨어
30 	코스피 50			33 	컴퓨터서비스
32 	코스피 배당		34 	통신장비
39 	KP200 건설기계		35 	정보기기
40 	KP200 조선운송		36 	반도체
41 	KP200 철강소재		37 	IT부품
42 	KP200 에너지화학	38 	KOSDAQ 100
43 	KP200 정보기술		39 	KOSDAQ MID 300
44 	KP200 금융		40 	KOSDAQ SMALL
45 	KP200 생활소비재	43 	코스닥 스타
46 	KP200 경기소비재	44 	오락문화
47 	동일가중 KP200		45 	프리미어
48 	동일가중 KP100		46 	우량기업부
49 	동일가중 KP50		47 	벤처기업부
　	　				48 	중견기업부
　	　				49 	기술성장기업부  */

//----------------------------------------------------------------------//
// 코스피 업종 지수 (u1)
//----------------------------------------------------------------------//

type S코스피_업종_지수_질의 struct {
	M업종_코드 string
}

type S코스피_업종_지수_응답 struct {
	M업종_코드 string
	M시각 time.Time
	M지수값 float64
	M등락_부호 string
	M등락폭 float64
	M거래량 int64
	M거래_대금 int64
	M개장값 float64
	M최고값 float64
	M최고값_시각 time.Time
	M최저값 float64
	M최저값_시간 time.Time
	M지수_등락율 float64
	M거래_비중 float64
}

//----------------------------------------------------------------------//
// 코스피 업종 지수 (k1)
//----------------------------------------------------------------------//

type S코스닥_업종_지수_질의 struct {
	M업종_코드 string
}

type S코스닥_업종_지수_응답 struct {
	M업종_코드 string
	M시각 time.Time
	M지수값 float64
	M등락_부호 string
	M등락폭 float64
	M거래량 int64
	M거래_대금 int64
	M개장값 float64
	M최고값 float64
	M최고값_시각 time.Time
	M최저값 float64
	M최저값_시간 time.Time
	M지수_등락율 float64
	M거래_비중 float64
}