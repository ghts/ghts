package shared_data

import (
	공용 "github.com/ghts/ghts/shared"
	zmq "github.com/pebbe/zmq4"
	
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
	주소정보_맵.Lock()
	주소정보_맵[P주소정보] = "tcp://127.0.0.1:3001"
	주소정보_맵[P테스트_결과] = "tcp://127.0.0.1:3002"
	주소정보_맵.Unlock()
	
	// 나머지 주소는 생성자 함수를 이용해서 자동 생성.
	f주소_추가(P종목정보)
	f주소_추가(P가격정보)
	f주소_추가(P가격정보_입수)
	f주소_추가(P가격정보_배포)
}

var 주소정보_맵 = make(map[string]string)
var 주소정보_맵_잠금 = &sync.RWMutex{}

var 포트번호 = 3010

func f주소_추가(주소 T주소_이름) {
	주소정보_맵_잠금.Lock()
	주소정보_맵[주소.(string)] = "tcp://127.0.0.1:" + strconv.Itoa(포트번호)
	주소정보_맵_잠금.Unlock()
	
	포트번호++
}

var 주소정보_관리_채널방식_Go루틴_실행_중 = false
var 주소정보_관리_채널방식_Go루틴_잠금 = &sync.RWMutex{}

var 주소정보_관리_소켓방식_Go루틴_실행_중 = false
var 주소정보_관리_소켓방식_Go루틴_잠금 = &sync.RWMutex{}

type S주소정보_질의 struct {
	M주소_이름 T주소_이름
	M회신_채널 chan I회신
} 

var Ch주소정보_질의채널 = chan S주소정보_질의

func F주소정보_관리_채널방식_Go루틴_실행_중() bool {
	주소정보_관리_채널방식_Go루틴_잠금.RLock()
	defer 주소정보_관리_채널방식_Go루틴_잠금.RUnlock()
	
	return 주소정보_관리_채널방식_Go루틴_실행_중
}

func F주소정보_관리_소켓방식_Go루틴_실행_중() bool {
	주소정보_관리_소켓방식_Go루틴_잠금.RLock()
	defer 주소정보_관리_소켓방식_Go루틴_잠금.RUnlock()
	
	return 주소정보_관리_소켓방식_Go루틴_실행_중
}

func F주소정보_관리_채널방식_Go루틴(실행_성공_회신_채널 chan bool) {
	if F주소정보_관리_채널방식_Go루틴_실행_중() {
		실행_성공_회신_채널 <- false
		return
	}

	주소정보_관리_채널방식_Go루틴_잠금.Lock()

	// Go루틴 시작하기 전에 마지막으로 1번 더 확인.
	if 주소정보_관리_Go루틴_실행_중 {
		// 그 짧은 시간 사이에 Go 루틴이 생성되었다면 바로 종료.
		주소정보_관리_Go루틴_잠금.Unlock()

		실행_성공_회신_채널 <- false
		return
	}

	주소정보_관리_Go루틴_실행_중 = true
	주소정보_관리_Go루틴_잠금.Unlock()
	
	// 준비 완료
	실행_성공_회신_채널 <- true
	
	for {
		select {
		case 주소정보_질의 := <-Ch주소정보_질의채널:
			주소정보_맵_잠금.RLock()
			주소, ok := 주소정보_맵[주소정보_질의.M주소_이름]
			실행_내역_맵[실행_내역.M_pid] = 실행_내역

			에러 = f실행_내역_맵_파일에_저장(실행_내역_맵)
			F에러_체크(에러)
		case <-종료_채널:
			주소정보_관리_채널방식_Go루틴_잠금.Lock()
			주소정보_관리_채널방식_Go루틴_실행_중 = false
			주소정보_관리_채널방식_Go루틴_잠금.Unlock()

			return
	}
	
}

func F주소정보_관리_소켓방식_Go루틴(실행_성공_회신_채널 chan bool) {
	if F주소정보_관리_Go루틴_실행_중() {
		실행_성공_회신_채널 <- false
		return
	}

	주소정보_관리_Go루틴_잠금.Lock()

	// Go루틴 시작하기 전에 마지막으로 1번 더 확인.
	if 주소정보_관리_Go루틴_실행_중 {
		// 그 짧은 시간 사이에 Go 루틴이 생성되었다면 바로 종료.
		주소정보_관리_Go루틴_잠금.Unlock()

		실행_성공_회신_채널 <- false
		return
	}

	주소정보_관리_Go루틴_실행_중 = true
	주소정보_관리_Go루틴_잠금.Unlock()
	
	// 소켓 초기화
	주소정보_REP, 에러 := zmq.NewSocket(zmq.REP)
	공용.F에러_체크(에러)
	defer 주소정보_REP.Close()
	
	주소정보_REP.Bind(공용.P주소_주소정보.String())

	종목정보_REP, 에러 := zmq.NewSocket(zmq.REP)
	defer 종목정보_REP.Close()
	if 에러 != nil {
		panic(에러)
	}
	종목정보_REP.Bind(공용.P주소_종목정보.String())

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
	
	// 주소 정보 서비스 제공 준비 완료.
	실행_성공_회신_채널 <- true

}

