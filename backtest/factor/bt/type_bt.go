package bt

import (
	"github.com/ghts/ghts/lib"
)

type I전략_식별_정보 interface {
	G전략명() string
	G계좌번호() string
	G전략_식별_문자열() string
}

type I전략_인수 interface {
	I전략_식별_정보
	G리밸런싱_주기() lib.T리밸런싱_주기
	G종목_수량() int
	G복합_등급_계산_함수() func(*mf.S종목별_데이터_정렬_도우미)
	G버퍼_퍼센트() float64 // 버퍼가 '0%'이면 버퍼룰 적용 안 함.
	G데이터_처리기() I데이터_처리기
	S데이터_처리기(I데이터_처리기)
	G포트폴리오() I포트폴리오
	S포트폴리오(i포트폴리오 I포트폴리오)
}
