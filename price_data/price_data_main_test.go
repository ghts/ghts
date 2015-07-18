package price_data

import (
	공용 "github.com/ghts/ghts/shared"
	공용_정보 "github.com/ghts/ghts/shared_data"
	zmq "github.com/pebbe/zmq4"
	
	"testing"
	"time"
)

func TestF가격정보_Go루틴(테스트 *testing.T) {
	공용.F메모("TODO : f새로운_가격정보_회신()")
	
	// a. 가격입수
   	// b. 가격배포 - 코딩 완료. TODO : 테스트 작성 및 디버깅
	// c. 가격정보 캐시 - 코딩 완료. TODO : 테스트 작성 및 디버깅
	// d. 가격정보 질의 응답 - 코딩 완료. TODO : 테스트 작성 및 디버깅
	
	
	f가격정보_Go루틴_종료_후_재시작()
	
	// 캐시 기능 쓰기, 읽기 테스트
	Ch가격정보 <- w
		
}

func f가격정보_Go루틴_종료_후_재시작() {
	if F가격정보_Go루틴_실행_중() {
		ch종료_가격정보_Go루틴 <- S비어있는_구조체{}
	}
		
	for F가격정보_Go루틴_실행_중() {
		time.Sleep(100 * time.Millisecond)
	}
}