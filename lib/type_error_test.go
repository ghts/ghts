package lib

import (
	"errors"
	"strings"
	"testing"
)

const 에러_메시지 = "에러 메시지 본문"

func TestS에러(t *testing.T) {
	에러 := f함수1()
	F테스트_에러발생(t, 에러)

	if 에러.(*S에러).메시지_출력_완료 {
		F테스트_참임(t, !strings.Contains(에러.Error(), 에러_메시지))
	} else {
		F테스트_참임(t, strings.Contains(에러.Error(), 에러_메시지))
	}

	함수명_모음 := []string{"f함수1", "f함수2", "f함수3", "f함수4", "f함수5"}

	for _, 함수명 := range 함수명_모음 {
		포함 := false
		for _, 호출경로 := range 에러.(*S에러).호출_경로_모음 {
			if strings.Contains(호출경로.M경로_문자열, 함수명) {
				포함 = true
			}
		}

		F테스트_참임(t, 포함)
	}
}

func f함수1() (에러 error) {
	defer S예외처리{M에러: &에러, M출력_숨김: true}.S실행()
	에러 = f함수2()
	return
}

func f함수2() (에러 error) {
	defer S예외처리{M에러: &에러, M출력_숨김: true}.S실행()
	에러 = f함수3()
	return
}

func f함수3() (에러 error) {
	defer S예외처리{M에러: &에러, M출력_숨김: true}.S실행()
	에러 = f함수4()
	return
}

func f함수4() (에러 error) {
	defer S예외처리{M에러: &에러, M출력_숨김: true}.S실행()
	에러 = f함수5()
	return
}

func f함수5() (에러 error) {
	defer S예외처리{M에러: &에러, M출력_숨김: true}.S실행()
	에러 = errors.New(에러_메시지)
	return
}
