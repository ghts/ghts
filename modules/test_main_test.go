package modules

import (
	공용 "github.com/ghts/ghts/shared"
	
	"os"
	"testing"
)


func TestMain(m *testing.M) {
	go F외부_프로세스_정리()
	
	샘플_종목_모음 := 공용.F샘플_종목_모음()
		
	for i:=0 ; i < len(샘플_종목_모음) ; i++ {
		종목정보_맵[샘플_종목_모음[i].G코드()] = 샘플_종목_모음[i]
	}
	
	공용.F테스트_모드_시작()
	defer 공용.F테스트_모드_종료()
	
	if 공용.F단일_스레드_모드임() {
		공용.F멀티_스레드_모드()
		defer 공용.F단일_스레드_모드()
	}
	
	os.Exit(m.Run())
}

