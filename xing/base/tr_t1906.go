package xt

import (
	"bytes"
	"encoding/binary"
	"github.com/ghts/ghts/lib"
	"time"
)

// t1906 ETF LP 호가 조회 응답
type T1906_ETF_LP_호가_조회_응답 struct {
	M종목코드         string
	M시각           time.Time
	M종목명          string
	M현재가          int64
	M상한가          int64
	M하한가          int64
	M시가           int64
	M고가           int64
	M저가           int64
	M전일대비구분       T전일대비_구분
	M전일대비등락폭      int64
	M등락율          float64
	M거래량          int64
	M전일종가         int64
	LP매도_잔량_모음    []int64
	LP매수_잔량_모음    []int64
	M매도_호가_모음     []int64
	M매수_호가_모음     []int64
	M매도_잔량_모음     []int64
	M매수_잔량_모음     []int64
	M매도_직전대비수량_모음 []int64
	M매수_직전대비수량_모음 []int64
	M매도호가수량합      int64
	M매수호가수량합      int64
	M직전매도대비수량합    int64
	M직전매수대비수량합    int64
	M예상체결가격       int64
	M예상체결수량       int64
	M예상체결전일구분     T전일대비_구분
	M예상체결전일대비     int64
	M예상체결등락율      float64
	M시간외매도잔량      int64
	M시간외매수잔량      int64
	M동시호가_구분      T동시호가_구분
}

func NewT1906InBlock(질의값 *lib.S질의값_단일_종목) (g *T1906InBlock) {
	g = new(T1906InBlock)
	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)

	f속성값_초기화(g)

	return g
}

