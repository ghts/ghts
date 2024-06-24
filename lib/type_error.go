package lib

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

func New에러(포맷_문자열or에러 interface{}, 추가_매개변수 ...interface{}) error {
	switch 변환값 := 포맷_문자열or에러.(type) {
	case nil:
		return nil
	case *S에러:
		return 변환값
	case S에러:
		return &변환값
	case error:
		if len(추가_매개변수) > 0 {
			panic(New에러("New에러() 예상하지 못한 추가 매개변수 : '%v'", len(추가_매개변수)))
		}

		에러 := new(S에러)
		에러.원래_에러 = 변환값
		에러.시점 = time.Now()
		에러.에러_메시지 = strings.TrimSpace(변환값.Error())
		에러.출력_완료 = false
		에러.호출_경로_모음 = F호출경로_모음()

		return 에러
	case string:
		에러 := new(S에러)
		에러.원래_에러 = nil
		에러.시점 = time.Now()
		에러.에러_메시지 = fmt.Sprintf(strings.TrimSpace(변환값), 추가_매개변수...)
		에러.출력_완료 = false
		에러.호출_경로_모음 = F호출경로_모음()

		return 에러
	default:
		panic(New에러("new에러() 예상하지 못한 자료형. '%T'", 포맷_문자열or에러))
	}
}

func New에러with출력(포맷_문자열or에러 interface{}, 추가_매개변수 ...interface{}) error {
	에러 := New에러(포맷_문자열or에러, 추가_매개변수...)
	F에러_출력(에러)

	return 에러
}

type S에러 struct {
	sync.Mutex
	원래_에러    error
	시점       time.Time
	에러_메시지   string
	출력_완료    bool
	호출_경로_모음 []string
}

func (s *S에러) Error() string {
	if s.출력_완료 {
		return ""
	}

	버퍼 := new(strings.Builder)

	if !strings.HasPrefix(s.에러_메시지, "\n") {
		버퍼.WriteString("\n")
	}

	버퍼.WriteString(s.에러_메시지)

	if !strings.HasSuffix(s.에러_메시지, "\n") {
		버퍼.WriteString("\n")
	}

	for _, 호출경로 := range s.호출_경로_모음 {
		버퍼.WriteString(호출경로)
		버퍼.WriteString("\n")
	}

	return 버퍼.String()
}

func (s *S에러) Is(에러값 error) bool {
	if s.원래_에러 != nil {
		return errors.Is(s.원래_에러, 에러값)
	} else if s.에러_메시지 == 에러값.Error() {
		return true
	}

	return false
}

func (s *S에러) Unwrap() error { return s.원래_에러 }

func (s *S에러) S출력_완료() {
	s.출력_완료 = true
}
