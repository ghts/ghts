package lib

import "testing"

func TestS작업(t *testing.T) {
	ch응답 := make(chan bool, 1)
	New작업(f작업_테스트용_함수, ch응답).S실행()

	F테스트_참임(t, <-ch응답)
}

func f작업_테스트용_함수(인수 ...interface{}) {
	인수[0].(chan bool) <- true
}
