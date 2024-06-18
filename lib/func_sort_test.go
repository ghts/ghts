package lib

import (
	"testing"
	"time"
)

func TestF정렬_시각(t *testing.T) {
	t.Parallel()

	어제 := F2일자(time.Now()).AddDate(0, 0, -1)
	오늘 := F2일자(time.Now())
	내일 := F2일자(time.Now()).AddDate(0, 0, 1)

	시각_모음 := []time.Time{오늘, 어제, 내일}
	시각_모음 = F시각_정순_정렬(시각_모음)

	F테스트_같음(t, 시각_모음[0], 어제)
	F테스트_같음(t, 시각_모음[1], 오늘)
	F테스트_같음(t, 시각_모음[2], 내일)
}