func NewT1906_ETF_LP_호가_조회_응답(b []byte) (s *T1906_ETF_LP_호가_조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT1906OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(T1906OutBlock)
	lib.F확인1(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g)) // 네트워크 전송 바이트 순서는 빅엔디언.

	s = new(T1906_ETF_LP_호가_조회_응답)
	s.M종목코드 = lib.F2문자열_공백_제거(g.Shcode)

	if 시각_문자열 := lib.F2문자열_공백_제거(g.Hotime); len(시각_문자열) <= 6 {
		s.M시각 = time.Time{}
	} else {
		s.M시각 = lib.F2일자별_시각_단순형_공백은_초기값(당일.G값(), "150405.999", 시각_문자열[:6]+"."+시각_문자열[6:])
	}

	s.M종목명 = lib.F2문자열_EUC_KR(g.Hname)
	s.M현재가 = lib.F확인2(lib.F2정수64(g.Price))
	s.M전일대비구분 = T전일대비_구분(lib.F확인2(lib.F2정수64(g.Sign)))
	s.M전일대비등락폭 = lib.F확인2(lib.F2정수64(g.Change))
	s.M등락율 = lib.F확인2(lib.F2실수_소숫점_추가(g.Diff, 2))
	s.M거래량 = lib.F확인2(lib.F2정수64(g.Volume))
	s.M전일종가 = lib.F확인2(lib.F2정수64(g.Jnilclose))
	s.M매도_호가_모음 = []int64{
		lib.F확인2(lib.F2정수64(g.Offerho1)),
		lib.F확인2(lib.F2정수64(g.Offerho2)),
		lib.F확인2(lib.F2정수64(g.Offerho3)),
		lib.F확인2(lib.F2정수64(g.Offerho4)),
		lib.F확인2(lib.F2정수64(g.Offerho5)),
		lib.F확인2(lib.F2정수64(g.Offerho6)),
		lib.F확인2(lib.F2정수64(g.Offerho7)),
		lib.F확인2(lib.F2정수64(g.Offerho8)),
		lib.F확인2(lib.F2정수64(g.Offerho9)),
		lib.F확인2(lib.F2정수64(g.Offerho10))}

	s.M매수_호가_모음 = []int64{
		lib.F확인2(lib.F2정수64(g.Bidho1)),
		lib.F확인2(lib.F2정수64(g.Bidho2)),
		lib.F확인2(lib.F2정수64(g.Bidho3)),
		lib.F확인2(lib.F2정수64(g.Bidho4)),
		lib.F확인2(lib.F2정수64(g.Bidho5)),
		lib.F확인2(lib.F2정수64(g.Bidho6)),
		lib.F확인2(lib.F2정수64(g.Bidho7)),
		lib.F확인2(lib.F2정수64(g.Bidho8)),
		lib.F확인2(lib.F2정수64(g.Bidho9)),
		lib.F확인2(lib.F2정수64(g.Bidho10))}

	s.LP매도_잔량_모음 = []int64{
		lib.F확인2(lib.F2정수64(g.Lp_offerrem1)),
		lib.F확인2(lib.F2정수64(g.Lp_offerrem2)),
		lib.F확인2(lib.F2정수64(g.Lp_offerrem3)),
		lib.F확인2(lib.F2정수64(g.Lp_offerrem4)),
		lib.F확인2(lib.F2정수64(g.Lp_offerrem5)),
		lib.F확인2(lib.F2정수64(g.Lp_offerrem6)),
		lib.F확인2(lib.F2정수64(g.Lp_offerrem7)),
		lib.F확인2(lib.F2정수64(g.Lp_offerrem8)),
		lib.F확인2(lib.F2정수64(g.Lp_offerrem9)),
		lib.F확인2(lib.F2정수64(g.Lp_offerrem10))}

	s.LP매수_잔량_모음 = []int64{
		lib.F확인2(lib.F2정수64(g.Lp_bidrem1)),
		lib.F확인2(lib.F2정수64(g.Lp_bidrem2)),
		lib.F확인2(lib.F2정수64(g.Lp_bidrem3)),
		lib.F확인2(lib.F2정수64(g.Lp_bidrem4)),
		lib.F확인2(lib.F2정수64(g.Lp_bidrem5)),
		lib.F확인2(lib.F2정수64(g.Lp_bidrem6)),
		lib.F확인2(lib.F2정수64(g.Lp_bidrem7)),
		lib.F확인2(lib.F2정수64(g.Lp_bidrem8)),
		lib.F확인2(lib.F2정수64(g.Lp_bidrem9)),
		lib.F확인2(lib.F2정수64(g.Lp_bidrem10))}

	s.M매도_잔량_모음 = []int64{
		lib.F확인2(lib.F2정수64(g.Offerrem1)),
		lib.F확인2(lib.F2정수64(g.Offerrem2)),
		lib.F확인2(lib.F2정수64(g.Offerrem3)),
		lib.F확인2(lib.F2정수64(g.Offerrem4)),
		lib.F확인2(lib.F2정수64(g.Offerrem5)),
		lib.F확인2(lib.F2정수64(g.Offerrem6)),
		lib.F확인2(lib.F2정수64(g.Offerrem7)),
		lib.F확인2(lib.F2정수64(g.Offerrem8)),
		lib.F확인2(lib.F2정수64(g.Offerrem9)),
		lib.F확인2(lib.F2정수64(g.Offerrem10))}

	s.M매수_잔량_모음 = []int64{
		lib.F확인2(lib.F2정수64(g.Bidrem1)),
		lib.F확인2(lib.F2정수64(g.Bidrem2)),
		lib.F확인2(lib.F2정수64(g.Bidrem3)),
		lib.F확인2(lib.F2정수64(g.Bidrem4)),
		lib.F확인2(lib.F2정수64(g.Bidrem5)),
		lib.F확인2(lib.F2정수64(g.Bidrem6)),
		lib.F확인2(lib.F2정수64(g.Bidrem7)),
		lib.F확인2(lib.F2정수64(g.Bidrem8)),
		lib.F확인2(lib.F2정수64(g.Bidrem9)),
		lib.F확인2(lib.F2정수64(g.Bidrem10))}

	s.M매도_직전대비수량_모음 = []int64{
		lib.F확인2(lib.F2정수64(g.Preoffercha1)),
		lib.F확인2(lib.F2정수64(g.Preoffercha2)),
		lib.F확인2(lib.F2정수64(g.Preoffercha3)),
		lib.F확인2(lib.F2정수64(g.Preoffercha4)),
		lib.F확인2(lib.F2정수64(g.Preoffercha5)),
		lib.F확인2(lib.F2정수64(g.Preoffercha6)),
		lib.F확인2(lib.F2정수64(g.Preoffercha7)),
		lib.F확인2(lib.F2정수64(g.Preoffercha8)),
		lib.F확인2(lib.F2정수64(g.Preoffercha9)),
		lib.F확인2(lib.F2정수64(g.Preoffercha10))}

	s.M매수_직전대비수량_모음 = []int64{
		lib.F확인2(lib.F2정수64(g.Prebidcha1)),
		lib.F확인2(lib.F2정수64(g.Prebidcha2)),
		lib.F확인2(lib.F2정수64(g.Prebidcha3)),
		lib.F확인2(lib.F2정수64(g.Prebidcha4)),
		lib.F확인2(lib.F2정수64(g.Prebidcha5)),
		lib.F확인2(lib.F2정수64(g.Prebidcha6)),
		lib.F확인2(lib.F2정수64(g.Prebidcha7)),
		lib.F확인2(lib.F2정수64(g.Prebidcha8)),
		lib.F확인2(lib.F2정수64(g.Prebidcha9)),
		lib.F확인2(lib.F2정수64(g.Prebidcha10))}

	s.M매도호가수량합 = lib.F확인2(lib.F2정수64(g.Offer))
	s.M매수호가수량합 = lib.F확인2(lib.F2정수64(g.Bid))
	s.M직전매도대비수량합 = lib.F확인2(lib.F2정수64(g.Preoffercha))
	s.M직전매수대비수량합 = lib.F확인2(lib.F2정수64(g.Prebidcha))
	s.M예상체결가격 = lib.F확인2(lib.F2정수64(g.Yeprice))
	s.M예상체결수량 = lib.F확인2(lib.F2정수64(g.Yevolume))
	s.M예상체결전일구분 = T전일대비_구분(lib.F확인2(lib.F2정수64(g.Yesign)))
	s.M예상체결전일대비 = lib.F확인2(lib.F2정수64(g.Yechange))
	s.M예상체결등락율 = lib.F확인2(lib.F2실수_소숫점_추가(g.Yediff, 2))
	s.M시간외매도잔량 = lib.F확인2(lib.F2정수64(g.Tmoffer))
	s.M시간외매수잔량 = lib.F확인2(lib.F2정수64(g.Tmbid))
	s.M동시호가_구분 = T동시호가_구분(lib.F확인2(lib.F2정수64(g.Ho_status)))
	s.M상한가 = lib.F확인2(lib.F2정수64(g.Uplmtprice))
	s.M하한가 = lib.F확인2(lib.F2정수64(g.Dnlmtprice))
	s.M시가 = lib.F확인2(lib.F2정수64(g.Open))
	s.M고가 = lib.F확인2(lib.F2정수64(g.High))
	s.M저가 = lib.F확인2(lib.F2정수64(g.Low))

	f속성값_초기화(g)

	return s, nil
}
