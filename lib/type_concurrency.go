package lib

func New작업(함수 func(...interface{}), 인수 ...interface{}) *S작업 {
	s := new(S작업)
	s.함수 = 함수
	s.인수 = 인수

	return s
}

type S작업 struct {
	함수 func(...interface{})
	인수 []interface{}
}

func (s *S작업) S실행() { s.함수(s.인수...) }
