package price_data_publish

import (
	공용 "github.com/gh-system/ghts/src/go/shared"
	zmq "github.com/pebbe/zmq4"
	"strconv"
	"testing"
	"time"
)

func f테스트용_가격정보_입수_모듈() {
	공용.F문자열_출력("f테스트용_가격정보_입수_모듈() 시작.")

	defer 공용.WaitGroup.Done()

	// 가격정보_송신_소켓
	가격정보_송신_소켓, 에러 := zmq.NewSocket(zmq.REQ)
	defer 가격정보_송신_소켓.Close()

	if 에러 != nil {
		공용.F문자열_출력("가격정보_송신_소켓 초기화 중 에러 발생. %s", 에러.Error())
		panic(에러)
	}

	가격정보_송신_소켓.Connect(공용.P가격정보_입수_주소)

	// 다른 모듈 초기화 할 동안 잠시 대기
	time.Sleep(2 * time.Second)

	공용.F문자열_출력("f테스트용_가격정보_입수_모듈() 초기화 완료.")

	var 메시지 []string
	var 구분 string
	var 에러_메시지 string

	for 반복횟수 := 0; 반복횟수 < 10000; 반복횟수++ {
		//공용.F문자열_출력("f테스트용_가격정보_입수_모듈() : 입수모듈 %v", 반복횟수 + 1)

		가격 := (반복횟수 + 1) * 10

		// 가격정보 송신
		메시지 = []string{공용.P메시지_구분_일반, strconv.Itoa(가격)}

		_, 에러 = 가격정보_송신_소켓.SendMessage(메시지)

		if 에러 != nil {
			공용.F문자열_출력("가격정보 송신 중 에러 발생.\n %v\n %v\n", 에러.Error(), 공용.F변수_내역_문자열(메시지[0], 메시지[1]))
			가격정보_송신_소켓.SendMessage([]string{공용.P회신_메시지_구분_에러, 에러.Error()})
			//panic(에러)
			continue
		}

		//공용.F문자열_출력("f테스트용_가격정보_입수_모듈() : SendMessage %v", 반복횟수 + 1)

		메시지, 에러 = 가격정보_송신_소켓.RecvMessage(0)

		//공용.F문자열_출력("f테스트용_가격정보_입수_모듈() : RecvMessage %v", 반복횟수 + 1)

		if 에러 != nil {
			공용.F문자열_출력("가격정보 송신 후 회신 수신 중 에러 발생.\n %v\n %v\n", 에러.Error(), 공용.F변수_내역_문자열(메시지[0], 메시지[1]))
			//panic(에러)
			continue
		}

		구분 = 메시지[0]
		에러_메시지 = 메시지[1]

		if 구분 == 공용.P회신_메시지_구분_에러 {
			공용.F문자열_출력("가격정보 송신 후 에러 메시지 수신.\n %v\n", 에러_메시지)
			//panic(에러_메시지)
			continue
		}
	}

	메시지 = []string{공용.P메시지_구분_종료, ""}
	가격정보_송신_소켓.SendMessage(메시지)
	가격정보_송신_소켓.RecvMessage(0)

	//공용.F문자열_출력("f테스트용_가격정보_입수_모듈() 종료 메시지 송신.")

	공용.F문자열_출력("f테스트용_가격정보_입수_모듈() 종료.")
}

func f테스트용_가격정보_구독_모듈() {
	공용.F문자열_출력("f테스트용_가격정보_구독_모듈() 시작.")

	defer 공용.WaitGroup.Done()

	가격정보_구독_소켓, 에러 := zmq.NewSocket(zmq.SUB)
	defer 가격정보_구독_소켓.Close()

	if 에러 != nil {
		공용.F문자열_출력("가격정보_구독_소켓 초기화 중 에러 발생. %v", 에러.Error())
		panic(에러)
	}

	가격정보_구독_소켓.Connect(공용.P가격정보_배포_주소)
	가격정보_구독_소켓.SetSubscribe("")

	공용.F문자열_출력("f테스트용_가격정보_구독_모듈() 초기화 완료.")

	var 메시지 []string
	var 구분, 데이터 string
	var 반복횟수 int = 1

	for {
		메시지, 에러 = 가격정보_구독_소켓.RecvMessage(0)

		if 에러 != nil {
			공용.F문자열_출력("가격정보_구독_소켓 메시지 수신 중 에러 발생. %v", 에러.Error())
			continue
		}

		구분 = 메시지[0]
		데이터 = 메시지[1]

		if 구분 == 공용.P메시지_구분_종료 {
			break
		}

		if 반복횟수%500 == 0 {
			데이터 = 데이터
			//공용.F문자열_출력("반복횟수 %v : 수신 데이터 %v", 반복횟수, 데이터)
		}

		//공용.F문자열_출력("f테스트용_가격정보_구독_모듈() : 구독모듈 %v, 데이터 %v", 반복횟수, 데이터)

		반복횟수++
	}

	공용.F문자열_출력("f테스트용_가격정보_구독_모듈() %v회 반복완료 후 종료.", 반복횟수-1) // 마지막 1번은 종료 신호이므로 제외해야 함.
}

func TestXYZ(t *testing.T) {
	공용.F멀티_스레드_모드()
	defer 공용.F단일_스레드_모드()

	공용.WaitGroup.Add(3)

	go F가격정보_배포_모듈()
	go f테스트용_가격정보_입수_모듈()
	go f테스트용_가격정보_구독_모듈()

	공용.WaitGroup.Wait()
}
