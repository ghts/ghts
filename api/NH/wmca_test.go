package NH

import (
	공용 "github.com/ghts/ghts/common"

	"testing"
)

func TestF로드가능(테스트 *testing.T) {
	공용.F테스트_참임(테스트, f로드가능())
}

func TestF접속하기(테스트 *testing.T) {
	f접속하기("ID", "PW", "CertPW")

	공용.F메모("\n" +
		"HWND를 1개를 공유해서 사용할 경우, 1번에 1개의 동작만 수행해야 함.\n" +
		"Go루틴을 사용해서 이러한 것을 구현하도록 할 것.\n" +
		"향후 필요하다고 판단될 경우에는  복수의 HWND를 사용하는 것을 고려해 볼 것.\n" +
		"복수의 HWND를 사용하더라도 DLL 호출은 1번에 1개의 동작만 수행해야 하는 것 아닐까?\n" +
		"API가 thread-safe라고 명시되지 않은 이상 동시 작업은 문제가 발생하기 마련임.\n")
}
