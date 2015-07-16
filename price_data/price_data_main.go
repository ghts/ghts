package price_data

import (
	공용 "github.com/ghts/ghts/shared"
	
	"math"
	"strconv"
	"time"
)

var Ch가격정보 = make(chan 공용.I질의, 1000)
var Ch가격정보_구독_채널 = make(chan (chan 공용.I가격정보), 100)
var Ch가격정보_구독_소켓 = make(chan string, 100)
var ch종료_가격정보_Go루틴 = make(chan 공용.S비어있는_구조체, 1)

var 가격정보_Go루틴_실행_중 = 공용.New안전한_bool(false)

func F가격정보_Go루틴_실행_중() bool {
	return 가격정보_Go루틴_실행_중.G값()
}

func F가격정보_Go루틴(go루틴_생성_결과 chan bool) {
	에러 := 가격정보_Go루틴_실행_중.S값(true)
	if 에러 != nil {
		go루틴_생성_결과 <- false; return
	}
	
	가격정보_맵 := make(map[string]공용.I가격정보)
	구독채널_맵 := make(map[chan 공용.I가격정보]공용.S비어있는_구조체)
	구독소켓_맵 := make(map[string]공용.S비어있는_구조체)
	공통_종료_채널 := 공용.F공통_종료_채널()
	
	// GC 압력을 줄이기 위한 재활용 변수
	질의 := 공용.New질의(공용.P메시지_GET)
	
	// 초기화 완료
	go루틴_생성_결과 <- true
	
	for {
		select {
		case 질의 = <-Ch가격정보:
			switch 질의.G구분() {
			case 공용.P메시지_GET:
				에러 = f캐시된_가격정보_회신(질의, 가격정보_맵)
				if 에러 == nil {
					break
				}
				
				가격정보, 에러 := f새로운_가격정보_회신(질의, 가격정보_맵)
				if 에러 != nil {
					공용.F문자열_출력(질의.String())
					공용.F에러_출력(에러.Error())
					break
				}
				
				에러 = f가격정보_배포(가격정보, 구독채널_맵, 구독소켓_맵)
				if 에러 != nil {
					공용.F문자열_출력(질의.String())
					공용.F에러_출력(에러.Error())
					break
				} 
			case 공용.P메시지_SET:
				가격정보, 에러 := f가격정보_캐시_저장(질의, 가격정보_맵)
				if 에러 != nil {
					공용.F문자열_출력(질의.String())
					공용.F에러_출력(에러.Error())
				}
				
				에러 = f가격정보_배포(가격정보, 구독채널_맵, 구독소켓_맵)
				if 에러 != nil {
					공용.F문자열_출력(질의.String())
					공용.F에러_출력(에러.Error())
					break
				}
			default:
				질의.S회신(공용.F에러_생성("예상치 못한 메시지 구분.\n%v", 질의))
			}
		case ch구독 := <-Ch가격정보_구독_채널:
			구독채널_맵[ch구독] = 공용.S비어있는_구조체{}
		case 소켓_주소 := <-Ch가격정보_구독_소켓:
			구독소켓_맵[소켓_주소] = 공용.S비어있는_구조체{}
		case <-공통_종료_채널:
			ch종료_가격정보_Go루틴 <- 공용.S비어있는_구조체{}
		case <-ch종료_가격정보_Go루틴:
			가격정보_Go루틴_실행_중.S값(false)
			return	
		}
	}
}

func f캐시된_가격정보_회신(질의 공용.I질의, 가격정보_맵 map[string]공용.I가격정보) error {
	// 0 : 종목코드, 1 : 유효기간 (초 단위), 2 : 소숫점 이하 자릿수
	if 에러 := 질의.G검사(공용.P메시지_GET, 3); 에러 != nil {
		질의.S회신(에러); return 에러
	}
	
	여기까지
	
	초_단위_유효기간, 에러 := strconv.ParseFloat(질의.G내용(1))
	if 에러 != nil {
		질의.S회신(에러); return 에러
	}
	
	if 초_단위_유효기간 < 0 {
		return F에러_생성("유효기간이 음수임. 가격정보 캐시 확인 중단.") 
	}
	
	종목코드 := 질의.G내용(0)
	가격정보, 존재함 := 가격정보_맵[종목코드]
	if !존재함 {
		에러 = 공용.F에러_생성("가격정보 캐시 데이터가 없음.\n%v", 종목코드, 질의)
		질의.S회신(에러); return 에러
	}
		
	차이 := math.Abs(time.Now().Sub(가격정보.G시점()).Seconds())
		
	if 차이 > 초_단위_유효기간 {
		에러 = 공용.F에러_생성("유효기간이 지났음. 차이 %v초.", 차이)
		질의.S회신(에러); return 에러
	}
		
	소숫점_이하_자릿수, 에러 := strconv.Atoi(질의.내용(2))
	if 에러 != nil {
		질의.S회신(에러); return 에러
	}
	
	질의.S회신(nil, 
		가격정보.G가격().G단위(), 
		가격정보.G가격().G실수_문자열(소숫점_이하_자릿수), 
		가격정보.G시점().Format(공용.P시간_형식))
	
	return nil
}

func f새로운_가격정보_회신(
	질의 공용.I질의, 가격정보_맵 map[string]공용.I가격정보) (공용.I가격정보, error) {
	F메모("f새로운_가격정보_회신(). 가격정보 입수 루틴 작성 후 진행 가능할 듯 함.")
	
	return nil, 공용.F에러_생성("아직 구현하지 못함.")
}

func f가격정보_캐시_저장(질의 공용.I질의, 가격정보_맵 map[string]공용.I가격정보) (공용.I가격정보, error) {
	// 0 : 종목코드, 1 : 통화단위, 2 : 금액
	if 에러 := 질의.G검사(공용.P메시지_SET, 3); 에러 != nil {
		질의.S에러(에러)
		return nil, 에러
	}
		
	종목코드 := 질의.G내용(0)
	통화 := 공용.New통화(질의.G내용(1), 질의.G내용(2))
	
	if 통화 == nil {
		에러 = 공용.F에러_생성("가격정보 SET 통화 생성 에러.\n%v", 질의)
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
	구독소켓_맵 map[string]공용.S비어있는_구조체) error {
	에러 := f가격정보_배포_Go채널(가격정보, 구독채널_맵)
	if 에러 != nil {
		return 에러
	}
	
	에러 = f가격정보_배포_zmq소켓(가격정보, 구독소켓_맵)
	if 에러 != nil {
		return 에러
	}
	
	return nil
}

func f가격정보_배포_Go채널(가격정보 공용.I가격정보, 
		구독채널_맵 map[chan 공용.I가격정보]공용.S비어있는_구조체) error {
	공용.F메모("f가격정보_배포_Go채널()")
	
	return 공용.F에러_생성("아직 구현하지 못함.")
}
		
func f가격정보_배포_zmq소켓(가격정보 공용.I가격정보, 
		구독소켓_맵 map[string]공용.S비어있는_구조체) error {
	공용.F메모("f가격정보_배포_zmq소켓()")
	
	return 공용.F에러_생성("아직 구현하지 못함.")
}