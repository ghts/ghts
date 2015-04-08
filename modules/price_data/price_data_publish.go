package price_data_publish

import (
	공용 "github.com/gh-system/ghts/shared"
	zmq "github.com/pebbe/zmq4"
)

func F가격정보_배포_모듈() {
	// 가격정보_입수_소켓
	가격정보_입수_소켓, 에러 := zmq.NewSocket(zmq.REP)
	defer 가격정보_입수_소켓.Close()

	if 에러 != nil {
		공용.F문자열_출력("가격정보_입수_소켓 초기화 중 에러 발생. %s", 에러.Error())
		panic(에러)
	}

	// 가격정보_배포_소켓
	가격정보_배포_소켓, 에러 := zmq.NewSocket(zmq.PUB)
	defer 가격정보_배포_소켓.Close()

	if 에러 != nil {
		공용.F문자열_출력("가격정보_배포_소켓 초기화 중 에러 발생. %s", 에러.Error())
		panic(에러)
	}

	가격정보_입수_소켓.Bind(공용.P가격정보_입수_주소)
	가격정보_배포_소켓.Bind(공용.P가격정보_배포_주소)

	//공용.F문자열_출력("F가격정보_배포_모듈() 초기화 완료.")

	var 메시지 []string
	var 구분 string

	회신_OK := []string{공용.P메시지_구분_OK, ""}

	for {
		// 가격정보 입수
		메시지, 에러 = 가격정보_입수_소켓.RecvMessage(0)

		if 에러 != nil {
			공용.F문자열_출력("가격정보 입수 중 에러 발생.\n %v\n %v\n", 에러.Error(), 공용.F변수_내역_문자열(메시지[0], 메시지[1]))
			가격정보_입수_소켓.SendMessage([]string{공용.P메시지_구분_에러, 에러.Error()})
			//panic(에러)
			continue
		}

		가격정보_입수_소켓.SendMessage(회신_OK)

		// 가격정보 배포
		_, 에러 = 가격정보_배포_소켓.SendMessage(메시지)

		if 에러 != nil {
			공용.F문자열_출력("가격정보 배포 중 에러 발생.\n %v\n %v\n", 에러.Error(), 공용.F변수_내역_문자열(메시지[0], 메시지[1]))
			//panic(에러)
			continue
		}

		// 종료 메시지 수신하면 반복루프 종료
		if 구분 = 메시지[0]; 구분 == 공용.P메시지_구분_종료 {
			//공용.F문자열_출력("배포횟수 : %v", 디버깅용_반복횟수)
			break
		}
	}

	//공용.F문자열_출력("F가격정보_배포_모듈() 종료.")
}