package lib

import (
	"sort"
	"time"
)

type s시간_정렬_도우미 struct{ 시각_모음 []time.Time }

func (s s시간_정렬_도우미) Len() int { return len(s.시각_모음) }
func (s s시간_정렬_도우미) Swap(i, j int) {
	s.시각_모음[i], s.시각_모음[j] = s.시각_모음[j], s.시각_모음[i]
}
func (s s시간_정렬_도우미) Less(i, j int) bool {
	return s.시각_모음[i].Before(s.시각_모음[j])
}

func F시각_정순_정렬(시각_모음 []time.Time) []time.Time {
	sort.Sort(s시간_정렬_도우미{시각_모음})

	return 시각_모음
}

func F시각_역순_정렬(시각_모음 []time.Time) []time.Time {
	sort.Sort(sort.Reverse(s시간_정렬_도우미{시각_모음}))

	return 시각_모음
}
