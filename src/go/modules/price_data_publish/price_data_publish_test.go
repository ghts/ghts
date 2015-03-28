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
    
    // 가격정보_송신_소켓
    가격정보_송신_소켓, 에러 := zmq.NewSocket(zmq.REQ)
    defer 가격정보_송신_소켓.Close()
    
    if 에러 != nil {
        공용.F문자열_출력("가격정보_송신_소켓 초기화 중 예상하지 못한 에러 발생. %s", 에러.Error())
	    panic(에러)
    }
	
	가격정보_송신_소켓.Connect(공용.P가격정보_입수_주소)
	
	공용.F문자열_출력("f테스트용_가격정보_입수_모듈() 초기화 완료.")
	
	// 다른 모듈 초기화 할 동안 잠시 대기
	time.Sleep(2 * time.Second)
	
	var total int
	var 메시지 []string
	var 구분 string
	
	for 반복횟수 := 0; 반복횟수 < 10000; 반복횟수++ {
	    가격 := 반복횟수 * 10
	    
	    // 가격정보 송신
	    메시지 = []string{공용.P송신_메시지_구분_일반, strconv.Itoa(가격)}
	    
	    total, 에러 = 가격정보_송신_소켓.SendMessage(메시지)
	    
	    if 에러 != nil {
	        공용.F문자열_출력("가격정보 송신 중 에러 발생.\n %v\n %v\n", 에러.Error(), 공용.F변수_내역_문자열(메시지[0], 메시지[1]))
	        가격정보_송신_소켓.SendMessage([]string{공용.P회신_메시지_구분_에러, 에러.Error()})
	        //panic(에러)
	        continue
	    }
	    
	    공용.F문자열_출력("total : %v", total)	// 이게 뭐지?
	    
	    메시지, 에러 = 가격정보_송신_소켓.RecvMessage(0)
	    
	    if 에러 != nil {
	        공용.F문자열_출력("가격정보 송신 후 회신 수신 중 에러 발생.\n %v\n %v\n", 에러.Error(), 공용.F변수_내역_문자열(메시지[0], 메시지[1]))
	        //panic(에러)
	        continue
	    }
	    
	    if 구분 = 메시지[0]; 구분 == 공용.P회신_메시지_구분_에러 {
	        에러_메시지 := 메시지[1]
	        공용.F문자열_출력("가격정보 송신 후 에러 메시지 수신.\n %v\n", 에러_메시지)
	        //panic(에러_메시지)
	        continue
	    }
	}
	
	공용.F문자열_출력("f테스트용_가격정보_입수_모듈() 종료.")
}

func f테스트용_가격정보_구독_모듈() {
    // TODO
}

func TestXYZ(t *testing.T) {
	공용.F메모("가격정보 배포 PUB-SUB 구현 및 테스트.")
}

