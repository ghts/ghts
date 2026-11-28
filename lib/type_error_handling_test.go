package lib

import (
	"bytes"
	"github.com/ghts/ghts/lib/test_helper"
	"log"
	"strings"
	"testing"
)

// TestS예외처리_출력_숨김_실행중_변경 : 함수 실행 중 S예외처리.M출력_숨김의 값을 변경하면,
// 종료 시 S실행()에서 변경된 에러 출력 여부가 반영되는지 확인.
// M출력_숨김은 *bool이므로 struct가 defer 시점에 복사되어도, 가리키는 bool 값은 공유되고,
// S실행()은 함수 종료(실행) 시점에 값을 읽으므로 실행 중 변경이 반영된다.
func TestS예외처리_출력_숨김_실행중_변경(t *testing.T) {
	t.Run("출력에서_숨김으로_변경", func(t *testing.T) {
		버퍼 := F로그_출력_캡처(t)
		에러 := f출력_숨김_실행중_변경_도우미("테스트_에러_출력에서_숨김으로_변경", new(false), new(true))

		th.F테스트_에러발생(t, 에러)
		th.F테스트_거짓임(t, strings.Contains(버퍼.String(), "테스트_에러_출력에서_숨김으로_변경"), "실행 중 숨김으로 변경했으므로, 에러가 출력되면 안 된다.")
	})

	t.Run("숨김에서_출력으로_변경", func(t *testing.T) {
		버퍼 := F로그_출력_캡처(t)
		에러 := f출력_숨김_실행중_변경_도우미("테스트_에러_숨김에서_출력으로_변경", new(true), new(false))

		th.F테스트_에러발생(t, 에러)
		th.F테스트_참임(t, strings.Contains(버퍼.String(), "테스트_에러_숨김에서_출력으로_변경"), "실행 중 출력으로 변경했으므로, 에러가 출력되어야 한다.")
	})

	t.Run("nil이면_출력_기본값", func(t *testing.T) {
		버퍼 := F로그_출력_캡처(t)

		에러 := f출력_숨김_실행중_변경_도우미("테스트_에러_nil_기본값", nil, nil)

		th.F테스트_에러발생(t, 에러)
		th.F테스트_참임(t, strings.Contains(버퍼.String(), "테스트_에러_nil_기본값"), "M출력_숨김이 nil이면 출력(숨김 false)이 기본값이다.")
	})

	t.Run("변경없이_숨김", func(t *testing.T) {
		버퍼 := F로그_출력_캡처(t)
		에러 := f출력_숨김_실행중_변경_도우미("테스트_에러_변경없이_숨김", new(true), nil)

		th.F테스트_에러발생(t, 에러)
		th.F테스트_거짓임(t, strings.Contains(버퍼.String(), "테스트_에러_변경없이_숨김"), "변경 없으면 초기 숨김 상태를 유지해야 한다.")
	})
}

// f출력_숨김_실행중_변경_도우미 : M출력_숨김에 초기값_여부를 넣고, 함수 실행 중(반환 직전) M출력_숨김의 값을
// 실행중_값으로 변경한 뒤 에러를 발생시켜 반환.
// 초기값_여부, 실행중_값이 nil이면 각각 "M출력_숨김 nil(출력 기본값)", "실행 중 변경 없음"을 의미.
func f출력_숨김_실행중_변경_도우미(메시지 string, 초기값_여부 *bool, 실행중_값 *bool) (에러 error) {
	예외_처리 := S예외처리{M에러: &에러, M출력_숨김: 초기값_여부}
	defer 예외_처리.S실행()

	에러 = New에러(메시지)

	if 실행중_값 != nil {
		*예외_처리.M출력_숨김 = *실행중_값 // 함수 실행 중 M출력_숨김 값 변경
	}

	return
}

// F로그_출력_캡처 : 기본 로거 출력(S에러.S출력 → log.Println)을 버퍼로 캡처하고, 테스트 종료 시 이전 출력처 복원.
func F로그_출력_캡처(t testing.TB) *bytes.Buffer {
	t.Helper()

	버퍼 := new(bytes.Buffer)
	이전_출력 := log.Default().Writer()
	log.SetOutput(버퍼)

	t.Cleanup(func() { log.SetOutput(이전_출력) })

	return 버퍼
}
