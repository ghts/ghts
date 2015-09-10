package nh

import (
	공용 "github.com/ghts/ghts/common"
	내부 "github.com/ghts/ghts/api/nh/internal"
)

func F접속(아이디, 암호, 공인인증_암호 string) {
	질의 := 공용.New질의_가변형(내부.P30초, 공용.P메시지_GET, 아이디, 암호, 공인인증_암호)
	회신 := 질의.S질의(내부.Ch접속).G회신()
	
	// TODO
	회신.G에러()
} 

