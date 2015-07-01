package shared

import (
	"bytes"
	"fmt"
	"strings"
)

func init() {
	f메모()
}

func f메모() {
	버퍼 := new(bytes.Buffer)
	버퍼.WriteString("에러가 발생했는 데 원인도 모르겠고 막막하다면?\n")
	버퍼.WriteString("문제를 쪼개고, 단순화 시킨 후, 내부 실행과정을 추적하자.\n")
	버퍼.WriteString("이보다 달리 더 좋은 방법은 없다.\n")
	
	문자열 := 버퍼.String()
	
	if !strings.HasSuffix(문자열, "\n") {
		문자열 += "\n"
	}
	 
	fmt.Println(버퍼.String())
	fmt.Println("")
}