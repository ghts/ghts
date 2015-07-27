package price_data

import (
	공용 "github.com/ghts/ghts/shared"
	공용_정보 "github.com/ghts/ghts/shared_data"
	zmq "github.com/pebbe/zmq4"
)

type S구독소켓_등록 struct {
	M구독_소켓 *zmq.Socket
	M회신_채널 chan error
}

var 가격정보_배포_Go루틴_zmq소켓_실행_중 = 공용.New안전한_bool(false)
var ch제어_가격정보_배포_Go루틴_zmq소켓 = make(chan 공용.I질의)
var ch가격정보_배포_zmq소켓 = make(chan 공용.I질의, 10000)

func f가격정보_배포_Go루틴_zmq소켓(ch초기화 chan bool) {
	에러 := 가격정보_배포_Go루틴_zmq소켓_실행_중.S값(true)
	if 에러 != nil {
		ch초기화 <- false; return
	}
	
	// 가격정보_PUB 주소 알아내기
	공용_정보.F공용정보_모듈_실행()
	
	회신 := 공용.New질의(공용.P메시지_GET, 공용.P주소명_가격정보_배포).G회신(공용_정보.Ch주소)

	if 회신.G에러() != nil {
		공용.F에러_출력(회신.G에러())
		ch초기화 <- false; return
	}
	
	p주소_가격정보_배포 := 회신.G내용(0)
	
	// 가격정보_PUB 소켓 초기화
	가격정보_PUB, 에러 := zmq.NewSocket(zmq.PUB)
	if 에러 != nil {
		공용.F에러_출력(에러)
		ch초기화 <- false; return
	}
	
	에러 = 가격정보_PUB.Bind(p주소_가격정보_배포)
	if 에러 != nil {
		공용.F에러_출력(에러)
		ch초기화 <- false; return
	}
	
	defer 가격정보_PUB.Close()
	
	공통_종료_채널 := 공용.F공통_종료_채널()
	
	// 초기화 완료
	ch초기화 <- true
	
	for {
		select {
		case 질의 := <-ch가격정보_배포_zmq소켓:
			에러 := 질의.G검사(공용.P메시지_SET, 4)
			if 에러 != nil {
				질의.S회신(에러)
				continue
			}
			
			_, 에러 = 가격정보_PUB.SendMessage(질의.G내용_전체_가변형()...)
			질의.S회신(에러)
		case 질의 := <-ch제어_가격정보_배포_Go루틴_zmq소켓:
			switch 질의.G구분() {
			case 공용.P메시지_종료:
				return
			default:
				에러 = 공용.F에러_생성("예상치 못한 구분. %v", 질의.G구분())
				panic(에러)
			}
		case <-공통_종료_채널:
			return
		default:
		}
	}
}