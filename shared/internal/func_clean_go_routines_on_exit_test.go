package internal

import (
	"testing"
)

func TestF등록된_모든_Go루틴_종료(테스트 *testing.T) {
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
	
	F등록된_모든_Go루틴_종료()
	
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