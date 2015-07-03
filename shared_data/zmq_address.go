package shared_data

import (
	공용 "github.com/ghts/ghts/shared"
	//zmq "github.com/pebbe/zmq4"
	
	"strconv"
	"sync"
)

type T주소_이름 string

const (
	P주소정보 T주소_이름 = "주소정보"
	P테스트_결과 = "테스트_결과"

	P종목정보 = "종목정보"
	P가격정보 = "가격정보"
	P가격정보_입수 = "가격정보_입수"
	P가격정보_배포 = "가격정보_배포"
)

func init() {
	// 아래 2개의 주소는 고정.
	주소정보_맵[P주소정보] = "tcp://127.0.0.1:3001"
	주소정보_맵[P테스트_결과] = "tcp://127.0.0.1:3002"
	
	// 나머지 주소는 생성자 함수를 이용해서 자동 생성.
	f주소_추가(P종목정보)
	f주소_추가(P가격정보)
	f주소_추가(P가격정보_입수)
	f주소_추가(P가격정보_배포)
}

var 포트번호 = 3010

func f주소_추가(주소_이름 T주소_이름) {
	주소정보_맵[주소_이름] = "tcp://127.0.0.1:" + strconv.Itoa(포트번호)
	포트번호++
}

type I주소정보_질의 interface {
	G주소_이름() T주소_이름
	G회신_채널() chan 공용.I회신
}

func New주소정보_질의(주소_이름 T주소_이름, 회신_채널 chan 공용.I회신) I주소정보_질의 {
	return s주소정보_질의 {주소_이름: 주소_이름, 회신_채널: 회신_채널}
}

type s주소정보_질의 struct {
	주소_이름 T주소_이름
	회신_채널 chan 공용.I회신
}
func (this s주소정보_질의) G주소_이름() T주소_이름 { return this.주소_이름 }
func (this s주소정보_질의) G회신_채널() chan 공용.I회신 { return this.회신_채널 }

// 테스트에서 에러가 발생하지 않고 정상 동작한다면 가능한한 버퍼 추가할 것.
var Ch주소정보_질의채널 = make(chan I주소정보_질의)
var ch주소정보_종료채널 = 공용.F공통_종료_채널()

var 주소정보_맵 = make(map[T주소_이름]string)

var 주소정보_관리_채널방식_Go루틴_실행_중 = false
var 주소정보_관리_채널방식_Go루틴_잠금 = &sync.RWMutex{}

var 주소정보_관리_소켓방식_Go루틴_실행_중 = false
var 주소정보_관리_소켓방식_Go루틴_잠금 = &sync.RWMutex{}

func F주소정보_관리_채널방식_Go루틴(실행_성공_회신_채널 chan bool) {
	주소정보_관리_채널방식_Go루틴_잠금.RLock()
	실행_중 := 주소정보_관리_채널방식_Go루틴_실행_중
	주소정보_관리_채널방식_Go루틴_잠금.RUnlock()
	
	if 실행_중 {
		실행_성공_회신_채널 <- false
		return
	}
	
	주소정보_관리_채널방식_Go루틴_잠금.Lock()

	if 주소정보_관리_채널방식_Go루틴_실행_중 {
		주소정보_관리_채널방식_Go루틴_잠금.Unlock()
		실행_성공_회신_채널 <- false
		return
	} else {
		주소정보_관리_채널방식_Go루틴_실행_중 = true
		주소정보_관리_채널방식_Go루틴_잠금.Unlock()
		실행_성공_회신_채널 <- true
	}
	
	for {
		select {
		case 주소정보_질의 := <-Ch주소정보_질의채널:
			주소, 존재함 := 주소정보_맵[주소정보_질의.G주소_이름()]
			
			var 회신 공용.I회신
			
			if !존재함 {
				에러 := 공용.F에러_생성("존재하지 않는 주소 이름. %v", 주소정보_질의.G주소_이름())
				회신 = 공용.New회신(nil, 에러)
			} else {
				회신 = 공용.New회신([]string{주소}, nil)
			}
			
			주소정보_질의.G회신_채널() <- 회신
		case <-ch주소정보_종료채널:
			주소정보_관리_채널방식_Go루틴_잠금.Lock()
			주소정보_관리_채널방식_Go루틴_실행_중 = false
			주소정보_관리_채널방식_Go루틴_잠금.Unlock()

			return
		}
	}
}

func F주소정보_관리_소켓방식_Go루틴(실행_성공_회신_채널 chan bool) {
	주소정보_관리_소켓방식_Go루틴_잠금.RLock()
	실행_중 := 주소정보_관리_소켓방식_Go루틴_실행_중
	주소정보_관리_소켓방식_Go루틴_잠금.RUnlock()
	
	주소정보_관리_소켓방식_Go루틴_잠금.Lock()

	if 주소정보_관리_소켓방식_Go루틴_실행_중 {
		주소정보_관리_소켓방식_Go루틴_잠금.Unlock()
		실행_성공_회신_채널 <- false
		return
	} else {
		주소정보_관리_소켓방식_Go루틴_실행_중 = true
		주소정보_관리_소켓방식_Go루틴_잠금.Unlock()
	}
	
	기저_Go루틴_실행_채널 := make(chan bool)
	go F주소정보_관리_채널방식_Go루틴(기저_Go루틴_실행_채널)
	<-기저_Go루틴_실행_채널
	
	회신_채널 := make(chan 공용.I회신)
	Ch주소정보_질의채널 <- New주소정보_질의(P주소정보, 회신_채널)
	주소_질의_회신 := <-회신_채널
	공용.F에러_체크(주소_질의_회신.G에러())
	
	// 소켓 초기화
	주소정보_REP, 에러 := zmq.NewSocket(zmq.REP)
	공용.F에러_체크(에러)
	defer 주소정보_REP.Close()
	주소정보_REP.Bind(주소_질의_회신.G내용()[0])

	// 초기화 완료.
	실행_성공_회신_채널 <- true

	reactor := zmq.NewReactor()
	reactor.AddSocket(주소정보_REP, zmq.POLLIN, func(e zmq.State) error { return f주소정보_제공(주소정보_REP) })
	
	채널 = <-chan interface{}
	채널_id := reactor.AddChannel(채널, 

	for {
		에러 = reactor.Run(-1)
		if 에러 != nil {
			공용.F문자열_출력("reactor 재시작 : %v", 에러)
		}
	}
}
