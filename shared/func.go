package shared

import (
	"os"
	"os/exec"
)

func F파이썬_프로세스_실행(파일명 string, 실행옵션 ...string) {
	실행옵션 = append([]string{파일명}, 실행옵션...)
	F외부_프로세스_실행(p파이썬_경로, 실행옵션...)
}

func F외부_프로세스_실행(프로그램 string, 실행옵션 ...string) {
	외부_프로세스 := exec.Command(프로그램, 실행옵션...)
	외부_프로세스.Stdin = os.Stdin
	외부_프로세스.Stdout = os.Stdout
	외부_프로세스.Stderr = os.Stderr

	에러 := 외부_프로세스.Start()
	if 에러 != nil {
		panic(에러.Error())
	}
}

func F실행화일_검색(파일명 string) string {
	파일경로, 에러 := exec.LookPath(파일명)

	if 에러 != nil {
		F문자열_출력("'%v' : 파일을 찾을 수 없습니다.", 파일명)
		return ""
	}

	return 파일경로
}

func F문자열_복사(문자열 string) string {
	return (문자열 + " ")[:len(문자열)]
}