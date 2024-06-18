package lib

import (
	"sync"
	"time"
)

var (
	ch공통_종료_채널 = make(chan T신호)

	파일경로_맵_잠금 = sync.RWMutex{}
	파일경로_맵    = make(map[string]string)

	실행경로_수정_잠금 = new(sync.Mutex)

	// 이하 테스트 관련 함수 모음
	인터넷_접속_확인_잠금 sync.Mutex
	인터넷_접속됨      = true
	인터넷_접속_확인_완료 = false

	테스트_모드 = New안전한_bool(false)

	화면_출력_잠금 sync.Mutex

	문자열_출력_중복_방지_잠금 = new(sync.Mutex)
	문자열_출력_중복_방지_맵  = make(map[string]S비어있음)

	소켓_테스트용_주소_중복_방지_잠금 = new(sync.Mutex)
	소켓_테스트용_주소_중복_방지_맵  = make(map[string]S비어있음)

	체크포인트_잠금 = new(sync.Mutex)

	P한국 = time.FixedZone("UTC+9", 9*60*60)
)
