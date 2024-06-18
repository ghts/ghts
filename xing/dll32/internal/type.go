package dll32

import (
	"sync"
	"time"
)

func New콜백_대기_항목(식별번호 int, TR코드 string, 값 interface{}) *S콜백_대기_항목 {
	s := new(S콜백_대기_항목)
	s.M식별번호 = 식별번호
	s.M생성_시각 = time.Now()
	s.TR코드 = TR코드
	s.M값 = 값

	return s
}

type S콜백_대기_항목 struct {
	M식별번호  int
	M생성_시각 time.Time
	TR코드   string
	M값     interface{}
}

type S콜백_대기_저장소 struct {
	sync.Mutex
	저장소 map[int]*S콜백_대기_항목
}

func (s *S콜백_대기_저장소) G대기_항목(식별번호 int) *S콜백_대기_항목 {
	s.Lock()
	defer s.Unlock()

	return s.저장소[식별번호]
}

func (s *S콜백_대기_저장소) S추가(식별번호 int, 대기_항목 *S콜백_대기_항목) {
	s.Lock()
	defer s.Unlock()

	s.저장소[식별번호] = 대기_항목
}

func (s *S콜백_대기_저장소) S삭제(식별번호 int) {
	s.Lock()
	defer s.Unlock()

	delete(s.저장소, 식별번호)
}
