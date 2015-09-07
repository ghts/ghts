package price_data

import (
	공용 "github.com/ghts/ghts/common"

	"math"
	"runtime"
	"strconv"
	"time"
)

type s가격정보_배포_대기 struct {
	가격정보 공용.I가격정보
	ch구독 chan 공용.I가격정보
}

var Ch가격정보 = make(chan 공용.I질의, 1000)
var Ch가격정보_구독채널_등록 = make(chan (chan 공용.I가격정보)) // 버퍼 만들지 말 것.
var ch제어_가격정보_Go루틴 = make(chan 공용.I질의)

var ch가격정보_배포_대기열 = make(chan *s가격정보_배포_대기, 100000) // 버퍼가 아주 큰 대기열
var ch가격정보_배포_Go루틴_종료 chan 공용.S비어있는_구조체

var 가격정보_Go루틴_실행_중 = 공용.New안전한_bool(false)

func F가격정보_모듈_실행_중() bool {
	return 가격정보_Go루틴_실행_중.G값()
}

func F가격정보_모듈_실행() bool {
	if 가격정보_Go루틴_실행_중.G값() {
		return false
	}

	ch초기화_대기 := make(chan bool)
	go f가격정보_Go루틴(ch초기화_대기)
	return <-ch초기화_대기
}

// 가격정보에 관련된 모든 기능의 허브
// - 가격정보 캐시 : 증권사 서버로 보내는 질의 수량 최소화 및 성능 향상.
// - 가격정보 수신 : 가격정보 입수 기능이 가격정보를 PUSH할 경우 이를 캐시 데이터에 저장하고, 배포.
// - 가격정보 배포 : 질의 응답 및 수신 과정에서 새로운 가격정보를 입수하게 되면 이를 배포함.
//
// 추가 개발이 필요한 기능
// - 가격정보 입수 : 증권사 서버에서 가격을 수신하는 기능
// - 가격정보 질의 응답 : 가격정보 질의 응답 (캐시 데이터 확인 혹은 새로운 데이터 입수)

func f가격정보_Go루틴(ch초기화 chan bool) {
	// 이미 실행 중인 가격정보_Go루틴이 존재하는 지 확인
	에러 := 가격정보_Go루틴_실행_중.S값(true)
	if 에러 != nil {
		ch초기화 <- false
		return
	}

	// zmq소켓을 통해서 가격정보 배포하는 Go루틴 실행.
	ch초기화_대기 := make(chan bool)

	go f가격정보_배포_Go루틴_zmq소켓(ch초기화_대기)
	<-ch초기화_대기

	defer 공용.New질의(공용.P메시지_종료).S질의(ch제어_가격정보_배포_Go루틴_zmq소켓).G회신()

	// 변수 초기화
	가격정보_맵 := make(map[string]공용.I가격정보)
	구독채널_맵 := make(map[chan 공용.I가격정보]공용.S비어있는_구조체)
	ch공통_종료 := 공용.F공통_종료_채널()
	질의 := 공용.New질의(공용.P메시지_GET) // GC 압력을 줄이기 위한 재활용 변수

	// 가격정보 배포 Go루틴 시작
	가격정보_배포_Go루틴_수량 := runtime.NumCPU()

	for i := 0; i < 가격정보_배포_Go루틴_수량; i++ {
		go f가격정보_배포_Go루틴()
	}

	ch가격정보_배포_Go루틴_종료 = make(chan 공용.S비어있는_구조체)

	// 초기화 완료
	ch초기화 <- true

	for {
		select {
		case 질의 = <-Ch가격정보:
			switch 질의.G구분() {
			case 공용.P메시지_GET:
				가격정보, 에러 := f캐시된_가격정보_검색(질의, 가격정보_맵)
				if 에러 == nil {
					f가격정보_회신_보내기(질의, 가격정보)
					continue
				}

				공용.F메모("TODO : f증권사_API_가격정보_질의()")

				가격정보, 에러 = f증권사_API_가격정보_질의(질의)
				if 에러 != nil {
					공용.F문자열_출력(질의.String())
					공용.F에러_출력(에러)
					질의.S회신(에러)
					break
				}

				// 캐시에 저장
				가격정보_맵[가격정보.G종목코드()] = 가격정보

				f가격정보_회신_보내기(질의, 가격정보)

				f가격정보_배포_대기열에_추가(가격정보, 구독채널_맵)
			case 공용.P메시지_SET:
				가격정보, 에러 := f가격정보_캐시_저장(질의, 가격정보_맵)

				if 에러 != nil {
					공용.F문자열_출력(질의.String())
					공용.F에러_출력(에러)
					질의.S회신(에러)
					break
				}

				질의.S회신(nil)

				f가격정보_배포_대기열에_추가(가격정보, 구독채널_맵)

				회신 := 공용.New질의(공용.P메시지_SET, 가격정보.G종목코드(),
					가격정보.G가격().G단위(),
					가격정보.G가격().G문자열값(),
					가격정보.G시점()).S질의(ch가격정보_배포_zmq소켓).G회신()

				if 회신.G에러() != nil {
					공용.F에러_출력(회신.G에러())
				}
			default:
				질의.S회신(공용.F에러_생성("예상치 못한 메시지 구분.\n%v", 질의))
			}
		case ch구독 := <-Ch가격정보_구독채널_등록:
			구독채널_맵[ch구독] = 공용.S비어있는_구조체{}
		case 질의 := <-ch제어_가격정보_Go루틴:
			switch 질의.G구분() {
			case 공용.P메시지_초기화:
				가격정보_맵 = make(map[string]공용.I가격정보)
				구독채널_맵 = make(map[chan 공용.I가격정보]공용.S비어있는_구조체)
				질의.S회신(nil)
			case 공용.P메시지_종료:
				f가격정보_Go루틴_종료()
				질의.S회신(nil)
				return
			default:
				에러 := 공용.F에러_생성("예상치 못한 질의 메시지 구분값 %v", 질의.G구분())
				공용.F에러_출력(에러)
				질의.S회신(에러)
				panic("")
			}
		case <-ch공통_종료:
			f가격정보_Go루틴_종료()
			return
		}
	}
}

