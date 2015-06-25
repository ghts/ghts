package internal

import (
	"sync"
)

// 빈 구조체 struct {} 는 Go 언어에서 메모리 소모량이 가작 적은 데이터로 알려져 있음.
// 내용없이 타이밍에 맞게 신호만 보낼 때 유용하다고 알려져 있음.

type S비어있는_구조체 struct {}

var Ch종료_채널_등록 = make(chan (chan S비어있는_구조체), 100)
var Ch등록된_Go루틴_정리 = make(chan S비어있는_구조체)

var 종료_채널_관리_Go루틴_실행_중 = false
var 종료_채널_관리_Go루틴_잠금 = &sync.RWMutex{}


func F새로운_자동_종료_채널_생성_및_등록() chan S비어있는_구조체 {
	if !F종료_채널_관리_Go루틴_실행_중() {
		go F종료_채널_관리_Go루틴()
	}
	
	새로운_종료_채널 := make(chan S비어있는_구조체)
	
	Ch종료_채널_등록 <- 새로운_종료_채널
	
	return 새로운_종료_채널
}

func F종료_채널_관리_Go루틴_실행_중() bool {
	종료_채널_관리_Go루틴_잠금.RLock()
	defer 종료_채널_관리_Go루틴_잠금.RUnlock()
	
	return 종료_채널_관리_Go루틴_실행_중
}

func F종료_채널_관리_Go루틴() {
	if F종료_채널_관리_Go루틴_실행_중() { return }
	
	종료_채널_관리_Go루틴_잠금.Lock()
	
	if 종료_채널_관리_Go루틴_실행_중 {
		종료_채널_관리_Go루틴_잠금.Unlock()
		return 
	}
	
	종료_채널_관리_Go루틴_실행_중 = true
	종료_채널_관리_Go루틴_잠금.Unlock()
	
	defer func() {
		종료_채널_관리_Go루틴_잠금.Lock()
		종료_채널_관리_Go루틴_실행_중 = false
		종료_채널_관리_Go루틴_잠금.Unlock()
	}()
	
	비어있는_구조체 := S비어있는_구조체 {}
	종료_채널_맵 := make(map[chan S비어있는_구조체]S비어있는_구조체)
	
반복문:
	for {
		select {
		case 종료_채널 := <-Ch종료_채널_등록:
			종료_채널_맵[종료_채널] = 비어있는_구조체
		case <-Ch등록된_Go루틴_정리:
			for 종료_채널, _ := range 종료_채널_맵 {
				종료_채널 <- 비어있는_구조체
				close(종료_채널) 
			}
			
			break 반복문
		}	
	}
}