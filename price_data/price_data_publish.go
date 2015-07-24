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

var zmq소켓_가격정보_배포_Go루틴_실행_중 = 공용.New안전한_bool(false)
var Ch가격정보_구독소켓_등록 = make(chan S구독소켓_등록, 100)
var ch종료_zmq소켓_가격정보_배포_Go루틴 = make(chan 공용.S비어있는_구조체)
var ch가격정보_배포_zmq소켓 = make(chan 공용.I질의, 10000)

// zmq소켓은 동시에 사용하면 에러난다.
// 그렇다고 mutex로 보호하면 데드락이 생길 위험이 있다.
// 그래서, 별도의 go루틴으로 분리하고, 버퍼가 있는 채널을 이용해서,
// 데드락도 방지하면서, 동시 사용으로 인한 문제도 방지한다.
// 버퍼가 있는 채널도 블록될 가능성이 있으니, 이 문제에 대해서 추가 연구 필요함.
func f_zmq소켓_가격정보_배포_Go루틴(ch초기화 chan bool) {
	에러 := zmq소켓_가격정보_배포_Go루틴_실행_중.S값(true)
	if 에러 != nil {
		ch초기화 <- false; return
	}
	
	// 가격정보_PUB 주소 알아내기
	if !공용_정보.F공용_데이터_Go루틴_실행_중() {
		ch초기화_대기 := make(chan bool)
		go 공용_정보.F공용_데이터_Go루틴(ch초기화_대기)
		<-ch초기화_대기
	}
	
	회신 := 공용.New질의(공용.P메시지_GET, 공용.P주소명_가격정보_배포).G회신(공용_정보.Ch주소, 공용.P타임아웃_Go)
		
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
			
			_, 에러 = 가격정보_PUB.SendMessage(공용.F2인터페이스_모음(질의.G내용_전체()))
			질의.S회신(에러)
		case 구독신청 := <-Ch가격정보_구독소켓_등록:
			에러 := 구독신청.M구독_소켓.Connect(p주소_가격정보_배포)		
			구독신청.M회신_채널 <-에러
		case <-공통_종료_채널:
			ch종료_zmq소켓_가격정보_배포_Go루틴 <- 공용.S비어있는_구조체{}
		case <-ch종료_zmq소켓_가격정보_배포_Go루틴:
			return
		default:
		}
	}
}