func f캐시된_가격정보_검색(질의 공용.I질의, 가격정보_맵 map[string]공용.I가격정보) (공용.I가격정보, error) {
	// 0 : 종목코드, 1 : 유효기간 (초 단위)
	if 에러 := 질의.G검사(공용.P메시지_GET, 2); 에러 != nil {
		return nil, 에러
	}

	초_단위_유효기간, 에러 := strconv.ParseFloat(질의.G내용(1), 64)
	if 에러 != nil {
		return nil, 에러
	}

	if 초_단위_유효기간 < 0 {
		에러 := 공용.F에러_생성("유효기간이 음수임.")
		panic(에러)
		//return nil, 에러
	}

	종목코드 := 질의.G내용(0)
	가격정보, 존재함 := 가격정보_맵[종목코드]
	if !존재함 {
		에러 = 공용.F에러_생성("%v의 가격정보 캐시 데이터가 없음.\n%v", 종목코드, 질의)
		return nil, 에러
	}

	차이 := math.Abs(time.Now().Sub(가격정보.G시점()).Seconds())

	if 차이 > 초_단위_유효기간 {
		에러 = 공용.F에러_생성("유효기간이 지났음. 차이 %v초.", 차이)
		return nil, 에러
	}

	return 가격정보, nil
}

func f증권사_API_가격정보_질의(질의 공용.I질의) (공용.I가격정보, error) {
	return nil, 공용.F에러_생성("아직 구현하지 못함.")
}

func f가격정보_회신_보내기(질의 공용.I질의, 가격정보 공용.I가격정보) {
	질의.S회신(nil,
		가격정보.G가격().G단위(),
		가격정보.G가격().G문자열값(),
		가격정보.G시점().Format(공용.P시간_형식))
}

func f가격정보_캐시_저장(질의 공용.I질의, 가격정보_맵 map[string]공용.I가격정보) (공용.I가격정보, error) {
	// 0 : 종목코드, 1 : 통화단위, 2 : 금액, 3 : 시점
	에러 := 질의.G검사(공용.P메시지_SET, 4)
	if 에러 != nil {
		return nil, 에러
	}

	에러 = 공용.F통화단위_검사(질의.G내용(1))
	if 에러 != nil {
		return nil, 에러
	}

	금액, 에러 := strconv.ParseFloat(질의.G내용(2), 64)
	if 에러 != nil {
		return nil, 에러
	}

	종목코드 := 질의.G내용(0)
	통화단위 := 질의.G내용(1)
	통화 := 공용.New통화(통화단위, 금액)

	if 통화 == nil {
		return nil, 공용.F에러_생성("통화 생성 에러.\n%v", 질의)
	}

	시점, 에러 := time.Parse(공용.P시간_형식, 질의.G내용(3))
	if 에러 != nil {
		return nil, 에러
	}

	가격정보 := 공용.New가격정보(종목코드, 통화, 시점)

	가격정보_맵[종목코드] = 가격정보

	return 가격정보, nil
}

func f가격정보_배포_대기열에_추가(가격정보 공용.I가격정보, 구독채널_맵 map[chan 공용.I가격정보]공용.S비어있는_구조체) {
	for ch구독, _ := range 구독채널_맵 {
		대기열_구성요소 := s가격정보_배포_대기{가격정보: 가격정보, ch구독: ch구독}
		ch가격정보_배포_대기열 <- &대기열_구성요소
	}
}

func f가격정보_배포_Go루틴() {
	종료신호_수신 := false

	for {
		select {
		case s := <-ch가격정보_배포_대기열:
			f가격정보_배포_도우미(s.가격정보, s.ch구독)
		case <-ch가격정보_배포_Go루틴_종료:
			종료신호_수신 = true
		default:
			if len(ch가격정보_배포_대기열) == 0 && 종료신호_수신 {
				return
			}

			time.Sleep(50 * time.Millisecond)
		}
	}
}

func f가격정보_배포_도우미(가격정보 공용.I가격정보, ch구독 chan 공용.I가격정보) {
	select {
	case ch구독 <- 가격정보:
	case <-time.After(공용.P타임아웃_Go):
	}
}

func f가격정보_Go루틴_종료() {
	close(ch가격정보_배포_Go루틴_종료)

	가격정보_Go루틴_실행_중.S값(false)
}
