package lib

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func init() {
	f로그_파일_정리()
	f로그_초기화()
}

func f로그_파일_정리() {
	// 에러가 발생하더라도, 에러 발생하기전까지 읽어들였던 내용에 대해서라도 처리.
	파일_모음, _ := os.ReadDir(".")
	지금 := F지금()

	for _, 파일 := range 파일_모음 {
		파일명 := 파일.Name()

		if !strings.HasPrefix(파일명, "log_") || strings.HasSuffix(파일명, ".txt") {
			continue
		}

		생성_시각_문자열 := 파일명[4 : len(파일명)-4]
		if len(생성_시각_문자열) != 14 {
			fmt.Printf("예상하지 못한 생성 시각 문자열 길이 : %v", len(생성_시각_문자열))
		} else if 생성_시각, 에러 := F2포맷된_시각("20060102150405", 생성_시각_문자열); 에러 != nil {
			fmt.Printf("생성 시각 문자열 해석 오류 : %v", 생성_시각_문자열)
		} else if 생성_시각.Before(지금.AddDate(0, -1, 0)) {
			os.Remove(파일명) // 1달 지난 로그 파일 삭제
		}
	}
}

var 로그_파일 *os.File

func f로그_초기화() {
	if F환경변수("LOG_MODE") == "TEST" {
		F테스트용_로그_초기화()
		return
	}

	var 에러 error

	로그_파일명 := fmt.Sprintf("log_%v.txt", F지금().Format("20060102150405"))

	if 로그_파일, 에러 = os.OpenFile(로그_파일명, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644); 에러 != nil {
		panic(에러)
	} else {
		log.SetOutput(io.MultiWriter(os.Stdout, 로그_파일))
	}
}

func F테스트용_로그_초기화() {
	if 로그_파일 != nil {
		로그_파일명 := 로그_파일.Name()
		로그_파일.Close()
		os.Remove(로그_파일명)
	}
	
	log.SetOutput(os.Stdout)
}
