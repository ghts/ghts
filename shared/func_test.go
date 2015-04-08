package shared

import (
	zmq "github.com/pebbe/zmq4"
	"testing"
)

func TestF파이썬_프로세스_실행(테스트 *testing.T) {
	테스트_결과_회신_소켓, 에러 := zmq.NewSocket(zmq.REP)
	defer 테스트_결과_회신_소켓.Close()

	if 에러 != nil {
		F문자열_출력(에러.Error())
		테스트.Fail()
	}

	테스트_결과_회신_소켓.Bind(P테스트_결과_회신_주소)
	
	F파이썬_프로세스_실행("func_test.py", "exec_python_process", P테스트_결과_회신_주소)

	메시지, _ := 테스트_결과_회신_소켓.RecvMessage(0)
	구분 := 메시지[0]
	데이터 := 메시지[1]

	F테스트_같음(테스트, 구분, P메시지_구분_OK)
	F테스트_같음(테스트, 데이터, "")

	테스트_결과_회신_소켓.SendMessage([]string{P메시지_구분_OK, ""})
}
