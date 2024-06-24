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
		f에러_호출경로_추가(변환값, F호출경로_모음())
		return 변환값
	case S에러:
		f에러_호출경로_추가(&변환값, F호출경로_모음())
		return &변환값
	case error:
		if len(추가_매개변수) > 0 {
			panic(New에러("New에러() 예상하지 못한 추가 매개변수 : '%v'", len(추가_매개변수)))
		}

		에러 := new(S에러)
		에러.원래_에러 = 변환값
		에러.시점 = time.Now()
		에러.에러_메시지 = strings.TrimSpace(변환값.Error())
		에러.메시지_출력_완료 = false
		f에러_호출경로_추가(에러, F호출경로_모음())

		return 에러
	case string:
		에러 := new(S에러)
		에러.원래_에러 = nil
		에러.시점 = time.Now()
		에러.에러_메시지 = fmt.Sprintf(strings.TrimSpace(변환값), 추가_매개변수...)
		에러.메시지_출력_완료 = false
		f에러_호출경로_추가(에러, F호출경로_모음())

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
	원래_에러     error
	시점        time.Time
	에러_메시지    string
	메시지_출력_완료 bool
	호출_경로_모음  []*S호출경로
}

func (s *S에러) Error() string {
	버퍼 := new(strings.Builder)

	// 에러_메시지는 '메시지_출력_완료'를 통해서 중복 출력 방지
	if !s.메시지_출력_완료 {
		if !strings.HasPrefix(s.에러_메시지, "\n") {
			버퍼.WriteString("\n")
		}

		버퍼.WriteString(s.에러_메시지)

		if !strings.HasSuffix(s.에러_메시지, "\n") {
			버퍼.WriteString("\n")
		}
	}

	// 호출 경로는 매 항목마다 있는 'M출력완료'를 통해서 중복 출력 방지.
	for _, 호출경로 := range s.호출_경로_모음 {
		if 호출경로.M출력완료 {
			continue
		} else {
			버퍼.WriteString(호출경로.M경로_문자열)
			버퍼.WriteString("\n")
		}
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
	s.메시지_출력_완료 = true

	if len(s.호출_경로_모음) > 0 {
		for _, 호출경로 := range s.호출_경로_모음 {
			호출경로.M출력완료 = true
		}
	}
}

type S호출경로 struct {
	M경로_문자열 string
	M출력완료   bool
}

func f에러_호출경로_추가(에러 *S에러, 호출경로_문자열_모음 []string) {
	if 에러.호출_경로_모음 == nil {
		에러.호출_경로_모음 = make([]*S호출경로, 0)
	}

	for _, 호출경로_문자열 := range 호출경로_문자열_모음 {
		호출경로_문자열 = strings.TrimSpace(호출경로_문자열)

		if strings.Contains(에러.에러_메시지, 호출경로_문자열) {
			continue
		}

		이미_존재 := false

		for _, s호출경로 := range 에러.호출_경로_모음 {
			if strings.Contains(s호출경로.M경로_문자열, 호출경로_문자열) {
				이미_존재 = true
				break
			}
		}

		if 이미_존재 {
			continue
		}

		s호출경로 := &S호출경로{
			M경로_문자열: 호출경로_문자열,
			M출력완료:   false,
		}

		에러.호출_경로_모음 = append(에러.호출_경로_모음, s호출경로)
	}
}
