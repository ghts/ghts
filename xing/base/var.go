package xt

import (
	"github.com/ghts/ghts/lib"
	"sync"
)

var (
	V로그인_정보 *S로그인_정보 = nil

	전일, 당일 lib.I안전한_시각

	주소_설정_완료 = lib.New안전한_bool(false)
	주소_설정_잠금 sync.Mutex
)
