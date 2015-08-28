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


