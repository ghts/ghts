package shared_data

import (
	공용 "github.com/ghts/ghts/shared"

	"strconv"
	"sync"
)

func init() {
	// 아래 2개의 주소는 고정.
	주소정보_맵[P주소정보] = 공용.P주소정보_주소
	주소정보_맵[P테스트_결과] = 공용.P테스트_결과_주소

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

// 테스트에서 에러가 발생하지 않고 정상 동작한다면 가능한한 버퍼 추가할 것.
var Ch주소정보_질의_채널 = make(chan I주소정보_질의)
var ch주소정보_종료_채널 chan 공용.S비어있는_구조체 = nil

var 주소정보_맵 = make(map[T주소_이름]string)

var 주소정보_관리_Go루틴_실행_중 = false
var 주소정보_관리_Go루틴_잠금 = &sync.RWMutex{}

func F주소정보_관리_Go루틴_실행_중() bool {
	주소정보_관리_Go루틴_잠금.RLock()
	defer 주소정보_관리_Go루틴_잠금.RUnlock()
	
	return 주소정보_관리_Go루틴_실행_중
}

func F주소정보_관리_Go루틴(
		종료_채널 chan 공용.S비어있는_구조체, 
		실행_회신_채널 chan bool) {
	if F주소정보_관리_Go루틴_실행_중() {
		실행_회신_채널 <- false
		return
	}

	주소정보_관리_Go루틴_잠금.Lock()

	// 마지막으로 1번 더 확인
	if 주소정보_관리_Go루틴_실행_중 {
		주소정보_관리_Go루틴_잠금.Unlock()
		실행_회신_채널 <- false
		return
	}

	// 진짜로 실행	
	주소정보_관리_Go루틴_실행_중 = true
	주소정보_관리_Go루틴_잠금.Unlock()
	실행_회신_채널 <- true
	
	ch주소정보_종료_채널 = 종료_채널 
	
	for {
		select {
		case 주소정보_질의 := <-Ch주소정보_질의_채널:
			주소, 존재함 := 주소정보_맵[주소정보_질의.G주소_이름()]

			var 회신 공용.I회신

			if !존재함 {
				에러 := 공용.F에러_생성("존재하지 않는 주소 이름. %v", 주소정보_질의.G주소_이름())
				회신 = 공용.New회신(nil, 에러)
			} else {
				회신 = 공용.New회신([]string{주소}, nil)
			}

			주소정보_질의.G회신_채널() <- 회신
		case <-ch주소정보_종료_채널:
			주소정보_관리_Go루틴_잠금.Lock()
			주소정보_관리_Go루틴_실행_중 = false
			주소정보_관리_Go루틴_잠금.Unlock()

			return
		}
	}
}
