package lib

import "fmt"

// TR 및 응답 종류
const (
	TR조회 TR구분 = iota
	TR주문
	TR실시간_정보_구독
	TR실시간_정보_해지
	TR실시간_정보_일괄_해지
	TR접속
	TR접속됨
	TR접속_해제
	TR초기화
	TR종료
)

type TR구분 uint8

func (v TR구분) String() string {
	return TR구분_String(v)
}

// 증권사 API 패키지에서 오버라이드 될 수 있음.
var TR구분_String = func(v TR구분) string {
	switch v {
	case TR조회:
		return "TR조회"
	case TR주문:
		return "TR주문"
	case TR실시간_정보_구독:
		return "TR실시간 정보 구독"
	case TR실시간_정보_해지:
		return "TR실시간 정보 해지"
	case TR실시간_정보_일괄_해지:
		return "TR실시간 정보 일괄 해지"
	case TR접속:
		return "TR접속"
	case TR접속됨:
		return "TR접속됨"
	case TR접속_해제:
		return "TR접속 해제"
	case TR초기화:
		return "TR초기화"
	case TR종료:
		return "TR종료"
	default:
		return fmt.Sprintf("예상하지 못한 M값 : '%v'", uint8(v))
	}
}

const (
	TR응답_데이터 TR응답_구분 = iota
	TR응답_실시간_정보
	TR응답_메시지
	TR응답_완료
)

type TR응답_구분 int8

func (v TR응답_구분) String() string {
	switch v {
	case TR응답_데이터:
		return "TR응답_데이터"
	case TR응답_실시간_정보:
		return "TR응답_실시간_정보"
	case TR응답_메시지:
		return "TR응답_메시지"
	case TR응답_완료:
		return "TR응답_완료"
	default:
		return fmt.Sprintf("예상하지 못한 M값. %v", int8(v))
	}
}
