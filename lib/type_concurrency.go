package lib

func New작업(함수 func(...any), 인수 ...any) *S작업 {
	s := new(S작업)
	s.함수 = 함수
	s.인수 = 인수

	return s
}

type S작업 struct {
	함수 func(...any)
	인수 []any
}

func (s *S작업) S실행() { s.함수(s.인수...) }
