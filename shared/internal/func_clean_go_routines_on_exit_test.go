package internal

import (
	"testing"
)

func TestF등록된_Go루틴_종료(테스트 *testing.T) {
	// 공통 종료 채널을 이용하는 다른 테스트에 영향을 주지 않기 위해서
	// 새로운 채널을 이용해서 테스트를 진행함.
	ch공통_종료_채널_원본 := ch공통_종료_채널
	ch공통_종료_채널 = make(chan S비어있는_구조체)
	
	// 테스트 종료할 때, 공통 종료 채널을 원래대로 되돌려 놓음.
	defer func() {
		ch공통_종료_채널 = ch공통_종료_채널_원본 
	}()
	
	
	ch입력_모음 := make([](chan int), 10)
	ch출력 := make(chan int)
	ch공통_종료 := F공통_종료_채널()
	
	// Go루틴 10개 생성
	for i, _ := range ch입력_모음 {
		ch입력 := make(chan int)
		ch입력_모음[i] = ch입력
		
		go f테스트용_Go루틴(ch입력, ch출력, ch공통_종료)
	}
	
	// 모든 Go루틴 존재 확인
	for _, ch입력 := range ch입력_모음 {
		ch입력 <- 1
		F테스트_같음(테스트, <-ch출력, 1)
	}
	
	F등록된_Go루틴_종료()
	
	for range ch입력_모음 {
		F테스트_같음(테스트, <-ch출력, 999)
	}
}

func f테스트용_Go루틴(ch입력, ch출력 chan int, ch공통_종료 chan S비어있는_구조체) {
	for {
		select {
		case _, ok := <-ch입력:
			if !ok { ch출력 <- 10 }
			
			ch출력 <- 1
		case <-ch공통_종료:
			ch출력 <- 999
			return
		}
	}
}