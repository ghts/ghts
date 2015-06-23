package internal

import (
	zmq "github.com/pebbe/zmq4"

	"testing"

	//"fmt"
)

func TestF외부_프로세스_관리(테스트 *testing.T) {
	//테스트.Parallel()

	const 외부_프로세스_생성_수량 int = 10

	// zmq 소켓 초기화
	통보_REP, 에러 := zmq.NewSocket(zmq.REP)
	F테스트_에러없음(테스트, 에러)

	defer 통보_REP.Close()

	에러 = 통보_REP.Bind(P주소_테스트_결과_회신)
	F테스트_에러없음(테스트, 에러)

	// 외부 프로세스 관리 루틴 시작
	go F외부_프로세스_관리()
	<-Ch외부_프로세스_관리_go루틴_초기화_완료 // 초기화 될 때까지 대기

	for i := 0; i < 외부_프로세스_생성_수량; i++ {
		// 외부 프로세스 생성
		_, 에러 := F파이썬_프로세스_실행("long_running_python_script.py", P주소_테스트_결과_회신)
		F테스트_에러없음(테스트, 에러)

		// 외부 프로세스가 송신하는 메시지를 수신함으로써 외부 프로세스가 정상 동작하는 것 확인.
		통보_REP.RecvMessage(0)
		통보_REP.SendMessage("")
	}

	// 외부 프로세스 관리 루틴이 종료.
	Ch외부_프로세스_관리_go루틴_종료 <- true
	외부_프로세스_정리_횟수 := <-Ch외부_프로세스_관리_go루틴_종료_완료

	// 생성했던 외부 프로세스가 자동으로 정리(종료)되었는 지 확인.
	F테스트_같음(테스트, 외부_프로세스_정리_횟수, 외부_프로세스_생성_수량)
}
