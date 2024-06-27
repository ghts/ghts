package lib

import (
	"sync"
	"time"
)

type I송수신 interface {
	S송신(변환_형식 T변환, 값_모음 ...interface{}) error
	G수신() (*S바이트_변환_모음, error)
}

type I소켓 interface {
	I송수신
	S타임아웃(타임아웃 time.Duration) I소켓
	S옵션(옵션_모음 ...interface{})
	Close() error
}

type I소켓with컨텍스트 interface {
	I소켓
	G컨텍스트() (I송수신, error)
}

type I소켓_질의 interface {
	I소켓
	G질의_응답(변환_형식 T변환, 값_모음 ...interface{}) (*S바이트_변환_모음, error)
}

//goland:noinspection GoExportedFuncWithUnexportedType
func New소켓_저장소(수량 int, 생성함수 func() (I소켓_질의, error)) *s소켓_저장소 {
	s := new(s소켓_저장소)
	s.M저장소 = make(chan I소켓_질의, 수량)
	s.M생성함수 = 생성함수

	return s
}

type s소켓_저장소 struct {
	sync.Mutex
	M저장소  chan I소켓_질의
	M생성함수 func() (I소켓_질의, error)
}

var 생성_횟수 int = 1

func (s *s소켓_저장소) G소켓() I소켓_질의 {
	select {
	case <-Ch공통_종료():
		return nil
	case 소켓 := <-s.M저장소:
		//F체크포인트("재활용 성공. 남은 소켓 수량", len(s.M저장소))

		return 소켓
	default:
		s.Lock()
		defer s.Unlock()

		for i := 0; i < 3; i++ {
			if i소켓, 에러 := s.M생성함수(); 에러 == nil {
				F체크포인트(생성_횟수, "번째 소켓 생성 성공. ", i+1, "번째 시도")
				생성_횟수++

				return i소켓
			}

			F대기(P1초)
		}

		F체크포인트(생성_횟수, "번째 소켓 생성 실패")

		return nil
	}
}

func (s *s소켓_저장소) S회수(소켓 I소켓_질의) {
	if 소켓 == nil {
		return
	}

	소켓.S타임아웃(P30초)

	select {
	case s.M저장소 <- 소켓:
		return
	default:
		소켓.Close()
	}
}

func (s *s소켓_저장소) S정리() {
	for i := 0; i < len(s.M저장소); i++ {
		소켓 := <-s.M저장소
		소켓.Close()
	}
}
