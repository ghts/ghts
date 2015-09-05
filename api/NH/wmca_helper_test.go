package NH

import (
	공용 "github.com/ghts/ghts/common"

	"testing"
)

func TestP등락부호_상수(테스트 *testing.T) {

	등락부호_상수_모음 := make([]interface{}, 0)
	등락부호_상수_모음 = append(등락부호_상수_모음, P상한)
	등락부호_상수_모음 = append(등락부호_상수_모음, P상승)
	등락부호_상수_모음 = append(등락부호_상수_모음, P보합)
	등락부호_상수_모음 = append(등락부호_상수_모음, P하한)
	등락부호_상수_모음 = append(등락부호_상수_모음, P하락)

	for _, 등락부호_상수 := range 등락부호_상수_모음 {
		_, ok := 등락부호_상수.(byte)

		공용.F테스트_참임(테스트, ok)
	}
}

func TestP질의_종류_상수(테스트 *testing.T) {
	상수_모음 := make([]interface{}, 0)
	상수_모음 = append(상수_모음, P접속)
	상수_모음 = append(상수_모음, P접속_해제)
	상수_모음 = append(상수_모음, P조회)
	상수_모음 = append(상수_모음, P실시간_서비스_등록)
	상수_모음 = append(상수_모음, P실시간_서비스_해제)
	상수_모음 = append(상수_모음, P실시간_서비스_모두_해제)
	
	for _, 질의_종류_상수 := range 상수_모음 {
		_, ok := 질의_종류_상수.(T질의_종류)

		공용.F테스트_참임(테스트, ok)
	}
}