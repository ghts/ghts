package main

import (
	"os"

	lb "github.com/ghts/ghts/lib"
	dll32 "github.com/ghts/ghts/xing/dll32/internal"
)

func main() {
	if 에러 := dll32.F초기화(); 에러 != nil {
		lb.F에러_출력(에러)
		os.Exit(1)
	}

	dll32.F종료_대기()
}
