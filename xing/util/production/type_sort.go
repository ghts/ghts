package production

import "sort"

type S종목별_데이터_정렬_도우미 struct {
	M저장소   []I종목별_데이터
	Less함수 func(*S종목별_데이터_정렬_도우미, int, int) bool
}

func (s *S종목별_데이터_정렬_도우미) Len() int {
	return len(s.M저장소)
}
func (s *S종목별_데이터_정렬_도우미) Swap(i, j int) {
	s.M저장소[i], s.M저장소[j] = s.M저장소[j], s.M저장소[i]
}
func (s *S종목별_데이터_정렬_도우미) Less(i, j int) bool {
	return s.Less함수(s, i, j)
}

func (s *S종목별_데이터_정렬_도우미) S정렬_함수_설정(Less함수 func(*S종목별_데이터_정렬_도우미, int, int) bool) *S종목별_데이터_정렬_도우미 {
	s.Less함수 = Less함수

	return s
}

func (s *S종목별_데이터_정렬_도우미) S정렬() *S종목별_데이터_정렬_도우미 {
	sort.Sort(s)

	return s
}

func (s *S종목별_데이터_정렬_도우미) S역순_정렬() *S종목별_데이터_정렬_도우미 {
	sort.Sort(sort.Reverse(s))

	return s
}

func (s *S종목별_데이터_정렬_도우미) S상위_N개(수량 int) *S종목별_데이터_정렬_도우미 {
	if 수량 < len(s.M저장소) {
		s.M저장소 = s.M저장소[:수량-1]
	}

	return s
}
