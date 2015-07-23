package price_data

import (
	공용 "github.com/ghts/ghts/shared"
	//공용_정보 "github.com/ghts/ghts/shared_data"
	//zmq "github.com/pebbe/zmq4"
	
	"math"
	"testing"
	"time"
)

func TestF가격정보_Go루틴_수신_및_캐시(테스트 *testing.T) {
	공용.F메모("F가격정보_Go루틴() 새로운 가격정보 입수 기능")
	
	const 값_수량 = 10
	
	f가격정보_Go루틴_테스트_초기화()
	
	종목_모음 := 공용.F샘플_종목_모음()
	통화단위_모음 := 공용.F샘플_통화단위_모음()
	
	
	값_모음 := make(map[string][]interface{})
	r := 공용.F임의값_생성기()
	
	for i:=0 ; i < 값_수량 ; i++ {
		종목코드 := 종목_모음[r.Intn(len(종목_모음))].G코드()
		통화단위 := 통화단위_모음[r.Intn(len(통화단위_모음))]
		금액 := math.Trunc(r.Float64() * math.Pow10(r.Intn(10)) * 100) / 100 // 소숫점 2자리 임의 실수.
		
		값_모음[종목코드] = []interface{}{통화단위, 금액}
		
		// 값 쓰기
		회신 := 공용.New질의(공용.P메시지_SET, 종목코드, 통화단위, 금액).G회신(Ch가격정보, 공용.P타임아웃_Go)
		공용.F테스트_에러없음(테스트, 회신.G에러())
		공용.F테스트_같음(테스트, 회신.G구분(), 공용.P메시지_OK)
		공용.F테스트_같음(테스트, 회신.G길이(), 0)
	}
	
	// 값 읽기
	for 종목코드, 값 := range 값_모음 {
		통화단위 := 값[0].(string)
		금액 := 값[1].(float64)
		
		회신 := 공용.New질의(공용.P메시지_GET, 종목코드, 1).G회신(Ch가격정보, 공용.P타임아웃_Go)
		공용.F테스트_에러없음(테스트, 회신.G에러())
		공용.F테스트_같음(테스트, 회신.G구분(), 공용.P메시지_OK)
		공용.F테스트_같음(테스트, 회신.G길이(), 3)
		공용.F테스트_같음(테스트, 회신.G내용(0), 통화단위)
		공용.F테스트_같음(테스트, 회신.G내용(1), 공용.F실수2문자열(금액))
	
		시점, 에러 := time.Parse(공용.P시간_형식, 회신.G내용(2))
		공용.F테스트_에러없음(테스트, 에러)
		차이 := math.Abs(시점.Sub(time.Now()).Seconds())
		공용.F테스트_참임(테스트, 차이 < 1.0)
	}
	
	// 값 수정
	수정된_값_모음 := make(map[string][]interface{})
	
	for 종목코드, _ := range 값_모음 {
		통화단위 := 통화단위_모음[r.Intn(len(통화단위_모음))]
		금액 := math.Trunc(r.Float64() * math.Pow10(r.Intn(10)) * 100) / 100
		
		수정된_값_모음[종목코드] = []interface{}{통화단위, 금액}
		
		// 값 쓰기
		회신 := 공용.New질의(공용.P메시지_SET, 종목코드, 통화단위, 금액).G회신(Ch가격정보, 공용.P타임아웃_Go)
		공용.F테스트_에러없음(테스트, 회신.G에러())
		공용.F테스트_같음(테스트, 회신.G구분(), 공용.P메시지_OK)
		공용.F테스트_같음(테스트, 회신.G길이(), 0)
	}
	
	// 수정된 값 읽기
	for 종목코드, 값 := range 수정된_값_모음 {
		통화단위 := 값[0].(string)
		금액 := 값[1].(float64)
		
		회신 := 공용.New질의(공용.P메시지_GET, 종목코드, 1).G회신(Ch가격정보, 공용.P타임아웃_Go)
		공용.F테스트_에러없음(테스트, 회신.G에러())
		공용.F테스트_같음(테스트, 회신.G구분(), 공용.P메시지_OK)
		공용.F테스트_같음(테스트, 회신.G길이(), 3)
		공용.F테스트_같음(테스트, 회신.G내용(0), 통화단위)
		공용.F테스트_같음(테스트, 회신.G내용(1), 공용.F실수2문자열(금액))
	
		시점, 에러 := time.Parse(공용.P시간_형식, 회신.G내용(2))
		공용.F테스트_에러없음(테스트, 에러)
		차이 := math.Abs(시점.Sub(time.Now()).Seconds())
		공용.F테스트_참임(테스트, 차이 < 1.0)
	}
}

func TestF가격정보_Go루틴_구독채널_등록(테스트 *testing.T) {
	const 구독채널_수량 = 10
	const 가격정보_수량 = 20
	
	f가격정보_Go루틴_테스트_초기화()
	
	구독채널_모음 := make([]chan 공용.I가격정보, 구독채널_수량)
	
	for i:=0 ; i < 구독채널_수량 ; i++ {
		구독채널_모음[i] = make(chan 공용.I가격정보, 가격정보_수량 + 10)
		Ch가격정보_구독채널_등록 <- 구독채널_모음[i]
	}
	
	종목정보_모음 := 공용.F샘플_종목_모음()
	통화단위_모음 := 공용.F샘플_통화단위_모음()
	r := 공용.F임의값_생성기()
	가격정보_모음 := make([]공용.I가격정보, 가격정보_수량)
	
	for i:=0 ; i < 가격정보_수량 ; i++ {
		종목코드 := 종목정보_모음[r.Intn(len(종목정보_모음))].G코드()
		통화단위 := 통화단위_모음[r.Intn(len(통화단위_모음))]
		금액 := math.Trunc(r.Float64() * math.Pow10(r.Intn(10)) * 100) / 100
		가격정보 := 공용.New가격정보(종목코드, 공용.New통화(통화단위, 금액))
		가격정보_모음[i] = 가격정보
		
		회신 := 공용.New질의(공용.P메시지_SET, 종목코드, 통화단위, 금액).G회신(Ch가격정보, 공용.P타임아웃_Go)
		공용.F테스트_에러없음(테스트, 회신.G에러())
		공용.F테스트_같음(테스트, 회신.G길이(), 0)
	}
	
	현재_시점 := time.Now()
	
	for _, 구독채널 := range 구독채널_모음 {
		for i:=0 ; i < 가격정보_수량 ; i++ {
			가격정보 := <-구독채널
			
			공용.F테스트_같음(테스트, 가격정보.G종목코드(), 가격정보_모음[i].G종목코드())
			공용.F테스트_같음(테스트, 가격정보.G가격().G비교(가격정보_모음[i].G가격()), 공용.P같음)
			공용.F테스트_참임(테스트, 현재_시점.Sub(가격정보.G시점()).Seconds() < 10.0)
		}
	}
}

func f가격정보_Go루틴_테스트_초기화() {
	if !F가격정보_Go루틴_실행_중() {
		ch실행_대기 := make(chan bool)
		go F가격정보_Go루틴(ch실행_대기)
		<-ch실행_대기
	}
	
	_ = 공용.New질의(공용.P메시지_초기화).G회신(ch가격정보_Go루틴_제어, 공용.P타임아웃_Go)
}