package price_data

import (
	공용 "github.com/ghts/ghts/shared"
	공용_정보 "github.com/ghts/ghts/shared_data"
	zmq "github.com/pebbe/zmq4"
	
	"math"
	"strconv"
	"time"
)

type S구독소켓_등록 struct {
	M구독_소켓 *zmq.Socket
	M회신_채널 chan error
}

var Ch가격정보 = make(chan 공용.I질의, 1000)
var Ch가격정보_구독_채널 = make(chan (chan 공용.I가격정보), 100)
var Ch가격정보_구독_소켓 = make(chan S구독소켓_등록, 100)
var ch종료_가격정보_Go루틴 = make(chan 공용.S비어있는_구조체, 1)
var 가격정보_Go루틴_실행_중 = 공용.New안전한_bool(false)

func F가격정보_Go루틴_실행_중() bool {
	return 가격정보_Go루틴_실행_중.G값()
}

func F가격정보_Go루틴(ch초기화 chan bool) {
	에러 := 가격정보_Go루틴_실행_중.S값(true)
	if 에러 != nil {
		ch초기화 <- false; return
	}
	
	가격정보_맵 := make(map[string]공용.I가격정보)
	구독채널_맵 := make(map[chan 공용.I가격정보]공용.S비어있는_구조체)
	공통_종료_채널 := 공용.F공통_종료_채널()
	
	// 가격정보_PUB 소켓 초기화
	가격정보_PUB, 에러 := zmq.NewSocket(zmq.PUB)
	if 에러 != nil {
		공용.F에러_출력(에러)
		ch초기화 <- false; return
	}
	
	defer 가격정보_PUB.Close()
	
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
	
	에러 = 가격정보_PUB.Bind(p주소_가격정보_배포)
	if 에러 != nil {
		공용.F에러_출력(에러)
		ch초기화 <- false; return
	}
	
	// GC 압력을 줄이기 위한 재활용 변수
	질의 := 공용.New질의(공용.P메시지_GET)
	
	// 초기화 완료
	ch초기화 <- true
	
	for {
		select {
		case 질의 = <-Ch가격정보:
			switch 질의.G구분() {
			case 공용.P메시지_GET:
				에러 = f캐시된_가격정보_회신(질의, 가격정보_맵)
				if 에러 == nil {
					// 캐시된 데이터가 있으면 배포하지 않음
					break
				}
				
				공용.F메모("TODO : f새로운_가격정보_회신()")
				가격정보, 에러 := f새로운_가격정보_회신(질의, 가격정보_맵)
				if 에러 != nil {
					공용.F문자열_출력(질의.String())
					공용.F에러_출력(에러)
					break
				}
				
				에러 = f가격정보_배포(가격정보, 구독채널_맵, 가격정보_PUB)
				if 에러 != nil {
					공용.F문자열_출력(질의.String())
					공용.F에러_출력(에러)
					break
				} 
			case 공용.P메시지_SET:
				가격정보, 에러 := f가격정보_캐시_저장(질의, 가격정보_맵)
				if 에러 != nil {
					공용.F문자열_출력(질의.String())
					공용.F에러_출력(에러)
				}
				
				에러 = f가격정보_배포(가격정보, 구독채널_맵, 가격정보_PUB)
				if 에러 != nil {
					공용.F문자열_출력(질의.String())
					공용.F에러_출력(에러)
					break
				}
			default:
				질의.S회신(공용.F에러_생성("예상치 못한 메시지 구분.\n%v", 질의))
			}
		case ch구독 := <-Ch가격정보_구독_채널:
			구독채널_맵[ch구독] = 공용.S비어있는_구조체{}
		case 구독신청 := <-Ch가격정보_구독_소켓:
			에러 := 구독신청.M구독_소켓.Connect(p주소_가격정보_배포)		
			구독신청.M회신_채널 <-에러
		case <-공통_종료_채널:
			ch종료_가격정보_Go루틴 <- 공용.S비어있는_구조체{}
		case <-ch종료_가격정보_Go루틴:
			가격정보_Go루틴_실행_중.S값(false)
			return	
		}
	}
}

