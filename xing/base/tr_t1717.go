package xt

import (
	"bytes"
	"encoding/binary"
	"github.com/ghts/ghts/lib"
	"time"
)

type T1717_일별_누적_구분 uint8

const (
	P일별_순매수_T1717 T1717_일별_누적_구분 = iota
	P기간_누적_순매수_T1717
)

func (p T1717_일별_누적_구분) String() string {
	switch p {
	case P일별_순매수_T1717:
		return "일별 순매수"
	case P기간_누적_순매수_T1717:
		return "기간 누적 순매수"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

type T1717_종목별_매매주체_동향_질의값 struct {
	*lib.S질의값_단일_종목
	M일별_누적_구분 T1717_일별_누적_구분 // 0:일별 순매수, 1:기간 누적 순매수
	M시작일      time.Time
	M종료일      time.Time
}

type T1717_종목별_매매주체_동향_응답 struct {
	M반복값_모음 []*T1717_종목별_매매주체_동향_응답_반복값
}

type T1717_종목별_매매주체_동향_응답_반복값 struct {
	M종목코드         string
	M일자           time.Time
	M종가           float64
	M거래량          int64
	M사모펀드_순매수량    int64
	M증권_순매수량      int64
	M보험_순매수량      int64
	M투신_순매수량      int64
	M은행_순매수량      int64
	M종금_순매수량      int64
	M기금_순매수량      int64
	M기타법인_순매수량    int64
	M개인_순매수량      int64
	M등록_외국인_순매수량  int64
	M미등록_외국인_순매수량 int64
	M국가외_순매수량     int64
	M기관_순매수량      int64
	M외인계_순매수량     int64
	M기타계_순매수량     int64
	M사모펀드_단가      int64
	M증권_단가        int64
	M보험_단가        int64
	M투신_단가        int64
	M은행_단가        int64
	M종금_단가        int64
	M기금_단가        int64
	M기타법인_단가      int64
	M개인_단가        int64
	M등록_외국인_단가    int64
	M미등록_외국인_단가   int64
	M국가외_단가       int64
	M기관_단가        int64
	M외인계_단가       int64
	M기타계_단가       int64
}

func NewT1717InBlock(질의값 *T1717_종목별_매매주체_동향_질의값) (g *T1717InBlock) {
	g = new(T1717InBlock)
	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lib.F바이트_복사_문자열(g.Gubun[:], lib.F2문자열(int(질의값.M일별_누적_구분)))
	lib.F바이트_복사_문자열(g.Fromdt[:], 질의값.M시작일.Format("20060102"))
	lib.F바이트_복사_문자열(g.Todt[:], 질의값.M종료일.Format("20060102"))

	f속성값_초기화(g)

	return g
}

func NewT1717_종목별_매매주체_동향_응답(b []byte) (응답값 *T1717_종목별_매매주체_동향_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	나머지 := len(b) % SizeT1717OutBlock
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT1717OutBlock

	응답값 = new(T1717_종목별_매매주체_동향_응답)
	응답값.M반복값_모음 = make([]*T1717_종목별_매매주체_동향_응답_반복값, 수량)

	for i := 0; i < 수량; i++ {
		g := new(T1717OutBlock)
		lib.F확인1(binary.Read(버퍼, binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

		값 := new(T1717_종목별_매매주체_동향_응답_반복값)
		값.M일자 = lib.F2포맷된_일자_단순형_공백은_초기값("20060102", g.Date)
		값.M종가 = float64(lib.F확인2(lib.F2정수64_공백은_0(g.Close)))
		값.M거래량 = lib.F확인2(lib.F2정수64_공백은_0(g.Volume))
		값.M사모펀드_순매수량 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0000_vol))
		값.M증권_순매수량 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0001_vol))
		값.M보험_순매수량 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0002_vol))
		값.M투신_순매수량 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0003_vol))
		값.M은행_순매수량 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0004_vol))
		값.M종금_순매수량 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0005_vol))
		값.M기금_순매수량 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0006_vol))
		값.M기타법인_순매수량 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0007_vol))
		값.M개인_순매수량 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0008_vol))
		값.M등록_외국인_순매수량 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0009_vol))
		값.M미등록_외국인_순매수량 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0010_vol))
		값.M국가외_순매수량 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0011_vol))
		값.M기관_순매수량 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0018_vol))
		값.M외인계_순매수량 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0016_vol))
		값.M기타계_순매수량 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0017_vol))
		값.M사모펀드_단가 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0000_dan))
		값.M증권_단가 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0001_dan))
		값.M보험_단가 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0002_dan))
		값.M투신_단가 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0003_dan))
		값.M은행_단가 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0004_dan))
		값.M종금_단가 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0005_dan))
		값.M기금_단가 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0006_dan))
		값.M기타법인_단가 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0007_dan))
		값.M개인_단가 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0008_dan))
		값.M등록_외국인_단가 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0009_dan))
		값.M미등록_외국인_단가 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0010_dan))
		값.M국가외_단가 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0011_dan))
		값.M기관_단가 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0018_dan))
		값.M외인계_단가 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0016_dan))
		값.M기타계_단가 = lib.F확인2(lib.F2정수64_공백은_0(g.Tjj0017_dan))

		응답값.M반복값_모음[수량-1-i] = 값 // TR결과가 시간 역순 정렬되어 있으므로, 거꾸로 저장하여, 시간순이 되게 함.
	}

	return 응답값, nil
}