func f캐시된_가격정보_회신(질의 공용.I질의, 가격정보_맵 map[string]공용.I가격정보) error {
	// 0 : 종목코드, 1 : 유효기간 (초 단위)
	if 에러 := 질의.G검사(공용.P메시지_GET, 2); 에러 != nil {
		질의.S회신(에러); return 에러
	}
	
	초_단위_유효기간, 에러 := strconv.ParseFloat(질의.G내용(1), 64)
	if 에러 != nil {
		질의.S회신(에러); return 에러
	}
	
	if 초_단위_유효기간 <= 0 {
		return 공용.F에러_생성("유효기간이 0과 같거나 음수임. 가격정보 캐시 확인 중단.") 
	}
	
	종목코드 := 질의.G내용(0)
	가격정보, 존재함 := 가격정보_맵[종목코드]
	if !존재함 {
		에러 = 공용.F에러_생성("%v의 가격정보 캐시 데이터가 없음.\n%v", 종목코드, 질의)
		질의.S회신(에러); return 에러
	}
		
	차이 := math.Abs(time.Now().Sub(가격정보.G시점()).Seconds())
		
	if 차이 > 초_단위_유효기간 {
		에러 = 공용.F에러_생성("유효기간이 지났음. 차이 %v초.", 차이)
		질의.S회신(에러); return 에러
	}
	
	질의.S회신(nil, 
		가격정보.G가격().G단위(), 
		가격정보.G가격().G문자열값(), 
		가격정보.G시점().Format(공용.P시간_형식))
	
	return nil
}

func f새로운_가격정보_회신(
	질의 공용.I질의, 가격정보_맵 map[string]공용.I가격정보) (공용.I가격정보, error) {
	공용.F메모("TODO f새로운_가격정보_회신().\n가격정보 입수 루틴 작성 후 진행 가능.")
	
	return nil, 공용.F에러_생성("아직 구현하지 못함.")
}

func f가격정보_캐시_저장(질의 공용.I질의, 가격정보_맵 map[string]공용.I가격정보) (공용.I가격정보, error) {
	// 0 : 종목코드, 1 : 통화단위, 2 : 금액
	에러 := 질의.G검사(공용.P메시지_SET, 3)
	if 에러 != nil {
		질의.S회신(에러)
		return nil, 에러
	}
	
	에러 = 공용.F통화단위_검사(질의.G내용(1))
	if 에러 != nil {
		질의.S회신(에러)
		return nil, 에러
	}
	
	금액, 에러 := strconv.ParseFloat(질의.G내용(2), 64)
	if 에러 != nil {
		질의.S회신(에러)
		return nil, 에러
	}
		
	종목코드 := 질의.G내용(0)
	통화단위 := 질의.G내용(1)
	통화 := 공용.New통화(통화단위, 금액)
	
	if 통화 == nil {
		에러 := 공용.F에러_생성("통화 생성 에러.\n%v", 질의)
		질의.S회신(에러)
		return nil, 에러
	}
	
	가격정보 := 공용.New가격정보(종목코드, 통화)
	
	가격정보_맵[종목코드] = 가격정보
	
	질의.S회신(nil)
	return 가격정보, nil
}

func f가격정보_배포(가격정보 공용.I가격정보, 
	구독채널_맵 map[chan 공용.I가격정보]공용.S비어있는_구조체,
	가격정보_PUB *zmq.Socket) error {
	
	ch실행_결과 := make(chan error, 2) 
	
	go f가격정보_배포_Go채널(ch실행_결과, 가격정보, 구독채널_맵)
	go f가격정보_배포_zmq소켓(ch실행_결과, 가격정보, 가격정보_PUB)
	
	에러_모음 := make([]error, 0)
	
	for i:=0 ; i<2 ; i++ {
		에러 := <- ch실행_결과
		
		if 에러 != nil {
			에러_모음 = append(에러_모음, 에러)
		}
	}
	
	if len(에러_모음) == 0 {
		return nil
	}
	
	에러_문자열 := ""
	
	for i:=0 ; i<len(에러_모음); i++ {
		에러_문자열 += 에러_모음[i].Error() + "\n"
	}
	
	return 공용.F에러_생성(에러_문자열)
}

func f가격정보_배포_Go채널(
		ch실행_결과 chan error,
		가격정보 공용.I가격정보, 
		구독채널_맵 map[chan 공용.I가격정보]공용.S비어있는_구조체) {
	defer func() {
		r := recover()
		
		if r != nil {
			ch실행_결과 <- 공용.F에러_생성(공용.F포맷된_문자열("%v", r))
		} else {
			ch실행_결과 <- nil
		}	
	}()
	
	for ch구독, _ := range 구독채널_맵 {
		ch구독 <- 가격정보
	}
}
		
func f가격정보_배포_zmq소켓(
		ch실행_결과 chan error,
		가격정보 공용.I가격정보, 
		가격정보_PUB *zmq.Socket) {
	defer func() {
		r := recover()
		
		if r != nil {
			ch실행_결과 <- 공용.F에러_생성(공용.F포맷된_문자열("%v", r))
		} else {
			ch실행_결과 <- nil
		}	
	}()
	
	_, 에러 := 가격정보_PUB.SendMessage(
					가격정보.G종목코드(),
					가격정보.G가격().G단위(), 
					가격정보.G가격().G문자열값(),
					가격정보.G시점().Format(공용.P시간_형식)) 
	
	if 에러 != nil {
		panic(에러)
	}
}